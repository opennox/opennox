package nxz

import (
	"errors"
	"io"
	"math"
)

type Decoder struct {
	src1     []byte
	err      error
	buf0     [bufferSize]byte
	field4   uint32
	data8    decoderData
	field144 uint32
	field148 uint32
}

type decoderData struct {
	field0   commonData
	field4   [32]uint32
	field132 [tableSize]int16
}

type decoderRec struct {
	ind0 int16
	val2 int16
}

func NewDecoder(src []byte) *Decoder {
	dec := &Decoder{
		src1: src,
	}
	initDecData(&dec.data8)
	return dec
}

func initDecData(d *decoderData) {
	initCommon(&d.field0)
	d.field132 = nxz_table_3
	d.field4 = nxz_table_4
}

func (dec *Decoder) Free() {
	if dec == nil {
		return
	}
	freeDecData(&dec.data8)
}

func freeDecData(d *decoderData) {
	freeCommon(&d.field0)
}

func sub57DDE0(arr []decoderRec) {
	for i := 40; i > 0; i /= 3 {
		j := i + 1
		v11 := i + 1
		if i+1 <= len(arr) {
			for {
				v6 := j * 4
				v10 := j
				val := arr[j-1]
				if j > i {
					v7 := i * 4
					for {
						v8 := int32(arr[(v6-v7)/4-1].val2) - int32(val.val2)
						if v8 == 0 {
							v8 = int32(arr[(v6-v7)/4-1].ind0) - int32(val.ind0)
						}
						if v8 >= 0 {
							break
						}
						arr[v6/4-1] = arr[(v6-v7)/4-1]
						v6 -= v7
						v10 -= i
						if v10 <= i {
							break
						}
					}
					j = v11
				}
				j++
				arr[v10-1] = val
				v11 = j
				if j > len(arr) {
					break
				}
			}
		}
	}
}
func sub57DEA0(d *decoderData, arr []decoderRec) int {
	it := arr
	n := 0
	for i := range d.field0.field0 {
		p := &d.field0.field0[i]
		it[0] = decoderRec{ind0: int16(i), val2: *p}
		it = it[1:]
		n += int(*p)
		*p /= 2
	}
	sub57DDE0(arr[:tableSize])
	return n
}
func (dec *Decoder) Decode(dst []byte) (int, error) {
	if len(dec.src1) == 0 {
		if dec.err != nil {
			return 0, dec.err
		}
		return 0, io.EOF
	}
	n, err := dec.decode(dst)
	dec.err = err
	return n, err
}
func (dec *Decoder) readByte() (byte, bool) {
	if len(dec.src1) == 0 {
		return 0, false
	}
	b := dec.src1[0]
	dec.src1 = dec.src1[1:]
	return b, true
}
func (dec *Decoder) hasMore() bool {
	return len(dec.src1) > 0
}
func (dec *Decoder) readValA() int {
	if v8 := int32(dec.field148); v8 < 4 {
		b, ok := dec.readByte()
		if !ok {
			dec.field148 = 0
			return -1
		}
		dec.field144 |= uint32(int32(b) << (24 - v8))
		dec.field148 = uint32(v8 + 8)
	}
	val := dec.field144
	dec.field144 = val * 16
	dec.field148 -= 4
	return int(val >> 28)
}
func (dec *Decoder) readValB(v12 uint32) int {
	if dec.field148 >= v12 {
		val := int(dec.field144 >> (32 - v12))
		dec.field144 <<= v12
		dec.field148 -= v12
		return val
	} else if b, ok := dec.readByte(); ok {
		dec.field144 |= uint32(int32(b) << (24 - dec.field148))
		dec.field148 += 8
		val := int(dec.field144 >> (32 - v12))
		dec.field144 <<= v12
		dec.field148 -= v12
		return val
	} else {
		dec.field148 = 0
		return -1
	}
}
func (dec *Decoder) decode(dst []byte) (int, error) {
	var (
		v22  int32
		v28  int
		v29  int32
		v30  int32
		v31  int32
		v32  int32
		v33  int32
		v34  int32
		v35  int32
		v36  uint32
		v37  int32
		v38  int32
		v39  int32
		v40  int32
		v41  uint32
		v42  int32
		v43  int32
		v44  int32
		v45i int
		v46  int32
		v47  int32
		v49  int32
		v50  int32
		v51  int32
		v52  []byte
		v54  int32
		v55  int32
		v59  []byte
		v60  uint32
		v61i int
		v65  int32
		v66  int32
		v67  int32
		v68i int
		v70  int32
		v71  int32
	)
	var recs [tableSize]decoderRec
	dstInd := 0
	dec.field148 = 0
	for {
		if !dec.hasMore() {
			return 0, io.ErrUnexpectedEOF
		}
		v9 := dec.readValA()
		var v13 int
		if v12 := dec.data8.field4[2*v9]; v12 == 0 {
			v13 = int(dec.data8.field132[dec.data8.field4[v9*2+1]])
		} else {
			v15 := dec.readValB(v12)
			ind := int(uint32(v15) + dec.data8.field4[v9*2+1])
			if ind >= tableSize {
				return 0, errors.New("wrong table index")
			}
			v13 = int(dec.data8.field132[ind])
		}
		dec.data8.field0.field0[v13]++
		if v13 < 256 {
			if dstInd >= len(dst) {
				return 0, io.ErrShortBuffer
			}
			dst[dstInd] = uint8(int8(v13))
			dstInd++
			dec.buf0[dec.field4%bufferSize] = uint8(int8(v13))
			dec.field4++
			continue
		} else if v13 == 272 {
			sub57DEA0(&dec.data8, recs[:])
			for i := range dec.data8.field132 {
				dec.data8.field132[i] = recs[i].ind0
			}
			v22 = 0
			v23a := dec.data8.field4[:]
			v70 = 0
			v66 = 16
			for {
				var ind int
				for {
					if int(dec.field148) < 1 {
						b, ok := dec.readByte()
						if !ok {
							dec.field148 = 0
							break
						}
						dec.field144 |= uint32(b) << (24 - dec.field148)
						dec.field148 += 8
					}
					stop := (dec.field144 >> 31) != 0
					dec.field144 *= 2
					dec.field148--
					if stop {
						break
					}
					ind++
				}
				v22 += int32(ind)
				v23a[0] = uint32(v22)
				v23a[1] = uint32(v70)
				v23a = v23a[2:]
				v70 += 1 << v22
				if func() int32 {
					p := &v66
					*p--
					return *p
				}() == 0 {
					break
				}
			}
			continue
		}
		if v13 == 273 {
			return dstInd, nil
		}
		if v13 < 264 {
			v28 = v13 - 256
			goto LABEL_43
		}
		v29 = int32(nxz_table_1[v13].v1)
		v30 = int32(dec.field148)
		if v30 >= v29 {
			v31 = int32(dec.field144 >> uint32(32-v29))
			v33 = int32(dec.field144 << uint32(v29))
			dec.field148 -= uint32(v29)
			dec.field144 = uint32(v33)
		} else if b, ok := dec.readByte(); ok {
			v32 = int32(uint32(int32(b)<<(24-v30)) | dec.field144)
			dec.field144 = uint32(v32)
			dec.field148 = uint32(v30 + 8)
			v31 = int32(dec.field144 >> uint32(32-v29))
			v33 = int32(dec.field144 << uint32(v29))
			dec.field148 -= uint32(v29)
			dec.field144 = uint32(v33)
		} else {
			dec.field148 = 0
			v31 = -1
		}
		v28 = int(uint32(v31) + nxz_table_1[v13].v2)
	LABEL_43:
		v34 = int32(dec.field148)
		v67 = int32(v28)
		if v34 < 3 {
			b, ok := dec.readByte()
			if !ok {
				dec.field148 = 0
				v71 = -1
				goto LABEL_48
			}
			v35 = int32(uint32(int32(b)<<(24-v34)) | dec.field144)
			dec.field144 = uint32(v35)
			dec.field148 = uint32(v34 + 8)
		}
		v36 = dec.field144
		dec.field144 = v36 * 8
		v71 = int32(v36 >> 29)
		dec.field148 -= 3
	LABEL_48:
		v65 = 0
		v37 = int32(nxz_table_2[v71+1].v1 + 9)
		if v37 > 8 {
			v38 = int32(dec.field148)
			v37 = int32(nxz_table_2[v71+1].v1 + 1)
			if v38 >= 8 {
				v41 = dec.field144
				dec.field144 = v41 << 8
				dec.field148 -= 8
				v39 = int32(v41 >> 24)
			} else if b, ok := dec.readByte(); ok {
				v40 = int32(uint32(int32(b)<<(24-v38)) | dec.field144)
				dec.field144 = uint32(v40)
				dec.field148 = uint32(v38 + 8)
				v41 = dec.field144
				dec.field144 = v41 << 8
				dec.field148 -= 8
				v39 = int32(v41 >> 24)
			} else {
				dec.field148 = 0
				v39 = -1
			}
			v65 = v39 << v37
		}
		v42 = int32(dec.field148)
		if v42 >= v37 {
			v43 = int32(dec.field144 >> uint32(32-v37))
			v44 = int32(dec.field144 << uint32(v37))
			dec.field148 -= uint32(v37)
			dec.field144 = uint32(v44)
		} else if b, ok := dec.readByte(); ok {
			dec.field144 |= uint32(int32(b) << (24 - v42))
			dec.field148 = uint32(v42 + 8)
			v43 = int32(dec.field144 >> uint32(32-v37))
			v44 = int32(dec.field144 << uint32(v37))
			dec.field148 -= uint32(v37)
			dec.field144 = uint32(v44)
		} else {
			dec.field148 = 0
			v43 = -1
		}
		v45i = dstInd
		v46 = int32((nxz_table_2[v71+1].v2 << 9) + uint32(v65|v43))
		v47 = v67 + 4
		v68i = dstInd + int(v67) + 4
		if v68i > len(dst) {
			return 0, io.ErrShortBuffer
		}
		if v48 := int(dec.field4 - uint32(v46)); v47 >= v46 {
			v50 = int32(uint16(int16(v48)))
			var v53 int
			if uint32(int32(uint16(int16(v48)))+v46) <= bufferSize {
				v53 = int(v46)
				v52 = dec.buf0[v50:]
			} else {
				v51 = int32(bufferSize - uint32(uint16(int16(v48))))
				n := bufferSize - int(v48)
				copy(dst[dstInd:dstInd+n], dec.buf0[v48:v48+n])
				v52 = dec.buf0[:]
				v53 = int(v46 - v51)
				v45i = dstInd + int(v51)
			}
			copy(dst[v45i:v45i+v53], v52[:v53])
			v54 = 0
			v55 = v47 - v46
			if v47-v46 > 0 {
				for {
					dst[int32(dstInd)+v46+v54] = dst[dstInd+int(v54)]
					v54++
					if v54 >= v55 {
						break
					}
				}
			}
		} else {
			v49 = int32(uint16(int16(v48)))
			if uint32(v49+v47) <= bufferSize {
				copy(dst[dstInd:dstInd+int(v47)], dec.buf0[v49:v49+v47])
			} else {
				n := bufferSize - int(v49)
				copy(dst[dstInd:dstInd+n], dec.buf0[v49:int(v49)+n])
				copy(dst[dstInd+n:dstInd+int(v47)], dec.buf0[:int(v47)-n])
			}
		}
		if v57 := int(dec.field4 & math.MaxUint16); v57+int(v47) <= bufferSize {
			v61i = dstInd
			v60 = uint32(v47)
			v59 = dec.buf0[v57:]
		} else {
			n := bufferSize - v57
			copy(dec.buf0[v57:v57+n], dst[dstInd:dstInd+n])
			v59 = dec.buf0[:]
			v60 = uint32(int(v47) - n)
			v61i = dstInd + n
		}
		copy(v59[:v60], dst[v61i:v61i+int(v60)])
		dec.field4 += uint32(v47)
		dstInd = v68i
	}
}
