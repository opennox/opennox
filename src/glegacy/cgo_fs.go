package legacy

import (
	"errors"
	"io"
	"os"
	"sync"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/datapath"
	"github.com/noxworld-dev/opennox-lib/ifs"

	"github.com/noxworld-dev/opennox/v1/internal/binfile"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
)

var files struct {
	sync.RWMutex
	byHandle map[unsafe.Pointer]*binfile.File
}

type FILE = FILE

// nox_fs_root
func nox_fs_root() *char {
	return internCStr(datapath.Data())
}

// nox_fs_normalize
func nox_fs_normalize(path *char) *char {
	out := ifs.Normalize(GoString(path))
	return CString(out)
}

// nox_fs_remove
func nox_fs_remove(path *char) bool {
	return ifs.Remove(GoString(path)) == nil
}

// nox_fs_mkdir
func nox_fs_mkdir(path *char) bool {
	return ifs.Mkdir(GoString(path)) == nil
}

// nox_fs_set_workdir
func nox_fs_set_workdir(path *char) bool {
	return ifs.Chdir(GoString(path)) == nil
}

// nox_fs_copy
func nox_fs_copy(src, dst *char) bool {
	return ifs.Copy(GoString(src), GoString(dst)) == nil
}

// nox_fs_move
func nox_fs_move(src, dst *char) bool {
	return ifs.Rename(GoString(src), GoString(dst)) == nil
}

func convWhence(mode int) int {
	var whence int
	switch int(mode) {
	case SEEK_SET:
		whence = io.SeekStart
	case SEEK_CUR:
		whence = io.SeekCurrent
	case SEEK_END:
		whence = io.SeekEnd
	default:
		panic("unsupported seek mode")
	}
	return whence
}

// nox_fs_fseek
func nox_fs_fseek(f *FILE, off long, mode int) int {
	fp := fileByHandle(f)
	_, err := fp.Seek(int64(off), convWhence(mode))
	if err != nil {
		return -1
	}
	return 0
}

// nox_fs_ftell
func nox_fs_ftell(f *FILE) long {
	fp := fileByHandle(f)
	off, err := fp.Seek(0, io.SeekCurrent)
	if err != nil {
		e := int64(-1)
		return long(e)
	}
	return long(off)
}

// nox_fs_fsize
func nox_fs_fsize(f *FILE) long {
	fp := fileByHandle(f)
	size, err := fp.Size()
	if err != nil {
		e := int64(-1)
		return long(e)
	}
	return long(size)
}

// nox_fs_fread
func nox_fs_fread(f *FILE, dst unsafe.Pointer, sz int) int {
	fp := fileByHandle(f)
	n, _ := fp.Read(unsafe.Slice((*byte)(dst), sz))
	return n
}

// nox_fs_fwrite
func nox_fs_fwrite(f *FILE, dst unsafe.Pointer, sz int) int {
	fp := fileByHandle(f)
	n, _ := fp.Write(unsafe.Slice((*byte)(dst), sz))
	return n
}

// nox_fs_fgets
func nox_fs_fgets(f *FILE, dst *char, sz int) bool {
	fp := fileByHandle(f)
	out, err := fp.ReadString()
	if err != nil && !errors.Is(err, io.EOF) {
		return false
	}
	StrCopy(dst, sz, string(out))
	return bool(!errors.Is(err, io.EOF))
}

// nox_fs_fputs
func nox_fs_fputs(f *FILE, str *char) int {
	fp := fileByHandle(f)
	n, err := fp.WriteString(GoString(str))
	if err != nil {
		return -1
	}
	return n
}

// nox_fs_feof
func nox_fs_feof(f *FILE) bool {
	fp := fileByHandle(f)
	return fp.Err == io.EOF
}

func fileByHandle(f *FILE) *binfile.File {
	h := unsafe.Pointer(f)
	handles.AssertValidPtr(h)
	files.RLock()
	fp := files.byHandle[h]
	files.RUnlock()
	return fp
}

// nox_fs_close
func nox_fs_close(f *FILE) {
	if f == nil {
		return
	}
	h := unsafe.Pointer(f)
	handles.AssertValidPtr(h)
	files.Lock()
	defer files.Unlock()
	fp := files.byHandle[h]
	if fp != nil {
		_ = fp.Close()
		delete(files.byHandle, h)
	}
}

func Nox_fs_close(f *FILE) {
	nox_fs_close(f)
}

func NewFileHandle(f *binfile.File) *FILE {
	if f.Handle != nil {
		return (*FILE)(f.Handle)
	}
	f.Handle = handles.NewPtr()
	files.Lock()
	defer files.Unlock()
	if files.byHandle == nil {
		files.byHandle = make(map[unsafe.Pointer]*binfile.File)
	}
	files.byHandle[f.Handle] = f
	return (*FILE)(f.Handle)
}

// nox_fs_access
func nox_fs_access(path *char, mode int) int {
	_, err := ifs.Stat(GoString(path))
	if os.IsNotExist(err) {
		return -1
	} else if err != nil {
		return -2
	}
	return 0
}

// nox_fs_open
func nox_fs_open(path *char) *FILE {
	f, err := ifs.Open(GoString(path))
	if err != nil {
		return nil
	}
	return NewFileHandle(binfile.NewFile(f))
}

// nox_fs_open_text
func nox_fs_open_text(path *char) *FILE {
	f, err := ifs.Open(GoString(path))
	if err != nil {
		return nil
	}
	return NewFileHandle(binfile.NewTextFile(f))
}

// nox_fs_create
func nox_fs_create(path *char) *FILE {
	f, err := ifs.Create(GoString(path))
	if err != nil {
		return nil
	}
	return NewFileHandle(binfile.NewFile(f))
}

// nox_fs_create_text
func nox_fs_create_text(path *char) *FILE {
	f, err := ifs.Create(GoString(path))
	if err != nil {
		return nil
	}
	return NewFileHandle(binfile.NewTextFile(f))
}

// nox_fs_open_rw
func nox_fs_open_rw(path *char) *FILE {
	f, err := ifs.OpenFile(GoString(path), os.O_RDWR)
	if err != nil {
		return nil
	}
	return NewFileHandle(binfile.NewFile(f))
}