#include <float.h>
#include <math.h>
#include <time.h>

#include "GAME1.h"
#include "GAME1_1.h"
#include "GAME1_2.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_2.h"
#include "GAME2_3.h"
#include "GAME3_1.h"
#include "GAME3_2.h"
#include "GAME3_3.h"
#include "GAME4.h"
#include "GAME4_1.h"
#include "GAME4_2.h"
#include "GAME4_3.h"
#include "GAME5.h"
#include "GAME5_2.h"
#include "client__system__parsecmd.h"
#include "common__binfile.h"
#include "common__random.h"
#include "common__strman.h"

#include "client__gui__guiinv.h"
#include "client__shell__mainmenu.h"
#include "common__system__team.h"

#include "MixPatch.h"
#include "client__video__draw_common.h"
#include "common/fs/nox_fs.h"
#include "common__magic__speltree.h"
#include "operators.h"

extern uint32_t dword_5d4594_2516352;
extern uint32_t dword_5d4594_2516380;
extern uint32_t dword_5d4594_2523804;
extern uint32_t dword_5d4594_2516372;
extern uint32_t dword_5d4594_2523764;
extern uint32_t dword_5d4594_2523776;
extern uint32_t dword_5d4594_2516356;
extern uint64_t qword_581450_9544;
extern uint32_t dword_5d4594_2523780;
extern uint32_t dword_5d4594_2516344;
extern uint32_t dword_5d4594_2516328;
extern uint32_t dword_5d4594_2516348;
extern uint32_t dword_5d4594_2650652;
extern uint32_t dword_8531A0_2576;

//----- (00554040) --------------------------------------------------------
unsigned int nox_server_makeServerInfoPacket_554040(const char* inBuf, int inSz, char* out) {
	char buf[72];

	char* v3 = sub_416640();
	char* game = nox_xxx_cliGamedataGet_416590(0);
	if (!sub_43AF30() || sub_4D6F30()) {
		return 0;
	}
	char playerLimit = nox_xxx_servGetPlrLimit_409FA0();
	char playerCount = nox_common_playerInfoCount_416F40();
	if (nox_common_getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
		--playerCount;
		--playerLimit;
	}
	char* srvName = nox_xxx_serverOptionsGetServername_40A4C0();
	buf[0] = 0;
	buf[1] = 0;
	buf[2] = 13;
	buf[3] = playerCount;
	buf[4] = playerLimit;
	buf[5] = v3[101] & 0xF;
	buf[6] = ((unsigned char)v3[101]) >> 4;
	*(uint32_t*)&buf[7] = *((uint32_t*)game + 11);
	strcpy(&buf[10], nox_xxx_mapGetMapName_409B40());
	buf[19] = v3[102] | sub_43BE50_get_video_mode_id();
	buf[20] = v3[100];
	buf[21] = v3[100] & 0x10;
	*(uint32_t*)&buf[24] = *((uint32_t*)v3 + 12);
	unsigned int gameFlags = nox_common_gameFlags_getVal_40A5B0();
	if (nox_xxx_isQuest_4D6F50()) {
		gameFlags = (gameFlags & 0xFFFFFF7Fu) | 0x1000u;
		*(uint16_t*)&buf[68] = nox_game_getQuestStage_4E3CC0();
	}
	*(uint32_t*)&buf[28] = gameFlags;
	*(uint32_t*)&buf[32] = *((uint32_t*)game + 12);
	*(uint16_t*)&buf[36] = *(uint16_t*)(v3 + 105);
	*(uint16_t*)&buf[38] = *(uint16_t*)(v3 + 107);
	*(uint32_t*)&buf[40] = *((uint32_t*)v3 + 11);
	*(uint32_t*)&buf[44] = *(uint32_t*)(&inBuf[8]); // timestamp of the packet
	memcpy(&buf[48], game + 24, 20);
	memcpy(&out[0], buf, 72);
	strcpy(&out[72], srvName);
	return 72 + strlen(srvName) + 1;
}

unsigned int sub_554290() {
	unsigned int v0;     // edi
	int v1;              // ebx
	char* v2;            // esi
	unsigned int v3;     // eax

	v0 = -1;
	v1 = 0;
	v2 = nox_common_playerInfoGetFirst_416EA0();
	if (!v2) {
		return 0;
	}
	do {
		if (v2[2064] != 31 && sub_554240((unsigned char)v2[2064]) > 0) {
			v3 = sub_554240((unsigned char)v2[2064]);
			if (v3 < v0) {
				v0 = v3;
			}
			++v1;
		}
		v2 = nox_common_playerInfoGetNext_416EE0((int)v2);
	} while (v2);
	if (v1) {
		return v0;
	} else {
		return 0;
	}
}

int sub_554300() {
	int v0;     // ebx
	int v1;     // edi
	char* v2;   // esi

	v0 = 0;
	v1 = 0;
	v2 = nox_common_playerInfoGetFirst_416EA0();
	if (!v2) {
		return 0;
	}
	do {
		if (v2[2064] != 31 && (int)sub_554240((unsigned char)v2[2064]) > 0) {
			v0 += sub_554240((unsigned char)v2[2064]);
			++v1;
		}
		v2 = nox_common_playerInfoGetNext_416EE0((int)v2);
	} while (v2);
	if (v1) {
		return v0 / v1;
	} else {
		return 0;
	}
}

//----- (0056F1C0) --------------------------------------------------------
int sub_56F1C0() {
	int v0;     // eax
	int result; // eax

	v0 = time(0);
	sub_56FF00(v0);
	dword_5d4594_2516352 = 0;
	dword_5d4594_2516348 = gameFrame();
	dword_5d4594_2516344 = 0;
	*getMemU16Ptr(0x587000, 311204) = 0;
	dword_5d4594_2516356 = 657757279;
	dword_5d4594_2516348 ^= nox_xxx_protect_56F240();
	dword_5d4594_2516328 = ~dword_5d4594_2516348;
	*getMemU32Ptr(0x5D4594, 2516340) = nox_xxx_protectionCreateInt_56F400(0);
	sub_56F250();
	result = nox_xxx_protectionCreateInt_56F400(1);
	*getMemU32Ptr(0x5D4594, 2516332) = result;
	return result;
}

//----- (0056F240) --------------------------------------------------------
int nox_xxx_protect_56F240() { return sub_56FF80(1, -1); }

//----- (0056F250) --------------------------------------------------------
int sub_56F250() {
	int v0;     // esi
	int result; // eax

	v0 = 7;
	do {
		result = nox_xxx_protectionCreateStructForInt_56F280(*(int*)&dword_5d4594_2516356, 0);
		--v0;
		++dword_5d4594_2516356;
	} while (v0);
	return result;
}

//----- (0056F280) --------------------------------------------------------
int nox_xxx_protectionCreateStructForInt_56F280(int a1, int a2) {
	uint32_t* v2; // eax
	int v3;       // ecx
	int v4;       // ecx

	v2 = calloc(1u, 0x10u);
	if (!v2) {
		return 0;
	}
	v2[3] = 0;
	v2[2] = 0;
	v3 = a1 ^ dword_5d4594_2516348;
	*v2 = a1 ^ dword_5d4594_2516348;
	dword_5d4594_2516328 ^= v3;
	v2[1] = a2;
	v4 = a2 ^ dword_5d4594_2516348;
	v2[1] = a2 ^ dword_5d4594_2516348;
	dword_5d4594_2516328 ^= v4;
	return sub_56F2F0(v2);
}

//----- (0056F2F0) --------------------------------------------------------
int sub_56F2F0(uint32_t* a1) {
	int v1;     // esi
	short v2;   // di
	int result; // eax
	short i;    // ax
	int v5;     // eax

	v1 = dword_5d4594_2516344;
	v2 = 0;
	if (*getMemU16Ptr(0x587000, 311204)) {
		for (i = nox_common_randomInt_415FA0(0, *getMemU16Ptr(0x587000, 311204) - 1); v1; ++v2) {
			if (v2 == i) {
				break;
			}
			v1 = *(uint32_t*)(v1 + 8);
		}
		a1[3] = *(uint32_t*)(v1 + 12);
		if (!dword_5d4594_2516348) {
			nullsub_31(1);
		}
		a1[2] = v1;
		*(uint32_t*)(v1 + 12) = a1;
		v5 = a1[3];
		if (v5) {
			*(uint32_t*)(v5 + 8) = a1;
			++*getMemU16Ptr(0x587000, 311204);
		} else {
			++*getMemU16Ptr(0x587000, 311204);
			dword_5d4594_2516344 = a1;
		}
		result = 1;
	} else {
		++*getMemU16Ptr(0x587000, 311204);
		dword_5d4594_2516352 = a1;
		dword_5d4594_2516344 = a1;
		result = 1;
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(uint32_t);

//----- (0056F3B0) --------------------------------------------------------
uint32_t* sub_56F3B0() {
	uint32_t* result; // eax
	uint32_t* v1;     // esi

	result = *(uint32_t**)&dword_5d4594_2516344;
	if (dword_5d4594_2516344) {
		do {
			v1 = (uint32_t*)result[2];
			free(result);
			result = v1;
		} while (v1);
	}
	dword_5d4594_2516328 = 0;
	*getMemU16Ptr(0x587000, 311204) = 0;
	dword_5d4594_2516348 = 0;
	dword_5d4594_2516352 = 0;
	dword_5d4594_2516344 = 0;
	return result;
}

//----- (0056F400) --------------------------------------------------------
int nox_xxx_protectionCreateInt_56F400(int a1) {
	if (nox_xxx_protectionCreateStructForInt_56F280(dword_5d4594_2516356, a1)) {
		return (dword_5d4594_2516356)++;
	}
	nullsub_31(1);
	return 0;
}
// 560840: using guessed type void  nullsub_31(uint32_t);

//----- (0056F480) --------------------------------------------------------
int nox_xxx_protectionCreateStructForFloat_56F480(int a1, float val) {
	int a2 = *(int*)&val;
	uint32_t* v2; // eax
	int v3;       // ecx
	int v4;       // ecx

	v2 = calloc(1u, 0x10u);
	if (!v2) {
		return 0;
	}
	v2[3] = 0;
	v2[2] = 0;
	v3 = a1 ^ dword_5d4594_2516348;
	*v2 = a1 ^ dword_5d4594_2516348;
	dword_5d4594_2516328 ^= v3;
	v2[1] = a2;
	v4 = a2 ^ dword_5d4594_2516348;
	v2[1] = a2 ^ dword_5d4594_2516348;
	dword_5d4594_2516328 ^= v4;
	return sub_56F2F0(v2);
}

//----- (0056F4F0) --------------------------------------------------------
int sub_56F4F0(int* a1) {
	int result; // eax

	result = sub_56F510(*a1);
	if (result) {
		*a1 = 0;
	}
	return result;
}

//----- (0056F510) --------------------------------------------------------
int sub_56F510(int a1) {
	uint32_t* v1; // eax
	int v2;       // ecx
	int v3;       // ecx
	int v4;       // ecx
	int v5;       // ecx

	v1 = sub_56F590(a1);
	if (!v1) {
		return 0;
	}
	v2 = v1[3];
	if (v2) {
		*(uint32_t*)(v2 + 8) = v1[2];
	} else {
		dword_5d4594_2516344 = v1[2];
	}
	v3 = v1[2];
	if (v3) {
		*(uint32_t*)(v3 + 12) = v1[3];
	} else {
		dword_5d4594_2516352 = v1[3];
	}
	v4 = *v1 ^ dword_5d4594_2516328;
	dword_5d4594_2516328 = v4;
	v5 = v1[1] ^ v4;
	--*getMemU16Ptr(0x587000, 311204);
	dword_5d4594_2516328 = v5;
	free(v1);
	return 1;
}

//----- (0056F590) --------------------------------------------------------
uint32_t* sub_56F590(int a1) {
	uint32_t* result; // eax

	result = *(uint32_t**)&dword_5d4594_2516344;
	if (dword_5d4594_2516344) {
		while (*result != (a1 ^ dword_5d4594_2516348)) {
			result = (uint32_t*)result[2];
			if (!result) {
				nullsub_31(1);
				return 0;
			}
		}
	} else {
		nullsub_31(1);
		return 0;
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(uint32_t);

//----- (0056F5C0) --------------------------------------------------------
int nox_xxx_protectData_56F5C0() {
	int v0;          // ebx
	int v1;          // edi
	int v2;          // ebx
	int v3;          // edi
	unsigned int v4; // eax
	int i;           // ebp
	int v6;          // esi
	int v7;          // eax
	int* v8;         // eax
	int* v9;         // eax
	int v10;         // ecx
	int v11;         // ebp
	bool v12;        // zf
	int result;      // eax
	int* v14;        // [esp-14h] [ebp-14h]

	if (!dword_5d4594_2516348) {
		nullsub_31(1);
	}
	v0 = gameFrame();
	v1 = dword_5d4594_2516348;
	v2 = nox_xxx_protect_56F240() ^ v0;
	v3 = v2 ^ v1;
	dword_5d4594_2516328 = ~v2;
	v4 = *getMemU16Ptr(0x587000, 311204);
	for (i = 0; i < (*getMemU16Ptr(0x587000, 311204) >> 2); v4 = *getMemU16Ptr(0x587000, 311204)) {
		v6 = nox_common_randomInt_415FA0(0, v4 >> 1);
		v7 = nox_common_randomInt_415FA0((*getMemU16Ptr(0x587000, 311204) >> 1) + 1,
										 *getMemU16Ptr(0x587000, 311204) - 1);
		if (v6 != v7) {
			v14 = sub_56F6F0(v7);
			v8 = sub_56F6F0(v6);
			sub_56F720(v8, v14);
		}
		++i;
	}
	v9 = *(int**)&dword_5d4594_2516344;
	dword_5d4594_2516348 = 0;
	if (dword_5d4594_2516344) {
		do {
			v10 = v3 ^ *v9;
			v11 = v3 ^ v9[1];
			*v9 = v10;
			v9[1] = v11;
			dword_5d4594_2516328 ^= v10;
			dword_5d4594_2516328 ^= v9[1];
			v9 = (int*)v9[2];
		} while (v9);
	}
	result = v2 ^ dword_5d4594_2516348;
	v12 = v2 == dword_5d4594_2516348;
	++*getMemU32Ptr(0x5D4594, 2516364);
	dword_5d4594_2516348 ^= v2;
	if (v12) {
		nullsub_31(1);
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(uint32_t);

//----- (0056F6F0) --------------------------------------------------------
uint32_t* sub_56F6F0(int a1) {
	uint32_t* result; // eax
	int v2;           // ecx

	result = *(uint32_t**)&dword_5d4594_2516344;
	v2 = 0;
	if (dword_5d4594_2516344) {
		while (v2 != a1) {
			result = (uint32_t*)result[2];
			++v2;
			if (!result) {
				nullsub_31(1);
				return 0;
			}
		}
	} else {
		nullsub_31(1);
		return 0;
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(uint32_t);

//----- (0056F720) --------------------------------------------------------
void sub_56F720(int* a1, int* a2) {
	int v2; // edx
	int v3; // esi

	if (!a1 || !a2) {
		nullsub_31(1);
		return;
	}
	v2 = *a1;
	v3 = a1[1];
	*a1 = *a2;
	a1[1] = a2[1];
	*a2 = v2;
	a2[1] = v3;
	++*getMemU32Ptr(0x5D4594, 2516360);
	if (!dword_5d4594_2516348) {
		nullsub_31(1);
	}
}
// 560840: using guessed type void  nullsub_31(uint32_t);

//----- (0056F780) --------------------------------------------------------
uint32_t* sub_56F780(int a1, int a2) {
	uint32_t* result; // eax
	int v3;           // ecx

	result = (uint32_t*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = a2 ^ dword_5d4594_2516348;
			result[1] = a2 ^ dword_5d4594_2516348;
			dword_5d4594_2516328 ^= v3;
			result = (uint32_t*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F7D0) --------------------------------------------------------
uint32_t* nox_xxx_playerResetProtectionCRC_56F7D0(int a1, int a2) {
	uint32_t* result; // eax
	int v3;           // ecx

	result = (uint32_t*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = a2 ^ dword_5d4594_2516348;
			result[1] = a2 ^ dword_5d4594_2516348;
			dword_5d4594_2516328 ^= v3;
			result = (uint32_t*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F820) --------------------------------------------------------
uint32_t* sub_56F820(int a1, unsigned char a2) {
	uint32_t* result; // eax
	int v3;           // ecx

	result = (uint32_t*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = dword_5d4594_2516348 ^ a2;
			result[1] = v3;
			dword_5d4594_2516328 ^= v3;
			result = (uint32_t*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F870) --------------------------------------------------------
uint32_t* nox_xxx_protectPlayerHPMana_56F870(int a1, unsigned short a2) {
	uint32_t* result; // eax
	int v3;           // ecx

	result = (uint32_t*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = dword_5d4594_2516348 ^ a2;
			result[1] = v3;
			dword_5d4594_2516328 ^= v3;
			result = (uint32_t*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F8C0) --------------------------------------------------------
uint32_t* sub_56F8C0(int a1, float a2) {
	uint32_t* result; // eax
	uint32_t* v3;     // esi
	int v4;           // eax

	result = (uint32_t*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		v3 = result;
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v4 = dword_5d4594_2516348 ^ (unsigned long long)(long long)a2;
			v3[1] = v4;
			dword_5d4594_2516328 ^= v4;
			result = (uint32_t*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F920) --------------------------------------------------------
uint32_t* sub_56F920(int a1, int a2) {
	uint32_t* result; // eax
	int v3;           // ecx

	result = (uint32_t*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = dword_5d4594_2516348 ^ (a2 + (dword_5d4594_2516348 ^ result[1]));
			result[1] = v3;
			dword_5d4594_2516328 ^= v3;
			result = (uint32_t*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056F9E0) --------------------------------------------------------
uint32_t* nox_xxx_protectMana_56F9E0(int a1, short a2) {
	uint32_t* result; // eax
	int v3;           // ecx

	result = (uint32_t*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v3 = dword_5d4594_2516348 ^ (a2 + (dword_5d4594_2516348 ^ result[1]));
			result[1] = v3;
			dword_5d4594_2516328 ^= v3;
			result = (uint32_t*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056FA40) --------------------------------------------------------
uint32_t* sub_56FA40(int a1, float a2) {
	uint32_t* result; // eax
	uint32_t* v3;     // esi
	int v4;           // eax

	result = (uint32_t*)a1;
	if (a1 >= 657757279) {
		result = sub_56F590(a1);
		v3 = result;
		if (result) {
			dword_5d4594_2516328 ^= result[1];
			v4 = dword_5d4594_2516348 ^
				 (unsigned long long)(long long)((double)(unsigned int)(dword_5d4594_2516348 ^ result[1]) + a2);
			v3[1] = v4;
			dword_5d4594_2516328 ^= v4;
			result = (uint32_t*)nox_xxx_protectData_56F5C0();
		}
	}
	return result;
}

//----- (0056FAC0) --------------------------------------------------------
int nox_xxx_protectionStringCRC_56FAC0(int* a1, unsigned int a2) {
	int* v2;        // ecx
	int result;     // eax
	unsigned int i; // edx
	int v5;         // esi

	v2 = a1;
	result = 0;
	for (i = a2 >> 2; i; --i) {
		v5 = *v2;
		++v2;
		result ^= v5;
	}
	return result;
}

//----- (0056FAE0) --------------------------------------------------------
int nox_xxx_protectionStringCRCLen_56FAE0(int* a1, unsigned int a2) {
	int result; // eax

	result = 0;
	if (a1) {
		result = nox_xxx_protectionStringCRC_56FAC0(a1, a2);
	}
	return result;
}

//----- (0056FB00) --------------------------------------------------------
int sub_56FB00(int* a1, unsigned int a2, int a3) {
	uint32_t* v3; // esi

	if (a3 >= 657757279) {
		v3 = sub_56F590(a3);
		if (v3 && (dword_5d4594_2516348 ^ nox_xxx_protectionStringCRCLen_56FAE0(a1, a2)) == v3[1]) {
			return 1;
		}
		nullsub_31(1);
	}
	return 0;
}
// 560840: using guessed type void  nullsub_31(uint32_t);

//----- (0056FB60) --------------------------------------------------------
int sub_56FB60(nox_object_t* item) {
	int result; // eax
	int v2;     // ebx
	int v3;     // ebx
	int v4;     // ebx
	int* v5;    // edi
	int v6;     // eax
	int* v7;    // eax

	result = 0;
	if (item) {
		v2 = sub_4E4C00(item);
		v3 = (unsigned short)nox_xxx_unitGetHP_4EE780(item) ^ v2;
		v4 = sub_4E4C10(item) ^ v3;
		v5 = (int*)nox_object_getInitData_4E4C30(item);
		v6 = sub_4E4C50(item);
		if (v5 && v6 > 0) {
			v4 ^= nox_xxx_protectionStringCRC_56FAC0(v5, v6);
		}
		v7 = (int*)sub_4E4C80(item);
		if (v7) {
			if (strlen((const char*)v7)) {
				v4 ^= nox_xxx_protectionStringCRC_56FAC0(v7, strlen((const char*)v7));
			}
		}
		result = v4;
	}
	return result;
}

//----- (0056FBF0) --------------------------------------------------------
int nox_xxx_protect_56FBF0(int a1, nox_object_t* item) {
	int result;   // eax
	uint32_t* v3; // eax
	uint32_t* v4; // esi
	int v5;       // ecx

	result = a1;
	if (a1 >= 657757279) {
		v3 = sub_56F590(a1);
		v4 = v3;
		if (v3) {
			dword_5d4594_2516328 ^= v3[1];
			v5 = sub_56FB60(item) ^ v3[1];
			v4[1] = v5;
			result = v5 ^ dword_5d4594_2516328;
			dword_5d4594_2516328 ^= v5;
		} else {
			nullsub_31(1);
		}
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(uint32_t);

//----- (0056FC50) --------------------------------------------------------
int nox_xxx_protect_56FC50(int a1, const nox_object_t* object) {
	int* a2 = object;
	int result;   // eax
	uint32_t* v3; // eax
	uint32_t* v4; // esi
	int v5;       // ecx

	result = a1;
	if (a1 >= 657757279) {
		v3 = sub_56F590(a1);
		v4 = v3;
		if (v3) {
			dword_5d4594_2516328 ^= v3[1];
			v5 = sub_56FB60(a2) ^ v3[1];
			v4[1] = v5;
			result = v5 ^ dword_5d4594_2516328;
			dword_5d4594_2516328 ^= v5;
		} else {
			nullsub_31(1);
		}
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(uint32_t);

//----- (0056FCB0) --------------------------------------------------------
int sub_56FCB0(int a1, int a2) {
	int result; // eax

	result = 0;
	if (a2) {
		result = 1 << (a1 % 32);
	}
	return result;
}

//----- (0056FCE0) --------------------------------------------------------
int nox_xxx_playerAwardSpellProtectionCRC_56FCE0(int a1, int a2, int a3) {
	int result;   // eax
	uint32_t* v4; // eax
	uint32_t* v5; // esi

	result = a1;
	if (a1 >= 657757279) {
		v4 = sub_56F590(a1);
		v5 = v4;
		if (v4) {
			dword_5d4594_2516328 ^= v4[1];
			result = dword_5d4594_2516348 ^ (dword_5d4594_2516348 ^ v4[1] | sub_56FCB0(a2, a3));
			v5[1] = result;
			dword_5d4594_2516328 ^= result;
		} else {
			nullsub_31(1);
		}
	}
	return result;
}
// 560840: using guessed type void  nullsub_31(uint32_t);

//----- (0056FD50) --------------------------------------------------------
int nox_xxx_playerApplyProtectionCRC_56FD50(int a1, void* a2p, int a3) {
	int a2 = a2p;
	int v3;       // ebp
	uint32_t* v4; // eax
	int v5;       // esi
	int* v6;      // edi
	uint32_t* v8; // [esp+8h] [ebp+4h]

	v3 = 0;
	if (a1 >= 657757279) {
		v4 = sub_56F590(a1);
		v8 = v4;
		if (v4) {
			v5 = 1;
			if (a3 > 1) {
				v6 = (int*)(a2 + 4);
				do {
					v3 |= sub_56FCB0(v5++, *v6);
					++v6;
				} while (v5 < a3);
				v4 = v8;
			}
			if ((v3 ^ dword_5d4594_2516348) == v4[1]) {
				return 1;
			}
		}
		nullsub_31(1);
	}
	return 0;
}
// 560840: using guessed type void  nullsub_31(uint32_t);

//----- (0056FE30) --------------------------------------------------------
double nox_xxx_unkDoubleSmth_56FE30() {
	double v0; // st7

	*getMemU64Ptr(0x5D4594, 2516412) = *getMemU64Ptr(0x5D4594, 2516404);
	*getMemU64Ptr(0x5D4594, 2516404) = *getMemU64Ptr(0x5D4594, 2516396);
	*getMemU32Ptr(0x5D4594, 2516396) = *getMemU32Ptr(0x5D4594, 2516388);
	*getMemU32Ptr(0x5D4594, 2516400) = *getMemU32Ptr(0x5D4594, 2516392);
	v0 = *getMemDoublePtr(0x5D4594, 2516388) * *getMemDoublePtr(0x581450, 11376) +
		 *getMemDoublePtr(0x5D4594, 2516404) * *getMemDoublePtr(0x581450, 11368) +
		 *getMemDoublePtr(0x5D4594, 2516412) * *getMemDoublePtr(0x581450, 11360) +
		 *getMemDoublePtr(0x5D4594, 2516412) * *getMemDoublePtr(0x581450, 11352) + *getMemDoublePtr(0x5D4594, 2516420);
	floor(v0);
	*getMemDoublePtr(0x5D4594, 2516388) = v0 - v0;
	*getMemDoublePtr(0x5D4594, 2516420) = v0 * *getMemDoublePtr(0x581450, 11344);
	return *getMemDoublePtr(0x5D4594, 2516388);
}

//----- (0056FF00) --------------------------------------------------------
void sub_56FF00(int a1) {
	int v1;            // eax
	unsigned char* v2; // ecx
	unsigned int v3;   // eax
	int v4;            // esi

	v1 = a1;
	if (!a1) {
		v1 = -1;
	}
	v2 = getMemAt(0x5D4594, 2516388);
	do {
		v2 += 8;
		v3 = (((v1 << 13) ^ (unsigned int)v1) >> 17) ^ (v1 << 13) ^ v1;
		v1 = (32 * v3) ^ v3;
		*((double*)v2 - 1) = (double)(unsigned int)v1 * *getMemDoublePtr(0x581450, 11344);
	} while ((int)v2 < (int)getMemAt(0x5D4594, 2516428));
	v4 = 19;
	do {
		nox_xxx_unkDoubleSmth_56FE30();
		--v4;
	} while (v4);
	dword_5d4594_2516380 = 0;
	*getMemU32Ptr(0x5D4594, 2516376) = 99;
	dword_5d4594_2516372 = 100;
}

//----- (0056FF80) --------------------------------------------------------
int sub_56FF80(int a1, int a2) {
	long long v2; // rax
	int result;   // eax

	*getMemU32Ptr(0x5D4594, 2516376) = a2;
	dword_5d4594_2516380 = a1;
	dword_5d4594_2516372 = a2 - a1 + 1;
	v2 = (long long)(nox_xxx_unkDoubleSmth_56FE30() * (double)*(unsigned int*)&dword_5d4594_2516372);
	if ((unsigned int)v2 < *(int*)&dword_5d4594_2516372) {
		result = dword_5d4594_2516380 + v2;
	} else {
		result = dword_5d4594_2516372 + dword_5d4594_2516380;
	}
	return result;
}

//----- (00578B00) --------------------------------------------------------
unsigned int nox_xxx_netGetUnitCodeCli_578B00(int a1) {
	unsigned int result; // eax

	if (!a1) {
		return 0;
	}
	result = *(uint32_t*)(a1 + 128);
	if (result >= 0x8000) {
		return 0;
	}
	if (*(uint32_t*)(a1 + 112) & 0x20400000) {
		BYTE1(result) |= 0x80u;
	}
	return result;
}

//----- (00578B30) --------------------------------------------------------
int nox_xxx_netClearHighBit_578B30(short a1) { return a1 & 0x7FFF; }

//----- (00578B40) --------------------------------------------------------
int nox_xxx_packetDynamicUnitCode_578B40(int a1) {
	int result; // eax
	int v2;     // eax

	result = a1;
	if ((a1 & 0x8000) == 0x8000) {
		BYTE1(result) &= 0x7Fu;
		v2 = nox_xxx_netGetUnitByExtent_4ED020(result);
		if (v2) {
			result = *(uint32_t*)(v2 + 36);
		} else {
			result = 0;
		}
	}
	return result;
}

//----- (00578B70) --------------------------------------------------------
unsigned int nox_xxx_netTestHighBit_578B70(unsigned int a1) { return (a1 >> 15) & 1; }

//----- (00579870) --------------------------------------------------------
int nox_xxx_waypointNext_579870(int a1) {
	int result; // eax

	if (a1) {
		result = *(uint32_t*)(a1 + 484);
	} else {
		result = 0;
	}
	return result;
}

//----- (005798A0) --------------------------------------------------------
int sub_5798A0(int a1) {
	int result; // eax

	if (a1) {
		result = *(uint32_t*)(a1 + 484);
	} else {
		result = 0;
	}
	return result;
}

//----- (00579E70) --------------------------------------------------------
uint32_t* sub_579E70() {
	uint32_t* result; // eax

	result = calloc(1u, 0x204u);
	if (result) {
		result[120] |= 0x1000000u;
	}
	return result;
}

//----- (00579EE0) --------------------------------------------------------
int sub_579EE0(nox_waypoint_t* a1, unsigned char a2) { return (a2 & *(uint8_t*)((uint32_t)a1 + 477)) != 0; }

//----- (0057A160) --------------------------------------------------------
int nox_xxx_playerCanTalkMB_57A160(int a1) {
	int result; // eax

	if (a1 && nox_common_gameFlags_check_40A5C0(2)) {
		result = (*(uint32_t*)(a1 + 3680) >> 3) & 1;
	} else {
		result = 0;
	}
	return result;
}

//----- (0057A1B0) --------------------------------------------------------
char* sub_57A1B0(short a1) {
	int v1;            // ecx
	unsigned char* v2; // eax

	v1 = 0;
	v2 = getMemAt(0x587000, 312212);
	while ((a1 & 0x17F0) != *(uint32_t*)v2) {
		v2 += 8;
		++v1;
		if ((int)v2 >= (int)getMemAt(0x587000, 312268)) {
			return 0;
		}
	}
	return *(char**)getMemAt(0x587000, 312208 + 8 * v1);
}

//----- (0057A1E0) --------------------------------------------------------
char sub_57A1E0(int* a1, const char* a2, int* a3, char a4, short a5) {
	int* v5;           // esi
	int v6;            // ebx
	unsigned char v7;  // dl
	char* v8;          // edi
	int v9;            // ecx
	int v10;           // eax
	unsigned char v12; // cl
	char result;       // al
	char v14[256];     // [esp+14h] [ebp-200h]
	char v15[256];     // [esp+114h] [ebp-100h]

	v5 = a3;
	v6 = a5 & 0x17F0;
	if (a3) {
		sub_57ADF0(a3);
	}
	a1[6] = -1;
	a1[7] = -1;
	a1[8] = -1;
	a1[9] = -1;
	a1[10] = -1;
	a1[11] = -1;
	a1[12] = -1;
	if (a4 & 3) {
		v15[0] = 0;
		strcpy(v14, "maps\\");
		strncat(v14, (const char*)a1, 8u);
		strcat(v14, "\\");
		if (a4 & 2) {
			strcpy(v15, v14);
			if (a2) {
				strcat(v15, a2);
			} else {
				strcat(v15, "user.rul");
			}
			v10 = sub_57A3F0(v15, (int)a1, (int)a3, v6);
			v5 = a3;
		} else {
			v10 = 0;
		}
		if (a4 & 1 && !v10) {
			strncat(v14, (const char*)a1, 8u);
			strncat(v14, ".rul", 255);
			sub_57A3F0(v14, (int)a1, (int)v5, v6);
		}
	}
	if (dword_5d4594_2650652 && a4 & 4) {
		sub_57A3F0("internet.rul", (int)a1, (int)v5, v6);
	}
	result = a5;
	if (a5 & 0x40) {
		result = sub_453FA0((int)(a1 + 6), 132, 0);
	}
	return result;
}
// 57A1E0: using guessed type char var_100[256];

//----- (0057A3F0) --------------------------------------------------------
int sub_57A3F0(char* a1, int a2, int a3, int a4) {
	FILE* v4;        // eax
	FILE* v5;        // esi
	char* v6;        // eax
	char v8[256];    // [esp+4h] [ebp-300h]
	wchar2_t v9[256]; // [esp+104h] [ebp-200h]

	dword_5d4594_2523764 = 6128;
	v4 = nox_fs_open_text(a1);
	v5 = v4;
	if (!v4) {
		return 0;
	}
	if (!nox_fs_feof(v4)) {
		do {
			memset(v8, 0, sizeof(v8));
			nox_fs_fgets(v5, v8, 256);
			v6 = strchr(v8, 10);
			if (v6) {
				*v6 = 0;
			}
			if (v8[0]) {
				nox_swprintf(v9, L"%S", v8);
				sub_57A4D0(v9, a2, a3, a4);
			}
		} while (!nox_fs_feof(v5));
	}
	nox_fs_close(v5);
	return 1;
}

//----- (0057A4D0) --------------------------------------------------------
void sub_57A4D0(wchar2_t* a1, int a2, int a3, int a4) {
	unsigned char v4;      // bl
	int v5;                // edi
	const wchar2_t* result; // eax
	const wchar2_t* v7;     // esi
	int v8;                // ecx
	const wchar2_t* v9;     // eax
	wchar2_t* v10;          // esi
	unsigned char v11;     // [esp+10h] [ebp-284h]
	int v12[32];           // [esp+14h] [ebp-280h]
	wchar2_t v13[256];      // [esp+94h] [ebp-200h]

	v4 = 0;
	v5 = 0;
	v11 = 0;
	sub_416580();
	result = nox_wcscpy(v13, a1);
	if (true) { // TODO: was if (v13)
		if (v13[0] == 34) {
			result = nox_wcstok(&v13[1], L"\"\n\r");
			v7 = result;
			v5 = 1;
		} else {
			result = nox_wcstok(v13, L" \n\t\r");
			v7 = result;
		}
		if (v7) {
			do {
				v8 = v11;
				v11 = ++v4;
				v12[v8] = (int)v7;
				v9 = &v7[nox_wcslen(v7) + 1];
				if (v5) {
					++v9;
				}
				if (*v9 == 34) {
					result = nox_wcstok((wchar2_t*)v9 + 1, L"\"\n\r");
					v7 = result;
					v5 = 1;
				} else {
					result = nox_wcstok(0, L" \n\t\r");
					v7 = result;
					v5 = 0;
				}
			} while (v7);
			if (v4) {
				result = (const wchar2_t*)sub_57A620(v4, (const wchar2_t**)v12, a2, a4);
				if (!result) {
					if (a3) {
						v10 = (wchar2_t*)calloc(1, 0x20Cu);
						nox_wcscpy(v10 + 6, a1);
						nox_common_list_append_4258E0(a3, v10);
					}
				}
			}
		}
	}
}
// 57A4D0: using guessed type int var_280[32];

//----- (0057A620) --------------------------------------------------------
int sub_415EC0(char* a1);
int sub_415A30(char* a1);
int sub_57A620(unsigned char a1, const wchar2_t** a2, int a3, int a4) {
	const wchar2_t** v4; // ebp
	unsigned char v5;   // dl
	int v6;             // esi
	int v9;             // eax
	int v10;            // ebx
	unsigned char v11;  // cl
	int v12;            // eax
	int v14;            // esi
	int v15;            // ecx
	char v16[100];      // [esp+10h] [ebp-64h]
	unsigned char v17;  // [esp+7Ch] [ebp+8h]
	unsigned char v18;  // [esp+7Ch] [ebp+8h]

	v4 = a2;
	nox_sprintf(v16, "%S", *a2);
	v5 = 0;
	v17 = 0;
	do {
		if (!strcmp(*(const char**)getMemAt(0x587000, 312208 + 8 * v17), v16)) {
			dword_5d4594_2523764 = *getMemU32Ptr(0x587000, 312212 + 8 * v17);
			return a4 == dword_5d4594_2523764;
		}
		v17 = ++v5;
	} while (v5 < 7u);
	if (!(dword_5d4594_2523764 & a4) || _nox_wcsicmp(*v4, L"set") || a1 <= 1u) {
		return 0;
	}
	if (_nox_wcsicmp(v4[1], L"spell")) {
		if (_nox_wcsicmp(v4[1], L"weapon")) {
			if (_nox_wcsicmp(v4[1], L"armor")) {
				return 0;
			}
			if (a1 != 4) {
				return 0;
			}
			if (!nox_common_gameFlags_check_40A5C0(1)) {
				return 0;
			}
			nox_sprintf(v16, "%S", v4[2]);
			int v13 = sub_415EC0(v16);
			if (v13 < 0) {
				return 0;
			}
			v14 = sub_415D10(v13);
			if (!v14) {
				return 0;
			}
			if (_nox_wcsicmp(v4[3], L"off")) {
				v15 = v14 | *(uint32_t*)(a3 + 48);
			} else {
				v15 = ~v14 & *(uint32_t*)(a3 + 48);
			}
			*(uint32_t*)(a3 + 48) = v15;
		} else {
			if (a1 != 4) {
				return 0;
			}
			if (!nox_common_gameFlags_check_40A5C0(1)) {
				return 0;
			}
			nox_sprintf(v16, "%S", v4[2]);
			int v8 = sub_415A30(v16);
			if (v8 < 0) {
				return 0;
			}
			v9 = nox_xxx_ammoCheck_415880(v8);
			v10 = v9;
			if (!v9) {
				return 0;
			}
			v11 = 0;
			v18 = 0;
			if (v9 > 0) {
				do {
					v12 = v10 >> 8;
					if (v10 >> 8 > 0) {
						v10 >>= 8;
					}
					++v11;
				} while (v12 > 0);
				v18 = v11;
			}
			if (_nox_wcsicmp(v4[3], L"off")) {
				*(uint8_t*)(v18 + a3 + 43) |= v10;
			} else {
				*(uint8_t*)(v18 + a3 + 43) &= ~(uint8_t)v10;
			}
		}
	} else {
		if (a1 != 4) {
			return 0;
		}
		nox_sprintf(v16, "%S", v4[2]);
		v6 = nox_xxx_spellNameToN_4243F0(v16);
		if (!v6) {
			v6 = nox_xxx_spellByTitle_424960((wchar2_t*)v4[2]);
			if (!v6) {
				return 0;
			}
		}
		if (!nox_xxx_spellIsValid_424B50(v6)) {
			return 0;
		}
		if (nox_xxx_spellFlags_424A70(v6) & 0x7000000) {
			if (!_nox_wcsicmp(v4[3], L"off")) {
				sub_453FA0(a3 + 24, v6, 0);
			}
		}
	}
	if (a4 == dword_5d4594_2523764) {
		return 1;
	}
	return 0;
}

//----- (0057A950) --------------------------------------------------------
int sub_57A950(char* a1) {
	char v2[256]; // [esp+Ch] [ebp-100h]

	strcpy(v2, "maps\\");
	strncat(v2, a1, 256 - 6);
	v2[strlen(v2) - 4] = 0;
	strncat(v2, "\\", 256 - 4);
	strcat(v2, a1);
	return sub_4D0550(v2);
}

//----- (0057A9F0) --------------------------------------------------------
int sub_57A9F0(const char* a1, const char* a2) {
	char FileName[256]; // [esp+Ch] [ebp-100h]

	strcpy(FileName, "maps\\");
	strcat(FileName, a1);
	strcat(FileName, "\\");
	strcat(FileName, a2);
	return nox_fs_remove(FileName);
}

//----- (0057AAA0) --------------------------------------------------------
char* sub_4159B0(int a1);
char* sub_415E40(char* a1);
char sub_57AAA0(const char* a1, char* a2, int* a3) {
	FILE* v3;           // eax
	FILE* v4;           // ebp
	int* i;             // esi
	char* v6;           // eax
	int v7;             // esi
	int v8;             // edi
	char* v9;           // eax
	int v10;            // esi
	int v11;            // edi
	char* v12;          // eax
	int v13;            // ebx
	int v14;            // esi
	char* v15;          // edi
	char* v16;          // eax
	int v18;            // [esp+10h] [ebp-27Ch]
	// Keep each map header and its rule flags in one contiguous buffer.
	char v19[60];       // [esp+14h] [ebp-278h]
	char v21[60];       // [esp+50h] [ebp-23Ch]
	char v23[256];      // [esp+8Ch] [ebp-200h]
	char FileName[256]; // [esp+18Ch] [ebp-100h]

	LOBYTE(v3) = a2[52];
	if ((char)v3 >= 0) {
		strcpy(FileName, "maps\\");
		strncat(FileName, a2, 8u);
		strncat(FileName, "\\", 10u);
		strcat(FileName, a1);
		v3 = nox_fs_create_text(FileName);
		v4 = v3;
		if (v3) {
			if (dword_5d4594_2650652) {
				strcpy(v21, a2);
				strcpy(v19, a2);
				sub_57A1E0((int*)v21, 0, 0, 4, *((uint16_t*)a2 + 26));
				sub_57A1E0((int*)v19, 0, 0, 3, *((uint16_t*)a2 + 26));
			}
			if (a3) {
				for (i = nox_common_list_getFirstSafe_425890(a3); i; i = nox_common_list_getNextSafe_4258A0(i)) {
					nox_sprintf(v23, "%S\n", i + 3);
					nox_fs_fputs(v4, v23);
				}
			}
			v6 = sub_57A1B0(*((uint16_t*)a2 + 26));
			nox_fs_fputs(v4, v6);
			nox_fs_fputs(v4, "\n");
			v7 = 1;
			v8 = 136;
			do {
				if (nox_xxx_spellIsValid_424B50(v7) && !sub_454000((int)(a2 + 24), v7) &&
					nox_xxx_spellFlags_424A70(v7) & 0x7000000 &&
					(!dword_5d4594_2650652 || sub_454000((int)(v21 + 24), v7) || !sub_454000((int)(v19 + 24), v7))) {
					v9 = nox_xxx_spellNameByN_424870(v7);
					nox_sprintf(v23, "%s %s \"%s\" %s\n", "set", "spell", v9, "off");
					nox_fs_fputs(v4, v23);
				}
				++v7;
				--v8;
			} while (v8);
			v10 = 1;
			v11 = 26;
			do {
				if (!(v10 & *((uint32_t*)a2 + 12))) {
					v12 = sub_415E40((char*)v10);
					if (v12) {
						nox_sprintf(v23, "%s %s \"%s\" %s\n", "set", "Armor", v12, "off");
						nox_fs_fputs(v4, v23);
					}
				}
				v10 *= 2;
				--v11;
			} while (v11);
			v13 = 1;
			v14 = 1;
			v18 = 27;
			v15 = a2 + 44;
			do {
				if (!((unsigned char)v13 & (unsigned char)*v15)) {
					v16 = sub_4159B0((char*)v14);
					if (v16) {
						nox_sprintf(v23, "%s %s \"%s\" %s\n", "set", "weapon", v16, "off");
						nox_fs_fputs(v4, v23);
					}
				}
				if (v13 == 128) {
					v13 = 1;
					++v15;
				} else {
					v13 *= 2;
				}
				v14 *= 2;
				--v18;
			} while (v18);
			LOBYTE(v3) = 0;
			nox_fs_close(v4);
		}
	}
	return (char)v3;
}

//----- (0057ADF0) --------------------------------------------------------
int* sub_57ADF0(int* a1) {
	int* result; // eax
	int* v2;     // esi
	int* v3;     // edi

	result = nox_common_list_getFirstSafe_425890(a1);
	v2 = result;
	if (result) {
		do {
			v3 = nox_common_list_getNextSafe_4258A0(v2);
			nox_common_list_remove_425920((uint32_t**)v2);
			free(v2);
			v2 = v3;
		} while (v3);
	}
	return result;
}

//----- (0057AE30) --------------------------------------------------------
int sub_57AE30(const char* a1) {
	int v1;          // ebp
	const char** v2; // edi

	v1 = 0;
	v2 = (const char**)getMemAt(0x587000, 312208);
	while (strcmp(*v2, a1)) {
		v2 += 2;
		++v1;
		if ((int)v2 >= (int)getMemAt(0x587000, 312264)) {
			return 0;
		}
	}
	return *getMemU32Ptr(0x587000, 312212 + 8 * v1);
}

//----- (0057AEA0) --------------------------------------------------------
int nox_xxx_playerCheckSpellClass_57AEA0(int a1, int a2) {
	int v2;     // eax
	int result; // eax
	int v4;     // ecx

	v2 = nox_xxx_spellFlags_424A70(a2);
	if (a1 == 1) {
		v4 = 0x2000000;
	} else {
		if (a1 != 2) {
			return 9;
		}
		v4 = 0x4000000;
	}
	if (v2 & 0x1000000 || v4 & v2) {
		result = 0;
	} else {
		result = 9;
	}
	return result;
}

//----- (0057AF20) --------------------------------------------------------
int nox_xxx_get_57AF20() { return dword_5d4594_2523804; }

//----- (0057B0A0) --------------------------------------------------------
void sub_57B0A0() {
	int result;   // eax
	uint32_t* v1; // ecx

	result = dword_5d4594_2523804;
	if (!result) {
		return;
	}
	v1 = *(uint32_t**)&dword_5d4594_2523780;
	if (dword_5d4594_2523780 && (!*getMemU32Ptr(0x5D4594, 2523772) || *getMemU32Ptr(0x5D4594, 2523772) == 1)) {
		nox_xxx_netSendPointFx_522FF0(154, (float2*)(dword_5d4594_2523780 + 56));
		v1 = *(uint32_t**)&dword_5d4594_2523780;
	}
	if (dword_5d4594_2523776) {
		nox_xxx_delayedDeleteObject_4E5CC0(*(int*)&dword_5d4594_2523776);
		v1 = *(uint32_t**)&dword_5d4594_2523780;
	}
	dword_5d4594_2523776 = 0;
	if (v1) {
		nox_xxx_playerSetState_4FA020(v1, 13);
	}
	dword_5d4594_2523780 = 0;
	if (!sub_45D9B0()) {
		sub_413A00(0);
	}
	dword_5d4594_2523804 = 0;
}

//----- (0057B180) --------------------------------------------------------
long long nox_xxx___Getcvt_57B180() { return *getMemU64Ptr(0x5D4594, 2523788); }

//----- (0057B190) --------------------------------------------------------
int sub_57B190(unsigned short a1, unsigned short a2) {
	int result; // eax
	double v3;  // st7
	double v4;  // st6

	if (!a2) {
		return 4;
	}
	if (a1 == a2) {
		return 0;
	}
	v3 = (double)a1;
	v4 = (double)a2;
	if (v3 >= v4 * *(double*)&qword_581450_9544) {
		return 1;
	}
	result = 2;
	if (v3 < v4 * *getMemDoublePtr(0x581450, 9608)) {
		result = 3;
	}
	return result;
}

//----- (0057B3D0) --------------------------------------------------------
int nox_cheat_allowall = 0;

//----- (0057B400) --------------------------------------------------------
int nox_xxx_client_57B400(int a1) {
	int v1; // eax

	v1 = *getMemU32Ptr(0x5D4594, 2523876);
	if (!*getMemU32Ptr(0x5D4594, 2523876)) {
		v1 = nox_xxx_getTTByNameSpriteMB_44CFC0("Glyph");
		*getMemU32Ptr(0x5D4594, 2523876) = v1;
	}
	if (!dword_8531A0_2576) {
		return 0;
	}
	if (*(uint32_t*)(a1 + 108) != v1 || *(uint8_t*)(dword_8531A0_2576 + 2251) == 1) {
		return 1;
	}
	return 0;
}

//----- (0057B450) --------------------------------------------------------
int sub_57B450(nox_drawable* a1p) {
	int* a1 = a1p;
	int v1;           // eax
	unsigned char v2; // bl

	v1 = *getMemU32Ptr(0x5D4594, 2523880);
	if (!*getMemU32Ptr(0x5D4594, 2523880)) {
		v1 = nox_xxx_getTTByNameSpriteMB_44CFC0("Glyph");
		*getMemU32Ptr(0x5D4594, 2523880) = v1;
	}
	if (!a1 || !*getMemU32Ptr(0x852978, 8) || !dword_8531A0_2576 ||
		a1[27] == v1 && *(uint8_t*)(dword_8531A0_2576 + 2251) != 1) {
		return 0;
	}
	if (nox_cheat_allowall) {
		return 1;
	}
	v2 = 1 << *(uint8_t*)(dword_8531A0_2576 + 2251);
	return (v2 & (unsigned char)sub_57B370(a1[28], a1[29], a1[27])) != 0;
}

//----- (0057B810) --------------------------------------------------------
int nox_xxx_collideReflect_57B810(float* a1, int a2) {
	int result; // eax
	double v3;  // st7
	int v4;     // ecx

	result = a2;
	v3 = *(float*)a2;
	if (a1[1] * *a1 <= 0.0) {
		v4 = *(uint32_t*)(a2 + 4);
		*(float*)(a2 + 4) = *(float*)a2;
		*(uint32_t*)a2 = v4;
	} else {
		*(float*)a2 = -*(float*)(a2 + 4);
		*(float*)(a2 + 4) = -v3;
	}
	return result;
}

//----- (0057B850) --------------------------------------------------------
int nox_xxx_map_57B850(float2* a1, float* a2, float2* a3) {
	int result; // eax
	float v4;   // [esp+0h] [ebp-10h]
	float v5;   // [esp+4h] [ebp-Ch]
	float v6;   // [esp+8h] [ebp-8h]
	float v7;   // [esp+Ch] [ebp-4h]

	v4 = a2[5] + a1->field_0;
	v5 = a2[6] + a1->field_4;
	result = 0;
	if ((v5 - v4 + a3->field_0 - a3->field_4) * 0.70709997 < 0.0 &&
		(a2[8] + a1->field_4 - (a2[7] + a1->field_0) + a3->field_0 - a3->field_4) * 0.70709997 > 0.0) {
		v6 = a2[9] + a1->field_0;
		v7 = a2[10] + a1->field_4;
		if ((v7 + v6 - a3->field_0 - a3->field_4) * 0.70709997 > 0.0 &&
			(v5 + v4 - a3->field_0 - a3->field_4) * 0.70709997 < 0.0) {
			result = 1;
		}
	}
	return result;
}

//----- (0057B920) --------------------------------------------------------
int sub_57B920(void* a1) {
	int result; // eax

	result = 0;
	memset(a1, 0, 0x7F8u);
	return result;
}

//----- (0057B9A0) --------------------------------------------------------
char nox_xxx_cliGenerateAlias_57B9A0(int a1, int a2, int a3, unsigned int a4) {
	int v4; // eax
	int v5; // edx

	v4 = (unsigned char)a2;
	v5 = (unsigned char)a2;
	if ((unsigned char)a2 == 255 || !(uint8_t)a2) {
		v4 = 1;
		v5 = 1;
	}
	while ((*(unsigned short*)(a1 + 8 * v4) != a2 || *(unsigned short*)(a1 + 8 * v4 + 2) != a3) &&
		   *(uint32_t*)(a1 + 8 * v4 + 4) >= a4) {
		if (++v4 == 255) {
			v4 = 1;
		}
		if (v4 == v5) {
			LOBYTE(v4) = -1;
			return v4;
		}
	}
	return v4;
}

//----- (0057BA10) --------------------------------------------------------
int sub_57BA10(int a1, short a2, short a3, int a4) {
	int result; // eax

	result = a1;
	*(uint16_t*)a1 = a2;
	*(uint16_t*)(a1 + 2) = a3;
	*(uint32_t*)(a1 + 4) = a4;
	return result;
}

//----- (0057C090) --------------------------------------------------------
int nox_server_getNextMapGroup_57C090(int a1) {
	int result; // eax

	if (a1) {
		result = *(uint32_t*)(a1 + 88);
	} else {
		result = 0;
	}
	return result;
}

//----- (0057C790) --------------------------------------------------------
void sub_57C790(float4* a1, float2* a2, float2* a3, float a4) {
	double v4; // st5
	double v5; // st7
	double v6; // st6
	double v7; // st7
	float v8;  // [esp+0h] [ebp-10h]
	float v9;  // [esp+4h] [ebp-Ch]
	float v10; // [esp+8h] [ebp-8h]
	float v11; // [esp+14h] [ebp+4h]
	float v12; // [esp+18h] [ebp+8h]
	float v13; // [esp+20h] [ebp+10h]

	v8 = a1->field_8 - a1->field_0;
	v4 = a1->field_C - a1->field_4;
	v9 = v4;
	v5 = v4 * (a2->field_4 - a1->field_4) + v8 * (a2->field_0 - a1->field_0);
	v6 = a4 * a4;
	v10 = v5 * v8 / v6;
	a3->field_0 = v10 + a1->field_0;
	a3->field_4 = v5 * v9 / v6 + a1->field_4;
	if (a1->field_0 >= (double)a1->field_8) {
		v12 = a1->field_0;
		v13 = a1->field_8;
	} else {
		v13 = a1->field_0;
		v12 = a1->field_8;
	}
	if (a1->field_4 >= (double)a1->field_C) {
		v7 = a1->field_C;
		v11 = a1->field_4;
	} else {
		v7 = a1->field_4;
		v11 = a1->field_C;
	}
	if (a3->field_0 >= (double)v13) {
		if (a3->field_0 > (double)v12) {
			a3->field_0 = v12;
		}
	} else {
		a3->field_0 = v13;
	}
	if (v7 <= a3->field_4) {
		if (a3->field_4 > (double)v11) {
			a3->field_4 = v11;
		}
	} else {
		a3->field_4 = v7;
	}
}

//----- (0057C8A0) --------------------------------------------------------
int nox_xxx_mathPointOnTheLine_57C8A0(float4* a1, float2* a2, float2* a3) {
	float4* v3; // ecx
	float2* v4; // edx
	double v5;  // st7
	double v6;  // st6
	double v7;  // st6
	double v8;  // st7
	float v10;  // [esp+0h] [ebp-Ch]
	float v11;  // [esp+0h] [ebp-Ch]
	float v12;  // [esp+8h] [ebp-4h]
	float v13;  // [esp+8h] [ebp-4h]
	float v14;  // [esp+10h] [ebp+4h]
	float v15;  // [esp+10h] [ebp+4h]
	float v16;  // [esp+14h] [ebp+8h]
	float v17;  // [esp+18h] [ebp+Ch]

	v3 = a1;
	v4 = a3;
	v5 = a1->field_8 - a1->field_0;
	v6 = a1->field_C - a1->field_4;
	v12 = v6;
	v10 = v6 * v12 + v5 * v5;
	v7 = (a2->field_4 - a1->field_4) * v12 + (a2->field_0 - a1->field_0) * v5;
	v14 = v7;
	v13 = v14 * v12 / v10;
	a3->field_0 = v5 * v7 / v10 + v3->field_0;
	v15 = v13 + v3->field_4;
	a3->field_4 = v15;
	if (v3->field_0 >= (double)v3->field_8) {
		v8 = v3->field_8;
		v16 = v3->field_0;
	} else {
		v8 = v3->field_0;
		v16 = v3->field_8;
	}
	if (v3->field_4 >= (double)v3->field_C) {
		v11 = v3->field_4;
		v17 = v3->field_C;
	} else {
		v17 = v3->field_4;
		v11 = v3->field_C;
	}
	return v8 <= v4->field_0 && v4->field_0 <= (double)v16 && v15 >= (double)v17 && v15 <= (double)v11;
}

//----- (0057CDB0) --------------------------------------------------------
int sub_57F2A0(float2* a1, int a2, int a3);
char sub_57F1D0(float2* a1);
int sub_57CDB0(int2* a1, float* a2, float2* a3) {
	int2* v3;   // esi
	char v4;    // bl
	int result; // eax
	float2* v6; // eax
	float2* v7; // eax
	float2* v8; // eax
	char v9;    // [esp+10h] [ebp+4h]

	v3 = a1;
	v9 = sub_57F2A0((float2*)a2, a1->field_0, a1->field_4);
	v4 = sub_57F1D0((float2*)a2 + 1);
	switch (sub_57B500(v3->field_0, v3->field_4, 64)) {
	case 0:
		if (v9 != 1 && v9) {
			v8 = a3;
			v8->field_0 = 0.70709997;
			v8->field_4 = 0.70709997;
			return 1;
		}
		a3->field_0 = -0.70709997;
		a3->field_4 = -0.70709997;
		return 1;
	case 1:
		if (v9 == 1 || v9 == 2) {
			a3->field_0 = 0.70709997;
			a3->field_4 = -0.70709997;
			return 1;
		}
		a3->field_0 = -0.70709997;
		a3->field_4 = 0.70709997;
		return 1;
	case 2:
		switch (v9) {
		case 0:
			v8 = a3;
			a3->field_0 = -0.70709997;
			if (!(v4 & 2)) {
				a3->field_4 = -0.70709997;
				return 1;
			}
			v8->field_4 = 0.70709997;
			return 1;
		case 1:
			v6 = a3;
			if (!(v4 & 1)) {
				v6->field_0 = 0.70709997;
				v6->field_4 = -0.70709997;
				return 1;
			}
			a3->field_0 = -0.70709997;
			a3->field_4 = -0.70709997;
			result = 1;
			break;
		case 2:
			v8 = a3;
			a3->field_0 = 0.70709997;
			if (v4 & 1) {
				v8->field_4 = 0.70709997;
				return 1;
			}
			a3->field_4 = -0.70709997;
			return 1;
		case 3:
			v8 = a3;
			if (!(v4 & 4)) {
				v8->field_0 = 0.70709997;
				v8->field_4 = 0.70709997;
				return 1;
			}
			a3->field_0 = -0.70709997;
			a3->field_4 = 0.70709997;
			return 1;
		default:
			return 1;
		}
		return result;
	case 3:
		if (!v9) {
			v8 = a3;
			a3->field_0 = -0.70709997;
			if (!(v4 & 2)) {
				a3->field_4 = -0.70709997;
				return 1;
			}
			v8->field_4 = 0.70709997;
			return 1;
		}
		if (v9 != 1) {
			v8 = a3;
			v8->field_0 = 0.70709997;
			v8->field_4 = 0.70709997;
			return 1;
		}
		v7 = a3;
		if (!(v4 & 1)) {
			v7->field_0 = 0.70709997;
			v7->field_4 = -0.70709997;
			return 1;
		}
		v7->field_0 = -0.70709997;
		v7->field_4 = -0.70709997;
		return 1;
	case 4:
		if (v9 == 1) {
			v7 = a3;
			if (!(v4 & 1)) {
				v7->field_0 = 0.70709997;
				v7->field_4 = -0.70709997;
				return 1;
			}
			v7->field_0 = -0.70709997;
			v7->field_4 = -0.70709997;
			return 1;
		}
		if (v9 == 2) {
			v8 = a3;
			a3->field_0 = 0.70709997;
			if (v4 & 1) {
				v8->field_4 = 0.70709997;
				return 1;
			}
			a3->field_4 = -0.70709997;
			return 1;
		}
		a3->field_0 = -0.70709997;
		a3->field_4 = 0.70709997;
		return 1;
	case 5:
		if (v9 == 2) {
			v8 = a3;
			a3->field_0 = 0.70709997;
			if (v4 & 1) {
				v8->field_4 = 0.70709997;
				return 1;
			}
			a3->field_4 = -0.70709997;
			return 1;
		}
		if (v9 == 3) {
			v8 = a3;
			if (!(v4 & 4)) {
				v8->field_0 = 0.70709997;
				v8->field_4 = 0.70709997;
				return 1;
			}
			a3->field_0 = -0.70709997;
			a3->field_4 = 0.70709997;
			return 1;
		}
		a3->field_0 = -0.70709997;
		a3->field_4 = -0.70709997;
		return 1;
	case 6:
		if (!v9) {
			v8 = a3;
			a3->field_0 = -0.70709997;
			if (!(v4 & 2)) {
				a3->field_4 = -0.70709997;
				return 1;
			}
			v8->field_4 = 0.70709997;
			return 1;
		}
		if (v9 == 3) {
			v8 = a3;
			if (!(v4 & 4)) {
				v8->field_0 = 0.70709997;
				v8->field_4 = 0.70709997;
				return 1;
			}
			a3->field_0 = -0.70709997;
			a3->field_4 = 0.70709997;
			return 1;
		}
		v6 = a3;
		v6->field_0 = 0.70709997;
		v6->field_4 = -0.70709997;
		return 1;
	case 7:
		if (v9 == 1) {
			v7 = a3;
			if (!(v4 & 1)) {
				v7->field_0 = 0.70709997;
				v7->field_4 = -0.70709997;
				return 1;
			}
			v7->field_0 = -0.70709997;
			v7->field_4 = -0.70709997;
			return 1;
		}
		v8 = a3;
		if (v4 & 1) {
			v8->field_0 = 0.70709997;
			v8->field_4 = 0.70709997;
			return 1;
		}
		a3->field_0 = -0.70709997;
		a3->field_4 = 0.70709997;
		return 1;
	case 8:
		if (v9 == 2) {
			v8 = a3;
			a3->field_0 = 0.70709997;
			if (v4 & 1) {
				v8->field_4 = 0.70709997;
				return 1;
			}
			a3->field_4 = -0.70709997;
			return 1;
		} else {
			v8 = a3;
			a3->field_0 = -0.70709997;
			if (!(v4 & 1)) {
				v8->field_4 = 0.70709997;
				return 1;
			}
			a3->field_4 = -0.70709997;
			return 1;
		}
	case 9:
		if (v9 == 3) {
			v8 = a3;
			if (!(v4 & 4)) {
				v8->field_0 = 0.70709997;
				v8->field_4 = 0.70709997;
				return 1;
			}
			a3->field_0 = -0.70709997;
			a3->field_4 = 0.70709997;
			return 1;
		} else {
			v7 = a3;
			if (v4 & 4) {
				v7->field_0 = 0.70709997;
				v7->field_4 = -0.70709997;
				return 1;
			} else {
				v7->field_0 = -0.70709997;
				v7->field_4 = -0.70709997;
				return 1;
			}
		}
	case 0xA:
		if (v9) {
			if (v4 & 2) {
				a3->field_0 = 0.70709997;
				a3->field_4 = -0.70709997;
				return 1;
			}
			v8 = a3;
			v8->field_0 = 0.70709997;
			v8->field_4 = 0.70709997;
			return 1;
		} else {
			v8 = a3;
			a3->field_0 = -0.70709997;
			if (!(v4 & 2)) {
				a3->field_4 = -0.70709997;
				return 1;
			}
			v8->field_4 = 0.70709997;
			return 1;
		}
	default:
		return 0;
	}
}

//----- (0057F1D0) --------------------------------------------------------
char sub_57F1D0(float2* a1) {
	char v1;          // bl
	int v2;           // edi
	double v3;        // st7
	double v4;        // st6
	double v5;        // st7
	unsigned char v7; // [esp+Ch] [ebp-4h]

	v1 = 0;
	v2 = nox_float2int(a1->field_0);
	v7 = nox_float2int(a1->field_4) % 23;
	v3 = (double)(unsigned char)(v2 % 23);
	if (v3 >= 11.5) {
		v4 = (double)v7;
		if (v4 >= 11.5) {
			v1 = 4;
		}
		if (v4 <= 11.5) {
			v1 |= 1u;
		}
	}
	if (v3 <= 11.5) {
		v5 = (double)v7;
		if (v5 >= 11.5) {
			v1 |= 8u;
		}
		if (v5 <= 11.5) {
			v1 |= 2u;
		}
	}
	return v1;
}

//----- (0057F2A0) --------------------------------------------------------
int sub_57F2A0(float2* a1, int a2, int a3) {
	int v3;     // esi
	int v4;     // eax
	int result; // eax
	int v6;     // eax
	float v7;   // [esp+0h] [ebp-Ch]
	float v8;   // [esp+0h] [ebp-Ch]

	v7 = a1->field_0 - (double)(23 * a2);
	v3 = nox_float2int(v7);
	v8 = a1->field_4 - (double)(23 * a3);
	v4 = nox_float2int(v8);
	if (v3 <= v4) {
		LOBYTE(v4) = 22 - v3 <= v4;
		v6 = v4 - 1;
		LOBYTE(v6) = v6 & 0xFD;
		result = v6 + 3;
	} else {
		LOBYTE(v4) = 22 - v3 <= v4;
		result = v4 + 1;
	}
	return result;
}

void nullsub_10(uint32_t a1) {}

int sub_42CC50(void** this) { return sub_42C770(this); }

int nox_xxx_j_inventoryNameSignInit_467460(void) { return nox_xxx_inventoryNameSignInit_4671E0(); }

int nullsub_8(int a1, int a2) { return 0; }
void nullsub_28(uint32_t a1) {}
void nullsub_30(uint32_t a1) {}
void nullsub_29(void) {}
void nullsub_35(uint32_t a1, uint32_t a2) {}
void nullsub_24(uint32_t a1) {}
void nullsub_31(uint32_t a1) {}
void nullsub_9(uint32_t a1) {}
void nox_xxx_j_allocHitArray_511840(void) { nox_xxx_allocHitArray_5486D0(); }
