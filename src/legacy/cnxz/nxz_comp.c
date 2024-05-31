#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

uint32_t* sub57DD90(uint32_t* this);
void sub57DDC0(void** this);

#define HIBYTE(x) (((uint8_t*)(&x))[sizeof(x) - 1])
#define BYTE2(x) (((uint8_t*)(&x))[2])
#define LOWORD(x) (((uint16_t*)(&x))[0])
#define HIWORD(x) (((uint16_t*)(&x))[1])
#define SHIWORD(x) (((short*)(&x))[1])

static inline unsigned int __ROL4__(unsigned int value, int shift) {
	unsigned int c = shift & 31;
	if (c == 0)
		return value;
	return (value << c) | (value >> (32 - c));
}

typedef struct {
	uint32_t v1;
	uint32_t v2;
} uint32_vec2;

typedef struct {
	uint32_t v1;
	uint32_t v2;
	uint32_t v3;
} uint32_vec3;

uint32_vec2 nxz_table_5[128] = {
	{0x3, 0x0},
	{0x3, 0x1},
	{0x4, 0x4},
	{0x4, 0x5},
	{0x5, 0xc},
	{0x5, 0xd},
	{0x5, 0xe},
	{0x5, 0xf},
	{0x6, 0x20},
	{0x6, 0x21},
	{0x6, 0x22},
	{0x6, 0x23},
	{0x6, 0x24},
	{0x6, 0x25},
	{0x6, 0x26},
	{0x6, 0x27},
	{0x7, 0x50},
	{0x7, 0x51},
	{0x7, 0x52},
	{0x7, 0x53},
	{0x7, 0x54},
	{0x7, 0x55},
	{0x7, 0x56},
	{0x7, 0x57},
	{0x7, 0x58},
	{0x7, 0x59},
	{0x7, 0x5a},
	{0x7, 0x5b},
	{0x7, 0x5c},
	{0x7, 0x5d},
	{0x7, 0x5e},
	{0x7, 0x5f},
	{0x8, 0xc0},
	{0x8, 0xc1},
	{0x8, 0xc2},
	{0x8, 0xc3},
	{0x8, 0xc4},
	{0x8, 0xc5},
	{0x8, 0xc6},
	{0x8, 0xc7},
	{0x8, 0xc8},
	{0x8, 0xc9},
	{0x8, 0xca},
	{0x8, 0xcb},
	{0x8, 0xcc},
	{0x8, 0xcd},
	{0x8, 0xce},
	{0x8, 0xcf},
	{0x8, 0xd0},
	{0x8, 0xd1},
	{0x8, 0xd2},
	{0x8, 0xd3},
	{0x8, 0xd4},
	{0x8, 0xd5},
	{0x8, 0xd6},
	{0x8, 0xd7},
	{0x8, 0xd8},
	{0x8, 0xd9},
	{0x8, 0xda},
	{0x8, 0xdb},
	{0x8, 0xdc},
	{0x8, 0xdd},
	{0x8, 0xde},
	{0x8, 0xdf},
	{0x9, 0x1c0},
	{0x9, 0x1c1},
	{0x9, 0x1c2},
	{0x9, 0x1c3},
	{0x9, 0x1c4},
	{0x9, 0x1c5},
	{0x9, 0x1c6},
	{0x9, 0x1c7},
	{0x9, 0x1c8},
	{0x9, 0x1c9},
	{0x9, 0x1ca},
	{0x9, 0x1cb},
	{0x9, 0x1cc},
	{0x9, 0x1cd},
	{0x9, 0x1ce},
	{0x9, 0x1cf},
	{0x9, 0x1d0},
	{0x9, 0x1d1},
	{0x9, 0x1d2},
	{0x9, 0x1d3},
	{0x9, 0x1d4},
	{0x9, 0x1d5},
	{0x9, 0x1d6},
	{0x9, 0x1d7},
	{0x9, 0x1d8},
	{0x9, 0x1d9},
	{0x9, 0x1da},
	{0x9, 0x1db},
	{0x9, 0x1dc},
	{0x9, 0x1dd},
	{0x9, 0x1de},
	{0x9, 0x1df},
	{0x9, 0x1e0},
	{0x9, 0x1e1},
	{0x9, 0x1e2},
	{0x9, 0x1e3},
	{0x9, 0x1e4},
	{0x9, 0x1e5},
	{0x9, 0x1e6},
	{0x9, 0x1e7},
	{0x9, 0x1e8},
	{0x9, 0x1e9},
	{0x9, 0x1ea},
	{0x9, 0x1eb},
	{0x9, 0x1ec},
	{0x9, 0x1ed},
	{0x9, 0x1ee},
	{0x9, 0x1ef},
	{0x9, 0x1f0},
	{0x9, 0x1f1},
	{0x9, 0x1f2},
	{0x9, 0x1f3},
	{0x9, 0x1f4},
	{0x9, 0x1f5},
	{0x9, 0x1f6},
	{0x9, 0x1f7},
	{0x9, 0x1f8},
	{0x9, 0x1f9},
	{0x9, 0x1fa},
	{0x9, 0x1fb},
	{0x9, 0x1fc},
	{0x9, 0x1fd},
	{0x9, 0x1fe},
	{0x9, 0x1ff},
};

uint32_t nxz_table_6[274] = {
	0x140007,
	0x300008,
	0x310008,
	0x320008,
	0x330008,
	0x340008,
	0x350008,
	0x360008,
	0x370008,
	0x380008,
	0x390008,
	0x3a0008,
	0x3b0008,
	0x3c0008,
	0x3d0008,
	0x3e0008,
	0x3f0008,
	0x400008,
	0x410008,
	0x420008,
	0x430008,
	0x440008,
	0x450008,
	0x460008,
	0x470008,
	0x480008,
	0x490008,
	0x4a0008,
	0x4b0008,
	0x4c0008,
	0x4d0008,
	0x4e0008,
	0x150007,
	0x4f0008,
	0x500008,
	0x510008,
	0x520008,
	0x530008,
	0x540008,
	0x550008,
	0x560008,
	0x570008,
	0x580008,
	0x590008,
	0x5a0008,
	0x5b0008,
	0x5c0008,
	0x5d0008,
	0x160007,
	0x5e0008,
	0x5f0008,
	0x600008,
	0x610008,
	0x620008,
	0x630008,
	0x640008,
	0x650008,
	0x660008,
	0x670008,
	0x680008,
	0x690008,
	0x6a0008,
	0x6b0008,
	0x6c0008,
	0x6d0008,
	0x6e0008,
	0x6f0008,
	0x700008,
	0x710008,
	0x720008,
	0x730008,
	0x740008,
	0x750008,
	0x760008,
	0x770008,
	0x780008,
	0x790008,
	0x7a0008,
	0x7b0008,
	0x7c0008,
	0x7d0008,
	0x7e0008,
	0x7f0008,
	0x800008,
	0x810008,
	0x820008,
	0x830008,
	0x840008,
	0x850008,
	0x860008,
	0x870008,
	0x880008,
	0x890008,
	0x8a0008,
	0x8b0008,
	0x8c0008,
	0x8d0008,
	0x8e0008,
	0x8f0008,
	0x900008,
	0x910008,
	0x920008,
	0x930008,
	0x940008,
	0x950008,
	0x960008,
	0x970008,
	0x980008,
	0x990008,
	0x9a0008,
	0x9b0008,
	0x9c0008,
	0x9d0008,
	0x9e0008,
	0x9f0008,
	0xa00008,
	0xa10008,
	0xa20008,
	0xa30008,
	0xa40008,
	0xa50008,
	0xa60008,
	0xa70008,
	0xa80008,
	0xa90008,
	0xaa0008,
	0xab0008,
	0xac0008,
	0xad0008,
	0xae0008,
	0xaf0008,
	0xb00008,
	0xb10008,
	0xb20008,
	0xb30008,
	0xb40008,
	0xb50008,
	0xb60008,
	0xb70008,
	0xb80008,
	0xb90008,
	0xba0008,
	0xbb0008,
	0xbc0008,
	0xbd0008,
	0xbe0008,
	0xbf0008,
	0xc00008,
	0xc10008,
	0xc20008,
	0xc30008,
	0xc40008,
	0xc50008,
	0xc60008,
	0xc70008,
	0xc80008,
	0xc90008,
	0xca0008,
	0xcb0008,
	0xcc0008,
	0xcd0008,
	0xce0008,
	0xcf0008,
	0x1a00009,
	0x1a10009,
	0x1a20009,
	0x1a30009,
	0x1a40009,
	0x1a50009,
	0x1a60009,
	0x1a70009,
	0x1a80009,
	0x1a90009,
	0x1aa0009,
	0x1ab0009,
	0x1ac0009,
	0x1ad0009,
	0x1ae0009,
	0x1af0009,
	0x1b00009,
	0x1b10009,
	0x1b20009,
	0x1b30009,
	0x1b40009,
	0x1b50009,
	0x1b60009,
	0x1b70009,
	0x1b80009,
	0x1b90009,
	0x1ba0009,
	0x1bb0009,
	0x1bc0009,
	0x1bd0009,
	0x1be0009,
	0x1bf0009,
	0x1c00009,
	0x1c10009,
	0x1c20009,
	0x1c30009,
	0x1c40009,
	0x1c50009,
	0x1c60009,
	0x1c70009,
	0x1c80009,
	0x1c90009,
	0x1ca0009,
	0x1cb0009,
	0x1cc0009,
	0x1cd0009,
	0x1ce0009,
	0x1cf0009,
	0x1d00009,
	0x1d10009,
	0x1d20009,
	0x1d30009,
	0x1d40009,
	0x1d50009,
	0x1d60009,
	0x1d70009,
	0x1d80009,
	0x1d90009,
	0x1da0009,
	0x1db0009,
	0x1dc0009,
	0x1dd0009,
	0x1de0009,
	0x1df0009,
	0x1e00009,
	0x1e10009,
	0x1e20009,
	0x1e30009,
	0x1e40009,
	0x1e50009,
	0x1e60009,
	0x1e70009,
	0x1e80009,
	0x1e90009,
	0x1ea0009,
	0x1eb0009,
	0x1ec0009,
	0x1ed0009,
	0x1ee0009,
	0x1ef0009,
	0x1f00009,
	0x1f10009,
	0x1f20009,
	0x1f30009,
	0x1f40009,
	0x1f50009,
	0x1f60009,
	0x1f70009,
	0x1f80009,
	0x1f90009,
	0x1fa0009,
	0x1fb0009,
	0x170007,
	0x6,
	0x10006,
	0x20006,
	0x30006,
	0x80007,
	0x90007,
	0xa0007,
	0xb0007,
	0xc0007,
	0xd0007,
	0xe0007,
	0xf0007,
	0x100007,
	0x110007,
	0x120007,
	0x130007,
	0x1fc0009,
	0x1fd0009,
};

uint32_vec3 nxz_table_7[16] = {
	{0x108, 0x1, 0x0},
	{0x109, 0x2, 0x0},
	{0x109, 0x2, 0x2},
	{0x10a, 0x3, 0x0},
	{0x10a, 0x3, 0x2},
	{0x10a, 0x3, 0x4},
	{0x10a, 0x3, 0x6},
	{0x10b, 0x4, 0x0},
	{0x10b, 0x4, 0x2},
	{0x10b, 0x4, 0x4},
	{0x10b, 0x4, 0x6},
	{0x10b, 0x4, 0x8},
	{0x10b, 0x4, 0xa},
	{0x10b, 0x4, 0xc},
	{0x10b, 0x4, 0xe},
	{0x0, 0x3, 0x0},
};

//----- (0057DF00) --------------------------------------------------------
uint32_t* sub_57DF00(uint32_t* this) {
	uint32_t* v1;      // ebx
	unsigned char* v2; // eax
	unsigned char* v3; // edi
	uint32_t* result;  // eax

	v1 = this;
	sub57DD90(this);
	v1[1] = 4096;
	v2 = calloc(1, sizeof(nxz_table_6));
	v3 = v2;
	v1[2] = (uint32_t)v2;
	result = v1;
	memcpy(v3, nxz_table_6, sizeof(nxz_table_6));
	return result;
}

//----- (00578B80) --------------------------------------------------------
void* nxz_compress_new() {
	uint32_t* v0; // eax
	uint32_t* v2; // esi
	int v3;       // eax

	v0 = calloc(1, 0x18u);
	if (!v0) {
		return 0;
	}
	v2 = v0;
	*v0 = (uint32_t)calloc(1, 0x10000u);
	v2[1] = 0;
	sub_57DF00(v2 + 2);
	v2[5] = (uint32_t)calloc(1, 0x10000u);
	v3 = 0;
	do {
		v3 += 2;
		*(uint16_t*)(v3 + v2[5] - 2) = -1;
	} while (v3 < 0x10000);
	return v2;
}
// 5667CB: using guessed type void * calloc(1, unsigned int);

//----- (0057DF70) --------------------------------------------------------
void sub57DF70(void** this) {
	void** v1; // esi

	v1 = this;
	free(this[2]);
	sub57DDC0(v1);
}

//----- (0057D150) --------------------------------------------------------
void sub57D150(void** this) {
	void** v1; // esi

	v1 = this;
	free(this[5]);
	sub57DF70(v1 + 2);
	free(*v1);
}

//----- (00578BD0) --------------------------------------------------------
void nxz_compress_free(void* lpMem) {
	if (lpMem) {
		sub57D150((void**)lpMem);
		free(lpMem);
	}
}

//----- (0057DDE0) --------------------------------------------------------
unsigned int sub_57DDE0(int a1, int a2) {
	int v2;              // ecx
	int v3;              // esi
	int i;               // ebx
	int v5;              // eax
	int v6;              // edx
	int v7;              // edi
	int v8;              // esi
	unsigned int result; // eax
	int v10;             // [esp+10h] [ebp-8h]
	int v11;             // [esp+14h] [ebp-4h]
	int v12;             // [esp+1Ch] [ebp+4h]

	v2 = a1;
	v3 = a2;
	for (i = 40; i > 0; i /= 3) {
		v5 = i + 1;
		v11 = i + 1;
		if (i + 1 <= v3) {
			do {
				v6 = 4 * v5;
				v10 = v5;
				v12 = *(uint32_t*)(4 * v5 + v2);
				if (v5 > i) {
					v7 = 4 * i;
					do {
						v8 = *(short*)(v6 - v7 + v2 + 2) - SHIWORD(v12);
						if (!v8) {
							v8 = *(short*)(v6 - v7 + v2) - (short)v12;
						}
						if (v8 >= 0) {
							break;
						}
						*(uint32_t*)(v6 + v2) = *(uint32_t*)(v6 - v7 + v2);
						v6 -= v7;
						v10 -= i;
					} while (v10 > i);
					v5 = v11;
				}
				++v5;
				*(uint32_t*)(v2 + 4 * v10) = v12;
				v3 = a2;
				v11 = v5;
			} while (v5 <= a2);
		}
		result = (unsigned int)(i / 3) >> 31;
	}
	return result;
}

//----- (0057DEA0) --------------------------------------------------------
int sub_57DEA0(uint32_t* this, uint16_t* a2) {
	uint16_t* v2; // esi
	int v3;       // ebx
	int i;        // eax
	short* v5;    // edi
	short v6;     // dx

	v2 = a2;
	v3 = 0;
	for (i = 0; i < 274; ++i) {
		*v2 = i;
		v2 += 2;
		*(v2 - 1) = *(uint16_t*)(*this + 2 * i);
		v5 = (short*)(*this + 2 * i);
		v6 = *v5;
		v3 += v6;
		*v5 = v6 >> 1;
	}
	sub_57DDE0((int)(a2 - 2), 274);
	return v3;
}

//----- (0057DFC0) --------------------------------------------------------
int sub_57DFC0(uint32_t* this, int* a2) {
	uint32_t* v2;   // esi
	int v3;         // eax
	int v4;         // ecx
	int v5;         // edi
	int v6;         // ebp
	int v7;         // esi
	int v8;         // ebx
	int v9;         // eax
	int v10;        // edx
	int v11;        // eax
	char* v12;      // edx
	int v13;        // eax
	int* v14;       // edx
	int v15;        // edx
	int v16;        // ebx
	int v17;        // esi
	char* v18;      // eax
	int v19;        // ecx
	int v20;        // esi
	int v21;        // eax
	int v22;        // ecx
	int v23;        // eax
	char* v24;      // ecx
	int v25;        // ecx
	int v26;        // eax
	int* v27;       // edi
	int v28;        // ecx
	int* v29;       // eax
	int v30;        // ecx
	int* v31;       // eax
	int v32;        // ebx
	int v33;        // esi
	int v34;        // ecx
	int v35;        // edi
	int result;     // eax
	int v37;        // edx
	short v38;      // di
	int v39;        // esi
	char* v40;      // ecx
	int v41;        // eax
	short v42;      // bp
	int* v43;       // [esp+10h] [ebp-46Ch]
	int v44;        // [esp+10h] [ebp-46Ch]
	int v45;        // [esp+10h] [ebp-46Ch]
	int v46;        // [esp+14h] [ebp-468h]
	int v47;        // [esp+14h] [ebp-468h]
	int* v48;       // [esp+14h] [ebp-468h]
	int v49;        // [esp+18h] [ebp-464h]
	int v50;        // [esp+18h] [ebp-464h]
	int v51;        // [esp+18h] [ebp-464h]
	int v52;        // [esp+1Ch] [ebp-460h]
	int v53;        // [esp+1Ch] [ebp-460h]
	int v54;        // [esp+1Ch] [ebp-460h]
	int v55;        // [esp+20h] [ebp-45Ch]
	int v56;        // [esp+24h] [ebp-458h]
	int v57;        // [esp+24h] [ebp-458h]
	int v58;        // [esp+28h] [ebp-454h]
	int v59;        // [esp+2Ch] [ebp-450h]
	char v60[1100]; // [esp+30h] [ebp-44Ch]

	v2 = this;
	*(uint32_t*)v60 = (uint32_t)this;
	v3 = sub_57DEA0(this, (uint16_t*)&v60[4]);
	v4 = 0;
	v5 = 0;
	v6 = 16;
	v56 = v3;
	v2[1] = 4096;
	v55 = 0;
	v46 = 0;
	v52 = 16;
	v43 = a2 - 1;
	while (1) {
		v7 = 0;
		v58 = 0;
		v8 = 0;
		v59 = (v56 - v55) / v6;
		while (1) {
			v49 = 0;
			v9 = 1 << v4;
			if (v5 + (1 << v4) + v7 > 274) {
				v49 = 1;
				v9 = 274 - v5;
			}
			if (v7 < v9) {
				v10 = v7 + v5;
				v11 = v9 - v7;
				v7 += v11;
				v12 = &v60[4 * v10 + 6];
				do {
					v8 += *(short*)v12;
					v12 += 4;
					--v11;
				} while (v11);
				v6 = v52;
			}
			if (v49 || v4 >= 8 || v8 > v59) {
				break;
			}
			v58 = v8;
			++v4;
		}
		if (v4 && abs(v58 - v59) <= abs(v8 - v59)) {
			v8 = v58;
			--v4;
		}
		v13 = v46;
		if (v6 < 16) {
			v14 = v43;
			do {
				if (v4 >= *v14) {
					break;
				}
				v14[1] = *v14;
				--v13;
				--v14;
			} while (v13 > 0);
		}
		a2[v13] = v4;
		v55 += v8;
		++v43;
		v5 += 1 << v4;
		--v6;
		++v46;
		v52 = v6;
		if (v6 <= 2) {
			break;
		}
		v4 = 0;
	}
	v15 = 0;
	v16 = 0;
	v17 = 0;
	v44 = 0;
	v47 = 0;
	v50 = 0x7FFFFFFF;
	v53 = 0;
	if (v5 < 274) {
		v18 = &v60[4 * v5 + 6];
		v19 = 274 - v5;
		do {
			v17 += *(short*)v18;
			v18 += 4;
			--v19;
		} while (v19);
		v53 = v17;
	}
	v20 = 0;
	v21 = 1;
	if (v5 + 1 <= 274) {
		do {
			if (v15 < v21) {
				v22 = v15 + v5;
				v23 = v21 - v15;
				v15 += v23;
				v24 = &v60[4 * v22 + 6];
				do {
					v16 += *(short*)v24;
					v24 += 4;
					--v23;
				} while (v23);
			}
			v25 = 0;
			v26 = 274 - v15 - v5;
			if (v26 > 1) {
				do {
					++v25;
				} while (1 << v25 < v26);
			}
			if (v20 <= 8 && v25 <= 8) {
				if (v16 * v20 + v25 * (v53 - v16) >= v50) {
					break;
				}
				v50 = v16 * v20 + v25 * (v53 - v16);
				v44 = v20;
				v47 = v25;
			}
			v21 = 1 << ++v20;
		} while (v5 + (1 << v20) + v15 <= 274);
	}
	v27 = a2;
	v28 = 14;
	v29 = a2 + 13;
	do {
		if (v44 >= *v29) {
			break;
		}
		v29[1] = *v29;
		--v28;
		--v29;
	} while (v28 > 0);
	a2[v28] = v44;
	v30 = 15;
	v31 = a2 + 14;
	do {
		if (v47 >= *v31) {
			break;
		}
		v31[1] = *v31;
		--v30;
		--v31;
	} while (v30 > 0);
	a2[v30] = v47;
	v32 = 0;
	v33 = 0;
	v51 = 0;
	v45 = 0;
	v48 = a2;
	do {
		v34 = *v27;
		v35 = 1 << *v27;
		result = 274 - v32;
		v57 = v35;
		if (v35 < 274 - v32) {
			result = v35;
		}
		v37 = 0;
		v54 = result;
		if (result > 0) {
			v38 = v34 + 4;
			v39 = v33 << v34;
			v40 = &v60[4 * v32 + 4];
			do {
				v41 = *(short*)v40;
				v40 += 4;
				*(uint16_t*)(*(uint32_t*)(*(uint32_t*)v60 + 8) + 4 * v41) = v38;
				v42 = v37++ | v39;
				*(uint16_t*)(*(uint32_t*)(*(uint32_t*)v60 + 8) + 4 * v41 + 2) = v42;
				result = v54;
			} while (v37 < v54);
			v35 = v57;
			v33 = v45;
			v32 = v51;
		}
		v32 += v35;
		++v33;
		v27 = v48 + 1;
		v51 = v32;
		v45 = v33;
		++v48;
	} while (v33 < 16);
	return result;
}

//----- (0057E2C0) --------------------------------------------------------
int sub_57E2C0(uint32_t** this) {
	uint32_t** v1;     // esi
	uint32_t* v2;      // edi
	unsigned int v3;   // ebp
	int v4;            // eax
	int v5;            // edx
	unsigned short v6; // bx
	int v7;            // ebp
	uint32_t* v8;      // edx
	uint8_t* v9;       // ebx
	int v10;           // ecx
	uint32_t* v11;     // eax
	int v12;           // edi
	char* v13;         // ebx
	uint32_t* v14;     // edx
	int v15;           // eax
	char v16;          // cl
	int v17;           // eax
	int v18;           // ebp
	unsigned int v19;  // ecx
	int v20;           // ecx
	uint8_t* v21;      // eax
	unsigned int v22;  // ecx
	uint32_t* v23;     // eax
	int v24;           // ecx
	int result;        // eax
	int v26;           // [esp+10h] [ebp-44h]
	char v27[64];      // [esp+14h] [ebp-40h]

	v1 = this;
	*this[2] = 2;
	if ((int)--*this[2] <= 0) {
		sub_57E2C0(this);
	}
	++*((uint16_t*)v1[1] + 272);
	v2 = v1[6];
	v3 = (unsigned int)v1[5];
	v4 = (*v1)[2];
	v5 = *(short*)(v4 + 1088);
	v6 = *(uint16_t*)(v4 + 1090);
	v1[6] = (uint32_t*)((char*)v2 + v5);
	v7 = (v6 << (32 - (uint8_t)v2 - v5)) | v3;
	v1[5] = (uint32_t*)v7;
	if ((int)v2 + v5 >= 16) {
		*(uint8_t*)v1[3] = HIBYTE(v7);
		v8 = v1[5];
		v9 = (char*)v1[3] + 1;
		v1[3] = (uint32_t*)v9;
		*v9 = BYTE2(v8);
		v10 = (int)(v1[6] - 4);
		v11 = (uint32_t*)((uint32_t)v1[5] << 16);
		v1[3] = (uint32_t*)((char*)v1[3] + 1);
		v1[6] = (uint32_t*)v10;
		v1[5] = v11;
	}
	sub_57DFC0(*v1, (int*)v27);
	v12 = 0;
	v13 = v27;
	v26 = 16;
	do {
		v14 = v1[6];
		v15 = *(uint32_t*)v13 - v12;
		v12 = *(uint32_t*)v13;
		v16 = 32 - (uint8_t)v14 - ++v15;
		v17 = (int)v14 + v15;
		v18 = 1 << v16;
		v19 = (unsigned int)v1[5];
		v1[6] = (uint32_t*)v17;
		v20 = v18 | v19;
		v1[5] = (uint32_t*)v20;
		if (v17 >= 16) {
			*(uint8_t*)v1[3] = HIBYTE(v20);
			v21 = (char*)v1[3] + 1;
			v22 = (unsigned int)v1[5] >> 16;
			v1[3] = (uint32_t*)v21;
			*v21 = v22;
			v23 = v1[5];
			v24 = (int)(v1[6] - 4);
			v1[3] = (uint32_t*)((char*)v1[3] + 1);
			v1[6] = (uint32_t*)v24;
			v1[5] = (uint32_t*)((uint32_t)v23 << 16);
		}
		v13 += 4;
		result = --v26;
	} while (v26);
	return result;
}

//----- (0057E7D0) --------------------------------------------------------
int sub_57E7D0(uint32_t** this) {
	uint32_t** v1; // esi
	uint32_t* v2;  // edi
	int v3;        // edx
	uint32_t* v4;  // ecx
	uint8_t* v5;   // ebx
	int v6;        // ecx
	uint32_t* v7;  // eax
	uint32_t* v8;  // edi
	int v9;        // ebx

	v1 = this;
	if ((int)--*this[2] <= 0) {
		sub_57E2C0(this);
	}
	++*((uint16_t*)v1[1] + 273);
	v2 = v1[6];
	v3 = *(short*)((*v1)[2] + 1092);
	v1[5] = (uint32_t*)((*(unsigned short*)((*v1)[2] + 1094) << (32 - (uint8_t)v2 - v3)) | (unsigned int)v1[5]);
	v1[6] = (uint32_t*)((char*)v2 + v3);
	if ((int)v2 + v3 >= 16) {
		*(uint8_t*)v1[3] = (unsigned int)v1[5] >> 24;
		v4 = v1[5];
		v5 = (char*)v1[3] + 1;
		v1[3] = (uint32_t*)v5;
		*v5 = BYTE2(v4);
		v6 = (int)(v1[6] - 4);
		v7 = (uint32_t*)((uint32_t)v1[5] << 16);
		v1[3] = (uint32_t*)((char*)v1[3] + 1);
		v1[6] = (uint32_t*)v6;
		v1[5] = v7;
	}
	if ((int)v1[6] > 0) {
		do {
			*(uint8_t*)v1[3] = (unsigned int)v1[5] >> 24;
			v8 = v1[5];
			v9 = (int)(v1[6] - 2);
			v1[3] = (uint32_t*)((char*)v1[3] + 1);
			v1[6] = (uint32_t*)v9;
			v1[5] = (uint32_t*)((uint32_t)v8 << 8);
		} while (v9 > 0);
	}
	return (char*)v1[3] - (char*)v1[4];
}



//----- (0057E3F0) --------------------------------------------------------
unsigned int sub_57E3F0(uint32_t** this, unsigned int a2, int a3) {
	unsigned char* v3;   // ebx
	unsigned int result; // eax
	uint32_t** v5;       // esi
	bool v6;             // cf
	unsigned short v7;   // di
	int v8;              // eax
	uint32_t* v9;        // edi
	int v10;             // edx
	unsigned short v11;  // bp
	int v12;             // edx
	uint32_t* v13;       // edx
	uint8_t* v14;        // ebp
	int v15;             // ecx
	uint32_t* v16;       // eax
	unsigned int v17;    // [esp+Ch] [ebp+4h]

	v3 = (unsigned char*)a2;
	result = a2 + a3;
	v5 = this;
	v6 = a2 < a2 + a3;
	v17 = a2 + a3;
	if (v6) {
		do {
			v7 = *v3;
			if ((int)--*v5[2] <= 0) {
				sub_57E2C0(v5);
			}
			v8 = v7;
			++*((uint16_t*)v5[1] + v7);
			v9 = v5[6];
			v10 = (*v5)[2];
			v11 = *(uint16_t*)(v10 + 4 * v8 + 2);
			v12 = *(short*)(v10 + 4 * v8);
			v5[5] = (uint32_t*)((v11 << (32 - (uint8_t)v9 - v12)) | (unsigned int)v5[5]);
			v5[6] = (uint32_t*)((char*)v9 + v12);
			if ((int)v9 + v12 >= 16) {
				*(uint8_t*)v5[3] = (unsigned int)v5[5] >> 24;
				v13 = v5[5];
				v14 = (char*)v5[3] + 1;
				v5[3] = (uint32_t*)v14;
				*v14 = BYTE2(v13);
				v15 = (int)(v5[6] - 4);
				v16 = (uint32_t*)((uint32_t)v5[5] << 16);
				v5[3] = (uint32_t*)((char*)v5[3] + 1);
				v5[6] = (uint32_t*)v15;
				v5[5] = v16;
			}
			result = v17;
			++v3;
		} while ((unsigned int)v3 < v17);
	}
	return result;
}

//----- (0057F160) --------------------------------------------------------
uint32_t* sub_57F160(int this, int a2, int a3) {
	uint32_t* result; // eax
	int v4;           // edx
	uint8_t* v5;      // ecx
	int v6;           // edx
	int v7;           // ecx
	int v8;           // edx

	result = (uint32_t*)this;
	v4 = *(uint32_t*)(this + 24);
	*(uint32_t*)(this + 20) |= a3 << (32 - v4 - a2);
	*(uint32_t*)(this + 24) = v4 + a2;
	if (v4 + a2 >= 16) {
		**(uint8_t**)(this + 12) = *(uint32_t*)(this + 20) >> 24;
		v5 = (uint8_t*)(*(uint32_t*)(this + 12) + 1);
		v6 = result[5] >> 16;
		result[3] = (uint32_t)v5;
		*v5 = v6;
		v7 = result[5];
		v8 = result[6] - 16;
		++result[3];
		result[6] = v8;
		result[5] = v7 << 16;
	}
	return result;
}

//----- (0057E4C0) --------------------------------------------------------
int sub_57E4C0(uint32_t** this, unsigned int a2, int a3, unsigned int a4, unsigned int a5) {
	uint32_t** v5;      // esi
	unsigned short v6;  // di
	int v7;             // eax
	uint32_t* v8;       // edi
	unsigned int v9;    // ebp
	int v10;            // ecx
	int v11;            // edx
	unsigned short v12; // bx
	unsigned short v13; // bp
	unsigned int v14;   // ebx
	int v15;            // edi
	int v16;            // ecx
	uint32_t* v17;      // edx
	int v18;            // eax
	int v19;            // ebp
	unsigned int v20;   // edi
	int v21;            // ecx
	int v22;            // eax
	unsigned char* v23; // ebx
	unsigned short v24; // bp
	uint32_t* v25;      // edx
	int v26;            // eax
	unsigned int v27;   // edi
	int v28;            // ebx
	int v29;            // eax
	uint32_t* v30;      // ecx
	uint8_t* v31;       // ebx
	int v32;            // ecx
	uint32_t* v33;      // eax
	int v34;            // edx
	unsigned int v35;   // eax
	uint32_t* v36;      // edi
	unsigned int v37;   // ebp
	int v38;            // edx
	int v39;            // ebp
	uint32_t* v40;      // edx
	uint8_t* v41;       // ebp
	int v42;            // edx
	uint32_t* v43;      // ecx
	uint32_t* v44;      // edx
	int v45;            // edi
	int result;         // eax
	uint32_t* v47;      // edi
	char v48;           // cl
	int v49;            // edx
	unsigned int v50;   // ecx
	uint32_t* v51;      // edx
	uint8_t* v52;       // ebx
	int v53;            // ecx

	v5 = this;
	sub_57E3F0(this, a2, a3);
	if (a4 >= 8) {
		if (a4 >= 0x26) {
			v23 = (unsigned char*)&nxz_table_7[(a4 - 38) >> 5];
			v24 = nxz_table_7[(a4 - 38) >> 5].v1 + 4;
			if ((int)--*v5[2] <= 0) {
				sub_57E2C0(v5);
			}
			++*((uint16_t*)v5[1] + v24);
			sub_57F160((int)v5, *(short*)((*v5)[2] + 4 * v24), *(unsigned short*)((*v5)[2] + 4 * v24 + 2));
			v25 = v5[6];
			v26 = *((uint32_t*)v23 + 1) + 4;
			v27 = (unsigned int)v5[5];
			v28 = (((uint8_t)a4 - 38) & 0x1F | (16 * *((unsigned short*)v23 + 4))) << (32 - (uint8_t)v25 - v26);
			v29 = (int)v25 + v26;
			v5[6] = (uint32_t*)v29;
			v5[5] = (uint32_t*)(v28 | v27);
			if (v29 >= 16) {
				goto LABEL_14;
			}
		} else {
			v13 = nxz_table_7[(a4 - 8) >> 1].v1;
			v14 = (a4 - 8) & 1 | nxz_table_7[(a4 - 8) >> 1].v3;
			v15 = nxz_table_7[(a4 - 8) >> 1].v2;
			if ((int)--*v5[2] <= 0) {
				sub_57E2C0(v5);
			}
			++*((uint16_t*)v5[1] + v13);
			v16 = (*v5)[2];
			v17 = v5[6];
			v18 = v15 + *(short*)(v16 + 4 * v13);
			v19 = *(unsigned short*)(v16 + 4 * v13 + 2) << v15;
			v20 = (unsigned int)v5[5];
			v21 = 32 - (uint32_t)v17 - v18;
			v22 = (int)v17 + v18;
			v5[6] = (uint32_t*)v22;
			v5[5] = (uint32_t*)(((v14 | v19) << v21) | v20);
			if (v22 >= 16) {
				goto LABEL_14;
			}
		}
	} else {
		v6 = a4 + 256;
		if ((int)--*v5[2] <= 0) {
			sub_57E2C0(v5);
		}
		v7 = v6;
		++*((uint16_t*)v5[1] + v6);
		v8 = v5[6];
		v9 = (unsigned int)v5[5];
		v10 = (*v5)[2];
		v11 = *(short*)(v10 + 4 * v7);
		v12 = *(uint16_t*)(v10 + 4 * v7 + 2);
		v5[6] = (uint32_t*)((char*)v8 + v11);
		v5[5] = (uint32_t*)((v12 << (32 - (uint8_t)v8 - v11)) | v9);
		if ((int)v8 + v11 >= 16) {
		LABEL_14:
			*(uint8_t*)v5[3] = (unsigned int)v5[5] >> 24;
			v30 = v5[5];
			v31 = (char*)v5[3] + 1;
			v5[3] = (uint32_t*)v31;
			*v31 = BYTE2(v30);
			v32 = (int)(v5[6] - 4);
			v33 = (uint32_t*)((uint32_t)v5[5] << 16);
			v5[3] = (uint32_t*)((char*)v5[3] + 1);
			v5[6] = (uint32_t*)v32;
			v5[5] = v33;
			goto LABEL_15;
		}
	}
LABEL_15:
	v34 = nxz_table_5[a5 >> 9].v1 + 9;
	v35 = a5 & 0x1FF | (nxz_table_5[a5 >> 9].v2 << 9);
	if (v34 <= 16) {
		v47 = v5[6];
		v48 = 32 - (uint8_t)v47 - v34;
		v49 = (int)v47 + v34;
		result = v35 << v48;
		v50 = (unsigned int)v5[5];
		v5[6] = (uint32_t*)v49;
		v5[5] = (uint32_t*)(result | v50);
		if (v49 < 16) {
			return result;
		}
		goto LABEL_21;
	}
	v36 = v5[6];
	v37 = (unsigned int)v5[5];
	v38 = nxz_table_5[a5 >> 9].v1 - 7;
	v5[6] = (uint32_t*)((char*)v36 + v38);
	v39 = (v35 >> 16 << (32 - (uint8_t)v36 - v38)) | v37;
	v5[5] = (uint32_t*)v39;
	if ((int)v36 + v38 >= 16) {
		*(uint8_t*)v5[3] = HIBYTE(v39);
		v40 = v5[5];
		v41 = (char*)v5[3] + 1;
		v5[3] = (uint32_t*)v41;
		*v41 = BYTE2(v40);
		v42 = (int)(v5[6] - 4);
		v43 = (uint32_t*)((uint32_t)v5[5] << 16);
		v5[3] = (uint32_t*)((char*)v5[3] + 1);
		v5[6] = (uint32_t*)v42;
		v5[5] = v43;
	}
	v44 = v5[6];
	v45 = ((unsigned short)v35 << (16 - (uint8_t)v44)) | (unsigned int)v5[5];
	result = (int)(v44 + 4);
	v5[5] = (uint32_t*)v45;
	v5[6] = v44 + 4;
	if ((int)(v44 + 4) >= 16) {
	LABEL_21:
		*(uint8_t*)v5[3] = (unsigned int)v5[5] >> 24;
		v51 = v5[5];
		v52 = (char*)v5[3] + 1;
		v5[3] = (uint32_t*)v52;
		*v52 = BYTE2(v51);
		v53 = (int)(v5[6] - 4);
		result = (uint32_t)v5[5] << 16;
		v5[3] = (uint32_t*)((char*)v5[3] + 1);
		v5[6] = (uint32_t*)v53;
		v5[5] = (uint32_t*)result;
	}
	return result;
}

//----- (0057D1C0) --------------------------------------------------------
int nxz_compress(void* a1p, uint8_t* a2p, uint8_t* a3p, int a4p) {
	void** this = a1p;
	int a2 = (int)a2p;
	unsigned char* a3 = a3p;
	int a4 = a4p;
	void** v4;           // ebx
	void* v5;            // ecx
	int v6;              // eax
	unsigned char* v7;   // esi
	int v8;              // ecx
	int v9;              // eax
	int v10;             // edx
	int v11;             // edi
	int v12;             // ebp
	uint16_t* v13;       // ecx
	unsigned int v14;    // eax
	int v15;             // esi
	uint16_t* v16;       // eax
	int v17;             // ecx
	int v18;             // eax
	int v19;             // ebp
	int v20;             // eax
	unsigned char* i;    // eax
	int v22;             // ecx
	int v23;             // ecx
	int v24;             // eax
	int v25;             // edx
	int v26;             // eax
	int v27;             // ecx
	unsigned char* v28;  // eax
	int v29;             // edx
	int v30;             // eax
	unsigned char v31;   // cl
	int v32;             // eax
	int v33;             // esi
	int v34;             // esi
	char* v35;           // eax
	int v36;             // eax
	int v37;             // esi
	int v38;             // ecx
	unsigned char* v39;  // edx
	unsigned char* v40;  // edi
	unsigned char* v41;  // eax
	unsigned int v42;    // ecx
	int v43;             // edx
	int v44;             // eax
	int v45;             // edi
	int v46;             // edx
	int v47;             // edi
	int v48;             // eax
	char* v49;           // edi
	unsigned char* v50;  // esi
	unsigned int v51;    // eax
	char* v52;           // edi
	unsigned char* v53;  // esi
	char v54;            // cl
	int v55;             // esi
	int v56;             // ecx
	int v57;             // ebp
	unsigned char* v58;  // edx
	int v59;             // ecx
	int v60;             // esi
	unsigned char* v61;  // eax
	unsigned int v62;    // ebp
	int v63;             // ecx
	int v64;             // eax
	int v65;             // esi
	int v66;             // ecx
	int v67;             // edx
	int v68;             // edi
	signed int v69;      // ebp
	int v70;             // eax
	int v71;             // esi
	int v72;             // ecx
	unsigned char* v73;  // edi
	unsigned char* v74;  // eax
	unsigned int v75;    // ecx
	int v76;             // edx
	int v77;             // eax
	int v78;             // edi
	int v79;             // edx
	int v80;             // eax
	char* v81;           // edi
	unsigned char* v82;  // esi
	unsigned int v83;    // eax
	char* v84;           // edi
	int v85;             // ebp
	unsigned char* v86;  // edx
	int v87;             // esi
	int v88;             // ecx
	unsigned char* v89;  // edi
	unsigned char* v90;  // eax
	unsigned int v91;    // ebp
	int v92;             // ecx
	int v93;             // eax
	int v94;             // edi
	int v95;             // eax
	int v96;             // ecx
	int v97;             // edx
	int v98;             // edi
	signed int v99;      // ebp
	int v100;            // eax
	unsigned char* v101; // esi
	int v102;            // edi
	int v103;            // edx
	unsigned char* v104; // edi
	unsigned int v105;   // ecx
	int v106;            // edi
	unsigned char* v107; // edx
	int v108;            // ebp
	signed int v110;     // [esp+10h] [ebp-5Ch]
	int v111;            // [esp+14h] [ebp-58h]
	int v112;            // [esp+18h] [ebp-54h]
	int v113;            // [esp+18h] [ebp-54h]
	int v114;            // [esp+18h] [ebp-54h]
	int v115;            // [esp+18h] [ebp-54h]
	int v116;            // [esp+18h] [ebp-54h]
	int v117;            // [esp+1Ch] [ebp-50h]
	unsigned char* v118; // [esp+20h] [ebp-4Ch]
	int v119;            // [esp+24h] [ebp-48h]
	int v120;            // [esp+24h] [ebp-48h]
	int v121;            // [esp+24h] [ebp-48h]
	unsigned short v122; // [esp+24h] [ebp-48h]
	int v123;            // [esp+24h] [ebp-48h]
	int v124;            // [esp+24h] [ebp-48h]
	int v125;            // [esp+28h] [ebp-44h]
	int v126;            // [esp+2Ch] [ebp-40h]
	int v127;            // [esp+30h] [ebp-3Ch]
	short v128;          // [esp+34h] [ebp-38h]
	int v129;            // [esp+38h] [ebp-34h]
	int v130;            // [esp+38h] [ebp-34h]
	int v131;            // [esp+38h] [ebp-34h]
	int v132;            // [esp+38h] [ebp-34h]
	void* v133;          // [esp+3Ch] [ebp-30h]
	int v134;            // [esp+40h] [ebp-2Ch]
	int v135[10];        // [esp+44h] [ebp-28h]
	int v136;            // [esp+70h] [ebp+4h]
	unsigned char* v137; // [esp+70h] [ebp+4h]
	unsigned char* v138; // [esp+78h] [ebp+Ch]
	unsigned char* v139; // [esp+78h] [ebp+Ch]
	unsigned char* v140; // [esp+78h] [ebp+Ch]
	int v141;            // [esp+78h] [ebp+Ch]
	unsigned char* v142; // [esp+78h] [ebp+Ch]
	unsigned char* v143; // [esp+78h] [ebp+Ch]

	v4 = this;
	v5 = this[2];
	v135[0] = (int)(v4 + 2);
	v135[2] = (int)(v4 + 3);
	v135[4] = a2;
	v135[3] = a2;
	v6 = 0;
	v135[1] = (int)v5;
	v135[5] = 0;
	v135[6] = 0;
	v135[9] = 0;
	v136 = 0;
	v118 = &a3[a4];
	if ((unsigned int)a4 >= 5) {
		v7 = a3;
		if (a3 < a3 + 5) {
			do {
				v8 = *v7++;
				v6 = __ROL4__(v8 ^ v6, 5);
			} while (v7 < a3 + 5);
			v136 = v6;
		}
	}
	v9 = a4;
	v111 = a4;
	if (a4 < 5) {
		goto LABEL_144;
	}
	v10 = v136;
	while (2) {
		v11 = 0;
		v126 = v9 - 5;
		if (v9 - 5 >= 64) {
			v126 = 64;
		}
		v110 = 0;
		v125 = 0;
		while (1) {
			HIWORD(v15) = 0;
			v12 = 0;
			v13 = v4[5];
			v14 = (unsigned int)(214013 * v10 + 2531011) >> 17;
			LOWORD(v15) = v13[v14];
			v16 = &v13[v14];
			v17 = (unsigned int)v4[1] & 0xFFFF;
			v117 = v17;
			*v16 = v17;
			if ((unsigned short)v15 == 0xFFFF || v15 == v17) {
				goto LABEL_62;
			}
			v18 = (unsigned short)(v17 - v15);
			v19 = v18;
			if (v18 >= v111 - v11) {
				v19 = v111 - v11;
			}
			if (v19 >= 521) {
				v18 = 521;
			} else if (v18 >= v111 - v11) {
				v18 = v111 - v11;
			}
			v112 = v18;
			v138 = &a3[v11];
			v119 = 0x10000 - v15;
			if (0x10000 - v15 < v18) {
				v22 = v15;
				if (v15 >= 0x10000) {
				LABEL_33:
					v23 = 0;
					if (v18 + v15 - 0x10000 <= 0) {
					LABEL_23:
						v12 = v18;
						goto LABEL_24;
					}
					while (*((uint8_t*)*v4 + v23) == v138[v119 + v23]) {
						v18 = v112;
						if (++v23 >= v112 + v15 - 0x10000) {
							goto LABEL_23;
						}
					}
					v12 = v119 + v23;
				} else {
					while (*((uint8_t*)*v4 + v22) == v138[v22 - v15]) {
						if (++v22 >= 0x10000) {
							v10 = v136;
							goto LABEL_33;
						}
					}
					v10 = v136;
					v12 = v22 - v15;
				}
			} else {
				v12 = 0;
				if (v18 <= 0) {
					goto LABEL_23;
				}
				while (*((uint8_t*)*v4 + v15 + v12) == a3[v11 + v12]) {
					v18 = v112;
					if (++v12 >= v112) {
						goto LABEL_23;
					}
				}
			}
		LABEL_24:
			if ((unsigned short)(v15 + v12) == v117) {
				v20 = v111 + -v12 - v11;
				if (521 - v12 < v20) {
					v20 = 521 - v12;
				}
				v120 = v20;
				v113 = 0;
				if (v20 > 0) {
					for (i = &a3[v11]; *i == i[v12]; ++i) {
						if (++v113 >= v120) {
							break;
						}
					}
				}
				v12 += v113;
			}
			if (v12 >= 3) {
				v114 = 521 - v12;
				if (521 - v12 >= v11) {
					v114 = v11;
				}
				v24 = (unsigned short)((char*)v4[1] - v15);
				if (v114 >= v24 - v12) {
					v121 = (unsigned short)((char*)v4[1] - v15) - v12;
				} else {
					v121 = v114;
				}
				v25 = 0x10000 - v24;
				if (v121 < 0x10000 - v24) {
					v25 = v114;
					if (v114 < v24 - v12) {
					LABEL_53:
						v26 = 0;
						if (v25 <= 0) {
							goto LABEL_152;
						}
						v122 = v15 - 1;
						v139 = v138 - 1;
						do {
							if (*((uint8_t*)*v4 + v122) != *v139) {
								break;
							}
							++v26;
							--v122;
							--v139;
						} while (v26 < v114);
						if (v26 <= 0) {
						LABEL_152:
							v10 = v136;
						} else {
							v11 -= v26;
							v27 = (int)v4[1] - v26;
							LOWORD(v15) = v15 - v26;
							v12 += v26;
							v4[1] = (void*)v27;
							v28 = &a3[v11];
							LOWORD(v117) = v27;
							v10 = 0;
							if (v28 < v28 + 5) {
								do {
									v29 = *v28++ ^ v10;
									v10 = __ROL4__(v29, 5);
								} while (v28 < &a3[v11 + 5]);
							}
							v136 = v10;
							v125 = 1;
						}
						goto LABEL_62;
					}
					v25 = (unsigned short)((char*)v4[1] - v15) - v12;
				}
				v114 = v25;
				goto LABEL_53;
			}
		LABEL_62:
			v30 = v110;
			if (v110 >= 4) {
				v115 = v11;
				if (v12 <= v110) {
					v4[1] = v133;
					v137 = &a3[v127];
					v55 = __ROL4__(v134 ^ a3[v127 + 5] ^ __ROL4__(a3[v127], 25), 5);
					sub_57E4C0((uint32_t**)v135, (unsigned int)a3, v127, v110 - 4,
							   (unsigned short)((uint16_t)v133 - v128));
					v56 = &v118[-v127] - a3;
					v57 = v110 - 2;
					if (v110 - 2 >= v56 - 2) {
						v57 = v56 - 2;
					}
					v58 = &a3[v127];
					v59 = (int)v4[1] + 2;
					v141 = (int)v4[1] + 2;
					if (v57 > 0) {
						if (v57 <= 1024) {
							v64 = __ROL4__(v55 ^ v137[6] ^ __ROL4__(v137[1], 25), 5);
							v65 = 0;
							v130 = (int)(v137 + 7);
							while (1) {
								*((uint16_t*)v4[5] + ((unsigned int)(214013 * v64 + 2531011) >> 17)) = v65 + v59;
								v66 = *(unsigned char*)(v130 + v65);
								v67 = __ROL4__(*(unsigned char*)(v130 + v65++ - 5), 25);
								v64 = __ROL4__(v64 ^ v66 ^ v67, 5);
								if (v65 >= v57) {
									break;
								}
								LOWORD(v59) = v141;
							}
							v58 = &a3[v127];
							v136 = v64;
						} else {
							v60 = 0;
							v61 = &v137[v57 + 2];
							v62 = (unsigned int)(v61 + 5);
							if (v61 < v61 + 5) {
								do {
									v63 = *v61++;
									v60 = __ROL4__(v63 ^ v60, 5);
								} while ((unsigned int)v61 < v62);
							}
							v136 = v60;
						}
					} else {
						v136 = 0;
					}
					v68 = (unsigned int)v4[1] & 0xFFFF;
					if (v68 + v110 <= 0x10000) {
						v69 = v110;
						memcpy((char*)*v4 + v68, v58, v110);
					} else {
						memcpy((char*)*v4 + v68, v58, 0x10000 - v68);
						v69 = v110;
						memcpy(*v4, &v58[0x10000 - v68], v110 - (0x10000 - v68));
					}
					v4[1] = (char*)v4[1] + v69;
					a3 += v69 + v127;
					goto LABEL_143;
				}
				sub_57E4C0((uint32_t**)v135, (unsigned int)a3, v11, v12 - 4, (unsigned short)(v117 - v15));
				v36 = &v118[-v11] - a3 - 1;
				if (v12 - 1 < v36) {
					v36 = v12 - 1;
				}
				v37 = 0;
				v38 = (int)v4[1] + 1;
				v39 = &a3[v11];
				v123 = v36;
				v129 = (int)v4[1] + 1;
				v140 = &a3[v11];
				v40 = &a3[v11];
				if (v36 > 0) {
					if (v36 <= 1024) {
						v44 = __ROL4__(v136 ^ v40[5] ^ __ROL4__(*v40, 25), 5);
						v45 = (int)(v40 + 1);
						if (v123 > 0) {
							while (1) {
								*((uint16_t*)v4[5] + ((unsigned int)(214013 * v44 + 2531011) >> 17)) = v37 + v38;
								v46 =
									v44 ^ *(unsigned char*)(v45 + 5 + v37) ^ __ROL4__(*(unsigned char*)(v45 + v37), 25);
								++v37;
								v44 = __ROL4__(v46, 5);
								if (v37 >= v123) {
									break;
								}
								LOWORD(v38) = v129;
							}
						}
						v136 = v44;
					} else {
						v41 = &v40[v36 + 1];
						v42 = (unsigned int)(v41 + 5);
						if (v41 < v41 + 5) {
							do {
								v43 = *v41++;
								v37 = __ROL4__(v43 ^ v37, 5);
							} while ((unsigned int)v41 < v42);
						}
						v136 = v37;
					}
					v39 = v140;
				} else {
					v136 = 0;
				}
				v47 = (unsigned int)v4[1] & 0xFFFF;
				if (v47 + v12 > 0x10000) {
					v48 = 0x10000 - v47;
					memcpy((char*)*v4 + v47, v39, 0x10000 - v47);
					v49 = (char*)*v4;
					v50 = &v140[v48];
					v51 = v12 - v48;
					memcpy(*v4, v50, 4 * (v51 >> 2));
					v53 = &v50[4 * (v51 >> 2)];
					v52 = &v49[4 * (v51 >> 2)];
					v54 = v51;
				LABEL_118:
					memcpy(v52, v53, v54 & 3);
					v4[1] = (char*)v4[1] + v12;
					a3 += v115 + v12;
					goto LABEL_143;
				}
			LABEL_117:
				v84 = (char*)*v4 + v47;
				memcpy(v84, v39, 4 * ((unsigned int)v12 >> 2));
				v53 = &v39[4 * ((unsigned int)v12 >> 2)];
				v52 = &v84[4 * ((unsigned int)v12 >> 2)];
				v54 = v12;
				goto LABEL_118;
			}
			if (v12 < 4) {
				goto LABEL_66;
			}
			if (v125) {
				v115 = v11;
				sub_57E4C0((uint32_t**)v135, (unsigned int)a3, v11, v12 - 4, (unsigned short)(v117 - v15));
				v70 = &v118[-v11] - a3 - 1;
				if (v12 - 1 < v70) {
					v70 = v12 - 1;
				}
				v71 = 0;
				v72 = (int)v4[1] + 1;
				v39 = &a3[v11];
				v124 = v70;
				v131 = (int)v4[1] + 1;
				v142 = &a3[v11];
				v73 = &a3[v11];
				if (v70 > 0) {
					if (v70 <= 1024) {
						v77 = __ROL4__(v136 ^ v73[5] ^ __ROL4__(*v73, 25), 5);
						v78 = (int)(v73 + 1);
						if (v124 > 0) {
							while (1) {
								*((uint16_t*)v4[5] + ((unsigned int)(214013 * v77 + 2531011) >> 17)) = v71 + v72;
								v79 =
									v77 ^ *(unsigned char*)(v78 + 5 + v71) ^ __ROL4__(*(unsigned char*)(v78 + v71), 25);
								++v71;
								v77 = __ROL4__(v79, 5);
								if (v71 >= v124) {
									break;
								}
								LOWORD(v72) = v131;
							}
						}
						v136 = v77;
					} else {
						v74 = &v73[v70 + 1];
						v75 = (unsigned int)(v74 + 5);
						if (v74 < v74 + 5) {
							do {
								v76 = *v74++;
								v71 = __ROL4__(v76 ^ v71, 5);
							} while ((unsigned int)v74 < v75);
						}
						v136 = v71;
					}
					v39 = v142;
				} else {
					v136 = 0;
				}
				v47 = (unsigned int)v4[1] & 0xFFFF;
				if (v47 + v12 > 0x10000) {
					v80 = 0x10000 - v47;
					memcpy((char*)*v4 + v47, v39, 0x10000 - v47);
					v81 = (char*)*v4;
					v82 = &v142[v80];
					v83 = v12 - v80;
					memcpy(*v4, v82, 4 * (v83 >> 2));
					v53 = &v82[4 * (v83 >> 2)];
					v52 = &v81[4 * (v83 >> 2)];
					v54 = v83;
					goto LABEL_118;
				}
				goto LABEL_117;
			}
			v30 = v12;
			v110 = v12;
			v128 = v15;
			v127 = v11;
			v133 = v4[1];
			v134 = v10;
		LABEL_66:
			if (v11 + 1 > v126) {
				break;
			}
			v31 = a3[v11];
			v32 = a3[v11 + 5];
			v33 = a3[v11++];
			v34 = v32 ^ __ROL4__(v33, 25);
			v35 = (char*)v4[1];
			v10 = __ROL4__(v10 ^ v34, 5);
			v4[1] = v35 + 1;
			v136 = v10;
			*((uint8_t*)*v4 + (unsigned short)v35) = v31;
		}
		v116 = v11;
		if (v30 < 4) {
			if (v11 + 5 >= v111 && v111 <= 64) {
				v100 = v111 - v11;
				v101 = &a3[v11];
				v102 = (unsigned int)v4[1] & 0xFFFF;
				if (v102 + v100 <= 0x10000) {
					v105 = v100;
					v104 = (unsigned char*)*v4 + v102;
				} else {
					v103 = 0x10000 - v102;
					memcpy((char*)*v4 + v102, v101, 0x10000 - v102);
					v104 = (unsigned char*)*v4;
					v105 = v100 - v103;
					v101 += v103;
				}
				memcpy(v104, v101, v105);
				v4[1] = (char*)v4[1] + v100;
				v11 = v111;
			}
			sub_57E3F0((uint32_t**)v135, (unsigned int)a3, v11);
			a3 += v11;
		} else {
			sub_57E4C0((uint32_t**)v135, (unsigned int)a3, v11, v110 - 4, (unsigned short)((char*)v4[1] - v128));
			v85 = v110 - 1;
			if (v110 - 1 >= &v118[-v11] - a3 - 1) {
				v85 = &v118[-v11] - a3 - 1;
			}
			v86 = &a3[v11];
			v87 = 0;
			v88 = (int)v4[1] + 1;
			v132 = (int)v4[1] + 1;
			v143 = &a3[v11];
			v89 = &a3[v11];
			if (v85 > 0) {
				if (v85 <= 1024) {
					v93 = v136 ^ v89[5] ^ __ROL4__(*v86, 25);
					v94 = (int)(v89 + 1);
					v95 = __ROL4__(v93, 5);
					while (1) {
						*((uint16_t*)v4[5] + ((unsigned int)(214013 * v95 + 2531011) >> 17)) = v87 + v88;
						v96 = *(unsigned char*)(v94 + 5 + v87);
						v97 = __ROL4__(*(unsigned char*)(v94 + v87++), 25);
						v95 = __ROL4__(v95 ^ v96 ^ v97, 5);
						if (v87 >= v85) {
							break;
						}
						LOWORD(v88) = v132;
					}
					v86 = v143;
					v136 = v95;
				} else {
					v90 = &v86[v85 + 1];
					v91 = (unsigned int)(v90 + 5);
					if (v90 < v90 + 5) {
						do {
							v92 = *v90++;
							v87 = __ROL4__(v92 ^ v87, 5);
						} while ((unsigned int)v90 < v91);
					}
					v136 = v87;
				}
			} else {
				v136 = 0;
			}
			v98 = (unsigned int)v4[1] & 0xFFFF;
			if (v98 + v110 <= 0x10000) {
				v99 = v110;
				memcpy((char*)*v4 + v98, v86, v110);
			} else {
				memcpy((char*)*v4 + v98, v86, 0x10000 - v98);
				v99 = v110;
				memcpy(*v4, &v86[0x10000 - v98], v110 - (0x10000 - v98));
			}
			v4[1] = (char*)v4[1] + v99;
			a3 += v116 + v99;
		}
	LABEL_143:
		v9 = v118 - a3;
		v111 = v118 - a3;
		if (v118 - a3 >= 5) {
			v10 = v136;
			continue;
		}
		break;
	}
LABEL_144:
	if (v111) {
		v106 = (unsigned int)v4[1] & 0xFFFF;
		if (v106 + v111 <= 0x10000) {
			v108 = v111;
			memcpy((char*)*v4 + v106, a3, v111);
			v107 = a3;
		} else {
			v107 = a3;
			memcpy((char*)*v4 + v106, a3, 0x10000 - v106);
			v108 = v111;
			memcpy(*v4, &a3[0x10000 - v106], v111 - (0x10000 - v106));
		}
		v4[1] = (char*)v4[1] + v108;
		sub_57E3F0((uint32_t**)v135, (unsigned int)v107, v108);
	}
	return sub_57E7D0((uint32_t**)v135);
}
