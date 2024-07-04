#include <math.h>

#include "compat.h"

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
#include "GAME3_3.h"
#include "GAME4_1.h"
#include "GAME5_2.h"
#include "client__draw__debugdraw.h"
#include "client__draw__staticdraw.h"
#include "client__drawable__drawable.h"

#include "client__gui__gadgets__listbox.h"
#include "client__gui__guibook.h"
#include "client__gui__guicon.h"
#include "client__gui__guishop.h"
#include "client__gui__servopts__guiserv.h"
#include "client__gui__tooltip.h"
#include "client__gui__window.h"

#include "client__video__draw_common.h"

#include "common/fs/nox_fs.h"
#include "common__binfile.h"
#include "common__magic__speltree.h"
#include "common__net_list.h"
#include "common__strman.h"
#include "input.h"
#include "input_common.h"
#include "operators.h"

extern uint32_t dword_8531A0_2576;
extern uint32_t dword_5d4594_1098456;
extern uint32_t dword_5d4594_1193352;
extern uint32_t dword_5d4594_1096636;
extern uint32_t dword_5d4594_3807116;
extern uint32_t dword_5d4594_1098620;
extern uint32_t dword_5d4594_1123520;
extern uint32_t dword_5d4594_1193188;
extern uint32_t dword_5d4594_1098596;
extern uint32_t dword_5d4594_3799452;
extern uint32_t dword_5d4594_1098600;
extern uint32_t dword_5d4594_3807152;
extern uint32_t dword_5d4594_1098616;
extern uint32_t dword_5d4594_3807136;
extern uint32_t dword_5d4594_1098604;
extern uint32_t dword_5d4594_3804684;
extern uint32_t dword_5d4594_3807140;
extern uint32_t nox_xxx_waypointCounterMB_587000_154948;
extern uint32_t dword_5d4594_1098592;
extern uint32_t dword_5d4594_1098580;
extern uint32_t dword_5d4594_3798812;
extern uint32_t dword_5d4594_3798800;
extern uint32_t dword_5d4594_3798828;
extern uint64_t qword_581450_9552;
extern uint64_t qword_581450_9544;
extern uint32_t dword_5d4594_1098624;
extern uint32_t dword_5d4594_3798816;
extern uint32_t dword_5d4594_3798808;
extern uint32_t dword_5d4594_3798832;
extern uint32_t dword_5d4594_1193384;
extern uint32_t dword_5d4594_1193348;
extern uint32_t dword_5d4594_1193360;
extern uint32_t nox_client_highResFrontWalls_80820;
extern uint32_t dword_5d4594_3798836;
extern uint32_t dword_5d4594_1107036;
extern uint32_t dword_5d4594_251572;
extern uint32_t dword_5d4594_1098628;
extern uint32_t dword_5d4594_3798804;
extern uint32_t dword_5d4594_1098576;
extern uint32_t dword_5d4594_3798820;
extern uint32_t dword_5d4594_3798824;
extern void* dword_587000_155144;
extern uint32_t dword_5d4594_3798840;
extern void* dword_5d4594_1123524;
extern uint32_t dword_5d4594_1193380;
extern int nox_win_width;
extern int nox_win_height;

extern uint32_t nox_color_white_2523948;
extern uint32_t nox_color_yellow_2589772;

extern nox_render_data_t* nox_draw_curDrawData_3799572;

extern obj_5D4594_2650668_t** ptr_5D4594_2650668;

extern uint32_t nox_tile_def_cnt;
extern nox_tileDef_t nox_tile_defs_arr[176];

uint8_t** nox_pixbuffer_rows_3798784 = 0;

void* dword_5d4594_1096640 = 0;

void (*func_587000_154940)(int2*, uint32_t, uint32_t) = nox_xxx_tileDraw_4815E0;
int (*func_587000_154944)(int, int) = nox_xxx_drawTexEdgesProbably_481900;

void* nox_client_spriteUnderCursorXxx_1096644 = 0;
uint32_t nox_client_highResFloors_154952 = 1;
void* nox_video_tileBuf_ptr_3798796 = 0;
void* nox_video_tileBuf_end_3798844 = 0;

//----- (00476080) --------------------------------------------------------
int sub_476080(unsigned char* a1) {
	int v1;     // esi
	int v2;     // ecx
	int result; // eax
	int v4;     // edx
	int v5;     // ecx

	if (!*getMemU32Ptr(0x852978, 8)) {
		return 23 * a1[6] + 11;
	}
	switch (*a1) {
	case 0u:
	case 3u:
	case 0xBu:
		v1 = -23;
		v2 = 23 * a1[5] + 22;
		result = 23 * a1[6];
		break;
	case 1u:
	case 4u:
	case 0xCu:
		v1 = 23;
		v2 = 23 * a1[5];
		result = 23 * a1[6];
		break;
	default:
		return 23 * a1[6] + 11;
	}
	v4 = *(uint32_t*)(*getMemU32Ptr(0x852978, 8) + 12) - v2;
	v5 = v1 * (*(uint32_t*)(*getMemU32Ptr(0x852978, 8) + 16) - result) - 23 * v4;
	if (v1 < 0) {
		v5 = 23 * v4 - v1 * (*(uint32_t*)(*getMemU32Ptr(0x852978, 8) + 16) - result);
	}
	if (v5 < 0) {
		result += 22;
	}
	return result;
}

//----- (004761B0) --------------------------------------------------------
int sub_4761B0(nox_drawable* a1p) {
	int a1 = a1p;
	int result; // eax
	int v2;     // edx
	int v3;     // ecx
	int v4;     // edx

	if (!*getMemU32Ptr(0x852978, 8)) {
		return *(uint32_t*)(a1 + 16) + *getMemIntPtr(0x587000, 196188 + 8 * *(unsigned char*)(a1 + 299)) / 2;
	}
	result = *(uint32_t*)(a1 + 16);
	v2 = 8 * *(unsigned char*)(a1 + 299);
	v3 = (*(uint32_t*)(*getMemU32Ptr(0x852978, 8) + 16) - result) * *getMemIntPtr(0x587000, 196184 + v2) -
		 (*(uint32_t*)(*getMemU32Ptr(0x852978, 8) + 12) - *(uint32_t*)(a1 + 12)) * *getMemIntPtr(0x587000, 196188 + v2);
	if (*getMemIntPtr(0x587000, 196184 + v2) < 0) {
		v3 = (*(uint32_t*)(*getMemU32Ptr(0x852978, 8) + 12) - *(uint32_t*)(a1 + 12)) *
				 *getMemIntPtr(0x587000, 196188 + 8 * *(unsigned char*)(a1 + 299)) -
			 (*(uint32_t*)(*getMemU32Ptr(0x852978, 8) + 16) - result) *
				 *getMemIntPtr(0x587000, 196184 + 8 * *(unsigned char*)(a1 + 299));
	}
	v4 = result + *getMemIntPtr(0x587000, 196188 + 8 * *(unsigned char*)(a1 + 299));
	if (v3 >= 0) {
		if (v4 <= result) {
			result += *getMemIntPtr(0x587000, 196188 + 8 * *(unsigned char*)(a1 + 299));
		}
	} else if (v4 > result) {
		result += *getMemIntPtr(0x587000, 196188 + 8 * *(unsigned char*)(a1 + 299));
	}
	return result;
}

//----- (00476AE0) --------------------------------------------------------
void* sub_476AE0(nox_draw_viewport_t* vp, nox_drawable* dr) {
	unsigned char* a2 = dr;
	unsigned char* v2;                        // ebx
	int result;                               // eax
	int (*result2)(int*, int);                // eax
	int v4;                                   // eax
	int v5;                                   // edi
	int v6;                                   // ecx
	int v7;                                   // ebp
	char* v8;                                 // esi
	int v9;                                   // edx
	int v10;                                  // ecx
	unsigned int v11;                         // edi
	uint8_t* v12;                             // esi
	int v13;                                  // ebx
	int v14;                                  // edx
	unsigned int v15;                         // ecx
	unsigned int v16;                         // ebx
	int v17;                                  // ecx
	int v18;                                  // eax
	int v19;                                  // eax
	int v20;                                  // ebp
	int v21;                                  // edi
	int v22;                                  // ebp
	int v23;                                  // ebx
	int v24;                                  // esi
	int v25;                                  // [esp+10h] [ebp-1Ch]
	int v26;                                  // [esp+14h] [ebp-18h]
	int v27;                                  // [esp+18h] [ebp-14h]
	int v28;                                  // [esp+18h] [ebp-14h]
	unsigned int v29;                         // [esp+1Ch] [ebp-10h]
	int v30;                                  // [esp+20h] [ebp-Ch]
	int v31;                                  // [esp+24h] [ebp-8h]
	int v32;                                  // [esp+28h] [ebp-4h]
	void (*v33)(unsigned int, uint8_t*, int); // [esp+34h] [ebp+8h]

	v2 = a2;
	result2 = (int (*)(int*, int)) * ((uint32_t*)a2 + 75);
	if (result2 == nox_thing_static_draw) {
		if (*((uint32_t*)a2 + 28) & 0x40000 && !(*((uint32_t*)a2 + 30) & 0x1000000)) {
			return result2;
		}
		v4 = *(uint32_t*)(*((uint32_t*)a2 + 76) + 4);
	} else {
		v4 = *(uint32_t*)(*(uint32_t*)(*((uint32_t*)a2 + 76) + 4) + 4 * *((uint32_t*)a2 + 77));
	}
	result = nox_video_getImagePixdata_42FB30(v4);
	if (result) {
		v33 = sub_476D70;
		v5 = *(uint32_t*)result;
		v6 = *((uint32_t*)result + 1);
		v27 = *((uint32_t*)result + 1);
		v7 = *((uint32_t*)result + 2) + *((uint32_t*)v2 + 3) - *v2;
		v8 = (char*)result + 16;
		result =
			(int)(*((uint32_t*)result + 3) + *((uint32_t*)v2 + 4) - *((short*)v2 + 53) - *((short*)v2 + 52) - v2[1]);
		v31 = v5;
		if (v7 < *(int*)&dword_5d4594_3798820 || v7 + v5 >= *(int*)&dword_5d4594_3798820 + dword_5d4594_3798800 ||
			(v9 = dword_5d4594_3798824, (int)result < *(int*)&dword_5d4594_3798824) ||
			(int)result + v6 >= *(int*)&dword_5d4594_3798824 + dword_5d4594_3798808) {
			*((uint32_t*)v2 + 86) = 0;
		} else {
			v10 = nox_xxx_waypointCounterMB_587000_154948;
			if (*(int*)&nox_xxx_waypointCounterMB_587000_154948 <= 0) {
				*((uint32_t*)v2 + 86) = 0;
				v9 = dword_5d4594_3798824;
				v10 = nox_xxx_waypointCounterMB_587000_154948;
			}
			if (v10 - *((int*)v2 + 86) > 1 || v10 <= 0) {
				v11 = nox_video_tileBuf_end_3798844;
				v12 = v8 + 1;
				v29 = nox_video_tileBuf_end_3798844;
				v26 = (uint32_t)nox_video_tileBuf_end_3798844 - (uint32_t)nox_video_tileBuf_ptr_3798796;
				v13 = dword_5d4594_3798804 * ((uint32_t)result + dword_5d4594_3798840 - v9);
				v14 = v7 + dword_5d4594_3798836 - dword_5d4594_3798820;
				v15 = v13 + (uint32_t)nox_video_tileBuf_ptr_3798796 + 2 * v14;
				v25 = v13 + (uint32_t)nox_video_tileBuf_ptr_3798796 + 2 * v14;
				if (v15 >= nox_video_tileBuf_end_3798844) {
					v15 -= v26;
					v25 = v15;
				}
				result = (int)(v27 - 1);
				if (v27) {
					v30 = v27;
					do {
						v16 = v15;
						v28 = v31;
						if (v31 > 0) {
							do {
								v17 = (unsigned char)v12[1];
								v18 = *v12 & 0xF;
								v12 += 2;
								v19 = v18 - 1;
								v32 = v17;
								v20 = 2 * v17;
								if (v19) {
									if (v19 == 1) {
										if (v16 >= v11 || v16 + v20 < v11) {
											v33(v16, v12, 2 * v17);
											v12 += v20;
											v16 += v20;
										} else {
											v21 = v16 + v20 - v29;
											v22 = v29 - v16;
											v33(v16, v12, v29 - v16);
											v23 = nox_video_tileBuf_ptr_3798796;
											v24 = (int)&v12[v22];
											v33(nox_video_tileBuf_ptr_3798796, (uint8_t*)v24, v21);
											v12 = (uint8_t*)(v21 + v24);
											v16 = v21 + v23;
											v11 = v29;
										}
									}
								} else {
									v16 += v20;
									if (v16 >= v11) {
										v16 -= v26;
									}
								}
								v28 -= v32;
							} while (v28 > 0);
							v15 = v25;
						}
						v15 += dword_5d4594_3798804;
						v25 = v15;
						if (v15 >= v11) {
							v15 -= v26;
							v25 = v15;
						}
						result = (int)--v30;
					} while (v30);
				}
			} else {
				*((uint32_t*)v2 + 86) = v10;
			}
		}
	}
	return result;
}

//----- (00476D70) --------------------------------------------------------
short sub_476D70(uint32_t* a1, int* a2, unsigned int a3) {
	uint32_t* v3;  // edi
	signed int v4; // ecx
	int* v5;       // esi
	int v6;        // eax

	v3 = a1;
	v4 = a3 >> 2;
	v5 = a2;
	if (a3 >> 2) {
		do {
			v6 = *v5;
			++v5;
			*v3 = v6;
			++v3;
		} while (v4-- > 1);
	}
	if (a3 & 3) {
		LOWORD(v6) = *(uint16_t*)v5;
		*(uint16_t*)v3 = *(uint16_t*)v5;
	}
	return v6;
}

//----- (00476E00) --------------------------------------------------------
int nox_client_setPhonemeFrame_476E00(int a1) {
	int result; // eax

	result = a1;
	*getMemU32Ptr(0x5D4594, 1096596 + 4 * a1) = gameFrame();
	return result;
}

//----- (00476E20) --------------------------------------------------------
uint32_t* sub_476E20() {
	int v0;       // esi
	char* v1;     // eax
	uint32_t* v2; // esi

	v0 = 0;
	while (1) {
		v1 = nox_xxx_gLoadImg_42F970(*(const char**)getMemAt(0x587000, 151272 + v0));
		*getMemU32Ptr(0x5D4594, 1096564 + v0) = v1;
		if (!v1) {
			break;
		}
		v0 += 4;
		if (v0 >= 32) {
			v2 = nox_window_new(0, 64, (nox_win_width - 100) / 2, (nox_win_height - 100) / 2, 1, 1, 0);
			nox_window_set_all_funcs(v2, 0, sub_476E90, 0);
			return v2;
		}
	}
	return 0;
}

//----- (00476E90) --------------------------------------------------------
int sub_476E90() {
	unsigned char* v0; // edi
	int v1;            // esi

	v0 = getMemAt(0x587000, 151208);
	v1 = 0;
	do {
		if (*getMemU32Ptr(0x5D4594, 1096596 + v1)) {
			nox_client_drawImageAt_47D2C0(*getMemU32Ptr(0x5D4594, 1096564 + v1),
										  nox_win_width / 2 + *(uint32_t*)v0 - 16,
										  *((uint32_t*)v0 + 1) + nox_win_height / 2 - 41);
			if ((unsigned int)(gameFrame() - *getMemU32Ptr(0x5D4594, 1096596 + v1)) > 3) {
				*getMemU32Ptr(0x5D4594, 1096596 + v1) = 0;
			}
		}
		v0 += 8;
		v1 += 4;
	} while ((int)v0 < (int)getMemAt(0x587000, 151272));
	return 1;
}

//----- (00476F40) --------------------------------------------------------
unsigned int nox_xxx_packetGetMarshall_476F40() {
	unsigned int result; // eax

	if (dword_5d4594_1096640) {
		result = nox_xxx_netGetUnitCodeCli_578B00(*(int*)&dword_5d4594_1096640);
	} else {
		result = 0;
	}
	return result;
}

//----- (00476FA0) --------------------------------------------------------
void nox_xxx_clientEnumHover_476FA0() {
	int4 v2; // [esp+10h] [ebp-10h]

	if (!*getMemU32Ptr(0x5D4594, 1096632)) {
		*getMemU32Ptr(0x5D4594, 1096632) = nox_xxx_getNameId_4E3AA0("Glyph");
	}
	nox_point mpos = nox_client_getMousePos_4309F0();
	sub_473970(&mpos, &mpos);
	dword_5d4594_1096640 = 0;
	nox_client_spriteUnderCursorXxx_1096644 = 0;
	*getMemU32Ptr(0x5D4594, 1096628) = 0;
	v2.field_0 = mpos.x - 96;
	v2.field_8 = mpos.x + 96;
	v2.field_C = mpos.y + 96;
	v2.field_4 = mpos.y - 96;
	dword_5d4594_1096636 = 0;
	nox_xxx_forEachSprite_49AB00(&v2, nox_xxx_clientOnCursorHover_477050, &mpos);
}

//----- (00477050) --------------------------------------------------------
int nox_xxx_client_4984B0_drawable(nox_drawable* dr);
void nox_xxx_clientOnCursorHover_477050(int arg0, int a2) {
	int v2;    // esi
	int v3;    // eax
	int v4;    // eax
	int v5;    // ecx
	char* v6;  // eax
	int v7;    // eax
	float v8;  // ebp
	int v9;    // edi
	int v10;   // ebx
	int v11;   // eax
	int v12;   // edx
	int v13;   // edi
	int v14;   // ebx
	float v15; // eax
	int v16;   // edx
	int v17;   // eax
	int v18;   // edi
	int v19;   // ebx
	int v20;   // eax
	int v21;   // ecx
	int v22;   // edi
	float v23; // [esp+0h] [ebp-24h]
	float v24; // [esp+0h] [ebp-24h]
	float v25; // [esp+0h] [ebp-24h]
	float v26; // [esp+0h] [ebp-24h]
	float2 a3; // [esp+14h] [ebp-10h]
	float2 a1; // [esp+1Ch] [ebp-8h]
	int v29;   // [esp+28h] [ebp+4h]

	if (!*getMemU32Ptr(0x5D4594, 1096648)) {
		*getMemU32Ptr(0x5D4594, 1096648) = nox_xxx_getTTByNameSpriteMB_44CFC0("Polyp");
	}
	v2 = arg0;
	if (arg0 == *getMemU32Ptr(0x852978, 8)) {
		return;
	}
	v3 = *(uint32_t*)(arg0 + 120);
	if (!((v3 & 0x8000) == 0 && (!nox_client_drawable_testBuff_4356C0(arg0, 0) ||
								 nox_client_drawable_testBuff_4356C0(*getMemIntPtr(0x852978, 8), 21)))) {
		return;
	}
	v4 = *(uint32_t*)(arg0 + 112);
	if (!(!(v4 & 2) || (v5 = *(uint32_t*)(arg0 + 116), !(v5 & 0x4000)))) {
		return;
	}
	if (!(v4 & 0x80400206 || *(uint32_t*)(arg0 + 108) == *getMemU32Ptr(0x5D4594, 1096648))) {
		return;
	}
	if (!nox_xxx_client_4984B0_drawable(arg0)) {
		return;
	}
	if (!(!(*(uint8_t*)(arg0 + 112) & 4) ||
		  (v6 = nox_common_playerInfoGetByID_417040(*(uint32_t*)(arg0 + 128))) != 0 && !(v6[3680] & 1))) {
		return;
	}
	v7 = *(uint32_t*)(arg0 + 112);
	if (!((!(v7 & 0x400000) || (*(uint8_t*)(arg0 + 116) & 0x80)) && (!(v7 & 2) || *(uint32_t*)(arg0 + 276) != 10))) {
		return;
	}
	v23 = (double)*(int*)(arg0 + 16) - *(float*)(arg0 + 100) - (double)*(short*)(arg0 + 104);
	v29 = nox_float2int(v23);
	v24 = (double)*(int*)(v2 + 16) - *(float*)(v2 + 96) - (double)*(short*)(v2 + 104);
	v8 = COERCE_FLOAT(nox_float2int(v24));
	a3.field_0 = v8;
	if (*(uint32_t*)(v2 + 44) == 2) {
		v25 = *(float*)(v2 + 48) * *(float*)(v2 + 48);
		LODWORD(a3.field_0) = nox_float2int(v25);
		v17 = nox_float2int(*(float*)(v2 + 48));
		v18 = *(uint32_t*)(v2 + 12);
		v19 = *(uint32_t*)(v2 + 12) - v17;
		v20 = v18 + nox_float2int(*(float*)(v2 + 48));
		v21 = *(uint32_t*)(a2 + 4);
		if (v21 <= SLODWORD(v8)) {
			v8 = *(float*)&v29;
			if (v21 >= v29) {
				if (*(int*)a2 <= v19 || *(int*)a2 >= v20) {
					return;
				}
				goto LABEL_38;
			}
		}
		v15 = a3.field_0;
		v16 = (*(uint32_t*)a2 - v18) * (*(uint32_t*)a2 - v18) + (v21 - LODWORD(v8)) * (v21 - LODWORD(v8));
	} else {
		if (*(uint32_t*)(v2 + 44) != 3) {
			return;
		}
		a1.field_0 = (double)*(int*)(v2 + 12);
		a1.field_4 = (double)SLODWORD(a3.field_0);
		a3.field_0 = (double)*(int*)a2;
		a3.field_4 = (double)*(int*)(a2 + 4);
		if (nox_xxx_map_57B850(&a1, (float*)(v2 + 44), &a3) ||
			(a1.field_4 = (double)v29, nox_xxx_map_57B850(&a1, (float*)(v2 + 44), &a3)) ||
			(v9 = *(uint32_t*)(v2 + 12) + nox_float2int(*(float*)(v2 + 72)),
			 v10 = v29 + nox_float2int(*(float*)(v2 + 76)), v11 = LODWORD(v8) + nox_float2int(*(float*)(v2 + 76)),
			 *(uint32_t*)a2 > v9) &&
				*(uint32_t*)a2 < *(int*)(v2 + 12) && (v12 = *(uint32_t*)(a2 + 4), v12 > v10) && v12 < v11) {
			goto LABEL_38;
		}
		v13 = *(uint32_t*)(v2 + 12) + nox_float2int(*(float*)(v2 + 80));
		v14 = v29 + nox_float2int(*(float*)(v2 + 84));
		LODWORD(v15) = LODWORD(v8) + nox_float2int(*(float*)(v2 + 84));
		if (*(int*)a2 < *(int*)(v2 + 12)) {
			return;
		}
		if (*(int*)a2 >= v13) {
			return;
		}
		v16 = *(uint32_t*)(a2 + 4);
		if (v16 <= v14) {
			return;
		}
	}
	if (v16 >= SLODWORD(v15)) {
		return;
	}
LABEL_38:
	v26 = (double)*(short*)(v2 + 104) + (double)*(int*)(v2 + 16) + *(float*)(v2 + 96);
	v22 = nox_float2int(v26);
	if (v22 > *getMemIntPtr(0x5D4594, 1096628)) {
		*getMemU32Ptr(0x5D4594, 1096628) = v22;
		dword_5d4594_1096640 = v2;
	}
	if (v2 != *getMemU32Ptr(0x852978, 8) && v22 > *(int*)&dword_5d4594_1096636 && nox_xxx_client_57B400(v2)) {
		if (dword_8531A0_2576 && *(uint8_t*)(dword_8531A0_2576 + 2251) == 1 &&
			*(uint32_t*)(v2 + 108) == *getMemU32Ptr(0x5D4594, 1096632)) {
			if (!nox_client_spriteUnderCursorXxx_1096644) {
				nox_client_spriteUnderCursorXxx_1096644 = v2;
				dword_5d4594_1096636 = 0;
			}
		} else {
			dword_5d4594_1096636 = v22;
			nox_client_spriteUnderCursorXxx_1096644 = v2;
		}
	}
}

//----- (00477600) --------------------------------------------------------
int nox_xxx_guiCursor_477600() { return *getMemU32Ptr(0x5D4594, 1096672); }

//----- (004776B0) --------------------------------------------------------
void nox_xxx_cursorSetTooltip_4776B0(wchar2_t* a1) {
	if (a1) {
		if ((int)nox_wcslen(a1) >= 256) {
			nox_wcsncpy((wchar2_t*)getMemAt(0x5D4594, 1096676), a1, 0xFFu);
			*getMemU16Ptr(0x5D4594, 1097186) = 0;
		} else {
			nox_wcscpy((wchar2_t*)getMemAt(0x5D4594, 1096676), a1);
		}
	} else {
		*getMemU16Ptr(0x5D4594, 1096676) = 0;
	}
}

//----- (00478030) --------------------------------------------------------
int sub_478030() { return dword_5d4594_1098624; }

//----- (00478040) --------------------------------------------------------
int sub_478040() {
	short v2; // [esp+0h] [ebp-2h]

	if (!dword_5d4594_1098624) {
		return 0;
	}
	v2 = 4809;
	nox_netlist_addToMsgListCli_40EBC0(31, 0, &v2, 2);
	sub_467680();
	return 1;
}

//----- (00478080) --------------------------------------------------------
int sub_478080(int a1) {
	char* v1;   // eax
	int result; // eax

	if (dword_5d4594_1098624 && (v1 = sub_4780A0(a1)) != 0) {
		result = *(uint32_t*)v1;
	} else {
		result = 0;
	}
	return result;
}

//----- (004780A0) --------------------------------------------------------
char* sub_4780A0(int a1) {
	int v1;            // ebp
	unsigned char* v2; // edi
	int v3;            // esi
	unsigned char* v4; // edx
	int v5;            // eax
	unsigned char* v6; // ecx

	v1 = 0;
	v2 = getMemAt(0x5D4594, 1098644);
	while (1) {
		v3 = 0;
		v4 = v2;
		do {
			if (*((uint32_t*)v4 - 1)) {
				v5 = 0;
				v6 = v4;
				while (*(uint32_t*)v6 != a1) {
					++v5;
					v6 += 4;
					if (v5 >= 32) {
						goto LABEL_7;
					}
				}
				return (char*)getMemAt(0x5D4594, 1098636 + 140 * (v1 + 10 * v3));
			}
		LABEL_7:
			++v3;
			v4 += 1400;
		} while (v3 < 6);
		v2 += 140;
		++v1;
		if ((int)v2 < (int)getMemAt(0x5D4594, 1100044)) {
			continue;
		}
		return 0;
	}
}

//----- (00478110) --------------------------------------------------------
int sub_478110() {
	uint32_t* v0;      // eax
	int result;        // eax
	uint32_t* v2;      // eax
	unsigned char* v3; // edx
	unsigned char* v4; // eax
	int v5;            // ecx
	uint32_t* v6;      // esi
	int v7;            // [esp+4h] [ebp-8h]
	int v8;            // [esp+8h] [ebp-4h]

	*getMemU32Ptr(0x5D4594, 1098500) = nox_win_width;
	*getMemU32Ptr(0x5D4594, 1098504) = nox_win_height;
	*getMemU32Ptr(0x5D4594, 1098524) = nox_win_width;
	*getMemU32Ptr(0x5D4594, 1098492) = 0;
	*getMemU32Ptr(0x5D4594, 1098496) = 0;
	*getMemU32Ptr(0x5D4594, 1098528) = nox_win_height;
	*getMemU32Ptr(0x5D4594, 1098508) = 0;
	*getMemU32Ptr(0x5D4594, 1098512) = 0;
	v0 = nox_new_window_from_file("Shop.wnd", sub_478480);
	dword_5d4594_1098576 = v0;
	if (!v0) {
		return 0;
	}
	nox_window_set_all_funcs(v0, sub_478650, sub_478970, 0);
	v2 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1098576, 3806);
	nox_gui_winSetFunc96_46B070((int)v2, sub_478E50);
	v3 = getMemAt(0x5D4594, 1098636);
	do {
		v4 = v3;
		v5 = 6;
		do {
			*(uint32_t*)v4 = 0;
			v4 += 1400;
			--v5;
		} while (v5);
		v3 += 140;
	} while ((int)v3 < (int)getMemAt(0x5D4594, 1100036));
	sub_478F10();
	nox_window_get_size(*(int*)&dword_5d4594_1098576, &v8, &v7);
	nox_window_setPos_46A9B0(*(uint32_t**)&dword_5d4594_1098576, nox_win_width - v8, nox_win_height - v7);
	nox_window_set_hidden(*(int*)&dword_5d4594_1098576, 1);
	nox_xxx_wnd_46ABB0(*(int*)&dword_5d4594_1098576, 0);
	*getMemU32Ptr(0x5D4594, 1098400) = nox_xxx_gLoadImg_42F970("ShopBase");
	*getMemU32Ptr(0x5D4594, 1098404) = nox_xxx_gLoadImg_42F970("ShopTradeMode");
	*getMemU32Ptr(0x5D4594, 1098408) = nox_xxx_gLoadImg_42F970("ShopIdentifyMode");
	*getMemU32Ptr(0x5D4594, 1098412) = nox_xxx_gLoadImg_42F970("ShopRepairMode");
	*getMemU32Ptr(0x5D4594, 1098416) = nox_xxx_gLoadImg_42F970("ShopRepairMode");
	*getMemU32Ptr(0x5D4594, 1098420) = nox_xxx_gLoadImg_42F970("ShopExitMode");
	*getMemU32Ptr(0x5D4594, 1098424) = nox_xxx_gLoadImg_42F970("ShopInventoryBar1");
	*getMemU32Ptr(0x5D4594, 1098428) = nox_xxx_gLoadImg_42F970("ShopInventoryBar2");
	*getMemU32Ptr(0x5D4594, 1098432) = nox_xxx_gLoadImg_42F970("ShopInventorySlider");
	*getMemU32Ptr(0x5D4594, 1098436) = nox_xxx_gLoadImg_42F970("ShopInventorySliderSelected");
	*getMemU32Ptr(0x5D4594, 1098448) = nox_xxx_gLoadImg_42F970("ShopInventoryUp");
	*getMemU32Ptr(0x5D4594, 1098452) = nox_xxx_gLoadImg_42F970("ShopInventoryUpSelected");
	*getMemU32Ptr(0x5D4594, 1098440) = nox_xxx_gLoadImg_42F970("ShopInventorydown");
	*getMemU32Ptr(0x5D4594, 1098444) = nox_xxx_gLoadImg_42F970("ShopInventorydownSelected");
	dword_5d4594_1098456 = nox_xxx_gLoadImg_42F970("ShopTextBorder");
	*getMemU32Ptr(0x5D4594, 1098460) = nox_xxx_gLoadImg_42F970("ShopkeeperPic");
	*getMemU32Ptr(0x5D4594, 1098464) = nox_xxx_gLoadImg_42F970("ShopkeeperWarriorPic");
	*getMemU32Ptr(0x5D4594, 1098468) = nox_xxx_gLoadImg_42F970("ShopkeeperConjurerPic");
	*getMemU32Ptr(0x5D4594, 1098472) = nox_xxx_gLoadImg_42F970("ShopkeeperWizardPic");
	*getMemU32Ptr(0x5D4594, 1098476) = nox_xxx_gLoadImg_42F970("ShopkeeperLandOfTheDeadPic");
	*getMemU32Ptr(0x5D4594, 1098480) = nox_xxx_gLoadImg_42F970("ShopkeeperMagicShopPic");
	dword_5d4594_1098580 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1098576, 3807);
	*getMemU32Ptr(0x5D4594, 1098584) = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1098576, 3808);
	*getMemU32Ptr(0x5D4594, 1098588) = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1098576, 3809);
	v6 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1098576, 3806);
	nox_client_wndGetPosition_46AA60(v6, getMemAt(0x5D4594, 1098380), getMemAt(0x5D4594, 1098384));
	nox_window_get_size((int)v6, getMemAt(0x5D4594, 1098388), getMemAt(0x5D4594, 1098392));
	*getMemU32Ptr(0x5D4594, 1098388) += *getMemU32Ptr(0x5D4594, 1098380);
	*getMemU32Ptr(0x5D4594, 1098392) += *getMemU32Ptr(0x5D4594, 1098384);
	*(uint32_t*)(*(uint32_t*)(dword_5d4594_1098580 + 400) + 8) = 16;
	*(uint32_t*)(*(uint32_t*)(dword_5d4594_1098580 + 400) + 12) = 12;
	nox_xxx_wndSetOffsetMB_46AE40(*(uint32_t*)(dword_5d4594_1098580 + 400), 0, -15);
	sub_4B5700(*(int*)&dword_5d4594_1098580, 0, 0, *getMemIntPtr(0x5D4594, 1098432), *getMemIntPtr(0x5D4594, 1098436),
			   *getMemIntPtr(0x5D4594, 1098436));
	result = 1;
	dword_5d4594_1098592 = *(uint32_t*)(*(uint32_t*)(dword_5d4594_1098580 + 32) + 4);
	return result;
}

//----- (00478480) --------------------------------------------------------
int sub_478480(int a1, int a2, int* a3, int a4) {
	int result; // eax
	int v5;     // esi
	int v7;     // ecx
	int v8;     // ecx

	if (a2 != 16391) {
		if (a2 == 16393) {
			dword_5d4594_1107036 = dword_5d4594_1098592 - a4;
			return 0;
		}
		return 0;
	}
	if (sub_45D9B0()) {
		return 0;
	}
	v5 = nox_xxx_wndGetID_46B0A0(a3);
	nox_xxx_clientPlaySoundSpecial_452D80(766, 100);
	switch (v5) {
	case 3801:
		if (!dword_5d4594_1098624) {
			return 0;
		}
		sub_478040();
		result = 0;
		break;
	case 3802:
		if (!dword_5d4594_1098624) {
			return 0;
		}
		if (dword_5d4594_1098628 == 4) {
			sub_467680();
		}
		nox_client_setCursorType_477610(12);
		dword_5d4594_1098628 = 3;
		if (sub_467C80()) {
			return 0;
		}
		sub_467BB0();
		result = 0;
		break;
	case 3803:
		if (!dword_5d4594_1098624) {
			return 0;
		}
		sub_467650();
		dword_5d4594_1098628 = 4;
		result = 0;
		break;
	case 3804:
		if (dword_5d4594_1098624) {
			if (dword_5d4594_1098628 == 4) {
				sub_467680();
			}
			nox_client_setCursorType_477610(11);
			dword_5d4594_1098628 = 2;
		}
		return 0;
	case 3808:
		if (*(int*)&dword_5d4594_1107036 - 50 >= 0) {
			v8 = dword_5d4594_1107036 - 50 - (dword_5d4594_1107036 - 50) % 50;
		} else {
			v8 = 0;
		}
		dword_5d4594_1107036 = v8;
		nox_window_call_field_94(*(int*)&dword_5d4594_1098580, 16394, dword_5d4594_1098592 - v8, 0);
		result = 0;
		break;
	case 3809:
		if (*(int*)&dword_5d4594_1107036 + 50 <= *(int*)&dword_5d4594_1098592) {
			v7 = dword_5d4594_1107036 + 50 - (dword_5d4594_1107036 + 50) % 50;
			dword_5d4594_1107036 = dword_5d4594_1107036 + 50 - (dword_5d4594_1107036 + 50) % 50;
		} else {
			v7 = dword_5d4594_1098592;
			dword_5d4594_1107036 = dword_5d4594_1098592;
		}
		nox_window_call_field_94(*(int*)&dword_5d4594_1098580, 16394, dword_5d4594_1098592 - v7, 0);
		result = 0;
		break;
	default:
		return 0;
	}
	return result;
}
// 478585: variable 'v6' is possibly undefined

//----- (00478650) --------------------------------------------------------
int sub_478650(int a1, int a2, unsigned int a3) {
	int v4;  // eax
	int2 v5; // [esp+0h] [ebp-8h]

	v5.field_0 = (unsigned short)a3;
	v5.field_4 = a3 >> 16;
	if (!sub_45D9B0()) {
		switch (a2) {
		case 5:
			v4 = nox_xxx_pointInRect_4281F0(&v5, (int4*)getMemAt(0x5D4594, 1098380));
			if (v4) {
				if (dword_5d4594_1098628 == 2) {
					sub_478730(&v5.field_0);
				}
			}
			break;
		case 19:
			if (dword_5d4594_1098628 == 2) {
				nox_window_call_field_94(*(int*)&dword_5d4594_1098576, 16391, *getMemIntPtr(0x5D4594, 1098584), 0);
				return 1;
			}
			break;
		case 20:
			if (dword_5d4594_1098628 == 2) {
				nox_window_call_field_94(*(int*)&dword_5d4594_1098576, 16391, *getMemIntPtr(0x5D4594, 1098588), 0);
				return 1;
			}
			break;
		default:
			return 0;
		}
	}
	return 1;
}
// 478703: variable 'v4' is possibly undefined

//----- (00478850) --------------------------------------------------------
void sub_478850(int a1, short a2, int a3, int a4) {
	if (a4) {
		if (a4 == 1) {
			nox_client_tradeXxxBuyAccept_478880(a3, a2);
		} else {
			sub_4788F0(a3, a4);
		}
	}
}

//----- (00478970) --------------------------------------------------------
int sub_478970() {
	int result; // eax
	int2 a1;    // [esp+0h] [ebp-8h]

	a1.field_0 = nox_win_width - NOX_DEFAULT_WIDTH;
	a1.field_4 = nox_win_height - NOX_DEFAULT_HEIGHT;
	nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1098400), nox_win_width - NOX_DEFAULT_WIDTH,
								  nox_win_height - NOX_DEFAULT_HEIGHT);
	if (dword_5d4594_1098628 == 2) {
		nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1098404), a1.field_0, a1.field_4);
		sub_478C80();
		result = 1;
	} else if (dword_5d4594_1098628 == 3) {
		nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1098408), a1.field_0, a1.field_4);
		sub_478B10(&a1);
		result = 1;
	} else if (dword_5d4594_1098628 == 1) {
		nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1098412), a1.field_0, a1.field_4);
		sub_478A70(&a1);
		result = 1;
	} else {
		if (dword_5d4594_1098628 == 4) {
			nox_client_drawImageAt_47D2C0(*getMemIntPtr(0x5D4594, 1098416), a1.field_0, a1.field_4);
			sub_478BC0(&a1.field_0);
		}
		result = 1;
	}
	return result;
}

//----- (00478A70) --------------------------------------------------------
int sub_478A70(int2* a1) {
	uint32_t* v1; // esi
	int result;   // eax
	int v3;       // [esp+4h] [ebp-10h]
	int v4;       // [esp+8h] [ebp-Ch]
	int v5;       // [esp+Ch] [ebp-8h]
	int v6;       // [esp+10h] [ebp-4h]

	v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1098576, 3806);
	nox_client_wndGetPosition_46AA60(v1, &v5, &v6);
	nox_window_get_size((int)v1, &v4, &v3);
	nox_xxx_drawSetTextColor_434390(nox_color_white_2523948);
	nox_client_drawImageAt_47D2C0(dword_5d4594_1098456, a1->field_0, a1->field_4);
	result = dword_5d4594_1098604;
	if (dword_5d4594_1098604) {
		result = nox_xxx_drawStringWrap_43FAF0(0, *(uint16_t**)&dword_5d4594_1098604, v5 + 8, v6 + 8, v4 - 16, v3 - 16);
	}
	return result;
}

//----- (00478C80) --------------------------------------------------------
int sub_478C80() {
	int v0;            // edi
	int v1;            // ebx
	unsigned char* v2; // esi
	int v3;            // edi
	int v4;            // ebp
	int v6;            // [esp+10h] [ebp-58h]
	unsigned char* v7; // [esp+14h] [ebp-54h]
	unsigned int v8;   // [esp+18h] [ebp-50h]
	int v9;            // [esp+1Ch] [ebp-4Ch]
	int v10;           // [esp+20h] [ebp-48h]
	wchar2_t v11[32];   // [esp+28h] [ebp-40h]

	v8 = sub_4674A0();
	nox_xxx_wndDraw_49F7F0();
	nox_client_copyRect_49F6F0(*getMemIntPtr(0x5D4594, 1098380), *getMemIntPtr(0x5D4594, 1098384),
							   *getMemU32Ptr(0x5D4594, 1098388) - *getMemU32Ptr(0x5D4594, 1098380),
							   *getMemU32Ptr(0x5D4594, 1098392) - *getMemU32Ptr(0x5D4594, 1098384));
	v0 = *getMemU32Ptr(0x5D4594, 1098380);
	v10 = *getMemU32Ptr(0x5D4594, 1098380);
	v1 = *getMemU32Ptr(0x5D4594, 1098384) - dword_5d4594_1107036;
	v2 = getMemAt(0x5D4594, 1098640);
	v9 = nox_xxx_guiFontHeightMB_43F320(0);
	v6 = 0;
	v7 = getMemAt(0x5D4594, 1098640);
	do {
		if (v1 > *getMemIntPtr(0x5D4594, 1098384) - 50) {
			nox_client_drawImageAt_47D2C0(*getMemU32Ptr(0x5D4594, 1098424 + 4 * (v6 % 2)), v0, v1);
			v3 = v0 + 5;
			v4 = 6;
			do {
				if (*(uint32_t*)v2) {
					*(uint32_t*)(*((uint32_t*)v2 - 1) + 12) = v3 + 20;
					*(uint32_t*)(*((uint32_t*)v2 - 1) + 16) = v1 + 25;
					(*(void (**)(unsigned char*, uint32_t))(*((uint32_t*)v2 - 1) + 300))(getMemAt(0x5D4594, 1098492),
																						 *((uint32_t*)v2 - 1));
					if (v8 < *((uint32_t*)v2 + 33)) {
						nox_client_drawRectFilledAlpha_49CF10(v3 - 5, v1, 50, 50);
					}
					if (*(uint32_t*)v2 > 1u) {
						nox_swprintf(v11, L"%d", *(uint32_t*)v2);
						nox_xxx_drawSetTextColor_434390(nox_color_white_2523948);
						nox_xxx_drawStringWrap_43FAF0(0, v11, v3, v1 + 5, 320, 0);
					}
					nox_swprintf(v11, L"%d", *((uint32_t*)v2 + 33));
					nox_xxx_drawSetTextColor_434390(nox_color_yellow_2589772);
					nox_xxx_drawStringWrap_43FAF0(0, v11, v3, v1 - v9 + 45, 320, 0);
				}
				v3 += 50;
				v2 += 1400;
				--v4;
			} while (v4);
			v2 = v7;
			v0 = v10;
		}
		v1 += 50;
		if (v1 >= *getMemIntPtr(0x5D4594, 1098392)) {
			break;
		}
		v2 += 140;
		++v6;
		v7 = v2;
	} while ((int)v2 < (int)getMemAt(0x5D4594, 1100040));
	return sub_49F860();
}

//----- (00478E50) --------------------------------------------------------
int sub_478E50(int a1, int a2, unsigned int a3) {
	int v3;       // ecx
	int v4;       // edx
	int v5;       // edx
	wchar2_t** v6; // eax
	wchar2_t* v7;  // eax

	if (dword_5d4594_1098628 == 2) {
		v3 = ((unsigned short)a3 - *getMemU32Ptr(0x5D4594, 1098380)) / 50;
		v4 = (int)((a3 >> 16) - *getMemU32Ptr(0x5D4594, 1098384) + dword_5d4594_1107036) / 50;
		if (v3 >= 6) {
			v3 = 5;
		}
		if (v4 >= 10) {
			v4 = 9;
		}
		v5 = 35 * (v4 + 10 * v3);
		v6 = (wchar2_t**)getMemAt(0x5D4594, 1098636 + 4 * v5);
		if (*getMemU32Ptr(0x5D4594, 1098640 + 4 * v5)) {
			*((uint32_t*)*v6 + 32) = v6[2];
			v7 = nox_xxx_clientAskInfoMb_4BF050(*v6);
			nox_xxx_cursorSetTooltip_4776B0(v7);
		}
	}
	return 1;
}

//----- (00478F10) --------------------------------------------------------
uint32_t* sub_478F10() {
	unsigned char* v0; // esi
	int v1;            // ebx
	void* v2;          // edi
	uint32_t* result;  // eax
	unsigned char* v4; // [esp+10h] [ebp-4h]

	v4 = getMemAt(0x5D4594, 1098636);
	do {
		v0 = v4;
		v1 = 6;
		do {
			if (*(uint32_t*)v0) {
				nox_xxx_spriteDelete_45A4B0(*(uint64_t**)v0);
			}
			*(uint32_t*)v0 = 0;
			v2 = v0 + 8;
			*((uint32_t*)v0 + 1) = 0;
			v0 += 1400;
			memset(v2, 0, 0x80u);
			*((uint32_t*)v0 - 316) = 0;
			--v1;
		} while (v1);
		result = v4 + 140;
		v4 += 140;
	} while ((int)v4 < (int)getMemAt(0x5D4594, 1100036));
	dword_5d4594_1107036 = 0;
	return result;
}

//----- (00478F80) --------------------------------------------------------
int sub_478F80() {
	int result; // eax

	sub_478F10();
	sub_44D8F0();
	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1098576);
	result = 0;
	dword_5d4594_1098576 = 0;
	dword_5d4594_1098624 = 0;
	dword_5d4594_1098628 = 1;
	dword_5d4594_1098596 = 0;
	dword_5d4594_1098600 = 0;
	dword_5d4594_1098604 = 0;
	*getMemU32Ptr(0x5D4594, 1098608) = 0;
	dword_5d4594_1098616 = 0;
	return result;
}

//----- (004790F0) --------------------------------------------------------
char* nox_xxx_getShopPic_4790F0(int a1) {
	uint32_t* v1; // esi
	char* result; // eax

	v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1098576, 3805);
	if (!*getMemU32Ptr(0x5D4594, 1107040)) {
		*getMemU32Ptr(0x5D4594, 1098396) = nox_xxx_getTTByNameSpriteMB_44CFC0("Shopkeeper");
		*getMemU32Ptr(0x5D4594, 1098560) = nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperWarriorsRealm");
		*getMemU32Ptr(0x5D4594, 1098556) = nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperConjurerRealm");
		*getMemU32Ptr(0x5D4594, 1098564) = nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperWizardRealm");
		*getMemU32Ptr(0x5D4594, 1098572) = nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperLandOfTheDead");
		*getMemU32Ptr(0x5D4594, 1098568) = nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperMagicShop");
		*getMemU32Ptr(0x5D4594, 1098484) = nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperPurple");
		*getMemU32Ptr(0x5D4594, 1097292) = nox_xxx_getTTByNameSpriteMB_44CFC0("ShopkeeperYellow");
		*getMemU32Ptr(0x5D4594, 1107040) = 1;
	}
	if (a1 == *getMemU32Ptr(0x5D4594, 1098396)) {
		result = nox_xxx_gLoadImg_42F970("ShopKeeperPic");
		v1[15] = result;
	} else if (a1 == *getMemU32Ptr(0x5D4594, 1098560)) {
		result = nox_xxx_gLoadImg_42F970("ShopKeeperWarriorPic");
		v1[15] = result;
	} else if (a1 == *getMemU32Ptr(0x5D4594, 1098556)) {
		result = nox_xxx_gLoadImg_42F970("ShopKeeperConjurerPic");
		v1[15] = result;
	} else if (a1 == *getMemU32Ptr(0x5D4594, 1098564)) {
		result = nox_xxx_gLoadImg_42F970("ShopKeeperWizardPic");
		v1[15] = result;
	} else if (a1 == *getMemU32Ptr(0x5D4594, 1098572)) {
		result = nox_xxx_gLoadImg_42F970("ShopKeeperLandOfTheDeadPic");
		v1[15] = result;
	} else if (a1 == *getMemU32Ptr(0x5D4594, 1098568)) {
		result = nox_xxx_gLoadImg_42F970("ShopKeeperMagicShopPic");
		v1[15] = result;
	} else if (a1 == *getMemU32Ptr(0x5D4594, 1098484)) {
		result = nox_xxx_gLoadImg_42F970("ShopKeeperPurplePic");
		v1[15] = result;
	} else {
		if (a1 == *getMemU32Ptr(0x5D4594, 1097292)) {
			result = nox_xxx_gLoadImg_42F970("ShopKeeperBrownPic");
		} else {
			result = nox_xxx_gLoadImg_42F970("ShopKeeperPic");
		}
		v1[15] = result;
	}
	return result;
}

//----- (00479280) --------------------------------------------------------
void sub_479280() {
	if (dword_5d4594_1098624) {
		sub_467680();
		dword_5d4594_1098624 = 0;
		dword_5d4594_1098628 = 0;
		sub_478F10();
		sub_44D8F0();
		nox_window_set_hidden(*(int*)&dword_5d4594_1098576, 1);
		nox_xxx_wnd_46ABB0(*(int*)&dword_5d4594_1098576, 0);
		sub_467C10();
		nox_client_setCursorType_477610(0);
		if (!nox_client_getRenderGUI() && *getMemU32Ptr(0x5D4594, 1098612) == 1) {
			nox_client_setRenderGUI(1);
		}
	}
}

//----- (00479300) --------------------------------------------------------
uint32_t* sub_479300(int a1, int a2, int a3, short a4, int a5) {
	uint32_t* result; // eax
	uint32_t* v6;     // ebx
	int* v7;          // edi
	int i;            // esi

	result = sub_4793C0(a1);
	v6 = result;
	if (result) {
		if (!*result) {
			result = nox_new_drawable_for_thing(a1);
			*v6 = result;
			if (!result) {
				return result;
			}
			result[30] |= 0x40000000u;
			*(uint16_t*)(*v6 + 294) = a4;
			*(uint16_t*)(*v6 + 292) = a4;
			if (*(uint32_t*)(*v6 + 112) & 0x13001000) {
				v7 = (int*)(*v6 + 432);
				for (i = 0; i < 4; ++i) {
					if (*(char*)(i + a5) == -1) {
						*v7 = 0;
					} else {
						*v7 = nox_xxx_modifGetDescById_413330(*(unsigned char*)(i + a5));
					}
					++v7;
				}
			}
			v6[1] = 0;
		}
		v6[v6[1] + 2] = a2;
		result = (uint32_t*)(v6[1] + 1);
		v6[34] = a3;
		v6[1] = result;
	}
	return result;
}

//----- (004793C0) --------------------------------------------------------
char* sub_4793C0(int a1) {
	int v1;            // ebx
	unsigned char* v2; // esi
	int v3;            // ecx
	unsigned char* v4; // eax

	v1 = 0;
	v2 = getMemAt(0x5D4594, 1098636);
	while (1) {
		v3 = 0;
		v4 = v2;
		do {
			if (*((uint32_t*)v4 + 1) && *(uint32_t*)(*(uint32_t*)v4 + 108) == a1 &&
				!(*(uint32_t*)(*(uint32_t*)v4 + 112) & 0x4000000)) {
				return (char*)getMemAt(0x5D4594, 1098636 + 140 * (v1 + 10 * v3));
			}
			++v3;
			v4 += 1400;
		} while (v3 < 6);
		v2 += 140;
		++v1;
		if ((int)v2 < (int)getMemAt(0x5D4594, 1100036)) {
			continue;
		}
		break;
	}
	return sub_479430();
}

//----- (00479430) --------------------------------------------------------
char* sub_479430() {
	int v0;            // esi
	unsigned char* v1; // edx
	int v2;            // eax
	unsigned char* v3; // ecx

	v0 = 0;
	v1 = getMemAt(0x5D4594, 1098640);
	while (1) {
		v2 = 0;
		v3 = v1;
		do {
			if (!*(uint32_t*)v3) {
				return (char*)getMemAt(0x5D4594, 1098636 + 140 * (v0 + 10 * v2));
			}
			++v2;
			v3 += 1400;
		} while (v2 < 6);
		v1 += 140;
		++v0;
		if ((int)v1 < (int)getMemAt(0x5D4594, 1100040)) {
			continue;
		}
		break;
	}
	return 0;
}

//----- (00479480) --------------------------------------------------------
char* sub_479480(int a1) {
	char* result; // eax
	char* v2;     // esi

	result = sub_4780A0(a1);
	v2 = result;
	if (result) {
		sub_4794D0((int)result, a1);
		result = (char*)(*((uint32_t*)v2 + 1) - 1);
		*((uint32_t*)v2 + 1) = result;
		if (!result) {
			result = (char*)nox_xxx_spriteDelete_45A4B0(*(uint64_t**)v2);
			*(uint32_t*)v2 = 0;
			*((uint32_t*)v2 + 34) = 0;
		}
	}
	return result;
}

//----- (004794D0) --------------------------------------------------------
int sub_4794D0(int a1, int a2) {
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

//----- (00479590) --------------------------------------------------------
int sub_479590() { return dword_5d4594_1098628; }

//----- (004795A0) --------------------------------------------------------
void sub_4795A0(int a1) {
	if (dword_5d4594_1098628 == 4) {
		if (a1 != 4) {
			sub_467680();
		}
		dword_5d4594_1098628 = a1;
	} else {
		dword_5d4594_1098628 = a1;
	}
}

//----- (00479690) --------------------------------------------------------
int sub_479690(int a1, short a2, short a3, int a4) {
	int result; // eax

	result = a4;
	dword_5d4594_1098616 = 0;
	if (a4) {
		if (a4 == 1) {
			result = nox_client_tradeXxxSellAccept_4796D0(a2);
		} else {
			result = sub_479700(a3, a4);
		}
	}
	return result;
}

//----- (004796D0) --------------------------------------------------------
int nox_client_tradeXxxSellAccept_4796D0(short a1) {
	char v2[4]; // [esp+0h] [ebp-4h]

	v2[0] = -55;
	v2[1] = 24;
	*(uint16_t*)&v2[2] = a1;
	return nox_netlist_addToMsgListCli_40EBC0(31, 0, v2, 4);
}

//----- (00479700) --------------------------------------------------------
int sub_479700(short a1, char a2) {
	char v3[5]; // [esp+0h] [ebp-8h]

	v3[0] = -55;
	v3[1] = 25;
	*(uint16_t*)&v3[2] = a1;
	v3[4] = a2;
	return nox_netlist_addToMsgListCli_40EBC0(31, 0, v3, 5);
}

//----- (00479810) --------------------------------------------------------
void sub_479810() { dword_5d4594_1098620 = 0; }

//----- (00479820) --------------------------------------------------------
int sub_479820(int a1, short a2) {
	int result; // eax

	result = sub_479840(a2);
	dword_5d4594_1098620 = 0;
	return result;
}

//----- (00479840) --------------------------------------------------------
int sub_479840(short a1) {
	char v2[4]; // [esp+0h] [ebp-4h]

	v2[0] = -55;
	v2[1] = 26;
	*(uint16_t*)&v2[2] = a1;
	return nox_netlist_addToMsgListCli_40EBC0(31, 0, v2, 4);
}

//----- (00479870) --------------------------------------------------------
int sub_479870() { return dword_5d4594_1098628 == 2; }

//----- (00479880) --------------------------------------------------------
bool sub_479880(uint32_t* a1) { return nox_xxx_pointInRect_4281F0((int2*)a1, (int4*)getMemAt(0x5D4594, 1098380)); }

//----- (004798A0) --------------------------------------------------------
int sub_4798A0(uint32_t* a1) {
	int v1;            // eax
	int v2;            // ecx
	int v3;            // edx
	int v4;            // edx
	unsigned char* v5; // eax

	v1 = nox_xxx_pointInRect_4281F0((int2*)a1, (int4*)getMemAt(0x5D4594, 1098380));
	if (!v1) {
		return 0;
	}
	v2 = (*a1 - *getMemU32Ptr(0x5D4594, 1098380)) / 50;
	v3 = (a1[1] - *getMemU32Ptr(0x5D4594, 1098384) + dword_5d4594_1107036) / 50;
	if (v2 >= 6) {
		v2 = 5;
	}
	if (v3 >= 10) {
		v3 = 9;
	}
	v4 = 35 * (v3 + 10 * v2);
	v5 = getMemAt(0x5D4594, 1098636 + 4 * v4);
	if (!*getMemU32Ptr(0x5D4594, 1098640 + 4 * v4)) {
		return 0;
	}
	*(uint32_t*)(*(uint32_t*)v5 + 128) = *((uint32_t*)v5 + 2);
	return *(uint32_t*)v5;
}
// 4798B5: variable 'v1' is possibly undefined

//----- (00479950) --------------------------------------------------------
int sub_479950() {
	void* v2; // [esp+0h] [ebp-4h]

	if (wndIsShown_nox_xxx_wndIsShown_46ACC0(*(int*)&dword_5d4594_1123524) == 1) {
		return 0;
	}
	*getMemU8Ptr(0x5D4594, 1123516) = 0;
	LOWORD(v2) = 720;
	BYTE2(v2) = 0;
	nox_netlist_addToMsgListCli_40EBC0(31, 0, &v2, 3);
	return 1;
}

//----- (004799A0) --------------------------------------------------------
int sub_4799A0() {
	int result;    // eax
	uint32_t* v1;  // edi
	uint32_t* v2;  // ebp
	uint32_t* v3;  // eax
	int v4;        // esi
	uint32_t* v5;  // ebx
	char* v6;      // eax
	uint32_t* v7;  // eax
	char* v8;      // [esp-18h] [ebp-1Ch]
	char* v9;      // [esp-14h] [ebp-18h]
	uint32_t* v10; // [esp+0h] [ebp-4h]

	*getMemU32Ptr(0x5D4594, 1107052) = nox_color_rgb_4344A0(240, 128, 64);
	result = nox_new_window_from_file("Dialog.wnd", nox_xxx_guiDialog_479B00);
	dword_5d4594_1123524 = result;
	if (result) {
		nox_xxx_wndSetWindowProc_46B300(result, sub_479BE0);
		nox_xxx_wndSetDrawFn_46B340(*(int*)&dword_5d4594_1123524, sub_479CB0);
		nox_gui_winSetFunc96_46B070(*(int*)&dword_5d4594_1123524, sub_479D00);
		v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1123524, 3904);
		v2 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1123524, 3903);
		v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1123524, 3902);
		v3 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1123524, 3901);
		v4 = (int)v3;
		v5 = (uint32_t*)v3[8];
		v9 = nox_xxx_gLoadImg_42F970("UISliderLit");
		v8 = nox_xxx_gLoadImg_42F970("UISliderLit");
		v6 = nox_xxx_gLoadImg_42F970("UISlider");
		sub_4B5700((int)v1, 0, 0, (int)v6, (int)v8, (int)v9);
		nox_xxx_wnd_46B280((int)v1, v4);
		nox_xxx_wnd_46B280((int)v2, v4);
		nox_xxx_wnd_46B280((int)v10, v4);
		v5[9] = v1;
		v5[7] = v2;
		v5[8] = v10;
		v7 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1123524, 3906);
		nox_xxx_wndSetDrawFn_46B340((int)v7, sub_479C40);
		nox_xxx_wnd_46ABB0(*(int*)&dword_5d4594_1123524, 0);
		nox_window_set_hidden(*(int*)&dword_5d4594_1123524, 1);
		dword_5d4594_1123520 = 0;
		result = 1;
	}
	return result;
}

//----- (00479B00) --------------------------------------------------------
int nox_xxx_guiDialog_479B00(int a1, int a2, int* a3, int a4) {
	int v3;     // esi
	int result; // eax

	if (a2 != 16391) {
		return 0;
	}
	v3 = nox_xxx_wndGetID_46B0A0(a3);
	if (sub_45D9B0()) {
		return 0;
	}
	nox_xxx_clientPlaySoundSpecial_452D80(766, 100);
	switch (v3) {
	case 3906:
		sub_479950();
		result = 0;
		break;
	case 3907:
		nox_xxx_playDialogFile_44D900(*getMemIntPtr(0x5D4594, 1115312), 100);
		result = 0;
		break;
	case 3908:
		*getMemU8Ptr(0x5D4594, 1123516) = 1;
		LOWORD(a2) = 720;
		BYTE2(a2) = 1;
		nox_netlist_addToMsgListCli_40EBC0(31, 0, &a2, 3);
		result = 0;
		break;
	case 3909:
		*getMemU8Ptr(0x5D4594, 1123516) = 2;
		BYTE2(a2) = 2;
		LOWORD(a2) = 720;
		nox_netlist_addToMsgListCli_40EBC0(31, 0, &a2, 3);
		return 0;
	default:
		return 0;
	}
	return result;
}
// 479B4D: variable 'v4' is possibly undefined

//----- (00479BE0) --------------------------------------------------------
int sub_479BE0(uint32_t* a1, int a2, unsigned int a3, int a4) {
	switch (a2) {
	case 5:
	case 6:
	case 7:
	case 9:
	case 10:
	case 11:
	case 13:
	case 14:
	case 15:
		nox_xxx_wndPointInWnd_46AAB0(a1, (unsigned short)a3, a3 >> 16);
		break;
	default:
		return 1;
	}
	return 1;
}

//----- (00479C40) --------------------------------------------------------
int nox_xxx_wndButtonDrawNoImg_4A81D0(int a1, int a2);
int sub_479C40(uint32_t* a1, int a2) {
	char v2;   // bl
	int yTop;  // [esp+8h] [ebp-8h]
	int xLeft; // [esp+Ch] [ebp-4h]

	v2 = nox_xxx_bookGet_430B40_get_mouse_prev_seq();
	if (!sub_44D930() && (v2 & 0x7Fu) < 0x1E && v2 & 8) {
		nox_client_wndGetPosition_46AA60(a1, &xLeft, &yTop);
		sub_49CD30(xLeft, yTop, a1[2], a1[3] - 2, *getMemIntPtr(0x5D4594, 1107052), 4);
	}
	return nox_xxx_wndButtonDrawNoImg_4A81D0((int)a1, a2);
}

//----- (00479CB0) --------------------------------------------------------
int sub_479CB0(int a1, int a2) {
	int v2; // esi
	int v4; // [esp+4h] [ebp-8h]
	int v5; // [esp+8h] [ebp-4h]

	v2 = *(uint32_t*)(a2 + 24);
	nox_client_wndGetPosition_46AA60(*(uint32_t**)&dword_5d4594_1123524, &v4, &v5);
	nox_client_drawImageAt_47D2C0(v2, nox_win_width - NOX_DEFAULT_WIDTH, nox_win_height - NOX_DEFAULT_HEIGHT);
	return 1;
}

//----- (00479D00) --------------------------------------------------------
int sub_479D00() { return 1; }

//----- (00479D10) --------------------------------------------------------
int sub_479D10() {
	int result; // eax

	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1123524);
	result = 0;
	dword_5d4594_1123524 = 0;
	dword_5d4594_1123520 = 0;
	return result;
}

//----- (0047A260) --------------------------------------------------------
int sub_47A260() { return dword_5d4594_1123520; }

//----- (0047D380) --------------------------------------------------------
int sub_47D380(int a1, int a2) {
	int v2; // eax
	int v3; // ecx
	int v4; // edx
	int v5; // esi

	v2 = a1;
	v3 = a2;
	if (a1 > a2) {
		v2 = a2;
		v3 = a1;
	}
	v4 = nox_draw_curDrawData_3799572->clip.min_x;
	if (v2 >= v4) {
		if (v2 >= nox_draw_curDrawData_3799572->clip.max_x) {
			return 0;
		}
	} else {
		v2 = nox_draw_curDrawData_3799572->clip.min_x;
	}
	v5 = nox_draw_curDrawData_3799572->clip.max_x;
	if (v3 < v5) {
		if (v3 < v4) {
			return 0;
		}
	} else {
		v3 = nox_draw_curDrawData_3799572->clip.max_x;
	}
	if (v2 == v3) {
		return 0;
	}
	if (v2 != v4 || v3 != v5) {
		*getMemU32Ptr(0x973F18, 52) = v2;
		*getMemU32Ptr(0x973F18, 12) = v3;
		dword_5d4594_3799452 = 1;
	}
	return 1;
}

//----- (0047DBC0) --------------------------------------------------------
unsigned char sub_47DBC0() { return getMemByte(0x5D4594, 1193128); }

//----- (0047FCE0) --------------------------------------------------------
int sub_47FCE0(uint32_t* a1, int a2) {
	int v2;            // edx
	unsigned char* v3; // eax
	int v4;            // eax
	int v5;            // esi

	v2 = 0;
	if (*(int*)&dword_5d4594_3804684 > 0) {
		v3 = getMemAt(0x973F18, 6092);
		while (*((uint32_t*)v3 - 1) != a1[3] || *(uint32_t*)v3 != a1[2] || *((uint32_t*)v3 + 1) != a1[21] ||
			   *((uint32_t*)v3 + 2) != a1[26]) {
			++v2;
			v3 += 16;
			if (v2 >= *(int*)&dword_5d4594_3804684) {
				goto LABEL_8;
			}
		}
		return 1;
	}
LABEL_8:
	v4 = 16 * dword_5d4594_3804684;
	v5 = dword_5d4594_3804684 + 1;
	*getMemU32Ptr(0x973F18, 6088 + v4) = a1[3];
	*getMemU32Ptr(0x973F18, 6092 + v4) = a1[2];
	*getMemU32Ptr(0x973F18, 6096 + v4) = a1[21];
	*getMemU32Ptr(0x973F18, 6100 + v4) = a1[26];
	dword_5d4594_3804684 = v5;
	return 1;
}

// 4514E0: using guessed type void  nullsub_4(uint32_t, uint32_t, uint32_t, uint32_t);

//----- (00480220) --------------------------------------------------------
uint8_t* sub_480220(uint8_t* a1, uint8_t* a2) {
	unsigned int v2; // edx
	uint8_t* result; // eax

	result = a2;
	*a1 = 8 * *a2;
	LOWORD(v2) = *(uint16_t*)a2;
	a1[1] = (v2 >> 3) & 0xFC;
	a1[2] = a2[1] & 0xF8;
	return result;
}
// 480232: variable 'v2' is possibly undefined

//----- (00480250) --------------------------------------------------------
uint16_t* sub_480250(uint8_t* a1, uint16_t* a2) {
	uint16_t* result; // eax

	result = a2;
	*a2 = (*a1 >> 3) | (8 * (a1[1] & 0xFC | (32 * (a1[2] & 0xF8))));
	return result;
}

//----- (00480EF0) --------------------------------------------------------
int nox_getBackbufferPitch();
int* nox_xxx_edgeDraw_480EF0(int a1, int a2, int a3, int* a4, int* a5, int a6, int a7, int a8, int a9, int a10) {
	int* result;                                // eax
	int v10;                                    // ebx
	char v11;                                   // cl
	int v12;                                    // esi
	int v13;                                    // edx
	int* v14;                                   // ebp
	int v15;                                    // edi
	char* v16;                                  // ebp
	int v17;                                    // ecx
	int v18;                                    // esi
	int v19;                                    // esi
	int v20;                                    // edx
	int v21;                                    // edx
	int v22;                                    // ecx
	int v23;                                    // edx
	int v24;                                    // eax
	int v25;                                    // ecx
	int v26;                                    // esi
	int v27;                                    // ecx
	int i;                                      // esi
	char v29;                                   // al
	unsigned int v30;                           // ecx
	const void* v31;                            // esi
	int v32;                                    // edi
	int j;                                      // esi
	char v34;                                   // al
	char v35;                                   // dl
	int v36;                                    // ecx
	int v37;                                    // esi
	unsigned char v38;                          // dl
	char v39;                                   // al
	int v40;                                    // [esp+10h] [ebp-48h]
	int v41;                                    // [esp+14h] [ebp-44h]
	int v43;                                    // [esp+1Ch] [ebp-3Ch]
	int v44;                                    // [esp+20h] [ebp-38h]
	int2 v45;                                   // [esp+24h] [ebp-34h]
	int2 v46;                                   // [esp+2Ch] [ebp-2Ch]
	int v47;                                    // [esp+34h] [ebp-24h]
	int v48;                                    // [esp+38h] [ebp-20h]
	int v49;                                    // [esp+3Ch] [ebp-1Ch]
	int v50[3];                                 // [esp+40h] [ebp-18h]
	int v51[3];                                 // [esp+4Ch] [ebp-Ch]
	char* v52;                                  // [esp+60h] [ebp+8h]
	int v53;                                    // [esp+64h] [ebp+Ch]
	int v54;                                    // [esp+64h] [ebp+Ch]
	int v55;                                    // [esp+70h] [ebp+18h]
	int v56;                                    // [esp+74h] [ebp+1Ch]
	unsigned char v57;                          // [esp+7Ch] [ebp+24h]
	unsigned char v58;                          // [esp+7Ch] [ebp+24h]
	unsigned char v59;                          // [esp+7Ch] [ebp+24h]
	unsigned char v60;                          // [esp+7Ch] [ebp+24h]
	unsigned char v61;                          // [esp+7Ch] [ebp+24h]

	result = (int*)a1;
	v10 = 0;
	v44 = 0;
	if (!result) {
		return 0;
	}
	v11 = nox_video_bag_image_type(a1);
	v45.field_4 = 0;
	v45.field_0 = 0;
	v46.field_4 = 0;
	v46.field_0 = 0;
	if ((v11 & 0x3F) != 3) {
		return result;
	}
	result = (int*)nox_video_getImagePixdata_42FB30(a1);
	if (!result) {
		return result;
	}
	v12 = *result;
	v13 = result[1] - a9;
	v43 = *result;
	v40 = *result;
	v41 = v13;
	if (v13 <= 0) {
		return result;
	}
	v14 = result + 3;
	result = (int*)(result[2] + a2);
	v15 = *v14;
	v16 = (char*)v14 + 5;
	v17 = v15 + a3;
	v53 = v15 + a3;
	if (!((int)result <= *(int*)&dword_5d4594_3807116 && v17 <= *(int*)&dword_5d4594_3807152)) {
		return result;
	}
	if ((int)result < *(int*)&dword_5d4594_3807140) {
		if ((int)result + v12 < *(int*)&dword_5d4594_3807140) {
			return result;
		}
		v10 = *(int*)&dword_5d4594_3807140 - (int)result;
		v40 = v12 - (*(int*)&dword_5d4594_3807140 - (int)result);
		result = *(int**)&dword_5d4594_3807140;
	}
	if ((int)result + v40 > *(int*)&dword_5d4594_3807116) {
		v40 = *(int*)&dword_5d4594_3807116 - (int)result;
	}
	if (v17 < *(int*)&dword_5d4594_3807136) {
		if (v13 + v17 < *(int*)&dword_5d4594_3807136) {
			return result;
		}
		v18 = *(int*)&dword_5d4594_3807136 - v17;
		v17 = *(int*)&dword_5d4594_3807136;
		v13 -= v18;
		v53 = *(int*)&dword_5d4594_3807136;
		v41 = v13;
		v44 = v18;
	}
	if (v13 + v17 > *(int*)&dword_5d4594_3807152) {
		v13 = *(int*)&dword_5d4594_3807152 - v17;
		v41 = *(int*)&dword_5d4594_3807152 - v17;
	}
	if (!(v40 > 0 && v13 > 0)) {
		return result;
	}
	v19 = a6;
	if (a6 < v17) {
		v19 = v17;
	}
	if (v13 + v17 > v19) {
		v41 = v19 - v17;
	}
	if (a7 <= (int)result + v10) {
		v20 = 0;
	} else {
		v20 = a7 - v10 - (int)result;
		v10 = a7 - (int)result;
		v40 -= v20;
	}
	if (a8 < (int)result + v40 + v10) {
		v40 = a8 - v10 - (int)result;
	}
	if (v40 <= 0) {
		return result;
	}
	v21 = (int)nox_pixbuffer_rows_3798784[v17] + 2 * ((int)result + v20);
	v22 = a4[1];
	v52 = (char*)v21;
	v23 = *a4;
	v49 = a4[2] << 8;
	v23 <<= 8;
	v24 = (*a5 << 8) - v23;
	v47 = v23;
	v48 = v22 << 8;
	v25 = a5[1] << 8;
	v26 = a5[2] << 8;
	v50[0] = v24 / v43;
	v50[1] = (v25 - v48) / v43;
	v50[2] = (v26 - v49) / v43;
	if (v44 > 0) {
		v27 = v44;
		do {
			for (i = v43; i > 0; i -= v57) {
				v29 = *v16;
				v57 = v16[1];
				v16 += 2;
				LOBYTE(a1) = v29;
				if (v29 == 2) {
					v16 += 2 * v57;
				}
			}
			--v27;
		} while (v27);
	}
	v46.field_4 = v53;
	sub_473970(&v46, &v45);
	v30 = 0;
	v31 = 0;
	result = (int*)(v41 - 1);
	if (!v41) {
		return result;
	}
	v56 = v41;
	int bpitch = nox_getBackbufferPitch();
	while (1) {
		v32 = 0;
		v54 = 0;
		v55 = 0;
		if (nox_client_highResFrontWalls_80820 || !(v45.field_4 & 1)) {
			if (v10 <= 0) {
				v35 = a1;
			} else {
				do {
					v35 = *v16;
					v59 = v16[1];
					v16 += 2;
					LOBYTE(a1) = v35;
					if (v35 == 2) {
						v16 += 2 * v59;
					}
					v32 += v59;
				} while (v32 < v10);
			}
			if (v32 <= v10) {
				v36 = v40;
			} else {
				v36 = v40;
				v37 = v40;
				if (v32 - v10 <= v40) {
					v37 = v32 - v10;
				}
				if (v35 == 2) {
					v51[0] = v47 + v10 * v50[0];
					v51[1] = v48 + v10 * v50[1];
					v51[2] = v49 + v10 * v50[2];
					sub_480860(v52, &v16[-2 * (v32 - v10)], v37, v51, v50);
					v36 = v40;
					v54 = 2 * v37;
				}
				v55 = 0;
			}
			if (v32 < v36 + v10) {
				do {
					v38 = v16[1];
					LOBYTE(a1) = *v16;
					v16 += 2;
					v60 = v38;
					if ((uint8_t)a1 == 2) {
						if (v38 > v10 + v36 - v32) {
							v55 = v32 + v38 - v36 - v10;
							v60 = -(char)(v32 + -(char)v36 - v10);
						}
						v51[0] = v47 + v50[0] * v32;
						v51[2] = v49 + v50[2] * v32;
						v51[1] = v48 + v50[1] * v32;
						sub_480860(&v52[2 * (v32 - v10)], v16, v60, v51, v50);
						v36 = v40;
						v16 += 2 * v60;
					}
					v32 += v60;
					v54 += 2 * v60;
				} while (v32 < v36 + v10);
				if (v55) {
					if ((uint8_t)a1 == 2) {
						v16 += 2 * v55;
					}
					v32 += v55;
				}
			}
			for (; v32 < v43; v32 += v61) {
				v39 = *v16;
				v61 = v16[1];
				v16 += 2;
				LOBYTE(a1) = v39;
				if (v39 == 2) {
					v16 += 2 * v61;
				}
			}
		} else {
			if (v31 && v30) {
				memcpy(v52, v31, v30);
			}
			for (j = v43; j > 0; j -= v58) {
				v34 = *v16;
				v58 = v16[1];
				v16 += 2;
				LOBYTE(a1) = v34;
				if (v34 == 2) {
					v16 += 2 * v58;
				}
			}
		}
		v31 = v52;
		v52 += bpitch;
		++v45.field_4;
		result = (int*)--v56;
		if (!v56) {
			break;
		}
		v30 = v54;
	}
	return result;
}

//----- (00481410) --------------------------------------------------------
void sub_481410() { nox_xxx_waypointCounterMB_587000_154948 = -1; }

//----- (004815E0) --------------------------------------------------------
char nox_xxx_tileDraw_4815E0(uint32_t* a1, int a2) {
	unsigned int v2; // ebx
	int v3;          // eax
	char* v4;        // esi
	unsigned int v5; // edi
	int v6;          // eax
	int v7;          // edx
	unsigned int v8; // ebp
	signed int v9;   // eax
	char* v10;       // edi
	char* v11;       // esi
	char v12;        // cl
	char* v14;       // [esp+10h] [ebp+4h]
	int i;           // [esp+14h] [ebp+8h]

	v2 = dword_5d4594_3798804 * (dword_5d4594_3798840 + a1[1] - dword_5d4594_3798824) + (uint32_t)nox_video_tileBuf_ptr_3798796 +
		 ((dword_5d4594_3798836 + *a1 - dword_5d4594_3798820) << getMemByte(0x973F18, 7696));
	if (v2 >= nox_video_tileBuf_end_3798844) {
		v2 += (uint32_t)nox_video_tileBuf_ptr_3798796 - (uint32_t)nox_video_tileBuf_end_3798844;
	}
	v3 = nox_video_getImagePixdata_42FB30(a2);
	v4 = (char*)v3;
	v14 = (char*)v3;
	if (v3) {
		v5 = nox_video_tileBuf_end_3798844;
		if (v2 + *getMemIntPtr(0x973CE0, 376) >= nox_video_tileBuf_end_3798844) {
			v6 = 0;
			for (i = 0;; v6 = i) {
				v7 = *getMemU32Ptr(0x973CE0, 192 + v6) << getMemByte(0x973F18, 7696);
				v8 = *getMemU32Ptr(0x973CE0, 384 + v6) << getMemByte(0x973F18, 7696);
				if (v2 + v7 + v8 < v5) {
					memcpy((void*)(v7 + v2), v4, 4 * (v8 >> 2));
					v11 = &v4[4 * (v8 >> 2)];
					v10 = (char*)(v7 + v2 + 4 * (v8 >> 2));
					v12 = v8;
				} else {
					v9 = v5 - v7 - v2;
					if (v9 > 0) {
						memcpy((void*)(v7 + v2), v4, v9);
						memcpy(nox_video_tileBuf_ptr_3798796, &v14[v9], v8 - v9);
						v5 = nox_video_tileBuf_end_3798844;
						v2 += (uint32_t)nox_video_tileBuf_ptr_3798796 - (uint32_t)nox_video_tileBuf_end_3798844;
						goto LABEL_13;
					}
					v2 += (uint32_t)nox_video_tileBuf_ptr_3798796 - v5;
					memcpy((void*)(v7 + v2), v4, 4 * (v8 >> 2));
					v11 = &v4[4 * (v8 >> 2)];
					v10 = (char*)(v7 + v2 + 4 * (v8 >> 2));
					v12 = v8;
				}
				memcpy(v10, v11, v12 & 3);
				v5 = nox_video_tileBuf_end_3798844;
			LABEL_13:
				v2 += dword_5d4594_3798804;
				v14 += v8;
				LOBYTE(v3) = i + 4;
				i += 4;
				if (i >= 184) {
					return v3;
				}
				v4 = v14;
			}
		} else {
			LOBYTE(v3) = sub_4831C0(v3, v2);
		}
	}
	return v3;
}

//----- (00481770) --------------------------------------------------------
bool get_nox_client_texturedFloors_154956();
uint32_t dword_5d4594_1193156 = 0;
char* sub_481770(uint32_t* a1, int a2, unsigned short a3) {
	unsigned char v4; // cl
	unsigned int v5;  // edx
	unsigned int v6;  // esi
	int v7;           // ebx
	int v8;           // eax
	int v9;           // edi
	int v10;          // ebx
	char* result;     // eax
	int v12;          // [esp+Ch] [ebp+Ch]

	nox_tileDef_t* tile = &nox_tile_defs_arr[a3];
	if (!get_nox_client_texturedFloors_154956() && (tile->field_58 & 1) == 1) {
		dword_5d4594_1193156 = 1;
	}
	int cl = tile->color_48;
	v4 = getMemByte(0x973F18, 7696);
	v5 = nox_video_tileBuf_end_3798844;
	v6 = dword_5d4594_3798804 * (dword_5d4594_3798840 + a1[1] - dword_5d4594_3798824) + (uint32_t)nox_video_tileBuf_ptr_3798796 +
		 ((dword_5d4594_3798836 + *a1 - dword_5d4594_3798820) << getMemByte(0x973F18, 7696));
	if (v6 >= nox_video_tileBuf_end_3798844) {
		v6 += (uint32_t)nox_video_tileBuf_ptr_3798796 - (uint32_t)nox_video_tileBuf_end_3798844;
	}
	if (v6 + *getMemU32Ptr(0x973CE0, 376) < nox_video_tileBuf_end_3798844) {
		result = (char*)sub_484450(cl, v6);
	} else {
		v7 = 0;
		v12 = 0;
		while (1) {
			v8 = *getMemU32Ptr(0x973CE0, 192 + v7) << v4;
			v9 = *getMemU32Ptr(0x973CE0, 384 + v7) << v4;
			if (v6 + v9 + v8 < v5) {
				result = (char*)sub_49D1C0(v6 + v8, cl, v9);
				v5 = nox_video_tileBuf_end_3798844;
			} else {
				v10 = v5 - v8 - v6;
				if (v10 <= 0) {
					v6 += (uint32_t)nox_video_tileBuf_ptr_3798796 - v5;
					result = (char*)sub_49D1C0(v6 + v8, cl, v9);
					v5 = nox_video_tileBuf_end_3798844;
					v7 = v12;
				} else {
					sub_49D1C0(v6 + v8, cl, v5 - v8 - v6);
					sub_49D1C0(nox_video_tileBuf_ptr_3798796, cl, v9 - v10);
					v5 = nox_video_tileBuf_end_3798844;
					v7 = v12;
					result = (char*)((uint32_t)nox_video_tileBuf_ptr_3798796 - (uint32_t)nox_video_tileBuf_end_3798844);
					v6 += (uint32_t)nox_video_tileBuf_ptr_3798796 - (uint32_t)nox_video_tileBuf_end_3798844;
				}
			}
			v7 += 4;
			v6 += dword_5d4594_3798804;
			v12 = v7;
			if (v7 >= 184) {
				break;
			}
			v4 = getMemByte(0x973F18, 7696);
		}
	}
	return result;
}

//----- (00481900) --------------------------------------------------------
char nox_xxx_drawTexEdgesProbably_481900(uint32_t* a1, uint32_t* a2) {
	int v2;            // ebx
	int v3;            // esi
	int v4;            // edx
	int v5;            // ebp
	int v6;            // eax
	int v7;            // edi
	int v8;            // edi
	int v9;            // eax
	uint8_t* v10;      // ebx
	int v11;           // ebp
	unsigned int v12;  // esi
	unsigned int v13;  // edi
	char* v14;         // ebp
	char* v15;         // edx
	int v16;           // ebx
	char* v17;         // edi
	unsigned int v18;  // ecx
	char* v19;         // esi
	unsigned int v20;  // eax
	int v22;           // eax
	int v24;           // [esp+10h] [ebp-1Ch]
	unsigned char v25; // [esp+14h] [ebp-18h]
	char* v26;         // [esp+14h] [ebp-18h]
	char* v27;         // [esp+18h] [ebp-14h]
	int v28;           // [esp+1Ch] [ebp-10h]
	unsigned char v29; // [esp+20h] [ebp-Ch]
	int v30;           // [esp+28h] [ebp-4h]
	unsigned int v31;  // [esp+30h] [ebp+4h]
	unsigned char v32; // [esp+34h] [ebp+8h]
	char* v33;         // [esp+34h] [ebp+8h]
	int* addr;

	nox_tileDef_t* tile = &nox_tile_defs_arr[a2[0]];

	v2 = dword_5d4594_3798824;
	v3 = tile->data_32[a2[1] + tile->field_46];
	v4 = a2[2];
	v5 = dword_5d4594_3798836;
	addr = (*getMemU32Ptr(0x85B3FC, 28676 + 60 * v4) + 4 * (a2[3] + *getMemU16Ptr(0x85B3FC, 28690 + 60 * v4)));
	v6 = *(uint32_t*)addr;
	v7 = dword_5d4594_3798840;
	*getMemU32Ptr(0x5D4594, 2523980 + 4 * v4) = 1;
	v8 = dword_5d4594_3798804 * (v7 + a1[1] - v2) + (uint32_t)nox_video_tileBuf_ptr_3798796 +
		 ((v5 + *a1 - dword_5d4594_3798820) << getMemByte(0x973F18, 7696));
	v9 = nox_video_getImagePixdata_42FB30(v6);
	if (v9) {
		v32 = *(uint8_t*)v9;
		v25 = *(uint8_t*)(v9 + 1);
		v10 = (uint8_t*)(v9 + 2);
		v9 = nox_video_getImagePixdata_42FB30(v3);
		v11 = v9;
		if (v9) {
			v12 = nox_video_tileBuf_end_3798844;
			v9 = v32;
			v13 = dword_5d4594_3798804 * v32 + v8;
			v31 = v13;
			v14 = (char*)((*getMemU32Ptr(0x973CE0, +4 * v32) << getMemByte(0x973F18, 7696)) + v11);
			v27 = v14;
			if (v13 >= nox_video_tileBuf_end_3798844) {
				v13 += (uint32_t)nox_video_tileBuf_ptr_3798796 - (uint32_t)nox_video_tileBuf_end_3798844;
				v31 = v13;
			}
			v28 = v32;
			v30 = v25;
			if (v32 <= (int)v25) {
				do {
					v24 = *getMemU32Ptr(0x973CE0, 384 + 4 * v9);
					v26 = v14;
					v15 = (char*)(v13 + (*getMemU32Ptr(0x973CE0, 192 + 4 * v9) << getMemByte(0x973F18, 7696)));
					if (v24 > 0) {
						do {
							LOBYTE(v9) = *v10;
							v29 = v10[1];
							v33 = v10 + 2;
							v16 = v29 << getMemByte(0x973F18, 7696);
							switch ((uint8_t)v9) {
							case 2:
								if ((unsigned int)&v15[v16] < v12) {
									v19 = v33;
									v18 = v29 << getMemByte(0x973F18, 7696);
									v17 = v15;
								} else {
									memcpy(v15, v33, v12 - (uint32_t)v15);
									v17 = nox_video_tileBuf_ptr_3798796;
									v18 = v16 - (v12 - (uint32_t)v15);
									v19 = &v33[v12 - (uint32_t)v15];
								}
								memcpy(v17, v19, v18);
								v12 = nox_video_tileBuf_end_3798844;
								v13 = v31;
								v33 += v16;
								break;
							case 3:
								if ((unsigned int)&v15[v16] < v12) {
									memcpy(v15, v14, v16);
								} else {
									v20 = v12 - (uint32_t)v15;
									memcpy(v15, v14, v20);
									v14 = v26;
									memcpy(nox_video_tileBuf_ptr_3798796, &v26[v20], v16 - v20);
								}
								v12 = nox_video_tileBuf_end_3798844;
								v13 = v31;
								break;
							case 1:
								break;
							default:
								return v9;
							}
							v15 += v16;
							v22 = v24 - v29;
							v14 += v16;
							v24 -= v29;
							v26 = v14;
							if ((unsigned int)v15 >= v12) {
								v15 += (uint32_t)nox_video_tileBuf_ptr_3798796 - v12;
							}
							v10 = v33;
						} while (v22 > 0);
						v9 = v28;
					}
					v14 = &v27[*getMemU32Ptr(0x973CE0, 384 + 4 * v9) << getMemByte(0x973F18, 7696)];
					v13 += dword_5d4594_3798804;
					v27 += *getMemU32Ptr(0x973CE0, 384 + 4 * v9) << getMemByte(0x973F18, 7696);
					v31 = v13;
					if (v13 >= v12) {
						v13 += (uint32_t)nox_video_tileBuf_ptr_3798796 - v12;
						v31 = v13;
					}
					v28 = ++v9;
				} while (v9 <= v30);
			}
		}
	}
	return v9;
}

//----- (00481BF0) --------------------------------------------------------
void nox_xxx_tileCallDrawEdges_481BF0(int a1, int a2) {
	int i; // esi

	for (i = a2; i; i = *(uint32_t*)(i + 16)) {
		func_587000_154944(a1, i);
	}
}
// 5ACD40: invalid function type has been ignored

//----- (00481C20) --------------------------------------------------------
void nox_xxx_tileDrawMB_481C20_A(nox_draw_viewport_t* vp, int v3) {
	int v17; // esi
	int v63; // [esp-8h] [ebp-54h]
	int v16; // edi
	int v62; // [esp-8h] [ebp-54h]
	int v14; // esi
	int v15; // edi
	int v13; // ecx
	int v12; // eax
	int v11; // ebx
	int v4;  // eax
	int v6;  // ecx
	int v10; // esi
	int v9;  // edx
	int v7;  // edx
	int v8;  // eax
	int j;
	int2 v68; // [esp+1Ch] [ebp-30h]
	if (v3 >= *(int*)&dword_5d4594_3798820 + 23) {
		int v71 = nox_getBackbufWidth() + v3;
		if (v71 <= *(int*)&dword_5d4594_3798800 + dword_5d4594_3798820 - 46 ||
			*(int*)&dword_5d4594_3798812 + *(int*)&dword_5d4594_3798828 - 1 >= 128) {
			return;
		}
		if (v71 > *(int*)&dword_5d4594_3798800 + *(int*)&dword_5d4594_3798820) {
			nox_xxx_tileDrawImpl_4826A0(vp);
			return;
		}
		v7 = dword_5d4594_3798828 + 1;
		dword_5d4594_3798828 = v7;
		v8 = dword_5d4594_3798820 + 46;
		dword_5d4594_3798820 += 46;
		j = dword_5d4594_3798812 + v7 - 2;
		v9 = dword_5d4594_3798836 + 46;
		dword_5d4594_3798836 += 46;
		if (*(int*)&dword_5d4594_3798836 >= *(int*)&dword_5d4594_3798800) {
			dword_5d4594_3798836 = v9 - dword_5d4594_3798800;
			v10 = ++dword_5d4594_3798840;
			if (*(int*)&dword_5d4594_3798840 >= *(int*)&dword_5d4594_3798808) {
				dword_5d4594_3798840 = v10 - dword_5d4594_3798808;
			}
		}
		v4 = dword_5d4594_3798800 + v8 - 92;
	} else {
		if (*(int*)&dword_5d4594_3798828 <= 0) {
			return;
		}
		if (v3 < *(int*)&dword_5d4594_3798820 - 23) {
			nox_xxx_tileDrawImpl_4826A0(vp);
			return;
		}
		v4 = dword_5d4594_3798820 - 46;
		j = --dword_5d4594_3798828;
		dword_5d4594_3798820 -= 46;
		dword_5d4594_3798836 -= 46;
		if (*(int*)&dword_5d4594_3798836 < 0) {
			v6 = *(int*)&dword_5d4594_3798840 - 1;
			bool v5 = *(int*)&dword_5d4594_3798840 - 1 < 0;
			dword_5d4594_3798836 += dword_5d4594_3798800;
			--*(int*)&dword_5d4594_3798840;
			if (v5) {
				*(int*)&dword_5d4594_3798840 = *(int*)&dword_5d4594_3798808 + v6;
			}
		}
	}
	int v76 = v4;
	sub_481410();
	v11 = dword_5d4594_3798824;
	unsigned int v74 = dword_5d4594_3798832;
	if (*(int*)&dword_5d4594_3798832 < *(int*)&dword_5d4594_3798832 + *(int*)&dword_5d4594_3798816) {
		v12 = 44 * dword_5d4594_3798832;
		for (int i = 44 * dword_5d4594_3798832;; v12 = i) {
			HIWORD(v13) = HIWORD(ptr_5D4594_2650668); // TODO: why it's doing it?
			v14 = v12 + (uint32_t)(ptr_5D4594_2650668[j]);
			if (*(uint8_t*)v14 & 2) {
				LOWORD(v13) = *(uint16_t*)(v14 + 24);
				v15 = *(unsigned short*)(v14 + 24);
				v62 = nox_tile_defs_arr[v15].data_32[*(uint32_t*)(v14 + 28) + nox_tile_defs_arr[v15].field_46];
				v68.field_0 = v76;
				v68.field_4 = v11 + 23;
				func_587000_154940(&v68, v62, v13);
				*getMemU32Ptr(0x85B3FC, 228 + 4 * v15) = 1;
				if (*(uint32_t*)(v14 + 40)) {
					nox_xxx_tileCallDrawEdges_481BF0((int)&v68, *(uint32_t*)(v14 + 40));
				}
			}
			if (*(uint8_t*)v14 & 1) {
				LOWORD(v13) = *(uint16_t*)(v14 + 4);
				v16 = *(unsigned short*)(v14 + 4);
				v63 = nox_tile_defs_arr[v16].data_32[*(uint32_t*)(v14 + 8) + nox_tile_defs_arr[v16].field_46];
				v68.field_0 = v76 + 23;
				v68.field_4 = v11;
				func_587000_154940(&v68, v63, v13);
				*getMemU32Ptr(0x85B3FC, 228 + 4 * v16) = 1;
				v17 = *(uint32_t*)(v14 + 20);
				if (v17) {
					nox_xxx_tileCallDrawEdges_481BF0((int)&v68, v17);
				}
			}
			v11 += 46;
			bool v18 = (int)++v74 < *(int*)&dword_5d4594_3798832 + dword_5d4594_3798816;
			i += 44;
			if (!v18) {
				break;
			}
		}
	}
}

void nox_xxx_tileDrawMB_481C20_B(nox_draw_viewport_t* vp, int v78) {
	int v33;  // esi
	int v65;  // [esp-8h] [ebp-54h]
	int v32;  // edi
	int2 v68; // [esp+1Ch] [ebp-30h]
	int v64;  // [esp-8h] [ebp-54h]
	int v31;  // edi
	int v28;  // esi
	char v29; // al
	int v30;  // esi
	int v27;  // ecx
	int v25;  // eax
	int v26;  // ebx
	int v21;  // eax
	int v20;  // edi
	int v19;  // esi
	int v23;  // edi
	int v24;  // ecx
	int v22;  // esi
	int v76;
	if ((int)v78 >= *(int*)&dword_5d4594_3798824 + 23) {
		if ((int)v78 + nox_getBackbufHeight() <= *(int*)&dword_5d4594_3798824 + *(int*)&dword_5d4594_3798808) {
			return;
		}
		v22 = dword_5d4594_3798832;
		if (*(int*)&dword_5d4594_3798832 + *(int*)&dword_5d4594_3798816 >= 128) {
			return;
		}
		if ((int)v78 + nox_getBackbufHeight() > *(int*)&dword_5d4594_3798824 + *(int*)&dword_5d4594_3798808 + 46) {
			nox_xxx_tileDrawImpl_4826A0(vp);
			return;
		}
		++dword_5d4594_3798832;
		v23 = dword_5d4594_3798824 + 46;
		v19 = dword_5d4594_3798816 + v22;
		v24 = dword_5d4594_3798840 + 46;
		dword_5d4594_3798824 += 46;
		dword_5d4594_3798840 += 46;
		if (*(int*)&dword_5d4594_3798840 >= *(int*)&dword_5d4594_3798808) {
			dword_5d4594_3798840 = v24 - dword_5d4594_3798808;
		}
		v76 = v23 + dword_5d4594_3798808 - 46;
	} else {
		if (*(int*)&dword_5d4594_3798832 <= 0) {
			return;
		}
		if ((int)v78 < *(int*)&dword_5d4594_3798824 - 23) {
			nox_xxx_tileDrawImpl_4826A0(vp);
			return;
		}
		v19 = dword_5d4594_3798832 - 1;
		v20 = dword_5d4594_3798824 - 46;
		v21 = dword_5d4594_3798840 - 46;
		bool v5 = *(int*)&dword_5d4594_3798840 - 46 < 0;
		--dword_5d4594_3798832;
		dword_5d4594_3798824 -= 46;
		dword_5d4594_3798840 -= 46;
		if (v5) {
			dword_5d4594_3798840 = dword_5d4594_3798808 + v21;
		}
		v76 = v20;
	}
	sub_481410();
	v25 = dword_5d4594_3798828;
	v26 = dword_5d4594_3798820;
	int i = dword_5d4594_3798828;
	int j = dword_5d4594_3798812 + dword_5d4594_3798828 - 1;
	if (dword_5d4594_3798828 < j) {
		v27 = 44 * v19;
		int v71 = 44 * v19;
		while (1) {
			v28 = ptr_5D4594_2650668[v25];
			v29 = *(uint8_t*)(v28 + v27);
			v30 = v27 + v28;
			if (v29 & 2) {
				LOWORD(v27) = *(uint16_t*)(v30 + 24);
				v31 = *(unsigned short*)(v30 + 24);
				v64 = nox_tile_defs_arr[v31].data_32[*(uint32_t*)(v30 + 28) + nox_tile_defs_arr[v31].field_46];
				v68.field_0 = v26;
				v68.field_4 = v76 + 23;
				func_587000_154940(&v68, v64, v27);
				*getMemU32Ptr(0x85B3FC, 228 + 4 * v31) = 1;
				if (*(uint32_t*)(v30 + 40)) {
					nox_xxx_tileCallDrawEdges_481BF0((int)&v68, *(uint32_t*)(v30 + 40));
				}
			}
			if (*(uint8_t*)v30 & 1) {
				LOWORD(v27) = *(uint16_t*)(v30 + 4);
				v32 = *(unsigned short*)(v30 + 4);
				v65 = nox_tile_defs_arr[v32].data_32[*(uint32_t*)(v30 + 8) + nox_tile_defs_arr[v32].field_46];
				v68.field_0 = v26 + 23;
				v68.field_4 = v76;
				func_587000_154940(&v68, v65, v27);
				*getMemU32Ptr(0x85B3FC, 228 + 4 * v32) = 1;
				v33 = *(uint32_t*)(v30 + 20);
				if (v33) {
					nox_xxx_tileCallDrawEdges_481BF0((int)&v68, v33);
				}
			}
			v26 += 46;
			if (++i >= j) {
				break;
			}
			v27 = v71;
			v25 = i;
		}
	}
}

//----- (00482570) --------------------------------------------------------
int nox_xxx_tileCheckRedrawMB_482570(nox_draw_viewport_t* vp) {
	uint32_t* a1 = vp;
	int v1;       // esi
	int v2;       // ebx
	int v3;       // eax
	int v4;       // edx
	int i;        // ebp
	int v6;       // edi
	uint32_t* v7; // esi
	int v8;       // eax
	int v10;      // [esp+10h] [ebp-8h]
	int v11;      // [esp+14h] [ebp-4h]
	int v12;      // [esp+1Ch] [ebp+4h]

	v1 = a1[5] - a1[1];
	v2 = (a1[4] - *a1 - 11) / 46;
	if (v2 < 0) {
		v2 = 0;
	}
	v12 = dword_5d4594_3798812 + v2 - 1;
	if (v12 >= 128) {
		v12 = 127;
		v2 = 127 - dword_5d4594_3798812;
	}
	v3 = (v1 - 11) / 46 - 1;
	if (v3 < 0) {
		v3 = 0;
	}
	v4 = dword_5d4594_3798816 + v3;
	v10 = dword_5d4594_3798816 + v3;
	if (*(int*)&dword_5d4594_3798816 + v3 >= 128) {
		v10 = 127;
		v3 = 127 - dword_5d4594_3798816;
		v4 = 127;
	}
	v11 = v3;
	if (v3 >= v4) {
		return 0;
	}
	for (i = 44 * v3; 1; i += 44) {
		v6 = v2;
		if (v2 < v12) {
			v7 = &ptr_5D4594_2650668[v2];
			while (1) {
				v8 = i + *v7;
				if (*(uint8_t*)v8 & 1) {
					if ((nox_tile_defs_arr[*(uint32_t*)(v8 + 4)].field_58 & 1) == 1) {
						return 1;
					}
				}
				if (*(uint8_t*)v8 & 2 && (nox_tile_defs_arr[*(uint32_t*)(v8 + 24)].field_58 & 1) == 1) {
					return 1;
				}
				++v6;
				++v7;
				if (v6 >= v12) {
					v4 = v10;
					goto LABEL_19;
				}
			}
		}
	LABEL_19:
		if (++v11 >= v4) {
			return 0;
		}
	}
}

//----- (004826A0) --------------------------------------------------------
int nox_xxx_tileDrawImpl_4826A0(nox_draw_viewport_t* vp) {
	uint32_t* a1 = vp;
	int v1;     // esi
	int v2;     // ebx
	int result; // eax
	int v4;     // ecx
	int v5;     // esi
	int v6;     // ebx
	int v7;     // ecx
	int v8;     // ebp
	int v9;     // eax
	int v10;    // ecx
	int v11;    // esi
	char v12;   // al
	int v13;    // esi
	int v14;    // edi
	int v15;    // edi
	int v16;    // esi
	bool v17;   // zf
	int v18;    // [esp-8h] [ebp-34h]
	int v19;    // [esp-8h] [ebp-34h]
	int v20;    // [esp+10h] [ebp-1Ch]
	int v21;    // [esp+14h] [ebp-18h]
	int v22;    // [esp+18h] [ebp-14h]
	int v23;    // [esp+1Ch] [ebp-10h]
	int2 v24;   // [esp+24h] [ebp-8h]
	int v25;    // [esp+30h] [ebp+4h]

	v1 = a1[4] - *a1;
	v2 = a1[5] - a1[1];
	dword_5d4594_3798836 = 0;
	dword_5d4594_3798840 = 0;
	sub_481410();
	memset(getMemAt(0x85B3FC, 228), 0, 0x2C0u);
	memset(getMemAt(0x5D4594, 2523980), 0, 0x100u);
	v25 = (v1 - 11) / 46;
	if (v25 < 0) {
		v25 = 0;
	}
	v20 = dword_5d4594_3798812 + v25 - 1;
	if (v20 >= 128) {
		v20 = 127;
		v25 = 127 - dword_5d4594_3798812;
	}
	result = (v2 - 11) / 46 - 1;
	if (result < 0) {
		result = 0;
	}
	v4 = dword_5d4594_3798816 + result;
	if (*(int*)&dword_5d4594_3798816 + result >= 128) {
		v4 = 127;
		result = 127 - dword_5d4594_3798816;
	}
	v5 = v25;
	dword_5d4594_3798824 = 46 * result - 11;
	v6 = 46 * result - 57;
	dword_5d4594_3798828 = v25;
	dword_5d4594_3798820 = 46 * v25 - 11;
	dword_5d4594_3798832 = result;
	if (result < v4) {
		v21 = 44 * result;
		v23 = v4 - result;
		v7 = v20;
		do {
			v8 = 46 * v25 - 57;
			v6 += 46;
			v9 = v5;
			v22 = v5;
			if (v5 < v7) {
				do {
					HIWORD(v10) = HIWORD(ptr_5D4594_2650668); // TODO: why it's doing it?
					v8 += 46;
					v11 = ptr_5D4594_2650668[v9];
					v12 = *(uint8_t*)(v11 + v21);
					v13 = v21 + v11;
					if (v12) {
						if (v12 & 2) {
							LOWORD(v10) = *(uint16_t*)(v13 + 24);
							v14 = *(unsigned short*)(v13 + 24);
							v18 = nox_tile_defs_arr[v14].data_32[*(uint32_t*)(v13 + 28) + nox_tile_defs_arr[v14].field_46];
							v24.field_0 = v8;
							v24.field_4 = v6 + 23;
							func_587000_154940(&v24, v18, v10);
							*getMemU32Ptr(0x85B3FC, 228 + 4 * v14) = 1;
							if (*(uint32_t*)(v13 + 40)) {
								nox_xxx_tileCallDrawEdges_481BF0((int)&v24, *(uint32_t*)(v13 + 40));
							}
						}
						if (*(uint8_t*)v13 & 1) {
							LOWORD(v10) = *(uint16_t*)(v13 + 4);
							v15 = *(unsigned short*)(v13 + 4);
							v19 = nox_tile_defs_arr[v15].data_32[*(uint32_t*)(v13 + 8) + nox_tile_defs_arr[v15].field_46];
							v24.field_0 = v8 + 23;
							v24.field_4 = v6;
							func_587000_154940(&v24, v19, v10);
							*getMemU32Ptr(0x85B3FC, 228 + 4 * v15) = 1;
							v16 = *(uint32_t*)(v13 + 20);
							if (v16) {
								nox_xxx_tileCallDrawEdges_481BF0((int)&v24, v16);
							}
						}
					}
					v7 = v20;
					v9 = ++v22;
				} while (v22 < v20);
				v5 = v25;
			}
			result = v23 - 1;
			v17 = v23 == 1;
			v21 += 44;
			--v23;
		} while (!v17);
	}
	return result;
}
// 4828A6: variable 'v10' is possibly undefined

//----- (004831C0) --------------------------------------------------------
short sub_4831C0(int a1, int a2) {
	int v2;       // edx
	int v3;       // edi
	int v4;       // edi
	int v5;       // edi
	int v6;       // edi
	int v7;       // edi
	int v8;       // edi
	int v9;       // edi
	int v10;      // edi
	int v11;      // edi
	int v12;      // edi
	int v13;      // edi
	int v14;      // edi
	int v15;      // edi
	int v16;      // edi
	int v17;      // edi
	int v18;      // edi
	int v19;      // edi
	int v20;      // edi
	int v21;      // edi
	int v22;      // edi
	int v23;      // edi
	int v24;      // edi
	int v25;      // edi
	int v26;      // edi
	int v27;      // edi
	int v28;      // edi
	int v29;      // edi
	int v30;      // edi
	int v31;      // edi
	int v32;      // edi
	int v33;      // edi
	int v34;      // edi
	int v35;      // edi
	int v36;      // edi
	int v37;      // edi
	int v38;      // edi
	int v39;      // edi
	int v40;      // edi
	int v41;      // edi
	int v42;      // edi
	int v43;      // edi
	int v44;      // edi
	int v45;      // edi
	int v46;      // edi
	short result; // ax

	v2 = dword_5d4594_3798804;
	*(uint16_t*)(a2 + 46) = *(uint16_t*)a1;
	v3 = v2 + a2;
	*(uint32_t*)(v3 + 44) = *(uint32_t*)(a1 + 2);
	*(uint16_t*)(v3 + 48) = *(uint16_t*)(a1 + 6);
	v4 = v2 + v2 + a2;
	*(uint16_t*)(v4 + 42) = *(uint16_t*)(a1 + 8);
	*(uint32_t*)(v4 + 44) = *(uint32_t*)(a1 + 10);
	*(uint32_t*)(v4 + 48) = *(uint32_t*)(a1 + 14);
	v5 = v2 + v4;
	*(uint32_t*)(v5 + 40) = *(uint32_t*)(a1 + 18);
	*(uint32_t*)(v5 + 44) = *(uint32_t*)(a1 + 22);
	*(uint32_t*)(v5 + 48) = *(uint32_t*)(a1 + 26);
	*(uint16_t*)(v5 + 52) = *(uint16_t*)(a1 + 30);
	v6 = v2 + v5;
	*(uint16_t*)(v6 + 38) = *(uint16_t*)(a1 + 32);
	*(uint32_t*)(v6 + 40) = *(uint32_t*)(a1 + 34);
	*(uint32_t*)(v6 + 44) = *(uint32_t*)(a1 + 38);
	*(uint32_t*)(v6 + 48) = *(uint32_t*)(a1 + 42);
	*(uint32_t*)(v6 + 52) = *(uint32_t*)(a1 + 46);
	v7 = v2 + v6;
	*(uint32_t*)(v7 + 36) = *(uint32_t*)(a1 + 50);
	*(uint32_t*)(v7 + 40) = *(uint32_t*)(a1 + 54);
	*(uint32_t*)(v7 + 44) = *(uint32_t*)(a1 + 58);
	*(uint32_t*)(v7 + 48) = *(uint32_t*)(a1 + 62);
	*(uint32_t*)(v7 + 52) = *(uint32_t*)(a1 + 66);
	*(uint16_t*)(v7 + 56) = *(uint16_t*)(a1 + 70);
	v8 = v2 + v7;
	*(uint16_t*)(v8 + 34) = *(uint16_t*)(a1 + 72);
	*(uint32_t*)(v8 + 36) = *(uint32_t*)(a1 + 74);
	*(uint32_t*)(v8 + 40) = *(uint32_t*)(a1 + 78);
	*(uint32_t*)(v8 + 44) = *(uint32_t*)(a1 + 82);
	*(uint32_t*)(v8 + 48) = *(uint32_t*)(a1 + 86);
	*(uint32_t*)(v8 + 52) = *(uint32_t*)(a1 + 90);
	*(uint32_t*)(v8 + 56) = *(uint32_t*)(a1 + 94);
	v9 = v2 + v8;
	*(uint32_t*)(v9 + 32) = *(uint32_t*)(a1 + 98);
	*(uint32_t*)(v9 + 36) = *(uint32_t*)(a1 + 102);
	*(uint32_t*)(v9 + 40) = *(uint32_t*)(a1 + 106);
	*(uint32_t*)(v9 + 44) = *(uint32_t*)(a1 + 110);
	*(uint32_t*)(v9 + 48) = *(uint32_t*)(a1 + 114);
	*(uint32_t*)(v9 + 52) = *(uint32_t*)(a1 + 118);
	*(uint32_t*)(v9 + 56) = *(uint32_t*)(a1 + 122);
	*(uint16_t*)(v9 + 60) = *(uint16_t*)(a1 + 126);
	v10 = v2 + v9;
	*(uint16_t*)(v10 + 30) = *(uint16_t*)(a1 + 128);
	*(uint32_t*)(v10 + 32) = *(uint32_t*)(a1 + 130);
	*(uint32_t*)(v10 + 36) = *(uint32_t*)(a1 + 134);
	*(uint32_t*)(v10 + 40) = *(uint32_t*)(a1 + 138);
	*(uint32_t*)(v10 + 44) = *(uint32_t*)(a1 + 142);
	*(uint32_t*)(v10 + 48) = *(uint32_t*)(a1 + 146);
	*(uint32_t*)(v10 + 52) = *(uint32_t*)(a1 + 150);
	*(uint32_t*)(v10 + 56) = *(uint32_t*)(a1 + 154);
	*(uint32_t*)(v10 + 60) = *(uint32_t*)(a1 + 158);
	v11 = v2 + v10;
	*(uint32_t*)(v11 + 28) = *(uint32_t*)(a1 + 162);
	*(uint32_t*)(v11 + 32) = *(uint32_t*)(a1 + 166);
	*(uint32_t*)(v11 + 36) = *(uint32_t*)(a1 + 170);
	*(uint32_t*)(v11 + 40) = *(uint32_t*)(a1 + 174);
	*(uint32_t*)(v11 + 44) = *(uint32_t*)(a1 + 178);
	*(uint32_t*)(v11 + 48) = *(uint32_t*)(a1 + 182);
	*(uint32_t*)(v11 + 52) = *(uint32_t*)(a1 + 186);
	*(uint32_t*)(v11 + 56) = *(uint32_t*)(a1 + 190);
	*(uint32_t*)(v11 + 60) = *(uint32_t*)(a1 + 194);
	*(uint16_t*)(v11 + 64) = *(uint16_t*)(a1 + 198);
	v12 = v2 + v11;
	*(uint16_t*)(v12 + 26) = *(uint16_t*)(a1 + 200);
	*(uint32_t*)(v12 + 28) = *(uint32_t*)(a1 + 202);
	*(uint32_t*)(v12 + 32) = *(uint32_t*)(a1 + 206);
	*(uint32_t*)(v12 + 36) = *(uint32_t*)(a1 + 210);
	*(uint32_t*)(v12 + 40) = *(uint32_t*)(a1 + 214);
	*(uint32_t*)(v12 + 44) = *(uint32_t*)(a1 + 218);
	*(uint32_t*)(v12 + 48) = *(uint32_t*)(a1 + 222);
	*(uint32_t*)(v12 + 52) = *(uint32_t*)(a1 + 226);
	*(uint32_t*)(v12 + 56) = *(uint32_t*)(a1 + 230);
	*(uint32_t*)(v12 + 60) = *(uint32_t*)(a1 + 234);
	*(uint32_t*)(v12 + 64) = *(uint32_t*)(a1 + 238);
	v13 = v2 + v12;
	*(uint32_t*)(v13 + 24) = *(uint32_t*)(a1 + 242);
	*(uint32_t*)(v13 + 28) = *(uint32_t*)(a1 + 246);
	*(uint32_t*)(v13 + 32) = *(uint32_t*)(a1 + 250);
	*(uint32_t*)(v13 + 36) = *(uint32_t*)(a1 + 254);
	*(uint32_t*)(v13 + 40) = *(uint32_t*)(a1 + 258);
	*(uint32_t*)(v13 + 44) = *(uint32_t*)(a1 + 262);
	*(uint32_t*)(v13 + 48) = *(uint32_t*)(a1 + 266);
	*(uint32_t*)(v13 + 52) = *(uint32_t*)(a1 + 270);
	*(uint32_t*)(v13 + 56) = *(uint32_t*)(a1 + 274);
	*(uint32_t*)(v13 + 60) = *(uint32_t*)(a1 + 278);
	*(uint32_t*)(v13 + 64) = *(uint32_t*)(a1 + 282);
	*(uint16_t*)(v13 + 68) = *(uint16_t*)(a1 + 286);
	v14 = v2 + v13;
	*(uint16_t*)(v14 + 22) = *(uint16_t*)(a1 + 288);
	*(uint32_t*)(v14 + 24) = *(uint32_t*)(a1 + 290);
	*(uint32_t*)(v14 + 28) = *(uint32_t*)(a1 + 294);
	*(uint32_t*)(v14 + 32) = *(uint32_t*)(a1 + 298);
	*(uint32_t*)(v14 + 36) = *(uint32_t*)(a1 + 302);
	*(uint32_t*)(v14 + 40) = *(uint32_t*)(a1 + 306);
	*(uint32_t*)(v14 + 44) = *(uint32_t*)(a1 + 310);
	*(uint32_t*)(v14 + 48) = *(uint32_t*)(a1 + 314);
	*(uint32_t*)(v14 + 52) = *(uint32_t*)(a1 + 318);
	*(uint32_t*)(v14 + 56) = *(uint32_t*)(a1 + 322);
	*(uint32_t*)(v14 + 60) = *(uint32_t*)(a1 + 326);
	*(uint32_t*)(v14 + 64) = *(uint32_t*)(a1 + 330);
	*(uint32_t*)(v14 + 68) = *(uint32_t*)(a1 + 334);
	v15 = v2 + v14;
	*(uint32_t*)(v15 + 20) = *(uint32_t*)(a1 + 338);
	*(uint32_t*)(v15 + 24) = *(uint32_t*)(a1 + 342);
	*(uint32_t*)(v15 + 28) = *(uint32_t*)(a1 + 346);
	*(uint32_t*)(v15 + 32) = *(uint32_t*)(a1 + 350);
	*(uint32_t*)(v15 + 36) = *(uint32_t*)(a1 + 354);
	*(uint32_t*)(v15 + 40) = *(uint32_t*)(a1 + 358);
	*(uint32_t*)(v15 + 44) = *(uint32_t*)(a1 + 362);
	*(uint32_t*)(v15 + 48) = *(uint32_t*)(a1 + 366);
	*(uint32_t*)(v15 + 52) = *(uint32_t*)(a1 + 370);
	*(uint32_t*)(v15 + 56) = *(uint32_t*)(a1 + 374);
	*(uint32_t*)(v15 + 60) = *(uint32_t*)(a1 + 378);
	*(uint32_t*)(v15 + 64) = *(uint32_t*)(a1 + 382);
	*(uint32_t*)(v15 + 68) = *(uint32_t*)(a1 + 386);
	*(uint16_t*)(v15 + 72) = *(uint16_t*)(a1 + 390);
	v16 = v2 + v15;
	*(uint16_t*)(v16 + 18) = *(uint16_t*)(a1 + 392);
	*(uint32_t*)(v16 + 20) = *(uint32_t*)(a1 + 394);
	*(uint32_t*)(v16 + 24) = *(uint32_t*)(a1 + 398);
	*(uint32_t*)(v16 + 28) = *(uint32_t*)(a1 + 402);
	*(uint32_t*)(v16 + 32) = *(uint32_t*)(a1 + 406);
	*(uint32_t*)(v16 + 36) = *(uint32_t*)(a1 + 410);
	*(uint32_t*)(v16 + 40) = *(uint32_t*)(a1 + 414);
	*(uint32_t*)(v16 + 44) = *(uint32_t*)(a1 + 418);
	*(uint32_t*)(v16 + 48) = *(uint32_t*)(a1 + 422);
	*(uint32_t*)(v16 + 52) = *(uint32_t*)(a1 + 426);
	*(uint32_t*)(v16 + 56) = *(uint32_t*)(a1 + 430);
	*(uint32_t*)(v16 + 60) = *(uint32_t*)(a1 + 434);
	*(uint32_t*)(v16 + 64) = *(uint32_t*)(a1 + 438);
	*(uint32_t*)(v16 + 68) = *(uint32_t*)(a1 + 442);
	*(uint32_t*)(v16 + 72) = *(uint32_t*)(a1 + 446);
	v17 = v2 + v16;
	*(uint32_t*)(v17 + 16) = *(uint32_t*)(a1 + 450);
	*(uint32_t*)(v17 + 20) = *(uint32_t*)(a1 + 454);
	*(uint32_t*)(v17 + 24) = *(uint32_t*)(a1 + 458);
	*(uint32_t*)(v17 + 28) = *(uint32_t*)(a1 + 462);
	*(uint32_t*)(v17 + 32) = *(uint32_t*)(a1 + 466);
	*(uint32_t*)(v17 + 36) = *(uint32_t*)(a1 + 470);
	*(uint32_t*)(v17 + 40) = *(uint32_t*)(a1 + 474);
	*(uint32_t*)(v17 + 44) = *(uint32_t*)(a1 + 478);
	*(uint32_t*)(v17 + 48) = *(uint32_t*)(a1 + 482);
	*(uint32_t*)(v17 + 52) = *(uint32_t*)(a1 + 486);
	*(uint32_t*)(v17 + 56) = *(uint32_t*)(a1 + 490);
	*(uint32_t*)(v17 + 60) = *(uint32_t*)(a1 + 494);
	*(uint32_t*)(v17 + 64) = *(uint32_t*)(a1 + 498);
	*(uint32_t*)(v17 + 68) = *(uint32_t*)(a1 + 502);
	*(uint32_t*)(v17 + 72) = *(uint32_t*)(a1 + 506);
	*(uint16_t*)(v17 + 76) = *(uint16_t*)(a1 + 510);
	v18 = v2 + v17;
	*(uint16_t*)(v18 + 14) = *(uint16_t*)(a1 + 512);
	*(uint32_t*)(v18 + 16) = *(uint32_t*)(a1 + 514);
	*(uint32_t*)(v18 + 20) = *(uint32_t*)(a1 + 518);
	*(uint32_t*)(v18 + 24) = *(uint32_t*)(a1 + 522);
	*(uint32_t*)(v18 + 28) = *(uint32_t*)(a1 + 526);
	*(uint32_t*)(v18 + 32) = *(uint32_t*)(a1 + 530);
	*(uint32_t*)(v18 + 36) = *(uint32_t*)(a1 + 534);
	*(uint32_t*)(v18 + 40) = *(uint32_t*)(a1 + 538);
	*(uint32_t*)(v18 + 44) = *(uint32_t*)(a1 + 542);
	*(uint32_t*)(v18 + 48) = *(uint32_t*)(a1 + 546);
	*(uint32_t*)(v18 + 52) = *(uint32_t*)(a1 + 550);
	*(uint32_t*)(v18 + 56) = *(uint32_t*)(a1 + 554);
	*(uint32_t*)(v18 + 60) = *(uint32_t*)(a1 + 558);
	*(uint32_t*)(v18 + 64) = *(uint32_t*)(a1 + 562);
	*(uint32_t*)(v18 + 68) = *(uint32_t*)(a1 + 566);
	*(uint32_t*)(v18 + 72) = *(uint32_t*)(a1 + 570);
	*(uint32_t*)(v18 + 76) = *(uint32_t*)(a1 + 574);
	v19 = v2 + v18;
	*(uint32_t*)(v19 + 12) = *(uint32_t*)(a1 + 578);
	*(uint32_t*)(v19 + 16) = *(uint32_t*)(a1 + 582);
	*(uint32_t*)(v19 + 20) = *(uint32_t*)(a1 + 586);
	*(uint32_t*)(v19 + 24) = *(uint32_t*)(a1 + 590);
	*(uint32_t*)(v19 + 28) = *(uint32_t*)(a1 + 594);
	*(uint32_t*)(v19 + 32) = *(uint32_t*)(a1 + 598);
	*(uint32_t*)(v19 + 36) = *(uint32_t*)(a1 + 602);
	*(uint32_t*)(v19 + 40) = *(uint32_t*)(a1 + 606);
	*(uint32_t*)(v19 + 44) = *(uint32_t*)(a1 + 610);
	*(uint32_t*)(v19 + 48) = *(uint32_t*)(a1 + 614);
	*(uint32_t*)(v19 + 52) = *(uint32_t*)(a1 + 618);
	*(uint32_t*)(v19 + 56) = *(uint32_t*)(a1 + 622);
	*(uint32_t*)(v19 + 60) = *(uint32_t*)(a1 + 626);
	*(uint32_t*)(v19 + 64) = *(uint32_t*)(a1 + 630);
	*(uint32_t*)(v19 + 68) = *(uint32_t*)(a1 + 634);
	*(uint32_t*)(v19 + 72) = *(uint32_t*)(a1 + 638);
	*(uint32_t*)(v19 + 76) = *(uint32_t*)(a1 + 642);
	*(uint16_t*)(v19 + 80) = *(uint16_t*)(a1 + 646);
	v20 = v2 + v19;
	*(uint16_t*)(v20 + 10) = *(uint16_t*)(a1 + 648);
	*(uint32_t*)(v20 + 12) = *(uint32_t*)(a1 + 650);
	*(uint32_t*)(v20 + 16) = *(uint32_t*)(a1 + 654);
	*(uint32_t*)(v20 + 20) = *(uint32_t*)(a1 + 658);
	*(uint32_t*)(v20 + 24) = *(uint32_t*)(a1 + 662);
	*(uint32_t*)(v20 + 28) = *(uint32_t*)(a1 + 666);
	*(uint32_t*)(v20 + 32) = *(uint32_t*)(a1 + 670);
	*(uint32_t*)(v20 + 36) = *(uint32_t*)(a1 + 674);
	*(uint32_t*)(v20 + 40) = *(uint32_t*)(a1 + 678);
	*(uint32_t*)(v20 + 44) = *(uint32_t*)(a1 + 682);
	*(uint32_t*)(v20 + 48) = *(uint32_t*)(a1 + 686);
	*(uint32_t*)(v20 + 52) = *(uint32_t*)(a1 + 690);
	*(uint32_t*)(v20 + 56) = *(uint32_t*)(a1 + 694);
	*(uint32_t*)(v20 + 60) = *(uint32_t*)(a1 + 698);
	*(uint32_t*)(v20 + 64) = *(uint32_t*)(a1 + 702);
	*(uint32_t*)(v20 + 68) = *(uint32_t*)(a1 + 706);
	*(uint32_t*)(v20 + 72) = *(uint32_t*)(a1 + 710);
	*(uint32_t*)(v20 + 76) = *(uint32_t*)(a1 + 714);
	*(uint32_t*)(v20 + 80) = *(uint32_t*)(a1 + 718);
	v21 = v2 + v20;
	*(uint32_t*)(v21 + 8) = *(uint32_t*)(a1 + 722);
	*(uint32_t*)(v21 + 12) = *(uint32_t*)(a1 + 726);
	*(uint32_t*)(v21 + 16) = *(uint32_t*)(a1 + 730);
	*(uint32_t*)(v21 + 20) = *(uint32_t*)(a1 + 734);
	*(uint32_t*)(v21 + 24) = *(uint32_t*)(a1 + 738);
	*(uint32_t*)(v21 + 28) = *(uint32_t*)(a1 + 742);
	*(uint32_t*)(v21 + 32) = *(uint32_t*)(a1 + 746);
	*(uint32_t*)(v21 + 36) = *(uint32_t*)(a1 + 750);
	*(uint32_t*)(v21 + 40) = *(uint32_t*)(a1 + 754);
	*(uint32_t*)(v21 + 44) = *(uint32_t*)(a1 + 758);
	*(uint32_t*)(v21 + 48) = *(uint32_t*)(a1 + 762);
	*(uint32_t*)(v21 + 52) = *(uint32_t*)(a1 + 766);
	*(uint32_t*)(v21 + 56) = *(uint32_t*)(a1 + 770);
	*(uint32_t*)(v21 + 60) = *(uint32_t*)(a1 + 774);
	*(uint32_t*)(v21 + 64) = *(uint32_t*)(a1 + 778);
	*(uint32_t*)(v21 + 68) = *(uint32_t*)(a1 + 782);
	*(uint32_t*)(v21 + 72) = *(uint32_t*)(a1 + 786);
	*(uint32_t*)(v21 + 76) = *(uint32_t*)(a1 + 790);
	*(uint32_t*)(v21 + 80) = *(uint32_t*)(a1 + 794);
	*(uint16_t*)(v21 + 84) = *(uint16_t*)(a1 + 798);
	v22 = v2 + v21;
	*(uint16_t*)(v22 + 6) = *(uint16_t*)(a1 + 800);
	*(uint32_t*)(v22 + 8) = *(uint32_t*)(a1 + 802);
	*(uint32_t*)(v22 + 12) = *(uint32_t*)(a1 + 806);
	*(uint32_t*)(v22 + 16) = *(uint32_t*)(a1 + 810);
	*(uint32_t*)(v22 + 20) = *(uint32_t*)(a1 + 814);
	*(uint32_t*)(v22 + 24) = *(uint32_t*)(a1 + 818);
	*(uint32_t*)(v22 + 28) = *(uint32_t*)(a1 + 822);
	*(uint32_t*)(v22 + 32) = *(uint32_t*)(a1 + 826);
	*(uint32_t*)(v22 + 36) = *(uint32_t*)(a1 + 830);
	*(uint32_t*)(v22 + 40) = *(uint32_t*)(a1 + 834);
	*(uint32_t*)(v22 + 44) = *(uint32_t*)(a1 + 838);
	*(uint32_t*)(v22 + 48) = *(uint32_t*)(a1 + 842);
	*(uint32_t*)(v22 + 52) = *(uint32_t*)(a1 + 846);
	*(uint32_t*)(v22 + 56) = *(uint32_t*)(a1 + 850);
	*(uint32_t*)(v22 + 60) = *(uint32_t*)(a1 + 854);
	*(uint32_t*)(v22 + 64) = *(uint32_t*)(a1 + 858);
	*(uint32_t*)(v22 + 68) = *(uint32_t*)(a1 + 862);
	*(uint32_t*)(v22 + 72) = *(uint32_t*)(a1 + 866);
	*(uint32_t*)(v22 + 76) = *(uint32_t*)(a1 + 870);
	*(uint32_t*)(v22 + 80) = *(uint32_t*)(a1 + 874);
	*(uint32_t*)(v22 + 84) = *(uint32_t*)(a1 + 878);
	v23 = v2 + v22;
	*(uint32_t*)(v23 + 4) = *(uint32_t*)(a1 + 882);
	*(uint32_t*)(v23 + 8) = *(uint32_t*)(a1 + 886);
	*(uint32_t*)(v23 + 12) = *(uint32_t*)(a1 + 890);
	*(uint32_t*)(v23 + 16) = *(uint32_t*)(a1 + 894);
	*(uint32_t*)(v23 + 20) = *(uint32_t*)(a1 + 898);
	*(uint32_t*)(v23 + 24) = *(uint32_t*)(a1 + 902);
	*(uint32_t*)(v23 + 28) = *(uint32_t*)(a1 + 906);
	*(uint32_t*)(v23 + 32) = *(uint32_t*)(a1 + 910);
	*(uint32_t*)(v23 + 36) = *(uint32_t*)(a1 + 914);
	*(uint32_t*)(v23 + 40) = *(uint32_t*)(a1 + 918);
	*(uint32_t*)(v23 + 44) = *(uint32_t*)(a1 + 922);
	*(uint32_t*)(v23 + 48) = *(uint32_t*)(a1 + 926);
	*(uint32_t*)(v23 + 52) = *(uint32_t*)(a1 + 930);
	*(uint32_t*)(v23 + 56) = *(uint32_t*)(a1 + 934);
	*(uint32_t*)(v23 + 60) = *(uint32_t*)(a1 + 938);
	*(uint32_t*)(v23 + 64) = *(uint32_t*)(a1 + 942);
	*(uint32_t*)(v23 + 68) = *(uint32_t*)(a1 + 946);
	*(uint32_t*)(v23 + 72) = *(uint32_t*)(a1 + 950);
	*(uint32_t*)(v23 + 76) = *(uint32_t*)(a1 + 954);
	*(uint32_t*)(v23 + 80) = *(uint32_t*)(a1 + 958);
	*(uint32_t*)(v23 + 84) = *(uint32_t*)(a1 + 962);
	*(uint16_t*)(v23 + 88) = *(uint16_t*)(a1 + 966);
	v24 = v2 + v23;
	*(uint16_t*)(v24 + 2) = *(uint16_t*)(a1 + 968);
	*(uint32_t*)(v24 + 4) = *(uint32_t*)(a1 + 970);
	*(uint32_t*)(v24 + 8) = *(uint32_t*)(a1 + 974);
	*(uint32_t*)(v24 + 12) = *(uint32_t*)(a1 + 978);
	*(uint32_t*)(v24 + 16) = *(uint32_t*)(a1 + 982);
	*(uint32_t*)(v24 + 20) = *(uint32_t*)(a1 + 986);
	*(uint32_t*)(v24 + 24) = *(uint32_t*)(a1 + 990);
	*(uint32_t*)(v24 + 28) = *(uint32_t*)(a1 + 994);
	*(uint32_t*)(v24 + 32) = *(uint32_t*)(a1 + 998);
	*(uint32_t*)(v24 + 36) = *(uint32_t*)(a1 + 1002);
	*(uint32_t*)(v24 + 40) = *(uint32_t*)(a1 + 1006);
	*(uint32_t*)(v24 + 44) = *(uint32_t*)(a1 + 1010);
	*(uint32_t*)(v24 + 48) = *(uint32_t*)(a1 + 1014);
	*(uint32_t*)(v24 + 52) = *(uint32_t*)(a1 + 1018);
	*(uint32_t*)(v24 + 56) = *(uint32_t*)(a1 + 1022);
	*(uint32_t*)(v24 + 60) = *(uint32_t*)(a1 + 1026);
	*(uint32_t*)(v24 + 64) = *(uint32_t*)(a1 + 1030);
	*(uint32_t*)(v24 + 68) = *(uint32_t*)(a1 + 1034);
	*(uint32_t*)(v24 + 72) = *(uint32_t*)(a1 + 1038);
	*(uint32_t*)(v24 + 76) = *(uint32_t*)(a1 + 1042);
	*(uint32_t*)(v24 + 80) = *(uint32_t*)(a1 + 1046);
	*(uint32_t*)(v24 + 84) = *(uint32_t*)(a1 + 1050);
	*(uint32_t*)(v24 + 88) = *(uint32_t*)(a1 + 1054);
	v25 = v2 + v24;
	*(uint16_t*)(v25 + 2) = *(uint16_t*)(a1 + 1058);
	*(uint32_t*)(v25 + 4) = *(uint32_t*)(a1 + 1060);
	*(uint32_t*)(v25 + 8) = *(uint32_t*)(a1 + 1064);
	*(uint32_t*)(v25 + 12) = *(uint32_t*)(a1 + 1068);
	*(uint32_t*)(v25 + 16) = *(uint32_t*)(a1 + 1072);
	*(uint32_t*)(v25 + 20) = *(uint32_t*)(a1 + 1076);
	*(uint32_t*)(v25 + 24) = *(uint32_t*)(a1 + 1080);
	*(uint32_t*)(v25 + 28) = *(uint32_t*)(a1 + 1084);
	*(uint32_t*)(v25 + 32) = *(uint32_t*)(a1 + 1088);
	*(uint32_t*)(v25 + 36) = *(uint32_t*)(a1 + 1092);
	*(uint32_t*)(v25 + 40) = *(uint32_t*)(a1 + 1096);
	*(uint32_t*)(v25 + 44) = *(uint32_t*)(a1 + 1100);
	*(uint32_t*)(v25 + 48) = *(uint32_t*)(a1 + 1104);
	*(uint32_t*)(v25 + 52) = *(uint32_t*)(a1 + 1108);
	*(uint32_t*)(v25 + 56) = *(uint32_t*)(a1 + 1112);
	*(uint32_t*)(v25 + 60) = *(uint32_t*)(a1 + 1116);
	*(uint32_t*)(v25 + 64) = *(uint32_t*)(a1 + 1120);
	*(uint32_t*)(v25 + 68) = *(uint32_t*)(a1 + 1124);
	*(uint32_t*)(v25 + 72) = *(uint32_t*)(a1 + 1128);
	*(uint32_t*)(v25 + 76) = *(uint32_t*)(a1 + 1132);
	*(uint32_t*)(v25 + 80) = *(uint32_t*)(a1 + 1136);
	*(uint32_t*)(v25 + 84) = *(uint32_t*)(a1 + 1140);
	*(uint32_t*)(v25 + 88) = *(uint32_t*)(a1 + 1144);
	v26 = v2 + v25;
	*(uint32_t*)(v26 + 4) = *(uint32_t*)(a1 + 1148);
	*(uint32_t*)(v26 + 8) = *(uint32_t*)(a1 + 1152);
	*(uint32_t*)(v26 + 12) = *(uint32_t*)(a1 + 1156);
	*(uint32_t*)(v26 + 16) = *(uint32_t*)(a1 + 1160);
	*(uint32_t*)(v26 + 20) = *(uint32_t*)(a1 + 1164);
	*(uint32_t*)(v26 + 24) = *(uint32_t*)(a1 + 1168);
	*(uint32_t*)(v26 + 28) = *(uint32_t*)(a1 + 1172);
	*(uint32_t*)(v26 + 32) = *(uint32_t*)(a1 + 1176);
	*(uint32_t*)(v26 + 36) = *(uint32_t*)(a1 + 1180);
	*(uint32_t*)(v26 + 40) = *(uint32_t*)(a1 + 1184);
	*(uint32_t*)(v26 + 44) = *(uint32_t*)(a1 + 1188);
	*(uint32_t*)(v26 + 48) = *(uint32_t*)(a1 + 1192);
	*(uint32_t*)(v26 + 52) = *(uint32_t*)(a1 + 1196);
	*(uint32_t*)(v26 + 56) = *(uint32_t*)(a1 + 1200);
	*(uint32_t*)(v26 + 60) = *(uint32_t*)(a1 + 1204);
	*(uint32_t*)(v26 + 64) = *(uint32_t*)(a1 + 1208);
	*(uint32_t*)(v26 + 68) = *(uint32_t*)(a1 + 1212);
	*(uint32_t*)(v26 + 72) = *(uint32_t*)(a1 + 1216);
	*(uint32_t*)(v26 + 76) = *(uint32_t*)(a1 + 1220);
	*(uint32_t*)(v26 + 80) = *(uint32_t*)(a1 + 1224);
	*(uint32_t*)(v26 + 84) = *(uint32_t*)(a1 + 1228);
	*(uint16_t*)(v26 + 88) = *(uint16_t*)(a1 + 1232);
	v27 = v2 + v26;
	*(uint16_t*)(v27 + 6) = *(uint16_t*)(a1 + 1234);
	*(uint32_t*)(v27 + 8) = *(uint32_t*)(a1 + 1236);
	*(uint32_t*)(v27 + 12) = *(uint32_t*)(a1 + 1240);
	*(uint32_t*)(v27 + 16) = *(uint32_t*)(a1 + 1244);
	*(uint32_t*)(v27 + 20) = *(uint32_t*)(a1 + 1248);
	*(uint32_t*)(v27 + 24) = *(uint32_t*)(a1 + 1252);
	*(uint32_t*)(v27 + 28) = *(uint32_t*)(a1 + 1256);
	*(uint32_t*)(v27 + 32) = *(uint32_t*)(a1 + 1260);
	*(uint32_t*)(v27 + 36) = *(uint32_t*)(a1 + 1264);
	*(uint32_t*)(v27 + 40) = *(uint32_t*)(a1 + 1268);
	*(uint32_t*)(v27 + 44) = *(uint32_t*)(a1 + 1272);
	*(uint32_t*)(v27 + 48) = *(uint32_t*)(a1 + 1276);
	*(uint32_t*)(v27 + 52) = *(uint32_t*)(a1 + 1280);
	*(uint32_t*)(v27 + 56) = *(uint32_t*)(a1 + 1284);
	*(uint32_t*)(v27 + 60) = *(uint32_t*)(a1 + 1288);
	*(uint32_t*)(v27 + 64) = *(uint32_t*)(a1 + 1292);
	*(uint32_t*)(v27 + 68) = *(uint32_t*)(a1 + 1296);
	*(uint32_t*)(v27 + 72) = *(uint32_t*)(a1 + 1300);
	*(uint32_t*)(v27 + 76) = *(uint32_t*)(a1 + 1304);
	*(uint32_t*)(v27 + 80) = *(uint32_t*)(a1 + 1308);
	*(uint32_t*)(v27 + 84) = *(uint32_t*)(a1 + 1312);
	v28 = v2 + v27;
	*(uint32_t*)(v28 + 8) = *(uint32_t*)(a1 + 1316);
	*(uint32_t*)(v28 + 12) = *(uint32_t*)(a1 + 1320);
	*(uint32_t*)(v28 + 16) = *(uint32_t*)(a1 + 1324);
	*(uint32_t*)(v28 + 20) = *(uint32_t*)(a1 + 1328);
	*(uint32_t*)(v28 + 24) = *(uint32_t*)(a1 + 1332);
	*(uint32_t*)(v28 + 28) = *(uint32_t*)(a1 + 1336);
	*(uint32_t*)(v28 + 32) = *(uint32_t*)(a1 + 1340);
	*(uint32_t*)(v28 + 36) = *(uint32_t*)(a1 + 1344);
	*(uint32_t*)(v28 + 40) = *(uint32_t*)(a1 + 1348);
	*(uint32_t*)(v28 + 44) = *(uint32_t*)(a1 + 1352);
	*(uint32_t*)(v28 + 48) = *(uint32_t*)(a1 + 1356);
	*(uint32_t*)(v28 + 52) = *(uint32_t*)(a1 + 1360);
	*(uint32_t*)(v28 + 56) = *(uint32_t*)(a1 + 1364);
	*(uint32_t*)(v28 + 60) = *(uint32_t*)(a1 + 1368);
	*(uint32_t*)(v28 + 64) = *(uint32_t*)(a1 + 1372);
	*(uint32_t*)(v28 + 68) = *(uint32_t*)(a1 + 1376);
	*(uint32_t*)(v28 + 72) = *(uint32_t*)(a1 + 1380);
	*(uint32_t*)(v28 + 76) = *(uint32_t*)(a1 + 1384);
	*(uint32_t*)(v28 + 80) = *(uint32_t*)(a1 + 1388);
	*(uint16_t*)(v28 + 84) = *(uint16_t*)(a1 + 1392);
	v29 = v2 + v28;
	*(uint16_t*)(v29 + 10) = *(uint16_t*)(a1 + 1394);
	*(uint32_t*)(v29 + 12) = *(uint32_t*)(a1 + 1396);
	*(uint32_t*)(v29 + 16) = *(uint32_t*)(a1 + 1400);
	*(uint32_t*)(v29 + 20) = *(uint32_t*)(a1 + 1404);
	*(uint32_t*)(v29 + 24) = *(uint32_t*)(a1 + 1408);
	*(uint32_t*)(v29 + 28) = *(uint32_t*)(a1 + 1412);
	*(uint32_t*)(v29 + 32) = *(uint32_t*)(a1 + 1416);
	*(uint32_t*)(v29 + 36) = *(uint32_t*)(a1 + 1420);
	*(uint32_t*)(v29 + 40) = *(uint32_t*)(a1 + 1424);
	*(uint32_t*)(v29 + 44) = *(uint32_t*)(a1 + 1428);
	*(uint32_t*)(v29 + 48) = *(uint32_t*)(a1 + 1432);
	*(uint32_t*)(v29 + 52) = *(uint32_t*)(a1 + 1436);
	*(uint32_t*)(v29 + 56) = *(uint32_t*)(a1 + 1440);
	*(uint32_t*)(v29 + 60) = *(uint32_t*)(a1 + 1444);
	*(uint32_t*)(v29 + 64) = *(uint32_t*)(a1 + 1448);
	*(uint32_t*)(v29 + 68) = *(uint32_t*)(a1 + 1452);
	*(uint32_t*)(v29 + 72) = *(uint32_t*)(a1 + 1456);
	*(uint32_t*)(v29 + 76) = *(uint32_t*)(a1 + 1460);
	*(uint32_t*)(v29 + 80) = *(uint32_t*)(a1 + 1464);
	v30 = v2 + v29;
	*(uint32_t*)(v30 + 12) = *(uint32_t*)(a1 + 1468);
	*(uint32_t*)(v30 + 16) = *(uint32_t*)(a1 + 1472);
	*(uint32_t*)(v30 + 20) = *(uint32_t*)(a1 + 1476);
	*(uint32_t*)(v30 + 24) = *(uint32_t*)(a1 + 1480);
	*(uint32_t*)(v30 + 28) = *(uint32_t*)(a1 + 1484);
	*(uint32_t*)(v30 + 32) = *(uint32_t*)(a1 + 1488);
	*(uint32_t*)(v30 + 36) = *(uint32_t*)(a1 + 1492);
	*(uint32_t*)(v30 + 40) = *(uint32_t*)(a1 + 1496);
	*(uint32_t*)(v30 + 44) = *(uint32_t*)(a1 + 1500);
	*(uint32_t*)(v30 + 48) = *(uint32_t*)(a1 + 1504);
	*(uint32_t*)(v30 + 52) = *(uint32_t*)(a1 + 1508);
	*(uint32_t*)(v30 + 56) = *(uint32_t*)(a1 + 1512);
	*(uint32_t*)(v30 + 60) = *(uint32_t*)(a1 + 1516);
	*(uint32_t*)(v30 + 64) = *(uint32_t*)(a1 + 1520);
	*(uint32_t*)(v30 + 68) = *(uint32_t*)(a1 + 1524);
	*(uint32_t*)(v30 + 72) = *(uint32_t*)(a1 + 1528);
	*(uint32_t*)(v30 + 76) = *(uint32_t*)(a1 + 1532);
	*(uint16_t*)(v30 + 80) = *(uint16_t*)(a1 + 1536);
	v31 = v2 + v30;
	*(uint16_t*)(v31 + 14) = *(uint16_t*)(a1 + 1538);
	*(uint32_t*)(v31 + 16) = *(uint32_t*)(a1 + 1540);
	*(uint32_t*)(v31 + 20) = *(uint32_t*)(a1 + 1544);
	*(uint32_t*)(v31 + 24) = *(uint32_t*)(a1 + 1548);
	*(uint32_t*)(v31 + 28) = *(uint32_t*)(a1 + 1552);
	*(uint32_t*)(v31 + 32) = *(uint32_t*)(a1 + 1556);
	*(uint32_t*)(v31 + 36) = *(uint32_t*)(a1 + 1560);
	*(uint32_t*)(v31 + 40) = *(uint32_t*)(a1 + 1564);
	*(uint32_t*)(v31 + 44) = *(uint32_t*)(a1 + 1568);
	*(uint32_t*)(v31 + 48) = *(uint32_t*)(a1 + 1572);
	*(uint32_t*)(v31 + 52) = *(uint32_t*)(a1 + 1576);
	*(uint32_t*)(v31 + 56) = *(uint32_t*)(a1 + 1580);
	*(uint32_t*)(v31 + 60) = *(uint32_t*)(a1 + 1584);
	*(uint32_t*)(v31 + 64) = *(uint32_t*)(a1 + 1588);
	*(uint32_t*)(v31 + 68) = *(uint32_t*)(a1 + 1592);
	*(uint32_t*)(v31 + 72) = *(uint32_t*)(a1 + 1596);
	*(uint32_t*)(v31 + 76) = *(uint32_t*)(a1 + 1600);
	v32 = v2 + v31;
	*(uint32_t*)(v32 + 16) = *(uint32_t*)(a1 + 1604);
	*(uint32_t*)(v32 + 20) = *(uint32_t*)(a1 + 1608);
	*(uint32_t*)(v32 + 24) = *(uint32_t*)(a1 + 1612);
	*(uint32_t*)(v32 + 28) = *(uint32_t*)(a1 + 1616);
	*(uint32_t*)(v32 + 32) = *(uint32_t*)(a1 + 1620);
	*(uint32_t*)(v32 + 36) = *(uint32_t*)(a1 + 1624);
	*(uint32_t*)(v32 + 40) = *(uint32_t*)(a1 + 1628);
	*(uint32_t*)(v32 + 44) = *(uint32_t*)(a1 + 1632);
	*(uint32_t*)(v32 + 48) = *(uint32_t*)(a1 + 1636);
	*(uint32_t*)(v32 + 52) = *(uint32_t*)(a1 + 1640);
	*(uint32_t*)(v32 + 56) = *(uint32_t*)(a1 + 1644);
	*(uint32_t*)(v32 + 60) = *(uint32_t*)(a1 + 1648);
	*(uint32_t*)(v32 + 64) = *(uint32_t*)(a1 + 1652);
	*(uint32_t*)(v32 + 68) = *(uint32_t*)(a1 + 1656);
	*(uint32_t*)(v32 + 72) = *(uint32_t*)(a1 + 1660);
	*(uint16_t*)(v32 + 76) = *(uint16_t*)(a1 + 1664);
	v33 = v2 + v32;
	*(uint16_t*)(v33 + 18) = *(uint16_t*)(a1 + 1666);
	*(uint32_t*)(v33 + 20) = *(uint32_t*)(a1 + 1668);
	*(uint32_t*)(v33 + 24) = *(uint32_t*)(a1 + 1672);
	*(uint32_t*)(v33 + 28) = *(uint32_t*)(a1 + 1676);
	*(uint32_t*)(v33 + 32) = *(uint32_t*)(a1 + 1680);
	*(uint32_t*)(v33 + 36) = *(uint32_t*)(a1 + 1684);
	*(uint32_t*)(v33 + 40) = *(uint32_t*)(a1 + 1688);
	*(uint32_t*)(v33 + 44) = *(uint32_t*)(a1 + 1692);
	*(uint32_t*)(v33 + 48) = *(uint32_t*)(a1 + 1696);
	*(uint32_t*)(v33 + 52) = *(uint32_t*)(a1 + 1700);
	*(uint32_t*)(v33 + 56) = *(uint32_t*)(a1 + 1704);
	*(uint32_t*)(v33 + 60) = *(uint32_t*)(a1 + 1708);
	*(uint32_t*)(v33 + 64) = *(uint32_t*)(a1 + 1712);
	*(uint32_t*)(v33 + 68) = *(uint32_t*)(a1 + 1716);
	*(uint32_t*)(v33 + 72) = *(uint32_t*)(a1 + 1720);
	v34 = v2 + v33;
	*(uint32_t*)(v34 + 20) = *(uint32_t*)(a1 + 1724);
	*(uint32_t*)(v34 + 24) = *(uint32_t*)(a1 + 1728);
	*(uint32_t*)(v34 + 28) = *(uint32_t*)(a1 + 1732);
	*(uint32_t*)(v34 + 32) = *(uint32_t*)(a1 + 1736);
	*(uint32_t*)(v34 + 36) = *(uint32_t*)(a1 + 1740);
	*(uint32_t*)(v34 + 40) = *(uint32_t*)(a1 + 1744);
	*(uint32_t*)(v34 + 44) = *(uint32_t*)(a1 + 1748);
	*(uint32_t*)(v34 + 48) = *(uint32_t*)(a1 + 1752);
	*(uint32_t*)(v34 + 52) = *(uint32_t*)(a1 + 1756);
	*(uint32_t*)(v34 + 56) = *(uint32_t*)(a1 + 1760);
	*(uint32_t*)(v34 + 60) = *(uint32_t*)(a1 + 1764);
	*(uint32_t*)(v34 + 64) = *(uint32_t*)(a1 + 1768);
	*(uint32_t*)(v34 + 68) = *(uint32_t*)(a1 + 1772);
	*(uint16_t*)(v34 + 72) = *(uint16_t*)(a1 + 1776);
	v35 = v2 + v34;
	*(uint16_t*)(v35 + 22) = *(uint16_t*)(a1 + 1778);
	*(uint32_t*)(v35 + 24) = *(uint32_t*)(a1 + 1780);
	*(uint32_t*)(v35 + 28) = *(uint32_t*)(a1 + 1784);
	*(uint32_t*)(v35 + 32) = *(uint32_t*)(a1 + 1788);
	*(uint32_t*)(v35 + 36) = *(uint32_t*)(a1 + 1792);
	*(uint32_t*)(v35 + 40) = *(uint32_t*)(a1 + 1796);
	*(uint32_t*)(v35 + 44) = *(uint32_t*)(a1 + 1800);
	*(uint32_t*)(v35 + 48) = *(uint32_t*)(a1 + 1804);
	*(uint32_t*)(v35 + 52) = *(uint32_t*)(a1 + 1808);
	*(uint32_t*)(v35 + 56) = *(uint32_t*)(a1 + 1812);
	*(uint32_t*)(v35 + 60) = *(uint32_t*)(a1 + 1816);
	*(uint32_t*)(v35 + 64) = *(uint32_t*)(a1 + 1820);
	*(uint32_t*)(v35 + 68) = *(uint32_t*)(a1 + 1824);
	v36 = v2 + v35;
	*(uint32_t*)(v36 + 24) = *(uint32_t*)(a1 + 1828);
	*(uint32_t*)(v36 + 28) = *(uint32_t*)(a1 + 1832);
	*(uint32_t*)(v36 + 32) = *(uint32_t*)(a1 + 1836);
	*(uint32_t*)(v36 + 36) = *(uint32_t*)(a1 + 1840);
	*(uint32_t*)(v36 + 40) = *(uint32_t*)(a1 + 1844);
	*(uint32_t*)(v36 + 44) = *(uint32_t*)(a1 + 1848);
	*(uint32_t*)(v36 + 48) = *(uint32_t*)(a1 + 1852);
	*(uint32_t*)(v36 + 52) = *(uint32_t*)(a1 + 1856);
	*(uint32_t*)(v36 + 56) = *(uint32_t*)(a1 + 1860);
	*(uint32_t*)(v36 + 60) = *(uint32_t*)(a1 + 1864);
	*(uint32_t*)(v36 + 64) = *(uint32_t*)(a1 + 1868);
	*(uint16_t*)(v36 + 68) = *(uint16_t*)(a1 + 1872);
	v37 = v2 + v36;
	*(uint16_t*)(v37 + 26) = *(uint16_t*)(a1 + 1874);
	*(uint32_t*)(v37 + 28) = *(uint32_t*)(a1 + 1876);
	*(uint32_t*)(v37 + 32) = *(uint32_t*)(a1 + 1880);
	*(uint32_t*)(v37 + 36) = *(uint32_t*)(a1 + 1884);
	*(uint32_t*)(v37 + 40) = *(uint32_t*)(a1 + 1888);
	*(uint32_t*)(v37 + 44) = *(uint32_t*)(a1 + 1892);
	*(uint32_t*)(v37 + 48) = *(uint32_t*)(a1 + 1896);
	*(uint32_t*)(v37 + 52) = *(uint32_t*)(a1 + 1900);
	*(uint32_t*)(v37 + 56) = *(uint32_t*)(a1 + 1904);
	*(uint32_t*)(v37 + 60) = *(uint32_t*)(a1 + 1908);
	*(uint32_t*)(v37 + 64) = *(uint32_t*)(a1 + 1912);
	v38 = v2 + v37;
	*(uint32_t*)(v38 + 28) = *(uint32_t*)(a1 + 1916);
	*(uint32_t*)(v38 + 32) = *(uint32_t*)(a1 + 1920);
	*(uint32_t*)(v38 + 36) = *(uint32_t*)(a1 + 1924);
	*(uint32_t*)(v38 + 40) = *(uint32_t*)(a1 + 1928);
	*(uint32_t*)(v38 + 44) = *(uint32_t*)(a1 + 1932);
	*(uint32_t*)(v38 + 48) = *(uint32_t*)(a1 + 1936);
	*(uint32_t*)(v38 + 52) = *(uint32_t*)(a1 + 1940);
	*(uint32_t*)(v38 + 56) = *(uint32_t*)(a1 + 1944);
	*(uint32_t*)(v38 + 60) = *(uint32_t*)(a1 + 1948);
	*(uint16_t*)(v38 + 64) = *(uint16_t*)(a1 + 1952);
	v39 = v2 + v38;
	*(uint16_t*)(v39 + 30) = *(uint16_t*)(a1 + 1954);
	*(uint32_t*)(v39 + 32) = *(uint32_t*)(a1 + 1956);
	*(uint32_t*)(v39 + 36) = *(uint32_t*)(a1 + 1960);
	*(uint32_t*)(v39 + 40) = *(uint32_t*)(a1 + 1964);
	*(uint32_t*)(v39 + 44) = *(uint32_t*)(a1 + 1968);
	*(uint32_t*)(v39 + 48) = *(uint32_t*)(a1 + 1972);
	*(uint32_t*)(v39 + 52) = *(uint32_t*)(a1 + 1976);
	*(uint32_t*)(v39 + 56) = *(uint32_t*)(a1 + 1980);
	*(uint32_t*)(v39 + 60) = *(uint32_t*)(a1 + 1984);
	v40 = v2 + v39;
	*(uint32_t*)(v40 + 32) = *(uint32_t*)(a1 + 1988);
	*(uint32_t*)(v40 + 36) = *(uint32_t*)(a1 + 1992);
	*(uint32_t*)(v40 + 40) = *(uint32_t*)(a1 + 1996);
	*(uint32_t*)(v40 + 44) = *(uint32_t*)(a1 + 2000);
	*(uint32_t*)(v40 + 48) = *(uint32_t*)(a1 + 2004);
	*(uint32_t*)(v40 + 52) = *(uint32_t*)(a1 + 2008);
	*(uint32_t*)(v40 + 56) = *(uint32_t*)(a1 + 2012);
	*(uint16_t*)(v40 + 60) = *(uint16_t*)(a1 + 2016);
	v41 = v2 + v40;
	*(uint16_t*)(v41 + 34) = *(uint16_t*)(a1 + 2018);
	*(uint32_t*)(v41 + 36) = *(uint32_t*)(a1 + 2020);
	*(uint32_t*)(v41 + 40) = *(uint32_t*)(a1 + 2024);
	*(uint32_t*)(v41 + 44) = *(uint32_t*)(a1 + 2028);
	*(uint32_t*)(v41 + 48) = *(uint32_t*)(a1 + 2032);
	*(uint32_t*)(v41 + 52) = *(uint32_t*)(a1 + 2036);
	*(uint32_t*)(v41 + 56) = *(uint32_t*)(a1 + 2040);
	v42 = v2 + v41;
	*(uint32_t*)(v42 + 36) = *(uint32_t*)(a1 + 2044);
	*(uint32_t*)(v42 + 40) = *(uint32_t*)(a1 + 2048);
	*(uint32_t*)(v42 + 44) = *(uint32_t*)(a1 + 2052);
	*(uint32_t*)(v42 + 48) = *(uint32_t*)(a1 + 2056);
	*(uint32_t*)(v42 + 52) = *(uint32_t*)(a1 + 2060);
	*(uint16_t*)(v42 + 56) = *(uint16_t*)(a1 + 2064);
	v43 = v2 + v42;
	*(uint16_t*)(v43 + 38) = *(uint16_t*)(a1 + 2066);
	*(uint32_t*)(v43 + 40) = *(uint32_t*)(a1 + 2068);
	*(uint32_t*)(v43 + 44) = *(uint32_t*)(a1 + 2072);
	*(uint32_t*)(v43 + 48) = *(uint32_t*)(a1 + 2076);
	*(uint32_t*)(v43 + 52) = *(uint32_t*)(a1 + 2080);
	v44 = v2 + v43;
	*(uint32_t*)(v44 + 40) = *(uint32_t*)(a1 + 2084);
	*(uint32_t*)(v44 + 44) = *(uint32_t*)(a1 + 2088);
	*(uint32_t*)(v44 + 48) = *(uint32_t*)(a1 + 2092);
	*(uint16_t*)(v44 + 52) = *(uint16_t*)(a1 + 2096);
	v45 = v2 + v44;
	*(uint16_t*)(v45 + 42) = *(uint16_t*)(a1 + 2098);
	*(uint32_t*)(v45 + 44) = *(uint32_t*)(a1 + 2100);
	*(uint32_t*)(v45 + 48) = *(uint32_t*)(a1 + 2104);
	v46 = v2 + v45;
	*(uint32_t*)(v46 + 44) = *(uint32_t*)(a1 + 2108);
	result = *(uint16_t*)(a1 + 2112);
	*(uint16_t*)(v46 + 48) = result;
	*(uint16_t*)(v2 + v46 + 46) = *(uint16_t*)(a1 + 2114);
	return result;
}

//----- (00484450) --------------------------------------------------------
int sub_484450(int a1, int a2) {
	int result; // eax
	int v3;     // edx
	int v4;     // edi
	int v5;     // edi
	int v6;     // edi
	int v7;     // edi
	int v8;     // edi
	int v9;     // edi
	int v10;    // edi
	int v11;    // edi
	int v12;    // edi
	int v13;    // edi
	int v14;    // edi
	int v15;    // edi
	int v16;    // edi
	int v17;    // edi
	int v18;    // edi
	int v19;    // edi
	int v20;    // edi
	int v21;    // edi
	int v22;    // edi
	int v23;    // edi
	int v24;    // edi
	int v25;    // edi
	int v26;    // edi
	int v27;    // edi
	int v28;    // edi
	int v29;    // edi
	int v30;    // edi
	int v31;    // edi
	int v32;    // edi
	int v33;    // edi
	int v34;    // edi
	int v35;    // edi
	int v36;    // edi
	int v37;    // edi
	int v38;    // edi
	int v39;    // edi
	int v40;    // edi
	int v41;    // edi
	int v42;    // edi
	int v43;    // edi
	int v44;    // edi
	int v45;    // edi
	int v46;    // edi
	int v47;    // edi

	result = a1;
	v3 = dword_5d4594_3798804;
	*(uint16_t*)(a2 + 46) = a1;
	v4 = v3 + a2;
	*(uint32_t*)(v4 + 44) = a1;
	*(uint16_t*)(v4 + 48) = a1;
	v5 = v3 + v3 + a2;
	*(uint16_t*)(v5 + 42) = a1;
	*(uint32_t*)(v5 + 44) = a1;
	*(uint32_t*)(v5 + 48) = a1;
	v6 = v3 + v5;
	*(uint32_t*)(v6 + 40) = a1;
	*(uint32_t*)(v6 + 44) = a1;
	*(uint32_t*)(v6 + 48) = a1;
	*(uint16_t*)(v6 + 52) = a1;
	v7 = v3 + v6;
	*(uint16_t*)(v7 + 38) = a1;
	*(uint32_t*)(v7 + 40) = a1;
	*(uint32_t*)(v7 + 44) = a1;
	*(uint32_t*)(v7 + 48) = a1;
	*(uint32_t*)(v7 + 52) = a1;
	v8 = v3 + v7;
	*(uint32_t*)(v8 + 36) = a1;
	*(uint32_t*)(v8 + 40) = a1;
	*(uint32_t*)(v8 + 44) = a1;
	*(uint32_t*)(v8 + 48) = a1;
	*(uint32_t*)(v8 + 52) = a1;
	*(uint16_t*)(v8 + 56) = a1;
	v9 = v3 + v8;
	*(uint16_t*)(v9 + 34) = a1;
	*(uint32_t*)(v9 + 36) = a1;
	*(uint32_t*)(v9 + 40) = a1;
	*(uint32_t*)(v9 + 44) = a1;
	*(uint32_t*)(v9 + 48) = a1;
	*(uint32_t*)(v9 + 52) = a1;
	*(uint32_t*)(v9 + 56) = a1;
	v10 = v3 + v9;
	*(uint32_t*)(v10 + 32) = a1;
	*(uint32_t*)(v10 + 36) = a1;
	*(uint32_t*)(v10 + 40) = a1;
	*(uint32_t*)(v10 + 44) = a1;
	*(uint32_t*)(v10 + 48) = a1;
	*(uint32_t*)(v10 + 52) = a1;
	*(uint32_t*)(v10 + 56) = a1;
	*(uint16_t*)(v10 + 60) = a1;
	v11 = v3 + v10;
	*(uint16_t*)(v11 + 30) = a1;
	*(uint32_t*)(v11 + 32) = a1;
	*(uint32_t*)(v11 + 36) = a1;
	*(uint32_t*)(v11 + 40) = a1;
	*(uint32_t*)(v11 + 44) = a1;
	*(uint32_t*)(v11 + 48) = a1;
	*(uint32_t*)(v11 + 52) = a1;
	*(uint32_t*)(v11 + 56) = a1;
	*(uint32_t*)(v11 + 60) = a1;
	v12 = v3 + v11;
	*(uint32_t*)(v12 + 28) = a1;
	*(uint32_t*)(v12 + 32) = a1;
	*(uint32_t*)(v12 + 36) = a1;
	*(uint32_t*)(v12 + 40) = a1;
	*(uint32_t*)(v12 + 44) = a1;
	*(uint32_t*)(v12 + 48) = a1;
	*(uint32_t*)(v12 + 52) = a1;
	*(uint32_t*)(v12 + 56) = a1;
	*(uint32_t*)(v12 + 60) = a1;
	*(uint16_t*)(v12 + 64) = a1;
	v13 = v3 + v12;
	*(uint16_t*)(v13 + 26) = a1;
	*(uint32_t*)(v13 + 28) = a1;
	*(uint32_t*)(v13 + 32) = a1;
	*(uint32_t*)(v13 + 36) = a1;
	*(uint32_t*)(v13 + 40) = a1;
	*(uint32_t*)(v13 + 44) = a1;
	*(uint32_t*)(v13 + 48) = a1;
	*(uint32_t*)(v13 + 52) = a1;
	*(uint32_t*)(v13 + 56) = a1;
	*(uint32_t*)(v13 + 60) = a1;
	*(uint32_t*)(v13 + 64) = a1;
	v14 = v3 + v13;
	*(uint32_t*)(v14 + 24) = a1;
	*(uint32_t*)(v14 + 28) = a1;
	*(uint32_t*)(v14 + 32) = a1;
	*(uint32_t*)(v14 + 36) = a1;
	*(uint32_t*)(v14 + 40) = a1;
	*(uint32_t*)(v14 + 44) = a1;
	*(uint32_t*)(v14 + 48) = a1;
	*(uint32_t*)(v14 + 52) = a1;
	*(uint32_t*)(v14 + 56) = a1;
	*(uint32_t*)(v14 + 60) = a1;
	*(uint32_t*)(v14 + 64) = a1;
	*(uint16_t*)(v14 + 68) = a1;
	v15 = v3 + v14;
	*(uint16_t*)(v15 + 22) = a1;
	*(uint32_t*)(v15 + 24) = a1;
	*(uint32_t*)(v15 + 28) = a1;
	*(uint32_t*)(v15 + 32) = a1;
	*(uint32_t*)(v15 + 36) = a1;
	*(uint32_t*)(v15 + 40) = a1;
	*(uint32_t*)(v15 + 44) = a1;
	*(uint32_t*)(v15 + 48) = a1;
	*(uint32_t*)(v15 + 52) = a1;
	*(uint32_t*)(v15 + 56) = a1;
	*(uint32_t*)(v15 + 60) = a1;
	*(uint32_t*)(v15 + 64) = a1;
	*(uint32_t*)(v15 + 68) = a1;
	v16 = v3 + v15;
	*(uint32_t*)(v16 + 20) = a1;
	*(uint32_t*)(v16 + 24) = a1;
	*(uint32_t*)(v16 + 28) = a1;
	*(uint32_t*)(v16 + 32) = a1;
	*(uint32_t*)(v16 + 36) = a1;
	*(uint32_t*)(v16 + 40) = a1;
	*(uint32_t*)(v16 + 44) = a1;
	*(uint32_t*)(v16 + 48) = a1;
	*(uint32_t*)(v16 + 52) = a1;
	*(uint32_t*)(v16 + 56) = a1;
	*(uint32_t*)(v16 + 60) = a1;
	*(uint32_t*)(v16 + 64) = a1;
	*(uint32_t*)(v16 + 68) = a1;
	*(uint16_t*)(v16 + 72) = a1;
	v17 = v3 + v16;
	*(uint16_t*)(v17 + 18) = a1;
	*(uint32_t*)(v17 + 20) = a1;
	*(uint32_t*)(v17 + 24) = a1;
	*(uint32_t*)(v17 + 28) = a1;
	*(uint32_t*)(v17 + 32) = a1;
	*(uint32_t*)(v17 + 36) = a1;
	*(uint32_t*)(v17 + 40) = a1;
	*(uint32_t*)(v17 + 44) = a1;
	*(uint32_t*)(v17 + 48) = a1;
	*(uint32_t*)(v17 + 52) = a1;
	*(uint32_t*)(v17 + 56) = a1;
	*(uint32_t*)(v17 + 60) = a1;
	*(uint32_t*)(v17 + 64) = a1;
	*(uint32_t*)(v17 + 68) = a1;
	*(uint32_t*)(v17 + 72) = a1;
	v18 = v3 + v17;
	*(uint32_t*)(v18 + 16) = a1;
	*(uint32_t*)(v18 + 20) = a1;
	*(uint32_t*)(v18 + 24) = a1;
	*(uint32_t*)(v18 + 28) = a1;
	*(uint32_t*)(v18 + 32) = a1;
	*(uint32_t*)(v18 + 36) = a1;
	*(uint32_t*)(v18 + 40) = a1;
	*(uint32_t*)(v18 + 44) = a1;
	*(uint32_t*)(v18 + 48) = a1;
	*(uint32_t*)(v18 + 52) = a1;
	*(uint32_t*)(v18 + 56) = a1;
	*(uint32_t*)(v18 + 60) = a1;
	*(uint32_t*)(v18 + 64) = a1;
	*(uint32_t*)(v18 + 68) = a1;
	*(uint32_t*)(v18 + 72) = a1;
	*(uint16_t*)(v18 + 76) = a1;
	v19 = v3 + v18;
	*(uint16_t*)(v19 + 14) = a1;
	*(uint32_t*)(v19 + 16) = a1;
	*(uint32_t*)(v19 + 20) = a1;
	*(uint32_t*)(v19 + 24) = a1;
	*(uint32_t*)(v19 + 28) = a1;
	*(uint32_t*)(v19 + 32) = a1;
	*(uint32_t*)(v19 + 36) = a1;
	*(uint32_t*)(v19 + 40) = a1;
	*(uint32_t*)(v19 + 44) = a1;
	*(uint32_t*)(v19 + 48) = a1;
	*(uint32_t*)(v19 + 52) = a1;
	*(uint32_t*)(v19 + 56) = a1;
	*(uint32_t*)(v19 + 60) = a1;
	*(uint32_t*)(v19 + 64) = a1;
	*(uint32_t*)(v19 + 68) = a1;
	*(uint32_t*)(v19 + 72) = a1;
	*(uint32_t*)(v19 + 76) = a1;
	v20 = v3 + v19;
	*(uint32_t*)(v20 + 12) = a1;
	*(uint32_t*)(v20 + 16) = a1;
	*(uint32_t*)(v20 + 20) = a1;
	*(uint32_t*)(v20 + 24) = a1;
	*(uint32_t*)(v20 + 28) = a1;
	*(uint32_t*)(v20 + 32) = a1;
	*(uint32_t*)(v20 + 36) = a1;
	*(uint32_t*)(v20 + 40) = a1;
	*(uint32_t*)(v20 + 44) = a1;
	*(uint32_t*)(v20 + 48) = a1;
	*(uint32_t*)(v20 + 52) = a1;
	*(uint32_t*)(v20 + 56) = a1;
	*(uint32_t*)(v20 + 60) = a1;
	*(uint32_t*)(v20 + 64) = a1;
	*(uint32_t*)(v20 + 68) = a1;
	*(uint32_t*)(v20 + 72) = a1;
	*(uint32_t*)(v20 + 76) = a1;
	*(uint16_t*)(v20 + 80) = a1;
	v21 = v3 + v20;
	*(uint16_t*)(v21 + 10) = a1;
	*(uint32_t*)(v21 + 12) = a1;
	*(uint32_t*)(v21 + 16) = a1;
	*(uint32_t*)(v21 + 20) = a1;
	*(uint32_t*)(v21 + 24) = a1;
	*(uint32_t*)(v21 + 28) = a1;
	*(uint32_t*)(v21 + 32) = a1;
	*(uint32_t*)(v21 + 36) = a1;
	*(uint32_t*)(v21 + 40) = a1;
	*(uint32_t*)(v21 + 44) = a1;
	*(uint32_t*)(v21 + 48) = a1;
	*(uint32_t*)(v21 + 52) = a1;
	*(uint32_t*)(v21 + 56) = a1;
	*(uint32_t*)(v21 + 60) = a1;
	*(uint32_t*)(v21 + 64) = a1;
	*(uint32_t*)(v21 + 68) = a1;
	*(uint32_t*)(v21 + 72) = a1;
	*(uint32_t*)(v21 + 76) = a1;
	*(uint32_t*)(v21 + 80) = a1;
	v22 = v3 + v21;
	*(uint32_t*)(v22 + 8) = a1;
	*(uint32_t*)(v22 + 12) = a1;
	*(uint32_t*)(v22 + 16) = a1;
	*(uint32_t*)(v22 + 20) = a1;
	*(uint32_t*)(v22 + 24) = a1;
	*(uint32_t*)(v22 + 28) = a1;
	*(uint32_t*)(v22 + 32) = a1;
	*(uint32_t*)(v22 + 36) = a1;
	*(uint32_t*)(v22 + 40) = a1;
	*(uint32_t*)(v22 + 44) = a1;
	*(uint32_t*)(v22 + 48) = a1;
	*(uint32_t*)(v22 + 52) = a1;
	*(uint32_t*)(v22 + 56) = a1;
	*(uint32_t*)(v22 + 60) = a1;
	*(uint32_t*)(v22 + 64) = a1;
	*(uint32_t*)(v22 + 68) = a1;
	*(uint32_t*)(v22 + 72) = a1;
	*(uint32_t*)(v22 + 76) = a1;
	*(uint32_t*)(v22 + 80) = a1;
	*(uint16_t*)(v22 + 84) = a1;
	v23 = v3 + v22;
	*(uint16_t*)(v23 + 6) = a1;
	*(uint32_t*)(v23 + 8) = a1;
	*(uint32_t*)(v23 + 12) = a1;
	*(uint32_t*)(v23 + 16) = a1;
	*(uint32_t*)(v23 + 20) = a1;
	*(uint32_t*)(v23 + 24) = a1;
	*(uint32_t*)(v23 + 28) = a1;
	*(uint32_t*)(v23 + 32) = a1;
	*(uint32_t*)(v23 + 36) = a1;
	*(uint32_t*)(v23 + 40) = a1;
	*(uint32_t*)(v23 + 44) = a1;
	*(uint32_t*)(v23 + 48) = a1;
	*(uint32_t*)(v23 + 52) = a1;
	*(uint32_t*)(v23 + 56) = a1;
	*(uint32_t*)(v23 + 60) = a1;
	*(uint32_t*)(v23 + 64) = a1;
	*(uint32_t*)(v23 + 68) = a1;
	*(uint32_t*)(v23 + 72) = a1;
	*(uint32_t*)(v23 + 76) = a1;
	*(uint32_t*)(v23 + 80) = a1;
	*(uint32_t*)(v23 + 84) = a1;
	v24 = v3 + v23;
	*(uint32_t*)(v24 + 4) = a1;
	*(uint32_t*)(v24 + 8) = a1;
	*(uint32_t*)(v24 + 12) = a1;
	*(uint32_t*)(v24 + 16) = a1;
	*(uint32_t*)(v24 + 20) = a1;
	*(uint32_t*)(v24 + 24) = a1;
	*(uint32_t*)(v24 + 28) = a1;
	*(uint32_t*)(v24 + 32) = a1;
	*(uint32_t*)(v24 + 36) = a1;
	*(uint32_t*)(v24 + 40) = a1;
	*(uint32_t*)(v24 + 44) = a1;
	*(uint32_t*)(v24 + 48) = a1;
	*(uint32_t*)(v24 + 52) = a1;
	*(uint32_t*)(v24 + 56) = a1;
	*(uint32_t*)(v24 + 60) = a1;
	*(uint32_t*)(v24 + 64) = a1;
	*(uint32_t*)(v24 + 68) = a1;
	*(uint32_t*)(v24 + 72) = a1;
	*(uint32_t*)(v24 + 76) = a1;
	*(uint32_t*)(v24 + 80) = a1;
	*(uint32_t*)(v24 + 84) = a1;
	*(uint16_t*)(v24 + 88) = a1;
	v25 = v3 + v24;
	*(uint16_t*)(v25 + 2) = a1;
	*(uint32_t*)(v25 + 4) = a1;
	*(uint32_t*)(v25 + 8) = a1;
	*(uint32_t*)(v25 + 12) = a1;
	*(uint32_t*)(v25 + 16) = a1;
	*(uint32_t*)(v25 + 20) = a1;
	*(uint32_t*)(v25 + 24) = a1;
	*(uint32_t*)(v25 + 28) = a1;
	*(uint32_t*)(v25 + 32) = a1;
	*(uint32_t*)(v25 + 36) = a1;
	*(uint32_t*)(v25 + 40) = a1;
	*(uint32_t*)(v25 + 44) = a1;
	*(uint32_t*)(v25 + 48) = a1;
	*(uint32_t*)(v25 + 52) = a1;
	*(uint32_t*)(v25 + 56) = a1;
	*(uint32_t*)(v25 + 60) = a1;
	*(uint32_t*)(v25 + 64) = a1;
	*(uint32_t*)(v25 + 68) = a1;
	*(uint32_t*)(v25 + 72) = a1;
	*(uint32_t*)(v25 + 76) = a1;
	*(uint32_t*)(v25 + 80) = a1;
	*(uint32_t*)(v25 + 84) = a1;
	*(uint32_t*)(v25 + 88) = a1;
	v26 = v3 + v25;
	*(uint16_t*)(v26 + 2) = a1;
	*(uint32_t*)(v26 + 4) = a1;
	*(uint32_t*)(v26 + 8) = a1;
	*(uint32_t*)(v26 + 12) = a1;
	*(uint32_t*)(v26 + 16) = a1;
	*(uint32_t*)(v26 + 20) = a1;
	*(uint32_t*)(v26 + 24) = a1;
	*(uint32_t*)(v26 + 28) = a1;
	*(uint32_t*)(v26 + 32) = a1;
	*(uint32_t*)(v26 + 36) = a1;
	*(uint32_t*)(v26 + 40) = a1;
	*(uint32_t*)(v26 + 44) = a1;
	*(uint32_t*)(v26 + 48) = a1;
	*(uint32_t*)(v26 + 52) = a1;
	*(uint32_t*)(v26 + 56) = a1;
	*(uint32_t*)(v26 + 60) = a1;
	*(uint32_t*)(v26 + 64) = a1;
	*(uint32_t*)(v26 + 68) = a1;
	*(uint32_t*)(v26 + 72) = a1;
	*(uint32_t*)(v26 + 76) = a1;
	*(uint32_t*)(v26 + 80) = a1;
	*(uint32_t*)(v26 + 84) = a1;
	*(uint32_t*)(v26 + 88) = a1;
	v27 = v3 + v26;
	*(uint32_t*)(v27 + 4) = a1;
	*(uint32_t*)(v27 + 8) = a1;
	*(uint32_t*)(v27 + 12) = a1;
	*(uint32_t*)(v27 + 16) = a1;
	*(uint32_t*)(v27 + 20) = a1;
	*(uint32_t*)(v27 + 24) = a1;
	*(uint32_t*)(v27 + 28) = a1;
	*(uint32_t*)(v27 + 32) = a1;
	*(uint32_t*)(v27 + 36) = a1;
	*(uint32_t*)(v27 + 40) = a1;
	*(uint32_t*)(v27 + 44) = a1;
	*(uint32_t*)(v27 + 48) = a1;
	*(uint32_t*)(v27 + 52) = a1;
	*(uint32_t*)(v27 + 56) = a1;
	*(uint32_t*)(v27 + 60) = a1;
	*(uint32_t*)(v27 + 64) = a1;
	*(uint32_t*)(v27 + 68) = a1;
	*(uint32_t*)(v27 + 72) = a1;
	*(uint32_t*)(v27 + 76) = a1;
	*(uint32_t*)(v27 + 80) = a1;
	*(uint32_t*)(v27 + 84) = a1;
	*(uint16_t*)(v27 + 88) = a1;
	v28 = v3 + v27;
	*(uint16_t*)(v28 + 6) = a1;
	*(uint32_t*)(v28 + 8) = a1;
	*(uint32_t*)(v28 + 12) = a1;
	*(uint32_t*)(v28 + 16) = a1;
	*(uint32_t*)(v28 + 20) = a1;
	*(uint32_t*)(v28 + 24) = a1;
	*(uint32_t*)(v28 + 28) = a1;
	*(uint32_t*)(v28 + 32) = a1;
	*(uint32_t*)(v28 + 36) = a1;
	*(uint32_t*)(v28 + 40) = a1;
	*(uint32_t*)(v28 + 44) = a1;
	*(uint32_t*)(v28 + 48) = a1;
	*(uint32_t*)(v28 + 52) = a1;
	*(uint32_t*)(v28 + 56) = a1;
	*(uint32_t*)(v28 + 60) = a1;
	*(uint32_t*)(v28 + 64) = a1;
	*(uint32_t*)(v28 + 68) = a1;
	*(uint32_t*)(v28 + 72) = a1;
	*(uint32_t*)(v28 + 76) = a1;
	*(uint32_t*)(v28 + 80) = a1;
	*(uint32_t*)(v28 + 84) = a1;
	v29 = v3 + v28;
	*(uint32_t*)(v29 + 8) = a1;
	*(uint32_t*)(v29 + 12) = a1;
	*(uint32_t*)(v29 + 16) = a1;
	*(uint32_t*)(v29 + 20) = a1;
	*(uint32_t*)(v29 + 24) = a1;
	*(uint32_t*)(v29 + 28) = a1;
	*(uint32_t*)(v29 + 32) = a1;
	*(uint32_t*)(v29 + 36) = a1;
	*(uint32_t*)(v29 + 40) = a1;
	*(uint32_t*)(v29 + 44) = a1;
	*(uint32_t*)(v29 + 48) = a1;
	*(uint32_t*)(v29 + 52) = a1;
	*(uint32_t*)(v29 + 56) = a1;
	*(uint32_t*)(v29 + 60) = a1;
	*(uint32_t*)(v29 + 64) = a1;
	*(uint32_t*)(v29 + 68) = a1;
	*(uint32_t*)(v29 + 72) = a1;
	*(uint32_t*)(v29 + 76) = a1;
	*(uint32_t*)(v29 + 80) = a1;
	*(uint16_t*)(v29 + 84) = a1;
	v30 = v3 + v29;
	*(uint16_t*)(v30 + 10) = a1;
	*(uint32_t*)(v30 + 12) = a1;
	*(uint32_t*)(v30 + 16) = a1;
	*(uint32_t*)(v30 + 20) = a1;
	*(uint32_t*)(v30 + 24) = a1;
	*(uint32_t*)(v30 + 28) = a1;
	*(uint32_t*)(v30 + 32) = a1;
	*(uint32_t*)(v30 + 36) = a1;
	*(uint32_t*)(v30 + 40) = a1;
	*(uint32_t*)(v30 + 44) = a1;
	*(uint32_t*)(v30 + 48) = a1;
	*(uint32_t*)(v30 + 52) = a1;
	*(uint32_t*)(v30 + 56) = a1;
	*(uint32_t*)(v30 + 60) = a1;
	*(uint32_t*)(v30 + 64) = a1;
	*(uint32_t*)(v30 + 68) = a1;
	*(uint32_t*)(v30 + 72) = a1;
	*(uint32_t*)(v30 + 76) = a1;
	*(uint32_t*)(v30 + 80) = a1;
	v31 = v3 + v30;
	*(uint32_t*)(v31 + 12) = a1;
	*(uint32_t*)(v31 + 16) = a1;
	*(uint32_t*)(v31 + 20) = a1;
	*(uint32_t*)(v31 + 24) = a1;
	*(uint32_t*)(v31 + 28) = a1;
	*(uint32_t*)(v31 + 32) = a1;
	*(uint32_t*)(v31 + 36) = a1;
	*(uint32_t*)(v31 + 40) = a1;
	*(uint32_t*)(v31 + 44) = a1;
	*(uint32_t*)(v31 + 48) = a1;
	*(uint32_t*)(v31 + 52) = a1;
	*(uint32_t*)(v31 + 56) = a1;
	*(uint32_t*)(v31 + 60) = a1;
	*(uint32_t*)(v31 + 64) = a1;
	*(uint32_t*)(v31 + 68) = a1;
	*(uint32_t*)(v31 + 72) = a1;
	*(uint32_t*)(v31 + 76) = a1;
	*(uint16_t*)(v31 + 80) = a1;
	v32 = v3 + v31;
	*(uint16_t*)(v32 + 14) = a1;
	*(uint32_t*)(v32 + 16) = a1;
	*(uint32_t*)(v32 + 20) = a1;
	*(uint32_t*)(v32 + 24) = a1;
	*(uint32_t*)(v32 + 28) = a1;
	*(uint32_t*)(v32 + 32) = a1;
	*(uint32_t*)(v32 + 36) = a1;
	*(uint32_t*)(v32 + 40) = a1;
	*(uint32_t*)(v32 + 44) = a1;
	*(uint32_t*)(v32 + 48) = a1;
	*(uint32_t*)(v32 + 52) = a1;
	*(uint32_t*)(v32 + 56) = a1;
	*(uint32_t*)(v32 + 60) = a1;
	*(uint32_t*)(v32 + 64) = a1;
	*(uint32_t*)(v32 + 68) = a1;
	*(uint32_t*)(v32 + 72) = a1;
	*(uint32_t*)(v32 + 76) = a1;
	v33 = v3 + v32;
	*(uint32_t*)(v33 + 16) = a1;
	*(uint32_t*)(v33 + 20) = a1;
	*(uint32_t*)(v33 + 24) = a1;
	*(uint32_t*)(v33 + 28) = a1;
	*(uint32_t*)(v33 + 32) = a1;
	*(uint32_t*)(v33 + 36) = a1;
	*(uint32_t*)(v33 + 40) = a1;
	*(uint32_t*)(v33 + 44) = a1;
	*(uint32_t*)(v33 + 48) = a1;
	*(uint32_t*)(v33 + 52) = a1;
	*(uint32_t*)(v33 + 56) = a1;
	*(uint32_t*)(v33 + 60) = a1;
	*(uint32_t*)(v33 + 64) = a1;
	*(uint32_t*)(v33 + 68) = a1;
	*(uint32_t*)(v33 + 72) = a1;
	*(uint16_t*)(v33 + 76) = a1;
	v34 = v3 + v33;
	*(uint16_t*)(v34 + 18) = a1;
	*(uint32_t*)(v34 + 20) = a1;
	*(uint32_t*)(v34 + 24) = a1;
	*(uint32_t*)(v34 + 28) = a1;
	*(uint32_t*)(v34 + 32) = a1;
	*(uint32_t*)(v34 + 36) = a1;
	*(uint32_t*)(v34 + 40) = a1;
	*(uint32_t*)(v34 + 44) = a1;
	*(uint32_t*)(v34 + 48) = a1;
	*(uint32_t*)(v34 + 52) = a1;
	*(uint32_t*)(v34 + 56) = a1;
	*(uint32_t*)(v34 + 60) = a1;
	*(uint32_t*)(v34 + 64) = a1;
	*(uint32_t*)(v34 + 68) = a1;
	*(uint32_t*)(v34 + 72) = a1;
	v35 = v3 + v34;
	*(uint32_t*)(v35 + 20) = a1;
	*(uint32_t*)(v35 + 24) = a1;
	*(uint32_t*)(v35 + 28) = a1;
	*(uint32_t*)(v35 + 32) = a1;
	*(uint32_t*)(v35 + 36) = a1;
	*(uint32_t*)(v35 + 40) = a1;
	*(uint32_t*)(v35 + 44) = a1;
	*(uint32_t*)(v35 + 48) = a1;
	*(uint32_t*)(v35 + 52) = a1;
	*(uint32_t*)(v35 + 56) = a1;
	*(uint32_t*)(v35 + 60) = a1;
	*(uint32_t*)(v35 + 64) = a1;
	*(uint32_t*)(v35 + 68) = a1;
	*(uint16_t*)(v35 + 72) = a1;
	v36 = v3 + v35;
	*(uint16_t*)(v36 + 22) = a1;
	*(uint32_t*)(v36 + 24) = a1;
	*(uint32_t*)(v36 + 28) = a1;
	*(uint32_t*)(v36 + 32) = a1;
	*(uint32_t*)(v36 + 36) = a1;
	*(uint32_t*)(v36 + 40) = a1;
	*(uint32_t*)(v36 + 44) = a1;
	*(uint32_t*)(v36 + 48) = a1;
	*(uint32_t*)(v36 + 52) = a1;
	*(uint32_t*)(v36 + 56) = a1;
	*(uint32_t*)(v36 + 60) = a1;
	*(uint32_t*)(v36 + 64) = a1;
	*(uint32_t*)(v36 + 68) = a1;
	v37 = v3 + v36;
	*(uint32_t*)(v37 + 24) = a1;
	*(uint32_t*)(v37 + 28) = a1;
	*(uint32_t*)(v37 + 32) = a1;
	*(uint32_t*)(v37 + 36) = a1;
	*(uint32_t*)(v37 + 40) = a1;
	*(uint32_t*)(v37 + 44) = a1;
	*(uint32_t*)(v37 + 48) = a1;
	*(uint32_t*)(v37 + 52) = a1;
	*(uint32_t*)(v37 + 56) = a1;
	*(uint32_t*)(v37 + 60) = a1;
	*(uint32_t*)(v37 + 64) = a1;
	*(uint16_t*)(v37 + 68) = a1;
	v38 = v3 + v37;
	*(uint16_t*)(v38 + 26) = a1;
	*(uint32_t*)(v38 + 28) = a1;
	*(uint32_t*)(v38 + 32) = a1;
	*(uint32_t*)(v38 + 36) = a1;
	*(uint32_t*)(v38 + 40) = a1;
	*(uint32_t*)(v38 + 44) = a1;
	*(uint32_t*)(v38 + 48) = a1;
	*(uint32_t*)(v38 + 52) = a1;
	*(uint32_t*)(v38 + 56) = a1;
	*(uint32_t*)(v38 + 60) = a1;
	*(uint32_t*)(v38 + 64) = a1;
	v39 = v3 + v38;
	*(uint32_t*)(v39 + 28) = a1;
	*(uint32_t*)(v39 + 32) = a1;
	*(uint32_t*)(v39 + 36) = a1;
	*(uint32_t*)(v39 + 40) = a1;
	*(uint32_t*)(v39 + 44) = a1;
	*(uint32_t*)(v39 + 48) = a1;
	*(uint32_t*)(v39 + 52) = a1;
	*(uint32_t*)(v39 + 56) = a1;
	*(uint32_t*)(v39 + 60) = a1;
	*(uint16_t*)(v39 + 64) = a1;
	v40 = v3 + v39;
	*(uint16_t*)(v40 + 30) = a1;
	*(uint32_t*)(v40 + 32) = a1;
	*(uint32_t*)(v40 + 36) = a1;
	*(uint32_t*)(v40 + 40) = a1;
	*(uint32_t*)(v40 + 44) = a1;
	*(uint32_t*)(v40 + 48) = a1;
	*(uint32_t*)(v40 + 52) = a1;
	*(uint32_t*)(v40 + 56) = a1;
	*(uint32_t*)(v40 + 60) = a1;
	v41 = v3 + v40;
	*(uint32_t*)(v41 + 32) = a1;
	*(uint32_t*)(v41 + 36) = a1;
	*(uint32_t*)(v41 + 40) = a1;
	*(uint32_t*)(v41 + 44) = a1;
	*(uint32_t*)(v41 + 48) = a1;
	*(uint32_t*)(v41 + 52) = a1;
	*(uint32_t*)(v41 + 56) = a1;
	*(uint16_t*)(v41 + 60) = a1;
	v42 = v3 + v41;
	*(uint16_t*)(v42 + 34) = a1;
	*(uint32_t*)(v42 + 36) = a1;
	*(uint32_t*)(v42 + 40) = a1;
	*(uint32_t*)(v42 + 44) = a1;
	*(uint32_t*)(v42 + 48) = a1;
	*(uint32_t*)(v42 + 52) = a1;
	*(uint32_t*)(v42 + 56) = a1;
	v43 = v3 + v42;
	*(uint32_t*)(v43 + 36) = a1;
	*(uint32_t*)(v43 + 40) = a1;
	*(uint32_t*)(v43 + 44) = a1;
	*(uint32_t*)(v43 + 48) = a1;
	*(uint32_t*)(v43 + 52) = a1;
	*(uint16_t*)(v43 + 56) = a1;
	v44 = v3 + v43;
	*(uint16_t*)(v44 + 38) = a1;
	*(uint32_t*)(v44 + 40) = a1;
	*(uint32_t*)(v44 + 44) = a1;
	*(uint32_t*)(v44 + 48) = a1;
	*(uint32_t*)(v44 + 52) = a1;
	v45 = v3 + v44;
	*(uint32_t*)(v45 + 40) = a1;
	*(uint32_t*)(v45 + 44) = a1;
	*(uint32_t*)(v45 + 48) = a1;
	*(uint16_t*)(v45 + 52) = a1;
	v46 = v3 + v45;
	*(uint16_t*)(v46 + 42) = a1;
	*(uint32_t*)(v46 + 44) = a1;
	*(uint32_t*)(v46 + 48) = a1;
	v47 = v3 + v46;
	*(uint32_t*)(v47 + 44) = a1;
	*(uint16_t*)(v47 + 48) = a1;
	*(uint16_t*)(v3 + v47 + 46) = a1;
	return result;
}

//----- (00484BE0) --------------------------------------------------------
uint32_t* nox_xxx_spriteChangeLightColor_484BE0(uint32_t* a1, int a2, int a3, int a4) {
	uint32_t* result; // eax

	result = a1;
	a1[4] = a2;
	*a1 = 2;
	a1[5] = a3;
	a1[6] = a4;
	return result;
}

//----- (00484C00) --------------------------------------------------------
long long sub_484C00(int a1, int a2) {
	long long result; // rax

	result = (long long)((double)a2 * 0.0027777778 * *(double*)&qword_581450_9552 + *(double*)&qword_581450_9544);
	*(uint16_t*)(a1 + 28) = result;
	*(uint32_t*)(a1 + 32) = 0;
	return result;
}

//----- (00484C30) --------------------------------------------------------
long long nox_xxx_spriteChangeLightSize_484C30(int a1, int a2) {
	long long result; // rax

	result = (long long)((double)a2 * 0.0027777778 * *(double*)&qword_581450_9552 + *(double*)&qword_581450_9544);
	*(uint16_t*)(a1 + 30) = result;
	return result;
}

//----- (00484CE0) --------------------------------------------------------
int sub_484CE0(int a1, float a2) {
	int result; // eax

	if (a2 > 63.0) {
		a2 = 63.0;
	}
	*(float*)(a1 + 4) = a2;
	result = sub_484C60(a2);
	*(uint32_t*)(a1 + 8) = result;
	return result;
}

//----- (00484D70) --------------------------------------------------------
int nox_xxx_spriteChangeIntensity_484D70_light_intensity(int a1, float a2) {
	int result; // eax

	if (a2 > 63.0) {
		a2 = 63.0;
	}
	*(float*)(a1 + 4) = a2;
	*(uint32_t*)(a1 + 12) = (long long)(a2 * *(double*)&qword_581450_9552 + *(double*)&qword_581450_9544);
	result = sub_484C60(a2);
	*(uint32_t*)(a1 + 8) = result;
	return result;
}


//----- (00485B30) --------------------------------------------------------
int nox_thing_read_floor_485B30(nox_memfile* f, char* a2) {
	int a1 = f;
	int v2;            // esi
	uint8_t* v3;       // edi
	uint8_t* v9;       // edx
	int v10;           // edi
	int* v12;          // eax
	int v13;           // ecx
	char* v14;         // eax
	char v15;          // cl
	int v16;           // [esp-4h] [ebp-40h]
	unsigned char v17; // [esp+10h] [ebp-2Ch]
	int i;             // [esp+10h] [ebp-2Ch]
	unsigned char v19; // [esp+14h] [ebp-28h]
	const char* v21;   // [esp+18h] [ebp-24h]
	char v22[32];      // [esp+1Ch] [ebp-20h]
	unsigned char v23; // [esp+40h] [ebp+4h]

	v2 = a1;
	v16 = a1;
	v3 = (uint8_t*)(*(uint32_t*)(a1 + 8) + 4);
	*(uint32_t*)(a1 + 8) = v3;
	LOBYTE(a1) = *v3;
	*(uint32_t*)(v2 + 8) = v3 + 1;
	nox_memfile_read(v22, 1u, (unsigned char)a1, v16);
	v22[(unsigned char)a1] = 0;
	int v7 = a1;
	if (nox_tile_def_cnt > 0) {
		int v5 = 0;
		for (v5 = 0; v5 < nox_tile_def_cnt; v5++) {
			nox_tileDef_t* p = &nox_tile_defs_arr[v5];
			if (strcmp(&p->name[0], v22) == 0) {
				v7 = v5;
				break;
			}
		}
		if (v5 == nox_tile_def_cnt) {
			return 0;
		}
	}
	v9 = (uint8_t*)(*(uint32_t*)(v2 + 8) + 12);
	*(uint32_t*)(v2 + 8) = v9;
	LOBYTE(v21) = *v9;
	*(uint32_t*)(v2 + 8) = v9 + 1;
	v19 = v9[1];
	*(uint32_t*)(v2 + 8) = v9 + 2;
	v17 = v9[2];
	*(uint32_t*)(v2 + 8) = v9 + 4;
	v10 = (unsigned char)v21 * v19 * v17;
	nox_tile_defs_arr[v7].data_32 = calloc(v10, 4);
	int v11 = 0;
	for (i = 0; v11 < v10; ) {
		v12 = *(int**)(v2 + 8);
		v13 = *v12;
		*(uint32_t*)(v2 + 8) = v12 + 1;
		*a2 = getMemByte(0x5D4594, 1193192);
		if (v13 == -1) {
			v14 = *(char**)(v2 + 8);
			v15 = *v14++;
			*(uint32_t*)(v2 + 8) = v14;
			LOBYTE(v21) = v15;
			v23 = *v14;
			*(uint32_t*)(v2 + 8) = v14 + 1;
			nox_memfile_read(a2, 1u, v23, v2);
			v13 = -1;
			a2[v23] = 0;
			v11 = i;
		}
		nox_tile_defs_arr[v7].data_32[v11] = nox_xxx_readImgMB_42FAA0(v13, v21, a2);
		v11++;
		i = v11;
	}
	return 1;
}
// 485CB9: variable 'v21' is possibly undefined
// 485B30: using guessed type char var_20[32];

//----- (00485D40) --------------------------------------------------------
int nox_thing_read_edge_485D40(nox_memfile* f, char* a2) {
	int a1 = f;
	int v2;            // esi
	uint8_t* v3;       // edi
	int v4;            // eax
	int v5;            // ebp
	const char* v6;    // ebx
	int v7;            // ebx
	int result;        // eax
	uint8_t* v9;       // edi
	char v10;          // cl
	unsigned char v11; // dl
	int v12;           // eax
	unsigned char v13; // di
	int v14;           // edi
	int v15;           // ebx
	int* v16;          // eax
	int v17;           // ecx
	char* v18;         // eax
	char v19;          // cl
	int* v20;          // eax
	int v21;           // ecx
	int v22;           // [esp-4h] [ebp-40h]
	int i;             // [esp+10h] [ebp-2Ch]
	int v24;           // [esp+14h] [ebp-28h]
	unsigned int v25;  // [esp+18h] [ebp-24h]
	char v26[64];      // [esp+1Ch] [ebp-20h]
	unsigned char v27; // [esp+40h] [ebp+4h]

	v2 = a1;
	v22 = a1;
	v3 = (uint8_t*)(*(uint32_t*)(a1 + 8) + 4);
	*(uint32_t*)(a1 + 8) = v3;
	LOBYTE(a1) = *v3;
	*(uint32_t*)(v2 + 8) = v3 + 1;
	nox_memfile_read(v26, 1u, (unsigned char)a1, v22);
	v4 = dword_5d4594_251572;
	v5 = 0;
	v26[(unsigned char)a1] = 0;
	if (v4 <= 0) {
		v7 = a1;
	} else {
		v6 = (const char*)getMemAt(0x85B3FC, 28644);
		while (1) {
			if (!strcmp(v6, v26)) {
				v7 = v5;
				break;
			}
			++v5;
			v6 += 60;
			if (v5 >= *(int*)&dword_5d4594_251572) {
				v7 = a1;
				break;
			}
		}
	}
	if (v5 == dword_5d4594_251572) {
		return 0;
	}
	v9 = (uint8_t*)(*(uint32_t*)(v2 + 8) + 9);
	*(uint32_t*)(v2 + 8) = v9;
	v25 = *v9;
	*(uint32_t*)(v2 + 8) = v9 + 2;
	v10 = v9[2];
	*(uint32_t*)(v2 + 8) = v9 + 3;
	if (v10 == 1) {
		return 0;
	}
	v11 = v9[3];
	v12 = (int)(v9 + 4);
	*(uint32_t*)(v2 + 8) = v9 + 4;
	v13 = v9[4];
	*(uint32_t*)(v2 + 8) = v12 + 1;
	v14 = 2 * v25 * (v11 + v13);
	result = (int)calloc(v14, 5);
	memset((int*)result, 0, 5 * v14);
	v24 = 15 * v7;
	*getMemU32Ptr(0x85B3FC, 28676 + 60 * v7) = result;
	if (result) {
		v15 = 0;
		for (i = 0; v15 < v14; *(uint32_t*)(*getMemU32Ptr(0x85B3FC, 28676 + 4 * v24) + 4 * v15 - 4) =
								   nox_xxx_readImgMB_42FAA0(v17, v25, a2)) {
			v16 = *(int**)(v2 + 8);
			v17 = *v16;
			*(uint32_t*)(v2 + 8) = v16 + 1;
			*a2 = getMemByte(0x5D4594, 1193196);
			if (v17 == -1) {
				v18 = *(char**)(v2 + 8);
				v19 = *v18++;
				*(uint32_t*)(v2 + 8) = v18;
				LOBYTE(v25) = v19;
				v27 = *v18;
				*(uint32_t*)(v2 + 8) = v18 + 1;
				nox_memfile_read(a2, 1u, v27, v2);
				v17 = -1;
				a2[v27] = 0;
				v15 = i;
			}
			i = ++v15;
		}
		v20 = *(int**)(v2 + 8);
		v21 = *v20;
		*(uint32_t*)(v2 + 8) = v20 + 1;
		result = v21 == 1162757152;
	}
	return result;
}
// 485EEB: variable 'v25' is possibly undefined
// 485D40: using guessed type char var_20[32];

//----- (00486060) --------------------------------------------------------
int nox_xxx_tile_486060() {
	int result; // eax

	dword_5d4594_1193188 = 1;
	result = 45 * dword_5d4594_3798804;
	*getMemU32Ptr(0x973CE0, 376) = 45 * dword_5d4594_3798804 + (46 << getMemByte(0x973F18, 7696));
	return result;
}

//----- (00486640) --------------------------------------------------------
int sub_486640(void* a1p, int a2) {
	int a1 = a1p;
	return a2 * (*(uint32_t*)(a1 + 36) >> 16) / 100;
}

//----- (004866D0) --------------------------------------------------------
int sub_4866D0(uint32_t* a1, int a2) { return *a1 + 36 * a2; }

//----- (00486A10) --------------------------------------------------------
unsigned int sub_486A10(int a1, void* a2) {
	void* v2;            // eax
	unsigned int result; // eax

	v2 = bsearch(a2, *(const void**)a1, *(uint32_t*)(a1 + 4), 0x24u, (int (*)(const void*, const void*))nox_strcmpi);
	if (v2) {
		result = ((unsigned int)v2 - *(uint32_t*)a1) / 0x24;
	} else {
		result = -1;
	}
	return result;
}

//----- (00486AA0) --------------------------------------------------------
int sub_486AA0(uint32_t* a1, int a2, uint32_t* a3) {
	uint32_t* v3; // eax
	int result;   // eax

	v3 = (uint32_t*)sub_4866D0(a1, a2);
	*a3 = 4;
	a3[2] = v3[6];
	a3[3] = ((v3[7] & 1) != 0) + 1;
	a3[6] = v3[8];
	if (v3[7] & 8) {
		result = 2;
		a3[1] = 2;
		a3[4] = 2;
	} else {
		a3[1] = 0;
		result = ((v3[7] & 4) != 0) + 1;
		a3[4] = result;
	}
	return result;
}

//----- (00486B60) --------------------------------------------------------
int sub_486B60(int a1, int a2) {
	int v2;       // ebp
	FILE* v3;     // eax
	FILE* v6;     // eax
	FILE* v7;     // esi
	int v8;       // edi
	int v9;       // eax
	int v10;      // ecx
	int v12;      // [esp+10h] [ebp-12Ch]
	char v13[8];  // [esp+14h] [ebp-128h]
	char v14[16]; // [esp+1Ch] [ebp-120h]
	int v15[12];  // [esp+2Ch] [ebp-110h]

	v12 = 1;
	v2 = sub_4866D0((uint32_t*)a1, a2);
	sub_486E00(a1);
	v3 = *(FILE**)(a1 + 268);
	*(uint32_t*)(a1 + 280) = v3;
	*(uint32_t*)(a1 + 284) = *(uint32_t*)(v2 + 20);
	if (nox_fs_fseek_start(v3, *(uint32_t*)(v2 + 16))) {
		v12 = 0;
	}
	if (!*(uint32_t*)(v2 + 20)) {
		v12 = 0;
	}
	if (!*(uint32_t*)(a1 + 276)) {
		return v12;
	}
	strcpy((char*)&v15[3], (const char*)(a1 + 8));
	strcat((char*)&v15[3], (const char*)v2);
	strcat((char*)&v15[3], ".wav");
	v6 = nox_fs_open((const char*)&v15[3]);
	v7 = v6;
	v8 = 0;
	*(uint32_t*)(a1 + 272) = v6;
	if (!v6) {
		return v12;
	}
	if (nox_binfile_fread_raw_40ADD0((char*)v15, 0xCu, 1u, v6) != 1 || v15[0] != 1179011410 || v15[2] != 1163280727) {
		printf("error: '%s' is bad - cannot read\n", &v15[3]);
		if (*(uint32_t*)(a1 + 272)) {
			nox_fs_close(*(FILE**)(a1 + 272));
			*(uint32_t*)(a1 + 272) = 0;
		}
		return v12;
	}
	if (nox_binfile_fread_raw_40ADD0(v13, 8u, 1u, v7) != 1) {
		goto LABEL_18;
	}
	while (1) {
		if (*(uint32_t*)v13 == 544501094) {
			nox_binfile_fread_raw_40ADD0(v14, 0x10u, 1u, v7);
			nox_fs_fseek_cur(v7, *(uint32_t*)&v13[4] - 16);
			goto LABEL_15;
		}
		if (*(uint32_t*)v13 == 1635017060) {
			break;
		}
		nox_fs_fseek_cur(v7, *(int*)&v13[4]);
	LABEL_15:
		if (nox_binfile_fread_raw_40ADD0(v13, 8u, 1u, v7) != 1) {
			goto LABEL_18;
		}
	}
	v8 = *(uint32_t*)&v13[4];
LABEL_18:
	*(uint32_t*)(v2 + 28) = 2;
	if (*(unsigned short*)&v14[12] / (int)*(unsigned short*)&v14[2] == 2) {
		*(uint32_t*)(v2 + 28) = 6;
	}
	if (*(uint16_t*)&v14[2] == 2) {
		v9 = *(uint32_t*)(v2 + 28);
		LOBYTE(v9) = v9 | 1;
		*(uint32_t*)(v2 + 28) = v9;
	}
	*(uint32_t*)(v2 + 24) = *(uint32_t*)&v14[4];
	v10 = *(uint32_t*)(a1 + 272);
	*(uint32_t*)(a1 + 284) = v8;
	*(uint32_t*)(a1 + 280) = v10;
	return 1;
}

//----- (00486DB0) --------------------------------------------------------
signed int sub_486DB0(int a1, char* a2, signed int a3) {
	signed int result; // eax
	signed int v4;     // eax

	if (!*(uint32_t*)(a1 + 280)) {
		return 0;
	}
	v4 = a3;
	if (a3 > *(int*)(a1 + 284)) {
		v4 = *(uint32_t*)(a1 + 284);
	}
	if (v4 <= 0 || (result = nox_binfile_fread_raw_40ADD0(a2, 1u, v4, *(FILE**)(a1 + 280)), result < 0)) {
		result = 0;
	}
	*(uint32_t*)(a1 + 284) -= result;
	return result;
}

//----- (00486E00) --------------------------------------------------------
FILE* sub_486E00(int a1) {
	FILE* result; // eax

	result = *(FILE**)(a1 + 272);
	*(uint32_t*)(a1 + 280) = 0;
	if (result) {
		nox_fs_close(result);
		result = 0;
		*(uint32_t*)(a1 + 272) = 0;
	}
	return result;
}

//----- (00486E30) --------------------------------------------------------
int sub_486E30(int a1, uint32_t* a2) {
	int result; // eax

	a2[33] = a1;
	++*(uint32_t*)(a1 + 192);
	++*(uint32_t*)(a1 + 212);
	nox_common_list_append_4258E0(a1 + 200, a2);
	result = *(uint32_t*)(a1 + 212) - 1;
	*(uint32_t*)(a1 + 212) = result;
	if (result < 0) {
		*(uint32_t*)(a1 + 212) = 0;
	}
	return result;
}

//----- (00486E90) --------------------------------------------------------
int sub_486E90(int a1) {
	int v1;     // esi
	int result; // eax

	v1 = *(uint32_t*)(a1 + 132);
	nox_common_list_remove_425920((uint32_t**)a1);
	--*(uint32_t*)(v1 + 192);
	++*(uint32_t*)(v1 + 212);
	nox_common_list_remove_425920((uint32_t**)a1);
	result = *(uint32_t*)(v1 + 212) - 1;
	*(uint32_t*)(v1 + 212) = result;
	if (result < 0) {
		*(uint32_t*)(v1 + 212) = 0;
	}
	return result;
}

//----- (00486FA0) --------------------------------------------------------
uint32_t* sub_486FA0(int a1) {
	uint32_t* result; // eax
	uint32_t* v2;     // edi
	int v3;           // eax

	result = sub_486FE0(a1);
	v2 = result;
	if (result) {
		v3 = *(uint32_t*)(a1 + 12);
		LOBYTE(v3) = v3 | 1;
		*(uint32_t*)(a1 + 12) = v3;
		sub_487050(v2);
		if (*(uint8_t*)(a1 + 8) & 2) {
			*getMemU32Ptr(0x5D4594, 1193332) = 1;
		}
		result = v2;
	}
	return result;
}

//----- (00486FE0) --------------------------------------------------------
uint32_t* sub_486FE0(int a1) {
	uint32_t* v1; // esi

	v1 = calloc(1, 0x58u);
	memset(v1, 0, 0x58u);
	sub_425770(v1);
	v1[4] = 0;
	v1[3] = a1;
	if (!(*(int (**)(uint32_t*))(a1 + 20))(v1)) {
		return v1;
	}
	if (v1) {
		sub_487030(v1);
	}
	return 0;
}

//----- (00487030) --------------------------------------------------------
void sub_487030(void* lpMem) {
	(*(void (**)(void*))(*((uint32_t*)lpMem + 3) + 24))(lpMem);
	*(uint32_t*)(*((uint32_t*)lpMem + 3) + 12) &= 0xFFFFFFFE;
	free(lpMem);
}

//----- (00487050) --------------------------------------------------------
void sub_487050(uint32_t* a1) { nox_common_list_append_4258E0(*(int*)&dword_587000_155144, a1); }

//----- (00487070) --------------------------------------------------------
void sub_487070(void* lpMem) {
	sub_487090((uint32_t**)lpMem);
	sub_487030(lpMem);
	*getMemU32Ptr(0x5D4594, 1193332) = 0;
}

//----- (00487090) --------------------------------------------------------
void sub_487090(uint32_t** a1) { nox_common_list_remove_425920(a1); }

//----- (004870A0) --------------------------------------------------------
void sub_4870A0() {
	int* v1; // edi
	int* v2; // esi
	int* v3; // [esp+4h] [ebp-4h]

	v1 = sub_4870E0((int*)&v3);
	if (v1) {
		do {
			v2 = sub_487100(&v3);
			sub_487070(v1);
			v1 = v2;
		} while (v2);
	}
}

//----- (004870E0) --------------------------------------------------------
int* sub_4870E0(int* a1) {
	int* result; // eax

	result = nox_common_list_getFirstSafe_425890(*(int**)&dword_587000_155144);
	*a1 = (int)result;
	return result;
}

//----- (00487100) --------------------------------------------------------
int* sub_487100(int** a1) {
	if (*a1) {
		*a1 = nox_common_list_getNextSafe_4258A0(*a1);
	}
	return *a1;
}

//----- (00487150) --------------------------------------------------------
uint32_t* sub_487150(int a1, const void* a2) {
	int v2;       // edi
	uint32_t* v3; // esi
	uint32_t* v4; // eax
	int v6;       // [esp+8h] [ebp-4h]

	v2 = a1;
	if (a1 == -1) {
		v2 = 0;
	}
	sub_487360(v2, (int**)&a1, &v6);
	if (!a1) {
		return 0;
	}
	v3 = *(uint32_t**)(a1 + 4 * v6 + 24);
	if (!v3) {
		v4 = sub_4871C0(a1, v6, a2);
		v3 = v4;
		if (!v4) {
			return 0;
		}
		v4[47] = v2;
		sub_487310(v4);
	}
	++v3[4];
	return v3;
}

//----- (004871C0) --------------------------------------------------------
uint32_t* sub_4871C0(int a1, int a2, const void* a3) {
	int v3;       // ebp
	uint32_t* v4; // esi

	v3 = *(uint32_t*)(a1 + 12);
	v4 = calloc(1, 0x108u);
	memset(v4, 0, 0x108u);
	sub_425770(v4);
	v4[6] = a2;
	v4[5] = a1;
	v4[4] = 0;
	++*(uint32_t*)(a1 + 16);
	*(uint32_t*)(a1 + 4 * a2 + 24) = v4;
	v4[64] = *(uint32_t*)(v3 + 36);
	nox_common_list_clear_425760(v4 + 50);
	sub_4864A0(v4 + 22);
	v4[53] = 0;
	v4[56] = 33;
	v4[60] = 0;
	v4[58] = 0;
	v4[62] = 0;
	v4[54] = sub_4873C0;
	v4[57] = 0;
	v4[61] = 0;
	v4[59] = 0;
	v4[63] = 0;
	nullsub_10(v4 + 15);
	nullsub_10(v4 + 8);
	if (a3) {
		sub_487590((int)v4, a3);
	}
	if (!(*(int (**)(uint32_t*))(v3 + 28))(v4)) {
		return v4;
	}
	if (v4) {
		sub_4872C0(v4);
	}
	return 0;
}
// 487CF0: using guessed type void  nullsub_10(uint32_t);

//----- (004872C0) --------------------------------------------------------
void sub_4872C0(void* lpMem) {
	int v1; // eax
	int v2; // ecx

	sub_487910((int)lpMem, -1);
	(*(void (**)(void*))(*(uint32_t*)(*((uint32_t*)lpMem + 5) + 12) + 32))(lpMem);
	*(uint32_t*)(*((uint32_t*)lpMem + 5) + 4 * *((uint32_t*)lpMem + 6) + 24) = 0;
	v1 = *((uint32_t*)lpMem + 5);
	v2 = *(uint32_t*)(v1 + 16) - 1;
	*(uint32_t*)(v1 + 16) = v2;
	if (v2 < 0) {
		*(uint32_t*)(*((uint32_t*)lpMem + 5) + 16) = 0;
	}
	free(lpMem);
}

//----- (00487310) --------------------------------------------------------
int sub_487310(uint32_t* a1) {
	int result; // eax

	++*(uint32_t*)((uint32_t)dword_587000_155144 + 24);
	nox_common_list_append_4258E0((uint32_t)dword_587000_155144 + 12, a1);
	result = *(uint32_t*)((uint32_t)dword_587000_155144 + 24) - 1;
	*(uint32_t*)((uint32_t)dword_587000_155144 + 24) = result;
	if (result < 0) {
		*(uint32_t*)((uint32_t)dword_587000_155144 + 24) = 0;
	}
	return result;
}

//----- (00487360) --------------------------------------------------------
int* sub_487360(int a1, int** a2, int* a3) {
	int* result; // eax
	int i;       // esi
	int v5;      // ecx
	int* v6;     // [esp+4h] [ebp-4h]

	result = sub_4870E0((int*)&v6);
	for (i = a1; result; result = sub_487100(&v6)) {
		v5 = result[5];
		if (i < v5) {
			break;
		}
		i -= v5;
	}
	*a2 = result;
	if (result) {
		result = a3;
		*a3 = i;
	} else {
		*a3 = -1;
	}
	return result;
}

//----- (004873C0) --------------------------------------------------------
int sub_4873C0(int a3) {
	int v1;           // esi
	long long v3;     // rax
	unsigned int v4;  // ecx
	int v5;           // ebp
	bool v6;          // cf
	unsigned int v7;  // ebx
	unsigned int v8;  // ecx
	int v9;           // edi
	unsigned int v10; // eax
	uint32_t* v11;    // ebx
	int v12;          // edi
	int v13;          // ebp
	uint32_t* v14;    // esi
	int v15;          // [esp+10h] [ebp-Ch]
	int v16;          // [esp+18h] [ebp-4h]
	int v17;          // [esp+20h] [ebp+4h]

	v1 = a3;
	if (*(uint32_t*)(a3 + 212)) {
		return -2146304000;
	}
	v3 = nox_platform_get_ticks();
	v4 = *(uint32_t*)(a3 + 248);
	v5 = v3;
	v6 = (unsigned int)v3 < v4;
	v7 = v3 - v4;
	v8 = *(uint32_t*)(a3 + 224);
	v16 = HIDWORD(v3);
	v9 = HIDWORD(v3) - (v6 + *(uint32_t*)(a3 + 252));
	v10 = *(uint32_t*)(a3 + 228);
	if (__PAIR64__(v9, v7) >= __PAIR64__(v10, v8)) {
		*(uint32_t*)(a3 + 232) = v7;
		*(uint32_t*)(a3 + 236) = v9;
		if (*(uint64_t*)(a3 + 240) > 10 * __PAIR64__(v10, v8)) {
			*(uint32_t*)(a3 + 240) = 0;
			*(uint32_t*)(a3 + 244) = 0;
		}
		if (__PAIR64__(v9, v7) > *(uint64_t*)(a3 + 240)) {
			*(uint32_t*)(a3 + 240) = v7;
			*(uint32_t*)(a3 + 244) = v9;
		}
		v11 = (uint32_t*)(a3 + 88);
		v15 = a3 + 88;
		sub_486520((unsigned int*)(a3 + 88));
		if (*(uint32_t*)(a3 + 184)) {
			sub_486520(*(unsigned int**)(a3 + 184));
		}
		if (!*(uint32_t*)(a3 + 184) || !(v17 = sub_486550(*(uint8_t**)(a3 + 184)))) {
			v17 = sub_486550((uint8_t*)(v1 + 88));
		}
		*(uint32_t*)(v1 + 248) = v5;
		*(uint32_t*)(v1 + 252) = v16;
		v12 = *(uint32_t*)(v1 + 200);
		if (v12 != v1 + 200) {
			do {
				v13 = *(uint32_t*)v12;
				if (*(uint8_t*)(v12 + 124) & 1 && *(uint32_t*)(v12 + 288)) {
					if ((sub_486520((unsigned int*)(v12 + 16)), v17) || sub_486550((uint8_t*)(v12 + 16)) ||
						*(uint32_t*)(v12 + 116) && sub_486550(*(uint8_t**)(v12 + 116)) ||
						*(uint32_t*)(v12 + 112) && sub_486550(*(uint8_t**)(v12 + 112))) {
						sub_4BD840(v12);
						(*(void (**)(int))(*(uint32_t*)(v12 + 172) + 32))(v12);
					}
				}
				v12 = v13;
			} while (v13 != v1 + 200);
			v11 = (uint32_t*)v15;
		}
		v14 = *(uint32_t**)(v1 + 184);
		if (v14) {
			sub_486620(v14);
		}
		sub_486620(v11);
	}
	return 0;
}

//----- (00487590) --------------------------------------------------------
int sub_487590(int a1, const void* a2) {
	int result; // eax

	result = a1;
	memcpy((void*)(a1 + 60), a2, 0x1Cu);
	return result;
}

//----- (004875B0) --------------------------------------------------------
int* sub_4875B0(int* a1) {
	int* result; // eax

	result = nox_common_list_getFirstSafe_425890((int*)((uint32_t)dword_587000_155144 + 12));
	*a1 = (int)result;
	return result;
}

//----- (004875D0) --------------------------------------------------------
int* sub_4875D0(int** a1) {
	if (*a1) {
		*a1 = nox_common_list_getNextSafe_4258A0(*a1);
	}
	return *a1;
}

//----- (004875F0) --------------------------------------------------------
int sub_4875F0() {
	int* v0;    // edi
	int* v1;    // esi
	int result; // eax
	int* v3;    // [esp+4h] [ebp-4h]

	++*(uint32_t*)((uint32_t)dword_587000_155144 + 24);
	v0 = sub_4875B0((int*)&v3);
	if (v0) {
		do {
			v1 = sub_4875D0(&v3);
			sub_487680(v0);
			v0 = v1;
		} while (v1);
	}
	result = *(uint32_t*)((uint32_t)dword_587000_155144 + 24) - 1;
	*(uint32_t*)((uint32_t)dword_587000_155144 + 24) = result;
	if (result < 0) {
		*(uint32_t*)((uint32_t)dword_587000_155144 + 24) = 0;
	}
	return result;
}

//----- (00487680) --------------------------------------------------------
void sub_487680(void* lpMem) {
	sub_4876A0((uint32_t**)lpMem);
	sub_4872C0(lpMem);
}

//----- (004876A0) --------------------------------------------------------
void* sub_4876A0(uint32_t** a1) {
	void* result; // eax

	++*(uint32_t*)((uint32_t)dword_587000_155144 + 24);
	nox_common_list_remove_425920(a1);
	result = (void*)(*(uint32_t*)((uint32_t)dword_587000_155144 + 24) - 1);
	*(uint32_t*)((uint32_t)dword_587000_155144 + 24) = result;
	if ((int)result < 0) {
		result = *(void**)&dword_587000_155144;
		*(uint32_t*)((uint32_t)dword_587000_155144 + 24) = 0;
	}
	return result;
}

//----- (00487750) --------------------------------------------------------
uint32_t* sub_487750(int a1) {
	uint32_t* v1; // eax
	uint32_t* v2; // esi

	if (*(int*)(a1 + 192) >= *(int*)(a1 + 196)) {
		return 0;
	}
	v1 = sub_4BD720(a1);
	v2 = v1;
	if (!v1) {
		return 0;
	}
	sub_486E30(a1, v1);
	return v2;
}

//----- (00487790) --------------------------------------------------------
int sub_487790(int a1, int a2) {
	int v2; // esi
	int v3; // edi

	v2 = 0;
	if (sub_487750(a1)) {
		v3 = a2;
		do {
			++v2;
			--v3;
		} while (v3 && sub_487750(a1));
	}
	return v2;
}

//----- (004877D0) --------------------------------------------------------
int* sub_4877D0(int a1, int* a2) {
	int* result; // eax

	result = nox_common_list_getFirstSafe_425890((int*)(a1 + 200));
	*a2 = (int)result;
	return result;
}

//----- (004877F0) --------------------------------------------------------
int* sub_4877F0(int** a1) {
	if (*a1) {
		*a1 = nox_common_list_getNextSafe_4258A0(*a1);
	}
	return *a1;
}

//----- (00487810) --------------------------------------------------------
int* sub_487810(int a1, int a2) {
	unsigned int v2; // esi
	int v3;          // edi
	int* v4;         // ebp
	int* result;     // eax
	int v6;          // ecx
	unsigned int v7; // edx
	int v8;          // [esp+10h] [ebp-8h]
	int* v9;         // [esp+14h] [ebp-4h]

	v2 = -1;
	if (a2 == -1) {
		a2 = 1;
	}
	v3 = 127;
	v4 = 0;
	v8 = 127;
	v9 = 0;
	for (result = sub_4877D0(a1, &a1); result; result = sub_4877F0((int**)&a1)) {
		if (result[3] == a2) {
			if (!(result[31] & 0x15)) {
				return result;
			}
			v6 = result[30];
			if (result[31] & 1) {
				if (v6 >= v3) {
					if (v6 == v3) {
						v7 = result[45];
						if (v7 < v2 && v2 - v7 >= 0x666) {
							v3 = result[30];
							v4 = result;
							v2 = result[45];
						}
					}
				} else {
					v2 = result[45];
					v3 = result[30];
					v4 = result;
				}
			} else if (v6 < v8) {
				v8 = result[30];
				v9 = result;
			}
		}
	}
	result = v9;
	if (!v9 || v8 > v3) {
		result = v4;
	}
	return result;
}

//----- (00487910) --------------------------------------------------------
int sub_487910(int a1, int a2) {
	int* v2; // edi
	int v3;  // ebx
	int* v4; // esi

	v2 = sub_4877D0(a1, &a1);
	if (!v2) {
		return 0;
	}
	v3 = a2;
	do {
		v4 = sub_4877F0((int**)&a1);
		if (v3 == -1 || v2[3] == v3) {
			sub_4BDA60(v2);
		}
		v2 = v4;
	} while (v4);
	return 0;
}

//----- (00487970) --------------------------------------------------------
int* sub_487970(int a1, int a2) {
	int* result; // eax
	int* v3;     // edi
	int v4;      // ebx
	int* v5;     // esi

	result = sub_4877D0(a1, &a1);
	v3 = result;
	if (result) {
		v4 = a2;
		do {
			result = sub_4877F0((int**)&a1);
			v5 = result;
			if (v4 == -1 || v3[3] == v4) {
				result = (int*)sub_4BDA80((int)v3);
			}
			v3 = v5;
		} while (v5);
	}
	return result;
}

//----- (00487C30) --------------------------------------------------------
void sub_487C30(uint32_t* a1) {
	*a1 = 0;
	a1[1] = 0;
	a1[5] = 0;
	a1[6] = 0;
	nox_common_list_clear_425760(a1 + 2);
}

//----- (00487C50) --------------------------------------------------------
int sub_487C50(int a1, uint32_t* a2) {
	int result; // eax

	nox_common_list_append_4258E0(a1 + 8, a2);
	result = a2[4] + *(uint32_t*)(a1 + 4);
	*(uint32_t*)(a1 + 4) = result;
	a2[5] = a1;
	return result;
}

//----- (00487C80) --------------------------------------------------------
int sub_487C80(int a1) { return nox_common_list_getNext_425940((int*)(a1 + 8)); }

//----- (00487D00) --------------------------------------------------------
int sub_487D00(uint32_t* a1) {
	int v1;     // edx
	int result; // eax

	v1 = a1[1];
	result = a1[2] * a1[3] * a1[4];
	a1[5] = result;
	if (v1 == 1) {
		result >>= 2;
		a1[5] = result;
	}
	return result;
}

//----- (00487D30) --------------------------------------------------------
uint32_t* sub_487D30(uint32_t* a1, int a2, int a3) {
	uint32_t* result; // eax

	a1[4] = a3;
	a1[3] = a2;
	result = sub_425770(a1);
	a1[5] = 0;
	return result;
}

//----- (00487D60) --------------------------------------------------------
int sub_487D60(int a1) {
	int result; // eax

	result = a1;
	*(uint32_t*)(a1 + 20) = 0;
	return result;
}

//----- (00487D70) --------------------------------------------------------
int nox_xxx_wndEditProc_487D70_key(uint32_t* a1, int v4, int a3, int a4) {
	switch (a3) {
	case 1:
	case 58:
	case 59:
	case 60:
	case 61:
	case 62:
	case 63:
	case 64:
	case 65:
	case 66:
	case 67:
	case 68:
	case 87:
	case 88:
	case 199:
	case 201:
	case 207:
	case 209:
		return 0;
	case 14:
	case 211:
		if (a3 != 14 && (nox_strman_get_lang_code() == 6 || nox_strman_get_lang_code() == 8)) {
			break;
		}
		if (a4 != 2) {
			return 1;
		}
		if (!*(uint16_t*)(v4 + 1054)) {
			short v6 = *(uint16_t*)(v4 + 1052);
			if (!v6) {
				return 1;
			}
			unsigned short v7 = v6 - 1;
			*(uint16_t*)(v4 + 1052) = v7;
			*(uint16_t*)(v4 + 2 * v7) = 0;
			return 1;
		}
		break;
	case 15:
	case 205:
	case 208:
		if (a4 != 2) {
			return 1;
		}
		nox_xxx_wndRetNULL_46A8A0();
		return 1;
	case 28:
	case 156:
		if (a4 != 2 || *(uint32_t*)(v4 + 1044)) {
			return 1;
		}
		nox_window_call_field_94(a1[13], 16415, (int)a1, 0);
		return 1;
	case 200:
	case 203:
		if (a4 != 2) {
			return 1;
		}
		nox_xxx_wndRetNULL_0_46A8B0();
		return 1;
	default:
		break;
	}
	if (nox_strman_get_lang_code() == 6 || nox_strman_get_lang_code() == 8) {
		if (!*(uint32_t*)(v4 + 1036) && !*(uint32_t*)(v4 + 1032) && !*(uint32_t*)(v4 + 1028)) {
			wchar2_t* v8 = nox_input_getStringBuffer_57011C();
			nox_wcscpy((wchar2_t*)(v4 + 512), (const wchar2_t*)v8);
			nox_input_freeStringBuffer_57011C(v8);
			*(uint16_t*)(v4 + 1054) = nox_wcslen((const wchar2_t*)(v4 + 512));
			if (0) { // if (!nox_xxx_string_5702B4(*(uint32_t***)&dword_5d4594_1193348))
				return 1;
			}
			// nox_xxx_string_570392(*(uint32_t***)&dword_5d4594_1193348);
			nox_window_set_hidden(*(uint32_t*)(v4 + 1048), 1);
			return 1;
		}
	}
	unsigned short v9 = nox_input_scanCodeToAlpha_47F950(a3);
	if (!v9 || a4 != 2) {
		return 1;
	}
	int v10 = 0;
	if (*(uint32_t*)(v4 + 1028)) {
		v10 = iswdigit(v9);
		if (!v10) {
			return 1;
		}
	} else if (*(uint32_t*)(v4 + 1032)) {
		v10 = iswalnum(v9);
		if (!v10) {
			return 1;
		}
	}
	int v11 = *(unsigned short*)(v4 + 1052);
	if ((unsigned short)v11 >= *(short*)(v4 + 1040) - 1) {
		return 1;
	}
	*(uint16_t*)(v4 + 2 * v11) = v9;
	*(uint16_t*)(v4 + 2 * (unsigned short)++*(uint16_t*)(v4 + 1052)) = 0;
	return 1;
}
int nox_xxx_wndEditProc_487D70(nox_window* a1p, int a2, int a3, int a4) {
	uint32_t* a1 = a1p;
	int v4;  // esi
			 //	int result;          // eax
			 //	short v6;          // ax
			 //	unsigned short v7; // ax
			 //	unsigned short v9; // di
			 //	int v10;             // eax
			 //	int v11;             // eax
	int v12; // eax
	int v13; // eax
	int v14; // eax
	int v15; // eax
	int v16; // [esp-10h] [ebp-20h]
	int v17; // [esp-10h] [ebp-20h]

	v4 = a1[8];
	switch (a2) {
	case 7:
		a1[9] |= 2u;
		nox_xxx_windowFocus_46B500((int)a1);
		return 1;
	case 8:
		v15 = a1[11];
		if (v15 & 0x100) {
			nox_window_call_field_94(a1[13], 0x4000, (int)a1, 0);
		}
		return 1;
	case 17:
		v12 = a1[11];
		if (!(v12 & 0x100)) {
			return 1;
		}
		v16 = a1[13];
		a1[9] |= 2u;
		nox_window_call_field_94(v16, 16389, (int)a1, 0);
		nox_xxx_windowFocus_46B500((int)a1);
		return 1;
	case 18:
		v13 = a1[11];
		if (v13 & 0x100) {
			v14 = a1[9];
			LOBYTE(v14) = v14 & 0xFD;
			v17 = a1[13];
			a1[9] = v14;
			nox_window_call_field_94(v17, 16390, (int)a1, 0);
		}
		return 1;
	case 21:
		return nox_xxx_wndEditProc_487D70_key(a1, v4, a3, a4);
	default:
		return 0;
	}
}

//----- (00488160) --------------------------------------------------------
int nox_xxx_wndEditDrawNoImage_488160(int a1, int a2) {
	int v2;          // edi
	int v3;          // edx
	int v4;          // ebx
	int v5;          // ecx
	int v6;          // kr04_4
	int v7;          // ebp
	wchar2_t* v8;     // ecx
	signed short v9; // ax
	size_t v10;      // edx
	size_t v11;      // eax
	short* v12;      // edi
	int i;           // ecx
	uint32_t* v14;   // eax
	int v15;         // eax
	int v16;         // ebx
	bool v17;        // zf
	int v19;         // [esp-28h] [ebp-260h]
	int v20;         // [esp-14h] [ebp-24Ch]
	int v21;         // [esp+10h] [ebp-228h]
	int v22;         // [esp+14h] [ebp-224h]
	wchar2_t* v23;    // [esp+18h] [ebp-220h]
	int v24;         // [esp+1Ch] [ebp-21Ch]
	wchar2_t* v25;    // [esp+20h] [ebp-218h]
	short* v26;      // [esp+24h] [ebp-214h]
	int yTop;        // [esp+28h] [ebp-210h]
	int xLeft;       // [esp+2Ch] [ebp-20Ch]
	int v29;         // [esp+30h] [ebp-208h]
	int v30;         // [esp+34h] [ebp-204h]
	short v31[256];  // [esp+38h] [ebp-200h]

	v2 = a1;
	v3 = *(uint32_t*)(a2 + 20);
	v25 = *(wchar2_t**)(a2 + 28);
	v23 = *(wchar2_t**)(a1 + 32);
	v22 = 0;
	v24 = 0;
	*((uint32_t*)v23 + 261) = 0;
	v30 = v3;
	v26 = v31;
	nox_client_wndGetPosition_46AA60((uint32_t*)a1, &xLeft, &yTop);
	v4 = xLeft;
	v21 = *(uint32_t*)(a1 + 8);
	v5 = nox_xxx_guiFontHeightMB_43F320(*(uint32_t*)(a2 + 200));
	v6 = *(uint32_t*)(a1 + 12);
	v29 = v5;
	v7 = yTop + v6 / 2 - v5 / 2;
	if (((*(uint32_t*)(a1 + 4) >> 13) & 1) == 1) {
		nox_draw_enableTextSmoothing_43F670(1);
	}
	if ((*(uint32_t*)(a1 + 4) >> 3) & 1) {
		if (*(uint8_t*)a2 & 2) {
			v25 = *(wchar2_t**)(a2 + 36);
		}
	} else {
		v30 = *(uint32_t*)(a2 + 44);
	}
	if (*(uint16_t*)(a2 + 72)) {
		nox_xxx_drawGetStringSize_43F840(*(uint32_t*)(a2 + 200), (unsigned short*)(a2 + 72), &v22, 0, 0);
		nox_xxx_drawSetTextColor_434390(*(uint32_t*)(a2 + 68));
		nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), (uint16_t*)(a2 + 72), v4 + 2, v7, v21, 0);
		v4 += v22 + 6;
		v21 += -6 - v22;
	}
	v8 = v23;
	v9 = v23[521];
	if (v9 > 0 && v21 > v9) {
		v21 = v9;
		v4 = xLeft + *(uint32_t*)(a1 + 8) - v9;
	}
	if (v30 != 0x80000000) {
		nox_client_drawSetColor_434460(v30);
		nox_client_drawRectFilledOpaque_49CE30(v4 + 1, yTop + 1, v21 - 2, *(uint32_t*)(a1 + 12) - 2);
		v8 = v23;
	}
	if (v25 != (wchar2_t*)0x80000000) {
		nox_client_drawSetColor_434460((int)v25);
		nox_client_drawBorderLines_49CC70(v4, yTop, v21, *(uint32_t*)(a1 + 12));
		v8 = v23;
	}
	if (*(uint32_t*)(a2 + 68) != 0x80000000) {
		if (*((uint32_t*)v8 + 256)) {
			v10 = nox_wcslen(v8);
			v11 = 0;
			if ((int)v10 > 0) {
				memset32(v31, 2752554, v10 >> 1);
				v12 = &v31[2 * (v10 >> 1)];
				for (i = v10 & 1; i; --i) {
					*v12 = 42;
					++v12;
				}
				v2 = a1;
				v11 = v10;
			}
			v31[v11] = 0;
		} else {
			nox_wcscpy((wchar2_t*)v31, v8);
		}
		nox_xxx_drawGetStringSize_43F840(*(uint32_t*)(a2 + 200), (unsigned short*)v31, &v22, 0, 0);
		v19 = *(uint32_t*)(a2 + 200);
		v25 = v23 + 256;
		nox_xxx_drawGetStringSize_43F840(v19, v23 + 256, &v24, 0, 0);
		if (((*(uint32_t*)(v2 + 4) >> 14) & 1) == 1 && v24 + v22 > 0 && v21 >= 10 && v24 + v22 + 10 > v21) {
			do {
				v20 = *(uint32_t*)(a2 + 200);
				++v26;
				nox_xxx_drawGetStringSize_43F840(v20, (unsigned short*)v26, &v22, 0, 0);
			} while (v24 + v22 + 10 > v21);
		}
		v14 = (uint32_t*)*((uint32_t*)v23 + 262);
		if (v14) {
			nox_window_setPos_46A9B0(v14, v4 + v22, v29 + v7);
		}
		nox_xxx_drawSetTextColor_434390(*(uint32_t*)(a2 + 68));
		nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), v26, v4 + 5, v7, 0, 0);
		v15 = nox_color_rgb_4344A0(192, 0, 192);
		nox_xxx_drawSetTextColor_434390(v15);
		nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), v25, v4 + v22 + 5, v7, 0, 0);
		v16 = v4 + v22 + v24 + 5;
		if (v2 == nox_xxx_wndGetFocus_46B4F0()) {
			v17 = ((*getMemU8Ptr(0x5D4594, 1193344))++ & 8) == 0;
			if (!v17) {
				nox_client_drawSetColor_434460(*(uint32_t*)(a2 + 68));
				nox_client_drawRectFilledOpaque_49CE30(v16, v7, 2, v29);
			}
		}
	}
	nox_draw_enableTextSmoothing_43F670(0);
	return 1;
}
// 488160: using guessed type wchar2_t var_200[256];

//----- (00488500) --------------------------------------------------------
nox_window* nox_gui_newEntryField_488500(nox_window* a1p, int a2, int a3, int a4, int a5, int a6, int a7, wchar2_t* a8) {
	int a1 = a1p;
	uint32_t* v8;     // esi
	bool v9;          // cc
	int* v10;         // ebx
	uint32_t* result; // eax
	uint32_t* v12;    // eax
	int v13;          // eax
	char v14[56];     // [esp+14h] [ebp-184h]
	char v15[332];    // [esp+4Ch] [ebp-14Ch]

	if (*(uint8_t*)(a7 + 8) & 0x80) {
		v8 = nox_window_new(a1, a2, a3, a4, a5, a6, nox_xxx_wndEditProcPre_488710);
		nox_xxx_wndEdit_488830((int)v8);
		if (!v8) {
			return v8;
		}
		if (!*(uint32_t*)(a7 + 16)) {
			*(uint32_t*)(a7 + 16) = v8;
		}
		nox_gui_windowCopyDrawData_46AF80((int)v8, (const void*)a7);
		memset(a8, 0, 0x100u);
		memset(a8 + 256, 0, 0x100u);
		a8[526] = nox_wcslen(a8);
		v9 = (short)a8[520] < 256;
		a8[527] = 0;
		*((uint32_t*)a8 + 261) = 0;
		if (!v9) {
			a8[520] = 256;
		}
		v10 = (int*)calloc(1, 0x420u);
		memcpy(v10, a8, 0x420u);
		v8[8] = v10;
		if (nox_strman_get_lang_code() != 8 && nox_strman_get_lang_code() != 6) {
			result = v8;
			v10[262] = 0;
			return result;
		}
		memset(v15, 0, sizeof(v15));
		*(uint32_t*)&v15[24] = 0;
		*(uint32_t*)&v15[28] = *(uint32_t*)(a7 + 68);
		*(uint32_t*)&v15[36] = 0x80000000;
		*(uint32_t*)&v15[52] = 0x80000000;
		*(uint32_t*)&v15[68] = *(uint32_t*)&v15[28];
		memset(v14, 0, sizeof(v14));
		*(uint32_t*)&v14[8] = 1;
		*(uint32_t*)&v14[12] = 1;
		*(uint32_t*)&v15[48] = 0;
		*(uint16_t*)&v14[2] = 0x0a;
		*(uint16_t*)&v14[0] = 0x80;
		*(uint32_t*)&v14[4] = 0;
		*(uint32_t*)&v14[16] = 0;
		*(uint16_t*)&v15[72] = 0;
		*(uint32_t*)&v15[8] = 288;
		v12 = nox_gui_newScrollListBox_4A4310(0, 17584, 0, a6, 110, 119, (int)v15, (short*)v14);
		v10[262] = (int)v12;
		if (v12) {
			nox_xxx_wndClearFlag_46AD80((int)v12, 128);
			nox_xxx_wndListboxInit_4A3C00(v10[262], *(uint32_t*)(v10[262] + 32));
			v13 = nox_color_rgb_4344A0(0, 0, 0);
			nox_xxx_wndSetRectColor2MB_46AFE0(v10[262], v13);
			return v8;
		}
	}
	return 0;
}

//----- (00488710) --------------------------------------------------------
int nox_xxx_wndEditProcPre_488710(int a1, unsigned int a2, wchar2_t* a3, int a4) {
	int v3; // esi
	int v4; // eax
	int v6; // eax
	int v7; // eax

	v3 = *(uint32_t*)(a1 + 32);
	if (a2 > 0x401D) {
		if (a2 == 16414) {
			nox_wcsncpy((wchar2_t*)v3, a3, 0xFFu);
			*(uint16_t*)(v3 + 510) = 0;
			*(uint16_t*)(v3 + 1052) = nox_wcslen((const wchar2_t*)v3);
			*(uint16_t*)(v3 + 512) = 0;
			*(uint16_t*)(v3 + 1054) = 0;
		}
		return 0;
	}
	if (a2 == 16413) {
		return *(uint32_t*)(a1 + 32);
	}
	if (a2 == 2) {
		nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)(v3 + 1048));
		free(*(void**)(a1 + 32));
		return 0;
	}
	if (a2 != 23) {
		return 0;
	}
	if (a3) {
		dword_5d4594_1193352 = a1;
		nox_input_enableTextEdit_5700CA();
		v6 = *(uint32_t*)(a1 + 36);
		LOBYTE(v6) = v6 | 6;
		*(uint32_t*)(a1 + 36) = v6;
	} else {
		nox_input_disableTextEdit_5700F6();
		dword_5d4594_1193352 = 0;
		v4 = *(uint32_t*)(a1 + 36);
		LOBYTE(v4) = v4 & 0xF9;
		*(uint32_t*)(a1 + 36) = v4;
		nox_window_set_hidden(*(uint32_t*)(v3 + 1048), 1);
		*(uint16_t*)(v3 + 512) = 0;
		*(uint16_t*)(v3 + 1054) = 0;
	}
	v7 = nox_xxx_wndGetID_46B0A0((int*)a1);
	nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16387, (int)a3, v7);
	return 1;
}

//----- (00488830) --------------------------------------------------------
int nox_xxx_wndEdit_488830(int a1) {
	int result; // eax

	result = a1;
	if (a1) {
		if ((signed char)*(uint8_t*)(a1 + 4) >= 0) {
			result = nox_window_set_all_funcs((uint32_t*)a1, nox_xxx_wndEditProc_487D70,
											  nox_xxx_wndEditDrawNoImage_488160, 0);
		} else {
			result = nox_window_set_all_funcs((uint32_t*)a1, nox_xxx_wndEditProc_487D70,
											  nox_xxx_wndEditDrawWithImage_488870, 0);
		}
	}
	return result;
}

//----- (00488870) --------------------------------------------------------
int nox_xxx_wndEditDrawWithImage_488870(int a1, int a2) {
	int v2;         // edi
	int v3;         // ebp
	int v4;         // ecx
	int v5;         // kr04_4
	int v6;         // ebx
	int v7;         // eax
	size_t v8;      // edx
	size_t v9;      // eax
	short* v10;     // edi
	int i;          // ecx
	int v12;        // eax
	bool v13;       // zf
	int v15;        // [esp-14h] [ebp-244h]
	int xLeft;      // [esp+10h] [ebp-220h]
	int v17;        // [esp+14h] [ebp-21Ch]
	short* v18;     // [esp+18h] [ebp-218h]
	int v19;        // [esp+1Ch] [ebp-214h]
	wchar2_t* v20;   // [esp+20h] [ebp-210h]
	int v21;        // [esp+24h] [ebp-20Ch]
	wchar2_t* v22;   // [esp+28h] [ebp-208h]
	int v23;        // [esp+2Ch] [ebp-204h]
	short v24[256]; // [esp+30h] [ebp-200h]

	v2 = a1;
	v20 = *(wchar2_t**)(a2 + 24);
	v22 = *(wchar2_t**)(a1 + 32);
	*((uint32_t*)v22 + 261) = 0;
	v18 = v24;
	nox_client_wndGetPosition_46AA60((uint32_t*)a1, &xLeft, &v21);
	v3 = *(uint32_t*)(a1 + 8);
	v4 = nox_xxx_guiFontHeightMB_43F320(*(uint32_t*)(a2 + 200));
	v5 = *(uint32_t*)(a1 + 12);
	v23 = v4;
	v6 = v21 + v5 / 2 - v4 / 2;
	if (((*(uint32_t*)(a1 + 4) >> 13) & 1) == 1) {
		nox_draw_enableTextSmoothing_43F670(1);
	}
	if ((*(uint32_t*)(a1 + 4) >> 3) & 1) {
		v7 = (int)v20;
	} else {
		v7 = *(uint32_t*)(a2 + 48);
	}
	if (v7) {
		nox_client_drawImageAt_47D2C0(v7, xLeft, v21);
	}
	nox_xxx_drawSetTextColor_434390(*(uint32_t*)(a2 + 68));
	if (*(uint16_t*)(a2 + 72)) {
		nox_xxx_drawGetStringSize_43F840(*(uint32_t*)(a2 + 200), (unsigned short*)(a2 + 72), &v17, 0, 0);
		nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), (uint16_t*)(a2 + 72), xLeft + 2, v6, v3, 0);
		v3 += -6 - v17;
		xLeft += v17 + 6;
	}
	if (*(uint32_t*)(a2 + 68) != 0x80000000) {
		if (*((uint32_t*)v22 + 256)) {
			v8 = nox_wcslen(v22);
			v9 = 0;
			if ((int)v8 > 0) {
				memset32(v24, 2752554, v8 >> 1);
				v10 = &v24[2 * (v8 >> 1)];
				for (i = v8 & 1; i; --i) {
					*v10 = 42;
					++v10;
				}
				v2 = a1;
				v9 = v8;
			}
			v24[v9] = 0;
		} else {
			nox_wcscpy((wchar2_t*)v24, v22);
		}
		nox_xxx_drawGetStringSize_43F840(*(uint32_t*)(a2 + 200), (unsigned short*)v24, &v17, 0, 0);
		v20 = v22 + 256;
		nox_xxx_drawGetStringSize_43F840(*(uint32_t*)(a2 + 200), v22 + 256, &v19, 0, 0);
		if (((*(uint32_t*)(v2 + 4) >> 14) & 1) == 1 && v17 + v19 > 0 && v17 + v19 + 10 > v3) {
			do {
				v15 = *(uint32_t*)(a2 + 200);
				++v18;
				nox_xxx_drawGetStringSize_43F840(v15, (unsigned short*)v18, &v17, 0, 0);
			} while (v19 + v17 + 10 > v3);
		}
		nox_xxx_drawSetTextColor_434390(*(uint32_t*)(a2 + 68));
		nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), v18, xLeft + 5, v6, v3, 0);
		v12 = nox_color_rgb_4344A0(192, 0, 192);
		nox_xxx_drawSetTextColor_434390(v12);
		nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), v20, v17 + xLeft + 5, v6, v3, 0);
		xLeft += v17 + v19 + 5;
		if (v2 == nox_xxx_wndGetFocus_46B4F0()) {
			v13 = ((*getMemU8Ptr(0x5D4594, 1193344))++ & 8) == 0;
			if (!v13) {
				nox_client_drawSetColor_434460(*(uint32_t*)(a2 + 68));
				nox_client_drawRectFilledOpaque_49CE30(xLeft, v6, 2, v23);
			}
		}
	}
	nox_draw_enableTextSmoothing_43F670(0);
	return 1;
}
// 488870: using guessed type wchar2_t var_200[256];

//----- (00488B60) --------------------------------------------------------
int sub_488B60() {
	int** v0; // eax

	v0 = (int**)calloc(1, 4u);
	if (v0) {
		dword_5d4594_1193348 = v0;
		// sub_570142(*(uint32_t***)&dword_5d4594_1193348, 11);
	} else {
		dword_5d4594_1193348 = 0;
		// sub_570142(0, 11);
	}
	return 1;
}

//----- (00488BA0) --------------------------------------------------------
int sub_488BA0() {
	void* v0; // esi

	v0 = *(void**)&dword_5d4594_1193348;
	if (dword_5d4594_1193348) {
		free(v0);
	}
	dword_5d4594_1193348 = 0;
	return 1;
}

//----- (00488BD0) --------------------------------------------------------
void nox_xxx_onChar_488BD0(unsigned short a1) {
	int v2; // esi
	int v3; // eax

	if (dword_5d4594_1193348) {
		if (dword_5d4594_1193352) {
			v2 = *(uint32_t*)(dword_5d4594_1193352 + 32);
			if (nox_strman_get_lang_code() == 6 || nox_strman_get_lang_code() == 8) {
				if (!*(uint32_t*)(v2 + 1036)) {
					if (!*(uint32_t*)(v2 + 1032)) {
						if (!*(uint32_t*)(v2 + 1028)) {
							*(uint32_t*)(v2 + 1044) = 1;
							switch (a1) {
							case 7u:
							case 8u:
							case 9u:
							case 0xBu:
							case 0xCu:
								break;
							case 0xAu:
							case 0xDu:
								*(uint32_t*)(v2 + 1044) = 0;
								break;
							default:
								v3 = *(unsigned short*)(v2 + 1052);
								if ((unsigned short)v3 < *(short*)(v2 + 1040) - 1) {
									*(uint16_t*)(v2 + 2 * v3) = a1;
									*(uint16_t*)(v2 + 2 * (unsigned short)++*(uint16_t*)(v2 + 1052)) = 0;
								}
								wchar2_t* v4 = nox_input_getStringBuffer_57011C();
								nox_wcscpy((wchar2_t*)(v2 + 512), (const wchar2_t*)v4);
								nox_input_freeStringBuffer_57011C(v4);
								*(uint16_t*)(v2 + 1054) = nox_wcslen((const wchar2_t*)(v2 + 512));
								// nox_xxx_string_570392(*(uint32_t***)&dword_5d4594_1193348);
								nox_window_set_hidden(*(uint32_t*)(v2 + 1048), 1);
								break;
							}
						}
					}
				}
			}
		}
	}
}

//----- (004896E0) --------------------------------------------------------
int sub_4896E0() {
	if (dword_5d4594_1193360) {
		free(*(void**)&dword_5d4594_1193360);
	}
	return 1;
}

//----- (00489870) --------------------------------------------------------
int sub_489870() {
	int v0;            // eax
	unsigned char* v1; // esi
	uint32_t* v2;      // eax
	const wchar2_t* v3; // eax
	unsigned int v4;   // eax
	int v5;            // edx
	char v6;           // cl
	int v7;            // eax

	v0 = 0;
	v1 = getMemAt(0x5D4594, 1193388 + 44 * v0);
	if (*getMemU32Ptr(0x5D4594, 1193372 + 4 * v0) == 2) {
		*(uint32_t*)v1 =
			(nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10028)->draw_data.field_0 >> 2) & 1;
		v2 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10031);
		v3 = (const wchar2_t*)nox_window_call_field_94((int)v2, 16413, 0, 0);
		*((uint32_t*)v1 + 4) = nox_wcstol(v3, 0, 10);
		*((uint32_t*)v1 + 1) =
			(nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10029)->draw_data.field_0 >> 2) & 1;
		v4 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10030)->draw_data.field_0;
		*((uint32_t*)v1 + 3) = 0;
		*((uint32_t*)v1 + 2) = (v4 >> 2) & 1;
		if (nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10015)->draw_data.field_0 & 4) {
			v5 = *((uint32_t*)v1 + 3);
			LOBYTE(v5) = v5 | 0x80;
			*((uint32_t*)v1 + 3) = v5;
			v6 = *((uint8_t*)(&nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10016)
								   ->draw_data.field_0));
			v7 = *((uint32_t*)v1 + 3);
			if (v6 & 4) {
				LOBYTE(v7) = v7 | 1;
			} else {
				LOBYTE(v7) = v7 | 2;
			}
			*((uint32_t*)v1 + 3) = v7;
		}
		*((uint32_t*)v1 + 5) =
			(nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10014)->draw_data.field_0 >> 2) & 1;
		*((uint32_t*)v1 + 10) =
			(nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10018)->draw_data.field_0 >> 2) & 1;
	}
	return nox_window_set_hidden(*(int*)&dword_5d4594_1193380, 1);
}

//----- (004899C0) --------------------------------------------------------
int nox_xxx_checkSomeFlagsOnJoin_4899C0(nox_gui_server_ent_t* srv) {
	int a1 = srv;
	int v1;            // eax
	int v2;            // edx
	int v3;            // eax
	unsigned char* v4; // ebp
	int v5;            // eax
	int v6;            // ebx
	unsigned int v7;   // eax
	char v9;           // al
	unsigned char v10; // cl
	unsigned char v11; // cl
	char v12;          // al
	int v13[15];       // [esp+Ch] [ebp-3Ch]
	unsigned char v14; // [esp+4Ch] [ebp+4h]
	unsigned char v15; // [esp+4Ch] [ebp+4h]

	v1 = 0;
	v2 = 11 * v1;
	v3 = *getMemU32Ptr(0x5D4594, 1193372 + 4 * v1);
	v4 = getMemAt(0x5D4594, 1193388 + 4 * v2);
	if (!v3) {
		return 1;
	}
	v5 = v3 - 1;
	if (v5) {
		if (v5 == 1) {
			v6 = a1;
			if (*(uint32_t*)v4) {
				v7 = *(uint32_t*)(a1 + 96);
				if (v7 > *((uint32_t*)v4 + 4) && v7 != 9999) {
					return 0;
				}
			}
			if (*((uint32_t*)v4 + 1) && *(uint8_t*)(a1 + 100) & 0x10) {
				return 0;
			}
			if (*((uint32_t*)v4 + 2) && *(uint8_t*)(a1 + 100) & 0x20) {
				return 0;
			}
			v9 = *(uint8_t*)(a1 + 102);
			if (v9 < 0 && *((uint32_t*)v4 + 3) > (v9 & 0x7F)) {
				return 0;
			}
			if (*((uint32_t*)v4 + 5)) {
				strcpy((char*)v13, (const char*)(a1 + 111));
				sub_57A1E0(v13, 0, 0, 5, *(uint16_t*)(a1 + 163));
				v10 = 0;
				v14 = 0;
				while (v13[v14 + 6] == *(uint32_t*)(4 * v14 + v6 + 135)) {
					v14 = ++v10;
					if (v10 >= 5u) {
						v11 = 0;
						v15 = 0;
						while (*((uint8_t*)&v13[11] + v15) == *(uint8_t*)(v15 + v6 + 155)) {
							v15 = ++v11;
							if (v11 >= 4u) {
								if (v13[12] == *(uint32_t*)(v6 + 159)) {
									goto LABEL_26;
								}
								return 0;
							}
						}
						return 0;
					}
				}
				return 0;
			}
		LABEL_26:
			if (*((uint32_t*)v4 + 10) && *(uint32_t*)(v6 + 48) != NOX_CLIENT_VERS_CODE) {
				return 0;
			}
		}
		return 1;
	}
	v12 = *(uint8_t*)(a1 + 100);
	if (v12 & 0x10) {
		return 0;
	}
	if (v12 & 0x20) {
		return 0;
	}
	return *(uint32_t*)(a1 + 48) == NOX_CLIENT_VERS_CODE;
}

//----- (00489B80) --------------------------------------------------------
uint32_t* sub_489B80(int a1) {
	uint32_t* result;  // eax
	int v2;            // ebx
	unsigned char* v3; // edi
	uint32_t* v4;      // eax
	uint32_t* v5;      // eax
	uint32_t* v6;      // eax
	uint32_t* v7;      // eax
	uint32_t* v8;      // eax
	uint32_t* v9;      // eax
	uint32_t* v10;     // esi
	int v11;           // ebx
	int v12;           // ebx
	wchar2_t v13[16];   // [esp+0h] [ebp-20h]

	result = nox_new_window_from_file("filter.wnd", nox_xxx_windowMplayFilterProc_489E70);
	dword_5d4594_1193380 = result;
	if (result) {
		dword_5d4594_1193384 = nox_xxx_wndGetChildByID_46B0C0(result, 10012);
		v2 = 0;
		v3 = getMemAt(0x5D4594, 1193388 + 44 * v2);
		sub_46B120(*(uint32_t**)&dword_5d4594_1193380, a1);
		sub_46B120(*(uint32_t**)&dword_5d4594_1193384, *(int*)&dword_5d4594_1193380);
		nox_xxx_wndSetProc_46B2C0(*(int*)&dword_5d4594_1193384, nox_xxx_windowMplayFilterProc_489E70);
		if (*(uint32_t*)v3) {
			v4 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10028);
			v4[9] |= 4u;
		}
		nox_swprintf(v13, L"%d", *((uint32_t*)v3 + 4));
		v5 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10031);
		nox_window_call_field_94((int)v5, 16414, (int)v13, -1);
		if (*((uint32_t*)v3 + 1)) {
			v6 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10029);
			v6[9] |= 4u;
		}
		if (*((uint32_t*)v3 + 2)) {
			v7 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10030);
			v7[9] |= 4u;
		}
		if (*((uint32_t*)v3 + 3)) {
			v8 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10015);
			v8[9] |= 4u;
		}
		if (v3[12] & 2) {
			v9 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10017);
		} else {
			v9 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10016);
		}
		v10 = v9;
		v9[9] |= 4u;
		if (*((uint32_t*)v3 + 5)) {
			v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10014);
			v10[9] |= 4u;
		}
		if (*((uint32_t*)v3 + 10)) {
			v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10018);
			v10[9] |= 4u;
		}
		v11 = *getMemU32Ptr(0x5D4594, 1193372 + 4 * v2);
		if (v11) {
			v12 = v11 - 1;
			if (v12) {
				if (v12 == 1) {
					v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10026);
					sub_489DC0();
				}
			} else {
				v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10025);
				nox_window_set_hidden(*(int*)&dword_5d4594_1193384, 1);
			}
		} else {
			v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10024);
			nox_window_set_hidden(*(int*)&dword_5d4594_1193384, 1);
		}
		v10[9] |= 4u;
		result = *(uint32_t**)&dword_5d4594_1193380;
	}
	return result;
}

//----- (00489DC0) --------------------------------------------------------
void sub_489DC0() {
	uint32_t* v0; // eax
	uint32_t* v1; // eax

	nox_window_set_hidden(*(int*)&dword_5d4594_1193384, 0);
	if (nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193384, 10028)->draw_data.field_0 & 4) {
		v0 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193384, 10031);
		nox_xxx_wnd_46ABB0((int)v0, 1);
	} else {
		v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193384, 10031);
		nox_xxx_wnd_46ABB0((int)v1, 0);
	}
	if (nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193384, 10015)->draw_data.field_0 & 4) {
		sub_46AD20(*(uint32_t**)&dword_5d4594_1193384, 10016, 10017, 1);
	} else {
		sub_46AD20(*(uint32_t**)&dword_5d4594_1193384, 10016, 10017, 0);
	}
}

//----- (00489E70) --------------------------------------------------------
int nox_xxx_windowMplayFilterProc_489E70(int a1, int a2, int* a3, int a4) {
	int v3;       // ebx
	int v4;       // esi
	int result;   // eax
	uint32_t* v6; // eax
	int v7;       // [esp-Ch] [ebp-10h]
	int v8;       // [esp-Ch] [ebp-10h]

	v3 = 0;
	if (a2 == 23) {
		return 1;
	}
	if (a2 != 16391) {
		return 0;
	}
	v4 = nox_xxx_wndGetID_46B0A0(a3);
	nox_xxx_clientPlaySoundSpecial_452D80(766, 100);
	switch (v4) {
	case 10015:
		sub_46AD20(*(uint32_t**)&dword_5d4594_1193380, 10016, 10017, ((unsigned int)~a3[9] >> 2) & 1);
		return 0;
	case 10024:
		v7 = dword_5d4594_1193384;
		*getMemU32Ptr(0x5D4594, 1193372 + 4 * v3) = 0;
		nox_window_set_hidden(v7, 1);
		result = 0;
		break;
	case 10025:
		v8 = dword_5d4594_1193384;
		*getMemU32Ptr(0x5D4594, 1193372 + 4 * v3) = 1;
		nox_window_set_hidden(v8, 1);
		result = 0;
		break;
	case 10026:
		*getMemU32Ptr(0x5D4594, 1193372 + 4 * v3) = 2;
		sub_489DC0();
		result = 0;
		break;
	case 10028:
		v6 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1193380, 10031);
		nox_xxx_wnd_46ABB0((int)v6, ((unsigned int)~a3[9] >> 2) & 1);
		result = 0;
		break;
	default:
		return 0;
	}
	return result;
}

//----- (00489FB0) --------------------------------------------------------
int sub_489FB0() {
	int result; // eax

	result = dword_5d4594_1193380;
	if (dword_5d4594_1193380) {
		sub_489870();
		nox_xxx_wndClearCaptureMain_46ADE0(*(int*)&dword_5d4594_1193380);
		result = nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1193380);
		dword_5d4594_1193380 = 0;
	}
	return result;
}

//----- (00489FF0) --------------------------------------------------------
int sub_489FF0(int a1, int a2, const void* a3) {
	int result; // eax

	*getMemU32Ptr(0x5D4594, 1193372 + 4 * a1) = a2;
	result = 11 * a1;
	memcpy(getMemAt(0x5D4594, 1193388 + 44 * a1), a3, 0x2Cu);
	return result;
}

//----- (0048A210) --------------------------------------------------------
int nox_xxx_setSomeFunc_48A210(int a1) {
	int result; // eax

	result = a1;
	*getMemU32Ptr(0x5D4594, 1193504) = a1;
	return result;
}
