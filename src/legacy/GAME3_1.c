#include <math.h>

#include "GAME1.h"
#include "GAME1_1.h"
#include "GAME1_2.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_1.h"
#include "GAME2_2.h"
#include "GAME2_3.h"
#include "GAME3.h"
#include "GAME3_1.h"
#include "GAME3_2.h"
#include "GAME3_3.h"
#include "GAME4.h"
#include "GAME5_2.h"
#include "client__draw__animdraw.h"
#include "client__drawable__drawable.h"
#include "client__gui__guisumn.h"
#include "client__gui__servopts__advserv.h"
#include "common__net_list.h"
#include "common__system__team.h"

#include "client__gui__guiinput.h"
#include "client__gui__servopts__objlst.h"
#include "client__gui__servopts__spelllst.h"
#include "client__gui__tooltip.h"
#include "client__gui__window.h"
#include "client__shell__inputcfg__inputcfg.h"
#include "client__shell__mainmenu.h"

#include "client__draw__plasma.h"
#include "client__drawable__update__charmup.h"
#include "client__drawable__update__fireball.h"
#include "client__system__ctrlevnt.h"

#include "client__shell__optsback.h"
#include "client__video__draw_common.h"
#include "common/fs/nox_fs.h"
#include "common__strman.h"
#include "input_common.h"
#include "operators.h"

extern uint32_t dword_8531A0_2572;
extern uint32_t dword_8531A0_2576;
extern uint32_t dword_587000_180476;
extern uint32_t dword_5d4594_1319260;
extern uint32_t dword_5d4594_1321196;
extern uint32_t dword_5d4594_1313880;
extern uint32_t dword_5d4594_1320944;
extern uint32_t dword_5d4594_1320948;
extern uint32_t dword_587000_183460;
extern uint32_t dword_5d4594_1316412;
extern uint32_t dword_5d4594_1522968;
extern uint32_t dword_5d4594_1321520;
extern uint32_t dword_5d4594_1319268;
extern uint32_t dword_5d4594_1321024;
extern uint32_t dword_5d4594_1320936;
extern uint32_t dword_587000_183456;
extern uint32_t dword_5d4594_1320988;
extern uint32_t dword_5d4594_1319248;
extern uint32_t dword_587000_180480;
extern uint32_t dword_5d4594_1319236;
extern uint32_t dword_5d4594_1320972;
extern uint32_t dword_5d4594_1316704;
extern uint32_t dword_5d4594_1321208;
extern uint32_t dword_5d4594_1321800;
extern uint32_t dword_5d4594_1321224;
extern uint32_t dword_5d4594_1319056;
extern uint32_t dword_5d4594_1320932;
extern uint32_t dword_5d4594_1321032;
extern uint32_t dword_5d4594_1319264;
extern uint32_t dword_5d4594_1321044;
extern uint32_t dword_5d4594_1319232;
extern uint32_t dword_5d4594_1321216;
extern nox_window* dword_5d4594_1321236;
extern nox_window* dword_5d4594_1321240;
extern nox_window* dword_5d4594_1321248;
extern nox_window* dword_5d4594_1321244;
nox_window* dword_5d4594_1522616 = 0;
nox_window* dword_5d4594_1522620 = 0;
nox_window* dword_5d4594_1522624 = 0;
nox_window* dword_5d4594_1522628 = 0;
extern uint32_t dword_5d4594_1320992;
extern uint32_t dword_5d4594_1321204;
extern uint32_t dword_5d4594_1316408;
extern uint64_t qword_581450_9512;
extern uint64_t qword_581450_9544;
extern uint32_t dword_5d4594_1320968;
extern uint32_t dword_5d4594_1319060;
extern uint32_t dword_5d4594_1522632;
extern uint32_t dword_5d4594_1321252;
extern uint32_t dword_5d4594_1522612;
extern uint32_t dword_5d4594_1522604;
extern uint32_t dword_5d4594_1321232;
extern uint32_t dword_5d4594_1316712;
extern uint32_t dword_5d4594_1320964;
extern uint32_t dword_5d4594_1316708;
extern uint32_t dword_5d4594_1321228;
extern uint32_t dword_5d4594_1321040;
extern uint32_t dword_5d4594_1316972;
extern uint32_t dword_5d4594_1320940;
extern uint32_t nox_player_netCode_85319C;
extern int nox_win_width;
extern int nox_win_height;

extern uint32_t nox_color_white_2523948;
extern uint32_t nox_color_blue_2650684;
extern uint32_t nox_color_yellow_2589772;
extern uint32_t nox_color_violet_2598268;
extern uint32_t nox_color_black_2650656;

nox_gui_animation* nox_wnd_xxx_1522608 = 0;

void* nox_gui_itemAmount_item_1319256 = 0;
void* nox_gui_itemAmount_dialog_1319228 = 0;

//----- (004B9470) --------------------------------------------------------
int sub_4B9470(const char** a1) {
	const char* v1;    // esi
	int v2;            // edx
	unsigned char* v3; // edi

	if (!a1) {
		return 0;
	}
	v1 = *(const char**)getMemAt(0x587000, 177488);
	v2 = 0;
	if (!*getMemU32Ptr(0x587000, 177488)) {
		return 0;
	}
	v3 = getMemAt(0x587000, 177488);
	while (strcmp(*a1, v1)) {
		v1 = (const char*)*((uint32_t*)v3 + 2);
		v3 += 8;
		++v2;
		if (!v1) {
			return 0;
		}
	}
	return *getMemU32Ptr(0x587000, 177492 + 8 * v2);
}

//----- (004B94E0) --------------------------------------------------------
int sub_4B94E0(nox_drawable* dr) {
	int a1 = dr;
	int result; // eax

	if (*(uint32_t*)(a1 + 112) & 0x10000000) {
		result = sub_4B9470(*(const char***)(a1 + 436));
	} else {
		result = 0;
	}
	return result;
}

//----- (004B95D0) --------------------------------------------------------
uint32_t* sub_4B95D0(nox_drawable* dr) {
	uint32_t* result; // eax
	int v2;           // edx
	int v3;           // ecx
	uint32_t* v4;     // ebx
	int v5;           // edi
	uint8_t* v6;      // esi
	int* v7;          // edi
	uint32_t** v8;    // esi
	int v9;           // ebx

	int a1 = dr;

	result = nox_xxx_getProjectileClassById_413250(*(uint32_t*)(a1 + 108));
	v4 = result;
	if (result) {
		v5 = 1;
		v6 = result + 4;
		do {
			LOBYTE(v3) = v6[1];
			LOBYTE(v2) = *v6;
			LOBYTE(result) = *(v6 - 1);
			nox_draw_setMaterial_4340A0(v5++, (int)result, v2, v3);
			v6 += 3;
		} while (v5 < 7);
		v7 = v4 + 9;
		v8 = (uint32_t**)(a1 + 432);
		v9 = 4;
		do {
			result = *v8;
			if (*v8) {
				LOBYTE(v3) = *((uint8_t*)result + 26);
				LOBYTE(v2) = *((uint8_t*)result + 25);
				LOBYTE(result) = *((uint8_t*)result + 24);
				nox_draw_setMaterial_4340A0(*v7, (int)result, v2, v3);
			}
			++v8;
			++v7;
			--v9;
		} while (v9);
	}
	return result;
}
// 4B9604: variable 'v2' is possibly undefined
// 4B9604: variable 'v3' is possibly undefined

//----- (004B9650) --------------------------------------------------------
uint32_t* sub_4B9650(int a1) {
	uint32_t* result; // eax
	int v2;           // edx
	int v3;           // ecx
	int v4;           // edi
	uint8_t* v5;      // esi

	result = nox_xxx_getProjectileClassById_413250(a1);
	if (result) {
		v4 = 1;
		v5 = result + 4;
		do {
			LOBYTE(v3) = v5[1];
			LOBYTE(v2) = *v5;
			LOBYTE(result) = *(v5 - 1);
			nox_draw_setMaterial_4340A0(v4++, (int)result, v2, v3);
			v5 += 3;
		} while (v4 < 7);
	}
	return result;
}
// 4B9677: variable 'v2' is possibly undefined
// 4B9677: variable 'v3' is possibly undefined

//----- (004B96F0) --------------------------------------------------------
uint32_t* sub_4B96F0(nox_drawable* dr) {
	uint32_t* result; // eax
	int v2;           // edx
	int v3;           // ecx
	uint32_t* v4;     // ebx
	int v5;           // edi
	uint8_t* v6;      // esi
	int* v7;          // edi
	uint32_t** v8;    // esi
	int v9;           // ebx

	int a1 = dr;

	result = nox_xxx_equipClothFindDefByTT_413270(*(uint32_t*)(a1 + 108));
	v4 = result;
	if (result) {
		v5 = 1;
		v6 = result + 4;
		do {
			LOBYTE(v3) = v6[1];
			LOBYTE(v2) = *v6;
			LOBYTE(result) = *(v6 - 1);
			nox_draw_setMaterial_4340A0(v5++, (int)result, v2, v3);
			v6 += 3;
		} while (v5 < 7);
		v7 = v4 + 9;
		v8 = (uint32_t**)(a1 + 432);
		v9 = 4;
		do {
			result = *v8;
			if (*v8) {
				LOBYTE(v3) = *((uint8_t*)result + 26);
				LOBYTE(v2) = *((uint8_t*)result + 25);
				LOBYTE(result) = *((uint8_t*)result + 24);
				nox_draw_setMaterial_4340A0(*v7, (int)result, v2, v3);
			}
			++v8;
			++v7;
			--v9;
		} while (v9);
	}
	return result;
}
// 4B9724: variable 'v2' is possibly undefined
// 4B9724: variable 'v3' is possibly undefined

//----- (004BA670) --------------------------------------------------------
void sub_4BA670(int a1, int a2, int a3, int a4, int a5) {
	int v5;             // ebx
	int v6;             // edi
	int v7;             // esi
	int v8;             // eax
	double v9;          // st7
	double v10;         // st6
	double v11;         // st5
	double v12;         // st4
	double v13;         // st5
	double v14;         // st5
	double v15;         // rt0
	int v16;            // eax
	int v17;            // ecx
	int v18;            // esi
	int v19;            // edx
	unsigned char* v20; // eax
	float v21;          // [esp-Ch] [ebp-44h]
	int2 a4a;           // [esp+18h] [ebp-20h]
	int2 a3a;           // [esp+20h] [ebp-18h]
	int2 a2a;           // [esp+28h] [ebp-10h]
	int2 a1a;           // [esp+30h] [ebp-8h]
	float v26;          // [esp+40h] [ebp+8h]

	v5 = a5;
	v6 = a4 - a2;
	v7 = a5 - a3;
	v8 = sub_48C6B0(a4 - a2, a5 - a3);
	dword_5d4594_1316408 = v8 / 40 + 1;
	if (v8 / 40 + 2 >= 30) {
		dword_5d4594_1316408 = 28;
	}
	a1a.field_0 = a2;
	v9 = *getMemFloatPtr(0x587000, 194136 + 8 * a1);
	v10 = *getMemFloatPtr(0x587000, 194140 + 8 * a1);
	v11 = (double)(a4 - a2);
	v12 = (double)v7;
	a1a.field_4 = a3;
	a2a.field_0 = a4;
	a2a.field_4 = v5;
	*(float*)&dword_5d4594_1313880 = v12;
	v26 = sqrt(v12 * *(float*)&dword_5d4594_1313880 + v11 * v11) + 0.0099999998;
	*getMemFloatPtr(0x5D4594, 1313876) = v11 / v26;
	v13 = *(float*)&dword_5d4594_1313880 / v26;
	*(float*)&dword_5d4594_1313880 = v13;
	v14 = v13 * v10 + *getMemFloatPtr(0x5D4594, 1313876) * v9;
	if (v14 < 0.0) {
		v14 = v14 * 0.2;
	}
	v15 = (1.0 - v14) * (double)v8 * 2.3;
	*getMemFloatPtr(0x5D4594, 1313868) = v9 * v15;
	*getMemFloatPtr(0x5D4594, 1313872) = v10 * v15;
	a3a.field_0 = nox_float2int(*getMemFloatPtr(0x5D4594, 1313868));
	v16 = nox_float2int(*getMemFloatPtr(0x5D4594, 1313872));
	a4a.field_0 = v6;
	v17 = 0;
	a3a.field_4 = v16;
	a4a.field_4 = v7;
	a5 = 0;
	do {
		if (v17) {
			*getMemFloatPtr(0x5D4594, 1313856 + 4 * v17) = *getMemFloatPtr(0x5D4594, 1313856 + 4 * v17) + 0.25;
		} else {
			*getMemFloatPtr(0x5D4594, 1313856) = *getMemFloatPtr(0x5D4594, 1313856) + 0.2;
		}
		v18 = dword_5d4594_1316408;
		if (*getMemFloatPtr(0x5D4594, 1313856 + 4 * v17) >= 1.0) {
			v19 = dword_5d4594_1316408 + 1;
			if (dword_5d4594_1316408 + 1 > 0) {
				v20 = getMemAt(0x5D4594, 1313872 + 28 * (v19 + 30 * v17));
				do {
					*((uint32_t*)v20 + 6) = *((uint32_t*)v20 - 1);
					*((uint32_t*)v20 + 7) = *(uint32_t*)v20;
					*((uint32_t*)v20 + 8) = *((uint32_t*)v20 + 1);
					*((uint32_t*)v20 + 9) = *((uint32_t*)v20 + 2);
					*((uint32_t*)v20 + 5) = *((uint32_t*)v20 - 2);
					v20 -= 28;
					--v19;
				} while (v19);
			}
			*getMemU32Ptr(0x5D4594, 1313856 + 4 * v17) = 0;
			*getMemU32Ptr(0x5D4594, 1313908 + 840 * v17) = 0;
		}
		v21 = *getMemFloatPtr(0x5D4594, 1313856 + 4 * v17);
		dword_5d4594_1316412 = 0;
		sub_4BEDE0(&a1a, &a2a, &a3a, &a4a, v18, v21, sub_4BA8B0, (int)&a5);
		v17 = ++a5;
	} while (a5 < 3);
}
// 4BA6B4: variable 'v8' is possibly undefined
// 4BAA32: variable 'v12' is possibly undefined
// 4BAA46: variable 'v14' is possibly undefined
// 4BAA70: variable 'v16' is possibly undefined
// 4BAA84: variable 'v18' is possibly undefined

//----- (004BAB30) --------------------------------------------------------
int nox_xxx_prepareLightningEffects_4BAB30() {
	int result; // eax

	*getMemU32Ptr(0x5D4594, 1316464) = nox_color_rgb_4344A0(255, 255, 255);
	*getMemU32Ptr(0x5D4594, 1316488) = nox_color_rgb_4344A0(128, 128, 255);
	*getMemU32Ptr(0x5D4594, 1316424) = nox_color_rgb_4344A0(128, 128, 255);
	*getMemU32Ptr(0x5D4594, 1316428) = nox_color_rgb_4344A0(64, 64, 255);
	*getMemU32Ptr(0x5D4594, 1316516) = nox_color_rgb_4344A0(200, 200, 255);
	*getMemU32Ptr(0x5D4594, 1316512) = nox_color_rgb_4344A0(128, 128, 255);
	*getMemU32Ptr(0x5D4594, 1316496) = nox_color_rgb_4344A0(255, 255, 255);
	*getMemU32Ptr(0x5D4594, 1316468) = nox_color_rgb_4344A0(255, 255, 0);
	*getMemU32Ptr(0x5D4594, 1316460) = nox_color_rgb_4344A0(30, 160, 30);
	*getMemU32Ptr(0x5D4594, 1316444) = nox_color_rgb_4344A0(60, 140, 60);
	*getMemU32Ptr(0x5D4594, 1316504) = nox_color_rgb_4344A0(40, 225, 40);
	*getMemU32Ptr(0x5D4594, 1316480) = nox_color_rgb_4344A0(150, 220, 150);
	*getMemU32Ptr(0x5D4594, 1316520) = nox_xxx_getTTByNameSpriteMB_44CFC0("BlueSpark");
	*getMemU32Ptr(0x5D4594, 1316524) = nox_xxx_getTTByNameSpriteMB_44CFC0("YellowSpark");
	result = nox_xxx_getTTByNameSpriteMB_44CFC0("GreenSpark");
	*getMemU32Ptr(0x5D4594, 1316528) = result;
	return result;
}
// 4BBC74: variable 'v10' is possibly undefined
// 4BBC74: variable 'v9' is possibly undefined
// 4BC002: variable 'v5' is possibly undefined
// 4BC002: variable 'v3' is possibly undefined

//----- (004BC720) --------------------------------------------------------
int sub_4BC720(int a1) {
	int result; // eax

	result = *(uint32_t*)(a1 + 304);
	*(uint32_t*)(a1 + 432) = *(unsigned char*)(result + 27) * (*(unsigned char*)(result + 32) + 1);
	return result;
}

//----- (004BC920) --------------------------------------------------------
int nox_xxx_updDrawMonsterGen_4BC920() { return 1; }
// 4BCABD: variable 'v5' is possibly undefined
// 4BCABD: variable 'v4' is possibly undefined

//----- (004BD280) --------------------------------------------------------
uint32_t* sub_4BD280(int a1, int a2) {
	int v2;           // esi
	uint32_t* result; // eax
	uint32_t* v4;     // ecx
	int v5;           // edi

	v2 = a2 + 4;
	result = calloc(1, a1 * (a2 + 4) + 4);
	if (result) {
		v4 = result + 1;
		*result = result + 1;
		if (a1 != 1) {
			v5 = a1 - 1;
			do {
				--v5;
				*v4 = (char*)v4 + v2;
				v4 = (uint32_t*)((char*)v4 + v2);
			} while (v5);
		}
		*v4 = 0;
	}
	return result;
}

//----- (004BD2D0) --------------------------------------------------------
void sub_4BD2D0(void* lpMem) { free(lpMem); }

//----- (004BD2E0) --------------------------------------------------------
uint32_t* sub_4BD2E0(uint32_t** a1) {
	uint32_t* result; // eax
	uint32_t* v2;     // edx

	result = *a1;
	if (*a1) {
		v2 = (uint32_t*)*result;
		++result;
		*a1 = v2;
	}
	return result;
}

//----- (004BD300) --------------------------------------------------------
int sub_4BD300(uint32_t* a1, int a2) {
	int result; // eax

	result = a2 - 4;
	*(uint32_t*)(a2 - 4) = *a1;
	*a1 = a2 - 4;
	return result;
}

//----- (004BD340) --------------------------------------------------------
uint32_t* sub_4BD340(int a1, int a2, int a3, int a4) {
	uint32_t* v4; // esi

	v4 = calloc(1, 0x1Cu);
	memset(v4, 0, 0x1Cu);
	*v4 = a1;
	v4[6] = a4;
	v4[1] = sub_4BD280(a2 / (a4 + 24), a4 + 24);
	v4[2] = sub_4BD280(a3, 84);
	nox_common_list_clear_425760(v4 + 3);
	if (v4[1] && v4[2]) {
		return v4;
	}
	sub_4BD3C0(v4);
	return 0;
}

//----- (004BD3C0) --------------------------------------------------------
void sub_4BD3C0(void* lpMem) {
	int i; // eax

	for (i = nox_common_list_getNext_425940((int*)lpMem + 3); i; i = nox_common_list_getNext_425940((int*)lpMem + 3)) {
		sub_4BD690(i);
	}
	if (*((uint32_t*)lpMem + 1)) {
		sub_4BD2D0(*((void**)lpMem + 1));
	}
	if (*((uint32_t*)lpMem + 2)) {
		sub_4BD2D0(*((void**)lpMem + 2));
	}
	free(lpMem);
}

//----- (004BD420) --------------------------------------------------------
uint32_t* sub_4BD420(int a1, int a2) {
	uint32_t* result; // eax

	result = *(uint32_t**)(a1 + 12);
	if (result == (uint32_t*)(a1 + 12)) {
		return 0;
	}
	while (result[4] != a2 || !result[5]) {
		result = (uint32_t*)*result;
		if (result == (uint32_t*)(a1 + 12)) {
			return 0;
		}
	}
	return result;
}

//----- (004BD470) --------------------------------------------------------
uint32_t* sub_4BD470(uint32_t** a1, int a2) {
	uint32_t* v2;  // eax
	uint32_t* v3;  // edi
	uint32_t* v5;  // ebx
	uint32_t* v6;  // eax
	uint32_t* v7;  // ebp
	char* v8;      // edi
	signed int v9; // eax
	uint32_t* v10; // [esp+18h] [ebp+8h]

	v2 = sub_4BD420((int)a1, a2);
	v3 = v2;
	if (v2) {
		nox_common_list_remove_425920((uint32_t**)v2);
		sub_425900(a1 + 3, v3);
		return v3;
	}
	if (!sub_486B60((int)*a1, a2)) {
		return 0;
	}
	v5 = sub_4BD2E0((uint32_t**)a1[2]);
	if (!v5) {
		sub_4BD600((int)a1);
		v5 = sub_4BD2E0((uint32_t**)a1[2]);
		if (!v5) {
			sub_486E00((int)*a1);
			return 0;
		}
	}
	v5[4] = a2;
	v5[13] = a1;
	sub_425770(v5);
	v5[3] = 0;
	sub_487C30(v5 + 6);
	nullsub_10(v5 + 14);
	v5[11] = v5 + 14;
	v6 = (uint32_t*)(*a1)[71];
	v10 = (uint32_t*)(*a1)[71];
	if (!v6) {
		sub_486AA0(*a1, v5[4], v5 + 14);
		sub_425900(a1 + 3, v5);
		v5[5] = 1;
		sub_486E00((int)*a1);
		return v5;
	}
	while (1) {
		v7 = a1[6];
		if ((int)v7 > (int)v6) {
			v7 = v6;
		}
		v8 = (char*)sub_4BD2E0((uint32_t**)a1[1]);
		if (!v8) {
			while (sub_4BD600((int)a1)) {
				v8 = (char*)sub_4BD2E0((uint32_t**)a1[1]);
				if (v8) {
					goto LABEL_17;
				}
			}
			sub_4BD690((int)v5);
			return 0;
		}
	LABEL_17:
		sub_487D30(v8, (int)(v8 + 24), (int)v7);
		sub_487C50((int)(v5 + 6), v8);
		v9 = sub_486DB0((int)*a1, v8 + 24, (signed int)v7);
		if ((uint32_t*)v9 != v7) {
			sub_4BD690((int)v5);
			return 0;
		}
		v10 = (uint32_t*)((char*)v10 - v9);
		if (!v10) {
			sub_486AA0(*a1, v5[4], v5 + 14);
			sub_425900(a1 + 3, v5);
			v5[5] = 1;
			sub_486E00((int)*a1);
			return v5;
		}
		v6 = v10;
	}
}
// 487CF0: using guessed type void  nullsub_10(uint32_t);

//----- (004BD600) --------------------------------------------------------
int sub_4BD600(int a1) {
	int v1; // esi

	v1 = sub_425960(a1 + 12);
	if (!v1) {
		return 0;
	}
	while (sub_4BD680(v1)) {
		v1 = sub_425960(v1);
		if (!v1) {
			return 0;
		}
	}
	sub_4BD690(v1);
	return 1;
}

//----- (004BD650) --------------------------------------------------------
int sub_4BD650(int a1) {
	int result; // eax

	result = a1;
	++*(uint32_t*)(a1 + 12);
	return result;
}

//----- (004BD660) --------------------------------------------------------
int sub_4BD660(int a1) {
	int result; // eax

	result = *(uint32_t*)(a1 + 12) - 1;
	*(uint32_t*)(a1 + 12) = result;
	if (result < 0) {
		*(uint32_t*)(a1 + 12) = 0;
	}
	return result;
}

//----- (004BD680) --------------------------------------------------------
int sub_4BD680(int a1) { return *(uint32_t*)(a1 + 12); }

//----- (004BD690) --------------------------------------------------------
int sub_4BD690(int a1) {
	uint32_t** i; // esi

	if (*(uint32_t*)(a1 + 4) != a1) {
		nox_common_list_remove_425920((uint32_t**)a1);
	}
	for (i = (uint32_t**)nox_common_list_getNext_425940((int*)(a1 + 32)); i;
		 i = (uint32_t**)nox_common_list_getNext_425940((int*)(a1 + 32))) {
		nox_common_list_remove_425920(i);
		sub_487D60((int)i);
		sub_4BD300(*(uint32_t**)(*(uint32_t*)(a1 + 52) + 4), (int)i);
	}
	nullsub_9(a1 + 24);
	return sub_4BD300(*(uint32_t**)(*(uint32_t*)(a1 + 52) + 8), a1);
}
// 487CA0: using guessed type void  nullsub_9(uint32_t);

//----- (004BD710) --------------------------------------------------------
int sub_4BD710(int a1) { return a1 + 24; }

//----- (004BD720) --------------------------------------------------------
uint32_t* sub_4BD720(int a1) {
	uint32_t* v1; // esi

	v1 = calloc(1, 0x138u);
	memset(v1, 0, 0x138u);
	sub_425770(v1);
	sub_4BDC00((int)(v1 + 30));
	sub_4864A0(v1 + 44);
	sub_4BD7C0(v1);
	v1[33] = a1;
	v1[43] = *(uint32_t*)(a1 + 256);
	if (!(*(int (**)(uint32_t*))(*(uint32_t*)(a1 + 256) + 4))(v1)) {
		return v1;
	}
	if (v1) {
		sub_4BD7A0(v1);
	}
	return 0;
}

//----- (004BD7A0) --------------------------------------------------------
void sub_4BD7A0(void* lpMem) {
	(*(void (**)(void*))(*((uint32_t*)lpMem + 43) + 8))(lpMem);
	free(lpMem);
}

//----- (004BD7C0) --------------------------------------------------------
uint32_t* sub_4BD7C0(uint32_t* a1) {
	uint32_t* result; // eax

	a1[69] = sub_4BD8C0;
	a1[70] = sub_4BD940;
	a1[71] = sub_4BD9B0;
	a1[34] = 0;
	a1[35] = 0;
	a1[36] = 0;
	a1[38] = 0;
	a1[3] = 1;
	sub_4BDC00((int)(a1 + 30));
	a1[30] = 0;
	a1[29] = *getMemU32Ptr(0x5D4594, 1193340);
	a1[28] = 0;
	result = sub_4864A0(a1 + 4);
	a1[72] = 0;
	return result;
}

//----- (004BD840) --------------------------------------------------------
void sub_4BD840(int a3) {
	int v1;           // ebp
	unsigned int* v2; // edi
	uint32_t* v3;     // esi
	int result;       // eax

	v1 = *(uint32_t*)(a3 + 132);
	v2 = (unsigned int*)(a3 + 176);
	sub_4864A0((uint32_t*)(a3 + 176));
	sub_486570((unsigned int*)(a3 + 176), (uint32_t*)(a3 + 16));
	sub_486620((uint32_t*)(a3 + 16));
	if (*(uint32_t*)(a3 + 112)) {
		sub_486570(v2, *(uint32_t**)(a3 + 112));
		sub_486620(*(uint32_t**)(a3 + 112));
	}
	v3 = *(uint32_t**)(a3 + 116);
	if (v3) {
		sub_486570(v2, v3);
	}
	sub_486570(v2, (uint32_t*)(v1 + 88));
	result = *(uint32_t*)(v1 + 184);
	if (result) {
		 sub_486570(v2, *(uint32_t**)(v1 + 184));
	}
}

//----- (004BD8C0) --------------------------------------------------------
int sub_4BD8C0(int a1) {
	int (*v1)(int); // eax
	int result;     // eax
	int v3;         // eax
	int v4;         // eax

	v1 = *(int (**)(int))(a1 + 136);
	if (v1) {
		result = v1(a1);
		if (result) {
			*(uint32_t*)(a1 + 300) = 0;
			*(uint32_t*)(a1 + 304) = 0;
			*(uint32_t*)(a1 + 296) = 0;
			return result;
		}
	} else {
		if (*(uint32_t*)(a1 + 292)) {
			v3 = nox_common_list_getNext_425940(*(int**)(a1 + 292));
			*(uint32_t*)(a1 + 292) = v3;
			if (v3) {
				*(uint32_t*)(a1 + 296) = *(uint32_t*)(v3 + 12);
				v4 = *(uint32_t*)(v3 + 16);
				*(uint32_t*)(a1 + 300) = v4;
				*(uint32_t*)(a1 + 304) = v4;
				return 0;
			}
		}
		*(uint32_t*)(a1 + 300) = 0;
	}
	return 0;
}

//----- (004BD940) --------------------------------------------------------
int sub_4BD940(int a1) {
	void (*v1)(int); // eax

	if (*(uint32_t*)(a1 + 128)) {
		if (*(int*)(a1 + 128) != -1) {
			--*(uint32_t*)(a1 + 128);
		}
		sub_4BDB90((uint32_t*)a1, *(uint32_t**)(a1 + 288));
	} else {
		sub_4BDB90((uint32_t*)a1, 0);
	}
	v1 = *(void (**)(int))(a1 + 140);
	if (v1) {
		v1(a1);
	}
	if (*(uint32_t*)(a1 + 288)) {
		(*(void (**)(int))(*(uint32_t*)(a1 + 172) + 36))(a1);
	}
	return 0;
}

//----- (004BD9B0) --------------------------------------------------------
int sub_4BD9B0(uint32_t* a2) {
	int v1;               // eax
	int (*v2)(uint32_t*); // eax
	int result;           // eax

	a2[72] = 0;
	v1 = a2[31];
	LOBYTE(v1) = v1 & 0xFA;
	a2[31] = v1;
	a2[32] = 0;
	sub_4864A0(a2 + 4);
	v2 = (int (*)(uint32_t*))a2[36];
	if (v2) {
		result = v2(a2);
	} else {
		result = 0;
	}
	return result;
}

//----- (004BDA60) --------------------------------------------------------
void sub_4BDA60(void* lpMem) {
	sub_4BDA80((int)lpMem);
	sub_486E90((int)lpMem);
	sub_4BD7A0(lpMem);
}

//----- (004BDA80) --------------------------------------------------------
int sub_4BDA80(int a1) {
	int result = 0;      // eax
	int (*result2)(int); // eax

	if (*(uint8_t*)(a1 + 124) & 5) {
		(*(void (**)(int))(*(uint32_t*)(a1 + 172) + 16))(a1);
	}
	result2 = *(int (**)(int))(a1 + 148);
	if (result2) {
		result = (int)result2(a1);
	}
	*(uint32_t*)(a1 + 288) = 0;
	return result;
}

//----- (004BDB20) --------------------------------------------------------
int sub_4BDB20(int a1) {
	int result; // eax

	result = a1;
	*(uint32_t*)(a1 + 124) |= 0x10u;
	return result;
}

//----- (004BDB30) --------------------------------------------------------
int sub_4BDB30(int a1) {
	int result; // eax

	result = a1;
	*(uint32_t*)(a1 + 124) &= 0xFFFFFFEF;
	return result;
}

//----- (004BDB40) --------------------------------------------------------
int sub_4BDB40(int a2) {
	int result; // eax

	if (*(uint8_t*)(a2 + 124) & 5) {
		return -2146500608;
	}
	if (!*(uint32_t*)(a2 + 288)) {
		return -2147024896;
	}
	sub_486520((unsigned int*)(a2 + 16));
	sub_4BD840(a2);
	result = (*(int (**)(int))(*(uint32_t*)(a2 + 172) + 12))(a2);
	if (!result) {
		*(uint32_t*)(a2 + 124) |= 1u;
	}
	return result;
}

//----- (004BDB90) --------------------------------------------------------
void sub_4BDB90(uint32_t* a1, uint32_t* a2) {
	int v2;       // eax
	int v3;       // eax
	uint32_t* v4; // eax
	int v5;       // eax

	a1[72] = a2;
	if (a2) {
		v2 = sub_487C80((int)a2);
		a1[73] = v2;
		if (v2) {
			a1[74] = *(uint32_t*)(v2 + 12);
			v3 = *(uint32_t*)(v2 + 16);
			a1[75] = v3;
			a1[76] = v3;
			*a2 = 0;
		} else {
			v4 = (uint32_t*)a1[72];
			a1[74] = *v4;
			v5 = v4[1];
			a1[75] = v5;
			a1[76] = v5;
		}
	}
}

//----- (004BDC00) --------------------------------------------------------
int sub_4BDC00(int a1) {
	int result; // eax

	result = a1;
	*(uint32_t*)(a1 + 8) = 0;
	*(uint32_t*)(a1 + 4) = 0;
	return result;
}

//----- (004BDC10) --------------------------------------------------------
int nox_xxx_loadAdvancedWnd_4BDC10(int* a1) {
	dword_5d4594_1316708 = nox_new_window_from_file("advanced.wnd", nox_xxx_windowAdvancedServProc_4BDDB0);
	sub_46B120(*(uint32_t**)&dword_5d4594_1316708, 0);
	sub_46C690(*(int*)&dword_5d4594_1316708);
	nox_xxx_wndShowModalMB_46A8C0(*(int*)&dword_5d4594_1316708);
	nox_xxx_wndSetWindowProc_46B300(*(int*)&dword_5d4594_1316708, sub_4BDDA0);
	nox_xxx_windowFocus_46B500(*(int*)&dword_5d4594_1316708);
	return sub_4BDC70(a1);
}

//----- (004BDC70) --------------------------------------------------------
int sub_4BDC70(int* a1) {
	uint32_t* v1; // eax
	uint32_t* v2; // eax
	uint32_t* v3; // eax

	if (nox_common_gameFlags_check_40A5C0(1)) {
		v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316708, 10167);
		v1[9] |= 4u;
		dword_5d4594_1316704 = 0;
	} else {
		v2 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316708, 10164);
		v2[9] |= 4u;
		v3 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316708, 10167);
		nox_xxx_wnd_46ABB0((int)v3, 0);
		dword_5d4594_1316704 = 1;
	}
	sub_453F70(a1 + 6);
	sub_4535E0(a1 + 11);
	sub_4535F0(a1[12]);
	return sub_4BDD10();
}

//----- (004BDD10) --------------------------------------------------------
int sub_4BDD10() {
	int v1;   // eax
	char* v2; // eax

	switch (dword_5d4594_1316704) {
	case 0:
		v2 = sub_4165B0();
		v1 = sub_4CEBA0(*(int*)&dword_5d4594_1316708, v2);
		dword_5d4594_1316712 = v1;
		break;
	case 1:
		dword_5d4594_1316712 = nox_xxx_guiSpelllistLoad_453850(*(int*)&dword_5d4594_1316708);
		return nox_xxx_windowFocus_46B500(*(int*)&dword_5d4594_1316712);
	case 2:
		v1 = nox_xxx_guiObjlistLoad_4530C0(*(int*)&dword_5d4594_1316708, 0x1000000);
		dword_5d4594_1316712 = v1;
		break;
	case 3:
		v1 = nox_xxx_guiObjlistLoad_4530C0(*(int*)&dword_5d4594_1316708, 0x2000000);
		dword_5d4594_1316712 = v1;
		break;
	default:
		return nox_xxx_windowFocus_46B500(*(int*)&dword_5d4594_1316712);
	}
	return nox_xxx_windowFocus_46B500(*(int*)&dword_5d4594_1316712);
}

//----- (004BDDA0) --------------------------------------------------------
int sub_4BDDA0() { return 1; }

//----- (004BDDB0) --------------------------------------------------------
int nox_xxx_windowAdvancedServProc_4BDDB0(int a1, int a2, int* a3, int a4) {
	int v3;     // esi
	int result; // eax
	char* v5;   // ebx

	if (a2 == 23) {
		return 1;
	}
	if (a2 != 16391) {
		return 1;
	}
	v3 = nox_xxx_wndGetID_46B0A0(a3);
	nox_xxx_clientPlaySoundSpecial_452D80(766, 100);
	switch (v3) {
	case 10148:
		v5 = sub_4165B0();
		memcpy(v5 + 24, sub_453F90(), 0x14u);
		*((uint32_t*)v5 + 11) = *(uint32_t*)sub_453600();
		*((uint32_t*)v5 + 12) = sub_453610();
		sub_4BDF30();
		return 1;
	case 10164:
		if (dword_5d4594_1316712) {
			nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1316712);
			dword_5d4594_1316712 = 0;
		}
		dword_5d4594_1316704 = 1;
		sub_4BDD10();
		result = 1;
		break;
	case 10165:
		if (dword_5d4594_1316712) {
			nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1316712);
			dword_5d4594_1316712 = 0;
		}
		dword_5d4594_1316704 = 2;
		sub_4BDD10();
		result = 1;
		break;
	case 10166:
		if (dword_5d4594_1316712) {
			nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1316712);
			dword_5d4594_1316712 = 0;
		}
		dword_5d4594_1316704 = 3;
		sub_4BDD10();
		result = 1;
		break;
	case 10167:
		if (dword_5d4594_1316712) {
			nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1316712);
			dword_5d4594_1316712 = 0;
		}
		dword_5d4594_1316704 = 0;
		sub_4BDD10();
		result = 1;
		break;
	default:
		return 1;
	}
	return result;
}

//----- (004BDF30) --------------------------------------------------------
int sub_4BDF30() {
	int result; // eax

	result = dword_5d4594_1316708;
	if (dword_5d4594_1316708) {
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1316708);
		nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1316708);
		dword_5d4594_1316708 = 0;
		dword_5d4594_1316712 = 0;
		result = nox_xxx_windowFocus_46B500(0);
	}
	return result;
}

//----- (004BDF70) --------------------------------------------------------
int sub_4BDF70(int* a1) {
	int result; // eax

	result = *(int*)&dword_5d4594_1316708;
	if (dword_5d4594_1316708) {
		result = sub_4BDF90(a1);
	}
	return result;
}

//----- (004BDF90) --------------------------------------------------------
int sub_4BDF90(int* a1) {
	int result = 0;       // eax
	int (*result2)(void); // eax

	sub_453F70(a1 + 6);
	sub_4535E0(a1 + 11);
	sub_4535F0(a1[12]);
	result2 = *(int (**)(void))getMemAt(0x587000, 180016 + 4 * dword_5d4594_1316704);
	if (result2) {
		result = result2();
	}
	return result;
}

//----- (004BDFD0) --------------------------------------------------------
int sub_4BDFD0() {
	char* v0;     // eax
	int v1;       // esi
	int v2;       // edi
	uint32_t* v3; // eax
	uint32_t* v4; // eax
	int v5;       // eax
	uint32_t* v6; // esi
	uint32_t* v7; // eax
	int v9;       // [esp+8h] [ebp-8h]
	int v10;      // [esp+Ch] [ebp-4h]

	v0 = sub_416640();
	v1 = nox_strman_get_lang_code();
	v2 = (int)v0;
	if (nox_xxx_guiFontHeightMB_43F320(0) > 10) {
		v1 = 2;
	}
	if (0) {
		v3 = nox_new_window_from_file(*(const char**)getMemAt(0x587000, 180088 + 4 * v1), sub_4BE330);
	} else {
		v3 = nox_new_window_from_file(*(const char**)getMemAt(0x587000, 180048 + 4 * v1), sub_4BE330);
	}
	dword_5d4594_1316972 = v3;
	sub_46B120(v3, 0);
	sub_46C690(*(int*)&dword_5d4594_1316972);
	nox_xxx_wndShowModalMB_46A8C0(*(int*)&dword_5d4594_1316972);
	nox_xxx_windowFocus_46B500(*(int*)&dword_5d4594_1316972);
	nox_xxx_wndSetWindowProc_46B300(*(int*)&dword_5d4594_1316972, sub_4BE320);
	v4 = nox_xxx_wndGetChildByID_46B0C0(0, 10100);
	nox_gui_getWindowOffs_46AA20((int)v4, &v10, &v9);
	if (0) {
		v5 = v9 + 55;
	} else {
		v5 = v9 + 80;
	}
	nox_window_setPos_46A9B0(*(uint32_t**)&dword_5d4594_1316972, v10 + 15, v5);
	v6 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2104);
	nox_xxx_wnd_46B280((int)v6, *(int*)&dword_5d4594_1316972);
	nox_xxx_wndSetProc_46B2C0((int)v6, sub_4BE330);
	if (0) {
		v7 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2119);
		nox_xxx_wnd_46B280((int)v7, *(int*)&dword_5d4594_1316972);
	}
	return sub_4BE120(v2);
}

//----- (004BE120) --------------------------------------------------------
int sub_4BE120(int a1) {
	uint32_t* v1;    // eax
	int v2;          // ecx
	unsigned int v3; // ecx
	uint32_t* v4;    // eax
	int v5;          // ecx
	unsigned int v6; // ecx
	uint32_t* v7;    // esi
	uint32_t* v8;    // eax
	uint32_t* v9;    // eax
	uint32_t* v10;   // eax
	uint32_t* v11;   // eax
	uint32_t* v12;   // eax
	int result;      // eax
	uint32_t* v14;   // eax
	int v15;         // [esp-14h] [ebp-40h]
	wchar2_t v16[16]; // [esp+Ch] [ebp-20h]

	v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2102);
	v2 = v1[9];
	if (*(uint32_t*)(a1 + 58)) {
		v3 = v2 | 4;
	} else {
		v3 = v2 & 0xFFFFFFFB;
	}
	v1[9] = v3;
	v4 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2103);
	v5 = v4[9];
	if (*(uint32_t*)(a1 + 62)) {
		v6 = v5 | 4;
	} else {
		v6 = v5 & 0xFFFFFFFB;
	}
	v4[9] = v6;
	v7 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2110);
	nox_swprintf(v16, L"%d", *(uint32_t*)(a1 + 70));
	nox_window_call_field_94((int)v7, 16414, (int)v16, -1);
	switch (*(uint32_t*)(a1 + 66)) {
	case 0:
		v8 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2110);
		nox_xxx_wnd_46ABB0((int)v8, 0);
		v15 = 2106;
		v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, v15);
		v10[9] |= 4u;
		break;
	case 1:
		v9 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2110);
		nox_xxx_wnd_46ABB0((int)v9, 0);
		v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2107);
		v10[9] |= 4u;
		break;
	case 2:
		v11 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2110);
		nox_xxx_wnd_46ABB0((int)v11, 0);
		v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2108);
		v10[9] |= 4u;
		break;
	case 3:
		v12 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2110);
		nox_xxx_wnd_46ABB0((int)v12, 1);
		v15 = 2109;
		v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, v15);
		v10[9] |= 4u;
		break;
	default:
		break;
	}
	if (0) {
		v14 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2119);
		nox_window_call_field_94((int)v14, 16394, *(uint32_t*)(a1 + 74), 0);
		result = sub_4BE2C0(*(uint32_t*)(a1 + 74));
	}
	return 0;
}

//----- (004BE320) --------------------------------------------------------
int sub_4BE320() { return 1; }

//----- (004BE330) --------------------------------------------------------
int sub_4BE330(int a1, unsigned int a2, int* a3, int a4) {
	uint32_t* v4;       // esi
	char* v5;           // edi
	int result;         // eax
	const wchar2_t* v7;  // eax
	int v8;             // eax
	int v9;             // esi
	char* v10;          // eax
	char* v11;          // eax
	char* v12;          // esi
	uint32_t* v13;      // eax
	char* v14;          // esi
	uint32_t* v15;      // eax
	char* v16;          // esi
	uint32_t* v17;      // eax
	char* v18;          // esi
	uint32_t* v19;      // eax
	char* v20;          // edi
	const wchar2_t* v21; // eax
	int v22;            // esi

	if (a2 > 0x4007) {
		if (a2 == 16393) {
			sub_4BE2C0(a4);
			nox_xxx_gameSetAudioFadeoutMb_501AC0(a4);
		} else if (a2 == 16415) {
			v20 = sub_416640();
			v21 = (const wchar2_t*)nox_window_call_field_94((int)a3, 16413, 0, 0);
			if (v21) {
				if (*v21) {
					v22 = nox_wcstol(v21, 0, 10);
					if (v22 < 0) {
						v22 = 0;
					}
					if (nox_xxx_wndGetID_46B0A0(a3) == 2110) {
						*(uint32_t*)(v20 + 70) = v22;
						return 1;
					}
				}
			}
		}
		return 1;
	}
	if (a2 != 16391) {
		if (a2 != 23 && a2 == 16387) {
			v4 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, a4);
			v5 = sub_416640();
			if (!v4) {
				return 0;
			}
			if ((unsigned short)a3 == 1) {
				return 0;
			}
			v7 = (const wchar2_t*)nox_window_call_field_94((int)v4, 16413, 0, 0);
			if (v7 && *v7) {
				v8 = nox_wcstol(v7, 0, 10);
				if (v8 < 0) {
					v8 = 0;
				}
				if (a4 == 2110) {
					*(uint32_t*)(v5 + 70) = v8;
					return 1;
				}
			}
		}
		return 1;
	}
	v9 = nox_xxx_wndGetID_46B0A0(a3);
	nox_xxx_clientPlaySoundSpecial_452D80(766, 100);
	switch (v9) {
	case 2102:
		v10 = sub_416640();
		*(uint32_t*)(v10 + 58) ^= 1u;
		result = 1;
		break;
	case 2103:
		v11 = sub_416640();
		*(uint32_t*)(v11 + 62) ^= 1u;
		result = 1;
		break;
	case 2106:
		v12 = sub_416640();
		v13 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2110);
		nox_xxx_wnd_46ABB0((int)v13, 0);
		*(uint32_t*)(v12 + 66) = 0;
		sub_40A6A0(1);
		result = 1;
		break;
	case 2107:
		v14 = sub_416640();
		v15 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2110);
		nox_xxx_wnd_46ABB0((int)v15, 0);
		*(uint32_t*)(v14 + 66) = 1;
		sub_40A6A0(1);
		result = 1;
		break;
	case 2108:
		v16 = sub_416640();
		v17 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2110);
		nox_xxx_wnd_46ABB0((int)v17, 0);
		*(uint32_t*)(v16 + 66) = 2;
		sub_40A6A0(1);
		result = 1;
		break;
	case 2109:
		v18 = sub_416640();
		v19 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1316972, 2110);
		nox_xxx_wnd_46ABB0((int)v19, 1);
		*(uint32_t*)(v18 + 66) = 3;
		sub_40A6A0(1);
		result = 1;
		break;
	case 2130:
		sub_4BE610();
		result = 1;
		break;
	default:
		return 1;
	}
	return result;
}

//----- (004BE610) --------------------------------------------------------
int sub_4BE610() {
	int result; // eax

	result = dword_5d4594_1316972;
	if (dword_5d4594_1316972) {
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1316972);
		nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1316972);
		dword_5d4594_1316972 = 0;
		result = nox_xxx_windowFocus_46B500(0);
	}
	return result;
}

//----- (004BE800) --------------------------------------------------------
int sub_4BE800(int a1) {
	int result; // eax

	result = a1;
	*getMemU32Ptr(0x5D4594, 1316980) = a1;
	return result;
}

//----- (004BE810) --------------------------------------------------------
char sub_4BE810(int a1, int a2, int a3, char a4) {
	char result; // al

	*getMemU32Ptr(0x5D4594, 1316988) = a1;
	result = a4;
	*getMemU32Ptr(0x5D4594, 1316984) = a2;
	*getMemU32Ptr(0x5D4594, 1316992) = a3;
	*getMemU8Ptr(0x5D4594, 1316996) = a4;
	return result;
}

//----- (004BEAD0) --------------------------------------------------------
void sub_4BEAD0(int2* a1, int2* a2, int2* a3, int2* a4, int a5, int a6) {
	int v6;             // eax
	double v7;          // st7
	unsigned char* v8;  // ecx
	int v9;             // edi
	int v10;            // eax
	int v11;            // esi
	int v12;            // ebx
	int v13;            // edi
	int v14;            // edx
	int v15;            // ebx
	unsigned char* v16; // eax
	int v17;            // ecx
	double v18;         // st7
	double v19;         // st7
	double v20;         // st6
	double v21;         // st5
	double v22;         // st4
	double v23;         // st3
	double v24;         // st2
	double v25;         // st1
	double v26;         // st0
	int v27;            // ebp
	int v28;            // ebx
	int i;              // eax
	int v30;            // esi
	int v31;            // ecx
	int v32;            // eax
	bool v33;           // zf
	int v34;            // [esp+4h] [ebp-38h]
	int v35;            // [esp+8h] [ebp-34h]
	int v36;            // [esp+Ch] [ebp-30h]
	int v37;            // [esp+10h] [ebp-2Ch]
	int v38;            // [esp+14h] [ebp-28h]
	int v39;            // [esp+18h] [ebp-24h]
	int v40[8];         // [esp+1Ch] [ebp-20h]
	int v41;            // [esp+40h] [ebp+4h]
	int v42;            // [esp+44h] [ebp+8h]
	float v43;          // [esp+44h] [ebp+8h]
	int v44;            // [esp+48h] [ebp+Ch]
	int v45;            // [esp+4Ch] [ebp+10h]
	float v46;          // [esp+4Ch] [ebp+10h]
	int v47;            // [esp+4Ch] [ebp+10h]
	int v48;            // [esp+50h] [ebp+14h]

	v36 = a1->field_0;
	v41 = a1->field_4;
	v6 = a4->field_4;
	v34 = a4->field_0;
	v7 = 1.0 / (double)a5;
	v35 = a2->field_0;
	v39 = a2->field_4;
	v42 = a3->field_0;
	v38 = a3->field_4;
	v44 = v6;
	v8 = getMemAt(0x581450, 9876);
	v45 = 0;
	do {
		v9 = *((uint32_t*)v8 + 1);
		v10 = *((uint32_t*)v8 + 2);
		v37 = v9;
		v11 = v39 * *(uint32_t*)v8;
		v12 = v35 * *(uint32_t*)v8 + v34 * v10 + v42 * v9;
		v13 = v36;
		v14 = v41 * *((uint32_t*)v8 - 1);
		v15 = v36 * *((uint32_t*)v8 - 1) + v12;
		v8 += 16;
		v40[v45 + 4] = v15;
		*(int*)((char*)&v40[0] + v45 * 4) = v14 + v11 + v38 * v37 + v44 * v10;
		++v45;
	} while ((int)v8 < (int)getMemAt(0x581450, 9940));
	*getMemFloatPtr(0x587000, 180484) = v7;
	v16 = getMemAt(0x587000, 180468);
	v17 = 0;
	*(float*)&dword_587000_180480 = v7 * v7;
	*(float*)&dword_587000_180476 = v7 * v7 * v7;
	*getMemFloatPtr(0x587000, 180496) = *(float*)&dword_587000_180480 + *(float*)&dword_587000_180480;
	v18 = *(float*)&dword_587000_180476 * 6.0;
	*getMemFloatPtr(0x587000, 180492) = v18;
	*getMemFloatPtr(0x587000, 180508) = v18;
	v19 = (double)v40[7];
	v20 = (double)v40[6];
	v21 = (double)v40[5];
	v22 = (double)v40[4];
	v23 = (double)v40[3];
	v24 = (double)v40[2];
	v43 = (double)v40[1];
	v46 = (double)v40[0];
	do {
		v25 = v22 * *((float*)v16 - 2);
		v26 = v21 * *((float*)v16 - 1);
		v16 += 16;
		*(float*)((char*)&v40[0] + v17 * 4) = v25 + v26 + v20 * *((float*)v16 - 4) + v19 * *((float*)v16 - 3);
		++v17;
		*(float*)&v40[v17 + 3] =
			v46 * *((float*)v16 - 6) + v43 * *((float*)v16 - 5) + v24 * *((float*)v16 - 4) + v23 * *((float*)v16 - 3);
	} while ((int)v16 < (int)getMemAt(0x587000, 180532));
	if (a5 > 0) {
		v47 = a5;
		do {
			*(float*)v40 = *(float*)v40 + *(float*)&v40[1];
			*(float*)&v40[1] = *(float*)&v40[2] + *(float*)&v40[1];
			*(float*)&v40[2] = *(float*)&v40[3] + *(float*)&v40[2];
			*(float*)&v40[4] = *(float*)&v40[4] + *(float*)&v40[5];
			*(float*)&v40[5] = *(float*)&v40[6] + *(float*)&v40[5];
			*(float*)&v40[6] = *(float*)&v40[7] + *(float*)&v40[6];
			v27 = nox_float2int(*(float*)v40);
			v28 = nox_float2int(*(float*)&v40[4]);
			nox_client_drawAddPoint_49F500(v13, v41);
			nox_client_drawAddPoint_49F500(v27, v28);
			nox_client_drawSetColor_434460(*getMemIntPtr(0x5D4594, 1316980));
			nox_client_drawLineFromPoints_49E4B0();
			if (a6) {
				v48 = 0;
				for (i = v13 - v28;; i = v13 - v28) {
					v30 = v48 != 0 ? 1 : -1;
					if (i >= 0) {
						v31 = i;
					} else {
						v31 = v28 - v13;
					}
					if (i < 0) {
						i = v28 - v13;
					}
					if (v31 <= i) {
						nox_client_drawAddPoint_49F500(v30 + v27, v28);
						nox_client_drawAddPoint_49F500(v13 + v30, v41);
					} else {
						nox_client_drawAddPoint_49F500(v27, v30 + v28);
						nox_client_drawAddPoint_49F500(v13, v41 + v30);
					}
					nox_client_drawLineFromPoints_49E4B0();
					if (++v48 >= 2) {
						break;
					}
				}
			}
			if (*getMemU32Ptr(0x5D4594, 1316988)) {
				nox_client_drawAddPoint_49F500(v13, v41);
				nox_client_drawAddPoint_49F500(v27, v28);
				nox_draw_set54RGB32_434040(*getMemIntPtr(0x5D4594, 1316984));
				sub_434080(*getMemIntPtr(0x5D4594, 1316992));
				v32 = getMemByte(0x5D4594, 1316996);
				sub_49E4F0(v32);
			}
			v13 = v27;
			v33 = v47 == 1;
			v41 = v28;
			--v47;
		} while (!v33);
	}
}

//----- (004BEDE0) --------------------------------------------------------
void sub_4BEDE0(int2* a1, int2* a2, int2* a3, int2* a4, int a5, float a6, int a7, int a8) {
	int v8;             // ecx
	double v9;          // st7
	unsigned char* v10; // ecx
	int v11;            // esi
	int v12;            // eax
	int v13;            // ebx
	int v14;            // eax
	int v15;            // esi
	double v16;         // st6
	double v17;         // st7
	double v18;         // st7
	float v19;          // [esp+0h] [ebp-54h]
	float v20;          // [esp+0h] [ebp-54h]
	float v21;          // [esp+10h] [ebp-44h]
	float v22;          // [esp+10h] [ebp-44h]
	int v23;            // [esp+14h] [ebp-40h]
	float v24;          // [esp+14h] [ebp-40h]
	int v25;            // [esp+18h] [ebp-3Ch]
	float v26;          // [esp+18h] [ebp-3Ch]
	int v27;            // [esp+1Ch] [ebp-38h]
	float v28;          // [esp+1Ch] [ebp-38h]
	int v29;            // [esp+20h] [ebp-34h]
	float v30;          // [esp+20h] [ebp-34h]
	int v31;            // [esp+24h] [ebp-30h]
	float v32;          // [esp+24h] [ebp-30h]
	int v33;            // [esp+28h] [ebp-2Ch]
	float v34;          // [esp+28h] [ebp-2Ch]
	int2 v35[5];        // [esp+2Ch] [ebp-28h]
	int2* v36;          // [esp+58h] [ebp+4h]
	float v37;          // [esp+58h] [ebp+4h]
	int2* v38;          // [esp+5Ch] [ebp+8h]
	float v39;          // [esp+5Ch] [ebp+8h]
	int v40;            // [esp+60h] [ebp+Ch]
	float v41;          // [esp+60h] [ebp+Ch]
	int v42;            // [esp+64h] [ebp+10h]
	float v43;          // [esp+64h] [ebp+10h]
	float v44;          // [esp+68h] [ebp+14h]
	float v45;          // [esp+6Ch] [ebp+18h]

	v8 = a1->field_0;
	v33 = a1->field_4;
	v35[2].field_4 = a1->field_4;
	v25 = v8;
	v35[2].field_0 = v8;
	v27 = a2->field_0;
	v9 = 1.0 / (double)a5;
	v23 = a3->field_0;
	v29 = a3->field_4;
	v21 = 0.0;
	v40 = a4->field_0;
	v31 = a2->field_4;
	v42 = a4->field_4;
	v10 = getMemAt(0x581450, 9876);
	v38 = 0;
	do {
		v11 = *((uint32_t*)v10 + 1);
		v12 = *((uint32_t*)v10 + 2);
		v36 = (int2*)*((uint32_t*)v10 - 1);
		v13 = v27 * *(uint32_t*)v10 + v25 * (uint32_t)v36 + v23 * v11 + v40 * v12;
		v14 = v33 * (uint32_t)v36 + v31 * *(uint32_t*)v10 + v29 * v11 + v42 * v12;
		v10 += 16;
		*(int*)((char*)&v35[0].field_0 + (uint32_t)v38) = v13;
		*(int*)((char*)&v35[3].field_0 + (uint32_t)v38) = v14;
		v38 = (int2*)((char*)v38 + 4);
	} while ((int)v10 < (int)getMemAt(0x581450, 9940));
	v15 = a5;
	if (a5 > 0) {
		v16 = v9 * a6;
		v41 = v16 + v9;
		v45 = (double)v35[0].field_0;
		v43 = (double)v35[0].field_4;
		v34 = (double)v35[1].field_0;
		v32 = (double)v35[1].field_4;
		v28 = (double)v35[3].field_0;
		v30 = (double)v35[3].field_4;
		v26 = (double)v35[4].field_0;
		v24 = (double)v35[4].field_4;
		do {
			v17 = v41 + v21;
			v22 = v17;
			if (v17 > 1.0) {
				v22 = 1.0;
			}
			v18 = v22 * v22;
			v44 = v18;
			v37 = v18 * v22;
			v19 = v43 * v44 + v45 * v37 + v34 * v22 + v32;
			v35[0].field_0 = nox_float2int(v19);
			v20 = v30 * v44 + v28 * v37 + v26 * v22 + v24;
			v35[0].field_4 = nox_float2int(v20);
			v39 = v16;
			v21 = v22 - v39;
			((void (*)(int2*, int2*, int))a7)(&v35[2], v35, a8);
			--v15;
			v35[2].field_0 = v35[0].field_0;
			v35[2].field_4 = v35[0].field_4;
		} while (v15);
	}
}

//----- (004BF010) --------------------------------------------------------
int nox_xxx_clientReportSecondaryWeapon_4BF010(int a1) {
	char v3[3]; // [esp+0h] [ebp-4h]
	v3[0] = -32;
	*(uint16_t*)&v3[1] = nox_xxx_netGetUnitCodeCli_578B00(a1);
	return nox_xxx_netClientSend2_4E53C0(31, v3, 3, 0, 1);
}

//----- (004BF7E0) --------------------------------------------------------
short sub_4BF7E0(uint32_t* a1) {
	int v1;             // esi
	int v2;             // ebp
	int v3;             // eax
	int v4;             // edi
	int v5;             // ebx
	int v6;             // eax
	int v7;             // eax
	int v8;             // ebx
	int v9;             // eax
	int v10;            // ebx
	int v11;            // eax
	int v13;            // [esp+10h] [ebp-4h]
	unsigned char* v14; // [esp+18h] [ebp+4h]
	unsigned char* v15; // [esp+18h] [ebp+4h]

	v1 = a1[1] + 15;
	v2 = *a1 + 11;
	v13 = a1[1] + 15;
	nox_client_drawSetColor_434460(nox_color_black_2650656);
	nox_client_drawRectFilledOpaque_49CE30(v2, v1, 200, 200);
	LOWORD(v3) = *getMemU16Ptr(0x852978, 8);
	if (*getMemU32Ptr(0x852978, 8)) {
		v4 = dword_8531A0_2576;
		if (dword_8531A0_2576) {
			nox_draw_setMaterial_4341D0(1, *(uint32_t*)(dword_8531A0_2576 + 2296));
			nox_draw_setMaterial_4341D0(2, *(uint32_t*)(v4 + 2304));
			nox_draw_setMaterial_4341D0(3, *(uint32_t*)(v4 + 2312));
			nox_draw_setMaterial_4341D0(4, *(uint32_t*)(v4 + 2308));
			nox_draw_setMaterial_4341D0(5, *(uint32_t*)(v4 + 2300));
			nox_draw_setMaterial_4341D0(6, *(uint32_t*)(v4 + 2292));
			if (*(uint32_t*)(v4 + 2292) == *(uint32_t*)(v4 + 2296)) {
				nox_client_drawImageAt_47D2C0(*getMemU32Ptr(0x973A20, 24 + 4 * *(unsigned char*)(v4 + 2252)), v2, v1);
			} else {
				nox_client_drawImageAt_47D2C0(*getMemU32Ptr(0x973A20, 16 + 4 * *(unsigned char*)(v4 + 2252)), v2, v1);
			}
			v5 = 0;
			v14 = getMemAt(0x973A20, 32 + 104 * *(unsigned char*)(v4 + 2252));
			do {
				if (*(uint32_t*)v4 & (1 << v5) && !((1 << v5) & 0x3000000)) {
					v6 = sub_415CD0((char*)(1 << v5));
					sub_4BF9F0(1 << v5, v6, v2, v13, (int)v14, v5, 0);
				}
				++v5;
			} while (v5 < 26);
			if (*(uint8_t*)v4 & 2) {
				v7 = sub_415CD0((char*)2);
				sub_4BF9F0(2, v7, v2, v13, (int)v14, 0, 1);
			}
			v8 = 0;
			do {
				if (*(uint32_t*)v4 & (1 << v8) && (1 << v8) & 0x3000000) {
					v9 = sub_415CD0((char*)(1 << v8));
					sub_4BF9F0(1 << v8, v9, v2, v13, (int)v14, v8, 0);
				}
				++v8;
			} while (v8 < 26);
			v10 = 0;
			v15 = getMemAt(0x973A20, 256 + 108 * *(unsigned char*)(v4 + 2252));
			do {
				v3 = *(uint32_t*)(v4 + 4);
				if (v3 & (1 << v10)) {
					v11 = sub_415840((char*)(1 << v10));
					LOWORD(v3) = sub_4BF9F0(1 << v10, v11, v2, v13, (int)v15, v10, 0);
				}
				++v10;
			} while (v10 < 27);
		}
	}
	return v3;
}

//----- (004BF9F0) --------------------------------------------------------
short sub_4BF9F0(int a1, int a2, int a3, int a4, int a5, int a6, int a7) {
	int v7;        // eax
	int v8;        // esi
	uint32_t* v9;  // eax
	int v10;       // edx
	int v11;       // ecx
	uint32_t* v12; // ebx
	int v13;       // ebp
	int v14;       // edi
	uint8_t* v15;  // esi
	int* v16;      // edi
	uint8_t** v17; // esi
	int v18;       // ebx
	uint8_t* v19;  // eax

	v7 = sub_461600(a2);
	v8 = v7;
	if (v7) {
		if (*(uint32_t*)(v7 + 112) & 0x2000000) {
			v9 = nox_xxx_equipClothFindDefByTT_413270(*(uint32_t*)(v7 + 108));
		} else {
			v9 = nox_xxx_getProjectileClassById_413250(*(uint32_t*)(v7 + 108));
		}
		v12 = v9;
		if (v9) {
			v13 = v8 + 432;
			v14 = 1;
			v15 = v9 + 4;
			do {
				LOBYTE(v9) = v15[1];
				LOBYTE(v11) = *v15;
				LOBYTE(v10) = *(v15 - 1);
				nox_draw_setMaterial_4340A0(v14++, v10, v11, (int)v9);
				v15 += 3;
			} while (v14 < 7);
			v16 = v12 + 9;
			v17 = (uint8_t**)v13;
			v18 = 4;
			do {
				v19 = *v17;
				if (*v17) {
					LOBYTE(v11) = v19[26];
					LOBYTE(v10) = v19[25];
					LOBYTE(v19) = v19[24];
					nox_draw_setMaterial_4340A0(*v16, (int)v19, v10, v11);
				}
				++v17;
				++v16;
				--v18;
			} while (v18);
		}
		if (a7) {
			nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1319052), a3, a4);
		} else {
			nox_client_drawImageAt_47D2C0(*(uint32_t*)(a5 + 4 * a6), a3, a4);
		}
	}
	return v7;
}
// 4BFA4C: variable 'v10' is possibly undefined
// 4BFA4C: variable 'v11' is possibly undefined

//----- (004BFAD0) --------------------------------------------------------
int sub_4BFAD0() {
	int v0;         // esi
	int v1;         // ebx
	int i;          // edi
	char* v3;       // eax
	const char* v4; // ecx
	int v5;         // ebp
	int v6;         // ebp

	v0 = 0;
	v1 = 0;
	for (i = 0; i < 8; i += 4) {
		v3 = nox_xxx_gLoadImg_42F970(*(const char**)getMemAt(0x587000, 180960 + i));
		v4 = *(const char**)getMemAt(0x587000, 180968 + i);
		*getMemU32Ptr(0x973A20, 16 + i) = v3;
		*getMemU32Ptr(0x973A20, 24 + i) = nox_xxx_gLoadImg_42F970(v4);
		v5 = 26;
		do {
			*getMemU32Ptr(0x973A20, 32 + v1) = nox_xxx_gLoadImg_42F970(*(const char**)getMemAt(0x587000, 180976 + v1));
			v1 += 4;
			--v5;
		} while (v5);
		v6 = 27;
		do {
			*getMemU32Ptr(0x973A20, 256 + v0) = nox_xxx_gLoadImg_42F970(*(const char**)getMemAt(0x587000, 181184 + v0));
			v0 += 4;
			--v6;
		} while (v6);
	}
	*getMemU32Ptr(0x5D4594, 1319052) = nox_xxx_gLoadImg_42F970("MaleMedievalCloakTop");
	return 1;
}

//----- (004BFB70) --------------------------------------------------------
void sub_4BFB70(int a1) {
	if (dword_5d4594_1319056) {
		dword_5d4594_1319056 = a1;
	} else {
		if (a1 == 1) {
			nox_xxx_clientPlaySoundSpecial_452D80(1022, 100);
		}
		dword_5d4594_1319056 = a1;
	}
}

//----- (004BFBB0) --------------------------------------------------------
void sub_4BFBB0(uint32_t* a1) {
	if (dword_5d4594_1319056) {
		if (dword_5d4594_1319056 == 1) {
			if (!a1) {
				sub_4BFC70();
				sub_4BFB70(0);
			}
		}
	} else if (a1 == (uint32_t*)1) {
		sub_4BFBF0();
		sub_4BFB70(1);
	}
}

//----- (004BFBF0) --------------------------------------------------------
int sub_4BFBF0() {
	int result; // eax
	int v1;     // [esp+0h] [ebp-8h]
	int v2;     // [esp+4h] [ebp-4h]

	result = dword_5d4594_1319060;
	if (dword_5d4594_1319060) {
		nox_window_set_hidden(*(int*)&dword_5d4594_1319060, 0);
		nox_xxx_wnd_46ABB0(*(int*)&dword_5d4594_1319060, 1);
		nox_window_get_size(*(int*)&dword_5d4594_1319060, &v2, &v1);
		nox_window_setPos_46A9B0(*(uint32_t**)&dword_5d4594_1319060, nox_win_width / 2 - v2 / 2,
								 nox_win_height / 2 - v1 / 2);
		result = nox_xxx_windowFocus_46B500(0);
	}
	return result;
}

//----- (004BFC70) --------------------------------------------------------
int sub_4BFC70() {
	int result; // eax

	result = dword_5d4594_1319060;
	if (dword_5d4594_1319060) {
		nox_window_set_hidden(*(int*)&dword_5d4594_1319060, 1);
		result = nox_xxx_windowFocus_46B500(0);
	}
	return result;
}

//----- (004BFC90) --------------------------------------------------------
int sub_4BFC90() {
	int result; // eax

	result = nox_new_window_from_file("SKey.wnd", sub_4BFCD0);
	dword_5d4594_1319060 = result;
	if (result) {
		sub_4BFB70(0);
		sub_4BFC70();
		result = 1;
	}
	return result;
}

//----- (004BFCD0) --------------------------------------------------------
int sub_4BFCD0(int a1, int a2, int* a3, int a4) {
	int v3; // esi

	if (a2 == 16391) {
		v3 = nox_xxx_wndGetID_46B0A0(a3);
		nox_xxx_clientPlaySoundSpecial_452D80(766, 100);
		if (v3 == 10803) {
			sub_4BFC70();
		}
	}
	return 0;
}

//----- (004BFD10) --------------------------------------------------------
void sub_4BFD10() {
	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1319060);
	dword_5d4594_1319060 = 0;
	sub_4BFB70(0);
}

//----- (004BFD30) --------------------------------------------------------
int sub_4BFD30() { return dword_5d4594_1319056; }

//----- (004BFD40) --------------------------------------------------------
void sub_4BFD40() {
	if (dword_5d4594_1319268 == 1) {
		nox_window_set_hidden(nox_gui_itemAmount_dialog_1319228, 1);
		nox_xxx_wnd_46ABB0(nox_gui_itemAmount_dialog_1319228, 0);
		nox_xxx_wnd_46C6E0(nox_gui_itemAmount_dialog_1319228);
		if (nox_gui_itemAmount_item_1319256) {
			nox_xxx_spriteDelete_45A4B0(nox_gui_itemAmount_item_1319256);
		}
		nox_gui_itemAmount_item_1319256 = 0;
		dword_5d4594_1319268 = 0;
	} else {
		nox_xxx_wnd_46ABB0(nox_gui_itemAmount_dialog_1319228, 1);
		sub_46C690(nox_gui_itemAmount_dialog_1319228);
		nox_xxx_wndShowModalMB_46A8C0(nox_gui_itemAmount_dialog_1319228);
		dword_5d4594_1319268 = 1;
	}
}

//----- (004BFDD0) --------------------------------------------------------
int sub_4BFDD0(uint32_t* a1, int a2, unsigned int a3) {
	switch (a2) {
	case 5:
	case 9:
	case 13:
		if (!nox_xxx_wndPointInWnd_46AAB0(a1, (unsigned short)a3, a3 >> 16)) {
			sub_4BFE40();
		}
		return 1;
	case 6:
	case 7:
	case 10:
	case 11:
	case 14:
	case 15:
		return 1;
	default:
		return 0;
	}
}

//----- (004BFE40) --------------------------------------------------------
int sub_4BFE40() {
	const wchar2_t* v0; // eax
	unsigned int v1;   // esi
	int v3;            // [esp+0h] [ebp-10h]
	int v4;            // [esp+4h] [ebp-Ch]
	int2 v5;           // [esp+8h] [ebp-8h]

	if (dword_5d4594_1319268 != 1) {
		return 0;
	}
	v0 = (const wchar2_t*)sub_46AF00(*(int*)&dword_5d4594_1319232);
	v1 = nox_wcstol(v0, 0, 10);
	if (v1 > *(int*)&dword_5d4594_1319248) {
		v1 = dword_5d4594_1319248;
	}
	nox_gui_getWindowOffs_46AA20(nox_gui_itemAmount_dialog_1319228, &v3, &v4);
	v5.field_0 = v3 + dword_587000_183456;
	v5.field_4 = v4 + dword_587000_183460;
	if (*getMemU32Ptr(0x5D4594, 1319100)) {
		(*(void (**)(uint32_t, uint32_t, uint32_t, uint32_t, uint32_t))getMemAt(0x5D4594, 1319100))(
			&v5, *getMemU32Ptr(0x5D4594, 1319244), *getMemU32Ptr(0x5D4594, 1319240), v1,
			*getMemU32Ptr(0x5D4594, 1319252));
	}
	sub_4BFD40();
	return 1;
}

//----- (004BFEF0) --------------------------------------------------------
int nox_gui_itemAmount_init_4BFEF0() {
	dword_5d4594_1319264 = 0;
	nox_gui_itemAmount_dialog_1319228 = nox_new_window_from_file("MultMove.wnd", sub_4C01C0);
	if (!nox_gui_itemAmount_dialog_1319228) {
		return 0;
	}
	nox_window_set_all_funcs(nox_gui_itemAmount_dialog_1319228, sub_4BFDD0, sub_4C0030, 0);
	dword_5d4594_1319232 = nox_xxx_wndGetChildByID_46B0C0(nox_gui_itemAmount_dialog_1319228, 3601);
	dword_5d4594_1319236 = nox_xxx_wndGetChildByID_46B0C0(nox_gui_itemAmount_dialog_1319228, 3607);
	nox_window_set_hidden(nox_gui_itemAmount_dialog_1319228, 1);
	nox_xxx_wnd_46ABB0(nox_gui_itemAmount_dialog_1319228, 0);
	*getMemU32Ptr(0x5D4594, 1319108) = 0;
	*getMemU32Ptr(0x5D4594, 1319112) = 0;
	*getMemU32Ptr(0x5D4594, 1319116) = nox_win_width;
	*getMemU32Ptr(0x5D4594, 1319120) = nox_win_height;
	*getMemU32Ptr(0x5D4594, 1319140) = nox_win_width;
	*getMemU32Ptr(0x5D4594, 1319144) = nox_win_height;
	*getMemU32Ptr(0x5D4594, 1319124) = 0;
	*getMemU32Ptr(0x5D4594, 1319128) = 0;
	*getMemU32Ptr(0x5D4594, 1319196) = nox_xxx_gLoadImg_42F970("MultiMoveBase");
	*getMemU32Ptr(0x5D4594, 1319200) = nox_xxx_gLoadImg_42F970("MultiMoveUpLit");
	*getMemU32Ptr(0x5D4594, 1319204) = nox_xxx_gLoadImg_42F970("MultiMoveDownLit");
	*getMemU32Ptr(0x5D4594, 1319208) = nox_xxx_gLoadImg_42F970("MultiMoveYesPressed");
	*getMemU32Ptr(0x5D4594, 1319212) = nox_xxx_gLoadImg_42F970("MultiMoveNoPressed");
	*getMemU32Ptr(0x5D4594, 1319216) = nox_xxx_gLoadImg_42F970("MultiMoveBaseNoTag");
	*getMemU32Ptr(0x5D4594, 1319220) = nox_xxx_gLoadImg_42F970("MultiMoveYesPressedNoTag");
	*getMemU32Ptr(0x5D4594, 1319224) = nox_xxx_gLoadImg_42F970("MultiMoveNoPressedNoTag");
	return 1;
}

//----- (004C0030) --------------------------------------------------------
int sub_4C0030(int a1) {
	int v1; // ecx
	int v2; // ecx
	int v3; // edx
	int v5; // [esp+0h] [ebp-8h]
	int v6; // [esp+4h] [ebp-4h]

	nox_client_drawRectFilledAlpha_49CF10(0, 0, nox_win_width, nox_win_height);
	nox_gui_getWindowOffs_46AA20(a1, &v5, &v6);
	v1 = *getMemU32Ptr(0x5D4594, 1319196);
	if (!dword_5d4594_1319264) {
		v1 = *getMemU32Ptr(0x5D4594, 1319216);
	}
	nox_client_drawImageAt_47D2C0(v1, v5, v6);
	*(uint32_t*)((uint32_t)nox_gui_itemAmount_item_1319256 + 12) = v5 + dword_587000_183456;
	*(uint32_t*)((uint32_t)nox_gui_itemAmount_item_1319256 + 16) = v6 + dword_587000_183460;
	(*(void (**)(unsigned char*, uint32_t))((uint32_t)nox_gui_itemAmount_item_1319256 + 300))(
		getMemAt(0x5D4594, 1319108), nox_gui_itemAmount_item_1319256);
	if (nox_xxx_wndGetChildByID_46B0C0(nox_gui_itemAmount_dialog_1319228, 3603)->draw_data.field_0 & 4) {
		nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1319204), v5, v6);
	}
	if (nox_xxx_wndGetChildByID_46B0C0(nox_gui_itemAmount_dialog_1319228, 3602)->draw_data.field_0 & 4) {
		nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1319200), v5, v6);
	}
	if (nox_xxx_wndGetChildByID_46B0C0(nox_gui_itemAmount_dialog_1319228, 3604)->draw_data.field_0 & 4) {
		v2 = *getMemU32Ptr(0x5D4594, 1319208);
		if (!dword_5d4594_1319264) {
			v2 = *getMemU32Ptr(0x5D4594, 1319220);
		}
		nox_client_drawImageAt_47D2C0(v2, v5, v6);
	}
	if (nox_xxx_wndGetChildByID_46B0C0(nox_gui_itemAmount_dialog_1319228, 3605)->draw_data.field_0 & 4) {
		v3 = *getMemU32Ptr(0x5D4594, 1319212);
		if (!dword_5d4594_1319264) {
			v3 = *getMemU32Ptr(0x5D4594, 1319224);
		}
		nox_client_drawImageAt_47D2C0(v3, v5, v6);
	}
	return 1;
}

//----- (004C01C0) --------------------------------------------------------
int sub_4C01C0(int a1, int a2, int* a3, int a4) {
	int v3;            // esi
	int result;        // eax
	const wchar2_t* v5; // eax
	unsigned int v6;   // esi
	const wchar2_t* v7; // eax
	unsigned int v8;   // esi
	const wchar2_t* v9; // eax
	int v10;           // eax
	int v11;           // esi
	int v12;           // [esp+4h] [ebp-10h]
	int v13;           // [esp+8h] [ebp-Ch]
	int2 v14;          // [esp+Ch] [ebp-8h]

	if (a2 != 16391 || dword_5d4594_1319268 != 1) {
		return 0;
	}
	v3 = nox_xxx_wndGetID_46B0A0(a3);
	nox_xxx_clientPlaySoundSpecial_452D80(766, 100);
	switch (v3) {
	case 3602:
		v7 = (const wchar2_t*)sub_46AF00(*(int*)&dword_5d4594_1319232);
		v8 = nox_wcstol(v7, 0, 10) + 1;
		if (v8 > *(int*)&dword_5d4594_1319248) {
			return 0;
		}
		nox_swprintf((wchar2_t*)getMemAt(0x5D4594, 1319164), L"%d", v8);
		sub_46AEE0(*(int*)&dword_5d4594_1319232, (int)getMemAt(0x5D4594, 1319164));
		if (!dword_5d4594_1319264) {
			return 0;
		}
		nox_swprintf((wchar2_t*)getMemAt(0x5D4594, 1319068), L"%d", dword_5d4594_1319260 * v8);
		sub_46AEE0(*(int*)&dword_5d4594_1319236, (int)getMemAt(0x5D4594, 1319068));
		result = 0;
		break;
	case 3603:
		v9 = (const wchar2_t*)sub_46AF00(*(int*)&dword_5d4594_1319232);
		v10 = nox_wcstol(v9, 0, 10);
		v11 = v10;
		if (v10 > 1) {
			nox_swprintf((wchar2_t*)getMemAt(0x5D4594, 1319164), L"%d", v10 - 1);
			sub_46AEE0(*(int*)&dword_5d4594_1319232, (int)getMemAt(0x5D4594, 1319164));
			if (dword_5d4594_1319264) {
				nox_swprintf((wchar2_t*)getMemAt(0x5D4594, 1319068), L"%d", dword_5d4594_1319260 * (v11 - 1));
				sub_46AEE0(*(int*)&dword_5d4594_1319236, (int)getMemAt(0x5D4594, 1319068));
			}
		}
		return 0;
	case 3604:
	case 3606:
		v5 = (const wchar2_t*)sub_46AF00(*(int*)&dword_5d4594_1319232);
		v6 = nox_wcstol(v5, 0, 10);
		if (v6 > *(int*)&dword_5d4594_1319248) {
			v6 = dword_5d4594_1319248;
		}
		nox_gui_getWindowOffs_46AA20(nox_gui_itemAmount_dialog_1319228, &v12, &v13);
		v14.field_0 = v12 + dword_587000_183456;
		v14.field_4 = v13 + dword_587000_183460;
		if (*getMemU32Ptr(0x5D4594, 1319160)) {
			(*(void (**)(uint32_t, uint32_t, uint32_t, uint32_t, uint32_t))getMemAt(0x5D4594, 1319160))(
				&v14, *getMemU32Ptr(0x5D4594, 1319244), *getMemU32Ptr(0x5D4594, 1319240), v6,
				*getMemU32Ptr(0x5D4594, 1319252));
		}
		sub_4BFD40();
		result = 0;
		break;
	case 3605:
		sub_4BFE40();
		result = 0;
		break;
	default:
		return 0;
	}
	return result;
}

//----- (004C03E0) --------------------------------------------------------
void nox_gui_itemAmount_free_4C03E0() {
	if (nox_gui_itemAmount_item_1319256) {
		nox_xxx_spriteDelete_45A4B0(nox_gui_itemAmount_item_1319256);
	}
	nox_gui_itemAmount_item_1319256 = 0;
	nox_xxx_windowDestroyMB_46C4E0(nox_gui_itemAmount_dialog_1319228);
	nox_gui_itemAmount_dialog_1319228 = 0;
	dword_5d4594_1319232 = 0;
	dword_5d4594_1319236 = 0;
	dword_5d4594_1319264 = 0;
}

//----- (004C0430) --------------------------------------------------------
int nox_gui_itemAmountDialog_4C0430(wchar2_t* title, int x, int y, int a4, int a5, const void* a6, int a7, int a8,
									void* accept, void* cancel) {
	uint32_t* v10; // eax
	int result;    // eax

	v10 = nox_xxx_wndGetChildByID_46B0C0(nox_gui_itemAmount_dialog_1319228, 3606);
	sub_46AEE0(v10, title);
	nox_gui_itemAmount_item_1319256 = nox_new_drawable_for_thing(a5);
	*(uint32_t*)((uint32_t)nox_gui_itemAmount_item_1319256 + 120) |= 0x40000000u;
	if (a6) {
		memcpy((void*)((uint32_t)nox_gui_itemAmount_item_1319256 + 432), a6, 0x14u);
	}
	*getMemU32Ptr(0x5D4594, 1319160) = accept;
	*getMemU32Ptr(0x5D4594, 1319100) = cancel;
	*getMemU32Ptr(0x5D4594, 1319240) = a5;
	*getMemU32Ptr(0x5D4594, 1319244) = a4;
	dword_5d4594_1319248 = a7;
	*getMemU32Ptr(0x5D4594, 1319252) = a8;
	sub_4BFD40();
	sub_4C0560(x, y);
	nox_swprintf((wchar2_t*)getMemAt(0x5D4594, 1319164), L"%d", 1);
	sub_46AEE0(*(int*)&dword_5d4594_1319232, (int)getMemAt(0x5D4594, 1319164));
	if (dword_5d4594_1319264) {
		nox_swprintf((wchar2_t*)getMemAt(0x5D4594, 1319068), L"%d", dword_5d4594_1319260);
		result = sub_46AEE0(*(int*)&dword_5d4594_1319236, (int)getMemAt(0x5D4594, 1319068));
	} else {
		nox_swprintf((wchar2_t*)getMemAt(0x5D4594, 1319068), (const wchar2_t*)getMemAt(0x5D4594, 1319272));
		result = sub_46AEE0(*(int*)&dword_5d4594_1319236, (int)getMemAt(0x5D4594, 1319068));
	}
	return result;
}

//----- (004C0560) --------------------------------------------------------
int sub_4C0560(int a1, int a2) {
	int v2; // ecx
	int v3; // eax
	int v5; // [esp+8h] [ebp-8h]
	int v6; // [esp+Ch] [ebp-4h]

	nox_window_get_size(nox_gui_itemAmount_dialog_1319228, &v5, &v6);
	v2 = a1 - dword_587000_183456;
	v3 = a2 - dword_587000_183460;
	if (a1 - *(int*)&dword_587000_183456 < 0) {
		v2 = 0;
	}
	if (v3 < 0) {
		v3 = 0;
	}
	if (v2 + v5 >= nox_win_width) {
		v2 = nox_win_width - v5;
	}
	if (v3 + v6 >= nox_win_height) {
		v3 = nox_win_height - v6;
	}
	return nox_window_setPos_46A9B0(nox_gui_itemAmount_dialog_1319228, v2, v3);
}

//----- (004C05F0) --------------------------------------------------------
int sub_4C05F0(int a1, int a2) {
	int result; // eax

	result = a1;
	dword_5d4594_1319264 = a1;
	dword_5d4594_1319260 = a2;
	return result;
}

//----- (004C0610) --------------------------------------------------------
int nox_xxx_func_4C0610() {
	char v1[2]; // [esp+0h] [ebp-2h]

	v1[0] = -55;
	v1[1] = 14;
	return nox_netlist_addToMsgListCli_40EBC0(31, 0, v1, 2);
}

//----- (004C0630) --------------------------------------------------------
int sub_4C0630(int a1, unsigned int a2, unsigned int a3) {
	int v3;        // eax
	uint32_t* v4;  // ecx
	uint32_t* v5;  // eax
	int v6;        // esi
	int v7;        // edi
	uint32_t* v8;  // eax
	int v9;        // esi
	int v10;       // edi
	uint32_t* v12; // eax
	uint32_t* v13; // eax
	uint32_t* v14; // esi
	int v15;       // eax
	int2 a1a;      // [esp+4h] [ebp-8h]

	a1a.field_4 = a3 >> 16;
	a1a.field_0 = (unsigned short)a3;
	if (a2 == 5) {
		v12 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3704);
		if (nox_xxx_wndPointInWnd_46AAB0(v12, a1a.field_0, a1a.field_4)) {
			LOBYTE(v13) = sub_4C0910(&a1a);
			v14 = v13;
			if (v13[1]) {
				nox_xxx_wndSetCaptureMain_46ADC0(*(int*)&dword_5d4594_1320940);
				nox_xxx_setKeybTimeout_4160D0(2);
				*(int2*)getMemAt(0x5D4594, 1319276) = a1a;
				dword_5d4594_1320968 = *v14;
				*(uint32_t*)(dword_5d4594_1320968 + 128) = v14[v14[1] + 1];
				v14[v14[1] + 1] = 0;
				nox_xxx_cursorSetDraggedItem_477690(*(int*)&dword_5d4594_1320968);
				v15 = v14[1] - 1;
				v14[1] = v15;
				if (!v15) {
					*v14 = 0;
				}
				dword_5d4594_1320972 = v14;
				*getMemU32Ptr(0x5D4594, 1320304) = 0;
				nox_xxx_clientPlaySoundSpecial_452D80(791, 100);
			}
		}
		return 1;
	}
	if (a2 <= 5 || a2 > 7) {
		return 0;
	}
	v3 = nox_xxx_wndGetCaptureMain_46AE00();
	v4 = *(uint32_t**)&dword_5d4594_1320940;
	if (v3 == dword_5d4594_1320940) {
		nox_xxx_wndClearCaptureMain_46ADE0(*(int*)&dword_5d4594_1320940);
		v4 = *(uint32_t**)&dword_5d4594_1320940;
	}
	if (!dword_5d4594_1320968) {
		return 1;
	}
	if (*getMemU32Ptr(0x5D4594, 1320304)) {
		if (*getMemU32Ptr(0x5D4594, 1320304) == 1) {
			v8 = nox_xxx_wndGetChildByID_46B0C0(v4, 3705);
			if (!nox_xxx_wndPointInWnd_46AAB0(v8, a1a.field_0, a1a.field_4) ||
				(v9 = *getMemU32Ptr(0x5D4594, 1319276) - a1a.field_0,
				 v10 = *getMemU32Ptr(0x5D4594, 1319280) - a1a.field_4,
				 !nox_xxx_checkKeybTimeout_4160F0(3u, gameFPS() / 3u)) &&
					v9 * v9 + v10 * v10 < 100) {
				nox_xxx_clientTrade_0_4C08E0(*(int*)&dword_5d4594_1320968);
			}
		}
	} else {
		v5 = nox_xxx_wndGetChildByID_46B0C0(v4, 3704);
		if (nox_xxx_wndPointInWnd_46AAB0(v5, a1a.field_0, a1a.field_4)) {
			v6 = *getMemU32Ptr(0x5D4594, 1319276) - a1a.field_0;
			v7 = *getMemU32Ptr(0x5D4594, 1319280) - a1a.field_4;
			if (!nox_xxx_checkKeybTimeout_4160F0(2u, gameFPS() / 3u) && v6 * v6 + v7 * v7 < 100) {
				nox_xxx_clientTrade_0_4C08E0(*(int*)&dword_5d4594_1320968);
			}
		} else {
			nox_xxx_clientTrade_0_4C08E0(*(int*)&dword_5d4594_1320968);
		}
	}
	*(uint32_t*)(dword_5d4594_1320972 + 4 * (*(uint32_t*)(dword_5d4594_1320972 + 4))++ + 8) =
		*(uint32_t*)(dword_5d4594_1320968 + 128);
	if (*(uint32_t*)(dword_5d4594_1320972 + 4) == 1) {
		**(uint32_t**)&dword_5d4594_1320972 = dword_5d4594_1320968;
	}
	nox_xxx_cursorResetDraggedItem_4776A0();
	dword_5d4594_1320968 = 0;
	dword_5d4594_1320972 = 0;
	return 1;
}
// 4C0836: variable 'v13' is possibly undefined

//----- (004C08E0) --------------------------------------------------------
int nox_xxx_clientTrade_0_4C08E0(int a1) {
	char v2[4]; // [esp+0h] [ebp-4h]

	v2[0] = -55;
	v2[1] = 16;
	*(uint16_t*)&v2[2] = nox_xxx_netGetUnitCodeCli_578B00(a1);
	return nox_netlist_addToMsgListCli_40EBC0(31, 0, v2, 4);
}

//----- (004C0910) --------------------------------------------------------
char sub_4C0910(int2* a1) {
	uint32_t* v1; // eax
	int v2;       // edi
	int v3;       // ebx
	int i;        // esi
	int v5;       // eax
	int v7;       // [esp+10h] [ebp-1Ch]
	int v8;       // [esp+14h] [ebp-18h]
	int v9;       // [esp+18h] [ebp-14h]
	int4 v10;     // [esp+1Ch] [ebp-10h]

	v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3704);
	nox_client_wndGetPosition_46AA60(v1, &v8, &v9);
	v2 = 0;
	v7 = 0;
	while (2) {
		v3 = 0;
		for (i = 0; i < 100; i += 50) {
			v10.field_0 = i + v8;
			v10.field_8 = i + v8 + 50;
			v10.field_4 = v2 + v9;
			v10.field_C = v2 + v9 + 50;
			v5 = nox_xxx_pointInRect_4281F0(a1, &v10);
			if (v5) {
				return (char)getMemAt(0x5D4594, 1319284 + 140 * (v7 + 2 * v3));
			}
			++v3;
		}
		v2 += 50;
		++v7;
		if (v2 < 100) {
			continue;
		}
		break;
	}
	return v5;
}
// 4C097E: variable 'v5' is possibly undefined

//----- (004C0C90) --------------------------------------------------------
int sub_4C0C90(int a1, int a2, int* a3, int a4) {
	int v3; // esi

	if (a2 == 16391) {
		v3 = nox_xxx_wndGetID_46B0A0(a3);
		nox_xxx_clientPlaySoundSpecial_452D80(766, 100);
		if (v3 == 3708) {
			nox_xxx_clientTrade_4C0CE0();
		} else if (v3 == 3710) {
			nox_xxx_func_4C0610();
			return 0;
		}
	}
	return 0;
}

//----- (004C0CE0) --------------------------------------------------------
int nox_xxx_clientTrade_4C0CE0() {
	char v1[2]; // [esp+0h] [ebp-2h]

	v1[0] = -55;
	v1[1] = 17;
	return nox_netlist_addToMsgListCli_40EBC0(31, 0, v1, 2);
}

//----- (004C0D00) --------------------------------------------------------
int sub_4C0D00() {
	uint32_t* v0;       // eax
	uint32_t* v1;       // eax
	uint32_t* v2;       // eax
	uint32_t* v3;       // eax
	uint32_t* v4;       // eax
	uint32_t* v5;       // eax
	unsigned char* v6;  // esi
	int v7;             // ebp
	int v8;             // ebx
	int v9;             // edi
	uint32_t* v10;      // eax
	unsigned char* v11; // esi
	int v12;            // ebp
	int v13;            // ebx
	int v14;            // edi
	unsigned char* v16; // [esp+10h] [ebp-58h]
	unsigned char* v17; // [esp+10h] [ebp-58h]
	int v18;            // [esp+14h] [ebp-54h]
	int v19;            // [esp+18h] [ebp-50h]
	int v20;            // [esp+1Ch] [ebp-4Ch]
	int v21;            // [esp+20h] [ebp-48h]
	int v22;            // [esp+24h] [ebp-44h]
	wchar2_t v23[32];    // [esp+28h] [ebp-40h]

	nox_gui_getWindowOffs_46AA20(*(int*)&dword_5d4594_1320940, &v21, &v22);
	nox_window_get_size(*(int*)&dword_5d4594_1320940, &v20, &v19);
	nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1320164), v21, v22);
	v0 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3711);
	nox_window_call_field_94((int)v0, 16385, (int)getMemAt(0x5D4594, 1320240), 0);
	v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3712);
	nox_window_call_field_94((int)v1, 16385, (int)getMemAt(0x5D4594, 1320868), 0);
	v2 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3713);
	nox_window_call_field_94((int)v2, 16385, (int)getMemAt(0x5D4594, 1320100), 0);
	nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1320184), v21 + *getMemU32Ptr(0x587000, 183696) - 64,
								  v22 + *getMemU32Ptr(0x587000, 183700) - 64);
	nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1320184), v21 + *getMemU32Ptr(0x587000, 183704) - 64,
								  v22 + *getMemU32Ptr(0x587000, 183708) - 64);
	v3 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3708);
	if (dword_5d4594_1320944) {
		nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1320172), v21, v22);
	} else if (v3[9] & 4) {
		nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1320168), v21, v22);
	}
	if (dword_5d4594_1320948) {
		nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1320176), v21, v22);
	}
	v4 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3710);
	if (*getMemU32Ptr(0x5D4594, 1320960) || v4[9] & 4) {
		nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1320180), v21, v22);
	}
	v5 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3704);
	nox_client_wndGetPosition_46AA60(v5, &v21, &v22);
	v6 = getMemAt(0x5D4594, 1319284);
	v18 = nox_xxx_guiFontHeightMB_43F320(0);
	v7 = 0;
	v16 = getMemAt(0x5D4594, 1319284);
	do {
		v8 = 0;
		do {
			if (*((uint32_t*)v6 + 1)) {
				*(uint32_t*)(*(uint32_t*)v6 + 12) = v8 + v21 + 25;
				*(uint32_t*)(*(uint32_t*)v6 + 16) = v22 + v7 + 25;
				(*(void (**)(unsigned char*, uint32_t))(*(uint32_t*)v6 + 300))(getMemAt(0x5D4594, 1320188),
																			   *(uint32_t*)v6);
				nox_swprintf(v23, L"%d", *((uint32_t*)v6 + 1));
				nox_xxx_drawSetTextColor_434390(nox_color_white_2523948);
				nox_xxx_drawString_43F6E0(0, (short*)v23, v8 + v21 + 5, v22 + v7 + 5);
				v9 = v22 + v7 + 50;
				nox_swprintf(v23, L"%d", *((uint32_t*)v6 + 34));
				nox_xxx_drawSetTextColor_434390(nox_color_yellow_2589772);
				nox_xxx_drawString_43F6E0(0, (short*)v23, v8 + v21 + 5, v9 - v18 - 5);
			}
			v8 += 50;
			v6 += 280;
		} while (v8 < 100);
		v7 += 50;
		v6 = v16 + 140;
		v16 += 140;
	} while ((int)v16 < (int)getMemAt(0x5D4594, 1319564));
	v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3705);
	nox_client_wndGetPosition_46AA60(v10, &v21, &v22);
	v11 = getMemAt(0x5D4594, 1320308);
	v12 = 0;
	v17 = getMemAt(0x5D4594, 1320308);
	do {
		v13 = 0;
		do {
			if (*((uint32_t*)v11 + 1)) {
				*(uint32_t*)(*(uint32_t*)v11 + 12) = v13 + v21 + 25;
				*(uint32_t*)(*(uint32_t*)v11 + 16) = v22 + v12 + 25;
				(*(void (**)(unsigned char*, uint32_t))(*(uint32_t*)v11 + 300))(getMemAt(0x5D4594, 1320188),
																				*(uint32_t*)v11);
				nox_swprintf(v23, L"%d", *((uint32_t*)v11 + 1));
				nox_xxx_drawSetTextColor_434390(nox_color_white_2523948);
				nox_xxx_drawString_43F6E0(0, (short*)v23, v13 + v21 + 5, v22 + v12 + 5);
				v14 = v22 + v12 + 50;
				nox_swprintf(v23, L"%d", *((uint32_t*)v11 + 34));
				nox_xxx_drawSetTextColor_434390(nox_color_yellow_2589772);
				nox_xxx_drawString_43F6E0(0, (short*)v23, v13 + v21 + 5, v14 - v18 - 5);
			}
			v13 += 50;
			v11 += 280;
		} while (v13 < 100);
		v12 += 50;
		v11 = v17 + 140;
		v17 += 140;
	} while ((int)v17 < (int)getMemAt(0x5D4594, 1320588));
	return 1;
}

//----- (004C1120) --------------------------------------------------------
int sub_4C1120(int a1, int a2, wchar2_t** a3) {
	uint32_t* v3; // eax
	wchar2_t** v4; // eax
	uint32_t* v5; // eax
	wchar2_t* v6;  // eax
	int2 a1a;     // [esp+0h] [ebp-8h]

	a1a.field_0 = (unsigned short)a3;
	a1a.field_4 = (unsigned int)a3 >> 16;
	v3 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3704);
	if (nox_xxx_wndPointInWnd_46AAB0(v3, (unsigned short)a3, (unsigned int)a3 >> 16)) {
		LOBYTE(v4) = sub_4C0910(&a1a);
	} else {
		v5 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3705);
		if (nox_xxx_wndPointInWnd_46AAB0(v5, a1a.field_0, a1a.field_4)) {
			LOBYTE(v4) = sub_4C11E0(&a1a);
		} else {
			v4 = a3;
		}
	}
	if (*v4) {
		*((uint32_t*)*v4 + 32) = v4[2];
		v6 = nox_xxx_clientAskInfoMb_4BF050(*v4);
		nox_xxx_cursorSetTooltip_4776B0(v6);
	}
	return 1;
}
// 4C11AC: variable 'v4' is possibly undefined

//----- (004C11E0) --------------------------------------------------------
char sub_4C11E0(uint32_t* a1) {
	uint32_t* v1; // eax
	int v2;       // edi
	int v3;       // ebx
	int i;        // esi
	int v5;       // eax
	int v7;       // [esp+10h] [ebp-1Ch]
	int v8;       // [esp+14h] [ebp-18h]
	int v9;       // [esp+18h] [ebp-14h]
	int4 v10;     // [esp+1Ch] [ebp-10h]

	v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3705);
	nox_client_wndGetPosition_46AA60(v1, &v8, &v9);
	v2 = 0;
	v7 = 0;
	while (2) {
		v3 = 0;
		for (i = 0; i < 100; i += 50) {
			v10.field_0 = i + v8;
			v10.field_8 = i + v8 + 50;
			v10.field_4 = v2 + v9;
			v10.field_C = v2 + v9 + 50;
			v5 = nox_xxx_pointInRect_4281F0((int2*)a1, &v10);
			if (v5) {
				return (char)getMemAt(0x5D4594, 1320308 + 140 * (v7 + 2 * v3));
			}
			++v3;
		}
		v2 += 50;
		++v7;
		if (v2 < 100) {
			continue;
		}
		break;
	}
	return v5;
}
// 4C124E: variable 'v5' is possibly undefined

//----- (004C12A0) --------------------------------------------------------
int nox_xxx_closeP2PTradeWnd_4C12A0() {
	int result; // eax

	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1320940);
	result = 0;
	dword_5d4594_1320940 = 0;
	dword_5d4594_1320964 = 0;
	return result;
}

//----- (004C12C0) --------------------------------------------------------
int sub_4C12C0() { return dword_5d4594_1320964; }

//----- (004C12D0) --------------------------------------------------------
int nox_xxx_showP2PTradeWnd_4C12D0() {
	int result; // eax

	result = dword_5d4594_1320964;
	if (dword_5d4594_1320964) {
		result = nox_xxx_wndShowModalMB_46A8C0(*(int*)&dword_5d4594_1320940);
	}
	return result;
}

//----- (004C1320) --------------------------------------------------------
int nox_xxx_netP2PStartTrade_4C1320(int a1) {
	int v1;       // esi
	uint32_t* v3; // eax
	uint32_t* v4; // eax
	int v5;       // [esp+4h] [ebp-8h]
	int v6;       // [esp+8h] [ebp-4h]

	v1 = dword_8531A0_2576;
	if (!dword_8531A0_2576) {
		return 0;
	}
	if (dword_5d4594_1320964 == 1) {
		return 0;
	}
	dword_5d4594_1320964 = 1;
	sub_4C1410();
	nox_wcscpy((wchar2_t*)getMemAt(0x5D4594, 1319844), (const wchar2_t*)(a1 + 2));
	nox_xxx_wnd_46ABB0(*(int*)&dword_5d4594_1320940, 1);
	nox_xxx_showP2PTradeWnd_4C12D0();
	nox_window_get_size(*(int*)&dword_5d4594_1320940, &v6, &v5);
	nox_window_setPos_46A9B0(*(uint32_t**)&dword_5d4594_1320940, 198, 193);
	v3 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3702);
	nox_window_call_field_94((int)v3, 16385, v1 + 4704, 0);
	v4 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3703);
	nox_window_call_field_94((int)v4, 16385, (int)getMemAt(0x5D4594, 1319844), 0);
	sub_467BB0();
	return 1;
}

//----- (004C1410) --------------------------------------------------------
int sub_4C1410() {
	uint32_t* v0;      // eax
	uint32_t* v1;      // eax
	unsigned char* v2; // ebp
	unsigned char* v3; // esi
	int v4;            // ebx
	unsigned char* v5; // ebp
	unsigned char* v6; // esi
	int v7;            // ebx
	uint32_t* v8;      // eax
	uint32_t* v9;      // eax
	uint32_t* v10;     // eax
	int result;        // eax
	wchar2_t v12[64];   // [esp+10h] [ebp-80h]

	*getMemU16Ptr(0x5D4594, 1320240) = 0;
	*getMemU16Ptr(0x5D4594, 1320868) = 0;
	*getMemU16Ptr(0x5D4594, 1320100) = 0;
	v0 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3702);
	nox_window_call_field_94((int)v0, 16385, (int)getMemAt(0x5D4594, 1320976), 0);
	v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3703);
	nox_window_call_field_94((int)v1, 16385, (int)getMemAt(0x5D4594, 1320980), 0);
	v2 = getMemAt(0x5D4594, 1319284);
	do {
		v3 = v2;
		v4 = 2;
		do {
			if (*(uint32_t*)v3) {
				nox_xxx_spriteDelete_45A4B0(*(uint64_t**)v3);
			}
			*(uint32_t*)v3 = 0;
			*((uint32_t*)v3 + 1) = 0;
			v3 += 280;
			--v4;
		} while (v4);
		v2 += 140;
	} while ((int)v2 < (int)getMemAt(0x5D4594, 1319564));
	v5 = getMemAt(0x5D4594, 1320308);
	do {
		v6 = v5;
		v7 = 2;
		do {
			if (*(uint32_t*)v6) {
				nox_xxx_spriteDelete_45A4B0(*(uint64_t**)v6);
			}
			*(uint32_t*)v6 = 0;
			*((uint32_t*)v6 + 1) = 0;
			v6 += 280;
			--v7;
		} while (v7);
		v5 += 140;
	} while ((int)v5 < (int)getMemAt(0x5D4594, 1320588));
	nox_swprintf(v12, (const wchar2_t*)getMemAt(0x5D4594, 1319972), 0);
	v8 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3711);
	nox_window_call_field_94((int)v8, 16385, (int)v12, 0);
	v9 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3712);
	nox_window_call_field_94((int)v9, 16385, (int)v12, 0);
	v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3713);
	result = nox_window_call_field_94((int)v10, 16385, (int)v12, 0);
	dword_5d4594_1320944 = 0;
	dword_5d4594_1320948 = 0;
	dword_5d4594_1320968 = 0;
	dword_5d4594_1320972 = 0;
	dword_5d4594_1320932 = 0;
	dword_5d4594_1320936 = 0;
	return result;
}

//----- (004C1590) --------------------------------------------------------
int sub_4C1590() {
	int result; // eax

	result = dword_5d4594_1320964;
	if (dword_5d4594_1320964) {
		sub_4C1410();
		nox_window_set_hidden(*(int*)&dword_5d4594_1320940, 1);
		nox_xxx_wnd_46ABB0(*(int*)&dword_5d4594_1320940, 0);
		dword_5d4594_1320964 = 0;
		result = sub_467C10();
	}
	return result;
}

//----- (004C1710) --------------------------------------------------------
int sub_4C1710(int a1, int a2) {
	int result;   // eax
	uint32_t* i;  // ecx
	uint32_t* v4; // ecx
	int v5;       // edx

	result = 0;
	for (i = (uint32_t*)(a1 + 8); *i != a2; ++i) {
		if (++result >= 32) {
			return result;
		}
	}
	if (result < 31) {
		v4 = (uint32_t*)(a1 + 4 * result + 8);
		v5 = 31 - result;
		do {
			result = v4[1];
			*v4 = result;
			++v4;
			--v5;
		} while (v5);
	}
	*(uint32_t*)(a1 + 132) = 0;
	return result;
}

//----- (004C1760) --------------------------------------------------------
int sub_4C1760(int a1, int a2) {
	int v2;      // eax
	uint32_t* i; // ecx

	if (!*(uint32_t*)(a1 + 4)) {
		return 0;
	}
	v2 = 0;
	for (i = (uint32_t*)(a1 + 8); *i != a2; ++i) {
		if (++v2 >= 32) {
			return 0;
		}
	}
	return 1;
}

//----- (004C1790) --------------------------------------------------------
char* nox_xxx_tradeClientAddItem_4C1790(int a1) {
	char* result; // eax
	int v2;       // ebp
	char* v3;     // esi
	uint32_t* v4; // eax
	int v5;       // ecx
	int* v6;      // ebx
	uint8_t* v7;  // edi
	int v8;       // eax
	int v9;       // eax
	int v10;      // [esp+10h] [ebp+4h]

	result = *(char**)&dword_5d4594_1320964;
	if (!dword_5d4594_1320964) {
		return result;
	}
	v2 = a1;
	dword_5d4594_1320944 = 0;
	dword_5d4594_1320948 = 0;
	if (*(uint8_t*)(a1 + 2) == 1) {
		v3 = *(char**)&dword_5d4594_1320932;
		if (dword_5d4594_1320932) {
			result = (char*)sub_4C18E0(*(unsigned short*)(a1 + 3), *(uint32_t**)&dword_5d4594_1320932);
			if (result) {
				goto LABEL_12;
			}
			result = sub_4C1910(*(unsigned short*)(a1 + 3));
		} else {
			result = sub_4C1910(*(unsigned short*)(a1 + 3));
		}
	} else {
		v3 = *(char**)&dword_5d4594_1320936;
		if (dword_5d4594_1320936) {
			result = (char*)sub_4C18E0(*(unsigned short*)(a1 + 3), *(uint32_t**)&dword_5d4594_1320936);
			if (result) {
				goto LABEL_12;
			}
			result = sub_4C19C0(*(unsigned short*)(a1 + 3));
		} else {
			result = sub_4C19C0(*(unsigned short*)(a1 + 3));
		}
	}
	v3 = result;
LABEL_12:
	if (v3) {
		if (!*(uint32_t*)v3) {
			v4 = nox_new_drawable_for_thing(*(unsigned short*)(a1 + 3));
			*(uint32_t*)v3 = v4;
			if (v4[28] & 0x13001000) {
				v5 = -11 - a1;
				v6 = v4 + 108;
				v7 = (uint8_t*)(a1 + 11);
				v10 = -11 - a1;
				do {
					if (*v7 == -1) {
						*v6 = 0;
					} else {
						v8 = nox_xxx_modifGetDescById_413330((unsigned char)*v7);
						v5 = v10;
						*v6 = v8;
					}
					++v7;
					++v6;
				} while ((int)&v7[v5] < 4);
			}
			*((uint32_t*)v3 + 1) = 0;
			*((uint32_t*)v3 + 34) = 0;
		}
		*(uint32_t*)&v3[4 * *((uint32_t*)v3 + 1) + 8] = *(unsigned short*)(v2 + 5);
		v9 = *((uint32_t*)v3 + 34);
		++*((uint32_t*)v3 + 1);
		result = (char*)(*(uint32_t*)(v2 + 7) + v9);
		*((uint32_t*)v3 + 34) = result;
		dword_5d4594_1320932 = 0;
		dword_5d4594_1320936 = 0;
	}
	return result;
}

//----- (004C18E0) --------------------------------------------------------
int sub_4C18E0(int a1, uint32_t* a2) {
	if (!a2[1]) {
		return 1;
	}
	if (*(uint32_t*)(*a2 + 108) != a1 || *(uint32_t*)(*a2 + 112) & 0x13001000) {
		return 0;
	}
	return 1;
}

//----- (004C1910) --------------------------------------------------------
char* sub_4C1910(int a1) {
	int v1;            // edi
	unsigned char* v2; // esi
	int v3;            // ecx
	unsigned char* v4; // edx
	int v5;            // eax
	int v6;            // esi
	unsigned char* v7; // edx
	int v8;            // eax
	unsigned char* v9; // ecx

	v1 = 0;
	v2 = getMemAt(0x5D4594, 1319284);
	while (2) {
		v3 = 0;
		v4 = v2;
		do {
			v5 = *(uint32_t*)v4;
			if (*(uint32_t*)v4 && *(uint32_t*)(v5 + 108) == a1 && !(*(uint32_t*)(v5 + 112) & 0x13001000)) {
				return (char*)getMemAt(0x5D4594, 1319284 + 140 * (v1 + 2 * v3));
			}
			++v3;
			v4 += 280;
		} while (v3 < 2);
		v2 += 140;
		++v1;
		if ((int)v2 < (int)getMemAt(0x5D4594, 1319564)) {
			continue;
		}
		break;
	}
	v6 = 0;
	v7 = getMemAt(0x5D4594, 1319284);
	while (2) {
		v8 = 0;
		v9 = v7;
		do {
			if (!*(uint32_t*)v9) {
				return (char*)getMemAt(0x5D4594, 1319284 + 140 * (v6 + 2 * v8));
			}
			++v8;
			v9 += 280;
		} while (v8 < 2);
		v7 += 140;
		++v6;
		if ((int)v7 < (int)getMemAt(0x5D4594, 1319564)) {
			continue;
		}
		break;
	}
	return 0;
}

//----- (004C19C0) --------------------------------------------------------
char* sub_4C19C0(int a1) {
	int v1;            // edi
	unsigned char* v2; // esi
	int v3;            // ecx
	unsigned char* v4; // edx
	int v5;            // eax
	int v6;            // esi
	unsigned char* v7; // edx
	int v8;            // eax
	unsigned char* v9; // ecx

	v1 = 0;
	v2 = getMemAt(0x5D4594, 1320308);
	while (2) {
		v3 = 0;
		v4 = v2;
		do {
			v5 = *(uint32_t*)v4;
			if (*(uint32_t*)v4 && *(uint32_t*)(v5 + 108) == a1 && !(*(uint32_t*)(v5 + 112) & 0x13001000)) {
				return (char*)getMemAt(0x5D4594, 1320308 + 140 * (v1 + 2 * v3));
			}
			++v3;
			v4 += 280;
		} while (v3 < 2);
		v2 += 140;
		++v1;
		if ((int)v2 < (int)getMemAt(0x5D4594, 1320588)) {
			continue;
		}
		break;
	}
	v6 = 0;
	v7 = getMemAt(0x5D4594, 1320308);
	while (2) {
		v8 = 0;
		v9 = v7;
		do {
			if (!*(uint32_t*)v9) {
				return (char*)getMemAt(0x5D4594, 1320308 + 140 * (v6 + 2 * v8));
			}
			++v8;
			v9 += 280;
		} while (v8 < 2);
		v7 += 140;
		++v6;
		if ((int)v7 < (int)getMemAt(0x5D4594, 1320588)) {
			continue;
		}
		break;
	}
	return 0;
}

//----- (004C1B50) --------------------------------------------------------
int sub_4C1B50(int a1) {
	int result; // eax

	result = dword_5d4594_1320964;
	if (dword_5d4594_1320964) {
		nox_swprintf((wchar2_t*)getMemAt(0x5D4594, 1320240), L"%d", *(uint32_t*)(a1 + 2));
		if (*(uint32_t*)(a1 + 6)) {
			nox_swprintf((wchar2_t*)getMemAt(0x5D4594, 1320868), L"(%d)", *(uint32_t*)(a1 + 6));
		} else {
			nox_wcscpy((wchar2_t*)getMemAt(0x5D4594, 1320868), (const wchar2_t*)getMemAt(0x5D4594, 1320984));
		}
		result = nox_swprintf((wchar2_t*)getMemAt(0x5D4594, 1320100), L"%d", *(uint32_t*)(a1 + 10));
	}
	return result;
}

//----- (004C1BC0) --------------------------------------------------------
int sub_4C1BC0(int a1) {
	unsigned int v1; // edx
	int result;      // eax

	result = dword_5d4594_1320964;
	if (dword_5d4594_1320964) {
		result = a1;
		dword_5d4594_1320944 = *(uint8_t*)(a1 + 2) & 1;
		LOBYTE(v1) = *(uint8_t*)(a1 + 2);
		dword_5d4594_1320948 = (v1 >> 1) & 1;
	}
	return result;
}
// 4C1BDC: variable 'v1' is possibly undefined

//----- (004C1BF0) --------------------------------------------------------
int nox_xxx_prepareP2PTrade_4C1BF0() {
	int result;   // eax
	int v1;       // esi
	uint32_t* v2; // eax
	uint32_t* v3; // eax

	result = dword_5d4594_1320964;
	v1 = dword_8531A0_2576;
	if (dword_5d4594_1320964) {
		if (dword_8531A0_2576) {
			sub_4C1410();
			v2 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3702);
			nox_window_call_field_94((int)v2, 16385, v1 + 4704, 0);
			v3 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1320940, 3703);
			result = nox_window_call_field_94((int)v3, 16385, (int)getMemAt(0x5D4594, 1319844), 0);
		}
	}
	return result;
}

//----- (004C1CA0) --------------------------------------------------------
int sub_4C1CA0(int a1) {
	int result; // eax

	*getMemU32Ptr(0x587000, 184448) = a1;
	if (a1 == 3) {
		nox_xxx_wndSetIcon_46AE60(*(int*)&dword_5d4594_1321040, *getMemIntPtr(0x5D4594, 1321020));
		sub_46AEC0(*(int*)&dword_5d4594_1321040, *getMemIntPtr(0x5D4594, 1321016));
		result = nox_xxx_wndSetIconLit_46AEA0(*(int*)&dword_5d4594_1321040, *getMemIntPtr(0x5D4594, 1321016));
	} else if (a1 == 4) {
		nox_xxx_wndSetIcon_46AE60(*(int*)&dword_5d4594_1321040, *getMemIntPtr(0x5D4594, 1321028));
		sub_46AEC0(*(int*)&dword_5d4594_1321040, *(int*)&dword_5d4594_1321024);
		result = nox_xxx_wndSetIconLit_46AEA0(*(int*)&dword_5d4594_1321040, *(int*)&dword_5d4594_1321024);
	} else {
		result = a1 - 5;
		if (a1 == 5) {
			nox_xxx_wndSetIcon_46AE60(*(int*)&dword_5d4594_1321040, *getMemIntPtr(0x5D4594, 1321012));
			sub_46AEC0(*(int*)&dword_5d4594_1321040, *getMemIntPtr(0x5D4594, 1321008));
			result = nox_xxx_wndSetIconLit_46AEA0(*(int*)&dword_5d4594_1321040, *getMemIntPtr(0x5D4594, 1321008));
		}
	}
	return result;
}

//----- (004C1FE0) --------------------------------------------------------
int nox_xxx_guiDrawSummonBox_4C1FE0(uint32_t* a1) {
	unsigned char* v2; // ebx
	int v3;            // ebp
	int v4;            // eax
	int v5;            // esi
	int v6;            // eax
	int v7;            // edi
	int v8;            // ebp
	int v9;            // edi
	int v10;           // edi
	int v11;           // ecx
	int v13;           // ebx
	int* v14;          // esi
	uint32_t* v15;     // eax
	int v17;           // [esp-18h] [ebp-4Ch]
	int v18;           // [esp-14h] [ebp-48h]
	char v19;          // [esp+7h] [ebp-2Dh]
	int v20;           // [esp+8h] [ebp-2Ch]
	int v21;           // [esp+Ch] [ebp-28h]
	int v22 = 0;       // [esp+10h] [ebp-24h]
	int v23 = 0;       // [esp+14h] [ebp-20h]
	int2 a1a;          // [esp+1Ch] [ebp-18h]
	int v26;           // [esp+24h] [ebp-10h]
	int v27;           // [esp+28h] [ebp-Ch]
	int v28;           // [esp+2Ch] [ebp-8h]
	int v29;           // [esp+30h] [ebp-4h]

	nox_point mpos = nox_client_getMousePos_4309F0();
	nox_point v24 = mpos;
	if (getMemByte(0x5D4594, 1321200) == 1) {
		dword_5d4594_1320992 += 20;
		if (dword_5d4594_1320992 >= *getMemIntPtr(0x5D4594, 1321004)) {
			dword_5d4594_1320992 = *getMemU32Ptr(0x5D4594, 1321004);
			*getMemU8Ptr(0x5D4594, 1321200) = 2;
		}
	} else if (getMemByte(0x5D4594, 1321200) == 3) {
		dword_5d4594_1320992 -= 20;
		if (dword_5d4594_1320992 <= *getMemIntPtr(0x5D4594, 1321000)) {
			dword_5d4594_1320992 = *getMemU32Ptr(0x5D4594, 1321000);
			*getMemU8Ptr(0x5D4594, 1321200) = 0;
			nox_window_set_hidden(*(int*)&dword_5d4594_1321032, 1);
			nox_window_set_hidden(*(int*)&dword_5d4594_1321040, 1);
		}
	}
	nox_window_setPos_46A9B0(*(uint32_t**)&dword_5d4594_1321032, *(int*)&dword_5d4594_1320988,
							 *(int*)&dword_5d4594_1320992);
	nox_window_setPos_46A9B0(*(uint32_t**)&dword_5d4594_1321040, dword_5d4594_1320988 + 27, dword_5d4594_1320992 + 12);
	nox_client_wndGetPosition_46AA60(a1, &v28, &v29);
	nox_window_get_size((int)a1, &v21, &v20);
	nox_xxx_guiFontHeightMB_43F320(0);
	if (*getMemU32Ptr(0x5D4594, 1320996)) {
		nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1320996), *(int*)&dword_5d4594_1320988,
									  *(int*)&dword_5d4594_1320992);
	}
	v2 = getMemAt(0x5D4594, 1321064);
	do {
		if (*((uint32_t*)v2 - 1)) {
			switch (v2[8]) {
			case 1u:
				v22 = 1;
				v23 = 1;
				break;
			case 2u:
				v22 = 1;
				v23 = 2;
				break;
			case 4u:
				v22 = 2;
				v23 = 2;
				break;
			}
			v3 = v28 + 38 * *(uint32_t*)v2 + 2;
			v4 = *((uint32_t*)v2 + 1);
			a1a.field_0 = v28 + 38 * *(uint32_t*)v2 + 2;
			v5 = v29 + 38 * v4 + 2;
			v21 = 38 * v22 - 4;
			v20 = 38 * v23 - 4;
			if (*((uint32_t*)v2 + 4)) {
				nox_client_drawSetColor_434460(nox_color_yellow_2589772);
				v18 = v20;
				v17 = v21;
				*((uint32_t*)v2 + 4) = 0;
				nox_client_drawRectFilledOpaque_49CE30(v3, v5, v17, v18);
			} else {
				v6 = nox_xxx_guiDrawSummon_4C2440(*((uint32_t*)v2 - 2));
				if (v6) {
					nox_client_drawImageAt_47D2C0(v6, v3, v5);
				} else {
					v7 = v3 + v21 / 2;
					v8 = v5 + v20 / 2;
					nox_client_drawSetColor_434460(*getMemIntPtr(0x85B3FC, 956));
					nox_client_drawPoint_4B0BC0(v7, v8, 9);
					nox_video_drawCircleColored_4C3270(v7, v8, 9, *getMemIntPtr(0x852978, 4));
					v3 = a1a.field_0;
				}
			}
			if (sub_495180(*((uint32_t*)v2 - 3), &v27, &v26, &v19)) {
				if ((uint16_t)v26) {
					v9 = v20 * (unsigned short)v27 / (unsigned short)v26;
				} else {
					v9 = 0;
				}
				if (v19) {
					nox_client_drawSetColor_434460(*getMemIntPtr(0x85B3FC, 984));
				} else {
					nox_client_drawSetColor_434460(nox_color_violet_2598268);
				}
				nox_client_drawRectFilledOpaque_49CE30(v21 + v3 - 2, v5, 2, v20);
				if (v19) {
					nox_client_drawSetColor_434460(dword_8531A0_2572);
				} else {
					nox_client_drawSetColor_434460(*getMemIntPtr(0x85B3FC, 940));
				}
				nox_client_drawRectFilledOpaque_49CE30(v21 + v3 - 2, v20 + v5 - v9, 2, v9);
			}
			mpos = v24;
		}
		v2 += 32;
	} while ((int)v2 < (int)getMemAt(0x5D4594, 1321192));
	v10 = nox_xxx_wndPointInWnd_46AAB0(*(uint32_t**)&dword_5d4594_1321040, mpos.x, mpos.y);
	if (nox_xxx_wndPointInWnd_46AAB0(a1, mpos.x, mpos.y) || v10 || *getMemU32Ptr(0x5D4594, 1321212) == 1) {
		v11 = mpos.y;
		a1a.field_0 = (mpos.x - v28) / 38;
		a1a.field_4 = (v11 - v29) / 38;
		v13 = nox_xxx_wndSummonGet_4C2410(&a1a);
		*getMemU32Ptr(0x5D4594, 1321212) = 0;
		v14 = (int*)sub_4C2D60();
		if (v14) {
			do {
				v15 = nox_xxx_netSpriteByCodeDynamic_45A6F0(*v14);
				if (v15) {
					if (v14 == (int*)v13 || v10) {
						v15[30] |= 0x40000000u;
						*getMemU32Ptr(0x5D4594, 1321212) = 1;
					} else {
						v15[30] &= 0xBFFFFFFF;
					}
				}
				v14 = (int*)sub_4C2D90((int)v14);
			} while (v14);
			mpos = v24;
		}
	}
	if (dword_5d4594_1321044 && !nox_xxx_wndPointInWnd_46AAB0(*(uint32_t**)&dword_5d4594_1321044, mpos.x, mpos.y)) {
		nox_xxx_guiHideSummonWindow_4C2470();
	}
	return 1;
}

//----- (004C2410) --------------------------------------------------------
int nox_xxx_wndSummonGet_4C2410(int2* a1) {
	int result; // eax
	int v2;     // edx
	int v3;     // ecx

	result = 0;
	if (a1) {
		v2 = a1->field_0;
		if (a1->field_0 >= 0 && v2 < 2) {
			v3 = a1->field_4;
			if (v3 >= 0 && v3 < 2) {
				result = *getMemU32Ptr(0x5D4594, 1321180 + 4 * (v3 + 2 * v2));
			}
		}
	}
	return result;
}

//----- (004C2440) --------------------------------------------------------
int nox_xxx_guiDrawSummon_4C2440(int a1) {
	const char* v1; // eax
	int v2;         // eax

	v1 = (const char*)nox_get_thing_name(a1);
	if (!v1) {
		return 0;
	}
	v2 = nox_xxx_guide_427010(v1);
	return sub_427430(v2);
}

//----- (004C2470) --------------------------------------------------------
int nox_xxx_guiHideSummonWindow_4C2470() {
	int result; // eax

	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1321044);
	dword_5d4594_1321044 = 0;
	dword_5d4594_1321204 = 0;
	nox_xxx_clientPlaySoundSpecial_452D80(920, 100);
	result = dword_5d4594_1321040;
	*(uint32_t*)(dword_5d4594_1321040 + 36) &= 0xFFFFFFFD;
	return result;
}

//----- (004C24A0) --------------------------------------------------------
int sub_4C24A0() { return 1; }

//----- (004C24B0) --------------------------------------------------------
int nox_xxx_wndSummonBigButtonProc_4C24B0(int a1, int a2, unsigned int a3) {
	int2 a1a; // [esp+0h] [ebp-8h]

	switch (a2) {
	case 5:
	case 6:
		return 1;
	case 7:
		dword_5d4594_1321204 = 0;
		a1a.field_4 = a3 >> 16;
		a1a.field_0 = (unsigned short)a3;
		nox_xxx_wndSummonCreateList_4C2560(&a1a);
		*(uint32_t*)(dword_5d4594_1321040 + 36) |= 2u;
		return 1;
	case 17:
		*(uint32_t*)(dword_5d4594_1321040 + 36) |= 2u;
		return 0;
	case 18:
		*(uint32_t*)(dword_5d4594_1321040 + 36) &= 0xFFFFFFFD;
		return 0;
	default:
		return 0;
	}
}

//----- (004C2A00) --------------------------------------------------------
int sub_4C2A00(int a1, int a2, int a3, int a4, short* a5) {
	unsigned char* v5; // esi

	nox_xxx_drawSetTextColor_434390(a4);
	v5 = getMemAt(0x587000, 184520);
	do {
		nox_xxx_drawString_43F6E0(0, a5, a1 + *(uint32_t*)v5, a2 + *((uint32_t*)v5 + 1));
		v5 += 8;
	} while ((int)v5 < (int)getMemAt(0x587000, 184552));
	nox_xxx_drawSetTextColor_434390(a3);
	return nox_xxx_drawString_43F6E0(0, a5, a1, a2);
}

void nox_client_orderCreature(int creature, int command) {
	unsigned char buf[4] = {0};
	buf[0] = 0x78;
	if (creature) {
		*(uint16_t*)(&buf[1]) = *(uint16_t*)creature;
	} else {
		if (command == 1) {
			return;
		}
		*(uint16_t*)(&buf[1]) = 0;
	}
	buf[3] = command;
	nox_netlist_addToMsgListCli_40EBC0(31, 0, buf, 4);
	nox_xxx_guiHideSummonWindow_4C2470();
	if (!command) {
		nox_xxx_clientPlaySoundSpecial_452D80(777, 100);
	} else {
		nox_xxx_clientPlaySoundSpecial_452D80(898, 100);
	}
}

//----- (004C2A60) --------------------------------------------------------
int nox_xxx_clientOrderCreature_4C2A60(int a1, unsigned int a2) {
	int result; // eax

	if (a2 >= 5) {
		if (a2 <= 6) {
			return 1;
		}
		if (a2 == 7) {
			LOBYTE(a2) = 120;
			result = *(uint32_t*)(a1 + 32);
			if (result == 2 || (!dword_5d4594_1321204 && result == 1)) {
				return 1;
			}
			nox_client_orderCreature(dword_5d4594_1321204, *(uint32_t*)(a1 + 32));
			return 1;
		}
	}
	return 0;
}

//----- (004C2B10) --------------------------------------------------------
int nox_xxx_wndSummonProc_4C2B10(uint32_t* a1, unsigned int a2, unsigned int a3) {
	int2 v5;  // [esp+0h] [ebp-10h]
	int2 a1a; // [esp+8h] [ebp-8h]

	if (a2 >= 5) {
		if (a2 <= 6) {
			return 1;
		}
		if (a2 == 7) {
			v5.field_4 = a3 >> 16;
			v5.field_0 = (unsigned short)a3;
			nox_client_wndGetPosition_46AA60(a1, &a2, &a3);
			a1a.field_0 = (v5.field_0 - (int)a2) / 38;
			a1a.field_4 = (v5.field_4 - (int)a3) / 38;
			dword_5d4594_1321204 = nox_xxx_wndSummonGet_4C2410(&a1a);
			if (dword_5d4594_1321204) {
				nox_xxx_wndSummonCreateList_4C2560(&v5);
			}
			return 1;
		}
	}
	return 0;
}

//----- (004C2BD0) --------------------------------------------------------
int sub_4C2BD0() { return 0; }

//----- (004C2BE0) --------------------------------------------------------
int sub_4C2BE0() { return 1; }

//----- (004C2BF0) --------------------------------------------------------
int* sub_4C2BF0() {
	int* v0;     // edx
	int* result; // eax
	int v2;      // ecx

	v0 = getMemIntPtr(0x5D4594, 1321180);
	do {
		result = v0;
		v2 = 2;
		do {
			*result = 0;
			result += 2;
			--v2;
		} while (v2);
		++v0;
	} while ((int)v0 < (int)getMemAt(0x5D4594, 1321188));
	return result;
}

//----- (004C2C20) --------------------------------------------------------
int sub_4C2C20(uint32_t* a1, int a2, unsigned int a3) {
	wchar2_t* v3; // eax
	int2 a2a;    // [esp+0h] [ebp-8h]

	a2a.field_4 = a3 >> 16;
	a2a.field_0 = (unsigned short)a3;
	v3 = (wchar2_t*)sub_4C2C60(a1, &a2a);
	nox_xxx_cursorSetTooltip_4776B0(v3);
	return 1;
}

//----- (004C2C60) --------------------------------------------------------
int sub_4C2C60(uint32_t* a1, int2* a2) {
	int v4;     // eax
	int result; // eax
	int v6;     // [esp+4h] [ebp-Ch]
	int2 a1a;   // [esp+8h] [ebp-8h]

	nox_client_wndGetPosition_46AA60(a1, &a1, &v6);
	a1a.field_0 = (a2->field_0 - (int)a1) / 38;
	a1a.field_4 = (a2->field_4 - (int)v6) / 38;
	v4 = nox_xxx_wndSummonGet_4C2410(&a1a);
	if (v4) {
		result = nox_get_thing_pretty_name(*(uint32_t*)(v4 + 4));
	} else {
		result = 0;
	}
	return result;
}

//----- (004C2D60) --------------------------------------------------------
char* sub_4C2D60() {
	int v0;            // eax
	unsigned char* v1; // ecx

	v0 = 0;
	v1 = getMemAt(0x5D4594, 1321060);
	while (!*(uint32_t*)v1) {
		v1 += 32;
		++v0;
		if ((int)v1 >= (int)getMemAt(0x5D4594, 1321188)) {
			return 0;
		}
	}
	return (char*)getMemAt(0x5D4594, 1321052 + 32 * v0);
}

//----- (004C2D90) --------------------------------------------------------
char* sub_4C2D90(int a1) {
	int v1;            // eax
	unsigned char* v2; // ecx

	v1 = *(uint32_t*)(a1 + 24) + 1;
	if (v1 >= 4) {
		return 0;
	}
	v2 = getMemAt(0x5D4594, 1321060 + 32 * v1);
	while (!*(uint32_t*)v2) {
		v2 += 32;
		++v1;
		if ((int)v2 >= (int)getMemAt(0x5D4594, 1321188)) {
			return 0;
		}
	}
	return (char*)getMemAt(0x5D4594, 1321052 + 32 * v1);
}

//----- (004C2DD0) --------------------------------------------------------
int sub_4C2DD0(int a1) {
	int v1; // eax

	v1 = dword_5d4594_1321208;
	if (!dword_5d4594_1321208) {
		v1 = nox_xxx_getNameId_4E3AA0("CarnivorousPlant");
		dword_5d4594_1321208 = v1;
	}
	return *(uint32_t*)(a1 + 4) != v1;
}

//----- (004C2E00) --------------------------------------------------------
int sub_4C2E00() {
	char* v0; // esi

	if (!dword_5d4594_1321208) {
		dword_5d4594_1321208 = nox_xxx_getNameId_4E3AA0("CarnivorousPlant");
	}
	v0 = sub_4C2D60();
	if (!v0) {
		return 0;
	}
	while (!sub_4C2DD0((int)v0)) {
		v0 = sub_4C2D90((int)v0);
		if (!v0) {
			return 0;
		}
	}
	return 1;
}

//----- (004C2E50) --------------------------------------------------------
char nox_xxx_cliSummonCreat_4C2E50(int a1, int a2, int a3) {
	char* v3; // eax

	v3 = sub_4C2D60();
	if (v3) {
		while (*(uint32_t*)v3 != a1 || *((uint32_t*)v3 + 1) != a2) {
			v3 = sub_4C2D90((int)v3);
			if (!v3) {
				goto LABEL_5;
			}
		}
		return (char)v3;
	}
LABEL_5:
	v3 = sub_4C2F20();
	if (v3) {
		*(uint32_t*)v3 = a1;
		*((uint32_t*)v3 + 1) = a2;
		v3[20] = sub_4C2EF0(a2);
		sub_4C2F70();
		LOBYTE(v3) = getMemByte(0x5D4594, 1321200);
		if (!getMemByte(0x5D4594, 1321200) || getMemByte(0x5D4594, 1321200) == 3) {
			if (!a3) {
				nox_xxx_clientPlaySoundSpecial_452D80(801, 100);
			}
			*getMemU8Ptr(0x5D4594, 1321200) = 1;
			nox_window_set_hidden(*(int*)&dword_5d4594_1321032, 0);
			LOBYTE(v3) = nox_window_set_hidden(*(int*)&dword_5d4594_1321040, 0);
		}
		++dword_5d4594_1321196;
	}
	return (char)v3;
}

//----- (004C2EF0) --------------------------------------------------------
int sub_4C2EF0(int a1) {
	int v1; // eax
	int v3; // eax

	v1 = nox_get_thing(a1)->sub_class;
	if (v1 & 1) {
		return 1;
	}
	v3 = -((v1 & 2) != 0);
	LOBYTE(v3) = v3 & 0xFE;
	return v3 + 4;
}

//----- (004C2F20) --------------------------------------------------------
char* sub_4C2F20() {
	int v0;            // esi
	unsigned char* v1; // eax
	char* result;      // eax

	v0 = 0;
	v1 = getMemAt(0x5D4594, 1321060);
	while (*(uint32_t*)v1) {
		v1 += 32;
		++v0;
		if ((int)v1 >= (int)getMemAt(0x5D4594, 1321188)) {
			return 0;
		}
	}
	memset(getMemAt(0x5D4594, 1321052 + 32 * v0), 0, 0x20u);
	*getMemU32Ptr(0x5D4594, 1321060 + 32 * v0) = 1;
	result = (char*)getMemAt(0x5D4594, 1321052 + 32 * v0);
	*getMemU32Ptr(0x5D4594, 1321076 + 32 * v0) = v0;
	return result;
}

//----- (004C2F70) --------------------------------------------------------
char* sub_4C2F70() {
	int i;        // edi
	char* result; // eax
	char* j;      // esi

	sub_4C2BF0();
	for (i = 1;; i = 4) {
		while (1) {
			result = sub_4C2D60();
			for (j = result; result; j = result) {
				if ((unsigned char)j[20] == i) {
					sub_4C2FD0((int)j);
				}
				result = sub_4C2D90((int)j);
			}
			if (i != 1) {
				break;
			}
			i = 2;
		}
		if (i != 2) {
			break;
		}
	}
	return result;
}

//----- (004C2FD0) --------------------------------------------------------
int sub_4C2FD0(int a1) {
	int v1;     // esi
	int* v2;    // edi
	int v3;     // edx
	int result; // eax

	v1 = 0;
	v2 = (int*)(a1 + 12);
	while (1) {
		v3 = *(unsigned char*)(a1 + 20);
		*v2 = *getMemU32Ptr(0x587000, 184456 + 8 * v1);
		*(uint32_t*)(a1 + 16) = *getMemU32Ptr(0x587000, 184460 + 8 * v1);
		if (sub_4C30C0(v2, v3)) {
			break;
		}
		result = *(unsigned char*)(a1 + 20);
		v1 += result;
		if (v1 >= 4) {
			return result;
		}
	}
	return sub_4C3030((int*)(a1 + 12), *(unsigned char*)(a1 + 20), a1);
}

//----- (004C3030) --------------------------------------------------------
int sub_4C3030(int* a1, int a2, int a3) {
	int v3;            // esi
	int result;        // eax
	int v5;            // ebx
	int v6;            // eax
	unsigned char* v7; // ecx

	if (a2 == 1) {
		v3 = 1;
		a2 = 1;
	} else if (a2 == 2) {
		v3 = 1;
		a2 = 2;
	} else {
		v3 = 2;
		if (a2 == 4) {
			a2 = 2;
		} else {
			v3 = a2;
		}
	}
	result = a2;
	v5 = a1[1];
	if (v5 < v5 + a2) {
		do {
			v6 = *a1;
			if (*a1 < *a1 + v3) {
				v7 = getMemAt(0x5D4594, 1321180 + 4 * (v5 + 2 * v6));
				do {
					*(uint32_t*)v7 = a3;
					++v6;
					v7 += 8;
				} while (v6 < v3 + *a1);
			}
			++v5;
			result = a1[1] + a2;
		} while (v5 < result);
	}
	return result;
}

//----- (004C30C0) --------------------------------------------------------
int sub_4C30C0(int* a1, int a2) {
	int v2;            // ecx
	int v3;            // eax
	int v4;            // esi
	int v5;            // ebx
	int v6;            // edi
	int v7;            // edx
	int v8;            // eax
	unsigned char* v9; // ecx

	switch (a2) {
	case 1:
		v2 = 1;
		v3 = 1;
		break;
	case 2:
		v2 = 1;
		v3 = 2;
		break;
	case 4:
		v2 = 2;
		v3 = 2;
		break;
	default:
		v2 = a2;
		v3 = a2;
		break;
	}
	v4 = a1[1];
	v5 = v4 + v3;
	if (v4 >= v4 + v3) {
		return 1;
	}
	v6 = *a1;
	v7 = *a1 + v2;
	while (1) {
		v8 = *a1;
		if (v6 < v7) {
			v9 = getMemAt(0x5D4594, 1321180 + 4 * (v4 + 2 * v6));
			while (!*(uint32_t*)v9) {
				++v8;
				v9 += 8;
				if (v8 >= v7) {
					goto LABEL_14;
				}
			}
			return 0;
		}
	LABEL_14:
		if (++v4 >= v5) {
			return 1;
		}
	}
}

//----- (004C3140) --------------------------------------------------------
void nox_xxx_cliSummonOnDieOrBanish_4C3140(int a1, void* a2) {
	int* result;  // eax
	int* v3;      // esi
	uint32_t* v4; // eax

	result = (int*)sub_4C31D0(a1);
	v3 = result;
	if (!result) {
		return;
	}
	if (result == *(int**)&dword_5d4594_1321204) {
		nox_xxx_guiHideSummonWindow_4C2470();
	}
	v4 = nox_xxx_netSpriteByCodeDynamic_45A6F0(a1);
	if (v4) {
		v4[30] &= 0xBFFFFFFF;
	}
	sub_4C3030(v3 + 3, *((unsigned char*)v3 + 20), 0);
	sub_4C3210((int)v3);
	sub_4C2F70();
	--dword_5d4594_1321196;
	if (!dword_5d4594_1321196) {
		if (!a2) {
			nox_xxx_clientPlaySoundSpecial_452D80(802, 100);
		}
		*getMemU8Ptr(0x5D4594, 1321200) = 3;
	}
}

//----- (004C31D0) --------------------------------------------------------
char* sub_4C31D0(int a1) {
	int v1;            // eax
	unsigned char* v2; // ecx

	v1 = 0;
	v2 = getMemAt(0x5D4594, 1321052);
	while (!*((uint32_t*)v2 + 2) || *(uint32_t*)v2 != a1) {
		v2 += 32;
		++v1;
		if ((int)v2 >= (int)getMemAt(0x5D4594, 1321180)) {
			return 0;
		}
	}
	return (char*)getMemAt(0x5D4594, 1321052 + 32 * v1);
}

//----- (004C3210) --------------------------------------------------------
int sub_4C3210(int a1) {
	int result; // eax

	result = a1;
	*(uint32_t*)(a1 + 8) = 0;
	return result;
}

//----- (004C3220) --------------------------------------------------------
int nox_xxx_sprite_4C3220(nox_drawable* a1) { return sub_4C31D0(*(uint32_t*)((uint32_t)a1 + 128)) != 0; }

//----- (004C3260) --------------------------------------------------------
int sub_4C3260() { return sub_4C2D60() != 0; }

//----- (004C32A0) --------------------------------------------------------
int nox_xxx_spriteDrawCircleMB_4C32A0(int a1, int a2, int a3, int a4) {
	unsigned char* v4; // ebx
	int v5;            // esi
	int v6;            // edi
	int v8;            // [esp+10h] [ebp-8h]
	int v9;            // [esp+14h] [ebp-4h]

	v8 = (a3 * *getMemU32Ptr(0x587000, 192088)) >> 4;
	v9 = (a3 * *getMemU32Ptr(0x587000, 192092)) >> 4;
	nox_client_drawSetColor_434460(a4);
	nox_client_drawEnableAlpha_434560(1);
	v4 = getMemAt(0x587000, 192220);
	do {
		v5 = (a3 * *((uint32_t*)v4 - 1)) >> 4;
		v6 = (*(uint32_t*)v4 * a3) >> 4;
		nox_client_drawAddPoint_49F500(a1 + v8, a2 + v9);
		nox_client_drawAddPoint_49F500(v5 + a1, v6 + a2);
		nox_client_drawLineFromPoints_49E4B0();
		v4 += 128;
		v8 = v5;
		v9 = v6;
	} while ((int)v4 < (int)getMemAt(0x587000, 194140));
	nox_client_drawAddPoint_49F500(a1 + ((a3 * *getMemU32Ptr(0x587000, 192088)) >> 4),
								   a2 + ((a3 * *getMemU32Ptr(0x587000, 192092)) >> 4));
	nox_client_drawAddPoint_49F500(a1 + v5, a2 + v6);
	nox_client_drawLineFromPoints_49E4B0();
	nox_client_drawEnableAlpha_434560(0);
	return 1;
}

//----- (004C3390) --------------------------------------------------------
int sub_4C3390() {
	*getMemU32Ptr(0x5D4594, 1321220) = nox_xxx_gLoadImg_42F970("VoteInProgress");
	dword_5d4594_1321216 = nox_window_new(0, 136, nox_win_width - 50, nox_win_height / 2 - 100, 50, 50, 0);
	nox_xxx_wndSetIcon_46AE60(*(int*)&dword_5d4594_1321216, *getMemIntPtr(0x5D4594, 1321220));
	nox_window_set_all_funcs(*(uint32_t**)&dword_5d4594_1321216, 0, sub_4C3410, 0);
	nox_window_set_hidden(*(int*)&dword_5d4594_1321216, 1);
	return 1;
}

//----- (004C3410) --------------------------------------------------------
int sub_4C3410(int* a1) {
	int* v1; // esi
	int v2;  // edx
	int v4;  // [esp+4h] [ebp-4h]

	v1 = a1;
	nox_client_wndGetPosition_46AA60(a1, &a1, &v4);
	v2 = v1[25];
	a1 = (int*)((char*)a1 + v1[24]);
	nox_client_drawImageAt_47D2C0(v1[15], (int)a1, v2 + v4);
	return 1;
}

//----- (004C3460) --------------------------------------------------------
int sub_4C3460(int a1) { return nox_window_set_hidden(*(int*)&dword_5d4594_1321216, a1); }

//----- (004C34A0) --------------------------------------------------------
int sub_4C34A0() {
	int result; // eax

	result = nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1321216);
	dword_5d4594_1321216 = 0;
	return result;
}

//----- (004C3500) --------------------------------------------------------
int sub_4C3500() {
	uint32_t* v0; // ecx

	v0 = nox_new_window_from_file("yesno.wnd", 0);
	dword_5d4594_1321224 = v0;
	if (v0) {
		nox_window_setPos_46A9B0(v0, (nox_win_width - 320) / 2, (nox_win_height - 240) / 2);
		v0 = *(uint32_t**)&dword_5d4594_1321224;
	}
	nox_window_set_hidden((int)v0, 1);
	return dword_5d4594_1321224;
}

//----- (004C35B0) --------------------------------------------------------
void nox_common_writecfgfile(char* a1);
int sub_4C35B0(int a1) {
	if (wndIsShown_nox_xxx_wndIsShown_46ACC0(*(int*)&dword_5d4594_1321228)) {
		return 0;
	}
	sub_413A00(0);
	if (a1) {
		nox_window_set_hidden(*(int*)&dword_5d4594_1321228, 1);
	} else {
		sub_4C3620();
		nox_common_writecfgfile("nox.cfg");
		nox_window_set_hidden(*(int*)&dword_5d4594_1321228, 1);
		sub_472280();
		sub_4ADA40();
	}
	return 1;
}

//----- (004C3A60) --------------------------------------------------------
int sub_4C3A60(uint32_t* a1, unsigned int a2, unsigned int a3, int a4) {
	if (a2 < 0x13 || a2 > 0x14) {
		return nox_xxx_wndListboxProcWithoutData10_4A28E0(a1, a2, a3, a4);
	} else {
		return 0;
	}
}

//----- (004C3EB0) --------------------------------------------------------
int sub_4C3EB0(int a1, int a2, unsigned int a3, int a4) {
	switch (a2) {
	case 6:
	case 7:
		sub_4C4100(0x10000);
		nox_xxx_windowFocus_46B500(0);
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1321232);
		nox_window_set_hidden(*(int*)&dword_5d4594_1321232, 1);
		return 1;
	case 10:
	case 11:
		sub_4C4100(0x10002);
		nox_xxx_windowFocus_46B500(0);
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1321232);
		nox_window_set_hidden(*(int*)&dword_5d4594_1321232, 1);
		return 1;
	case 14:
	case 15:
		sub_4C4100(0x10001);
		nox_xxx_windowFocus_46B500(0);
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1321232);
		nox_window_set_hidden(*(int*)&dword_5d4594_1321232, 1);
		return 1;
	case 19:
		sub_4C4100(0x10003);
		nox_xxx_windowFocus_46B500(0);
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1321232);
		nox_window_set_hidden(*(int*)&dword_5d4594_1321232, 1);
		return 1;
	case 20:
		sub_4C4100(0x10004);
		nox_xxx_windowFocus_46B500(0);
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1321232);
		nox_window_set_hidden(*(int*)&dword_5d4594_1321232, 1);
		return 1;
	case 21:
		if (a3 == 1) {
			if (a4 == 2) {
				nox_xxx_windowFocus_46B500(0);
				nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1321232);
				nox_window_set_hidden(*(int*)&dword_5d4594_1321232, 1);
				if (dword_5d4594_1321252) {
					nox_window_call_field_94(*(int*)&dword_5d4594_1321252, 16403, -1, 0);
				}
			}
			return 1;
		} else if (a4 == 1 && sub_4C3FC0(a3)) {
			sub_4C4100(0x10000);
			nox_xxx_windowFocus_46B500(0);
			nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1321232);
			nox_window_set_hidden(*(int*)&dword_5d4594_1321232, 1);
			return 1;
		} else {
			return 0;
		}
	default:
		return 0;
	}
}

//----- (004C3FC0) --------------------------------------------------------
int sub_4C3FC0(unsigned int a1) {
	const wchar2_t* v6 = nox_xxx_keybind_titleByKeyZero_42EA00(a1);
	if (!v6) {
		return 0;
	}
	if (!dword_5d4594_1321252) {
		return 1;
	}
	int v7 = 0;
	int v8 = *(uint32_t*)((uint32_t)dword_5d4594_1321244 + 32);
	if (*(uint16_t*)(v8 + 44) > 0) {
		do {
			const wchar2_t* v9 = (const wchar2_t*)nox_window_call_field_94(*(int*)&dword_5d4594_1321244, 16406, v7, 0);
			if (!nox_wcscmp(v9, v6)) {
				nox_window_call_field_94(*(int*)&dword_5d4594_1321244, 16407, (int)getMemAt(0x587000, 185444), v7);
			}
			++v7;
		} while (v7 < *(short*)(v8 + 44));
	}
	int v10 = 0;
	int v11 = *(uint32_t*)((uint32_t)dword_5d4594_1321248 + 32);
	if (*(uint16_t*)(v11 + 44) > 0) {
		do {
			const wchar2_t* v12 = (const wchar2_t*)nox_window_call_field_94(*(int*)&dword_5d4594_1321248, 16406, v10, 0);
			if (!nox_wcscmp(v12, v6)) {
				nox_window_call_field_94(*(int*)&dword_5d4594_1321248, 16407, (int)getMemAt(0x587000, 185448), v10);
			}
			++v10;
		} while (v10 < *(short*)(v11 + 44));
	}
	nox_window_call_field_94(*(int*)&dword_5d4594_1321252, 16407, (int)v6,
							 *(uint32_t*)(*(uint32_t*)(dword_5d4594_1321252 + 32) + 48));
	nox_window_call_field_94(*(int*)&dword_5d4594_1321252, 16403, -1, 0);
	dword_5d4594_1321252 = 0;
	return 1;
}

//----- (004C4100) --------------------------------------------------------
int sub_4C4100(unsigned int a1) {
	char* v1;          // edi
	int v2;            // esi
	int v3;            // ebx
	const wchar2_t* v4; // eax
	int v5;            // esi
	int v6;            // ebx
	const wchar2_t* v7; // eax

	v1 = nox_xxx_keybind_titleByKey_42EA00(a1);
	if (!dword_5d4594_1321252) {
		return 1;
	}
	v2 = 0;
	v3 = *(uint32_t*)((uint32_t)dword_5d4594_1321244 + 32);
	if (*(uint16_t*)(v3 + 44) > 0) {
		do {
			v4 = (const wchar2_t*)nox_window_call_field_94(*(int*)&dword_5d4594_1321244, 16406, v2, 0);
			if (!nox_wcscmp(v4, (const wchar2_t*)v1)) {
				nox_window_call_field_94(*(int*)&dword_5d4594_1321244, 16407, (int)getMemAt(0x587000, 185452), v2);
			}
			++v2;
		} while (v2 < *(short*)(v3 + 44));
	}
	v5 = 0;
	v6 = *(uint32_t*)((uint32_t)dword_5d4594_1321248 + 32);
	if (*(uint16_t*)(v6 + 44) > 0) {
		do {
			v7 = (const wchar2_t*)nox_window_call_field_94(*(int*)&dword_5d4594_1321248, 16406, v5, 0);
			if (!nox_wcscmp(v7, (const wchar2_t*)v1)) {
				nox_window_call_field_94(*(int*)&dword_5d4594_1321248, 16407, (int)getMemAt(0x587000, 185456), v5);
			}
			++v5;
		} while (v5 < *(short*)(v6 + 44));
	}
	nox_window_call_field_94(*(int*)&dword_5d4594_1321252, 16407, (int)v1,
							 *(uint32_t*)(*(uint32_t*)(dword_5d4594_1321252 + 32) + 48));
	nox_window_call_field_94(*(int*)&dword_5d4594_1321252, 16403, -1, 0);
	dword_5d4594_1321252 = 0;
	return 1;
}

//----- (004C4220) --------------------------------------------------------
int sub_4C4220() {
	int result; // eax

	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1321228);
	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1321232);
	result = 0;
	dword_5d4594_1321228 = 0;
	dword_5d4594_1321232 = 0;
	dword_5d4594_1321236 = 0;
	dword_5d4594_1321240 = 0;
	dword_5d4594_1321244 = 0;
	dword_5d4594_1321248 = 0;
	return result;
}

//----- (004C4260) --------------------------------------------------------
void sub_4C4260() {
	sub_4C3B70();
	nox_xxx_wndShowModalMB_46A8C0(*(int*)&dword_5d4594_1321228);
	sub_413A00(1);
}

//----- (004C4280) --------------------------------------------------------
int sub_4C4280() { return wndIsShown_nox_xxx_wndIsShown_46ACC0(*(int*)&dword_5d4594_1321228) == 0; }

//----- (004C4770) --------------------------------------------------------
void nox_xxx_drawObject_4C4770_draw(nox_draw_viewport_t* vp, nox_drawable* dr, void* img) {
	int* a1 = vp;
	unsigned char* a2 = dr;
	int a3 = img;
	unsigned char* v3;  // esi
	char* v4;           // eax
	uint32_t* v5;       // edi
	uint32_t* v6;       // eax
	int v7;             // ecx
	int v8;             // edi
	int v9;             // ecx
	int v10;            // ebp
	int v11;            // edx
	int v12;            // eax
	int v13;            // ebx
	int v14;            // edx
	int v15;            // ebx
	int v16;            // edi
	int v17;            // ebx
	unsigned char v18;  // al
	unsigned char v19;  // cl
	int v20;            // ecx
	int v21;            // edx
	int v22;            // ecx
	int v23;            // edi
	int v24;            // ebx
	int v25;            // ebx
	int v26;            // edi
	int v27;            // ecx
	unsigned char* v28; // edi
	short v29;          // ax
	int v30;            // ebp
	int v31;            // edx
	int v32;            // eax
	int v33;            // eax
	int v34;            // edx
	int v35;            // ecx
	double v36;         // st7
	double v37;         // st7
	unsigned char v38;  // bl
	unsigned char v39;  // al
	int v40;            // edx
	int v41;            // ecx
	int v42;            // edi
	int v43;            // ebx
	int v44;            // eax
	int v45;            // eax
	int v46;            // eax
	int v47;            // ebp
	int v48;            // eax
	int v49;            // eax
	int v50;            // eax
	int v53;            // [esp+18h] [ebp-30h]
	int v54;            // [esp+1Ch] [ebp-2Ch]
	int2 a1a;           // [esp+20h] [ebp-28h]
	int2 a2a;           // [esp+28h] [ebp-20h]
	int a4;             // [esp+30h] [ebp-18h]
	int a3a;            // [esp+34h] [ebp-14h]
	int v59a;           // [esp+38h] [ebp-10h]
	int v59b;           // [esp+38h] [ebp-10h]
	long long v60;      // [esp+40h] [ebp-8h]
	unsigned char v61;  // [esp+4Ch] [ebp+4h]
	unsigned char v62;  // [esp+4Ch] [ebp+4h]
	unsigned char v63;  // [esp+50h] [ebp+8h]
	unsigned char v64;  // [esp+50h] [ebp+8h]

	v53 = 0;
	v54 = 0;
	if (!dword_5d4594_1321520) {
		dword_5d4594_1321520 = nox_xxx_getTTByNameSpriteMB_44CFC0("Ghost");
	}
	v3 = a2;
	if (a2[112] & 4) {
		v4 = nox_common_playerInfoGetByID_417040(*((uint32_t*)a2 + 32));
		if (nox_player_netCode_85319C == *((uint32_t*)a2 + 32)) {
			if (v4 && v4[3680] & 1) {
				v54 = 1;
			}
		} else if (v4 && v4[3680] & 1) {
			return;
		}
		v5 = nox_xxx_objGetTeamByNetCode_418C80(nox_player_netCode_85319C);
		if (v5) {
			v6 = nox_xxx_objGetTeamByNetCode_418C80(*((uint32_t*)a2 + 32));
			if (v6) {
				if (nox_player_netCode_85319C == *((uint32_t*)a2 + 32) ||
					nox_xxx_servCompareTeams_419150((int)v5, (int)v6)) {
					v53 = 1;
				}
			}
		}
	}
	v7 = *((uint32_t*)a2 + 3);
	a3a = 0;
	v8 = *a1;
	v9 = v7 - a1[4];
	v10 = a1[1];
	v11 = *((short*)a2 + 53);
	v59a = v9 + *a1 - *a2;
	v12 = *((uint32_t*)a2 + 4);
	v13 = *((uint32_t*)a2 + 4) - v11 - a2[1] - *((short*)a2 + 52);
	v14 = a1[5];
	v59b = v10 + v13 - v14;
	a4 = nox_win_width;
	if (*getMemU32Ptr(0x587000, 80808)) {
		v15 = *((uint32_t*)a2 + 30);
		if (!(v15 & 0x40000000) && !(v15 & 1) && (a2[112] & 0x80u) != 0) {
			v16 = v9 + v8;
			a1a.field_0 = v16;
			v17 = v10 + v12 - v14;
			v18 = a2[299];
			v19 = a2[299];
			a1a.field_4 = v17;
			v20 = 8 * v19;
			v21 = v16 + *getMemIntPtr(0x587000, 196184 + v20);
			v22 = v17 + *getMemIntPtr(0x587000, 196188 + v20);
			a2a.field_0 = v21;
			a2a.field_4 = v22;
			if (v18 < 0x18u && v18 && (v18 < 8u || v18 > 0x10u)) {
				v23 = (v21 + v16) >> 1;
				v24 = (v22 + v17) >> 1;
				if (sub_4C5630(v23 - 5, v23 - 5, v24)) {
					a1a.field_0 -= 2;
					a2a.field_0 -= 2;
				} else if (sub_4C5630(v23 + 5, v23 + 5, v24)) {
					a1a.field_0 += 2;
					a2a.field_0 += 2;
				}
			} else {
				v25 = (v22 + v17) >> 1;
				v26 = (v21 + v16) >> 1;
				if (sub_4C5630(v26, v26, v25 - 5)) {
					v27 = a2a.field_4 - 2;
					a1a.field_4 -= 2;
					a2a.field_4 = v27;
				} else if (sub_4C5630(v26, v26, v25 + 5)) {
					v27 = a2a.field_4 + 2;
					a1a.field_4 += 2;
					a2a.field_4 = v27;
				}
			}
			if (!sub_4C42A0(&a1a, &a2a, &a3a, &a4)) {
				return;
			}
		}
	}
	if ((*((uint32_t*)a2 + 28) & 0x80000) && (*((uint32_t*)a2 + 30) & 0x1000000) || (*((uint32_t*)a2 + 30) & 0x40000000)) {
		v28 = getMemAt(0x587000, 185472);
	} else {
		v28 = sub_469920((int2*)(a2 + 12));
	}
	v29 = *((uint16_t*)a2 + 52);
	int a2b = 0;
	if (v29 >= 0) {
		v30 = 0;
	} else {
		if (*getMemIntPtr(0x5D4594, 1321512) < 0 && a2 == *(unsigned char**)getMemAt(0x5D4594, 1321516)) {
			a2b = 1;
		}
		v30 = -v29;
	}
	*getMemU32Ptr(0x5D4594, 1321516) = a2;
	v31 = *((short*)a2 + 52);
	a1a.field_0 = v30;
	*getMemU32Ptr(0x5D4594, 1321512) = v31;
	if (nox_client_drawable_testBuff_4356C0((int)a2, 25)) {
		nox_xxx_draw_434600(1);
		nox_draw_setColorMultAndIntensity_433E40(nox_color_blue_2650684);
	} else if ((a2[112] & 2) && (v32 = *((uint32_t*)a2 + 30), v32 & 0x40000000) && !(v32 & 0x8020)) {
		nox_xxx_draw_434600(1);
		nox_draw_setColorMultAndIntensity_433E40(nox_color_blue_2650684);
	} else {
		sub_4345F0(1);
		LOBYTE(v34) = v28[8];
		v33 = v28[4];
		LOBYTE(v35) = *v28;
		nox_draw_setColorMultAndIntensityRGB_433CD0(v35, v33, v34);
	}
	if (*((uint32_t*)a2 + 120)) {
		v60 = (unsigned int)(gameFrame() - *((uint32_t*)a2 + 85));
		v36 = (double)v60;
		v60 = gameFPS();
		v37 = 1.0 - v36 / (double)(int)gameFPS();
		if (v37 < 0.0) {
			v37 = 0.001;
		}
		v38 = (long long)(v37 * 255.0);
		v63 = (long long)(v37 * 255.0);
		if (*((uint32_t*)v3 + 27) == dword_5d4594_1321520) {
			v39 = sub_4C4EC0(a1, (int)v3);
			if (v39 < v38) {
				v64 = v39;
				nox_client_drawEnableAlpha_434560(1);
				nox_client_drawSetAlpha_434580(v64);
			} else {
				nox_client_drawEnableAlpha_434560(1);
				nox_client_drawSetAlpha_434580(v63);
			}
		} else {
			if (*((uint32_t*)v3 + 30) & 0x4000000) {
				v63 = v38 >> 1;
			}
			nox_client_drawEnableAlpha_434560(1);
			nox_client_drawSetAlpha_434580(v63);
		}
	} else if (*((uint32_t*)a2 + 27) == dword_5d4594_1321520) {
		v61 = sub_4C4EC0(a1, (int)a2);
		nox_client_drawEnableAlpha_434560(1);
		nox_client_drawSetAlpha_434580(v61);
	} else if (*((uint32_t*)a2 + 30) & 0x4000000) {
		nox_client_drawEnableAlpha_434560(1);
		nox_client_drawSetAlpha_434580(0x80u);
	}
	if (!(!nox_client_drawable_testBuff_4356C0((int)v3, 0) && !v54 &&
		(*((uint32_t*)v3 + 27) != dword_5d4594_1321520 || !*getMemU32Ptr(0x852978, 8) ||
		 !nox_client_drawable_testBuff_4356C0(*getMemIntPtr(0x852978, 8), 21)))) {
		if (dword_8531A0_2576 && (*(uint8_t*)(dword_8531A0_2576 + 3680) & 1)) {
			nox_client_drawEnableAlpha_434560(1);
			nox_client_drawSetAlpha_434580(0x80u);
		} else {
			v40 = *((uint32_t*)v3 + 8);
			v41 = *((uint32_t*)v3 + 9);
			v42 = *((uint32_t*)v3 + 3) - v40;
			v43 = *((uint32_t*)v3 + 4) - v41;
			v44 = *((uint32_t*)v3 + 3) - v40;
			if (v42 < 0) {
				v44 = v40 - *((uint32_t*)v3 + 3);
			}
			if (v44 < 4) {
				v45 = *((uint32_t*)v3 + 4) - v41;
				if (v43 < 0) {
					v45 = v41 - *((uint32_t*)v3 + 4);
				}
				if (v45 < 4) {
					v46 = *getMemU32Ptr(0x852978, 8);
					if (!*getMemU32Ptr(0x852978, 8)) {
						goto LABEL_82;
					}
					if (!nox_client_drawable_testBuff_4356C0(*getMemIntPtr(0x852978, 8), 21)) {
						LOWORD(v4) = v53;
						if (!v53) {
							return;
						}
					}
				}
			}
			v46 = *getMemU32Ptr(0x852978, 8);
		LABEL_82:
			if (v42 || v43) {
				v47 = *((uint32_t*)v3 + 5) - *((uint32_t*)v3 + 10);
				if (!v47) {
					v47 = 1;
				}
				v48 = nox_double2int(sqrt((double)(v42 * v42 + v43 * v43))) / v47;
				v30 = a1a.field_0;
				v42 = v48;
				v46 = *getMemU32Ptr(0x852978, 8);
			}
			if (!v46) {
				goto LABEL_123;
			}
			if (v54) {
				LOBYTE(v49) = -128;
				v62 = v49;
				goto LABEL_B;
			}
			if (!(v53 || !nox_client_drawable_testBuff_4356C0(v46, 21))) {
				nox_xxx_draw_434600(1);
				nox_draw_setColorMultAndIntensity_433E40(dword_8531A0_2572);
				v62 = -1;
				nox_client_drawEnableAlpha_434560(1);
				nox_client_drawSetAlpha_434580(v62);
				goto LABEL_A;
			}
		LABEL_123:
			if (v42 < 8) {
				v49 = (v42 << 7) / 8;
				v62 = v49;
				if (!(uint8_t)v49) {
					LOBYTE(v49) = 1;
					v62 = v49;
				}
			} else {
				LOBYTE(v49) = -128;
				v62 = v49;
			}
		LABEL_B:
			if (v53 && (unsigned char)v49 <= 1u) {
				if ((v50 = *((uint32_t*)v3 + 28), v50 & 2) && (*((uint32_t*)v3 + 69) == 8) ||
					(v50 & 4) && !*((uint32_t*)v3 + 69)) {
					nox_xxx_draw_434600(1);
					nox_draw_setColorMultAndIntensity_433E40(dword_8531A0_2572);
					v62 = -128;
				}
			}
			nox_client_drawEnableAlpha_434560(1);
			nox_client_drawSetAlpha_434580(v62);
		}
	}
LABEL_A:
	if (!(v3[112] & 4) && nox_client_drawable_testBuff_4356C0((int)v3, 23) && !nox_common_gameFlags_check_40A5C0(2048)) {
		nox_xxx_draw_434600(1);
		if ((unsigned char)gameFrame() & 1) {
			nox_draw_setColorMultAndIntensity_433E40(nox_color_white_2523948);
		} else {
			nox_draw_setColorMultAndIntensity_433E40(nox_color_blue_2650684);
		}
	}
	if (!nox_client_drawable_testBuff_4356C0((int)v3, 23) && nox_client_drawable_testBuff_4356C0((int)v3, 11)) {
		nox_xxx_draw_434600(1);
		nox_draw_setColorMultAndIntensity_433E40(*getMemIntPtr(0x85B3FC, 956));
	}
	if (a2b) {
		nox_xxx_wndDraw_49F7F0();
		sub_49F7C0_def();
	} else {
		sub_47D370(v30);
	}
	nox_client_drawImageAt_47D2C0(a3, v59a, v59b);
	sub_4345F0(0);
	nox_client_drawEnableAlpha_434560(0);
	nox_xxx_draw_434600(0);
	if (a2b) {
		sub_49F860();
	}
	*((uint16_t*)v3 + 1) = 1;
	LOWORD(v4) = *getMemU16Ptr(0x973F18, 88);
	*((uint16_t*)v3 + 2) = *getMemU16Ptr(0x973F18, 88);
	*((uint16_t*)v3 + 3) = *getMemU16Ptr(0x973F18, 76);
	*((uint32_t*)v3 + 2) = a3;
}
// 4C4AD0: variable 'v35' is possibly undefined
// 4C4AD0: variable 'v34' is possibly undefined

//----- (004C4EC0) --------------------------------------------------------
char sub_4C4EC0(uint32_t* a1, int a2) {
	int v3; // ecx
	int v4; // eax
	int v5; // ecx

	if (*getMemU32Ptr(0x852978, 8) && nox_client_drawable_testBuff_4356C0(*getMemIntPtr(0x852978, 8), 21)) {
		return -1;
	}
	v3 = *(uint32_t*)(a2 + 12) - a1[8] / 2 - a1[4];
	v4 = *(uint32_t*)(a2 + 16) - a1[9] / 2 - a1[5];
	v5 = 128 - ((v3 * v3 + v4 * v4) << 7) / *getMemU32Ptr(0x587000, 185464);
	if (v5 < 0) {
		return 0;
	}
	if (v5 > 128) {
		LOBYTE(v5) = -128;
	}
	return v5;
}

//----- (004C4F40) --------------------------------------------------------
short nox_xxx_drawShinySpot_4C4F40(nox_draw_viewport_t* vp, nox_drawable* dr) {
	uint32_t* a1 = vp;
	int a2 = dr;
	char* v2;        // eax
	int v3;          // edi
	signed int v4;   // esi
	unsigned int v5; // eax
	signed int v6;   // edx

	v2 = *(char**)getMemAt(0x5D4594, 1321524);
	if (!*getMemU32Ptr(0x5D4594, 1321524)) {
		v2 = nox_xxx_gLoadAnim_42FA20("ShinySpot");
		*getMemU32Ptr(0x5D4594, 1321524) = v2;
	}
	v3 = *((uint32_t*)v2 + 24);
	v4 = *(unsigned char*)(v3 + 8);
	v5 = (gameFrame() + *(uint32_t*)(a2 + 128)) / (unsigned int)(8 * v4);
	v6 = ((gameFrame() + *(uint32_t*)(a2 + 128)) % (unsigned int)(8 * v4)) >> 1;
	if (v6 < v4) {
		nox_client_drawImageAt_47D2C0(
			*(uint32_t*)(*(uint32_t*)(v3 + 4) + 4 * v6), *(uint32_t*)(a2 + 12) - a1[4] + *a1 - 64,
			a1[1] - *(short*)(a2 + 106) - *(short*)(a2 + 104) - a1[5] + *(uint32_t*)(a2 + 16) - 64);
	}
	return v5;
}

//----- (004C4FD0) --------------------------------------------------------
int nox_xxx_colorInit_4C4FD0() {
	int result; // eax

	*getMemU32Ptr(0x5D4594, 1321536) = nox_color_rgb_4344A0(255, 200, 255);
	*getMemU32Ptr(0x5D4594, 1321796) = nox_color_rgb_4344A0(255, 0, 255);
	result = nox_color_rgb_4344A0(100, 40, 100);
	*getMemU32Ptr(0x5D4594, 1321532) = result;
	return result;
}

//----- (004C5020) --------------------------------------------------------
int sub_4C5020(int a1) {
	int result; // eax

	result = dword_5d4594_1321800;
	if (*(int*)&dword_5d4594_1321800 < 32) {
		result = dword_5d4594_1321800 + 1;
		*getMemU32Ptr(0x5D4594, 1321532 + 8 * result) = *(uint32_t*)(a1 + 1);
		*getMemU32Ptr(0x5D4594, 1321536 + 8 * result) = *(uint32_t*)(a1 + 5);
		dword_5d4594_1321800 = result;
	}
	return result;
}

//----- (004C5050) --------------------------------------------------------
void sub_4C5050() { dword_5d4594_1321800 = 0; }

//----- (004C5060) --------------------------------------------------------
int sub_498C20(int2* a1, int2* a2, int a3);
int sub_4C5060(nox_draw_viewport_t* a1p) {
	uint32_t* a1 = a1p;
	int result;         // eax
	unsigned short* v2; // ebp
	int v3;             // esi
	int v4;             // edi
	int v5;             // ebx
	int v6;             // ecx
	int v7;             // edi
	int v8;             // esi
	int v10;            // [esp+0h] [ebp-24h]
	int2 a1a;           // [esp+4h] [ebp-20h]
	int2 v12;           // [esp+Ch] [ebp-18h]
	int2 a2;            // [esp+14h] [ebp-10h]
	int2 v14;           // [esp+1Ch] [ebp-8h]

	result = *getMemU32Ptr(0x852978, 8);
	if (*getMemU32Ptr(0x852978, 8)) {
		result = dword_5d4594_1321800;
		v10 = 0;
		if (dword_5d4594_1321800 > 0) {
			v2 = getMemU16Ptr(0x5D4594, 1321540);
			do {
				v3 = *a1 + *v2 - a1[4];
				v4 = a1[1] + v2[1] - a1[5];
				v5 = sub_4992B0(v3, v4);
				v6 = v2[3] - v2[1];
				if (v3 <= 0 || v3 >= a1[8] - 1 || v4 <= 0 || v4 >= a1[9] - 1) {
					v5 = 0;
				}
				v14.field_0 = v3 + v2[2] - *v2;
				v14.field_4 = v4 + v6;
				v12.field_0 = v3;
				v12.field_4 = v4;
				v7 = sub_498C20(&v12, &v14, (int)a1);
				if (v7) {
					v8 = 0;
					for (a1a = v12; v8 < v7; v5 = 1 - v5) {
						*(nox_point*)&a2 = sub_499290(v8);
						if (v5) {
							sub_4C51D0(&a1a, &a2);
						}
						a1a = a2;
						++v8;
					}
					if (v5) {
						sub_4C51D0(&a1a, &v14);
					}
				} else if (v5) {
					sub_4C51D0(&v12, &v14);
				}
				result = v10 + 1;
				v2 += 4;
				++v10;
			} while (v10 < *(int*)&dword_5d4594_1321800);
		}
	}
	return result;
}
// 4C5151: variable 'v9' is possibly undefined

//----- (004C51D0) --------------------------------------------------------
int sub_4C51D0(int2* a1, int2* a2) {
	int v2;  // esi
	int v3;  // edi
	int v4;  // ebx
	int v5;  // ebp
	int v6;  // esi
	int v7;  // ecx
	int v8;  // eax
	int v9;  // ebx
	int v11; // ebp

	v2 = a1->field_4;
	v3 = a1->field_0;
	v4 = a2->field_0 - a1->field_0;
	v5 = a2->field_4 - v2;
	nox_client_drawSetColor_434460(*getMemIntPtr(0x5D4594, 1321532));
	nox_client_drawEnableAlpha_434560(1);
	nox_client_drawAddPoint_49F500(v3, v2);
	nox_xxx_rasterPointRel_49F570(v4, v5);
	nox_client_drawLineFromPoints_49E4B0();
	nox_client_drawEnableAlpha_434560(0);
	v6 = v2 - 22;
	nox_client_drawSetColor_434460(*getMemIntPtr(0x5D4594, 1321536));
	nox_client_drawAddPoint_49F500(v3, v6);
	nox_xxx_rasterPointRel_49F570(v4, v5);
	nox_client_drawLineFromPoints_49E4B0();
	nox_client_drawSetColor_434460(*getMemIntPtr(0x5D4594, 1321796));
	v7 = v4;
	if (v4 < 0) {
		v7 = -v4;
	}
	v8 = v5;
	if (v5 < 0) {
		v8 = -v5;
	}
	if (v7 <= v8) {
		nox_client_drawAddPoint_49F500(v3 - 1, v6);
		v11 = v6 + v5;
		nox_client_drawAddPoint_49F500(v4 + v3 - 1, v11);
		nox_client_drawLineFromPoints_49E4B0();
		nox_client_drawAddPoint_49F500(v3 + 1, v6);
		nox_client_drawAddPoint_49F500(v4 + v3 + 1, v11);
	} else {
		nox_client_drawAddPoint_49F500(v3, v6 - 1);
		v9 = v3 + v4;
		nox_client_drawAddPoint_49F500(v9, v6 + v5 - 1);
		nox_client_drawLineFromPoints_49E4B0();
		nox_client_drawAddPoint_49F500(v3, v6 + 1);
		nox_client_drawAddPoint_49F500(v9, v6 + v5 + 1);
	}
	return nox_client_drawLineFromPoints_49E4B0();
}

//----- (004CA540) --------------------------------------------------------
int nox_xxx_sprite_4CA540(uint32_t* a1, int a2) {
	uint32_t* v2; // esi
	double v3;    // st7
	double v4;    // st6
	double v5;    // st5
	double v6;    // st4
	int v7;       // eax
	int v8;       // edi
	int v9;       // eax
	float v11;    // [esp+0h] [ebp-14h]
	float v12;    // [esp+0h] [ebp-14h]
	float v13;    // [esp+10h] [ebp-4h]
	float v14;    // [esp+1Ch] [ebp+8h]

	v2 = (uint32_t*)a2;
	v3 = 0.0;
	v4 = *(float*)(a2 + 468);
	v5 = *(float*)(a2 + 472);
	v6 = 0.0;
	v7 = gameFrame() - *(uint32_t*)(a2 + 316) + 1;
	do {
		--v7;
		v13 = -(v5 * *(float*)(a2 + 476));
		v4 = v4 - v4 * *(float*)(a2 + 476);
		v5 = v5 + v13;
		v3 = v3 + v4;
		v6 = v6 + v5;
	} while (v7);
	v14 = v6;
	v11 = (double)(int)v2[81] + v3;
	v8 = nox_float2int(v11);
	v12 = (double)(int)v2[82] + v14;
	v9 = nox_float2int(v12);
	if (v8 > 0 && v9 > 0 && v8 < 5888 && v9 < 5888) {
		nox_xxx_updateSpritePosition_49AA90(v2, v8, v9);
		if (sub_4992B0(*a1 + v2[3] - a1[4], v2[4] + a1[1] - a1[5])) {
			return 1;
		}
	}
	nox_xxx_spriteDeleteStatic_45A4E0_drawable((int)v2);
	return 0;
}

//----- (004CA650) --------------------------------------------------------
int sub_4CA650(int a1, int a2) {
	int v2;             // esi
	int v3;             // eax
	int v4;             // ebx
	int v5;             // edi
	int v6;             // eax
	int v7;             // ecx
	int v8;             // eax
	int v9;             // ebp
	int v10;            // eax
	int v11;            // edi
	unsigned short v12; // dx
	int v13;            // ebp
	int v14;            // eax
	int result;         // eax
	int v16;            // [esp+10h] [ebp-10h]
	int v17;            // [esp+28h] [ebp+8h]

	v2 = a2;
	v3 = *(uint32_t*)(a2 + 16);
	v4 = *(unsigned short*)(a2 + 434) - v3;
	v5 = *(unsigned short*)(a2 + 432) - *(uint32_t*)(a2 + 12);
	v6 = sub_48C6B0(v5, v4);
	v7 = v6;
	v8 = *(unsigned char*)(a2 + 443);
	v17 = v8;
	++v7;
	v9 = *(uint32_t*)(v2 + 12);
	v10 = v5 * v8 / v7;
	v11 = *(uint32_t*)(v2 + 16);
	v16 = v9 + v10;
	v12 = *(uint16_t*)(v2 + 432);
	v13 = v9 - v12;
	v14 = v11 + v4 * v17 / v7;
	if (v7 <= 10 ||
		v13 * (v16 - v12) + (v11 - *(unsigned short*)(v2 + 434)) * (v14 - *(unsigned short*)(v2 + 434)) < 0) {
		nox_xxx_spriteDeleteStatic_45A4E0_drawable(v2);
		result = 0;
	} else {
		nox_xxx_updateSpritePosition_49AA90((uint32_t*)v2, v16, v14);
		result = 1;
	}
	return result;
}
// 4CA67E: variable 'v6' is possibly undefined

//----- (004CA720) --------------------------------------------------------
int sub_4CA720(int a1, int a2) {
	int v2;     // ebx
	int v3;     // ebp
	int v4;     // ecx
	int v5;     // edx
	int v6;     // ebx
	int v7;     // kr00_4
	int v8;     // ebx
	int result; // eax

	v2 = gameFrame() - *(uint32_t*)(a2 + 316);
	if (v2 >= 60 || (v3 = *(unsigned short*)(a2 + 432), abs(*(uint32_t*)(a2 + 12) - (unsigned short)v3) < 10) &&
						abs(*(uint32_t*)(a2 + 16) - *(unsigned short*)(a2 + 434)) < 10) {
		nox_xxx_spriteDeleteStatic_45A4E0_drawable(a2);
		result = 0;
	} else {
		v4 = *(short*)(a2 + 440) - v2 * *(short*)(a2 + 440) / 60;
		v5 = (v2 << 8) / 120;
		if (*(uint8_t*)(a2 + 443)) {
			v6 = v5 + *(unsigned char*)(a2 + 442);
		} else {
			v6 = *(unsigned char*)(a2 + 442) - v5;
		}
		if (v6 < 0) {
			v6 += (unsigned int)(255 - v6) >> 8 << 8;
		}
		if (v6 >= 256) {
			v6 += -256 * ((unsigned int)v6 >> 8);
		}
		v7 = v4 * *getMemIntPtr(0x587000, 192088 + 8 * v6);
		v8 = *(unsigned short*)(a2 + 434) + v4 * *getMemI32Ptr(0x587000, 192092 + 8 * v6) / 16;
		nox_xxx_updateSpritePosition_49AA90((uint32_t*)a2, v3 + v7 / 16, v8);
		*(uint32_t*)(a2 + 32) = v3 + v7 / 16;
		*(uint32_t*)(a2 + 36) = v8;
		result = 1;
	}
	return result;
}

//----- (004CAF10) --------------------------------------------------------
nox_window* nox_gui_newProgressBar_4CAF10(int a1, int a2, int a3, int a4, int a5, int a6, uint32_t* a7) {
	int v7;       // eax
	uint32_t* v8; // esi

	v7 = a7[2];
	if (!(v7 & 0x1000)) {
		return 0;
	}
	v8 = nox_window_new(a1, a2 & 0xFFFFFFF7, a3, a4, a5, a6, sub_4CAF80);
	sub_4CAFB0((int)v8);
	if (v8) {
		if (!a7[4]) {
			a7[4] = v8;
		}
		nox_gui_windowCopyDrawData_46AF80((int)v8, a7);
	}
	return v8;
}

//----- (004CAF80) --------------------------------------------------------
int sub_4CAF80(int a1, int a2, int a3, int a4) {
	if (a2 == 16416 && a3 >= 0 && a3 <= 100) {
		*(uint32_t*)(a1 + 32) = a3;
	}
	return 0;
}

//----- (004CAFB0) --------------------------------------------------------
int sub_4CAFB0(int a1) {
	int result; // eax

	result = a1;
	if (a1) {
		if ((signed char)*(uint8_t*)(a1 + 4) >= 0) {
			result = nox_window_set_all_funcs((uint32_t*)a1, 0, sub_4CAFF0, 0);
		} else {
			result = nox_window_set_all_funcs((uint32_t*)a1, 0, sub_4CB1A0, 0);
		}
	}
	return result;
}

//----- (004CAFF0) --------------------------------------------------------
int sub_4CAFF0(uint32_t* a1, uint32_t* a2) {
	int v2;          // ecx
	int v3;          // ebp
	int v4;          // edi
	int v5;          // edi
	int v6;          // edi
	int v7;          // ebp
	int xLeft;       // [esp+10h] [ebp-98h]
	int yTop;        // [esp+14h] [ebp-94h]
	int v11;         // [esp+18h] [ebp-90h]
	int v12;         // [esp+1Ch] [ebp-8Ch]
	int v13;         // [esp+20h] [ebp-88h]
	int v14;         // [esp+24h] [ebp-84h]
	wchar2_t v15[64]; // [esp+28h] [ebp-80h]

	v2 = a2[17];
	v3 = a2[9];
	v14 = a2[7];
	v4 = a2[5];
	v11 = v2;
	nox_client_wndGetPosition_46AA60(a1, &xLeft, &yTop);
	if (v4 != 0x80000000) {
		nox_client_drawSetColor_434460(v4);
		nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, a1[2], a1[3]);
	}
	if (v3 != 0x80000000) {
		v5 = a1[2] * a1[8] / 100;
		nox_client_drawSetColor_434460(v3);
		nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, v5, a1[3]);
	}
	if (v11 != 0x80000000) {
		if ((a1[1] & 0x2000) == 0x2000) {
			nox_draw_enableTextSmoothing_43F670(1);
		}
		nox_swprintf(v15, L"%i%%", a1[8]);
		nox_xxx_drawGetStringSize_43F840(a2[50], v15, &v13, &v12, 64);
		v6 = xLeft + a1[2] / 2 - v13 / 2;
		v7 = a1[3] / 2 - v12 / 2 + yTop + 1;
		nox_xxx_drawSetTextColor_434390(v11);
		nox_xxx_drawStringWrap_43FAF0(a2[50], v15, v6, v7, a1[2], 0);
		nox_draw_enableTextSmoothing_43F670(0);
	}
	if (v14 != 0x80000000) {
		nox_client_drawSetColor_434460(v14);
		nox_client_drawBorderLines_49CC70(xLeft, yTop, a1[2], a1[3]);
	}
	return 1;
}

//----- (004CB1A0) --------------------------------------------------------
int sub_4CB1A0(uint32_t* a1, int a2) {
	int v2;       // ebx
	uint32_t* v3; // esi
	int v4;       // edi
	int xLeft;    // [esp+Ch] [ebp-8h]
	int yTop;     // [esp+10h] [ebp-4h]

	v2 = *(uint32_t*)(a2 + 24);
	v3 = a1;
	nox_client_wndGetPosition_46AA60(a1, &xLeft, &yTop);
	nox_xxx_wndDraw_49F7F0();
	v4 = v3[2] * v3[8] / 100;
	nox_client_copyRect_49F6F0(xLeft, yTop, v4, v3[3]);
	if (v2 && v4 > 0) {
		nox_client_drawImageAt_47D2C0(v2, xLeft, yTop);
	}
	sub_49F860();
	return 1;
}

//----- (004CB880) --------------------------------------------------------
int sub_4CB880() {
	int result;    // eax
	uint32_t** v1; // eax
	uint32_t* v2;  // esi
	uint32_t* v3;  // esi
	int v4;        // esi
	int v5;        // edi
	uint32_t* v6;  // eax
	int v7;        // eax
	uint32_t* v8;  // eax
	uint32_t* v9;  // eax

	nox_game_addStateCode_43BDD0(900);
	result = nox_new_window_from_file("InputCfg.wnd", sub_4CBE70);
	dword_5d4594_1522604 = result;
	if (result) {
		nox_xxx_wndSetWindowProc_46B300(result, sub_4A18E0);
		result = nox_gui_makeAnimation_43C5B0(*(uint32_t**)&dword_5d4594_1522604, 0, 0, 0, -480, 0, 20, 0, -40);
		nox_wnd_xxx_1522608 = result;
		if (result) {
			nox_wnd_xxx_1522608->field_0 = 900;
			nox_wnd_xxx_1522608->field_12 = sub_4CBB70;
			nox_wnd_xxx_1522608->fnc_done_out = sub_4CBBB0;
			dword_5d4594_1522616 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1522604, 910);
			dword_5d4594_1522620 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1522604, 911);
			dword_5d4594_1522624 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1522604, 912);
			dword_5d4594_1522628 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1522604, 913);
			result = dword_5d4594_1522616;
			if (dword_5d4594_1522616) {
				v1 = *(uint32_t***)((uint32_t)dword_5d4594_1522616 + 32);
				*v1[7] = 921;
				*v1[8] = 922;
				*v1[9] = 920;
				nox_xxx_wndSetProc_46B2C0(*(int*)&dword_5d4594_1522616, sub_4CBF60);
				sub_46B120(*(uint32_t**)&dword_5d4594_1522620, *(int*)&dword_5d4594_1522616);
				sub_46B120(*(uint32_t**)&dword_5d4594_1522624, *(int*)&dword_5d4594_1522616);
				sub_46B120(*(uint32_t**)&dword_5d4594_1522628, *(int*)&dword_5d4594_1522616);
				nox_xxx_wndSetWindowProc_46B300(*(int*)&dword_5d4594_1522624, sub_4CC140);
				nox_xxx_wndSetWindowProc_46B300(*(int*)&dword_5d4594_1522628, sub_4CC140);
				v2 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1522604, 921);
				nox_window_call_field_94(*(int*)&dword_5d4594_1522620, 16408, (int)v2, 0);
				nox_window_call_field_94(*(int*)&dword_5d4594_1522624, 16408, (int)v2, 0);
				nox_window_call_field_94(*(int*)&dword_5d4594_1522628, 16408, (int)v2, 0);
				v3 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1522604, 922);
				nox_window_call_field_94(*(int*)&dword_5d4594_1522620, 16409, (int)v3, 0);
				nox_window_call_field_94(*(int*)&dword_5d4594_1522624, 16409, (int)v3, 0);
				nox_window_call_field_94(*(int*)&dword_5d4594_1522628, 16409, (int)v3, 0);
				sub_4CBBF0();
				v4 = 971;
				v5 = sub_47DBC0() + 971;
				if (v5 > 971) {
					do {
						v6 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1522604, v4);
						nox_xxx_wnd_46ABB0((int)v6, 1);
						++v4;
					} while (v4 < v5);
				}
				v7 = nox_client_mousePriKey_430AF0();
				v8 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1522604, v7 + 971);
				nox_window_call_field_94((int)v8, 16392, 1, 0);
				dword_5d4594_1522612 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1522604, 980);
				sub_46B120(*(uint32_t**)&dword_5d4594_1522612, 0);
				nox_xxx_wndSetProc_46B2C0(*(int*)&dword_5d4594_1522612, sub_4CBE70);
				nox_xxx_wndSetWindowProc_46B300(*(int*)&dword_5d4594_1522612, sub_4CC170);
				nox_window_set_hidden(*(int*)&dword_5d4594_1522612, 1);
				v9 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1522612, 981);
				sub_46AEE0((int)v9, (int)getMemAt(0x5D4594, 1522636));
				result = 1;
			}
		}
	}
	return result;
}
// 4A18E0: using guessed type int  sub_4A18E0(int, int, int, int);

//----- (004CBB70) --------------------------------------------------------
int sub_4CBB70() {
	sub_4CBD30();
	nox_common_writecfgfile("nox.cfg");
	nox_wnd_xxx_1522608->state = NOX_GUI_ANIM_OUT;
	sub_43BE40(2);
	nox_xxx_clientPlaySoundSpecial_452D80(923, 100);
	return 1;
}

//----- (004CBBB0) --------------------------------------------------------
int sub_4CBBB0() {
	int (*v0)(void); // esi

	v0 = nox_wnd_xxx_1522608->field_13;
	nox_gui_freeAnimation_43C570(nox_wnd_xxx_1522608);
	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1522604);
	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1522612);
	v0();
	return 1;
}

//----- (004CC140) --------------------------------------------------------
int sub_4CC140(uint32_t* a1, unsigned int a2, unsigned int a3, int a4) {
	if (a2 < 0x13 || a2 > 0x14) {
		return nox_xxx_wndListboxProcWithoutData10_4A28E0(a1, a2, a3, a4);
	} else {
		return 0;
	}
}

//----- (004CC170) --------------------------------------------------------
int sub_4CC170(int a1, int a2, char* a3, int a4) {
	switch (a2) {
	case 6:
	case 7:
		sub_4CC3C0(0x10000);
		nox_xxx_windowFocus_46B500(0);
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1522612);
		nox_window_set_hidden(*(int*)&dword_5d4594_1522612, 1);
		return 1;
	case 10:
	case 11:
		sub_4CC3C0(0x10002);
		nox_xxx_windowFocus_46B500(0);
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1522612);
		nox_window_set_hidden(*(int*)&dword_5d4594_1522612, 1);
		return 1;
	case 14:
	case 15:
		sub_4CC3C0(0x10001);
		nox_xxx_windowFocus_46B500(0);
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1522612);
		nox_window_set_hidden(*(int*)&dword_5d4594_1522612, 1);
		return 1;
	case 19:
		sub_4CC3C0(0x10003);
		nox_xxx_windowFocus_46B500(0);
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1522612);
		nox_window_set_hidden(*(int*)&dword_5d4594_1522612, 1);
		return 1;
	case 20:
		sub_4CC3C0(0x10004);
		nox_xxx_windowFocus_46B500(0);
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1522612);
		nox_window_set_hidden(*(int*)&dword_5d4594_1522612, 1);
		return 1;
	case 21:
		if (a3 == (char*)1) {
			if (a4 == 2) {
				nox_xxx_windowFocus_46B500(0);
				nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1522612);
				nox_window_set_hidden(*(int*)&dword_5d4594_1522612, 1);
				if (dword_5d4594_1522632) {
					nox_window_call_field_94(*(int*)&dword_5d4594_1522632, 16403, -1, 0);
				}
			}
			return 1;
		} else if (a4 == 1 && sub_4CC280(a3)) {
			sub_4CC3C0(0x10000);
			nox_xxx_windowFocus_46B500(0);
			nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1522612);
			nox_window_set_hidden(*(int*)&dword_5d4594_1522612, 1);
			return 1;
		} else {
			return 0;
		}
	default:
		return 0;
	}
}

//----- (004CC280) --------------------------------------------------------
int sub_4CC280(unsigned int a1) {
	const wchar2_t* v6 = nox_xxx_keybind_titleByKeyZero_42EA00(a1);
	if (!v6) {
		return 0;
	}
	if (!dword_5d4594_1522632) {
		return 1;
	}
	int v7 = 0;
	int v8 = *(uint32_t*)((uint32_t)dword_5d4594_1522624 + 32);
	if (*(uint16_t*)(v8 + 44) > 0) {
		do {
			const wchar2_t* v9 = (const wchar2_t*)nox_window_call_field_94(*(int*)&dword_5d4594_1522624, 16406, v7, 0);
			if (!nox_wcscmp(v9, v6)) {
				nox_window_call_field_94(*(int*)&dword_5d4594_1522624, 16407, (int)getMemAt(0x587000, 187824), v7);
			}
			++v7;
		} while (v7 < *(short*)(v8 + 44));
	}
	int v10 = 0;
	int v11 = *(uint32_t*)((uint32_t)dword_5d4594_1522628 + 32);
	if (*(uint16_t*)(v11 + 44) > 0) {
		do {
			const wchar2_t* v12 = (const wchar2_t*)nox_window_call_field_94(*(int*)&dword_5d4594_1522628, 16406, v10, 0);
			if (!nox_wcscmp(v12, v6)) {
				nox_window_call_field_94(*(int*)&dword_5d4594_1522628, 16407, (int)getMemAt(0x587000, 187828), v10);
			}
			++v10;
		} while (v10 < *(short*)(v11 + 44));
	}
	nox_window_call_field_94(*(int*)&dword_5d4594_1522632, 16407, (int)v6,
							 *(uint32_t*)(*(uint32_t*)(dword_5d4594_1522632 + 32) + 48));
	nox_window_call_field_94(*(int*)&dword_5d4594_1522632, 16403, -1, 0);
	dword_5d4594_1522632 = 0;
	return 1;
}

//----- (004CC3C0) --------------------------------------------------------
int sub_4CC3C0(unsigned int a1) {
	char* v1;          // edi
	int v2;            // esi
	int v3;            // ebx
	const wchar2_t* v4; // eax
	int v5;            // esi
	int v6;            // ebx
	const wchar2_t* v7; // eax

	v1 = nox_xxx_keybind_titleByKey_42EA00(a1);
	if (!dword_5d4594_1522632) {
		return 1;
	}
	v2 = 0;
	v3 = *(uint32_t*)((uint32_t)dword_5d4594_1522624 + 32);
	if (*(uint16_t*)(v3 + 44) > 0) {
		do {
			v4 = (const wchar2_t*)nox_window_call_field_94(*(int*)&dword_5d4594_1522624, 16406, v2, 0);
			if (!nox_wcscmp(v4, (const wchar2_t*)v1)) {
				nox_window_call_field_94(*(int*)&dword_5d4594_1522624, 16407, (int)getMemAt(0x587000, 187832), v2);
			}
			++v2;
		} while (v2 < *(short*)(v3 + 44));
	}
	v5 = 0;
	v6 = *(uint32_t*)((uint32_t)dword_5d4594_1522628 + 32);
	if (*(uint16_t*)(v6 + 44) > 0) {
		do {
			v7 = (const wchar2_t*)nox_window_call_field_94(*(int*)&dword_5d4594_1522628, 16406, v5, 0);
			if (!nox_wcscmp(v7, (const wchar2_t*)v1)) {
				nox_window_call_field_94(*(int*)&dword_5d4594_1522628, 16407, (int)getMemAt(0x587000, 187836), v5);
			}
			++v5;
		} while (v5 < *(short*)(v6 + 44));
	}
	nox_window_call_field_94(*(int*)&dword_5d4594_1522632, 16407, (int)v1,
							 *(uint32_t*)(*(uint32_t*)(dword_5d4594_1522632 + 32) + 48));
	nox_window_call_field_94(*(int*)&dword_5d4594_1522632, 16403, -1, 0);
	dword_5d4594_1522632 = 0;
	return 1;
}

//----- (004CCCF0) --------------------------------------------------------
int nox_xxx_updDrawUndeadKiller_4CCCF0() { return 1; }

//----- (004CCD00) --------------------------------------------------------
int sub_4CCD00(int a1, int a2) {
	unsigned int i; // ecx

	for (i = *(uint32_t*)(a2 + 432); i < gameFrame(); ++i) {
		if (*(float*)(a2 + 436) > 0.0) {
			*(float*)(a2 + 436) = *(float*)(a2 + 440) + *(float*)(a2 + 436);
			*(float*)(a2 + 440) = *(float*)(a2 + 440) - 1.0;
		}
		if (*(float*)(a2 + 436) <= 0.0) {
			*(uint32_t*)(a2 + 436) = 0;
			*(uint32_t*)(a2 + 440) = 0;
		}
	}
	*(uint16_t*)(a2 + 104) = (long long)*(float*)(a2 + 436);
	*(uint8_t*)(a2 + 296) = (long long)*(float*)(a2 + 440);
	*(uint32_t*)(a2 + 432) = gameFrame();
	return 1;
}

//----- (004CCDB0) --------------------------------------------------------
int nox_xxx_updDrawFist_4CCDB0(int a1, int a2) {
	unsigned int i; // edx
	double v3;      // st7
	double v4;      // st7
	double v5;      // st7

	for (i = *(uint32_t*)(a2 + 432); i < gameFrame(); ++i) {
		v3 = *(float*)(a2 + 436) + *(float*)(a2 + 440);
		*(float*)(a2 + 436) = v3;
		if (v3 >= 0.0) {
			*(float*)(a2 + 440) = *(float*)(a2 + 440) - 0.5;
		} else {
			v4 = -*(float*)(a2 + 440) * *(float*)(a2 + 444);
			*(uint32_t*)(a2 + 436) = 0;
			v5 = v4 * 0.1;
			*(float*)(a2 + 440) = v5;
			if (v5 < 2.0) {
				*(uint32_t*)(a2 + 436) = 0;
				*(uint32_t*)(a2 + 440) = 0;
			}
		}
	}
	*(uint16_t*)(a2 + 104) = (long long)*(float*)(a2 + 436);
	*(uint8_t*)(a2 + 296) = (long long)*(float*)(a2 + 440);
	*(uint32_t*)(a2 + 432) = gameFrame();
	return 1;
}

//----- (004CCE70) --------------------------------------------------------
int sub_4CCE70(int a1, uint32_t* a2) {
	if (!a2[120] && !nox_xxx_checkGameFlagPause_413A50()) {
		sub_4CCEA0(a2, 5);
	}
	return 1;
}

//----- (004CD090) --------------------------------------------------------
int sub_4CD090(int a1, uint32_t* a2) {
	if (!a2[120] && !nox_xxx_checkGameFlagPause_413A50()) {
		sub_4CCEA0(a2, 4);
	}
	return 1;
}

//----- (004CD0C0) --------------------------------------------------------
int sub_4CD0C0(int a1, uint32_t* a2) {
	if (!a2[120] && !nox_xxx_checkGameFlagPause_413A50()) {
		sub_4CCEA0(a2, 3);
	}
	return 1;
}

//----- (004CD0F0) --------------------------------------------------------
int sub_4CD0F0(int a1, uint32_t* a2) {
	if (!a2[120] && !nox_xxx_checkGameFlagPause_413A50()) {
		sub_4CCEA0(a2, 2);
	}
	return 1;
}

//----- (004CD120) --------------------------------------------------------
int sub_4CD120(int a1, uint32_t* a2) {
	if (!a2[120] && !nox_xxx_checkGameFlagPause_413A50()) {
		sub_4CCEA0(a2, 1);
	}
	return 1;
}

// 4CD249: variable 'v11' is possibly undefined
// 4CD25A: variable 'v13' is possibly undefined
// 4CD27E: variable 'v15' is possibly undefined
// 4CD28F: variable 'v16' is possibly undefined

//----- (004CD400) --------------------------------------------------------
int sub_4CD400(uint32_t* a1, int a2) {
	int v2; // eax

	v2 = dword_5d4594_1522968;
	if (!dword_5d4594_1522968) {
		v2 = nox_xxx_getTTByNameSpriteMB_44CFC0("CharmOrb");
		dword_5d4594_1522968 = v2;
	}
	sub_4CD150(v2, a1, a2, 1);
	sub_4CD150(*(int*)&dword_5d4594_1522968, a1, a2, 0);
	return 1;
}

// 4CD516: variable 'v5' is possibly undefined
// 4CD527: variable 'v7' is possibly undefined
// 4CD54B: variable 'v9' is possibly undefined
// 4CD55C: variable 'v11' is possibly undefined
