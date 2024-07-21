#include <math.h>

#include "GAME1.h"
#include "GAME1_1.h"
#include "GAME1_2.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME3_2.h"
#include "GAME3_3.h"
#include "GAME5.h"
#include "GAME5_2.h"
#include "client__system__parsecmd.h"
#include "common__log.h"
#include "common__object__armrlook.h"
#include "common__object__weaplook.h"
#include "common__system__team.h"

#include "client__gui__servopts__playrlst.h"
#include "client__system__ctrlevnt.h"
#include "client__video__draw_common.h"
#include "common__binfile.h"
#include "common__magic__speltree.h"
#include "common__net_list.h"
#include "common__random.h"
#include "common__strman.h"

#include "client__gui__window.h"
#include "client__io__win95__focus.h"
#include "common/fs/nox_fs.h"
#include "input_common.h"
#include "operators.h"

extern uint32_t dword_5d4594_527656;
extern uint32_t dword_5d4594_526276;
extern uint32_t dword_5d4594_251712;
extern uint32_t dword_5d4594_251708;
extern uint32_t dword_5d4594_251716;
extern uint32_t dword_5d4594_10984;
extern uint32_t dword_5d4594_251720;
extern uint32_t dword_5d4594_251744;
extern uint32_t dword_5d4594_3484;
extern void* dword_5d4594_251560;
extern uint32_t dword_5d4594_371692;
extern uint32_t nox_tile_def_cnt;
extern uint32_t dword_5d4594_251572;
extern uint32_t nox_player_netCode_85319C;

int nox_server_gameSettingsUpdated; // If you define it as 1-byte bool, the game will crash

extern obj_5D4594_2650668_t** ptr_5D4594_2650668;
extern int ptr_5D4594_2650668_cap;


uint32_t nox_tile_def_cnt = 0;
nox_tileDef_t nox_tile_defs_arr[176] = {0};

//----- (00409470) --------------------------------------------------------
int nox_xxx_parseString_409470(FILE* a1, uint8_t* a2) {
	uint8_t* v2;          // ebx
	int v3;               // ecx
	int v4;               // edi
	int v5;               // ebp
	int v6;               // eax
	uint16_t CharType[2]; // [esp+10h] [ebp-4h]

	v2 = a2;
	v3 = 0;
	*(uint32_t*)CharType = 0;
	v4 = 1;
	do {
		while (1) {
			v5 = v3;
			nox_binfile_fread_408E40((char*)CharType, 1, 1, a1);
			if (nox_binfile_lastErr_409370(a1) == -1) {
				return 0;
			}
			v3 = *(uint32_t*)CharType;
			v6 = isspace(CharType[0]);
			if (v6) {
				break;
			}
			v4 = 0;
			if (v3 != '/' || v5 != '/') {
				*v2++ = v3;
			} else {
				nox_binfile_skipLine_409520(a1);
				v2 = a2;
				v3 = *(uint32_t*)CharType;
				v4 = 1;
			}
		}
	} while (v4);
	*v2 = 0;
	return 1;
}

//----- (00409A70) --------------------------------------------------------
int sub_409A70(short a1) {
	int result;        // eax
	unsigned char* v2; // ecx

	result = 0;
	v2 = getMemAt(0x587000, 4704);
	while (*(uint32_t*)v2 != (a1 & 0x17F0)) {
		v2 += 4;
		++result;
		if ((int)v2 >= (int)getMemAt(0x587000, 4728)) {
			return 0;
		}
	}
	return result;
}

//----- (00409B30) --------------------------------------------------------
char* nox_server_currentMapGetFilename_409B30() { return (char*)getMemAt(0x5D4594, 2598188); }

//----- (00409B40) --------------------------------------------------------
char* nox_xxx_mapGetMapName_409B40() { return (char*)getMemAt(0x85B3FC, 36); }

//----- (00409B50) --------------------------------------------------------
unsigned int sub_409B50(const char* a1) {
	unsigned int result; // eax

	result = strlen(a1) + 1;
	memcpy(getMemAt(0x5D4594, 3452), a1, result);
	return result;
}

//----- (00409B80) --------------------------------------------------------
char* sub_409B80() { return (char*)getMemAt(0x5D4594, 3452); }

//----- (00409D70) --------------------------------------------------------
char* nox_xxx_gameSetMapPath_409D70(char* a1) {
	char* result;  // eax
	char* v2;      // eax
	signed int v3; // esi

	result = (char*)nox_strcmpi((const char*)getMemAt(0x5D4594, 2598188), a1);
	if (result) {
		strncpy((char*)getMemAt(0x5D4594, 2598188), a1, 0x50u);
		*getMemU8Ptr(0x5D4594, 2598267) = 0;
		v2 = strrchr(a1, 92);
		if (v2) {
			v3 = strlen(v2 + 1) - 4;
			result = strncpy((char*)getMemAt(0x85B3FC, 36), v2 + 1, v3);
		} else {
			v3 = strlen((const char*)getMemAt(0x5D4594, 2598188)) - 4;
			if (v3 < 0) {
				v3 = 0;
			}
			result = strncpy((char*)getMemAt(0x85B3FC, 36), (const char*)getMemAt(0x5D4594, 2598188), v3);
		}
		*getMemU8Ptr(0x85B3FC, 36 + v3) = 0;
		nox_server_gameSettingsUpdated = 1;
	}
	return result;
}

//----- (00409E40) --------------------------------------------------------
int sub_409E40(int a1) {
	int result; // eax

	result = a1;
	if (dword_5d4594_3484 != a1) {
		dword_5d4594_3484 = a1;
		nox_server_gameSettingsUpdated = 1;
	}
	return result;
}

//----- (00409E60) --------------------------------------------------------
int nox_xxx_getServerSubFlags_409E60() { return dword_5d4594_3484; }

//----- (00409E70) --------------------------------------------------------
int sub_409E70(int a1) {
	int result; // eax

	result = a1 & dword_5d4594_3484;
	if ((a1 & dword_5d4594_3484) != a1) {
		dword_5d4594_3484 |= a1;
		result = nox_common_gameFlags_check_40A5C0(1);
		if (result) {
			if (a1 & 0x2000) {
				result = sub_4D7EA0();
			}
		}
		nox_server_gameSettingsUpdated = 1;
	}
	return result;
}

//----- (00409EC0) --------------------------------------------------------
int sub_409EC0(int a1) {
	int result; // eax

	result = a1;
	if (dword_5d4594_3484 & a1) {
		result = ~a1;
		nox_server_gameSettingsUpdated = 1;
		dword_5d4594_3484 &= ~a1;
	}
	return result;
}

//----- (00409EF0) --------------------------------------------------------
int sub_409EF0(int a1) {
	int result; // eax

	dword_5d4594_3484 ^= a1;
	result = nox_common_gameFlags_check_40A5C0(1);
	if (result) {
		if (a1 & 0x2000) {
			result = dword_5d4594_3484;
			if (result & 0x2000) {
				result = sub_4D7EA0();
			}
		}
	}
	nox_server_gameSettingsUpdated = 1;
	return result;
}

//----- (00409F40) --------------------------------------------------------
int sub_409F40(int a1) {
	int result; // eax

	if (a1 == 0x2000 && nox_common_gameFlags_check_40A5C0(1056)) {
		result = 1;
	} else {
		result = (a1 & dword_5d4594_3484) != 0;
	}
	return result;
}

//----- (00409F80) --------------------------------------------------------
int nox_xxx_servSetPlrLimit_409F80(int a1) {
	int result; // eax

	result = a1;
	if (*getMemU32Ptr(0x5D4594, 3464) != a1) {
		*getMemU32Ptr(0x5D4594, 3464) = a1;
		nox_server_gameSettingsUpdated = 1;
	}
	return result;
}

//----- (00409FA0) --------------------------------------------------------
int nox_xxx_servGetPlrLimit_409FA0() { return *getMemU32Ptr(0x5D4594, 3464); }

//----- (0040A020) --------------------------------------------------------
short nox_xxx_servGamedataGet_40A020(short a1) { return *getMemU16Ptr(0x5D4594, 3488 + 2 * sub_409A70(a1)); }

//----- (0040A180) --------------------------------------------------------
unsigned char sub_40A180(short a1) { return getMemByte(0x5D4594, 3500 + sub_409A70(a1)); }

//----- (0040A1F0) --------------------------------------------------------
int sub_40A1F0(int a1) {
	int result; // eax

	*getMemU32Ptr(0x587000, 4660) = a1;
	result = nox_common_gameFlags_check_40A5C0(1);
	if (result) {
		result = nox_xxx_netTimerStatus_4D8F50(159, a1);
	}
	return result;
}

//----- (0040A220) --------------------------------------------------------
int sub_40A220() { return *getMemU32Ptr(0x587000, 4660); }

//----- (0040A230) --------------------------------------------------------
int sub_40A230() { return *getMemU32Ptr(0x5D4594, 3468) - (unsigned long long)nox_platform_get_ticks(); }

//----- (0040A250) --------------------------------------------------------
long long sub_40A250() {
	short v0;         // ax
	int v1;           // esi
	long long v2;     // kr00_8
	long long result; // rax

	v0 = nox_common_gameFlags_getVal_40A5B0();
	v1 = sub_409A70(v0);
	v2 = nox_platform_get_ticks();
	result = 60000 * getMemByte(0x5D4594, 3500 + v1);
	*getMemU64Ptr(0x5D4594, 3468) = result + v2;
	return result;
}

//----- (0040A310) --------------------------------------------------------
long long sub_40A310(int a1) {
	long long v1;     // kr00_8
	long long result; // rax

	v1 = nox_platform_get_ticks();
	result = a1;
	*getMemU64Ptr(0x5D4594, 3468) = a1 + v1;
	return result;
}

//----- (0040A340) --------------------------------------------------------
void nox_xxx_set3512_40A340(int a1) { *getMemU32Ptr(0x5D4594, 3512) = a1; }

//----- (0040A350) --------------------------------------------------------
int nox_xxx_get3512_40A350() { return *getMemU32Ptr(0x5D4594, 3512); }

//----- (0040A3C0) --------------------------------------------------------
unsigned int sub_40A3C0(unsigned int a1) {
	unsigned int result; // eax

	result = a1;
	if (a1 > 0x800 && a1 < 0x8000) {
		*getMemU32Ptr(0x587000, 4652) = a1;
	}
	return result;
}

//----- (0040A440) --------------------------------------------------------
char* nox_xxx_gameSetServername_40A440(char* a1) {
	char* result; // eax
	char v2[16];  // [esp+0h] [ebp-10h]

	result = a1;
	if (a1) {
		strncpy(v2, a1, 0xFu);
		v2[15] = 0;
		result = (char*)nox_strcmpi((const char*)getMemAt(0x5D4594, 1324), v2);
		if (result) {
			result = strncpy((char*)getMemAt(0x5D4594, 1324), v2, 0xFu);
			*getMemU8Ptr(0x5D4594, 1339) = 0;
			nox_server_gameSettingsUpdated = 1;
		}
	} else {
		*getMemU8Ptr(0x5D4594, 1324) = 0;
		nox_server_gameSettingsUpdated = 1;
	}
	return result;
}

//----- (0040A4C0) --------------------------------------------------------
char* nox_xxx_serverOptionsGetServername_40A4C0() { return (char*)getMemAt(0x5D4594, 1324); }

//----- (0040A5E0) --------------------------------------------------------
int nox_xxx_ruleSetNoRespawn_40A5E0(int a1) {
	int result; // eax

	result = a1;
	*getMemU32Ptr(0x5D4594, 3584) = a1;
	return result;
}

//----- (0040A5F0) --------------------------------------------------------
int nox_server_doPlayersAutoRespawn_40A5F0() {
	return !nox_common_gameFlags_check_40A5C0(4096) ? *getMemU32Ptr(0x5D4594, 3584) : 0;
}

//----- (0040A610) --------------------------------------------------------
wchar2_t* nox_xxx_sysopSetPass_40A610(wchar2_t* a1) { return nox_wcscpy((wchar2_t*)getMemAt(0x5D4594, 3540), a1); }

//----- (0040A630) --------------------------------------------------------
wchar2_t* nox_xxx_sysopGetPass_40A630() { return (wchar2_t*)getMemAt(0x5D4594, 3540); }

//----- (0040A670) --------------------------------------------------------
void nox_server_gameSettingsUpdated_40A670() { nox_server_gameSettingsUpdated = 1; }

//----- (0040A680) --------------------------------------------------------
int nox_server_gameDoSwitchMap_40A680() { return nox_server_gameSettingsUpdated; }

//----- (0040A690) --------------------------------------------------------
void nox_server_gameUnsetMapLoad_40A690() { nox_server_gameSettingsUpdated = 0; }

//----- (0040A6A0) --------------------------------------------------------
int sub_40A6A0(int a1) {
	int result; // eax

	result = a1;
	*getMemU32Ptr(0x5D4594, 3588) = a1;
	return result;
}

//----- (0040A6B0) --------------------------------------------------------
int sub_40A6B0() { return *getMemU32Ptr(0x5D4594, 3588); }

//----- (0040A6C0) --------------------------------------------------------
int nox_xxx_rateGet_40A6C0() { return *getMemU32Ptr(0x587000, 4728); }

//----- (0040A6D0) --------------------------------------------------------
int nox_xxx_rateUpdate_40A6D0(int a1) {
	int result; // eax

	result = *getMemU32Ptr(0x587000, 4728);
	if (a1 != *getMemU32Ptr(0x587000, 4728)) {
		result = nox_common_gameFlags_check_40A5C0(0x20000);
		if (result == 1) {
			result = nox_xxx_netNotifyRate_4D7F10(159);
		}
	}
	*getMemU32Ptr(0x587000, 4728) = a1;
	return result;
}

//----- (0040A710) --------------------------------------------------------
int sub_40A710(int a1) {
	int v1;            // ecx
	unsigned char* v2; // eax

	v1 = 0;
	v2 = getMemAt(0x587000, 4664);
	while (*(uint32_t*)v2 != a1) {
		v2 += 8;
		++v1;
		if ((int)v2 >= (int)getMemAt(0x587000, 4704)) {
			return 1;
		}
	}
	return *getMemU32Ptr(0x587000, 4668 + 8 * v1);
}

//----- (0040A740) --------------------------------------------------------
int sub_40A740() {
	char* v0; // esi

	v0 = sub_4165B0();
	return nox_common_gameFlags_check_40A5C0(128) && v0[53] < 0;
}

//----- (0040A770) --------------------------------------------------------
int sub_40A770() {
	int v0;   // edi
	char* v1; // esi
	int v2;   // eax
	int v3;   // ecx
	char* i;  // eax

	v0 = 0;
	if (!nox_xxx_CheckGameplayFlags_417DA0(4)) {
		for (i = nox_common_playerInfoGetFirst_416EA0(); i; i = nox_common_playerInfoGetNext_416EE0((int)i)) {
			if (!(i[3680] & 1) &&
				(i[2064] != 31 || !nox_common_getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING))) {
				++v0;
			}
		}
		return v0;
	}
	v1 = nox_server_teamFirst_418B10();
	if (!v1) {
		return v0;
	}
	do {
		v2 = nox_xxx_getFirstPlayerUnit_4DA7C0();
		if (v2) {
			while (1) {
				v3 = *(uint32_t*)(*(uint32_t*)(v2 + 748) + 276);
				if (!(*(uint8_t*)(v3 + 3680) & 1) &&
					(*(uint8_t*)(v3 + 2064) != 31 ||
					 !nox_common_getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING))) {
					break;
				}
				v2 = nox_xxx_getNextPlayerUnit_4DA7F0(v2);
				if (!v2) {
					goto LABEL_10;
				}
			}
			++v0;
		}
	LABEL_10:
		v1 = nox_server_teamNext_418B60((int)v1);
	} while (v1);
	return v0;
}

//----- (0040A830) --------------------------------------------------------
int nox_xxx_countNonEliminatedPlayersInTeam_40A830(nox_team_t* a1p) {
	int a1 = a1p;
	int v1; // edi
	int v2; // esi
	int v3; // eax

	v1 = 0;
	v2 = nox_xxx_getFirstPlayerUnit_4DA7C0();
	if (!v2) {
		return 0;
	}
	do {
		if (nox_xxx_teamCompare2_419180(v2 + 48, *(uint8_t*)(a1 + 57))) {
			v3 = *(uint32_t*)(*(uint32_t*)(v2 + 748) + 276);
			if (!(*(uint8_t*)(v3 + 3680) & 1) &&
				(*(uint8_t*)(v3 + 2064) != 31 ||
				 !nox_common_getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING))) {
				++v1;
			}
		}
		v2 = nox_xxx_getNextPlayerUnit_4DA7F0(v2);
	} while (v2);
	return v1;
}

//----- (0040A8A0) --------------------------------------------------------
int nox_xxx_gamePlayIsAnyPlayers_40A8A0() {
	int v0;     // ebx
	char* v1;   // edi
	int v2;     // esi
	int v3;     // eax
	int result; // eax
	char* i;    // eax
	int v6;     // ecx

	v0 = 0;
	if (!nox_xxx_CheckGameplayFlags_417DA0(4)) {
		for (i = nox_common_playerInfoGetFirst_416EA0(); i; i = nox_common_playerInfoGetNext_416EE0((int)i)) {
			if (*((uint32_t*)i + 514)) {
				v6 = *((uint32_t*)i + 920);
				if (!(v6 & 1) || v6 & 0x20) {
					++v0;
				}
			}
		}
		result = v0 < 1;
		LOBYTE(result) = v0 > 1;
		return result;
	}
	v1 = nox_server_teamFirst_418B10();
	if (!v1) {
		result = v0 < 1;
		LOBYTE(result) = v0 > 1;
		return result;
	}
	do {
		v2 = nox_xxx_getFirstPlayerUnit_4DA7C0();
		if (v2) {
			while (1) {
				if (nox_xxx_teamCompare2_419180(v2 + 48, v1[57])) {
					v3 = *(uint32_t*)(*(uint32_t*)(*(uint32_t*)(v2 + 748) + 276) + 3680);
					if (!(v3 & 1) || v3 & 0x20) {
						break;
					}
				}
				v2 = nox_xxx_getNextPlayerUnit_4DA7F0(v2);
				if (!v2) {
					goto LABEL_10;
				}
			}
			++v0;
		}
	LABEL_10:
		v1 = nox_server_teamNext_418B60((int)v1);
	} while (v1);
	return v0 > 1;
}

//----- (0040A970) --------------------------------------------------------
void sub_40A970() {
	char* i;  // esi
	int v1;   // eax
	float v3; // [esp+0h] [ebp-8h]
	float v4; // [esp+0h] [ebp-8h]

	*getMemU32Ptr(0x5D4594, 3520) = gameFrame();
	*getMemU32Ptr(0x5D4594, 3536) = 0;
	v3 = nox_xxx_gamedataGetFloat_419D40("SuddenDeathPlayerThreshold");
	*getMemU32Ptr(0x5D4594, 3476) = nox_float2int(v3);
	v4 = nox_xxx_gamedataGetFloat_419D40("SuddenDeathLifeTime");
	*getMemU32Ptr(0x5D4594, 1392) = nox_float2int(v4);
	for (i = nox_common_playerInfoGetFirst_416EA0(); i; i = nox_common_playerInfoGetNext_416EE0((int)i)) {
		v1 = *((uint32_t*)i + 920);
		if (v1 & 0x100) {
			nox_xxx_playerUnsetStatus_417530((int)i, 256);
		}
	}
	nox_common_gameFlags_unset_40A540(0x4000000);
}

//----- (0040AA00) --------------------------------------------------------
int sub_40AA00() { return 20 * gameFPS() < (unsigned int)(gameFrame() - *getMemU32Ptr(0x5D4594, 3520)); }

//----- (0040AA20) --------------------------------------------------------
int sub_40AA20() { return *getMemU32Ptr(0x5D4594, 3536); }

//----- (0040AA30) --------------------------------------------------------
int sub_40AA30(int a1) {
	int result; // eax

	result = a1;
	*getMemU32Ptr(0x5D4594, 3536) = a1;
	return result;
}

//----- (0040AA40) --------------------------------------------------------
int sub_40AA40() { return *getMemU32Ptr(0x5D4594, 3476); }

//----- (0040AA60) --------------------------------------------------------
int sub_40AA60(int a1) {
	int result; // eax

	result = a1;
	*getMemU32Ptr(0x5D4594, 3508) = a1;
	return result;
}

//----- (0040AA70) --------------------------------------------------------
int sub_40AA70(nox_playerInfo* pl) {
	int a1 = pl;
	char* v1;   // edi
	int result; // eax
	int v3;     // eax
	int v4;     // esi
	int v5;     // ebx
	char* v6;   // eax

	v1 = sub_416640();
	if (!a1) {
		goto LABEL_31;
	}
	if (nox_common_gameFlags_check_40A5C0(4096)) {
		return sub_40A770() < 6;
	}
	v3 = *(uint32_t*)(a1 + 3680);
	if (v3 & 0x100 && !*getMemU32Ptr(0x5D4594, 3508)) {
		return 0;
	}
	if (!sub_40A740() && !nox_common_gameFlags_check_40A5C0(0x8000)) {
		goto LABEL_31;
	}
	result = *(uint32_t*)(a1 + 2068);
	if (!result) {
		return result;
	}
	if (nox_server_teamByXxx_418AE0(*(uint32_t*)(a1 + 2068))) {
		goto LABEL_31;
	}
	v4 = (unsigned char)v1[52];
	if ((nox_common_gameFlags_check_40A5C0(96) ||
		 nox_common_gameFlags_check_40A5C0(16) && nox_xxx_CheckGameplayFlags_417DA0(4)) &&
		v4 > 2) {
		v4 = 2;
	}
	if ((unsigned char)sub_417DE0() >= v4) {
		return 0;
	}
	if (nox_common_gameFlags_check_40A5C0(96)) {
		v5 = (unsigned char)sub_417DE0();
		if (v5 >= sub_417DC0()) {
			return 0;
		}
	}
LABEL_31:
	if (nox_common_gameFlags_check_40A5C0(128)) {
		return 1;
	}
	if (!nox_common_gameFlags_check_40A5C0(1024)) {
		return 1;
	}
	v6 = nox_common_playerInfoGetFirst_416EA0();
	if (!v6) {
		return 1;
	}
	while (*((int*)v6 + 535) <= 0) {
		v6 = nox_common_playerInfoGetNext_416EE0((int)v6);
		if (!v6) {
			return 1;
		}
	}
	if (!sub_40AA00()) {
		return 1;
	}
	return 0;
}

//----- (0040E090) --------------------------------------------------------
void sub_40E090() { dword_5d4594_10984 = 0; }

//----- (004100F0) --------------------------------------------------------
int nox_xxx_cliCanTalkMB_4100F0(short* a1) {
	short v1;  // ax
	short* v2; // ecx

	v1 = *a1;
	v2 = a1 + 1;
	if (!*a1) {
		return 1;
	}
	while (!HIBYTE(v1)) {
		v1 = *v2;
		++v2;
		if (!v1) {
			return 1;
		}
	}
	return 0;
}

//----- (00410360) --------------------------------------------------------
uint8_t* nox_xxx_doorAttachWall_410360(int a1, int a2, int a3) {
	uint8_t* result; // eax
	char v4;         // cl

	result = nox_xxx_wallCreateAt_410250(a2, a3);
	if (result) {
		v4 = result[4] | 0x10;
		*((uint32_t*)result + 7) = a1;
		result[4] = v4;
	}
	return result;
}

//----- (00410390) --------------------------------------------------------
uint32_t* sub_410390(int a1, int a2, int a3) {
	uint32_t* v3;     // esi
	uint32_t* result; // eax
	int v5;           // edx
	char v6;          // cl
	int v7[2];        // [esp+Ch] [ebp-8h]

	v3 = (uint32_t*)nox_xxx_wall_4105E0(a2, a3);
	if (v3 || (result = nox_xxx_wallCreateAt_410250(a2, a3), (v3 = result) != 0)) {
		v5 = *(uint32_t*)(a1 + 16);
		v7[0] = *(uint32_t*)(a1 + 12);
		v6 = *((uint8_t*)v3 + 4);
		v3[8] = a1;
		v7[1] = v5;
		*((uint8_t*)v3 + 4) = v6 | 0x10;
		result = nox_xxx_polygonIsPlayerInPolygon_4217B0((int2*)v7, 0);
		if (result || (result = sub_421990((int2*)v7, 10.0, 0)) != 0) {
			*((uint8_t*)v3 + 8) = *((uint8_t*)result + 130);
		} else {
			*((uint8_t*)v3 + 8) = 1;
		}
	}
	return result;
}

//----- (00410550) --------------------------------------------------------
int sub_410550(short a1) {
	int* v1; // eax

	v1 = (int*)nox_xxx_wallSecretGetFirstWall_410780();
	if (!v1) {
		return 0;
	}
	while (*(uint16_t*)(v1[3] + 10) != a1) {
		v1 = (int*)nox_xxx_wallSecretNext_410790(v1);
		if (!v1) {
			return 0;
		}
	}
	return v1[3];
}

//----- (00410730) --------------------------------------------------------
uint32_t* sub_410730() {
	uint32_t* result; // eax
	uint32_t* v1;     // esi

	result = dword_5d4594_251560;
	if (dword_5d4594_251560) {
		do {
			v1 = (uint32_t*)*result;
			free(result);
			result = v1;
		} while (v1);
		dword_5d4594_251560 = 0;
	} else {
		dword_5d4594_251560 = 0;
	}
	return result;
}

//----- (00410760) --------------------------------------------------------
uint32_t* nox_xxx_wallSecretBlock_410760(uint32_t* a1) {
	uint32_t* result; // eax

	result = a1;
	*a1 = dword_5d4594_251560;
	dword_5d4594_251560 = a1;
	return result;
}

//----- (00410780) --------------------------------------------------------
void* nox_xxx_wallSecretGetFirstWall_410780() { return dword_5d4594_251560; }

//----- (00410790) --------------------------------------------------------
int nox_xxx_wallSecretNext_410790(int* a1) {
	int result; // eax

	if (a1) {
		result = *a1;
	} else {
		result = 0;
	}
	return result;
}

//----- (004107A0) --------------------------------------------------------
int* sub_4107A0(void* lpMem) {
	int* result; // eax
	int* v2;     // esi

	result = dword_5d4594_251560;
	v2 = 0;
	if (dword_5d4594_251560) {
		while (result != lpMem) {
			v2 = result;
			result = (int*)nox_xxx_wallSecretNext_410790(result);
			if (!result) {
				return result;
			}
		}
		if (result == dword_5d4594_251560) {
			dword_5d4594_251560 = nox_xxx_wallSecretNext_410790(result);
		} else {
			*v2 = nox_xxx_wallSecretNext_410790(result);
		}
		free(lpMem);
	}
	return result;
}

//----- (00410F60) --------------------------------------------------------
int nox_xxx_tileAlloc_410F60_init() {
	ptr_5D4594_2650668 = calloc(ptr_5D4594_2650668_cap, sizeof(void*));
	if (!ptr_5D4594_2650668) {
		return 0;
	}

	for (int i = 0; i < ptr_5D4594_2650668_cap; i++) {
		ptr_5D4594_2650668[i] = (obj_5D4594_2650668_t*)calloc(ptr_5D4594_2650668_cap, sizeof(obj_5D4594_2650668_t));
		if (!ptr_5D4594_2650668[i]) {
			return 0;
		}
	}
	return 1;
}

//----- (00410FC0) --------------------------------------------------------
void nox_xxx_tileFree_410FC0_free() {
	for (int i = 0; i < ptr_5D4594_2650668_cap; i++) {
		obj_5D4594_2650668_t* ptr = ptr_5D4594_2650668[i];
		if (ptr) {
			free(ptr);
		}
	}
}

//----- (00411160) --------------------------------------------------------
int nox_xxx_tileNFromPoint_411160(float2* a1) {
	float v12 = (a1->field_0 + 11.5) * 0.021739131;
	float v13 = (a1->field_4 + 11.5) * 0.021739131;

	int i = nox_float2int(v12);
	int j = nox_float2int(v13);

	float v14 = a1->field_0 + 11.5;
	float v15 = a1->field_4 + 11.5;

	int v4 = nox_float2int(v14) % 46;
	int v5 = nox_float2int(v15) % 46;

	if (i - 1 <= 0 || i >= 127 || j - 1 <= 0 || j >= 127) {
		return -1;
	}

	int result = 0;
	int v16[2] = {0};
	if (v4 <= v5) {
		if (46 - v4 <= v5) {
			obj_5D4594_2650668_t* obj = &ptr_5D4594_2650668[i][j];
			result = obj->field_6;
			if (obj->field_10) {
				v16[0] = v4;
				v16[1] = v5 - 23;
				result = sub_411350(obj->field_10, v16, result);
			}
		} else {
			obj_5D4594_2650668_t* obj = &ptr_5D4594_2650668[i - 1][j];
			result = obj->field_1;
			if (obj->field_5) {
				v16[1] = v5;
				v16[0] = v4 + 23;
				result = sub_411350(obj->field_5, v16, result);
			}
		}
	} else if (46 - v4 <= v5) {
		obj_5D4594_2650668_t* obj = &ptr_5D4594_2650668[i][j];
		result = obj->field_1;
		if (obj->field_5) {
			v16[1] = v5;
			v16[0] = v4 - 23;
			result = sub_411350(obj->field_5, v16, result);
		}
	} else {
		obj_5D4594_2650668_t* obj = &ptr_5D4594_2650668[i][j - 1];
		result = obj->field_6;
		if (obj->field_10) {
			v16[0] = v4;
			v16[1] = v5 + 23;
			result = sub_411350(obj->field_10, v16, result);
		}
	}
	return result;
}

//----- (00411350) --------------------------------------------------------
int sub_411350(int* a1, int* a2, int a3) {
	int* v3; // esi
	int v4;  // ebx
	int v5;  // eax

	v3 = a1;
	if (!a1) {
		return a3;
	}
	v4 = a3;
	do {
		v5 = sub_411490(v3[2], v3[3]);
		if (sub_4113A0(a2, v5)) {
			v4 = *v3;
		}
		v3 = (int*)v3[4];
	} while (v3);
	return v4;
}

//----- (004113A0) --------------------------------------------------------
int sub_4113A0(int* a1, int a2) {
	int v2 = *a1;
	int v3 = a1[1];

	bool v4;
	bool v5;
	if (*a1 == v3) {
		v4 = 1;
		v5 = 1;
	} else if (*a1 >= v3) {
		v4 = 1;
		v5 = 0;
	} else {
		v4 = 0;
		v5 = 1;
	}

	bool v6;
	bool v7;
	if (v2 == 46 - v3) {
		v6 = 1;
		v7 = 1;
	} else if (v2 >= 46 - v3) {
		v6 = 1;
		v7 = 0;
	} else {
		v6 = 0;
		v7 = 1;
	}

	switch (a2) {
	case 0:
		if (!v5 || !v6) {
			return 0;
		}
		return 1;
	case 1:
		if (!v5) {
			return 0;
		}
		return 1;
	case 2:
		if (!v5 || !v7) {
			return 0;
		}
		return 1;
	case 3:
		if (!v6) {
			return 0;
		}
		return 1;
	case 4:
		if (!v7) {
			return 0;
		}
		return 1;
	case 5:
		if (!v6 || !v4) {
			return 0;
		}
		return 1;
	case 6:
		if (!v4) {
			return 0;
		}
		return 1;
	case 7:
		if (!v4 || !v7) {
			return 0;
		}
		return 1;
	case 8:
		if (v5) {
			return 1;
		}
		if (!v6) {
			return 0;
		}
		return 1;
	case 9:
		if (!v5 && !v7) {
			return 0;
		}
		return 1;
	case 10:
		if (!v4 && !v7) {
			return 0;
		}
		return 1;
	case 11:
		if (v4) {
			return 1;
		}
		if (!v6) {
			return 0;
		}
		return 1;
	}
	return 0;
}

//----- (00411490) --------------------------------------------------------
int sub_411490(int a1, int a2) {
	int v2; // ecx
	int v3; // edx
	int v5; // esi

	v2 = getMemByte(0x85B3FC, 28696 + 60 * a1);
	v3 = getMemByte(0x85B3FC, 28697 + 60 * a1);
	if (v2 == 3 && v3 == 3) {
		return a2;
	}
	if (!a2) {
		return 0;
	}
	if (a2 <= v2 - 2) {
		return 1;
	}
	if (a2 == v2 - 1) {
		return 2;
	}
	v5 = v2 + 2 * v3 - 4;
	if (a2 < v5) {
		return ((((unsigned char)v2 ^ (unsigned char)a2) & 1) != 0) + 3;
	}
	if (a2 == v5) {
		return 5;
	} else if (a2 > 2 * (v3 + v2) - 6) {
		if (a2 == 2 * (v3 + v2) - 5) {
			return 7;
		}
		return a2 + 2 * (6 - v3 - v2);
	}
	return 6;
}

//----- (00411540) --------------------------------------------------------
int nox_thing_read_FLOR_411540(nox_memfile* f, uint8_t* a2) {
	if (nox_tile_def_cnt >= 176) {
		return 0;
	}
	int ind = nox_tile_def_cnt;
	nox_tileDef_t* p = &nox_tile_defs_arr[ind];

	nox_memfile_skip(f, 4);

	int sz = nox_memfile_read_u8(f);
	char buf[32];
	nox_memfile_read(buf, 1, sz, f);
	buf[sz] = 0;
	strncpy(&p->name[0], buf, 31);

	char v7 = nox_memfile_read_i8(f);
	char v8 = nox_memfile_read_i8(f);
	char v9 = nox_memfile_read_i8(f);
	uint8_t cbuf[6];    // [esp+0h] [ebp-28h]
	cbuf[0] = v7;
	cbuf[1] = v8;
	cbuf[2] = v9;
	if (v7 != -1 || cbuf[1] || v9 != -1) {
		// TODO: int addressing here doesn't look right!
		p->color_48 = nox_color_rgb_4344A0(*(int*)&cbuf[0], *(int*)&cbuf[1], *(int*)&cbuf[2]);
	} else {
		p->color_48 = 0x80000000;
	}
	p->field_58 = 0;
	if (nox_xxx_checkFacade_4117E0(p) == 1) {
		p->field_58 |= 1u;
	}

	p->field_36 = nox_memfile_read_i32(f);
	p->field_40 = nox_memfile_read_i32(f);

	uint8_t v14 = nox_memfile_read_u8(f);
	p->field_57 = v14;
	v14 = nox_memfile_read_u8(f);
	p->field_53 = v14;
	v14 = nox_memfile_read_u8(f);
	p->field_52 = v14;
	v14 = nox_memfile_read_u8(f);
	p->field_54 = v14;
	p->field_46 = 0;
	v14 = nox_memfile_read_u8(f);
	p->field_55 = v14;
	p->field_56 = v14;
	p->field_44 = (int)(p->field_52) * (int)(p->field_53);
	int v22 = p->field_53 * p->field_54;
	int v23 = p->field_52;
	p->data_32 = 0;
	int v24 = v23 * v22;
	if (v24 > 0) {
		for (int i = 0; i < v24; i++) {
			int v26 = nox_memfile_read_i32(f);
			*a2 = getMemByte(0x5D4594, 251576);
			if (v26 == -1) {
				nox_memfile_skip(f, 1);
				unsigned char v33 = nox_memfile_read_u8(f);
				nox_memfile_read(a2, 1, v33, f);
				a2[v33] = 0;
			}
		}
	}
	++nox_tile_def_cnt;
	return 1;
}
// 411618: variable 'v29' is possibly undefined
// 411540: using guessed type char var_20[32];

//----- (004117E0) --------------------------------------------------------
int nox_xxx_checkFacade_4117E0(nox_tileDef_t* tile) {
	const char* v1;    // eax
	unsigned char* v2; // edi

	v1 = *(const char**)getMemAt(0x587000, 26488);
	if (!*getMemU32Ptr(0x587000, 26488)) {
		return 0;
	}
	v2 = getMemAt(0x587000, 26488);
	while (strcmp(v1, &tile->name[0])) {
		v1 = (const char*)*((uint32_t*)v2 + 1);
		v2 += 4;
		if (!v1) {
			return 0;
		}
	}
	return 1;
}

//----- (00411850) --------------------------------------------------------
int nox_thing_read_EDGE_411850(nox_memfile* f, uint8_t* a2) {
	int v5;            // eax
	int v7;            // ecx
	int v9;            // ecx
	int v12;           // eax
	int v14;           // eax
	unsigned char v17; // dl
	unsigned char v19; // bl
	int v20;           // ecx
	int v21;           // ebx
	int v23;           // ecx
	int v26;           // ecx
	unsigned char v28; // [esp+14h] [ebp-24h]
	char v29[32];      // [esp+18h] [ebp-20h]
	unsigned char v30; // [esp+3Ch] [ebp+4h]
	unsigned char v31; // [esp+3Ch] [ebp+4h]

	if (dword_5d4594_251572 >= 64) {
		return 0;
	}
	nox_memfile_skip(f, 4);
	v30 = nox_memfile_read_u8(f);
	nox_memfile_read(v29, 1u, v30, f);
	v5 = 3 * dword_5d4594_251572;
	v29[v30] = 0;
	strcpy((char*)getMemAt(0x85B3FC, 28644 + 20 * v5), v29);
	v7 = nox_memfile_read_i32(f);
	*getMemU32Ptr(0x85B3FC, 28680 + 60 * dword_5d4594_251572) = v7;
	v9 = nox_memfile_read_i32(f);
	*getMemU32Ptr(0x85B3FC, 28684 + 60 * dword_5d4594_251572) = v9;
	LOBYTE(v9) = nox_memfile_read_u8(f);
	*getMemU8Ptr(0x85B3FC, 28701 + 60 * dword_5d4594_251572) = v9;
	LOBYTE(v9) = nox_memfile_read_u8(f);
	v28 = v9;
	v12 = 60 * dword_5d4594_251572;
	*getMemU8Ptr(0x85B3FC, 28698 + v12) = v9;
	*getMemU16Ptr(0x85B3FC, 28690 + v12) = 0;
	LOBYTE(v9) = nox_memfile_read_u8(f);
	v14 = 60 * dword_5d4594_251572;
	*getMemU8Ptr(0x85B3FC, 28699 + v14) = v9;
	*getMemU8Ptr(0x85B3FC, 28700 + v14) = v9;
	LOBYTE(v9) = nox_memfile_read_u8(f);
	if ((uint8_t)v9 == 1) {
		return 0;
	}
	v17 = nox_memfile_read_u8(f);
	*getMemU8Ptr(0x85B3FC, 28697 + 60 * dword_5d4594_251572) = v17;
	v19 = nox_memfile_read_u8(f);
	v20 = 60 * dword_5d4594_251572;
	*getMemU8Ptr(0x85B3FC, 28696 + v20) = v19;
	*getMemU16Ptr(0x85B3FC, 28688 + v20) = 2 * (v17 + v19);
	*getMemU32Ptr(0x85B3FC, 28676 + v20) = 0;
	if (2 * v28 * (v17 + v19) > 0) {
		v21 = 2 * v28 * (v17 + v19);
		do {
			v23 = nox_memfile_read_i32(f);
			*a2 = getMemByte(0x5D4594, 251580);
			if (v23 == -1) {
				nox_memfile_skip(f, 1);
				v31 = nox_memfile_read_u8(f);
				nox_memfile_read(a2, 1u, v31, f);
				a2[v31] = 0;
			}
			--v21;
		} while (v21);
	}
	v26 = nox_memfile_read_i32(f);
	if (v26 != 0x454E4420) { // "END "
		return 0;
	}
	++dword_5d4594_251572;
	return 1;
}
// 411850: using guessed type char var_20[32];

//----- (00411A90) --------------------------------------------------------
int nox_xxx_mapTileAllowTeleport_411A90(float2* a1) {
	if (*getMemIntPtr(0x587000, 26520) == -1) {
		for (int i = 0; i < 176; i++) {
			nox_tileDef_t* p = &nox_tile_defs_arr[i];
			if (!strcmp(&p->name[0], "WaterNoTeleport")) {
				*getMemU32Ptr(0x587000, 26516) = i;
			} else if (!strcmp(&p->name[0], "WaterDeepNoTeleport")) {
				*getMemU32Ptr(0x587000, 26520) = i;
			} else if (!strcmp(&p->name[0], "WaterShallowNoTeleport")) {
				*getMemU32Ptr(0x587000, 26524) = i;
			} else if (!strcmp(&p->name[0], "WaterSwampDeepNoTeleport")) {
				*getMemU32Ptr(0x587000, 26528) = i;
			} else if (!strcmp(&p->name[0], "WaterSwampShallowNoTeleport")) {
				*getMemU32Ptr(0x587000, 26532) = i;
			}
		}
	}
	int v3 = nox_xxx_tileNFromPoint_411160(a1);
	return v3 == *getMemU32Ptr(0x587000, 26516) || v3 == *getMemU32Ptr(0x587000, 26520) ||
		   v3 == *getMemU32Ptr(0x587000, 26524) || v3 == *getMemU32Ptr(0x587000, 26528) ||
		   v3 == *getMemU32Ptr(0x587000, 26532);
}

//----- (004133C0) --------------------------------------------------------
obj_412ae0_t* nox_xxx_modifNext_4133C0(obj_412ae0_t* a1) { return a1->field_34; }

//----- (004133D0) --------------------------------------------------------
int sub_4133D0(nox_object_t* a1p) {
	int a1 = a1p;
	int v1; // eax
	int v2; // esi
	int v3; // eax

	v1 = *getMemU32Ptr(0x5D4594, 251620);
	v2 = *(uint32_t*)(a1 + 692);
	if (!*getMemU32Ptr(0x5D4594, 251620)) {
		v3 = nox_xxx_modifGetIdByName_413290("Material7");
		v1 = nox_xxx_modifGetDescById_413330(v3);
		*getMemU32Ptr(0x5D4594, 251620) = v1;
	}
	return *(uint32_t*)(a1 + 8) & 0x13001000 && *(uint32_t*)(v2 + 4) == v1;
}

//----- (00413420) --------------------------------------------------------
int sub_413420(char a1) {
	unsigned char* v1; // esi
	int v2;            // ecx
	unsigned char* v3; // eax

	if (!*getMemU32Ptr(0x5D4594, 251624)) {
		v1 = getMemAt(0x587000, 27340);
		do {
			*(uint32_t*)v1 = nox_xxx_gLoadImg_42F970(*((const char**)v1 - 1));
			v1 += 20;
		} while ((int)v1 < (int)getMemAt(0x587000, 27460));
		*getMemU32Ptr(0x5D4594, 251624) = 1;
	}
	v2 = 0;
	v3 = getMemAt(0x587000, 27332);
	while (*v3 != a1) {
		v3 += 20;
		++v2;
		if ((int)v3 >= (int)getMemAt(0x587000, 27452)) {
			return 0;
		}
	}
	return *getMemU32Ptr(0x587000, 27340 + 20 * v2);
}

//----- (004134D0) --------------------------------------------------------
int sub_4134D0() {
	dword_5d4594_251708 = 0;
	dword_5d4594_251712 = 0;
	dword_5d4594_251716 = 0;
	dword_5d4594_251720 = 0;
	return 0;
}

//----- (004137E0) --------------------------------------------------------
void sub_4137E0() {
	if (dword_5d4594_251720) {
		dword_5d4594_251720 = 0;
		sub_43DBE0();
	}
}

//----- (00413890) --------------------------------------------------------
char* sub_413890() {
	unsigned char* v0; // edi
	int v1;            // ecx
	short v2;          // dx
	unsigned char v3;  // al

	*getMemU8Ptr(0x5D4594, 251636) = 0;
	*getMemU8Ptr(0x5D4594, 251637) = 0;
	if (!getMemByte(0x5D4594, 251636)) {
		return 0;
	}
	v0 = getMemAt(0x5D4594, 251637 + strlen((const char*)getMemAt(0x5D4594, 251636)));
	v1 = *getMemU32Ptr(0x587000, 32316);
	v2 = *getMemU16Ptr(0x587000, 32320);
	*(uint32_t*)--v0 = *getMemU32Ptr(0x587000, 32312);
	v3 = getMemByte(0x587000, 32322);
	*((uint32_t*)v0 + 1) = v1;
	*((uint16_t*)v0 + 4) = v2;
	v0[10] = v3;
	return (char*)getMemAt(0x5D4594, 251636);
}

//----- (004138E0) --------------------------------------------------------
void sub_4138E0(int a1) {
	*getMemU32Ptr(0x5D4594, 251740) = nox_xxx_checkGameFlagPause_413A50();
	sub_413A00(1);
}

//----- (00413900) --------------------------------------------------------
void sub_413900(int a1) {
	if (!nox_video_inFadeTransition_44E0D0()) {
		if (!*getMemU32Ptr(0x5D4594, 251740)) {
			sub_413A00(0);
		}
	}
}

//----- (00413920) --------------------------------------------------------
int sub_413920() {
	sub_42EBB0(1u, sub_413900, 0, "Pause");
	sub_42EBB0(2u, sub_4138E0, 0, "Pause");
	dword_5d4594_251744 = 0;
	return 1;
}

//----- (004139B0) --------------------------------------------------------
int sub_4139B0() { return dword_5d4594_251744 != 0; }

//----- (00413A50) --------------------------------------------------------
int nox_xxx_checkGameFlagPause_413A50() { return nox_common_gameFlags_check_40A5C0(0x40000); }

//----- (00413E30) --------------------------------------------------------
void nox_xxx_gameLoopMemDump_413E30() {
	signed int v0;     // ebx
	int v1;            // esi
	unsigned char* v2; // eax
	int v3;            // edx
	signed int v4;     // edi
	const char* v5;    // ebp
	unsigned char* v6; // eax
	signed int v7;     // ecx
	int v8;            // ebx
	signed int v9;     // [esp+8h] [ebp-8h]

	v0 = 0;
	v1 = 0;
	v9 = 0;
	v2 = getMemAt(0x5D4594, 252364);
	do {
		*(uint32_t*)v2 = 0;
		v2 += 84;
	} while ((int)v2 < (int)getMemAt(0x5D4594, 338380));
	v3 = *getMemU32Ptr(0x5D4594, 338304);
	if (*getMemU32Ptr(0x5D4594, 338304)) {
		do {
			v4 = 0;
			if (v0 > 0) {
				v5 = (const char*)getMemAt(0x5D4594, 252284);
				while (strcmp(v5, (const char*)(v3 + 20))) {
					v0 = v9;
					++v4;
					v5 += 84;
					if (v4 >= v9) {
						goto LABEL_10;
					}
				}
				v0 = v9;
				*getMemU32Ptr(0x5D4594, 252364 + 84 * v4) += *(uint32_t*)(v3 + 16);
			}
		LABEL_10:
			if (v4 == v0) {
				*getMemU32Ptr(0x5D4594, 252364 + 84 * v4) = *(uint32_t*)(v3 + 16);
				strcpy((char*)getMemAt(0x5D4594, 252284 + 84 * v4), (const char*)(v3 + 20));
				v9 = ++v0;
			}
			v3 = *(uint32_t*)(v3 + 4);
		} while (v3);
		v1 = 0;
	}
	qsort(getMemAt(0x5D4594, 252284), v0, 0x54u, sub_413F60);
	if (v0 > 0) {
		v6 = getMemAt(0x5D4594, 252364);
		v7 = v0;
		do {
			v8 = *(uint32_t*)v6;
			v6 += 84;
			v1 += v8;
			--v7;
		} while (v7);
	}
}

//----- (00413F60) --------------------------------------------------------
int sub_413F60(const void* a1, const void* a2) { return *((uint32_t*)a1 + 20) - *((uint32_t*)a2 + 20); }

//----- (00414C90) --------------------------------------------------------
char nox_xxx_initSinCosTables_414C90() {
	long long v0;      // rax
	unsigned char* v1; // esi
	unsigned char* v2; // esi
	int v4;            // [esp+0h] [ebp-4h]
	int v5;            // [esp+0h] [ebp-4h]

	LOBYTE(v0) = getMemByte(0x5D4594, 371240);
	if (!getMemByte(0x5D4594, 371240)) {
		*getMemU8Ptr(0x5D4594, 371240) = 1;
		v4 = 0;
		v1 = getMemAt(0x85B3FC, 12260);
		do {
			*(uint32_t*)v1 = (long long)(sin((double)v4 * 0.0015339808) * *getMemDoublePtr(0x581450, 7200));
			v1 += 4;
			++v4;
		} while ((int)v1 < (int)getMemAt(0x85B3FC, 28644));
		v5 = 0;
		v2 = getMemAt(0x5D4594, 338472);
		do {
			v0 = (long long)(acos((double)v5 * 0.00024414062 - 1.0) * *getMemDoublePtr(0x581450, 7184));
			*(uint32_t*)v2 = v0;
			v2 += 4;
			++v5;
		} while ((int)v2 < (int)getMemAt(0x5D4594, 371240));
	}
	return v0;
}

//----- (00414D40) --------------------------------------------------------
int nox_thing_skip_AUD_414D40(nox_memfile* f) {
	int v2;           // edx
	int v3;           // esi
	unsigned char v5; // dl

	v2 = nox_memfile_read_i32(f);
	if (v2 <= 0) {
		return 1;
	}
	v3 = v2;
	do {
		v5 = nox_memfile_read_u8(f);
		nox_memfile_skip(f, v5 + 9);
		while (1) {
			v5 = nox_memfile_read_u8(f);
			if (!v5) {
				break;
			}
			nox_memfile_skip(f, v5);
		}
		--v3;
	} while (v3);
	return 1;
}

//----- (00414DB0) --------------------------------------------------------
int nox_thing_read_FLOR_414DB0(nox_memfile* f) {
	unsigned char* v2;  // ecx
	unsigned char* v3;  // ecx
	unsigned char* v4;  // esi
	int v5;             // edx
	int v6;             // ecx
	int* v7;            // edx
	int v8;             // esi
	int v9;             // edx
	unsigned char* v10; // edx
	int* v11;           // ecx
	int v12;            // edx
	unsigned char v14;  // [esp+8h] [ebp+4h]
	unsigned char v15;  // [esp+8h] [ebp+4h]

	v2 = (unsigned char*)(f->cur + 4);
	f->cur = v2;
	v3 = &v2[*v2 + 13];
	f->cur = v3;
	v4 = (unsigned char*)(f->cur + 1);
	v14 = *v3;
	f->cur = v4;
	v15 = *v4;
	f->cur = v4 + 1;
	v4 += 3;
	v5 = v14 * v15;
	v14 = *(v4 - 2);
	f->cur = v4;
	v6 = v5 * v14;
	if (v6 > 0) {
		do {
			v7 = f->cur;
			v8 = *v7;
			v9 = (int)(v7 + 1);
			f->cur = v9;
			if (v8 == -1) {
				v10 = (unsigned char*)(v9 + 1);
				f->cur = v10;
				f->cur += *v10 + 1;
			}
			--v6;
		} while (v6);
	}
	v11 = f->cur;
	v12 = *v11;
	f->cur = v11 + 1;
	return v12 == 1162757152;
}

//----- (00414E70) --------------------------------------------------------
int nox_thing_read_EDGE_414E70(nox_memfile* f, void* a2) {
	int a1 = f;
	int v2;             // esi
	unsigned char* v3;  // edx
	unsigned char* v4;  // eax
	unsigned char v5;   // cl
	unsigned char v6;   // cl
	unsigned char* v7;  // eax
	unsigned char v9;   // dl
	unsigned char* v10; // eax
	int v11;            // eax
	int* v12;           // ecx
	int v13;            // edx
	int v14;            // ecx
	unsigned char* v15; // ecx
	int* v16;           // eax
	int v17;            // ecx
	int v18;            // [esp-4h] [ebp-10h]
	unsigned char v19;  // [esp+8h] [ebp-4h]
	unsigned char v20;  // [esp+10h] [ebp+4h]
	unsigned char v21;  // [esp+14h] [ebp+8h]

	v2 = a1;
	v18 = a1;
	v3 = (unsigned char*)(*(uint32_t*)(a1 + 8) + 4);
	*(uint32_t*)(a1 + 8) = v3;
	v20 = *v3;
	*(uint32_t*)(v2 + 8) = v3 + 1;
	nox_memfile_read(a2, 1u, v20, v18);
	v4 = (unsigned char*)(*(uint32_t*)(v2 + 8) + 9);
	*(uint32_t*)(v2 + 8) = v4;
	v5 = *v4;
	v4 += 2;
	v19 = v5;
	*(uint32_t*)(v2 + 8) = v4;
	v6 = *v4;
	v7 = v4 + 1;
	*(uint32_t*)(v2 + 8) = v7;
	if (v6 == 1) {
		return 0;
	}
	v9 = *v7;
	v10 = v7 + 1;
	*(uint32_t*)(v2 + 8) = v10;
	v21 = *v10;
	*(uint32_t*)(v2 + 8) = v10 + 1;
	v11 = 2 * v19 * (v9 + v21);
	if (v11 > 0) {
		do {
			v12 = *(int**)(v2 + 8);
			v13 = *v12;
			v14 = (int)(v12 + 1);
			*(uint32_t*)(v2 + 8) = v14;
			if (v13 == -1) {
				v15 = (unsigned char*)(v14 + 1);
				*(uint32_t*)(v2 + 8) = v15;
				*(uint32_t*)(v2 + 8) = &v15[*v15 + 1];
			}
			--v11;
		} while (v11);
	}
	v16 = *(int**)(v2 + 8);
	v17 = *v16;
	*(uint32_t*)(v2 + 8) = v16 + 1;
	return v17 == 1162757152;
}

//----- (00414F60) --------------------------------------------------------
int nox_thing_read_WALL_414F60(nox_memfile* f, void* a2) {
	uint32_t* a1 = f;
	uint32_t* v2;       // esi
	void* v3;           // edi
	unsigned char* v4;  // eax
	int v5;             // ebp
	unsigned char* v6;  // eax
	unsigned char* v7;  // eax
	unsigned char* v8;  // eax
	unsigned char* v9;  // eax
	int v10;            // ebp
	int v11;            // edi
	int v12;            // edx
	int* v13;           // eax
	int v14;            // ecx
	int v15;            // eax
	unsigned char* v16; // eax
	int* v17;           // eax
	int v18;            // ecx
	uint32_t* v20;      // [esp-4h] [ebp-18h]
	int v21;            // [esp+10h] [ebp-4h]
	unsigned char v22;  // [esp+18h] [ebp+4h]
	unsigned char v23;  // [esp+18h] [ebp+4h]
	unsigned char v24;  // [esp+18h] [ebp+4h]
	unsigned char v25;  // [esp+18h] [ebp+4h]
	unsigned char v26;  // [esp+18h] [ebp+4h]

	v2 = a1;
	v3 = a2;
	v20 = a1;
	v4 = (unsigned char*)(a1[2] + 4);
	a1[2] = v4;
	v22 = *v4;
	v2[2] = v4 + 1;
	nox_memfile_read(v3, 1u, v22, (int)v20);
	*((uint8_t*)v3 + v22) = 0;
	v2[2] += 14;
	nox_memfile_read64align_40AD60((char*)&a2, 1, 1, v2);
	v5 = 0;
	if ((uint8_t)a2) {
		while (v5 < 8) {
			v6 = (unsigned char*)v2[2];
			v23 = *v6;
			v2[2] = v6 + 1;
			nox_memfile_read(v3, 1u, v23, (int)v2);
			*((uint8_t*)v3 + v23) = 0;
			if (++v5 >= (unsigned char)a2) {
				goto LABEL_4;
			}
		}
		return 0;
	}
LABEL_4:
	v7 = (unsigned char*)v2[2];
	v24 = *v7;
	v2[2] = v7 + 1;
	nox_memfile_read(v3, 1u, v24, (int)v2);
	*((uint8_t*)v3 + v24) = 0;
	v8 = (unsigned char*)v2[2];
	v25 = *v8;
	v2[2] = v8 + 1;
	nox_memfile_read(v3, 1u, v25, (int)v2);
	*((uint8_t*)v3 + v25) = 0;
	v9 = (unsigned char*)v2[2];
	v26 = *v9;
	v2[2] = v9 + 1;
	nox_memfile_read(v3, 1u, v26, (int)v2);
	*((uint8_t*)v3 + v26) = 0;
	++v2[2];
	v10 = 15;
	do {
		nox_memfile_read64align_40AD60((char*)&v21, 1, 1, v2);
		if ((int)(unsigned char)v21 > 0) {
			v11 = (unsigned char)v21;
			do {
				v12 = 4;
				do {
					v13 = (int*)(v2[2] + 8);
					v2[2] = v13;
					v14 = *v13;
					v15 = (int)(v13 + 1);
					v2[2] = v15;
					if (v14 == -1) {
						v16 = (unsigned char*)(v15 + 1);
						v2[2] = v16;
						v2[2] = &v16[*v16 + 1];
					}
					--v12;
				} while (v12);
				--v11;
			} while (v11);
		}
		--v10;
	} while (v10);
	v17 = (int*)v2[2];
	v18 = *v17;
	v2[2] = v17 + 1;
	return v18 == 1162757152;
}

//----- (00415100) --------------------------------------------------------
int nox_thing_skip_spells_415100(nox_memfile* f) {
	int a1 = f;
	int* v1;            // ecx
	int v2;             // edx
	int v3;             // esi
	unsigned char* v4;  // edx
	int* v5;            // ecx
	int v6;             // edx
	int v7;             // ecx
	unsigned char* v8;  // ecx
	int* v9;            // ecx
	int v10;            // edx
	int v11;            // ecx
	unsigned char* v12; // ecx
	unsigned char* v13; // ecx
	short* v14;         // edx
	unsigned char* v15; // ecx
	unsigned char* v16; // ecx
	unsigned char* v17; // ecx

	v1 = *(int**)(a1 + 8);
	v2 = *v1;
	*(uint32_t*)(a1 + 8) = v1 + 1;
	if (v2 <= 0) {
		return 1;
	}
	v3 = v2;
	do {
		v4 = (unsigned char*)(**(unsigned char**)(a1 + 8) + 4 + *(uint32_t*)(a1 + 8));
		*(uint32_t*)(a1 + 8) = v4;
		v5 = (int*)(*v4 + 1 + *(uint32_t*)(a1 + 8));
		*(uint32_t*)(a1 + 8) = v5;
		v6 = *v5;
		v7 = (int)(v5 + 1);
		*(uint32_t*)(a1 + 8) = v7;
		if (v6 == -1) {
			v8 = (unsigned char*)(v7 + 1);
			*(uint32_t*)(a1 + 8) = v8;
			*(uint32_t*)(a1 + 8) += *v8 + 1;
		}
		v9 = *(int**)(a1 + 8);
		v10 = *v9;
		v11 = (int)(v9 + 1);
		*(uint32_t*)(a1 + 8) = v11;
		if (v10 == -1) {
			v12 = (unsigned char*)(v11 + 1);
			*(uint32_t*)(a1 + 8) = v12;
			*(uint32_t*)(a1 + 8) += *v12 + 1;
		}
		v13 = (unsigned char*)(*(uint32_t*)(a1 + 8) + 4);
		*(uint32_t*)(a1 + 8) = v13;
		v14 = (short*)(*v13 + 1 + *(uint32_t*)(a1 + 8));
		*(uint32_t*)(a1 + 8) = v14;
		v15 = (unsigned char*)v14 + *v14 + 2;
		*(uint32_t*)(a1 + 8) = v15;
		v16 = (unsigned char*)(*v15 + 1 + *(uint32_t*)(a1 + 8));
		*(uint32_t*)(a1 + 8) = v16;
		v17 = (unsigned char*)(*v16 + 1 + *(uint32_t*)(a1 + 8));
		*(uint32_t*)(a1 + 8) = v17;
		--v3;
		*(uint32_t*)(a1 + 8) += *v17 + 1;
	} while (v3);
	return 1;
}

//----- (00415240) --------------------------------------------------------
int nox_thing_read_image_415240(nox_memfile* f) {
	int a1 = f;
	int v1;             // ecx
	int* v2;            // eax
	int v3;             // edx
	int v4;             // ebp
	uint8_t* v6;        // esi
	unsigned char* v7;  // esi
	unsigned char v8;   // bl
	int v9;             // edi
	int* v10;           // edx
	int v11;            // esi
	int v12;            // edx
	unsigned char* v13; // edx
	unsigned char v14;  // [esp+0h] [ebp-4h]
	char v15;           // [esp+8h] [ebp+4h]

	v1 = a1;
	v2 = *(int**)(a1 + 8);
	v3 = *v2;
	*(uint32_t*)(a1 + 8) = v2 + 1;
	if (v3 <= 0) {
		return 1;
	}
	for (v4 = v3; v4; v4--) {
		v6 = (uint8_t*)(**(unsigned char**)(v1 + 8) + 1 + *(uint32_t*)(v1 + 8));
		*(uint32_t*)(v1 + 8) = v6;
		v15 = *v6;
		*(uint32_t*)(v1 + 8) = v6 + 1;
		v7 = v6 + 1;
		if (v15 == 1 || v15 != 2) {
			v14 = 1;
		} else {
			v8 = *v7;
			*(uint32_t*)(v1 + 8) = v7 + 2;
			v14 = v8;
			*(uint32_t*)(v1 + 8) += v7[2] + 1;
			if (!v8) {
				continue;
			}
		}
		v9 = v14;
		do {
			v10 = *(int**)(v1 + 8);
			v11 = *v10;
			v12 = (int)(v10 + 1);
			*(uint32_t*)(v1 + 8) = v12;
			if (v11 == -1) {
				v13 = (unsigned char*)(v12 + 1);
				*(uint32_t*)(v1 + 8) = v13;
				*(uint32_t*)(v1 + 8) += *v13 + 1;
			}
			--v9;
		} while (v9);
	}
	return 1;
}

//----- (00415320) --------------------------------------------------------
int nox_thing_read_ability_415320(nox_memfile* f) {
	int a1 = f;
	int* v1;            // ecx
	int v2;             // edx
	int v3;             // esi
	int* v4;            // ecx
	int v5;             // edx
	int v6;             // ecx
	unsigned char* v7;  // ecx
	int* v8;            // ecx
	int v9;             // edx
	int v10;            // ecx
	unsigned char* v11; // ecx
	int* v12;           // ecx
	int v13;            // edx
	int v14;            // ecx
	unsigned char* v15; // ecx
	short* v16;         // ecx
	unsigned char* v17; // ecx
	unsigned char* v18; // ecx
	unsigned char* v19; // ecx

	v1 = *(int**)(a1 + 8);
	v2 = *v1;
	*(uint32_t*)(a1 + 8) = v1 + 1;
	if (v2 <= 0) {
		return 1;
	}
	v3 = v2;
	do {
		v4 = (int*)(**(unsigned char**)(a1 + 8) + 2 + *(uint32_t*)(a1 + 8));
		*(uint32_t*)(a1 + 8) = v4;
		v5 = *v4;
		v6 = (int)(v4 + 1);
		*(uint32_t*)(a1 + 8) = v6;
		if (v5 == -1) {
			v7 = (unsigned char*)(v6 + 1);
			*(uint32_t*)(a1 + 8) = v7;
			*(uint32_t*)(a1 + 8) += *v7 + 1;
		}
		v8 = *(int**)(a1 + 8);
		v9 = *v8;
		v10 = (int)(v8 + 1);
		*(uint32_t*)(a1 + 8) = v10;
		if (v9 == -1) {
			v11 = (unsigned char*)(v10 + 1);
			*(uint32_t*)(a1 + 8) = v11;
			*(uint32_t*)(a1 + 8) += *v11 + 1;
		}
		v12 = *(int**)(a1 + 8);
		v13 = *v12;
		v14 = (int)(v12 + 1);
		*(uint32_t*)(a1 + 8) = v14;
		if (v13 == -1) {
			v15 = (unsigned char*)(v14 + 1);
			*(uint32_t*)(a1 + 8) = v15;
			*(uint32_t*)(a1 + 8) += *v15 + 1;
		}
		v16 = (short*)(**(unsigned char**)(a1 + 8) + 1 + *(uint32_t*)(a1 + 8));
		*(uint32_t*)(a1 + 8) = v16;
		v17 = (unsigned char*)v16 + *v16 + 2;
		*(uint32_t*)(a1 + 8) = v17;
		v18 = (unsigned char*)(*v17 + 1 + *(uint32_t*)(a1 + 8));
		*(uint32_t*)(a1 + 8) = v18;
		v19 = (unsigned char*)(*v18 + 1 + *(uint32_t*)(a1 + 8));
		*(uint32_t*)(a1 + 8) = v19;
		--v3;
		*(uint32_t*)(a1 + 8) += *v19 + 1;
	} while (v3);
	return 1;
}

//----- (00415660) --------------------------------------------------------
int nox_thing_read_audio_415660(nox_memfile* a1p, char* a2) {
	int a1 = a1p;
	int v2;  // ebx
	int* v3; // eax
	int v4;  // edi

	v2 = 0;
	v3 = *(int**)(a1 + 8);
	v4 = *v3;
	*(uint32_t*)(a1 + 8) = v3 + 1;
	if (v4 <= 0) {
		return 1;
	}
	while (sub_452BD0(a1, a2)) {
		if (++v2 >= v4) {
			return 1;
		}
	}
	return 0;
}

//----- (00415960) --------------------------------------------------------
int sub_415960(wchar2_t* a1) {
	int v1;             // edi
	const wchar2_t** v2; // eax
	unsigned char* v3;  // esi
	int v4;             // ecx

	v1 = 0;
	if (!*getMemU32Ptr(0x587000, 33392)) {
		return 0;
	}
	v2 = (const wchar2_t**)getMemAt(0x587000, 33392);
	v3 = getMemAt(0x587000, 33392);
	while (_nox_wcsicmp(a1, *v2)) {
		v4 = *((uint32_t*)v3 + 3);
		v3 += 12;
		++v1;
		v2 = (const wchar2_t**)v3;
		if (!v4) {
			return 0;
		}
	}
	return *getMemU32Ptr(0x587000, 33400 + 12 * v1);
}

//----- (004159F0) --------------------------------------------------------
int sub_4159F0(int a1) {
	int v1;           // ecx
	unsigned char* i; // eax
	int v3;           // esi

	v1 = 0;
	if (!*getMemU32Ptr(0x587000, 33392)) {
		return 0;
	}
	for (i = getMemAt(0x587000, 33392); *((uint32_t*)i + 2) != a1; i += 12) {
		v3 = *((uint32_t*)i + 3);
		++v1;
		if (!v3) {
			return 0;
		}
	}
	return *getMemU32Ptr(0x587000, 33392 + 12 * v1);
}

//----- (00415BD0) --------------------------------------------------------
double sub_415BD0(int a1) {
	float* v1;     // eax
	double result; // st7

	if (*(uint32_t*)(a1 + 8) & 0x2000000 &&
		(v1 = (float*)nox_xxx_equipClothFindDefByTT_413270(*(unsigned short*)(a1 + 4))) != 0) {
		result = v1[16];
	} else {
		result = 0.0;
	}
	return result;
}

//----- (00415C00) --------------------------------------------------------
double nox_xxx_itemApplyDefendEffect_415C00(int a1) {
	int* v1;                                               // edi
	float* v2;                                             // eax
	int v4;                                                // eax
	void (*v5)(int, int, uint32_t, int, uint32_t, float*); // ecx
	float v6;                                              // [esp+8h] [ebp-4h]

	v6 = 0.0;
	if (!(*(uint32_t*)(a1 + 8) & 0x2000000)) {
		return 0.0;
	}
	v1 = *(int**)(a1 + 692);
	v2 = (float*)nox_xxx_equipClothFindDefByTT_413270(*(unsigned short*)(a1 + 4));
	if (!v2) {
		return v6;
	}
	v6 = v2[16];
	v4 = *v1;
	if (*v1) {
		v5 = *(void (**)(int, int, uint32_t, int, uint32_t, float*))(v4 + 76);
		if (v5) {
			v5(v4, a1, 0, a1, 0, &v6);
		}
	}
	return v6;
}

//----- (00415DA0) --------------------------------------------------------
int sub_415DA0(wchar2_t* a1) {
	int v1;             // edi
	const wchar2_t** v2; // eax
	unsigned char* v3;  // esi
	int v4;             // ecx

	v1 = 0;
	if (!*getMemU32Ptr(0x587000, 35496)) {
		return 0;
	}
	v2 = (const wchar2_t**)getMemAt(0x587000, 35496);
	v3 = getMemAt(0x587000, 35496);
	while (_nox_wcsicmp(a1, *v2)) {
		v4 = *((uint32_t*)v3 + 3);
		v3 += 12;
		++v1;
		v2 = (const wchar2_t**)v3;
		if (!v4) {
			return 0;
		}
	}
	return *getMemU32Ptr(0x587000, 35504 + 12 * v1);
}

//----- (00415E80) --------------------------------------------------------
int sub_415E80(int a1) {
	int v1;           // ecx
	unsigned char* i; // eax
	int v3;           // esi

	v1 = 0;
	if (!*getMemU32Ptr(0x587000, 35496)) {
		return 0;
	}
	for (i = getMemAt(0x587000, 35496); *((uint32_t*)i + 2) != a1; i += 12) {
		v3 = *((uint32_t*)i + 3);
		++v1;
		if (!v3) {
			return 0;
		}
	}
	return *getMemU32Ptr(0x587000, 35496 + 12 * v1);
}

//----- (004161E0) --------------------------------------------------------
int sub_4161E0() {
	int v0;            // ebx
	int v1;            // eax
	char v2;           // al
	unsigned short v3; // dx
	short v4;          // ax
	char* v5;          // eax
	unsigned char v6;  // cl
	unsigned char v7;  // cl
	int result;        // eax
	unsigned char v9;  // [esp+8h] [ebp-3Ch]
	unsigned char v10; // [esp+8h] [ebp-3Ch]
	char v11[4];       // [esp+Ch] [ebp-38h]
	int v12;           // [esp+10h] [ebp-34h]
	char v13[12];      // [esp+14h] [ebp-30h]
	char v14[16];      // [esp+20h] [ebp-24h]
	int v15[5];        // [esp+30h] [ebp-14h]

	sub_454040(v15);
	sub_4536B0(v11);
	v12 = sub_453710();
	v0 = nox_common_playerInfoCount_416F40();
	v1 = nox_xxx_servGetPlrLimit_409FA0();
	if (nox_common_getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) {
		--v0;
		--v1;
	}
	if (getMemByte(0x5D4594, 371516 + 103) != v0) {
		*getMemU8Ptr(0x5D4594, 371516 + 103) = v0;
		dword_5d4594_371692 = 1;
	}
	if (getMemByte(0x5D4594, 371516 + 104) != v1) {
		*getMemU8Ptr(0x5D4594, 371516 + 104) = v1;
		dword_5d4594_371692 = 1;
	}
	if ((getMemByte(0x5D4594, 371516 + 102) & 0xEF) != sub_43BE50_get_video_mode_id()) {
		v2 = sub_43BE50_get_video_mode_id();
		dword_5d4594_371692 = 1;
		*getMemU8Ptr(0x5D4594, 371516 + 102) = getMemByte(0x5D4594, 371516 + 102) & 0x80 | v2;
	}
	v3 = nox_common_gameFlags_getVal_40A5B0();
	v4 = *getMemU16Ptr(0x5D4594, 371380 + 52);
	if ((*getMemU16Ptr(0x5D4594, 371380 + 52) ^ v3) & 0xFFF0) {
		*getMemU16Ptr(0x5D4594, 371380 + 52) = nox_common_gameFlags_getVal_40A5B0();
		v4 = *getMemU16Ptr(0x5D4594, 371380 + 52);
		dword_5d4594_371692 = 1;
	}
	if (*getMemU16Ptr(0x5D4594, 371380 + 54) != nox_xxx_servGamedataGet_40A020(v4)) {
		*getMemU16Ptr(0x5D4594, 371380 + 54) = nox_xxx_servGamedataGet_40A020(*getMemI16Ptr(0x5D4594, 371380 + 52));
		dword_5d4594_371692 = 1;
	}
	if (getMemByte(0x5D4594, 371380 + 56) != sub_40A180(*getMemI16Ptr(0x5D4594, 371380 + 52))) {
		*getMemU8Ptr(0x5D4594, 371380 + 56) = sub_40A180(*getMemI16Ptr(0x5D4594, 371380 + 52));
		dword_5d4594_371692 = 1;
	}
	strncpy(v14, (const char*)getMemAt(0x5D4594, 371380 + 9), 0xFu);
	v14[15] = 0;
	if (strcmp(v14, nox_xxx_serverOptionsGetServername_40A4C0())) {
		v5 = nox_xxx_serverOptionsGetServername_40A4C0();
		strncpy((char*)getMemAt(0x5D4594, 371380 + 9), v5, 0xFu);
		dword_5d4594_371692 = 1;
	}
	strcpy(v13, (const char*)getMemAt(0x5D4594, 371380));
	if (strcmp(v13, nox_xxx_mapGetMapName_409B40())) {
		strcpy((char*)getMemAt(0x5D4594, 371380), nox_xxx_mapGetMapName_409B40());
		dword_5d4594_371692 = 1;
	}
	v6 = 0;
	v9 = 0;
	do {
		if (*getMemU32Ptr(0x5D4594, 371380 + 24 + 4*v9) != v15[v9]) {
			break;
		}
		v9 = ++v6;
	} while (v6 < 5u);
	if (v6 != 5) {
		dword_5d4594_371692 = 1;
		memcpy(getMemAt(0x5D4594, 371380 + 24), v15, 0x14u);
	}
	if (!nox_common_gameFlags_check_40A5C0(1)) {
		return dword_5d4594_371692;
	}
	if (*getMemU32Ptr(0x5D4594, 371380 + 48) != v12) {
		*getMemU32Ptr(0x5D4594, 371380 + 48) = v12;
		dword_5d4594_371692 = 1;
	}
	v7 = 0;
	v10 = 0;
	do {
		if (getMemByte(0x5D4594, 371380 + 44 + v10) != v11[v10]) {
			break;
		}
		v10 = ++v7;
	} while (v7 < 4u);
	if (v7 == 4) {
		return dword_5d4594_371692;
	}
	*getMemU32Ptr(0x5D4594, 371380 + 44) = *(uint32_t*)v11;
	result = 1;
	dword_5d4594_371692 = 1;
	return result;
}
// 4161E0: using guessed type int var_14[5];

//----- (004164F0) --------------------------------------------------------
void sub_4164F0() { dword_5d4594_371692 = 0; }

//----- (00416580) --------------------------------------------------------
int sub_416580() { return *getMemU32Ptr(0x5D4594, 371688); }

//----- (00416590) --------------------------------------------------------
char* nox_xxx_cliGamedataGet_416590(int a1) { return (char*)getMemAt(0x5D4594, 371380 + 58*a1); }

//----- (004165B0) --------------------------------------------------------
char* sub_4165B0() { return (char*)getMemAt(0x5D4594, 371380 + 58 * *getMemU32Ptr(0x5D4594, 371688)); }

//----- (004165D0) --------------------------------------------------------
char* sub_4165D0(int a1) {
	*getMemU32Ptr(0x5D4594, 371688) = a1;
	return (char*)getMemAt(0x5D4594, 371380 + 58*a1);
}

//----- (004165F0) --------------------------------------------------------
int sub_4165F0(int a1, int a2) {
	int result; // eax

	result = a2;
	memcpy(getMemAt(0x5D4594, 371380 + 58*a2), getMemAt(0x5D4594, 371380 + 58*a1), 0x3Au);
	return result;
}

//----- (00416630) --------------------------------------------------------
char* sub_416630() { return (char*)getMemAt(0x5D4594, 371516 + 100); }

//----- (00416640) --------------------------------------------------------
void* sub_416640() { return getMemAt(0x5D4594, 371516); }

//----- (00416650) --------------------------------------------------------
int sub_416650() { return *getMemU32Ptr(0x5D4594, 371700); }

//----- (00416690) --------------------------------------------------------
void sub_416690() {
	char* v0;    // esi
	char* v1;    // ebx
	short v2;    // ax
	char v3[84]; // [esp+0h] [ebp-54h]

	if (0) {
		v0 = nox_xxx_cliGamedataGet_416590(0);
		sub_4161E0();
		v1 = sub_416630();
		memcpy(v1 + 11, v0, 0x3Au);
		if (nox_xxx_isQuest_4D6F50()) {
			v2 = *(uint16_t*)(v1 + 63) & 0xFF7F;
			HIBYTE(v2) |= 0x10u;
			*(uint16_t*)(v1 + 63) = v2;
			*(uint16_t*)(v1 + 65) = nox_game_getQuestStage_4E3CC0();
		}
		v3[sub_425550(v1, v3, 552)] = 0;
		sub_4164F0();
	}
}
// 416690: using guessed type char var_54[84];

//----- (00416720) --------------------------------------------------------
void sub_416720() {
	int v0 = 0;
	int v2 = sub_416900();
	while (v2) {
		int* v3 = sub_416910((int*)v2);
		if (*(uint32_t*)(v2 + 68) || (*(uint32_t*)(v2 + 64))) {
			if (nox_platform_get_ticks() > *(uint64_t*)(v2 + 64)) {
				sub_416820(v0);
			}
		}
		++v0;
		v2 = (int)v3;
	}
}

//----- (00416770) --------------------------------------------------------
int* sub_416770(int a1, wchar2_t* a2, const char* a3) {
	uint32_t* v3; // ebp

	v3 = calloc(1, 0x60u);
	sub_425770(v3);
	nox_wcscpy((wchar2_t*)v3 + 6, a2);
	if (a3) {
		strcpy((char*)v3 + 72, a3);
	} else {
		*((uint8_t*)v3 + 72) = 0;
	}
	nox_common_list_append_4258E0((int)getMemAt(0x5D4594, 371500), v3);
	if (a1) {
		*((uint64_t*)v3 + 8) = 60000 * a1 + nox_platform_get_ticks();
	} else {
		v3[16] = 0;
		v3[17] = 0;
	}
	return sub_455800();
}

//----- (00416820) --------------------------------------------------------
void sub_416820(int a1) {
	int* v2; // esi
	int v3;  // edi
	int v4;  // eax

	v2 = sub_416900();
	if (!v2) {
		return;
	}
	v3 = a1;
	while (1) {
		v4 = v3--;
		if (!v4) {
			break;
		}
		v2 = sub_416910(v2);
		if (!v2) {
			return;
		}
	}
	nox_common_list_remove_425920((uint32_t**)v2);
	free(v2);
}

//----- (00416860) --------------------------------------------------------
int* sub_416860(int a1) {
	int* result; // eax
	int* v2;     // esi
	int v3;      // edi
	int v4;      // eax

	result = sub_4168E0();
	v2 = result;
	if (result) {
		v3 = a1;
		while (1) {
			v4 = v3--;
			if (!v4) {
				break;
			}
			result = sub_4168F0(v2);
			v2 = result;
			if (!result) {
				return result;
			}
		}
		nox_common_list_remove_425920((uint32_t**)v2);
		free(v2);
	}
	return result;
}

//----- (004168A0) --------------------------------------------------------
int* sub_4168A0(wchar2_t* a1) {
	wchar2_t* v1; // esi

	v1 = (wchar2_t*)calloc(1, 0x40u);
	sub_425770(v1);
	nox_wcscpy(v1 + 6, a1);
	nox_common_list_append_4258E0((int)getMemAt(0x5D4594, 371364), v1);
	return sub_455800();
}

//----- (004168E0) --------------------------------------------------------
int* sub_4168E0() { return nox_common_list_getFirstSafe_425890(getMemIntPtr(0x5D4594, 371364)); }

//----- (004168F0) --------------------------------------------------------
int* sub_4168F0(int* a1) { return nox_common_list_getNextSafe_4258A0(a1); }

//----- (00416900) --------------------------------------------------------
int* sub_416900() { return nox_common_list_getFirstSafe_425890(getMemIntPtr(0x5D4594, 371500)); }

//----- (00416910) --------------------------------------------------------
int* sub_416910(int* a1) { return nox_common_list_getNextSafe_4258A0(a1); }

//----- (00416920) --------------------------------------------------------
int sub_416920() {
	nox_common_list_clear_425760(getMemAt(0x5D4594, 371364));
	nox_common_list_clear_425760(getMemAt(0x5D4594, 371500));
	return sub_4E41B0((char*)getMemAt(0x587000, 54280));
}

//----- (00416950) --------------------------------------------------------
int* sub_416950() {
	int* v0;     // esi
	int* v1;     // edi
	int* result; // eax
	int* v3;     // esi
	int* v4;     // edi

	sub_4E43F0("ban.txt");
	v0 = sub_4168E0();
	if (v0) {
		do {
			v1 = sub_4168F0(v0);
			nox_common_list_remove_425920((uint32_t**)v0);
			free(v0);
			v0 = v1;
		} while (v1);
	}
	result = sub_416900();
	v3 = result;
	if (result) {
		do {
			v4 = sub_416910(v3);
			nox_common_list_remove_425920((uint32_t**)v3);
			free(v3);
			v3 = v4;
		} while (v4);
	}
	return result;
}

//----- (004169C0) --------------------------------------------------------
int sub_4169C0() { return *getMemU32Ptr(0x5D4594, 371704); }

//----- (004169D0) --------------------------------------------------------
int nox_xxx_cliSetSettingsAcquired_4169D0(int a1) {
	int result; // eax

	result = a1;
	*getMemU32Ptr(0x5D4594, 371704) = a1;
	return result;
}

//----- (004169F0) --------------------------------------------------------
char* sub_4169F0() {
	char* result; // eax

	result = sub_416640();
	result[100] &= 0xEFu;
	return result;
}

//----- (00416E50) --------------------------------------------------------
char* nox_xxx_playerForceSendLessons_416E50(int a1) {
	char* result; // eax
	int* i;       // esi

	result = nox_common_playerInfoGetFirst_416EA0();
	for (i = (int*)result; result; i = (int*)result) {
		i[534] = 0;
		i[535] = 0;
		if (a1) {
			if (i[514]) {
				nox_xxx_netReportLesson_4D8EF0(i[514]);
			}
		}
		result = nox_common_playerInfoGetNext_416EE0((int)i);
	}
	return result;
}

//----- (004170D0) --------------------------------------------------------
char* nox_xxx_playerByName_4170D0(wchar2_t* a1) {
	char* v1; // esi

	if (!a1) {
		return 0;
	}
	v1 = nox_common_playerInfoGetFirst_416EA0();
	if (!v1) {
		return 0;
	}
	while (_nox_wcsicmp((const wchar2_t*)v1 + 2352, a1)) {
		v1 = nox_common_playerInfoGetNext_416EE0((int)v1);
		if (!v1) {
			return 0;
		}
	}
	return v1;
}

//----- (00417270) --------------------------------------------------------
int sub_417270(int a1) {
	int result; // eax
	int v2;     // edx
	int v3;     // ecx

	result = 0;
	if (a1 >= 0 && a1 < NOX_PLAYERINFO_MAX) {
		nox_playerInfo* pl = nox_common_playerInfoFromNumRaw(a1);
		v2 = pl->field_4580;
		if (v2) {
			v3 = *(uint32_t*)(v2 + 8);
			for (result = 1; v3 != v2; ++result) {
				v3 = *(uint32_t*)(v3 + 8);
			}
		}
	}
	return result;
}

// mark spell -- is tracked
//----- (004173D0) --------------------------------------------------------
int nox_xxx_playerMapTracksObj_4173D0(int a1, nox_object_t* a2p) {
	int a2 = a2p;
	int result; // eax
	int v3;     // ecx

	result = 0;
	if (a1 >= 0 && a1 < NOX_PLAYERINFO_MAX) {
		if (a2) {
			nox_playerInfo* pl = nox_common_playerInfoFromNumRaw(a1);
			v3 = pl->field_4580;
			if (v3) {
				while (*(uint32_t*)(v3 + 4) != a2) {
					v3 = *(uint32_t*)(v3 + 8);
					if (v3 == pl->field_4580 || !v3) {
						return result;
					}
				}
				result = 1;
			}
		}
	}
	return result;
}

//----- (00417470) --------------------------------------------------------
char* nox_xxx_netUnmarkMinimapSpec_417470(int a1, int a2) {
	char* result; // eax
	int i;        // esi

	result = nox_common_playerInfoGetFirst_416EA0();
	for (i = (int)result; result; i = (int)result) {
		nox_xxx_netUnmarkMinimapObj_417300(*(unsigned char*)(i + 2064), a1, a2);
		result = nox_common_playerInfoGetNext_416EE0(i);
	}
	return result;
}

//----- (004174B0) --------------------------------------------------------
char* nox_xxx_netMarkMinimapForAll_4174B0(int a1, int a2) {
	char* result; // eax
	int i;        // esi

	result = nox_common_playerInfoGetFirst_416EA0();
	for (i = (int)result; result; i = (int)result) {
		nox_xxx_netMarkMinimapObject_417190(*(unsigned char*)(i + 2064), a1, a2);
		result = nox_common_playerInfoGetNext_416EE0(i);
	}
	return result;
}

//----- (004174F0) --------------------------------------------------------
int nox_xxx_netNeedTimestampStatus_4174F0(nox_playerInfo* a1p, int a2) {
	int a1 = a1p;
	int result; // eax

	*(uint32_t*)(a1 + 3680) |= a2;
	result = nox_common_gameFlags_check_40A5C0(1);
	if (result) {
		if (a2 & 0x423) {
			result = nox_xxx_netReportPlayerStatus_417630(a1);
		}
	}
	return result;
}

//----- (00417530) --------------------------------------------------------
char nox_xxx_playerUnsetStatus_417530(nox_playerInfo* a1p, int a2) {
	int a1 = a1p;
	int v2;   // eax
	short v3; // ax

	*(uint32_t*)(a1 + 3680) &= ~a2;
	v2 = nox_common_gameFlags_check_40A5C0(1);
	if (v2) {
		if (a2 & 1) {
			v2 = nox_common_gameFlags_check_40A5C0(128);
			if (!v2) {
				v2 = nox_xxx_gamePlayIsAnyPlayers_40A8A0();
				if (v2) {
					v2 = sub_40A220();
					if (!v2) {
						v3 = nox_common_gameFlags_getVal_40A5B0();
						LOBYTE(v2) = sub_40A180(v3);
						if ((uint8_t)v2) {
							sub_40A250();
							LOBYTE(v2) = sub_40A1F0(1);
						}
					}
				}
			}
		}
		if (a2 & 0x423) {
			LOBYTE(v2) = nox_xxx_netReportPlayerStatus_417630(a1);
		}
	}
	return v2;
}
// 417577: variable 'v2' is possibly undefined

//----- (004175C0) --------------------------------------------------------
char* nox_xxx_sendAllClientStatus_4175C0(int a1) {
	char* result; // eax
	int i;        // esi
	int v3;       // [esp-18h] [ebp-24h]
	char v4[7];   // [esp+4h] [ebp-8h]

	result = nox_common_playerInfoGetFirst_416EA0();
	for (i = (int)result; result; i = (int)result) {
		v4[0] = 106;
		*(uint16_t*)&v4[1] = *(uint16_t*)(i + 2060);
		v3 = *(unsigned char*)(a1 + 2064);
		*(uint32_t*)&v4[3] = *(uint32_t*)(i + 3680) & 0x423;
		nox_xxx_netSendPacket1_4E5390(v3, (int)v4, 7, 0, 0);
		result = nox_common_playerInfoGetNext_416EE0(i);
	}
	return result;
}

//----- (00417630) --------------------------------------------------------
int nox_xxx_netReportPlayerStatus_417630(nox_playerInfo* pl) {
	int a1 = pl;
	short v1;   // cx
	int v2;     // edx
	char v4[7]; // [esp+0h] [ebp-8h]

	v1 = *(uint16_t*)(a1 + 2060);
	v2 = *(uint32_t*)(a1 + 3680) & 0x423;
	v4[0] = 106;
	*(uint16_t*)&v4[1] = v1;
	*(uint32_t*)&v4[3] = v2;
	return nox_xxx_netSendPacket1_4E5390(255, (int)v4, 7, 0, 0);
}

//----- (00417680) --------------------------------------------------------
void nox_xxx_cliPlayerRespawn_417680(int a1, char a2) {
	int v2;       // esi
	int v3;       // eax
	int v4;       // ecx
	uint32_t* v5; // edi
	int v6;       // eax
	int v7;       // ecx
	uint32_t* v8; // edi
	int v9;       // eax
	int v10;      // edi
	int v11;      // eax
	int v12;      // edx
	char v13;     // dl
	int v14;      // eax
	char v15;     // al
	int v16;      // eax
	int v17;      // eax
	int v18;      // eax
	int v19;      // eax
	char v20;     // cl
	int v21;      // eax
	int v22;      // eax
	char v23;     // cl
	int v24;      // eax
	int v25;      // [esp-Ch] [ebp-18h]
	int v26;      // [esp-Ch] [ebp-18h]

	v2 = a1;
	if (!a1) {
		return;
	}
	if (!nox_common_gameFlags_check_40A5C0(1)) {
		*(uint32_t*)(v2 + 4) = 0;
	}
	v3 = v2 + 2328;
	v4 = 27;
	do {
		v5 = (uint32_t*)v3;
		*(uint32_t*)(v3 - 4) = 0;
		v3 += 24;
		*v5 = 0;
		--v4;
		v5[1] = 0;
		v5[2] = 0;
		v5[3] = 0;
	} while (v4);
	if (!nox_common_gameFlags_check_40A5C0(1)) {
		*(uint32_t*)v2 = 0;
	}
	v6 = v2 + 2976;
	v7 = 26;
	do {
		v8 = (uint32_t*)v6;
		*(uint32_t*)(v6 - 4) = 0;
		v6 += 24;
		*v8 = 0;
		--v7;
		v8[1] = 0;
		v8[2] = 0;
		v8[3] = 0;
	} while (v7);
	v9 = nox_xxx_modifGetIdByName_413290("UserColor1");
	v10 = nox_xxx_modifGetDescById_413330(v9);
	if (!v10) {
		return;
	}
	if (*(uint8_t*)(v2 + 2251) || nox_common_gameFlags_check_40A5C0(2048)) {
		LOBYTE(a1) = -1;
		v11 = nox_xxx_modifGetDescById_413330(*(uint32_t*)(v10 + 4) + *(unsigned char*)(v2 + 2269));
		v12 = *(unsigned char*)(v2 + 2270);
		BYTE1(a1) = *(uint8_t*)(v11 + 4);
		BYTE2(a1) = *((uint8_t*)nox_xxx_modifGetDescById_413330(*(uint32_t*)(v10 + 4) + v12) + 4);
		HIBYTE(a1) = -1;
		if (a2 & 1) {
			nox_xxx_clientEquipWeaponArmor_417AA0(82, *(uint32_t*)(v2 + 2060), 1024, (int)&a1);
		}
	}
	LOBYTE(a1) = -1;
	BYTE1(a1) = *((uint8_t*)nox_xxx_modifGetDescById_413330(*(uint32_t*)(v10 + 4) + *(unsigned char*)(v2 + 2268)) + 4);
	HIWORD(a1) = -1;
	if (a2 & 2) {
		nox_xxx_clientEquipWeaponArmor_417AA0(82, *(uint32_t*)(v2 + 2060), 4, (int)&a1);
	}
	v13 = *((uint8_t*)nox_xxx_modifGetDescById_413330(*(uint32_t*)(v10 + 4) + *(unsigned char*)(v2 + 2272)) + 4);
	v14 = *(unsigned char*)(v2 + 2271);
	LOBYTE(a1) = v13;
	BYTE1(a1) = *((uint8_t*)nox_xxx_modifGetDescById_413330(*(uint32_t*)(v10 + 4) + v14) + 4);
	HIWORD(a1) = -1;
	if (a2 & 4) {
		nox_xxx_clientEquipWeaponArmor_417AA0(82, *(uint32_t*)(v2 + 2060), 1, (int)&a1);
	}
	v15 = *(uint8_t*)(v2 + 2251);
	a1 = -1;
	if (v15 == 1) {
		if (nox_common_gameFlags_check_40A5C0(2048)) {
			if (a2 & 8) {
				v16 = nox_xxx_modifGetIdByName_413290("ArmorQuality1");
				LOBYTE(a1) = *((uint8_t*)nox_xxx_modifGetDescById_413330(v16) + 4);
				nox_xxx_clientEquipWeaponArmor_417AA0(80, *(uint32_t*)(v2 + 2060), 0x8000, (int)&a1);
			}
		} else if (nox_common_gameFlags_check_40A5C0(4096)) {
			a1 = -1;
			v17 = nox_xxx_modifGetIdByName_413290("Replenishment1");
			BYTE2(a1) = *((uint8_t*)nox_xxx_modifGetDescById_413330(v17) + 4);
			nox_xxx_clientEquipWeaponArmor_417AA0(80, *(uint32_t*)(v2 + 2060), 0x10000, (int)&a1);
		} else if (a2 & 0x10) {
			nox_xxx_clientEquipWeaponArmor_417AA0(79, *(uint32_t*)(v2 + 2060), 0x4000, (int)&a1);
		}
	}
	if (!*(uint8_t*)(v2 + 2251)) {
		if (nox_common_gameFlags_check_40A5C0(2048)) {
			if (a2 & 0x20) {
				v18 = nox_xxx_modifGetIdByName_413290("ArmorQuality1");
				LOBYTE(a1) = *((uint8_t*)nox_xxx_modifGetDescById_413330(v18) + 4);
				v19 = nox_xxx_modifGetIdByName_413290("Material1");
				v20 = *((uint8_t*)nox_xxx_modifGetDescById_413330(v19) + 4);
				v21 = *(uint32_t*)(v2 + 2060);
				LOBYTE(a1) = v20;
				nox_xxx_clientEquipWeaponArmor_417AA0(80, v21, 256, (int)&a1);
			}
		} else if (nox_common_gameFlags_check_40A5C0(4096)) {
			v25 = *(uint32_t*)(v2 + 2060);
			a1 = -1;
			nox_xxx_clientEquipWeaponArmor_417AA0(80, v25, 256, (int)&a1);
		} else {
			if (a2 & 0x40) {
				nox_xxx_clientEquipWeaponArmor_417AA0(80, *(uint32_t*)(v2 + 2060), 512, (int)&a1);
			}
			if (a2 < 0) {
				nox_xxx_clientEquipWeaponArmor_417AA0(79, *(uint32_t*)(v2 + 2060), 0x1000000, (int)&a1);
			}
		}
	}
	if (*(uint8_t*)(v2 + 2251) == 2) {
		if (nox_common_gameFlags_check_40A5C0(2048)) {
			if (a2 & 8) {
				v22 = nox_xxx_modifGetIdByName_413290("ArmorQuality1");
				v23 = *((uint8_t*)nox_xxx_modifGetDescById_413330(v22) + 4);
				v24 = *(uint32_t*)(v2 + 2060);
				LOBYTE(a1) = v23;
				nox_xxx_clientEquipWeaponArmor_417AA0(80, v24, 0x8000, (int)&a1);
			}
		} else if (nox_common_gameFlags_check_40A5C0(4096)) {
			v26 = *(uint32_t*)(v2 + 2060);
			a1 = -1;
			nox_xxx_clientEquipWeaponArmor_417AA0(80, v26, 4, (int)&a1);
		}
	}
}

//----- (00417AA0) --------------------------------------------------------
char* nox_xxx_clientEquipWeaponArmor_417AA0(char a1, int a2, int a3, int a4) {
	char* result; // eax
	char* v5;     // edx
	int v6;       // ecx
	int v7;       // esi
	char* v8;     // edi
	char* v9;     // edx
	int v10;      // ecx
	int v11;      // esi
	char* v12;    // edi

	result = nox_common_playerInfoGetByID_417040(a2);
	if (result) {
		if (a1 == 81 || a1 == 80) {
			v9 = result + 2324;
			*((uint32_t*)result + 1) |= a3;
			v10 = 0;
			while (*(uint32_t*)v9) {
				++v10;
				v9 += 24;
				if (v10 >= 27) {
					return result;
				}
			}
			*(uint32_t*)&result[24 * v10 + 2324] = a3;
			v11 = 0;
			v12 = &result[24 * v10 + 2328];
			do {
				result = (char*)nox_xxx_modifGetDescById_413330(*(unsigned char*)(v11 + a4));
				*(uint32_t*)v12 = result;
				++v11;
				v12 += 4;
			} while (v11 < 4);
		} else {
			v5 = result + 2972;
			*(uint32_t*)result |= a3;
			v6 = 0;
			while (*(uint32_t*)v5) {
				++v6;
				v5 += 24;
				if (v6 >= 26) {
					return result;
				}
			}
			*(uint32_t*)&result[24 * v6 + 2972] = a3;
			v7 = 0;
			v8 = &result[24 * v6 + 2976];
			do {
				result = (char*)nox_xxx_modifGetDescById_413330(*(unsigned char*)(v7 + a4));
				*(uint32_t*)v8 = result;
				++v7;
				v8 += 4;
			} while (v7 < 4);
		}
	}
	return result;
}

//----- (00417B80) --------------------------------------------------------
char* sub_417B80(char a1, int a2, int a3) {
	char* result; // eax
	int v4;       // ecx
	int v5;       // edx
	int v6;       // ecx
	char* i;      // edx
	int v8;       // edx
	int v9;       // ecx
	char* j;      // edx

	result = nox_common_playerInfoGetByID_417040(a2);
	if (result) {
		v4 = ~a3;
		if (a1 == 84) {
			v5 = v4 & *((uint32_t*)result + 1);
			v6 = 0;
			*((uint32_t*)result + 1) = v5;
			for (i = result + 2324; *(uint32_t*)i != a3; i += 24) {
				if (++v6 >= 27) {
					return result;
				}
			}
			*(uint32_t*)&result[24 * v6 + 2324] = 0;
		} else {
			v8 = v4 & *(uint32_t*)result;
			v9 = 0;
			*(uint32_t*)result = v8;
			for (j = result + 2972; *(uint32_t*)j != a3; j += 24) {
				if (++v9 >= 26) {
					return result;
				}
			}
			*(uint32_t*)&result[24 * v9 + 2972] = 0;
		}
	}
	return result;
}

//----- (00417CF0) --------------------------------------------------------
int sub_417CF0() { return nox_server_teamsZzz_419030(0); }

//----- (00417DC0) --------------------------------------------------------
int sub_417DC0() { return dword_5d4594_526276; }

//----- (00417DE0) --------------------------------------------------------
char sub_417DE0() {
	char v0; // bl
	char* i; // eax

	v0 = 0;
	for (i = nox_server_teamFirst_418B10(); i; i = nox_server_teamNext_418B60((int)i)) {
		if (*((uint32_t*)i + 15)) {
			++v0;
		}
	}
	return v0;
}

//----- (00417EA0) --------------------------------------------------------
int nox_xxx_mapInfoSetCapflag_417EA0() {
	int v0; // eax
	int v1; // esi

	LOBYTE(v0) = sub_417EC0();
	v1 = v0;
	if (v0) {
		sub_455A50(2);
	}
	return v1;
}
// 417EA6: variable 'v0' is possibly undefined

//----- (00417EC0) --------------------------------------------------------
bool sub_417EC0() {
	int i; // eax

	dword_5d4594_526276 = 0;
	for (i = nox_server_getFirstObject_4DA790(); i; i = nox_server_getNextObject_4DA7A0(i)) {
		if (*(uint32_t*)(i + 8) & 0x10000000) {
			++dword_5d4594_526276;
		}
	}
	if (dword_5d4594_526276 > 0 && !nox_common_gameFlags_check_40A5C0(0x8000)) {
		sub_4181F0(0);
	}
	return dword_5d4594_526276 > 0;
}

//----- (00417F30) --------------------------------------------------------
char nox_xxx_mapInfoSetFlagball_417F30() {
	int v0; // eax

	LOBYTE(v0) = sub_417EC0();
	if (v0) {
		sub_455F60();
		LOBYTE(v0) = sub_417F50(0);
	}
	return v0;
}
// 417F37: variable 'v0' is possibly undefined

//----- (00417F50) --------------------------------------------------------
int sub_417F50(int a1) {
	int v1;       // esi
	int v2;       // eax
	int v3;       // esi
	int v4;       // ebx
	uint32_t* v6; // eax
	uint32_t* v7; // esi
	int v8;       // edi
	int i;        // edi
	float v10;    // [esp+0h] [ebp-14h]

	if (!dword_5d4594_527656) {
		dword_5d4594_527656 = nox_xxx_getNameId_4E3AA0("GameBallStart");
	}
	v1 = 0;
	v2 = nox_server_getFirstObject_4DA790();
	if (!v2) {
		return 0;
	}
	do {
		if (*(unsigned short*)(v2 + 4) == dword_5d4594_527656) {
			++v1;
		}
		v2 = nox_server_getNextObject_4DA7A0(v2);
	} while (v2);
	if (!v1) {
		return 0;
	}
	v3 = nox_common_randomInt_415FA0(0, v1 - 1);
	v4 = nox_server_getFirstObject_4DA790();
	if (!v4) {
		return 0;
	}
	while (1) {
		if (*(unsigned short*)(v4 + 4) == dword_5d4594_527656) {
			if (!v3) {
				break;
			}
			--v3;
		}
		v4 = nox_server_getNextObject_4DA7A0(v4);
		if (!v4) {
			return 0;
		}
	}
	v6 = nox_xxx_newObjectByTypeID_4E3810("GameBall");
	v7 = v6;
	if (!v6) {
		return 0;
	}
	v8 = v6[187];
	*(uint64_t*)(v8 + 8) = nox_platform_get_ticks();
	v10 = nox_xxx_gamedataGetFloat_419D40("FlagballPossDuration");
	*(uint32_t*)(v8 + 20) = nox_float2int(v10);
	*(float*)(v8 + 24) = nox_xxx_gamedataGetFloat_419D40("FlagballResetVel");
	nox_xxx_netMarkMinimapForAll_4174B0((int)v7, 1);
	nox_xxx_createAt_4DAA50((int)v7, 0, 0.0, 0.0);
	nox_xxx_unitClearOwner_4EC300((int)v7);
	sub_4EB9B0((int)v7, 0);
	sub_4E8290(0, 0);
	nox_xxx_unitMove_4E7010((int)v7, (float2*)(v4 + 56));
	v7[20] = 0;
	v7[21] = 0;
	v7[22] = 0;
	v7[25] = 0;
	if (a1) {
		for (i = nox_xxx_getFirstPlayerUnit_4DA7C0(); i; i = nox_xxx_getNextPlayerUnit_4DA7F0(i)) {
			if (*(uint32_t*)(*(uint32_t*)(*(uint32_t*)(i + 748) + 276) + 3628) == a1) {
				nox_xxx_playerCameraUnlock_4E6040(i);
				nox_xxx_playerCameraFollow_4E6060(i, (int)v7);
			}
		}
		nox_xxx_delayedDeleteObject_4E5CC0(a1);
	}
	return 1;
}

//----- (004180D0) --------------------------------------------------------
int nox_xxx_mapInfoSetKotr_4180D0() {
	int v0;   // esi
	int v1;   // edi
	char* i;  // eax
	int v3;   // ebp
	int v4;   // esi
	int v5;   // edi
	bool v6;  // zf
	char v7;  // al
	char* v8; // eax

	if (!*getMemU32Ptr(0x5D4594, 527652)) {
		*getMemU32Ptr(0x5D4594, 527652) = nox_xxx_getNameId_4E3AA0("Crown");
	}
	v0 = nox_server_getFirstObject_4DA790();
	if (v0) {
		do {
			v1 = nox_server_getNextObject_4DA7A0(v0);
			if (*(uint32_t*)(v0 + 8) & 0x10000000) {
				nox_xxx_delayedDeleteObject_4E5CC0(v0);
			}
			v0 = v1;
		} while (v1);
	}
	for (i = nox_server_teamFirst_418B10(); i; i = nox_server_teamNext_418B60((int)i)) {
		*((uint32_t*)i + 19) = 0;
	}
	v3 = 0;
	v4 = nox_server_getFirstObject_4DA790();
	if (!v4) {
		return 0;
	}
	do {
		v5 = nox_server_getNextObject_4DA7A0(v4);
		if (*(unsigned short*)(v4 + 4) != *getMemU32Ptr(0x5D4594, 527652)) {
			v4 = v5;
			continue;
		}
		++v3;
		*(uint32_t*)(*(uint32_t*)(v4 + 748) + 4) = 0;
		v6 = !nox_xxx_CheckGameplayFlags_417DA0(4);
		v7 = *(uint8_t*)(v4 + 52);
		if (v6) {
			if (v7) {
				nox_xxx_delayedDeleteObject_4E5CC0(v4);
				sub_4EC6A0(v4);
				v4 = v5;
				continue;
			}
			nox_xxx_netMarkMinimapForAll_4174B0(v4, 1);
			v4 = v5;
			continue;
		}
		if (!v7) {
			nox_xxx_delayedDeleteObject_4E5CC0(v4);
			sub_4EC6A0(v4);
			v4 = v5;
			continue;
		}
		v8 = nox_xxx_getTeamByID_418AB0(*(unsigned char*)(v4 + 52));
		if (v8) {
			*((uint32_t*)v8 + 19) = v4;
			nox_xxx_netMarkMinimapForAll_4174B0(v4, 1);
		}
		v4 = v5;
	} while (v5);
	if (!v3) {
		return 0;
	}
	return 1;
}

//----- (004181F0) --------------------------------------------------------
void sub_4181F0(int a1) {
	unsigned char v1;  // bl
	char* i;           // esi
	char* result;      // eax
	char* v4;          // esi
	int v5;            // ecx
	int v6;            // eax
	int v7;            // ecx
	bool v8;           // cc
	unsigned char v9;  // bl
	int v10;           // edi
	int* v11;          // eax
	int v12;           // edx
	int* v13;          // ebx
	char* v14;         // ebp
	int v15;           // esi
	int v16;           // edi
	int v17;           // edx
	char* v18;         // eax
	unsigned char v19; // [esp+10h] [ebp-88h]
	unsigned char v20; // [esp+10h] [ebp-88h]
	unsigned char v21; // [esp+14h] [ebp-84h]
	int v22[32];       // [esp+18h] [ebp-80h]

	v1 = 0;
	v19 = 0;
	if (a1) {
		for (i = nox_server_teamFirst_418B10(); i; i = nox_server_teamNext_418B60((int)i)) {
			sub_418D80((int)i);
		}
	}
	result = nox_common_playerInfoGetFirst_416EA0();
	v4 = result;
	if (!result) {
		return;
	}
	do {
		v5 = *((uint32_t*)v4 + 514);
		if (v5 && (*((uint32_t*)v4 + 515) != nox_player_netCode_85319C ||
				   !nox_common_getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING))) {
			v6 = *((uint32_t*)v4 + 920);
			if ((!(v6 & 1) || v6 & 0x20) && !nox_xxx_servObjectHasTeam_419130(v5 + 48)) {
				v7 = v19;
				v19 = ++v1;
				v22[v7] = *((uint32_t*)v4 + 514);
			}
		}
		result = nox_common_playerInfoGetNext_416EE0((int)v4);
		v4 = result;
	} while (result);
	if (!v1) {
		return;
	}
	v8 = v1 <= 1u;
	v9 = v19;
	if (!v8) {
		v10 = 50;
		do {
			v20 = nox_common_randomInt_415FA0(0, v9 - 1);
			v21 = nox_common_randomInt_415FA0(0, v9 - 1);
			v11 = &v22[v20];
			--v10;
			v12 = *v11;
			*v11 = v22[v21];
			v22[v21] = v12;
		} while (v10);
	}
	result = (char*)v9;
	if ((int)v9 <= 0) {
		return;
	}
	v13 = v22;
	v14 = result;
	do {
		v15 = *v13;
		v16 = *(uint32_t*)(*v13 + 748);
		if (sub_40A740()) {
			v17 = *(uint32_t*)(v16 + 276);
			result = *(char**)(v17 + 2068);
			if (result) {
				result = nox_server_teamByXxx_418AE0(*(uint32_t*)(v17 + 2068));
				if (result) {
					nox_xxx_createAtImpl_4191D0(result[57], v15 + 48, 1, *(uint32_t*)(v15 + 36), 0);
				}
			}
		} else {
			v18 = sub_4189D0();
			nox_xxx_createAtImpl_4191D0(v18[57], v15 + 48, 1, *(uint32_t*)(v15 + 36), 1);
		}
		++v13;
		--v14;
	} while (v14);
}
// 4181F0: using guessed type int var_80[32];

//----- (00418390) --------------------------------------------------------
int sub_418390() {
	if (!nox_xxx_getTeamCounter_417DD0()) {
		return 0;
	}
	nox_xxx_SetGameplayFlag_417D50(2);
	sub_4181F0(0);
	return 1;
}

//----- (004183C0) --------------------------------------------------------
int sub_4183C0() {
	int result; // eax
	int i;      // esi
	int v2;     // eax
	int v3;     // eax
	int v4;     // eax
	char* v5;   // ebp
	int v6;     // edi
	char* v7;   // ecx
	int v8;     // eax
	double v9;  // st7
	double v10; // st6
	double v11; // st5
	float v12;  // [esp+0h] [ebp-4h]

	result = nox_xxx_getFirstPlayerUnit_4DA7C0();
	for (i = result; result; i = result) {
		v2 = *(uint32_t*)(i + 748);
		if (!v2 || ((v3 = *(uint32_t*)(v2 + 276), *(uint32_t*)(v3 + 2060) != nox_player_netCode_85319C) ||
					!nox_common_getEngineFlag(NOX_ENGINE_FLAG_DISABLE_GRAPHICS_RENDERING)) &&
					   ((v4 = *(uint32_t*)(v3 + 3680), !(v4 & 1)) || v4 & 0x20)) {
			v5 = nox_xxx_getTeamByID_418AB0(*(unsigned char*)(i + 52));
			v6 = 0;
			v12 = 1000000000.0;
			v7 = nox_server_teamFirst_418B10();
			if (v7) {
				do {
					v8 = *((uint32_t*)v7 + 18);
					if (v8) {
						v9 = *(float*)(v8 + 60) - *(float*)(i + 60);
						v10 = *(float*)(v8 + 56) - *(float*)(i + 56);
						v11 = v10 * v10 + v9 * v9;
						if (v11 < v12) {
							v12 = v11;
							v6 = (int)v7;
						}
					}
					v7 = nox_server_teamNext_418B60((int)v7);
				} while (v7);
				if (v6) {
					if (v5) {
						if ((char*)v6 != v5) {
							sub_4196D0(i + 48, v6, *(uint32_t*)(i + 36), 0);
						}
					} else {
						nox_xxx_createAtImpl_4191D0(*(uint8_t*)(v6 + 57), i + 48, 1, *(uint32_t*)(i + 36), 0);
					}
				}
			}
		}
		result = nox_xxx_getNextPlayerUnit_4DA7F0(i);
	}
	return result;
}

//----- (004184D0) --------------------------------------------------------
void sub_4184D0(nox_team_t* a1p) {
	wchar2_t* a1 = a1p;
	int v1;       // edi
	uint32_t* v2; // ebp
	char v3[18];  // [esp+4h] [ebp-14h]

	if (a1) {
		sub_457230(a1);
		*(uint32_t*)&v3[6] = *((uint32_t*)a1 + 15);
		*(uint32_t*)&v3[2] = *((unsigned char*)a1 + 57);
		*(uint16_t*)v3 = 196;
		*(uint32_t*)&v3[10] = *((uint32_t*)a1 + 13);
		v3[14] = 0;
		v3[16] = *((uint8_t*)a1 + 56);
		v3[17] = *((uint8_t*)a1 + 68);
		if (nox_common_gameFlags_check_40A5C0(512)) {
			v3[14] = 1;
		}
		v3[15] = nox_wcslen(a1);
		v1 = 2 * (unsigned char)v3[15];
		v2 = calloc(1, v1 + 18);
		*(uint64_t*)v2 = *(uint64_t*)v3;
		v2[2] = *(uint32_t*)&v3[8];
		v2[3] = *(uint32_t*)&v3[12];
		*((uint16_t*)v2 + 8) = *(uint16_t*)&v3[16];
		memcpy((char*)v2 + 18, a1, 2 * (unsigned char)v3[15]);
		nox_xxx_netSendPacket1_4E5390(159, (int)v2, v1 + 18, 0, 1);
		free(v2);
	}
}

//----- (004185B0) --------------------------------------------------------
int nox_xxx_wndGuiTeamCreate_4185B0() {
	int result;  // eax
	int i;       // esi
	char* v2;    // edi
	int v3;      // eax
	char v4;     // bl
	wchar2_t* v5; // eax

	nox_xxx_SetGameplayFlag_417D50(4);
	nox_xxx_teamCreate_4186D0(0);
	nox_xxx_teamCreate_4186D0(0);
	result = nox_server_getFirstObject_4DA790();
	for (i = result; result; i = result) {
		if (*(uint32_t*)(i + 8) & 0x10000000) {
			v2 = nox_xxx_getTeamByID_418AB0(*(unsigned char*)(i + 52));
			if (v2) {
				v3 = sub_4ECBD0(i);
				v4 = v3;
				v5 = nox_server_teamTitle_418C20(v3);
				if (v5) {
					sub_418800((wchar2_t*)v2, v5, 1);
				}
				v2[56] = v4;
				sub_4184D0((wchar2_t*)v2);
				*((uint32_t*)v2 + 18) = i;
			}
		}
		result = nox_server_getNextObject_4DA7A0(i);
	}
	return result;
}

//----- (00418640) --------------------------------------------------------
int nox_xxx_teamAssignFlags_418640() {
	int result; // eax
	int i;      // esi
	char v2;    // bl
	char* v3;   // eax

	result = nox_server_getFirstObject_4DA790();
	for (i = result; result; i = result) {
		if (*(uint32_t*)(i + 8) & 0x10000000) {
			v2 = sub_4ECBD0(i);
			v3 = nox_xxx_getTeamByID_418AB0(*(unsigned char*)(i + 52));
			if (v3) {
				v3[56] = v2;
				*((uint32_t*)v3 + 18) = i;
			}
		}
		result = nox_server_getNextObject_4DA7A0(i);
	}
	return result;
}

//----- (00418690) --------------------------------------------------------
char* nox_xxx_toggleAllTeamFlags_418690(int a1) {
	char* result; // eax
	char* i;      // esi
	int v3;       // [esp-8h] [ebp-Ch]

	result = nox_server_teamFirst_418B10();
	for (i = result; result; i = result) {
		if (*((uint32_t*)i + 18)) {
			v3 = *((uint32_t*)i + 18);
			if (a1) {
				nox_xxx_objectSetOn_4E75B0(v3);
			} else {
				nox_xxx_objectSetOff_4E7600(v3);
			}
		}
		result = nox_server_teamNext_418B60((int)i);
	}
	return result;
}
