#include "memfile.h"
#include "proto.h"
#include <string.h>

//----- (0040ABF0) --------------------------------------------------------
nox_memfile* nox_memfile_load(const char* path, int a2) {
	nox_memfile* nf = (nox_memfile*)malloc(sizeof(nox_memfile));
	if (!nf) {
		return 0;
	}
	FILE* f = nox_xxx_openFileBin_408CC0(path, 0);
	if (!f) {
		nox_memfile_free(nf);
		return 0;
	}
	if (!nox_xxx_cryptOpen_408D40(f, a2)) {
		nox_memfile_free(nf);
		nox_xxx_fileBinClose_408D90(f);
		return 0;
	}
	nox_xxx_fseek_409050(f, 0, SEEK_END);
	nf->size = nox_xxx_fileBinGetInPtr_409390();
	nox_xxx_fseek_409050(f, 0, SEEK_SET);
	nf->data = (char*)malloc(nf->size);
	if (nf->data == 0) {
		nox_memfile_free(nf);
		nox_xxx_fileBinClose_408D90(f);
		return 0;
	}
	int csize = nf->size;
	int x1 = nox_xxx_fread_408E40_fread(nf->data, 1, nf->size, f);
	if (x1 != csize) {
		nox_memfile_free(nf);
		nox_xxx_fileBinClose_408D90(f);
		return 0;
	}
	nf->cur = nf->data;
	nf->end = nf->data + csize;
	nox_xxx_fileBinClose_408D90(f);
	return nf;
}

//----- (0040ACA0) --------------------------------------------------------
void nox_memfile_free(nox_memfile* f) {
	if (f->data) {
		free(f->data);
	}
	f->data = NULL;
	f->cur = NULL;
	f->end = NULL;
	f->size = 0;
	free(f);
}

int8_t nox_memfile_read_i8(nox_memfile* f) {
	if (!f->data)
		return 0;
	int8_t v = *(int8_t*)f->cur;
	f->cur++;
	return v;
}

uint8_t nox_memfile_read_u8(nox_memfile* f) {
	if (!f->data)
		return 0;
	uint8_t v = *(uint8_t*)f->cur;
	f->cur++;
	return v;
}

int32_t nox_memfile_read_i32(nox_memfile* f) {
	if (!f->data)
		return 0;
	int32_t v = *(int32_t*)f->cur;
	f->cur += 4;
	return v;
}

uint32_t nox_memfile_read_u32(nox_memfile* f) {
	if (!f->data)
		return 0;
	uint32_t v = *(uint32_t*)f->cur;
	f->cur += 4;
	return v;
}

void nox_memfile_skip(nox_memfile* f, int n) {
	if (!f->data)
		return;
	f->cur += n;
}

//----- (0040ACC0) --------------------------------------------------------
unsigned int nox_memfile_read(void* dst, const unsigned int sz, const int cnt, nox_memfile* f) {
	unsigned int n = cnt * sz;
	if (f->cur + n > f->end)
		n = f->end - f->cur;
	memcpy(dst, f->cur, n);
	f->cur += n;
	return n / sz;
}

//----- (0040AD10) --------------------------------------------------------
unsigned int __cdecl nox_memfile_seek_40AD10(nox_memfile* memfile, const int offset, const int fromStartOrEnd) {
	if (fromStartOrEnd == 0) {
		memfile->cur = offset + memfile->data;
	} else if (fromStartOrEnd == 2) {
		memfile->cur = offset + memfile->end;
	}

	if (memfile->cur < memfile->data) {
		memfile->cur = memfile->data;
		return 0;
	}

	if (memfile->cur > memfile->end) {
		memfile->cur = memfile->end;
	}

	return memfile->cur - memfile->data;
}

//----- (0040AD60) --------------------------------------------------------
unsigned int nox_xxx_freadMB_40AD60(char* dest, int sz, int cnt, nox_memfile* f) {
    const size_t cur_offset = f->cur - f->data;
    const uint8_t read_past_8 = cur_offset & 0x7;

    char buf[8];
    if (read_past_8) {
        nox_memfile_read(&buf, 8 - read_past_8, 1, f);
    }

    unsigned int result = nox_memfile_read(&buf, 8, 1, f);
    if (result != 1) {
        return result;
    }

    memcpy(dest, &buf, cnt * sz);
    return 1;
}
