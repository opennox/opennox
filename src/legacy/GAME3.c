#include <math.h>
#include <sys/stat.h>

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
#include "GAME3_2.h"
#include "GAME4.h"
#include "GAME4_1.h"
#include "GAME5_2.h"
#include "client__draw__animdraw.h"
#include "client__drawable__drawable.h"
#include "client__video__draw_common.h"
#include "common__system__team.h"

#include "client__gui__guicon.h"
#include "client__gui__guiquit.h"
#include "client__gui__window.h"
#include "client__shell__mainmenu.h"
#include "client__shell__noxworld.h"
#include "client__shell__optsback.h"
#include "client__shell__selchar.h"
#include "client__shell__selclass.h"
#include "client__shell__selcolor.h"

#include "client__draw__drawrays.h"
#include "client__draw__fx.h"
#include "client__draw__lvupdraw.h"

#include "common/fs/nox_fs.h"
#include "common__binfile.h"
#include "common__crypt.h"
#include "input.h"
#include "input_common.h"
#include "operators.h"

extern uint32_t dword_8531A0_2572;
extern uint32_t dword_5d4594_1307792;
extern uint32_t dword_5d4594_1313788;
extern uint32_t dword_5d4594_1308136;
extern uint32_t nox_xxx_normalWndBits_587000_172880;
extern uint32_t dword_5d4594_1308104;
extern uint32_t dword_5d4594_1313532;
extern uint32_t dword_5d4594_1313564;
extern uint32_t dword_5d4594_1308096;
extern uint32_t dword_5d4594_1307724;
extern uint32_t dword_5d4594_1308112;
extern uint32_t nox_server_sendMotd_108752;
extern uint32_t dword_5d4594_1308148;
extern uint32_t dword_587000_171388;
extern uint32_t dword_5d4594_1308140;
extern uint32_t dword_5d4594_1308144;
extern uint32_t dword_5d4594_1308152;
extern uint32_t dword_5d4594_1313536;
extern uint32_t dword_5d4594_1308116;
extern uint32_t dword_5d4594_1313740;
extern uint32_t dword_5d4594_1308100;
extern uint32_t dword_5d4594_1308132;
extern uint32_t dword_5d4594_1308108;
extern uint32_t dword_5d4594_1308120;
extern uint32_t dword_5d4594_1308128;
extern uint32_t dword_5d4594_1308124;
extern uint32_t dword_5d4594_1309736;
extern uint32_t dword_5d4594_1309756;
extern uint32_t dword_5d4594_1309832;
extern uint32_t dword_5d4594_1313540;
extern uint32_t dword_5d4594_1309824;
extern uint32_t dword_5d4594_1307720;
extern uint32_t dword_5d4594_1307736;
extern uint32_t nox_server_connectionType_3596;
extern uint32_t dword_5d4594_1309828;
extern uint32_t dword_5d4594_1309836;
extern uint32_t dword_5d4594_1309728;
extern uint32_t dword_5d4594_1309732;
extern uint64_t qword_581450_9552;
extern void* dword_587000_93164;
extern uint32_t dword_5d4594_1307716;
extern void* dword_587000_127004;
extern uint32_t dword_5d4594_1308088;
extern uint32_t dword_5d4594_1309748;
extern uint32_t dword_5d4594_1309720;
extern uint32_t dword_5d4594_1308084;
extern uint32_t dword_5d4594_1309812;
extern uint32_t dword_5d4594_1309820;
extern uint32_t dword_5d4594_2650652;
extern uint32_t dword_5d4594_1307784;
extern int nox_win_width;
extern int nox_win_height;

extern uint32_t nox_color_white_2523948;
extern uint32_t nox_color_blue_2650684;
extern uint32_t nox_color_violet_2598268;
extern uint32_t nox_color_black_2650656;

nox_gui_animation* nox_wnd_xxx_1307732 = 0;
nox_gui_animation* nox_wnd_xxx_1308092 = 0;
nox_gui_animation* nox_wnd_xxx_1309740 = 0;

void* dword_5d4594_1308156 = 0;
void* dword_5d4594_1308160 = 0;
void* dword_5d4594_1308164 = 0;

//----- (004A2560) --------------------------------------------------------
int sub_4A2560(uint32_t* a1, int a2) {
	double v2; // st7
	double v3; // st6

	v2 = (double)(*(short*)(a2 + 44) - *a1);
	v3 = (double)(*(short*)(a2 + 46) - a1[1]);
	return sqrt(v3 * v3 + v2 * v2) <= *getMemDoublePtr(0x581450, 9720);
}

//----- (004A25C0) --------------------------------------------------------
int sub_4A25C0(uint32_t* a1, int* a2) {
	int v2;  // edi
	int* v3; // esi

	v2 = 0;
	v3 = nox_common_list_getFirstSafe_425890(a2);
	if (!v3) {
		return 0;
	}
	do {
		if (sub_4A2560(a1, (int)v3)) {
			++v2;
		}
		v3 = nox_common_list_getNextSafe_4258A0(v3);
	} while (v3);
	return v2;
}

//----- (004A2610) --------------------------------------------------------
int sub_4A2610(int a1, uint32_t* a2, int* a3) {
	int* i;             // esi
	int v4;             // eax
	uint32_t* v5;       // esi
	uint32_t* v6;       // ebx
	uint32_t* v7;       // edi
	uint32_t* v8;       // ebp
	char* v9;           // eax
	int v10;            // ebx
	unsigned char* v11; // esi
	uint32_t* v13;      // [esp+Ch] [ebp-150h]
	char* v14;          // [esp+10h] [ebp-14Ch]
	int v15[2];         // [esp+14h] [ebp-148h]
	char v16[64];       // [esp+1Ch] [ebp-140h]
	wchar2_t v17[128];   // [esp+5Ch] [ebp-100h]

	dword_5d4594_1307720 = 0;
	for (i = nox_common_list_getFirstSafe_425890(a3); i; i = nox_common_list_getNextSafe_4258A0(i)) {
		if (sub_4A2560(a2, (int)i)) {
			v4 = dword_5d4594_1307720;
			*getMemU32Ptr(0x5D4594, 1307316 + 4 * dword_5d4594_1307720) = i;
			dword_5d4594_1307720 = v4 + 1;
		}
	}
	if (dword_5d4594_1307720 > 0) {
		dword_5d4594_1307716 = nox_new_window_from_file("proxlist.wnd", *(uint32_t*)(a1 + 376));
		sub_4A2830(*a2 + 216, a2[1] + 27, v15);
		nox_window_setPos_46A9B0(*(uint32_t**)&dword_5d4594_1307716, v15[0], v15[1]);
		nox_xxx_wnd_46B280(*(int*)&dword_5d4594_1307716, a1);
		v5 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1307716, 10064);
		v6 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1307716, 10062);
		v13 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1307716, 10063);
		v7 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1307716, 10061);
		v8 = (uint32_t*)v7[8];
		v14 = nox_xxx_gLoadImg_42F970("UISlider");
		v9 = nox_xxx_gLoadImg_42F970("UISliderLit");
		sub_4B5700((int)v5, 0, 0, (int)v14, (int)v9, (int)v9);
		nox_xxx_wnd_46B280((int)v5, (int)v7);
		nox_xxx_wnd_46B280((int)v6, (int)v7);
		nox_xxx_wnd_46B280((int)v13, (int)v7);
		v8[9] = v5;
		v8[7] = v6;
		v8[8] = v13;
		*(uint32_t*)(v5[100] + 8) = 16;
		*(uint32_t*)(v5[100] + 12) = 10;
		v10 = 0;
		if (dword_5d4594_1307720 > 0) {
			v11 = getMemAt(0x5D4594, 1307316);
			do {
				if (*(uint8_t*)(*(uint32_t*)v11 + 120)) {
					strncpy(v16, (const char*)(*(uint32_t*)v11 + 120), 0xFu);
					v16[15] = 0;
				} else {
					nox_sprintAddrPort_43BC80(*(uint32_t*)v11 + 12, *(uint16_t*)(*(uint32_t*)v11 + 109), v16);
				}
				nox_swprintf(v17, L"%S   %dms", v16, *(uint32_t*)(*(uint32_t*)v11 + 96));
				nox_window_call_field_94((int)v7, 16397, (int)v17, -1);
				++v10;
				v11 += 4;
			} while (v10 < *(int*)&dword_5d4594_1307720);
		}
	}
	return dword_5d4594_1307716;
}

//----- (004A2830) --------------------------------------------------------
uint32_t* sub_4A2830(int a1, int a2, uint32_t* a3) {
	uint32_t* result; // eax

	result = a3;
	*a3 = a1 - 100;
	a3[1] = a2 - 20;
	if (a1 - 100 + 200 > 600) {
		*a3 = 400;
	}
	if (a2 - 20 + 200 > 451) {
		a3[1] = 251;
	}
	if (a3[1] < 27) {
		a3[1] = 27;
	}
	if ((int)*a3 < 216) {
		*a3 = 216;
	}
	return result;
}

//----- (004A2890) --------------------------------------------------------
int sub_4A2890() {
	int result; // eax

	result = dword_5d4594_1307716;
	if (dword_5d4594_1307716) {
		result = nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1307716);
		dword_5d4594_1307716 = 0;
	}
	return result;
}

//----- (004A28B0) --------------------------------------------------------
int sub_4A28B0() { return dword_5d4594_1307716 != 0; }

//----- (004A28C0) --------------------------------------------------------
int sub_4A28C0(int a1) {
	int result; // eax

	if (a1 < *(int*)&dword_5d4594_1307720) {
		result = *getMemU32Ptr(0x5D4594, 1307316 + 4 * a1);
	} else {
		result = 0;
	}
	return result;
}

//----- (004A28E0) --------------------------------------------------------
int nox_xxx_wndListboxProcWithoutData10_4A28E0(uint32_t* a1, int a2, unsigned int a3, int a4) {
	uint32_t* v4;     // edi
	int v5;           // esi
	int result;       // eax
	int v7;           // eax
	int v8;           // eax
	int v9;           // eax
	int v10;          // edx
	int v11;          // eax
	unsigned int v12; // ebx
	unsigned int v13; // ebx
	int v14;          // eax
	int v15;          // ecx
	int v16;          // eax
	unsigned int v17; // eax
	unsigned int v18; // ebx
	int v19;          // eax
	int v20;          // ecx
	int v21;          // eax
	int v22;          // eax

	v4 = a1;
	v5 = a1[8];
	switch (a2) {
	case 5:
	case 17:
	case 18:
		return 1;
	case 6:
	case 7:
		v12 = a3;
		a3 = *(uint32_t*)(v5 + 48);
		v13 = v12 >> 16;
		nox_client_wndGetPosition_46AA60(a1, &a1, &a4);
		if (*((uint16_t*)v4 + 54)) {
			v14 = nox_xxx_guiFontHeightMB_43F320(v4[59]);
			a4 += v14 + 1;
		}
		v15 = 0;
		*(uint32_t*)(v5 + 48) = -2;
		v16 = 0;
		while (2) {
			if (v16 > 0 && *(int*)(v16 + *(uint32_t*)(v5 + 24) - 524) > *(short*)(v5 + 52) + *(short*)(v5 + 54) ||
				v15 == *(short*)(v5 + 44)) {
				*(uint32_t*)(v5 + 48) = -1;
			} else if (*(int*)(v16 + *(uint32_t*)(v5 + 24)) <= (int)(v13 + *(short*)(v5 + 54) - a4)) {
				++v15;
				v16 += 524;
				continue;
			}
			break;
		}
		v17 = a3;
		if (v15 == a3 && !*(uint32_t*)(v5 + 20)) {
			*(uint32_t*)(v5 + 48) = -1;
		}
		if (*(int*)(v5 + 48) == -2 && v15 < *(short*)(v5 + 44)) {
			*(uint32_t*)(v5 + 48) = v15;
		}
		if (*(int*)(v5 + 48) < 0) {
			if (*(uint32_t*)(v5 + 20)) {
				*(uint32_t*)(v5 + 48) = v17;
			}
		}
		nox_window_call_field_94(v4[13], 16400, (int)v4, *(uint32_t*)(v5 + 48));
		return 1;
	case 8:
		v22 = a1[11];
		if (v22 & 0x100) {
			nox_window_call_field_94(a1[13], 0x4000, (int)a1, 0);
		}
		return 1;
	case 10:
	case 11:
		v18 = a3 >> 16;
		nox_client_wndGetPosition_46AA60(a1, &a3, &a4);
		if (*((uint16_t*)v4 + 54)) {
			v19 = nox_xxx_guiFontHeightMB_43F320(v4[59]);
			a4 += v19 + 1;
		}
		v20 = 0;
		v21 = 0;
		while ((v21 <= 0 || *(uint32_t*)(v21 + *(uint32_t*)(v5 + 24) - 524) <= *(short*)(v5 + 52) + *(short*)(v5 + 54)) &&
			   v20 != *(short*)(v5 + 44)) {
			if (*(uint32_t*)(v21 + *(uint32_t*)(v5 + 24)) > (int)(v18 + *(short*)(v5 + 54) - a4)) {
				nox_window_call_field_94(v4[13], 16401, (int)v4, v20);
				return 1;
			}
			++v20;
			v21 += 524;
		}
		v20 = -1;
		nox_window_call_field_94(v4[13], 16401, (int)v4, v20);
		return 1;
	case 19:
		v8 = *(uint32_t*)(v5 + 48);
		if (v8 == -1) {
			*(uint32_t*)(v5 + 48) = 0;
			nox_xxx_wndListBox_4A2D10((int)v4, 0, 1);
		} else if (v8 > 0) {
			v11 = v8 - 1;
			*(uint32_t*)(v5 + 48) = v11;
			if (*(uint32_t*)(*(uint32_t*)(v5 + 24) + 524 * v11) < *(short*)(v5 + 54)) {
				nox_xxx_wndListBox_4A2D10((int)v4, -1, 1);
			}
		}
		nox_window_call_field_94(v4[13], 16400, (int)v4, *(uint32_t*)(v5 + 48));
		return 1;
	case 20:
		v7 = *(uint32_t*)(v5 + 48);
		if (v7 == -1) {
			*(uint32_t*)(v5 + 48) = 0;
			nox_xxx_wndListBox_4A2D10((int)v4, 0, 1);
		} else if (v7 < *(short*)(v5 + 44) - 1) {
			v9 = v7 + 1;
			v10 = *(short*)(v5 + 52) + *(short*)(v5 + 54);
			*(uint32_t*)(v5 + 48) = v9;
			if (*(uint32_t*)(*(uint32_t*)(v5 + 24) + 524 * v9) > v10) {
				nox_xxx_wndListBox_4A2D10((int)v4, 1, 1);
			}
		}
		nox_window_call_field_94(v4[13], 16400, (int)v4, *(uint32_t*)(v5 + 48));
		return 1;
	case 21:
		switch (a3) {
		case 0xFu:
		case 0xCDu:
			if (a4 != 2) {
				return 1;
			}
			nox_xxx_wndRetNULL_46A8A0();
			return 1;
		case 0x1Cu:
		case 0x39u:
			if (a4 != 1) {
				return 1;
			}
			nox_window_call_field_94(a1[13], 16400, (int)a1, *(uint32_t*)(v5 + 48));
			return 1;
		case 0xC8u:
			if (a4 != 2) {
				return 1;
			}
			v8 = *(uint32_t*)(v5 + 48);
			if (v8 == -1) {
				*(uint32_t*)(v5 + 48) = 0;
				nox_xxx_wndListBox_4A2D10((int)v4, 0, 1);
			} else if (v8 > 0) {
				v11 = v8 - 1;
				*(uint32_t*)(v5 + 48) = v11;
				if (*(uint32_t*)(*(uint32_t*)(v5 + 24) + 524 * v11) < *(short*)(v5 + 54)) {
					nox_xxx_wndListBox_4A2D10((int)v4, -1, 1);
				}
			}
			nox_window_call_field_94(v4[13], 16400, (int)v4, *(uint32_t*)(v5 + 48));
			return 1;
		case 0xCBu:
			if (a4 != 2) {
				return 1;
			}
			nox_xxx_wndRetNULL_0_46A8B0();
			return 1;
		case 0xD0u:
			if (a4 != 2) {
				return 1;
			}
			v7 = *(uint32_t*)(v5 + 48);
			if (v7 == -1) {
				*(uint32_t*)(v5 + 48) = 0;
				nox_xxx_wndListBox_4A2D10((int)v4, 0, 1);
			} else if (v7 < *(short*)(v5 + 44) - 1) {
				v9 = v7 + 1;
				v10 = *(short*)(v5 + 52) + *(short*)(v5 + 54);
				*(uint32_t*)(v5 + 48) = v9;
				if (*(uint32_t*)(*(uint32_t*)(v5 + 24) + 524 * v9) > v10) {
					nox_xxx_wndListBox_4A2D10((int)v4, 1, 1);
				}
			}
			nox_window_call_field_94(v4[13], 16400, (int)v4, *(uint32_t*)(v5 + 48));
			return 1;
		default:
			return 0;
		}
	default:
		return 0;
	}
}

//----- (004A2D10) --------------------------------------------------------
int nox_xxx_wndListBox_4A2D10(int a1, int a2, int a3) {
	int v3;     // esi
	int v4;     // eax
	int v5;     // ecx
	int result; // eax
	int v7;     // ecx

	v3 = *(uint32_t*)(a1 + 32);
	v4 = a2 + sub_4A4800(*(uint32_t*)(a1 + 32));
	if (v4 >= 0) {
		v5 = *(short*)(v3 + 44);
		if (v4 >= v5) {
			v4 = v5 - 1;
		}
	} else {
		v4 = 0;
	}
	if (a3) {
		if (v4 <= 0) {
			*(uint16_t*)(v3 + 54) = 0;
		} else {
			*(uint16_t*)(v3 + 54) = *(uint16_t*)(*(uint32_t*)(v3 + 24) + 524 * v4 - 524) + 1;
		}
	}
	result = *(uint32_t*)(v3 + 36);
	if (result) {
		result = *(uint32_t*)(result + 32);
		v7 = *(uint32_t*)(v3 + 40) - *(short*)(v3 + 52) + 3;
		*(uint32_t*)(result + 4) = v7;
		if (v7 < 0) {
			*(uint32_t*)(result + 4) = 0;
		}
		*(float*)(result + 8) = (double)(int)(*(uint32_t*)(*(uint32_t*)(v3 + 36) + 12) -
											  *(uint32_t*)(*(uint32_t*)(*(uint32_t*)(v3 + 36) + 400) + 12)) /
								(double)*(int*)(result + 4);
		if (a3) {
			result = nox_window_call_field_94(*(uint32_t*)(v3 + 36), 16394,
											  *(uint32_t*)(result + 4) - *(short*)(v3 + 54), 0);
		}
	}
	return result;
}

//----- (004A2DE0) --------------------------------------------------------
int nox_xxx_wndListboxProcWithData10_4A2DE0(int a1, int a2, unsigned int a3, int a4) {
	int v4;           // ebp
	int v5;           // esi
	int result;       // eax
	unsigned int v7;  // edi
	int v8;           // eax
	unsigned int v9;  // ecx
	int v10;          // ebx
	int v11;          // eax
	int* v12;         // edi
	int v13;          // ecx
	int v14;          // eax
	int v15;          // edx
	unsigned int v16; // edi
	int v17;          // eax
	unsigned int v18; // edx
	int v19;          // ecx
	int v20;          // eax
	int v21;          // eax

	v4 = a1;
	v5 = *(uint32_t*)(a1 + 32);
	switch (a2) {
	case 5:
	case 17:
	case 18:
		return 1;
	case 6:
	case 7:
		v7 = a3 >> 16;
		nox_client_wndGetPosition_46AA60((uint32_t*)a1, &a2, &a3);
		if (*(uint16_t*)(a1 + 108)) {
			v8 = nox_xxx_guiFontHeightMB_43F320(*(uint32_t*)(a1 + 236));
			v9 = a3 + v8 + 1;
			a3 += v8 + 1;
		} else {
			v9 = a3;
		}
		v10 = 0;
		v11 = 0;
		while (2) {
			if (v11 > 0) {
				if (*(uint32_t*)(v11 + *(uint32_t*)(v5 + 24) - 524) > *(short*)(v5 + 52) + *(short*)(v5 + 54)) {
					v10 = -1;
					v4 = a1;
					break;
				}
				v4 = a1;
			}
			if (v10 == *(short*)(v5 + 44)) {
				v10 = -1;
				break;
			} else  if (*(uint32_t*)(v11 + *(uint32_t*)(v5 + 24)) <= (int)(v7 + *(short*)(v5 + 54) - v9)) {
				v4 = a1;
				++v10;
				v11 += 524;
			} else {
				v4 = a1;
				break;
			}
		}
		v12 = *(int**)(v5 + 48);
		v13 = 0;
		v14 = *v12;
		if ((int)*v12 >= 0) {
			v15 = *(uint32_t*)(v5 + 48);
			while (1) {
				if (v14 == v10) {
					sub_4A3090((short*)v5, v13);
					nox_window_call_field_94(*(uint32_t*)(v4 + 52), 16400, v4, v10);
					return 1;
				}
				v14 = *(uint32_t*)(v15 + 4);
				v15 += 4;
				++v13;
				if (v14 < 0) {
					break;
				}
			}
		}
		v12[v13] = v10;
		*(uint32_t*)(*(uint32_t*)(v5 + 48) + 4 * v13 + 4) = -1;
		nox_window_call_field_94(*(uint32_t*)(v4 + 52), 16400, v4, v10);
		return 1;
	case 8:
		v21 = *(uint32_t*)(a1 + 44);
		if (v21 & 0x100) {
			nox_window_call_field_94(*(uint32_t*)(a1 + 52), 0x4000, a1, 0);
		}
		return 1;
	case 10:
	case 11:
		v16 = a3 >> 16;
		nox_client_wndGetPosition_46AA60((uint32_t*)a1, &a2, &a3);
		if (*(uint16_t*)(a1 + 108)) {
			v17 = nox_xxx_guiFontHeightMB_43F320(*(uint32_t*)(a1 + 236));
			v18 = a3 + v17 + 1;
			a3 += v17 + 1;
		} else {
			v18 = a3;
		}
		v19 = 0;
		v20 = 0;
		while (2) {
			if (v20 > 0) {
				if (*(uint32_t*)(v20 + *(uint32_t*)(v5 + 24) - 524) > *(short*)(v5 + 52) + *(short*)(v5 + 54)) {
					v19 = -1;
					nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16401, a1, v19);
					return 1;
				}
				v4 = a1;
			}
			if (v19 != *(short*)(v5 + 44)) {
				if (*(uint32_t*)(v20 + *(uint32_t*)(v5 + 24)) <= (int)(v16 + *(short*)(v5 + 54) - v18)) {
					v4 = a1;
					++v19;
					v20 += 524;
					continue;
				}
				nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16401, a1, v19);
				return 1;
			}
			break;
		}
		nox_window_call_field_94(*(uint32_t*)(v4 + 52), 16401, v4, -1);
		return 1;
	case 19:
		if (!*(uint32_t*)(v5 + 28) || *(short*)(v5 + 54) <= 0) {
			return 1;
		}
		nox_xxx_wndListBox_4A2D10(a1, -1, 1);
		return 1;
	case 20:
		if (!*(uint32_t*)(v5 + 32) || *(short*)(v5 + 52) + *(short*)(v5 + 54) > *(int*)(v5 + 40)) {
			return 1;
		}
		nox_xxx_wndListBox_4A2D10(a1, 1, 1);
		return 1;
	case 21:
		if (a3 != 15) {
			return 0;
		}
		if (a4 != 2) {
			return 1;
		}
		nox_xxx_wndRetNULL_46A8A0();
		return 1;
	default:
		return 0;
	}
}

//----- (004A3090) --------------------------------------------------------
short* sub_4A3090(short* a1, int a2) {
	short* result; // eax

	result = a1;
	memcpy((void*)(*((uint32_t*)a1 + 12) + 4 * a2), (const void*)(*((uint32_t*)a1 + 12) + 4 * a2 + 4), 4 * (*a1 - a2));
	*(uint32_t*)(*((uint32_t*)a1 + 12) + 4 * *a1 - 4) = -1;
	return result;
}

//----- (004A30D0) --------------------------------------------------------
int nox_xxx_wndListboxProcPre_4A30D0(nox_window* win, unsigned int ev, uint32_t a3, int a4) {
	uint32_t* v9;  // eax
	uint32_t* v10; // eax
	uint32_t* v11; // eax
	int v12;       // eax
	short v13;     // ax
	uint32_t* v14; // eax
	int v15;       // esi
	wchar2_t* v16;  // ecx
	wchar2_t* v17;  // eax
	short v18;     // cx
	short v19;     // ax
	int v20;       // eax
	int v21;       // edx
	int v22;       // esi
	int v23;       // eax
	int v24;       // edx
	int v25;       // esi
	uint32_t* v26; // eax
	int v27;       // ecx
	int v28;       // edx
	short v29;     // cx
	short v30;     // ax
	int v31;       // esi
	short v32;     // ax
	wchar2_t** v33; // esi
	int v34;       // edx
	wchar2_t* v35;  // eax
	int v36;       // ecx
	int v37;       // esi
	int v38;       // ecx
	short v39;     // di
	int v40;       // eax
	uint32_t* v41; // eax
	int v42;       // esi
	int v43;       // ecx
	int v44;       // eax
	short v45;     // ax
	int v46;       // edx

	uint32_t* v4 = win;
	nox_scrollListBox_data* sdata = win->widget_data;
	int ind = 0;
	wchar2_t* wstr = 0;
	if (ev > 16402) {
		switch (ev) {
		case 0x4013u:
			ind = a3;
			if ((ind < 0) || (ind >= (int)sdata->count)) {
				if (sdata->field_4) {
					memset(sdata->field_12, 0xFF, 4 * (int)sdata->count);
					return 0;
				}
				sdata->field_12 = -1;
				return 0;
			}
			v37 = &sdata->items[ind];
			if (!*(uint16_t*)(v37 + 4)) {
				return 0;
			}
			if (sdata->field_4) {
				sdata->field_12[0] = ind;
				sdata->field_12[1] = -1;
				return 0;
			}
			v38 = sdata->field_13_1;
			sdata->field_12 = ind;
			if (*(uint32_t*)v37 < v38) {
				nox_window_call_field_94(win, 16412, ind, 0);
				return 0;
			}
			v39 = sdata->field_13_0;
			if (*(uint32_t*)v37 > v38 + v39) {
				if (ind <= 0) {
					sdata->field_13_1 = 0;
				} else {
					sdata->field_13_1 = *(uint16_t*)v37 - v39;
				}
				nox_xxx_wndListBox_4A2D10(win, 0, 1);
				return 0;
			}
			return 0;
		case 0x4014u:
			return sdata->field_12;
		case 0x4015u:
			ind = a3;
			if ((ind < 0) || (ind >= (int)sdata->count)) {
				if (!sdata->field_4) {
					return 0;
				}
				memset(sdata->field_12, 0xFF, 4 * (int)sdata->count);
				return 0;
			}
			if (!sdata->items[ind].text[0] || !sdata->field_4) {
				return 0;
			}
			v33 = (wchar2_t**)sdata->field_12;
			v34 = 0;
			v35 = *v33;
			if ((int)*v33 < 0) {
				v33[v34] = ind;
				sdata->field_12[v34 + 1] = -1;
				return 0;
			}
			v36 = sdata->field_12;
			while (v35 != ind) {
				v35 = *(wchar2_t**)(v36 + 4);
				v36 += 4;
				++v34;
				if ((int)v35 < 0) {
					v33[v34] = ind;
					sdata->field_12[v34 + 1] = -1;
					return 0;
				}
			}
			sub_4A3090(sdata, v34);
			return 0;
		case 0x4016u:
			ind = a3;
			if ((ind < 0) || (ind >= (int)sdata->count)) {
				return 0;
			}
			return sdata->items[ind].text;
		case 0x4017u:
			ind = a4;
			wstr = a3;
			if ((ind < 0) || (ind >= (int)sdata->count)) {
				return 0;
			}
			nox_wcsncpy(sdata->items[ind].text, wstr, 255);
			sdata->items[ind].text[nox_wcslen(wstr)] = 0; // TODO: potential overflow
			return 0;
		case 0x4018u:
			sdata->field_7 = a3;
			return 0;
		case 0x4019u:
			sdata->field_8 = a3;
			return 0;
		case 0x401Au:
			sdata->field_9 = a3;
			return 0;
		case 0x401Bu:
			ind = a3;
			v40 = sdata->field_11_0;
			if (v40 < ind) {
				return 0;
			}
			memmove(sdata->items, &sdata->items[ind], sizeof(nox_scrollListBox_item) * (unsigned int)(v40 - ind));
			sdata->field_11_0 -= (short)ind;
			sdata->field_11_1 = sdata->field_11_0;
			if (sdata->field_4) {
				v41 = sdata->field_12;
				v42 = 0;
				if ((int)*v41 >= 0) {
					do {
						v43 = v41[v42];
						if (ind < v43) {
							sub_4A3090(sdata, v42--);
						} else {
							v41[v42] = v43 - ind;
						}
						v41 = (uint32_t*)sdata->field_12;
						++v42;
					} while ((int)v41[v42] >= 0);
				}
			} else {
				v44 = sdata->field_12;
				if (v44 > 0) {
					sdata->field_12 = v44 - ind;
				}
			}
			if ((int)sdata->field_13_1 > 0) {
				nox_xxx_wndListBox_4A2D10(win, -ind, 1);
			}
			nox_xxx_wndListBox_4A3A70(win);
			return 1;
		case 0x401Cu:
			ind = a3;
			if ((ind - 1 >= 0) && (ind - 1 < (int)sdata->count)) {
				sdata->field_13_1 = sdata->items[ind - 1].field_0 + 1;
			} else {
				sdata->field_13_1 = 0;
			}
			v45 = sdata->field_13_0;
			v46 = sdata->field_10;
			if (sdata->field_13_1 + v45 >= v46) {
				sdata->field_13_1 = v46 - v45;
			}
			nox_xxx_wndListBox_4A2D10(win, 0, 1);
			return 0;
		default:
			return 0;
		}
		return 0; // unreachable
	}
	if (ev == 16402) {
		if ((int)a3 >= 0) {
			v32 = sdata->field_11_0;
			if ((int)a3 <= v32) {
				sdata->field_11_1 = (short)a3;
			} else {
				sdata->field_11_1 = v32;
			}
		} else {
			sdata->field_11_1 = 0;
		}
		return 0;
	}
	if (ev > 0x4007) {
		switch (ev) {
		case 0x4009u:
			v29 = sdata->field_13_0;
			v30 = *(uint16_t*)(*(uint32_t*)((uint32_t)sdata->field_9 + 32) + 4) - a4;
			v31 = sdata->field_10 - v29 + 1;
			sdata->field_13_1 = v30;
			if (v30 > v31) {
				sdata->field_13_1 = (int)sdata->field_10 - v29 + 1;
			}
			if ((int)sdata->field_13_1 < 0) {
				sdata->field_13_1 = 0;
			}
			nox_xxx_wndListBox_4A2D10(win, 0, 0);
			return 0;
		case 16397:
			v18 = sdata->field_11_1;
			v19 = sdata->field_11_0;
			if (v18 != v19) {
				if (v19 < (int)sdata->count) {
					v20 = v19 - 1;
					if (v20 >= v18) {
						v21 = 524 * v20;
						do {
							--v20;
							v22 = (uint32_t*)sdata->items + v21;
							v21 -= 524;
							memcpy((void*)(v22 + 524), (const void*)v22, 0x20Cu);
						} while (v20 >= sdata->field_11_1);
						v4 = win;
					}
					nox_xxx_wndListBoxAddLine_4A3AC0(a3, a4, v4);
				} else {
					if (!sdata->field_2) {
						return 0;
					}
					nox_window_call_field_94(win, 16411, 1, 0);
					v23 = sdata->field_11_0 - 1;
					if (v23 >= (int)sdata->field_11_1) {
						v24 = 524 * v23;
						do {
							--v23;
							v25 = (uint32_t)sdata->items + v24;
							v24 -= 524;
							memcpy((void*)(v25 + 524), (const void*)v25, 0x20Cu);
						} while (v23 >= sdata->field_11_1);
						v4 = win;
					}
					nox_xxx_wndListBoxAddLine_4A3AC0(a3, a4, v4);
				}
			} else if (v18 >= (int)sdata->count) {
				if (!sdata->field_2) {
					return 0;
				}
				nox_window_call_field_94(win, 16411, 1, 0);
				nox_xxx_wndListBoxAddLine_4A3AC0(a3, a4, win);
			} else {
				nox_xxx_wndListBoxAddLine_4A3AC0(a3, a4, win);
			}
			if (sdata->field_1) {
				while (sdata->items[sdata->field_11_1 - 1].field_0 >= sdata->field_13_1 + sdata->field_13_0) {
					nox_xxx_wndListBox_4A2D10(v4, 1, 1);
				}
			}
			v26 = (uint32_t*)sdata->field_12;
			if (sdata->field_4) {
				if ((int)*v26 >= 0) {
					v27 = 0;
					do {
						v28 = v26[v27];
						if ((int)sdata->field_11_1 < v28) {
							v26[v27] = v28 + 1;
						}
						v26 = (uint32_t*)sdata->field_12;
						++v27;
					} while ((int)v26[v27] >= 0);
					return 1;
				}
			} else if ((int)sdata->field_11_1 < (int)v26) {
				sdata->field_12 = (char*)v26 + 1;
			}
			return 1;
		case 0x400Eu:
			ind = a3;
			v12 = sdata->field_11_0;
			if ((ind < 0) || (ind >= v12)) {
				return 0;
			}
			for (int i = ind; i < v12 - 1; i++) {
				memcpy(&sdata->items[i], &sdata->items[i + 1], sizeof(nox_scrollListBox_item));
			}
			v13 = --sdata->field_11_0;
			memset(&sdata->items[v13], 0, sizeof(nox_scrollListBox_item));
			sdata->field_11_1 = v13;
			sdata->items[v13].text[0] = 0;
			if (sdata->field_4) {
				v14 = (uint32_t*)sdata->field_12;
				v15 = 0;
				if ((int)*v14 >= 0) {
					do {
						v16 = (wchar2_t*)v14[v15];
						if (ind >= (int)v16) {
							if (ind == v16) {
								sub_4A3090(sdata, v15--);
							}
						} else {
							v14[v15] = (char*)v16 - 1;
						}
						v14 = (uint32_t*)sdata->field_12;
						++v15;
					} while ((int)v14[v15] >= 0);
					nox_xxx_wndListBox_4A3A70(win);
					return 0;
				}
			} else {
				v17 = (wchar2_t*)sdata->field_12;
				if (ind < (int)v17) {
					sdata->field_12 = (char*)v17 - 1;
					nox_xxx_wndListBox_4A3A70(win);
					return 0;
				}
				if (ind == v17) {
					sdata->field_12 = -1;
				}
			}
			nox_xxx_wndListBox_4A3A70(win);
			return 0;
		case 0x400Fu:
			memset(sdata->items, 0, sizeof(nox_scrollListBox_item) * (int)sdata->count);
			if (a3 != 1) {
				sdata->field_13_1 = 0;
			}
			if (sdata->field_4) {
				memset(sdata->field_12, 0xFF, 4 * (int)sdata->count);
			} else {
				sdata->field_12 = -1;
			}
			sdata->field_11_1 = 0;
			sdata->field_11_0 = 0;
			sdata->field_10 = 0;
			nox_xxx_wndListBox_4A2D10(win, 0, 1);
			return 0;
		default:
			return 0;
		}
	}
	if (ev == 16391) {
		if (a3 == sdata->field_7) {
			if ((int)sdata->field_13_1 > 0) {
				nox_xxx_wndListBox_4A2D10(win, -1, 1);
				return 0;
			}
		} else if (a3 == sdata->field_8 && sdata->field_13_0 + sdata->field_13_1 <= sdata->field_10) {
			nox_xxx_wndListBox_4A2D10(win, 1, 1);
			return 0;
		}
		return 0;
	}
	if (ev == 16385) {
		wstr = a3;
		nox_wcsncpy(win->draw_data.text, wstr, 63);
		win->draw_data.text[nox_wcslen(wstr)] = 0; // TODO: this may overflow
		return 0;
	}
	if (ev == 16388) {
		v9 = sdata->field_7;
		if (v9) {
			nox_window_setPos_46A9B0(v9, (int)a3 - v9[2], 0);
		}
		v10 = sdata->field_8;
		if (v10) {
			nox_window_setPos_46A9B0(v10, (int)a3 - v10[2], a4 - v10[3]);
		}
		v11 = sdata->field_9;
		if (v11) {
			nox_window_setPos_46A9B0(v11, (int)a3 - v11[2], *(uint32_t*)((uint32_t)sdata->field_7 + 12));
			sub_46AB20(sdata->field_9, *(uint32_t*)((uint32_t)sdata->field_9 + 8),
					   a4 - 2 * *(uint32_t*)(*(uint32_t*)((uint32_t)sdata->field_9 + 400) + 12));
		}
		sdata->field_13_0 = a4;
		if (win->draw_data.text[0]) {
			sdata->field_13_0 -= nox_xxx_guiFontHeightMB_43F320(win->draw_data.font);
		}
		return 0;
	}
	if (ev > 0x4000) {
		return 0;
	}
	if (ev == 0x4000) {
		if (a3 == sdata->field_7) {
			if ((int)sdata->field_13_1 > 0) {
				nox_xxx_wndListBox_4A2D10(win, -1, 1);
				return 0;
			}
		} else if (a3 == sdata->field_8 && sdata->field_13_0 + sdata->field_13_1 <= sdata->field_10) {
			nox_xxx_wndListBox_4A2D10(win, 1, 1);
			return 0;
		}
		return 0;
	}
	if (ev == 2) {
		if (sdata) {
			free(sdata->items);
			if (sdata->field_4) {
				free(sdata->field_12);
			}
			free(sdata);
		}
		win->widget_data = 0;
		return 0;
	}
	if (ev == 23) {
		if (a3) {
			win->draw_data.field_0 |= 2;
		} else {
			win->draw_data.field_0 &= 0xFFFFFFFD;
		}
		int v7 = nox_xxx_wndGetID_46B0A0(win);
		nox_window_call_field_94(win->draw_data.win, 16387, a3, v7);
		return 1;
	}
	return 0;
}

//----- (004A3A70) --------------------------------------------------------
int nox_xxx_wndListBox_4A3A70(int a1) {
	int v1; // edi
	int v2; // eax
	int v3; // ecx
	int v4; // ebp
	int v5; // esi

	v1 = 0;
	v2 = 0;
	v3 = *(uint32_t*)(a1 + 32);
	if (*(uint16_t*)(v3 + 44) > 0) {
		v4 = 0;
		do {
			v5 = v4 + *(uint32_t*)(v3 + 24);
			++v1;
			v4 += 524;
			v2 += *(unsigned char*)(v5 + 520) + 1;
			*(uint32_t*)v5 = v2;
		} while (v1 < *(short*)(v3 + 44));
	}
	*(uint32_t*)(v3 + 40) = v2;
	return nox_xxx_wndListBox_4A2D10(a1, 0, 1);
}

//----- (004A3AC0) --------------------------------------------------------
int nox_xxx_wndListBoxAddLine_4A3AC0(wchar2_t* a1, int a2, uint32_t* a3) {
	uint32_t* v3; // ebx
	int v4;       // esi
	int v5;       // ebp
	wchar2_t* v6;  // edi
	wchar2_t* v7;  // eax

	v3 = a3;
	v4 = a3[8];
	v5 = a3[2] - 7;
	if (*(uint32_t*)(v4 + 12)) {
		v5 = a3[2] - 17;
	}
	if (a2 >= 17 || a2 < 0) {
		*(uint32_t*)(*(uint32_t*)(v4 + 24) + 524 * *(short*)(v4 + 46) + 516) = a3[26];
	} else {
		*(uint32_t*)(*(uint32_t*)(v4 + 24) + 524 * *(short*)(v4 + 46) + 516) =
			**(uint32_t**)getMemAt(0x85B3FC, 132 + 4 * a2);
	}
	v6 = (wchar2_t*)(*(uint32_t*)(v4 + 24) + 524 * *(short*)(v4 + 46) + 4);
	if (a1) {
		nox_wcsncpy(v6, a1, 0xFFu);
		v6[255] = 0;
		v7 = &v6[nox_wcslen(v6) - 1];
		if (*v7 == 10) {
			*v7 = 0;
		}
	} else {
		*v6 = 32;
		v6[1] = 0;
	}
	if ((v3[1] & 0x4000) == 0x4000) {
		*(uint8_t*)(*(uint32_t*)(v4 + 24) + 524 * *(short*)(v4 + 46) + 520) = nox_xxx_guiFontHeightMB_43F320(v3[59]);
	} else {
		nox_xxx_drawGetStringSize_43F840(v3[59], v6, 0, &a3, v5);
		*(uint8_t*)(*(uint32_t*)(v4 + 24) + 524 * *(short*)(v4 + 46) + 520) = (uint8_t)a3;
	}
	++*(uint16_t*)(v4 + 46);
	++*(uint16_t*)(v4 + 44);
	return nox_xxx_wndListBox_4A3A70((int)v3);
}

//----- (004A3C00) --------------------------------------------------------
void nox_xxx_wndListboxInit_4A3C00(int a1, int a2) {
	if (a1) {
		if ((signed char)*(uint8_t*)(a1 + 4) >= 0) {
			nox_xxx_wndSetDrawFn_46B340(a1, nox_xxx_wndListboxDrawNoImage_4A3C50);
		} else {
			nox_xxx_wndSetDrawFn_46B340(a1, nox_xxx_wndListboxDrawWithImage_4A3FC0);
		}
		if (*(uint32_t*)(a2 + 16)) {
			nox_xxx_wndSetWindowProc_46B300(a1, nox_xxx_wndListboxProcWithData10_4A2DE0);
		} else {
			nox_xxx_wndSetWindowProc_46B300(a1, nox_xxx_wndListboxProcWithoutData10_4A28E0);
		}
	}
}

//----- (004A3C50) --------------------------------------------------------
int nox_xxx_wndListboxDrawNoImage_4A3C50(uint32_t* a1, int a2) {
	int v2;         // ebp
	int v3;         // edi
	int v4;         // ebx
	int v5;         // eax
	int v6;         // edi
	int v7;         // edx
	int v8;         // edi
	int v9;         // ebx
	int v10;        // eax
	int v11;        // ecx
	int v12;        // edi
	int* v13;       // ecx
	int v14;        // eax
	short* v15;     // edi
	int v16;        // ebx
	int v17;        // ecx
	int v19;        // [esp-4h] [ebp-234h]
	int v20;        // [esp+10h] [ebp-220h]
	int v21;        // [esp+10h] [ebp-220h]
	int v22;        // [esp+14h] [ebp-21Ch]
	int v23;        // [esp+18h] [ebp-218h]
	int v24;        // [esp+18h] [ebp-218h]
	int yTop;       // [esp+1Ch] [ebp-214h]
	int xLeft;      // [esp+20h] [ebp-210h]
	int v27;        // [esp+24h] [ebp-20Ch]
	int i;          // [esp+28h] [ebp-208h]
	int v29;        // [esp+2Ch] [ebp-204h]
	short v30[256]; // [esp+30h] [ebp-200h]

	v2 = a1[8];
	v20 = *(uint32_t*)(a2 + 28);
	v23 = *(uint32_t*)(a2 + 20);
	nox_client_wndGetPosition_46AA60(a1, &xLeft, &yTop);
	v3 = nox_xxx_guiFontHeightMB_43F320(*(uint32_t*)(a2 + 200));
	v4 = a1[3];
	v22 = a1[2];
	if ((a1[1] & 0x2000) == 0x2000) {
		nox_draw_enableTextSmoothing_43F670(1);
	}
	nox_xxx_wndDraw_49F7F0();
	if (*(uint32_t*)(v2 + 12)) {
		v22 -= 10;
	}
	if (*(uint16_t*)(a2 + 72)) {
		nox_xxx_drawSetTextColor_434390(*(uint32_t*)(a2 + 68));
		nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), (uint16_t*)(a2 + 72), xLeft + 1, yTop, v22, v3);
		v4 += -1 - v3;
		yTop += v3 + 1;
	}
	if (a1[1] & 8) {
		if (*(uint8_t*)a2 & 2) {
			v6 = *(uint32_t*)(a2 + 36);
			v5 = v23;
			goto LABEL_13;
		}
		v5 = v23;
	} else {
		v5 = *(uint32_t*)(a2 + 44);
	}
	v6 = v20;
LABEL_13:
	if (v5 != 0x80000000) {
		nox_client_drawSetColor_434460(v5);
		nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, v22, v4);
	}
	if (v6 != 0x80000000) {
		nox_client_drawSetColor_434460(v6);
		nox_client_drawBorderLines_49CC70(xLeft, yTop, v22, v4);
	}
	nox_client_copyRect_49F6F0(xLeft, yTop, v22, v4);
	v7 = *(short*)(v2 + 54);
	v8 = yTop - v7;
	v21 = yTop - v7;
	if (*(uint32_t*)(a2 + 68) != 0x80000000) {
		v9 = 0;
		v24 = 0;
		for (i = 0;; i = v9) {
			if (v9 > 0 && *(uint32_t*)(*(uint32_t*)(v2 + 24) + v9 - 524) > *(short*)(v2 + 54) + *(short*)(v2 + 52) ||
				v24 == *(short*)(v2 + 44)) {
				break;
			}
			v10 = *(uint32_t*)(v2 + 24) + v9;
			if (*(uint32_t*)v10 < *(short*)(v2 + 54)) {
				v21 = v8 + *(unsigned char*)(v10 + 520) + 1;
				goto LABEL_38;
			}
			v11 = *(unsigned char*)(v10 + 520);
			v19 = *(uint32_t*)(v10 + 516);
			v12 = v11 + 1;
			v27 = v11 + 1;
			nox_xxx_drawSetTextColor_434390(v19);
			if (*(uint32_t*)(v2 + 16)) {
				v13 = *(int**)(v2 + 48);
				v14 = *v13;
				if ((int)*v13 >= 0) {
					while (v24 != v14) {
						v14 = v13[1];
						++v13;
						if (v14 < 0) {
							goto LABEL_32;
						}
					}
					if (*(uint32_t*)(a2 + 52) != 0x80000000) {
						nox_client_drawSetColor_434460(*(uint32_t*)(a2 + 52));
						nox_client_drawRectFilledOpaque_49CE30(xLeft, v21, v22, v12);
					}
				}
			} else if (v24 == *(uint32_t*)(v2 + 48)) {
				if (*(uint32_t*)(a2 + 52) != 0x80000000) {
					nox_client_drawSetColor_434460(*(uint32_t*)(a2 + 52));
					nox_client_drawRectFilledOpaque_49CE30(xLeft, v21, v22, v12);
				}
			}
		LABEL_32:
			nox_xxx_drawSetTextColor_434390(*(uint32_t*)(*(uint32_t*)(v2 + 24) + v9 + 516));
			if ((a1[1] & 0x4000) == 0x4000) {
				nox_wcscpy((wchar2_t*)v30, (const wchar2_t*)(*(uint32_t*)(v2 + 24) + v9 + 4));
				v15 = &v30[nox_wcslen((const wchar2_t*)v30)];
				v16 = v22 - 7;
				do {
					v17 = *(uint32_t*)(a2 + 200);
					*v15 = 0;
					--v15;
					nox_xxx_drawGetStringSize_43F840(v17, (unsigned short*)v30, &v29, 0, 0);
				} while (v29 > v16);
				nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), v30, xLeft + 5, v21 + 2, v16, v27);
				v9 = i;
				v12 = v27;
			} else {
				nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), (uint16_t*)(*(uint32_t*)(v2 + 24) + v9 + 4),
										  xLeft + 5, v21 + 2, v22 - 7, v12);
			}
			v21 += v12;
		LABEL_38:
			v8 = v21;
			v9 += 524;
			++v24;
		}
	}
	sub_49F860();
	nox_draw_enableTextSmoothing_43F670(0);
	return 1;
}
// 4A3C50: using guessed type wchar2_t var_200[256];

//----- (004A3FC0) --------------------------------------------------------
int nox_xxx_wndListboxDrawWithImage_4A3FC0(uint32_t* a1, int a2) {
	int v2;         // ecx
	int v3;         // ebp
	int v4;         // esi
	int v5;         // eax
	int v6;         // eax
	int v7;         // edx
	int v8;         // eax
	int v9;         // esi
	int v10;        // ebx
	int v11;        // eax
	int v12;        // ecx
	int v13;        // esi
	int* v14;       // ecx
	int v15;        // eax
	short* v16;     // esi
	int v17;        // ebx
	int v18;        // ecx
	int v20;        // [esp-24h] [ebp-254h]
	int v21;        // [esp-4h] [ebp-234h]
	int v22;        // [esp+10h] [ebp-220h]
	int yTop;       // [esp+14h] [ebp-21Ch]
	int xLeft;      // [esp+18h] [ebp-218h]
	int v25;        // [esp+1Ch] [ebp-214h]
	int v26;        // [esp+20h] [ebp-210h]
	int i;          // [esp+24h] [ebp-20Ch]
	int v28;        // [esp+28h] [ebp-208h]
	int v29;        // [esp+2Ch] [ebp-204h]
	short v30[256]; // [esp+30h] [ebp-200h]

	v2 = *(uint32_t*)(a2 + 24);
	v3 = a1[8];
	v28 = *(uint32_t*)(a2 + 28);
	i = v2;
	nox_client_wndGetPosition_46AA60(a1, &xLeft, &yTop);
	v4 = a1[3];
	v22 = a1[2];
	if ((a1[1] & 0x2000) == 0x2000) {
		nox_draw_enableTextSmoothing_43F670(1);
	}
	nox_xxx_wndDraw_49F7F0();
	if (*(uint32_t*)(v3 + 12)) {
		v22 -= 10;
	}
	if (a1[1] & 8) {
		v5 = i;
	} else {
		v5 = *(uint32_t*)(a2 + 48);
	}
	if (v5) {
		nox_client_drawImageAt_47D2C0(v5, xLeft, yTop);
	}
	if (*(uint16_t*)(a2 + 72)) {
		nox_xxx_drawSetTextColor_434390(*(uint32_t*)(a2 + 68));
		nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), (uint16_t*)(a2 + 72), xLeft + 1, yTop, v22, 0);
		v6 = nox_xxx_guiFontHeightMB_43F320(*(uint32_t*)(a2 + 200));
		v20 = *(uint32_t*)(a2 + 200);
		yTop += v6 + 1;
		v4 += -1 - nox_xxx_guiFontHeightMB_43F320(v20);
	}
	if (v28 != 0x80000000) {
		nox_client_drawSetColor_434460(v28);
		nox_client_drawBorderLines_49CC70(xLeft, yTop, v22, v4);
	}
	nox_client_copyRect_49F6F0(xLeft, yTop, v22, v4);
	v7 = *(short*)(v3 + 54);
	v8 = *(uint32_t*)(a2 + 68);
	v9 = yTop - v7;
	v25 = yTop - v7;
	if (v8 != 0x80000000) {
		v10 = 0;
		v26 = 0;
		for (i = 0;; i = v10) {
			if (v10 > 0 && *(uint32_t*)(*(uint32_t*)(v3 + 24) + v10 - 524) > *(short*)(v3 + 54) + *(short*)(v3 + 52) ||
				v26 == *(short*)(v3 + 44)) {
				break;
			}
			v11 = v10 + *(uint32_t*)(v3 + 24);
			if (*(uint32_t*)v11 < *(short*)(v3 + 54)) {
				v25 = v9 + *(unsigned char*)(v11 + 520) + 1;
				goto LABEL_35;
			}
			v12 = *(unsigned char*)(v11 + 520);
			v21 = *(uint32_t*)(v11 + 516);
			v13 = v12 + 1;
			v28 = v12 + 1;
			nox_xxx_drawSetTextColor_434390(v21);
			if (*(uint32_t*)(v3 + 16)) {
				v14 = *(int**)(v3 + 48);
				v15 = *v14;
				if ((int)*v14 >= 0) {
					while (v26 != v15) {
						v15 = v14[1];
						++v14;
						if (v15 < 0) {
							goto LABEL_29;
						}
					}
					if (*(uint32_t*)(a2 + 52) != 0x80000000) {
						nox_client_drawSetColor_434460(*(uint32_t*)(a2 + 52));
						nox_client_drawRectFilledOpaque_49CE30(xLeft, v25, v22, v13);
					}
				}
			} else if (v26 == *(uint32_t*)(v3 + 48)) {
				if (*(uint32_t*)(a2 + 52) != 0x80000000) {
					nox_client_drawSetColor_434460(*(uint32_t*)(a2 + 52));
					nox_client_drawRectFilledOpaque_49CE30(xLeft, v25, v22, v13);
				}
			}
		LABEL_29:
			nox_xxx_drawSetTextColor_434390(*(uint32_t*)(v10 + *(uint32_t*)(v3 + 24) + 516));
			if ((a1[1] & 0x4000) == 0x4000) {
				nox_wcscpy((wchar2_t*)v30, (const wchar2_t*)(v10 + *(uint32_t*)(v3 + 24) + 4));
				v16 = &v30[nox_wcslen((const wchar2_t*)v30)];
				v17 = v22 - 7;
				do {
					v18 = *(uint32_t*)(a2 + 200);
					*v16 = 0;
					--v16;
					nox_xxx_drawGetStringSize_43F840(v18, (unsigned short*)v30, &v29, 0, 0);
				} while (v29 > v17);
				nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), v30, xLeft + 5, v25 + 2, v17, v28);
				v10 = i;
				v13 = v28;
			} else {
				nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(a2 + 200), (uint16_t*)(v10 + *(uint32_t*)(v3 + 24) + 4),
										  xLeft + 5, v25 + 2, v22 - 7, v13);
			}
			v25 += v13;
		LABEL_35:
			v9 = v25;
			v10 += 524;
			++v26;
		}
	}
	sub_49F860();
	nox_draw_enableTextSmoothing_43F670(0);
	return 1;
}
// 4A3FC0: using guessed type wchar2_t var_200[256];

//----- (004A4800) --------------------------------------------------------
int sub_4A4800(int a1) {
	int result;   // eax
	uint32_t* v2; // ecx
	int v3;       // esi
	int v4;       // edi

	result = 0;
	v2 = *(uint32_t**)(a1 + 24);
	v3 = *(short*)(a1 + 54);
	if (*v2 <= v3) {
		while (result < *(short*)(a1 + 44)) {
			v4 = v2[131];
			v2 += 131;
			++result;
			if (v4 > v3) {
				return result;
			}
		}
		result = 0;
	}
	return result;
}

//----- (004A4840) --------------------------------------------------------
int nox_game_showSelClass_4A4840() {
	int result;   // eax
	uint32_t* v1; // eax
	uint32_t* v2; // eax
	uint32_t* v3; // eax

	sub_5007E0("*:*");
	sub_4A1BE0(1);
	dword_5d4594_1307724 = nox_xxx_getHostInfoPtr_431770();
	nox_game_addStateCode_43BDD0(600);
	result = nox_new_window_from_file("SelClass.wnd", sub_4A4A20);
	dword_5d4594_1307736 = result;
	if (result) {
		nox_xxx_wndSetWindowProc_46B300(result, sub_4A18E0);
		result = nox_gui_makeAnimation_43C5B0(*(uint32_t**)&dword_5d4594_1307736, 0, 0, 0, -460, 0, 20, 0, -40);
		nox_wnd_xxx_1307732 = result;
		if (result) {
			nox_wnd_xxx_1307732->field_0 = 600;
			nox_wnd_xxx_1307732->field_12 = sub_4A4970;
			nox_wnd_xxx_1307732->fnc_done_out = sub_4A49A0;
			v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1307736, 601);
			nox_xxx_wndSetDrawFn_46B340((int)v1, sub_4A49D0);
			v2 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1307736, 603);
			nox_xxx_wndSetDrawFn_46B340((int)v2, sub_4A49D0);
			v3 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1307736, 602);
			nox_xxx_wndSetDrawFn_46B340((int)v3, sub_4A49D0);
			*getMemU32Ptr(0x5D4594, 1307728) = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1307736, 610);
			nox_xxx_wndRetNULL_46A8A0();
			*getMemU32Ptr(0x5D4594, 1307740) = 0;
			sub_4A19F0("OptsBack.wnd:Back");
			sub_4602F0();
			result = 1;
		}
	}
	return result;
}
// 4A18E0: using guessed type int  sub_4A18E0(int, int, int, int);

//----- (004A4970) --------------------------------------------------------
int sub_4A4970() {
	nox_wnd_xxx_1307732->state = NOX_GUI_ANIM_OUT;
	sub_43BE40(2);
	nox_xxx_clientPlaySoundSpecial_452D80(923, 100);
	return 1;
}

//----- (004A49A0) --------------------------------------------------------
int sub_4A49A0() {
	int (*v0)(void); // esi

	v0 = nox_wnd_xxx_1307732->field_13;
	nox_gui_freeAnimation_43C570(nox_wnd_xxx_1307732);
	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1307736);
	v0();
	return 1;
}

//----- (004A49D0) --------------------------------------------------------
int sub_4A49D0(int yTop, int a2) {
	uint32_t* v1; // esi
	int xLeft;    // [esp+4h] [ebp-4h]

	v1 = (uint32_t*)yTop;
	if (*getMemU32Ptr(0x5D4594, 1307740) != *(uint32_t*)yTop) {
		nox_client_wndGetPosition_46AA60((uint32_t*)yTop, &xLeft, &yTop);
		nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, v1[6] - v1[4], v1[7] - v1[5]);
	}
	return 1;
}

//----- (004A5E90) --------------------------------------------------------
void sub_4A5E90_A();
int sub_4A5E90() {
	const char** i; // eax
	uint32_t* v1;   // eax
	uint32_t* v2;   // eax
	uint32_t* v3;   // eax
	uint32_t* v4;   // eax
	uint32_t* v5;   // eax
	uint32_t* v6;   // eax
	uint32_t* v7;   // eax
	uint32_t* v8;   // eax
	uint32_t* v9;   // eax
	uint32_t* v10;  // eax
	int result;     // eax

	if (*getMemU32Ptr(0x587000, 171384)) {
		sub_4A62B0();
	}
	sub_4A5E90_A();
	v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 720);
	dword_5d4594_1308096 = v1;
	v1[8] = 131074;
	v2 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 721);
	dword_5d4594_1308100 = v2;
	v2[8] = 589825;
	v3 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 722);
	dword_5d4594_1308104 = v3;
	v3[8] = 589825;
	v4 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 723);
	dword_5d4594_1308108 = v4;
	v4[8] = 589825;
	v5 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 724);
	dword_5d4594_1308112 = v5;
	v5[8] = 589825;
	v6 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 725);
	dword_5d4594_1308116 = v6;
	v6[8] = *getMemU16Ptr(0x587000, 171372) << 16;
	v7 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 726);
	dword_5d4594_1308120 = v7;
	v7[8] = *getMemU16Ptr(0x587000, 171374) << 16;
	v8 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 727);
	dword_5d4594_1308124 = v8;
	v8[8] = *getMemU16Ptr(0x587000, 171376) << 16;
	v9 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 728);
	dword_5d4594_1308128 = v9;
	v9[8] = *getMemU16Ptr(0x587000, 171378) << 16;
	v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 729);
	dword_5d4594_1308132 = v10;
	v10[8] = *getMemU16Ptr(0x587000, 171380) << 16;
	dword_5d4594_1308136 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 711);
	dword_5d4594_1308140 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 712);
	dword_5d4594_1308144 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 713);
	dword_5d4594_1308148 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 714);
	dword_5d4594_1308152 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, 751);
	result = dword_587000_171388;
	if (!dword_587000_171388) {
		nox_window_call_field_94(*(int*)&dword_5d4594_1308152, 16414, *(int*)&dword_5d4594_1307784, 0);
		sub_4A61E0(*(uint32_t**)&dword_5d4594_1308096, 2, (uint8_t*)(dword_5d4594_1307784 + 71));
		sub_4A61E0(*(uint32_t**)&dword_5d4594_1308100, 1, (uint8_t*)(dword_5d4594_1307784 + 68));
		sub_4A61E0(*(uint32_t**)&dword_5d4594_1308104, 1, (uint8_t*)(dword_5d4594_1307784 + 74));
		sub_4A61E0(*(uint32_t**)&dword_5d4594_1308108, 1, (uint8_t*)(dword_5d4594_1307784 + 77));
		sub_4A61E0(*(uint32_t**)&dword_5d4594_1308112, 1, (uint8_t*)(dword_5d4594_1307784 + 80));
		*(uint32_t*)(dword_5d4594_1308116 + 32) = *(unsigned char*)(dword_5d4594_1307784 + 83) << 16;
		*(uint32_t*)(dword_5d4594_1308120 + 32) = *(unsigned char*)(dword_5d4594_1307784 + 84) << 16;
		*(uint32_t*)(dword_5d4594_1308124 + 32) = *(unsigned char*)(dword_5d4594_1307784 + 85) << 16;
		*(uint32_t*)(dword_5d4594_1308128 + 32) = *(unsigned char*)(dword_5d4594_1307784 + 86) << 16;
		result = *(unsigned char*)(dword_5d4594_1307784 + 87) << 16;
		*(uint32_t*)(dword_5d4594_1308132 + 32) = result;
	}
	return result;
}

//----- (004A61E0) --------------------------------------------------------
unsigned char* sub_4A61E0(uint32_t* a1, int a2, unsigned char* a3) {
	unsigned int v3;       // esi
	unsigned char* result; // eax
	uint32_t* v5;          // eax
	uint32_t* v6;          // eax

	v3 = 0;
	result = getMemAt(0x5D4594, 1307797 + 96 * a2);
	while (*(result - 1) != *a3 || *result != a3[1] || result[1] != a3[2]) {
		++v3;
		result += 3;
		if (v3 >= 0x20) {
			goto LABEL_9;
		}
	}
	if (a2 == 1) {
		v5 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, *a1 - 10);
		nox_xxx_wnd_46ABB0((int)v5, 1);
		result = (unsigned char*)nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, *a1 + 10);
		*((uint32_t*)result + 9) |= 6u;
	}
LABEL_9:
	if (v3 == 32 && a2 == 1) {
		v6 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, *a1 - 10);
		result = (unsigned char*)nox_xxx_wnd_46ABB0((int)v6, 0);
		LOWORD(v3) = 9;
	}
	a1[8] = (unsigned short)a2 | ((unsigned short)v3 << 16);
	return result;
}

//----- (004A62B0) --------------------------------------------------------
void* nox_xxx_modifGetModifListByType_4133B0(int a1);
void sub_4A62B0() {
	unsigned char* v0; // esi
	int v1;            // eax
	int v3;            // ebp
	unsigned char* v5; // [esp+10h] [ebp-18h]
	int v6;            // [esp+14h] [ebp-14h]
	char v7[16];       // [esp+18h] [ebp-10h]

	v0 = getMemAt(0x5D4594, 1307796);
	v1 = 0;
	v5 = getMemAt(0x5D4594, 1307796);
	do {
		*v0 = 0;
		v6 = v1 + 1;
		v0[1] = 0;
		v0[2] = 0;
		nox_sprintf(v7, "UserColor%d", v1 + 1);
		obj_412ae0_t* v2 = nox_xxx_modifGetModifListByType_4133B0(2);
		if (v2) {
			while (strcmp(v7, v2->field_0)) {
				v2 = nox_xxx_modifNext_4133C0(v2);
				if (!v2) {
					goto LABEL_7;
				}
			}
			v3 = (int)(&v2->field_6);
			*(uint16_t*)v5 = *(uint16_t*)v3;
			v5[2] = *(uint8_t*)(v3 + 2);
		LABEL_7:
			v0 = v5;
		}
		v1 = v6;
		v0 += 3;
		v5 = v0;
	} while ((int)v0 < (int)getMemAt(0x5D4594, 1307892));
	*getMemU8Ptr(0x5D4594, 1307989) = -34;
	*getMemU8Ptr(0x5D4594, 1307991) = -34;
	*getMemU8Ptr(0x5D4594, 1307995) = -102;
	*getMemU8Ptr(0x5D4594, 1308000) = -102;
	*getMemU8Ptr(0x5D4594, 1308038) = -102;
	*getMemU8Ptr(0x5D4594, 1307988) = -23;
	*getMemU8Ptr(0x5D4594, 1307990) = -51;
	*getMemU8Ptr(0x5D4594, 1307992) = -74;
	*getMemU8Ptr(0x5D4594, 1307993) = -88;
	*getMemU8Ptr(0x5D4594, 1307994) = -38;
	*getMemU8Ptr(0x5D4594, 1307996) = 110;
	*getMemU8Ptr(0x5D4594, 1307997) = -20;
	*getMemU8Ptr(0x5D4594, 1307998) = -54;
	*getMemU8Ptr(0x5D4594, 1307999) = -78;
	*getMemU8Ptr(0x5D4594, 1308001) = -125;
	*getMemU8Ptr(0x5D4594, 1308002) = 105;
	*getMemU8Ptr(0x5D4594, 1308003) = 117;
	*getMemU8Ptr(0x5D4594, 1308004) = 95;
	*getMemU8Ptr(0x5D4594, 1308005) = 74;
	*getMemU8Ptr(0x5D4594, 1308006) = -17;
	*getMemU8Ptr(0x5D4594, 1308007) = -23;
	*getMemU8Ptr(0x5D4594, 1308008) = -63;
	*getMemU8Ptr(0x5D4594, 1308009) = -28;
	*getMemU8Ptr(0x5D4594, 1308010) = -28;
	*getMemU8Ptr(0x5D4594, 1308011) = -63;
	*getMemU8Ptr(0x5D4594, 1308012) = -53;
	*getMemU8Ptr(0x5D4594, 1308013) = -71;
	*getMemU8Ptr(0x5D4594, 1308014) = -100;
	*getMemU8Ptr(0x5D4594, 1308015) = -56;
	*getMemU8Ptr(0x5D4594, 1308016) = -92;
	*getMemU8Ptr(0x5D4594, 1308017) = 110;
	*getMemU8Ptr(0x5D4594, 1308018) = -87;
	*getMemU8Ptr(0x5D4594, 1308019) = -125;
	*getMemU8Ptr(0x5D4594, 1308020) = 79;
	*getMemU8Ptr(0x5D4594, 1308021) = -13;
	*getMemU8Ptr(0x5D4594, 1308022) = -98;
	*getMemU8Ptr(0x5D4594, 1308023) = 119;
	*getMemU8Ptr(0x5D4594, 1308024) = 115;
	*getMemU8Ptr(0x5D4594, 1308025) = 77;
	*getMemU8Ptr(0x5D4594, 1308026) = 34;
	*getMemU8Ptr(0x5D4594, 1308027) = 91;
	*getMemU8Ptr(0x5D4594, 1308028) = 55;
	*getMemU8Ptr(0x5D4594, 1308029) = 20;
	*getMemU8Ptr(0x5D4594, 1308030) = -7;
	*getMemU8Ptr(0x5D4594, 1308031) = -51;
	*getMemU8Ptr(0x5D4594, 1308032) = -118;
	*getMemU8Ptr(0x5D4594, 1308033) = -16;
	*getMemU8Ptr(0x5D4594, 1308034) = -21;
	*getMemU8Ptr(0x5D4594, 1308035) = -85;
	*getMemU8Ptr(0x5D4594, 1308036) = -122;
	*getMemU8Ptr(0x5D4594, 1308037) = -113;
	*getMemU8Ptr(0x5D4594, 1308039) = -121;
	*getMemU8Ptr(0x5D4594, 1308040) = 102;
	*getMemU8Ptr(0x5D4594, 1308041) = 67;
	*getMemU8Ptr(0x5D4594, 1308042) = 115;
	*getMemU8Ptr(0x5D4594, 1308043) = 80;
	*getMemU8Ptr(0x5D4594, 1308044) = 46;
	*getMemU8Ptr(0x5D4594, 1308045) = -91;
	*getMemU8Ptr(0x5D4594, 1308046) = 93;
	*getMemU8Ptr(0x5D4594, 1308047) = 70;
	*getMemU8Ptr(0x5D4594, 1308048) = 99;
	*getMemU8Ptr(0x5D4594, 1308049) = 65;
	*getMemU8Ptr(0x5D4594, 1308050) = 37;
	*getMemU8Ptr(0x5D4594, 1308051) = 83;
	*getMemU8Ptr(0x5D4594, 1308052) = 52;
	*getMemU8Ptr(0x5D4594, 1308053) = 29;
	*getMemU8Ptr(0x5D4594, 1308054) = -88;
	*getMemU8Ptr(0x5D4594, 1308055) = 127;
	*getMemU8Ptr(0x5D4594, 1308056) = 84;
	*getMemU8Ptr(0x5D4594, 1308057) = -85;
	*getMemU8Ptr(0x5D4594, 1308058) = -96;
	*getMemU8Ptr(0x5D4594, 1308059) = 96;
	*getMemU8Ptr(0x5D4594, 1308060) = -115;
	*getMemU8Ptr(0x5D4594, 1308061) = -120;
	*getMemU8Ptr(0x5D4594, 1308062) = 118;
	*getMemU8Ptr(0x5D4594, 1308063) = 127;
	*getMemU8Ptr(0x5D4594, 1308065) = 82;
	*getMemU8Ptr(0x5D4594, 1308064) = 116;
	*getMemU8Ptr(0x5D4594, 1308080) = 116;
	*getMemU8Ptr(0x5D4594, 1308082) = 116;
	*getMemU8Ptr(0x5D4594, 1307898) = -1;
	*getMemU8Ptr(0x5D4594, 1307904) = -1;
	*getMemU8Ptr(0x5D4594, 1308066) = 96;
	*getMemU8Ptr(0x5D4594, 1307908) = 68;
	*getMemU8Ptr(0x5D4594, 1307909) = 68;
	*getMemU8Ptr(0x5D4594, 1307926) = 68;
	*getMemU8Ptr(0x5D4594, 1307936) = 68;
	*getMemU8Ptr(0x5D4594, 1308067) = 84;
	*getMemU8Ptr(0x5D4594, 1307937) = -42;
	*getMemU8Ptr(0x5D4594, 1307938) = -42;
	*getMemU8Ptr(0x5D4594, 1307939) = -42;
	*getMemU8Ptr(0x5D4594, 1308068) = 51;
	*getMemU8Ptr(0x5D4594, 1308069) = -94;
	*getMemU8Ptr(0x5D4594, 1308070) = -127;
	*getMemU8Ptr(0x5D4594, 1308071) = 100;
	*getMemU8Ptr(0x5D4594, 1308072) = 78;
	*getMemU8Ptr(0x5D4594, 1308073) = 64;
	*getMemU8Ptr(0x5D4594, 1308074) = 46;
	*getMemU8Ptr(0x5D4594, 1308075) = 65;
	*getMemU8Ptr(0x5D4594, 1308076) = 50;
	*getMemU8Ptr(0x5D4594, 1308077) = 45;
	*getMemU8Ptr(0x5D4594, 1308078) = -85;
	*getMemU8Ptr(0x5D4594, 1308079) = -98;
	*getMemU8Ptr(0x5D4594, 1308081) = -81;
	*getMemU8Ptr(0x5D4594, 1308083) = -125;
	*getMemU8Ptr(0x5D4594, 1307892) = -13;
	*getMemU8Ptr(0x5D4594, 1307893) = -73;
	*getMemU8Ptr(0x5D4594, 1307894) = -97;
	*getMemU8Ptr(0x5D4594, 1307895) = -57;
	*getMemU8Ptr(0x5D4594, 1307896) = -124;
	*getMemU8Ptr(0x5D4594, 1307897) = 58;
	*getMemU8Ptr(0x5D4594, 1307899) = -86;
	*getMemU8Ptr(0x5D4594, 1307900) = 86;
	*getMemU8Ptr(0x5D4594, 1307901) = -110;
	*getMemU8Ptr(0x5D4594, 1307902) = 121;
	*getMemU8Ptr(0x5D4594, 1307903) = 106;
	*getMemU8Ptr(0x5D4594, 1307905) = -5;
	*getMemU8Ptr(0x5D4594, 1307906) = -70;
	*getMemU8Ptr(0x5D4594, 1307907) = -9;
	*getMemU8Ptr(0x5D4594, 1307910) = 63;
	*getMemU8Ptr(0x5D4594, 1307911) = -69;
	*getMemU8Ptr(0x5D4594, 1307912) = -17;
	*getMemU8Ptr(0x5D4594, 1307913) = -60;
	*getMemU8Ptr(0x5D4594, 1307914) = -62;
	*getMemU8Ptr(0x5D4594, 1307915) = -79;
	*getMemU8Ptr(0x5D4594, 1307916) = -3;
	*getMemU8Ptr(0x5D4594, 1307917) = -10;
	*getMemU8Ptr(0x5D4594, 1307918) = 102;
	*getMemU8Ptr(0x5D4594, 1307919) = 115;
	*getMemU8Ptr(0x5D4594, 1307920) = 77;
	*getMemU8Ptr(0x5D4594, 1307921) = 34;
	*getMemU8Ptr(0x5D4594, 1307922) = -27;
	*getMemU8Ptr(0x5D4594, 1307923) = 121;
	*getMemU8Ptr(0x5D4594, 1307924) = 25;
	*getMemU8Ptr(0x5D4594, 1307925) = 82;
	*getMemU8Ptr(0x5D4594, 1307927) = 60;
	*getMemU8Ptr(0x5D4594, 1307928) = -21;
	*getMemU8Ptr(0x5D4594, 1307929) = -27;
	*getMemU8Ptr(0x5D4594, 1307930) = -119;
	*getMemU8Ptr(0x5D4594, 1307931) = -109;
	*getMemU8Ptr(0x5D4594, 1307932) = 0;
	*getMemU8Ptr(0x5D4594, 1307933) = 0;
	*getMemU8Ptr(0x5D4594, 1307934) = 0;
	*getMemU8Ptr(0x5D4594, 1307935) = 37;
	*getMemU8Ptr(0x5D4594, 1307940) = -107;
	*getMemU8Ptr(0x5D4594, 1307941) = 108;
	*getMemU8Ptr(0x5D4594, 1307942) = 43;
	*getMemU8Ptr(0x5D4594, 1307943) = 52;
	*getMemU8Ptr(0x5D4594, 1307944) = 25;
	*getMemU8Ptr(0x5D4594, 1307945) = 0;
	*getMemU8Ptr(0x5D4594, 1307946) = -106;
	*getMemU8Ptr(0x5D4594, 1307947) = 55;
	*getMemU8Ptr(0x5D4594, 1307952) = -106;
	*getMemU8Ptr(0x5D4594, 1307955) = -125;
	*getMemU8Ptr(0x5D4594, 1307956) = 104;
	*getMemU8Ptr(0x5D4594, 1307961) = 117;
	*getMemU8Ptr(0x5D4594, 1307962) = 117;
	*getMemU8Ptr(0x5D4594, 1307963) = 117;
	*getMemU8Ptr(0x5D4594, 1307981) = 104;
	*getMemU8Ptr(0x5D4594, 1307948) = 0;
	*getMemU8Ptr(0x5D4594, 1307950) = 0;
	*getMemU8Ptr(0x5D4594, 1307984) = 0;
	*getMemU32Ptr(0x587000, 171384) = 0;
	*getMemU8Ptr(0x5D4594, 1307949) = 9;
	*getMemU8Ptr(0x5D4594, 1307951) = 12;
	*getMemU8Ptr(0x5D4594, 1307953) = -110;
	*getMemU8Ptr(0x5D4594, 1307954) = 81;
	*getMemU8Ptr(0x5D4594, 1307957) = -77;
	*getMemU8Ptr(0x5D4594, 1307958) = 121;
	*getMemU8Ptr(0x5D4594, 1307959) = -91;
	*getMemU8Ptr(0x5D4594, 1307960) = 66;
	*getMemU8Ptr(0x5D4594, 1307964) = 100;
	*getMemU8Ptr(0x5D4594, 1307965) = -126;
	*getMemU8Ptr(0x5D4594, 1307966) = 125;
	*getMemU8Ptr(0x5D4594, 1307967) = 56;
	*getMemU8Ptr(0x5D4594, 1307968) = 48;
	*getMemU8Ptr(0x5D4594, 1307969) = 40;
	*getMemU8Ptr(0x5D4594, 1307970) = 108;
	*getMemU8Ptr(0x5D4594, 1307971) = 81;
	*getMemU8Ptr(0x5D4594, 1307972) = 58;
	*getMemU8Ptr(0x5D4594, 1307973) = 52;
	*getMemU8Ptr(0x5D4594, 1307974) = 40;
	*getMemU8Ptr(0x5D4594, 1307975) = 39;
	*getMemU8Ptr(0x5D4594, 1307976) = -113;
	*getMemU8Ptr(0x5D4594, 1307977) = -114;
	*getMemU8Ptr(0x5D4594, 1307978) = 124;
	*getMemU8Ptr(0x5D4594, 1307979) = 56;
	*getMemU8Ptr(0x5D4594, 1307980) = 29;
	*getMemU8Ptr(0x5D4594, 1307982) = 15;
	*getMemU8Ptr(0x5D4594, 1307983) = 88;
	*getMemU8Ptr(0x5D4594, 1307985) = 87;
	*getMemU8Ptr(0x5D4594, 1307986) = 53;
	*getMemU8Ptr(0x5D4594, 1307987) = 53;
}

//----- (004A6890) --------------------------------------------------------
int sub_4A6890() {
	nox_wnd_xxx_1308092->state = NOX_GUI_ANIM_OUT;
	sub_43BE40(2);
	nox_xxx_clientPlaySoundSpecial_452D80(923, 100);
	sub_4A68C0();
	return 1;
}

//----- (004A6B50) --------------------------------------------------------
int sub_4A6B50(wchar2_t* a1) {
	wchar2_t* v1;    // esi
	int v2;         // ebp
	int v3;         // ebx
	signed int v4;  // eax
	short* v5;      // edi
	wchar2_t v6;     // ax
	signed int v7;  // esi
	wchar2_t* i;     // edi
	int v10;        // [esp+10h] [ebp-3Ch]
	signed int v11; // [esp+14h] [ebp-38h]
	short v12[26];  // [esp+18h] [ebp-34h]

	v1 = a1;
	v2 = 0;
	v3 = 1;
	v10 = 0;
	v4 = nox_wcslen(a1);
	if (v4 >= 1) {
		v5 = v12;
		v11 = v4;
		do {
			if (iswspace(*v1)) {
				if (!v3) {
					*v5 = *v1;
					++v5;
					++v10;
				}
			} else {
				if (v3 == 1) {
					v6 = *v1;
					if (*v1 == 42 || v6 == 63 || v6 == 60 || v6 == 62 || v6 == 92 || v6 == 47 || v6 == 58 || v6 == 34 ||
						v6 == 124) {
						*v1 = 45;
					}
				}
				*v5 = *v1;
				++v5;
				++v10;
				v3 = 0;
				v2 = 1;
			}
			++v1;
			--v11;
		} while (v11);
		v12[v10] = 0;
		if (v2) {
			nox_wcscpy(a1, (const wchar2_t*)v12);
			v7 = nox_wcslen((const wchar2_t*)v12) - 1;
			if ((int)v7 >= 0) {
				for (i = &a1[v7]; iswspace(*i); --i) {
					if (--v7 < 0) {
						return v2;
					}
				}
				a1[v7 + 1] = 0;
			}
		}
	}
	return v2;
}
// 4A6B50: using guessed type wchar2_t var_34[26];

//----- (004A6C90) --------------------------------------------------------
int sub_4A6C90() {
	int (*v0)(void); // esi

	v0 = nox_wnd_xxx_1308092->field_13;
	nox_gui_freeAnimation_43C570(nox_wnd_xxx_1308092);
	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1308084);
	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1308088);
	if (v0) {
		v0();
	} else {
		nox_client_resetScreenParticles_431510();
		nox_gui_draw();
		if (!nox_common_gameFlags_check_40A5C0(0x2000)) {
			nox_client_guiXxxDestroy_4A24A0();
			return 1;
		}
		nox_xxx_serverHost_43B4D0();
		return 1;
	}
	return 1;
}

//----- (004A6D20) --------------------------------------------------------
int sub_4A6D20(int a1, int a2) {
	int v1;            // esi
	unsigned short v2; // di
	int v3;            // ebx
	int v4;            // ecx
	int v5;            // edx
	int v6;            // edi
	int v7;            // ebp
	int v8;            // eax
	int v10;           // [esp-10h] [ebp-28h]
	int yTop;          // [esp+10h] [ebp-8h]
	int xLeft;         // [esp+14h] [ebp-4h]
	int v13;           // [esp+1Ch] [ebp+4h]

	v1 = a1;
	v2 = *(uint16_t*)(a1 + 32);
	v3 = *(uint32_t*)(a1 + 32) >> 16;
	nox_client_wndGetPosition_46AA60((uint32_t*)a1, &xLeft, &yTop);
	v4 = *(uint32_t*)(a1 + 20);
	v5 = *(uint32_t*)(a1 + 16);
	v6 = (unsigned short)v3 + 32 * v2;
	v13 = *(uint32_t*)(a1 + 28) - v4;
	v7 = *(uint32_t*)(v1 + 24) - v5;
	LOBYTE(v5) = getMemByte(0x5D4594, 1307798 + 3 * v6);
	LOBYTE(v4) = getMemByte(0x5D4594, 1307797 + 3 * v6);
	v10 = v5;
	LOBYTE(v5) = getMemByte(0x5D4594, 1307796 + 3 * v6);
	v8 = nox_color_rgb_4344A0(v5, v4, v10);
	nox_client_drawSetColor_434460(v8);
	nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, v7, v13);
	return 1;
}

//----- (004A6DC0) --------------------------------------------------------
int sub_4A6DC0(uint32_t* a1, int a2) {
	int v1;   // eax
	int v2;   // ecx
	int v3;   // edx
	int v4;   // esi
	int v5;   // eax
	int v6;   // ecx
	int v7;   // edx
	int v8;   // eax
	int v9;   // edx
	int v10;  // eax
	int v11;  // ecx
	int v12;  // eax
	int v13;  // edx
	int v14;  // ecx
	int v15;  // eax
	int v16;  // eax
	int v17;  // edx
	int v18;  // ecx
	int v19;  // eax
	int v20;  // eax
	int v21;  // ebx
	int v22;  // eax
	int v23;  // edx
	int* v24; // ebp
	int v25;  // ecx
	int v26;  // eax
	int v27;  // edi
	int i;    // esi
	int v29;  // eax
	int v30;  // eax
	int v31;  // ecx
	int v32;  // eax
	int v33;  // edi
	int j;    // esi
	int v35;  // eax
	int v36;  // eax
	int v37;  // ecx
	int v38;  // eax
	int v39;  // eax
	int v40;  // ecx
	int v41;  // edx
	int v42;  // edi
	int k;    // esi
	int v44;  // eax
	int v45;  // eax
	int v46;  // ecx
	int v47;  // eax
	int v48;  // eax
	int v49;  // ecx
	int v50;  // edx
	int v52;  // [esp-24h] [ebp-3Ch]
	int v53;  // [esp-24h] [ebp-3Ch]
	int v54;  // [esp+10h] [ebp-8h]
	int v55;  // [esp+14h] [ebp-4h]

	nox_client_wndGetPosition_46AA60(a1, &v54, &v55);
	v1 = ((unsigned short)(*(uint32_t*)(dword_5d4594_1308096 + 32) >> 16)) +
		 32 * *(unsigned short*)(dword_5d4594_1308096 + 32);
	v2 = getMemByte(0x5D4594, 1307798 + 3 * v1);
	LOBYTE(v3) = getMemByte(0x5D4594, 1307797 + 3 * v1);
	LOBYTE(v1) = getMemByte(0x5D4594, 1307796 + 3 * v1);
	v4 = nox_color_rgb_4344A0(v1, v3, v2);
	nox_draw_setMaterial_4341D0(6, v4);
	v5 = ((unsigned short)(*(uint32_t*)(dword_5d4594_1308100 + 32) >> 16)) +
		 32 * *(unsigned short*)(dword_5d4594_1308100 + 32);
	v6 = getMemByte(0x5D4594, 1307798 + 3 * v5);
	LOBYTE(v7) = getMemByte(0x5D4594, 1307797 + 3 * v5);
	LOBYTE(v5) = getMemByte(0x5D4594, 1307796 + 3 * v5);
	v8 = nox_color_rgb_4344A0(v5, v7, v6);
	nox_draw_setMaterial_4341D0(1, v8);
	v9 = dword_5d4594_1308144;
	if (*(uint8_t*)(dword_5d4594_1308144 + 4) & 8) {
		v10 = (unsigned short)(*(uint32_t*)(dword_5d4594_1308108 + 32) >> 16);
		v11 = v10 + 32 * *(unsigned short*)(dword_5d4594_1308108 + 32);
		LOBYTE(v9) = getMemByte(0x5D4594, 1307797 + 3 * v11);
		v12 = nox_color_rgb_4344A0(
			getMemByte(0x5D4594, 1307796 + 3 * v11), v9,
			getMemByte(0x5D4594, 1307798 + 3 * (v10 + 32 * *(unsigned short*)(dword_5d4594_1308108 + 32))));
		nox_draw_setMaterial_4341D0(2, v12);
	} else {
		nox_draw_setMaterial_4341D0(2, v4);
	}
	nox_draw_setMaterial_4341D0(3, v4);
	v13 = dword_5d4594_1308148;
	if (*(uint8_t*)(dword_5d4594_1308148 + 4) & 8) {
		v14 = (unsigned short)(*(uint32_t*)(dword_5d4594_1308112 + 32) >> 16);
		v15 = v14 + 32 * *(unsigned short*)(dword_5d4594_1308112 + 32);
		LOBYTE(v13) = getMemByte(0x5D4594, 1307797 + 3 * v15);
		LOBYTE(v15) = getMemByte(0x5D4594, 1307796 + 3 * v15);
		v16 = nox_color_rgb_4344A0(
			v15, v13, getMemByte(0x5D4594, 1307798 + 3 * (v14 + 32 * *(unsigned short*)(dword_5d4594_1308112 + 32))));
		nox_draw_setMaterial_4341D0(4, v16);
	} else {
		nox_draw_setMaterial_4341D0(4, v4);
	}
	v17 = dword_5d4594_1308140;
	if (*(uint8_t*)(dword_5d4594_1308140 + 4) & 8) {
		v18 = (unsigned short)(*(uint32_t*)(dword_5d4594_1308104 + 32) >> 16);
		v19 = v18 + 32 * *(unsigned short*)(dword_5d4594_1308104 + 32);
		LOBYTE(v17) = getMemByte(0x5D4594, 1307797 + 3 * v19);
		LOBYTE(v19) = getMemByte(0x5D4594, 1307796 + 3 * v19);
		v20 = nox_color_rgb_4344A0(
			v19, v17, getMemByte(0x5D4594, 1307798 + 3 * (v18 + 32 * *(unsigned short*)(dword_5d4594_1308104 + 32))));
		nox_draw_setMaterial_4341D0(5, v20);
	} else {
		nox_draw_setMaterial_4341D0(5, v4);
	}
	if (*(uint8_t*)(dword_5d4594_1308136 + 4) & 8) {
		nox_client_drawImageAt_47D2C0(*getMemU32Ptr(0x973A20, 16 + 4 * *(unsigned char*)(dword_5d4594_1307784 + 67)),
									  v54, v55);
	} else {
		nox_client_drawImageAt_47D2C0(*getMemU32Ptr(0x973A20, 24 + 4 * *(unsigned char*)(dword_5d4594_1307784 + 67)),
									  v54, v55);
	}
	v21 = 0;
	v22 = *(unsigned char*)(dword_5d4594_1307784 + 67);
	v23 = 3 * v22;
	v24 = getMemIntPtr(0x973A20, 32 + 104 * v22);
	do {
		v25 = v21;
		v26 = 1 << v21;
		if (1 << v21 == 4) {
			v27 = 1;
			for (i = 3; i < 21; i += 3) {
				v29 = dword_5d4594_1308156;
				LOBYTE(v25) = *(uint8_t*)(i + (uint32_t)dword_5d4594_1308156 + 14);
				LOBYTE(v23) = *(uint8_t*)(i + (uint32_t)dword_5d4594_1308156 + 13);
				LOBYTE(v29) = *(uint8_t*)(i + (uint32_t)dword_5d4594_1308156 + 12);
				nox_draw_setMaterial_4340A0(v27++, v29, v23, v25);
			}
			v30 = ((unsigned short)(*(uint32_t*)(dword_5d4594_1308116 + 32) >> 16)) +
				  32 * *(unsigned short*)(dword_5d4594_1308116 + 32);
			v31 = getMemByte(0x5D4594, 1307798 + 3 * v30);
			LOBYTE(v23) = getMemByte(0x5D4594, 1307797 + 3 * v30);
			LOBYTE(v30) = getMemByte(0x5D4594, 1307796 + 3 * v30);
			v32 = nox_color_rgb_4344A0(v30, v23, v31);
			nox_draw_setMaterial_4341D0(*(uint32_t*)((uint32_t)dword_5d4594_1308156 + 40), v32);
		} else if (v26 == 1024) {
			v33 = 1;
			for (j = 3; j < 21; j += 3) {
				v35 = dword_5d4594_1308160;
				LOBYTE(v25) = *(uint8_t*)((uint32_t)dword_5d4594_1308160 + j + 14);
				LOBYTE(v23) = *(uint8_t*)((uint32_t)dword_5d4594_1308160 + j + 13);
				LOBYTE(v35) = *(uint8_t*)((uint32_t)dword_5d4594_1308160 + j + 12);
				nox_draw_setMaterial_4340A0(v33++, v35, v23, v25);
			}
			v36 = ((unsigned short)(*(uint32_t*)(dword_5d4594_1308120 + 32) >> 16)) +
				  32 * *(unsigned short*)(dword_5d4594_1308120 + 32);
			v37 = getMemByte(0x5D4594, 1307798 + 3 * v36);
			LOBYTE(v23) = getMemByte(0x5D4594, 1307797 + 3 * v36);
			LOBYTE(v36) = getMemByte(0x5D4594, 1307796 + 3 * v36);
			v38 = nox_color_rgb_4344A0(v36, v23, v37);
			nox_draw_setMaterial_4341D0(*(uint32_t*)((uint32_t)dword_5d4594_1308160 + 40), v38);
			v39 = ((unsigned short)(*(uint32_t*)(dword_5d4594_1308124 + 32) >> 16)) +
				  32 * *(unsigned short*)(dword_5d4594_1308124 + 32);
			v40 = getMemByte(0x5D4594, 1307798 + 3 * v39);
			LOBYTE(v41) = getMemByte(0x5D4594, 1307797 + 3 * v39);
			LOBYTE(v39) = getMemByte(0x5D4594, 1307796 + 3 * v39);
			v52 = nox_color_rgb_4344A0(v39, v41, v40);
			nox_draw_setMaterial_4341D0(*(uint32_t*)((uint32_t)dword_5d4594_1308160 + 44), v52);
		} else {
			if (v26 != 1) {
				goto LABEL_27;
			}
			v42 = 1;
			for (k = 3; k < 21; k += 3) {
				v44 = dword_5d4594_1308164;
				LOBYTE(v25) = *(uint8_t*)(k + (uint32_t)dword_5d4594_1308164 + 14);
				LOBYTE(v23) = *(uint8_t*)(k + (uint32_t)dword_5d4594_1308164 + 13);
				LOBYTE(v44) = *(uint8_t*)(k + (uint32_t)dword_5d4594_1308164 + 12);
				nox_draw_setMaterial_4340A0(v42++, v44, v23, v25);
			}
			v45 = ((unsigned short)(*(uint32_t*)(dword_5d4594_1308128 + 32) >> 16)) +
				  32 * *(unsigned short*)(dword_5d4594_1308128 + 32);
			v46 = getMemByte(0x5D4594, 1307798 + 3 * v45);
			LOBYTE(v23) = getMemByte(0x5D4594, 1307797 + 3 * v45);
			LOBYTE(v45) = getMemByte(0x5D4594, 1307796 + 3 * v45);
			v47 = nox_color_rgb_4344A0(v45, v23, v46);
			nox_draw_setMaterial_4341D0(*(uint32_t*)((uint32_t)dword_5d4594_1308164 + 40), v47);
			v48 = ((unsigned short)(*(uint32_t*)(dword_5d4594_1308132 + 32) >> 16)) +
				  32 * *(unsigned short*)(dword_5d4594_1308132 + 32);
			v49 = getMemByte(0x5D4594, 1307798 + 3 * v48);
			LOBYTE(v50) = getMemByte(0x5D4594, 1307797 + 3 * v48);
			LOBYTE(v48) = getMemByte(0x5D4594, 1307796 + 3 * v48);
			v53 = nox_color_rgb_4344A0(v48, v50, v49);
			nox_draw_setMaterial_4341D0(*(uint32_t*)((uint32_t)dword_5d4594_1308164 + 36), v53);
		}
		nox_client_drawImageAt_47D2C0(*v24, v54, v55);
	LABEL_27:
		++v21;
		++v24;
	} while (v21 < 26);
	return 1;
}
// 4A6E13: variable 'v3' is possibly undefined
// 4A6E5A: variable 'v7' is possibly undefined
// 4A7032: variable 'v23' is possibly undefined
// 4A7032: variable 'v25' is possibly undefined
// 4A7158: variable 'v41' is possibly undefined
// 4A7227: variable 'v50' is possibly undefined

//----- (004A7270) --------------------------------------------------------
int sub_4A7270(int a1, int a2, unsigned int a3, int a4) {
	int v3;  // eax
	int2 v5; // [esp+0h] [ebp-8h]

	if (a2 != 5) {
		return 0;
	}
	v5.field_4 = a3 >> 16;
	v5.field_0 = (unsigned short)a3;
	v3 = nox_xxx_pointInRect_4281F0(&v5, (int4*)(a1 + 16));
	if (v3) {
		return 0;
	}
	sub_4A72D0(0xDEADu);
	return 1;
}
// 4A72AA: variable 'v3' is possibly undefined

//----- (004A72D0) --------------------------------------------------------
uint32_t* sub_4A72D0(unsigned short a1) {
	uint32_t* result; // eax

	nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1308088);
	result = (uint32_t*)nox_window_set_hidden(*(int*)&dword_5d4594_1308088, 1);
	if (a1 < 0x20u) {
		result = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308088, *(int*)&dword_5d4594_1307792);
		if (result) {
			result[8] = *getMemU16Ptr(0x5D4594, 1307788) | (a1 << 16);
		}
	}
	return result;
}

//----- (004A7330) --------------------------------------------------------
int sub_4A7330(int a1, int a2, int* a3, unsigned int a4) {
	int result;   // eax
	int v5;       // edi
	int v6;       // eax
	uint32_t* v7; // eax

	if (a2 == 16389) {
		nox_xxx_clientPlaySoundSpecial_452D80(920, 100);
		result = 1;
	} else if (a2 == 16391) {
		v5 = a4 >> 16;
		v6 = nox_xxx_wndGetID_46B0A0(a3);
		switch (v6) {
		case 720:
			dword_5d4594_1307792 = v6;
			sub_4A7530(2u);
			sub_4A7580((unsigned short)a4, v5);
			break;
		case 721:
		case 722:
		case 723:
		case 724:
			dword_5d4594_1307792 = v6;
			sub_4A7530(1u);
			sub_4A7580((unsigned short)a4, v5);
			break;
		case 725:
		case 726:
		case 727:
		case 728:
		case 729:
			dword_5d4594_1307792 = v6;
			sub_4A7530(0);
			sub_4A7580((unsigned short)a4, v5);
			break;
		case 731:
		case 732:
		case 733:
		case 734:
			v7 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308084, v6 - 20);
			if (v7) {
				nox_xxx_wnd_46ABB0((int)v7, ((unsigned int)~v7[1] >> 3) & 1);
			}
			break;
		case 761:
		case 762:
		case 763:
		case 764:
		case 765:
		case 766:
		case 767:
		case 768:
		case 769:
		case 770:
		case 771:
		case 772:
		case 773:
		case 774:
		case 775:
		case 776:
		case 777:
		case 778:
		case 779:
		case 780:
		case 781:
		case 782:
		case 783:
		case 784:
		case 785:
		case 786:
		case 787:
		case 788:
		case 789:
		case 790:
		case 791:
		case 792:
			sub_4A72D0(v6 - 761);
			break;
		case 799:
			if (*getMemU32Ptr(0x5D4594, 1308168) == 1) {
				nox_game_decStateInd_43BDC0();
			}
			nox_game_decStateInd_43BDC0();
			nox_game_decStateInd_43BDC0();
			dword_587000_171388 = 1;
			if (sub_4A75C0()) {
				if (*(uint8_t*)(dword_5d4594_1307784 + 66) == 0) {
					nox_xxx_gameSetMapPath_409D70("war01a.map");
				} else if (*(uint8_t*)(dword_5d4594_1307784 + 66) == 1) {
					nox_xxx_gameSetMapPath_409D70("wiz01a.map");
				} else if (*(uint8_t*)(dword_5d4594_1307784 + 66) == 2) {
					nox_xxx_gameSetMapPath_409D70("con01a.map");
				}
				sub_4A24C0(0);
				sub_4A6890();
				nox_wnd_xxx_1308092->field_13 = 0;
			}
			break;
		default:
			break;
		}
		nox_xxx_clientPlaySoundSpecial_452D80(921, 100);
		result = 1;
	} else {
		result = 0;
	}
	return result;
}

//----- (004A7530) --------------------------------------------------------
uint32_t* sub_4A7530(unsigned short a1) {
	int v1;           // esi
	uint32_t* result; // eax

	v1 = 761;
	*getMemU32Ptr(0x5D4594, 1307788) = a1;
	do {
		result = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1308088, v1);
		if (result) {
			result[8] = a1 | ((unsigned short)(v1 - 761) << 16);
		}
		++v1;
	} while (v1 <= 792);
	return result;
}

//----- (004A7580) --------------------------------------------------------
int sub_4A7580(int a1, int a2) {
	nox_xxx_wndShowModalMB_46A8C0(*(int*)&dword_5d4594_1308088);
	sub_46C690(*(int*)&dword_5d4594_1308088);
	return nox_window_setPos_46A9B0(*(uint32_t**)&dword_5d4594_1308088, a1 - *(uint32_t*)(dword_5d4594_1308088 + 8),
									a2 - *(uint32_t*)(dword_5d4594_1308088 + 12) / 2);
}

//----- (004A7A60) --------------------------------------------------------
int sub_4A7A60(int a1) {
	int result; // eax

	result = a1;
	dword_587000_171388 = a1;
	return result;
}

//----- (004A7A70) --------------------------------------------------------
int sub_4A7A70(int a1) {
	int result; // eax

	result = a1;
	*getMemU32Ptr(0x5D4594, 1308168) = a1;
	return result;
}

//----- (004A7A80) --------------------------------------------------------
int sub_4A7A80(const char* a1) {
	int result;        // eax
	unsigned int v2;   // ecx
	char v3;           // al
	unsigned char* v4; // edi
	const char* v5;    // esi

	if (!a1) {
		return 0;
	}
	v2 = strlen(a1) + 1;
	v3 = v2;
	v2 >>= 2;
	memcpy(getMemAt(0x5D4594, 1308644), a1, 4 * v2);
	v5 = &a1[4 * v2];
	v4 = getMemAt(0x5D4594, 1308644 + 4 * v2);
	LOBYTE(v2) = v3;
	result = 1;
	memcpy(v4, v5, v2 & 3);
	return result;
}

//----- (004A7AC0) --------------------------------------------------------
int sub_4A7AC0(const char* a1) {
	int result;        // eax
	unsigned int v2;   // ecx
	char v3;           // al
	unsigned char* v4; // edi
	const char* v5;    // esi

	if (!a1) {
		return 0;
	}
	v2 = strlen(a1) + 1;
	v3 = v2;
	v2 >>= 2;
	memcpy(getMemAt(0x5D4594, 1308172), a1, 4 * v2);
	v5 = &a1[4 * v2];
	v4 = getMemAt(0x5D4594, 1308172 + 4 * v2);
	LOBYTE(v2) = v3;
	result = 1;
	memcpy(v4, v5, v2 & 3);
	return result;
}

//----- (004A7B00) --------------------------------------------------------
int sub_4A7B00(const char* a1) {
	int result;        // eax
	unsigned int v2;   // ecx
	char v3;           // al
	unsigned char* v4; // edi
	const char* v5;    // esi

	if (!a1) {
		return 0;
	}
	v2 = strlen(a1) + 1;
	v3 = v2;
	v2 >>= 2;
	memcpy(getMemAt(0x5D4594, 1308352), a1, 4 * v2);
	v5 = &a1[4 * v2];
	v4 = getMemAt(0x5D4594, 1308352 + 4 * v2);
	LOBYTE(v2) = v3;
	result = 1;
	memcpy(v4, v5, v2 & 3);
	return result;
}

//----- (004A7B40) --------------------------------------------------------
int sub_4A7B40(char* a1) {
	const char* v2;    // eax
	int v3;            // edi
	unsigned char* v4; // esi

	if (!a1) {
		return 0;
	}
	v2 = *(const char**)getMemAt(0x587000, 171856);
	v3 = 0;
	if (*getMemU32Ptr(0x587000, 171856)) {
		v4 = getMemAt(0x587000, 171856);
		while (nox_strcmpi(v2, a1)) {
			v2 = (const char*)*((uint32_t*)v4 + 2);
			v4 += 8;
			++v3;
			if (!v2) {
				return 1;
			}
		}
		*getMemU32Ptr(0x5D4594, 1308184) = *getMemU32Ptr(0x587000, 171860 + 8 * v3);
	}
	return 1;
}

//----- (004A7BA0) --------------------------------------------------------
int sub_4A7BA0(char* a1) {
	int result; // eax

	result = (int)a1;
	if (a1) {
		*getMemU32Ptr(0x5D4594, 1308740) = atoi(a1);
		result = 1;
	}
	return result;
}

//----- (004A7BC0) --------------------------------------------------------
int sub_4A7BC0(const char* a1) {
	int result;        // eax
	unsigned int v2;   // ecx
	char v3;           // al
	unsigned char* v4; // edi
	const char* v5;    // esi

	if (!a1) {
		return 0;
	}
	v2 = strlen(a1) + 1;
	v3 = v2;
	v2 >>= 2;
	memcpy(getMemAt(0x5D4594, 1308324), a1, 4 * v2);
	v5 = &a1[4 * v2];
	v4 = getMemAt(0x5D4594, 1308324 + 4 * v2);
	LOBYTE(v2) = v3;
	result = 1;
	memcpy(v4, v5, v2 & 3);
	return result;
}

//----- (004A7C00) --------------------------------------------------------
int sub_4A7C00(const char* a1) {
	int result;        // eax
	unsigned int v2;   // ecx
	char v3;           // al
	unsigned char* v4; // edi
	const char* v5;    // esi

	if (!a1) {
		return 0;
	}
	v2 = strlen(a1) + 1;
	v3 = v2;
	v2 >>= 2;
	memcpy(getMemAt(0x5D4594, 1308364), a1, 4 * v2);
	v5 = &a1[4 * v2];
	v4 = getMemAt(0x5D4594, 1308364 + 4 * v2);
	LOBYTE(v2) = v3;
	result = 1;
	memcpy(v4, v5, v2 & 3);
	return result;
}

//----- (004A7C40) --------------------------------------------------------
int sub_4A7C40(char* a1) {
	int result; // eax

	result = (int)a1;
	if (a1) {
		*getMemU32Ptr(0x5D4594, 1308188) = atoi(a1);
		result = 1;
	}
	return result;
}

//----- (004A7C60) --------------------------------------------------------
int sub_4A7C60(char* a1) {
	char* v1;   // eax
	char* v2;   // eax

	if (!a1) {
		return 0;
	}
	*getMemU32Ptr(0x5D4594, 1308736) = 0;
	*getMemU32Ptr(0x5D4594, 1308732) = 0;
	v1 = strtok(a1, ",\t\r\n");
	if (v1) {
		*getMemU32Ptr(0x5D4594, 1308732) = atoi(v1);
	}
	v2 = strtok(0, " \t\r\n");
	if (v2) {
		*getMemU32Ptr(0x5D4594, 1308736) = atoi(v2);
	}
	if (*getMemU32Ptr(0x5D4594, 1308732) && *getMemU32Ptr(0x5D4594, 1308736)) {
		return 1;
	} else {
		return 0;
	}
}

//----- (004A7CE0) --------------------------------------------------------
int sub_4A7CE0(char* a1) {
	int result; // eax

	result = (int)a1;
	if (a1) {
		*getMemU32Ptr(0x5D4594, 1308728) = atoi(a1);
		result = 1;
	}
	return result;
}

//----- (004A7D00) --------------------------------------------------------
int sub_4A7D00(const char* a1) {
	int result;        // eax
	unsigned int v2;   // ecx
	char v3;           // al
	unsigned char* v4; // edi
	const char* v5;    // esi

	if (!a1) {
		return 0;
	}
	result = 0;
	if (strlen(a1) <= 0x80) {
		v2 = strlen(a1) + 1;
		v3 = v2;
		v2 >>= 2;
		memcpy(getMemAt(0x5D4594, 1308192), a1, 4 * v2);
		v5 = &a1[4 * v2];
		v4 = getMemAt(0x5D4594, 1308192 + 4 * v2);
		LOBYTE(v2) = v3;
		result = 1;
		memcpy(v4, v5, v2 & 3);
	}
	return result;
}

//----- (004A7D50) --------------------------------------------------------
int sub_4A7D50(char* a1) {
	int result; // eax

	result = (int)a1;
	if (a1) {
		*getMemU32Ptr(0x5D4594, 1308348) = atoi(a1);
		result = 1;
	}
	return result;
}

//----- (004A7EF0) --------------------------------------------------------
char* sub_4A7EF0() { return (char*)getMemAt(0x5D4594, 1308732); }

//----- (004A84E0) --------------------------------------------------------
int nox_xxx_wndRadioButtonProc_4A84E0(uint32_t* a1, int a2, int a3, int a4) {
	int v4;      // eax
	int v5;      // eax
	int v7;      // ecx
	int v8;      // edx
	int v9;      // eax
	int v10;     // edi
	uint32_t* i; // eax
	int v12;     // eax
	int v13;     // edi
	uint32_t* j; // eax

	switch (a2) {
	case 5:
		return 1;
	case 6:
	case 7:
		v9 = a1[9];
		v10 = a1[99];
		if (v9 & 4) {
			if (v9 & 2) {
				return 1;
			} else {
				return 0;
			}
		}
		nox_window_call_field_94(a1[13], 16391, (int)a1, a3);
		if (v10) {
			for (i = *(uint32_t**)(v10 + 400); i; i = (uint32_t*)i[97]) {
				if (i != a1 && i[10] == a1[10]) {
					i[9] &= 0xFFFFFFFB;
				}
			}
		}
		v12 = a1[9];
		LOBYTE(v12) = v12 | 4;
		a1[9] = v12;
		return 1;
	case 8:
		nox_window_call_field_94(a1[13], 0x4000, (int)a1, a3);
		return 1;
	case 17:
		v4 = a1[11];
		if (!(v4 & 0x100)) {
			return 1;
		}
		v5 = a1[9];
		LOBYTE(v5) = v5 | 2;
		a1[9] = v5;
		nox_window_call_field_94(a1[13], 16389, (int)a1, a3);
		nox_xxx_windowFocus_46B500((int)a1);
		return 1;
	case 18:
		v7 = a1[11];
		if (!(v7 & 0x100)) {
			return 1;
		}
		v8 = a1[13];
		a1[9] &= 0xFFFFFFFD;
		nox_window_call_field_94(v8, 16390, (int)a1, a3);
		return 1;
	case 21:
		switch (a3) {
		case 15:
		case 205:
		case 208:
			if (a4 != 2) {
				return 1;
			}
			nox_xxx_wndRetNULL_46A8A0();
			return 1;
		case 28:
		case 57:
			if (a4 != 2) {
				return 1;
			}
			v13 = a1[99];
			if (a1[9] & 4) {
				return 1;
			}
			nox_window_call_field_94(a1[13], 16391, (int)a1, 0);
			if (v13) {
				for (j = *(uint32_t**)(v13 + 400); j; j = (uint32_t*)j[97]) {
					if (j != a1 && j[10] == a1[10]) {
						j[9] &= 0xFFFFFFFB;
					}
				}
			}
			a1[9] ^= 4u;
			return 1;
		case 200:
		case 203:
			if (a4 == 2) {
				nox_xxx_wndRetNULL_0_46A8B0();
			}
			return 1;
		default:
			return 0;
		}
	default:
		return 0;
	}
}

//----- (004A87E0) --------------------------------------------------------
int nox_xxx_wndRadioButtonSetAllFn_4A87E0(int a1) {
	int result; // eax

	if ((signed char)*(uint8_t*)(a1 + 4) >= 0) {
		result = nox_window_set_all_funcs((uint32_t*)a1, nox_xxx_wndRadioButtonProc_4A84E0,
										  nox_xxx_wndRadioButtonDrawNoImg_4A8820, 0);
	} else {
		result = nox_window_set_all_funcs((uint32_t*)a1, nox_xxx_wndRadioButtonProc_4A84E0,
										  nox_xxx_wndRadioButtonDraw_4A8A20, 0);
	}
	return result;
}

//----- (004A8820) --------------------------------------------------------
int nox_xxx_wndRadioButtonDrawNoImg_4A8820(int a1, int a2) {
	int v2;  // esi
	int v3;  // ebx
	int v4;  // ebp
	int v5;  // edi
	int v6;  // eax
	int v7;  // edi
	int v8;  // eax
	int v9;  // eax
	int v11; // [esp-8h] [ebp-28h]
	int v12; // [esp+10h] [ebp-10h]
	int v13; // [esp+14h] [ebp-Ch]
	int v14; // [esp+18h] [ebp-8h]
	int v15; // [esp+1Ch] [ebp-4h]

	v2 = a2;
	v3 = a1;
	v12 = *(uint32_t*)(a2 + 28);
	a2 = *(uint32_t*)(a2 + 20);
	nox_client_wndGetPosition_46AA60((uint32_t*)a1, &v14, &v15);
	if (*(uint8_t*)(v3 + 4) & 8) {
		if (*(uint8_t*)v2 & 2) {
			v12 = *(uint32_t*)(v2 + 36);
		}
	} else {
		a2 = *(uint32_t*)(v2 + 44);
	}
	v4 = v14 + 4;
	v5 = *(uint32_t*)(v3 + 12) / 2 + v15 - 5;
	v13 = nox_xxx_guiFontHeightMB_43F320(*(uint32_t*)(v2 + 200));
	if (a2 != 0x80000000) {
		nox_client_drawSetColor_434460(a2);
		nox_client_drawRectFilledOpaque_49CE30(v4, v5, 10, 10);
	}
	if (v12 != 0x80000000) {
		nox_client_drawSetColor_434460(v12);
		nox_client_drawBorderLines_49CC70(v4, v5, 10, 10);
	}
	if (*(uint8_t*)v2 & 4 && *(uint32_t*)(v2 + 52) != 0x80000000) {
		nox_client_drawSetColor_434460(*(uint32_t*)(v2 + 52));
		nox_client_drawRectFilledOpaque_49CE30(v4 + 1, v5 + 1, 8, 8);
	}
	if ((*(uint32_t*)(v3 + 4) & 0x2000) == 0x2000) {
		nox_draw_enableTextSmoothing_43F670(1);
	}
	if (**(uint32_t**)(v3 + 32)) {
		v6 = *(uint32_t*)(v3 + 12);
		v14 += *(uint32_t*)(v3 + 8) / 2;
		v7 = v6 / 2 + v15;
		v8 = *(uint32_t*)(v2 + 200);
		v15 = v7;
		nox_xxx_drawGetStringSize_43F840(v8, (unsigned short*)(v2 + 72), &a1, 0, 0);
		if (*(uint32_t*)(v2 + 68) != 0x80000000) {
			nox_xxx_drawSetTextColor_434390(*(uint32_t*)(v2 + 68));
			v11 = *(uint32_t*)(v3 + 8);
			v9 = nox_xxx_guiFontHeightMB_43F320(*(uint32_t*)(v2 + 200));
			nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(v2 + 200), (uint16_t*)(v2 + 72), v14 - a1 / 2, v15 - v9 / 2, v11, 0);
		}
	} else if (*(uint32_t*)(v2 + 68) != 0x80000000) {
		nox_xxx_drawSetTextColor_434390(*(uint32_t*)(v2 + 68));
		nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(v2 + 200), (uint16_t*)(v2 + 72), v4 + 14, v5 - v13 / 2 + 5,
								  *(uint32_t*)(v3 + 8), 0);
	}
	nox_draw_enableTextSmoothing_43F670(0);
	return 1;
}

//----- (004A8A20) --------------------------------------------------------
int nox_xxx_wndRadioButtonDraw_4A8A20(int a1, int a2) {
	int v2;  // esi
	int v3;  // ebp
	int v4;  // ebx
	int v5;  // eax
	int v6;  // kr00_4
	int v7;  // edx
	int v8;  // eax
	int v9;  // kr10_4
	int v10; // eax
	int v12; // [esp-8h] [ebp-20h]
	int v13; // [esp+10h] [ebp-8h]
	int v14; // [esp+14h] [ebp-4h]

	v2 = a2;
	v3 = *(uint32_t*)(a2 + 32);
	v4 = *(uint32_t*)(a2 + 24);
	nox_client_wndGetPosition_46AA60((uint32_t*)a1, &v13, &v14);
	if (*(uint8_t*)(a1 + 4) & 8) {
		if (*(uint8_t*)v2 & 2) {
			v3 = *(uint32_t*)(v2 + 40);
		}
	} else {
		v4 = *(uint32_t*)(v2 + 48);
	}
	if (v4) {
		nox_client_drawImageAt_47D2C0(v4, v13 + *(uint32_t*)(v2 + 60), v14 + *(uint32_t*)(v2 + 64));
	}
	if (*(uint8_t*)v2 & 4) {
		v5 = *(uint32_t*)(v2 + 56);
		if (v5) {
			nox_client_drawImageAt_47D2C0(v5, v13 + *(uint32_t*)(v2 + 60), v14 + *(uint32_t*)(v2 + 64));
		}
	} else if (v3) {
		nox_client_drawImageAt_47D2C0(v3, v13 + *(uint32_t*)(v2 + 60), v14 + *(uint32_t*)(v2 + 64));
	}
	if ((*(uint32_t*)(a1 + 4) & 0x2000) == 0x2000) {
		nox_draw_enableTextSmoothing_43F670(1);
	}
	if (**(uint32_t**)(a1 + 32)) {
		v6 = *(uint32_t*)(a1 + 8);
		v7 = *(uint32_t*)(v2 + 200);
		v14 += *(uint32_t*)(a1 + 12) / 2;
		v13 += v6 / 2;
		nox_xxx_drawGetStringSize_43F840(v7, (unsigned short*)(v2 + 72), &a2, 0, 0);
		if (*(uint32_t*)(v2 + 68) != 0x80000000) {
			nox_xxx_drawSetTextColor_434390(*(uint32_t*)(v2 + 68));
			v12 = *(uint32_t*)(a1 + 8);
			v8 = nox_xxx_guiFontHeightMB_43F320(*(uint32_t*)(v2 + 200));
			nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(v2 + 200), (uint16_t*)(v2 + 72), v13 - a2 / 2, v14 - v8 / 2, v12, 0);
		}
	} else {
		v9 = *(uint32_t*)(a1 + 12) - nox_xxx_guiFontHeightMB_43F320(*(uint32_t*)(v2 + 200));
		v10 = *(uint32_t*)(v2 + 68);
		v14 += v9 / 2;
		v13 += 28;
		if (v10 != 0x80000000) {
			nox_xxx_drawSetTextColor_434390(v10);
			nox_xxx_drawStringWrap_43FAF0(*(uint32_t*)(v2 + 200), (uint16_t*)(v2 + 72), v13, v14, *(uint32_t*)(a1 + 8), 0);
		}
	}
	nox_draw_enableTextSmoothing_43F670(0);
	return 1;
}

//----- (004A93C0) --------------------------------------------------------
int nox_xxx_wndRadioButtonProcPre_4A93C0(int a1, int a2, wchar2_t* a3, int a4) {
	int v3;      // edi
	uint32_t* i; // eax
	int v5;      // eax
	int v7;      // eax
	int v8;      // eax

	if (a2 != 23) {
		if (a2 == 16385) {
			nox_wcsncpy((wchar2_t*)(a1 + 108), a3, 0x3Fu);
			*(uint16_t*)(a1 + 234) = 0;
		} else if (a2 == 16392) {
			v3 = *(uint32_t*)(a1 + 396);
			if (!(*(uint8_t*)(a1 + 36) & 4)) {
				if (a3 == (wchar2_t*)1) {
					nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16391, a1, 0);
				}
				if (v3) {
					for (i = *(uint32_t**)(v3 + 400); i; i = (uint32_t*)i[97]) {
						if (i != (uint32_t*)a1 && i[10] == *(uint32_t*)(a1 + 40)) {
							i[9] &= 0xFFFFFFFB;
						}
					}
				}
				v5 = *(uint32_t*)(a1 + 36);
				LOBYTE(v5) = v5 | 4;
				*(uint32_t*)(a1 + 36) = v5;
				return 0;
			}
		}
		return 0;
	}
	if (!a3) {
		v7 = *(uint32_t*)(a1 + 36);
		LOBYTE(v7) = v7 & 0xFD;
		*(uint32_t*)(a1 + 36) = v7;
	}
	v8 = nox_xxx_wndGetID_46B0A0((int*)a1);
	nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16387, (int)a3, v8);
	return 1;
}

//----- (004A9C80) --------------------------------------------------------
int nox_xxx_compassGenStrings_4A9C80() {
	int v0;            // edi
	unsigned char* v1; // esi
	int v2;            // edi
	unsigned char* v3; // esi
	char v5[64];       // [esp+8h] [ebp-40h]

	*getMemU32Ptr(0x5D4594, 1309664) = 0;
	v0 = 0;
	v1 = getMemAt(0x5D4594, 1309644);
	do {
		nox_sprintf(v5, "Compass%d", ++v0);
		*(uint32_t*)v1 = nox_xxx_gLoadImg_42F970(v5);
		v1 += 4;
	} while ((int)v1 < (int)getMemAt(0x5D4594, 1309660));
	v2 = 0;
	v3 = getMemAt(0x5D4594, 1309516);
	do {
		nox_sprintf(v5, "CompassMainArrow%d", ++v2);
		*(uint32_t*)v3 = nox_xxx_gLoadImg_42F970(v5);
		v3 += 4;
	} while ((int)v3 < (int)getMemAt(0x5D4594, 1309644));
	return 1;
}

//----- (004AA6B0) --------------------------------------------------------
int nox_game_showOptions_4AA6B0() {
	uint32_t* v1;     // esi
	char* v2;         // eax
	int v3;           // eax
	int v4;           // ecx
	unsigned int v5;  // ecx
	uint32_t* v6;     // esi
	char* v7;         // eax
	int v8;           // eax
	int v9;           // ecx
	unsigned int v10; // ecx
	uint32_t* v11;    // esi
	char* v12;        // eax
	int v13;          // eax
	int v14;          // ecx
	unsigned int v15; // ecx
	char* v16;        // [esp-Ch] [ebp-Ch]
	char* v17;        // [esp-Ch] [ebp-Ch]
	char* v18;        // [esp-Ch] [ebp-Ch]
	char* v19;        // [esp-8h] [ebp-8h]
	char* v20;        // [esp-8h] [ebp-8h]
	char* v21;        // [esp-8h] [ebp-8h]

	nox_game_addStateCode_43BDD0(300);
	dword_5d4594_1309720 = nox_new_window_from_file("Options.wnd", sub_4AABE0);
	if (!dword_5d4594_1309720) {
		return 0;
	}
	if (!nox_client_advVideoOpts_New_4CB590(dword_5d4594_1309720)) {
		return 0;
	}
	nox_xxx_wndSetWindowProc_46B300(*(int*)&dword_5d4594_1309720, sub_4A18E0);
	nox_draw_setTabWidth_43FE20(15);
	nox_wnd_xxx_1309740 =
		nox_gui_makeAnimation_43C5B0(*(uint32_t**)&dword_5d4594_1309720, 0, 0, 0, -480, 0, 20, 0, -40);
	if (!nox_wnd_xxx_1309740) {
		return 0;
	}
	nox_wnd_xxx_1309740->field_0 = 300;
	nox_wnd_xxx_1309740->field_12 = sub_4AA9C0;
	nox_wnd_xxx_1309740->fnc_done_out = sub_4AAA10;
	v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 351);
	*(uint32_t*)(v1[100] + 8) = 24;
	*(uint32_t*)(v1[100] + 12) = 20;
	v19 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v16 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v2 = nox_xxx_gLoadImg_42F970("OptionsVolumeSlider");
	sub_4B5700((int)v1, 0, 0, (int)v2, (int)v16, (int)v19);
	nox_window_call_field_94((int)v1, 16395, 0, 0x4000);
	nox_window_call_field_94((int)v1, 16394, *(uint32_t*)((uint32_t)dword_587000_127004 + 4) >> 16, 0);
	dword_5d4594_1309728 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 361);
	v3 = sub_453070();
	v4 = *(uint32_t*)(dword_5d4594_1309728 + 36);
	if (v3 == 1) {
		v5 = v4 | 4;
	} else {
		v5 = v4 & 0xFFFFFFFB;
	}
	*(uint32_t*)(dword_5d4594_1309728 + 36) = v5;
	v6 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 352);
	*(uint32_t*)(v6[100] + 8) = 24;
	*(uint32_t*)(v6[100] + 12) = 20;
	v20 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v17 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v7 = nox_xxx_gLoadImg_42F970("OptionsVolumeSlider");
	sub_4B5700((int)v6, 0, 0, (int)v7, (int)v17, (int)v20);
	nox_window_call_field_94((int)v6, 16395, 0, 0x4000);
	nox_window_call_field_94((int)v6, 16394, *(uint32_t*)((uint32_t)getMemAt(0x5D4594, 830980) + 4) >> 16, 0);
	dword_5d4594_1309732 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 362);
	v8 = sub_44D990();
	v9 = *(uint32_t*)(dword_5d4594_1309732 + 36);
	if (v8 == 1) {
		v10 = v9 | 4;
	} else {
		v10 = v9 & 0xFFFFFFFB;
	}
	*(uint32_t*)(dword_5d4594_1309732 + 36) = v10;
	v11 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 353);
	*(uint32_t*)(v11[100] + 8) = 24;
	*(uint32_t*)(v11[100] + 12) = 20;
	v21 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v18 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v12 = nox_xxx_gLoadImg_42F970("OptionsVolumeSlider");
	sub_4B5700((int)v11, 0, 0, (int)v12, (int)v18, (int)v21);
	nox_window_call_field_94((int)v11, 16395, 0, 0x4000);
	nox_window_call_field_94((int)v11, 16394, *(uint32_t*)((uint32_t)dword_587000_93164 + 4) >> 16, 0);
	dword_5d4594_1309736 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 363);
	v13 = sub_43DC30();
	v14 = *(uint32_t*)(dword_5d4594_1309736 + 36);
	if (v13 == 1) {
		v15 = v14 | 4;
	} else {
		v15 = v14 & 0xFFFFFFFB;
	}
	*(uint32_t*)(dword_5d4594_1309736 + 36) = v15;
	nox_xxx_wndRetNULL_46A8A0();
	sub_4A19F0("OptsBack.wnd:Back");
	sub_4A1A40(0);
	sub_4AAA70();
	return 1;
}
// 4A18E0: using guessed type int  sub_4A18E0(int, int, int, int);

void nox_video_setMenuOptions(nox_window* root);
//----- (004AAA70) --------------------------------------------------------
uint32_t* sub_4AAA70() {
	uint32_t* v0;     // eax
	uint32_t* v1;     // eax
	int v4;           // eax
	uint32_t* result; // eax
	int v6;           // [esp-4h] [ebp-10h]
	int v7;           // [esp-4h] [ebp-10h]

	nox_video_setMenuOptions(dword_5d4594_1309720);
	   // TODO: This is original handling of 8/16bit and windowed/fullscreen.
	   //       However, windowed/fullscreen switch is not present in options.wnd, thus it's not visible in the menu,
	   //       even though the code for it exists here.
	   //       So currently we just pretend that 8/16bit switch actually corresponds to windowed/fullscreen.
	   //       Later we should refactor it to inject those switches.
#if 0
	int v9a = nox_xxx_normalWndBits_587000_172880;
	if (v9a == 8) {
		v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 331);
	} else {
		v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 332);
	}
	nox_window_call_field_94((int)v1, 16392, 1, 0);
	if (nox_video_getFullScreen())
		v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 333);
	else
		v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 334);
	nox_window_call_field_94((int)v1, 16392, 1, 0);
#endif
	if (!nox_video_getFullScreen()) {
		v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 331);
	} else {
		v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 332);
	}
	nox_window_call_field_94((int)v1, 16392, 1, 0);
#if 0
    uint32_t* v2; // eax
    uint32_t* v3; // eax
    if (!nox_video_bagexists_4300D0(0))
    {
        v2 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t * *)& dword_5d4594_1309720, 331);
        nox_xxx_wndClearFlag_46AD80((int)v2, 8);
    }
    if (!nox_video_bagexists_4300D0(1))
    {
        v3 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t * *)& dword_5d4594_1309720, 332);
        nox_xxx_wndClearFlag_46AD80((int)v3, 8);
    }
#endif
	v4 = nox_video_getCutSize_4766D0();
	*getMemU32Ptr(0x587000, 172884) = v4;
	if (v4 > 69) {
		if (v4 <= 79) {
			result = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 312);
			goto LABEL_22;
		}
		if (v4 <= 89) {
			result = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, 313);
			goto LABEL_22;
		}
		v7 = 314;
	} else {
		v7 = 311;
	}
	result = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309720, v7);
LABEL_22:
	if (result) {
		result = (uint32_t*)nox_window_call_field_94((int)result, 16392, 1, 0);
	}
	return result;
}

int nox_gui_menu_proc_ext(int id);
void sub_4AA650();
//----- (004AABE0) --------------------------------------------------------
int sub_4AABE0(int a1, int a2, int* a3, int a4) {
	int v4;                               // eax
	int v5;                               // eax
	int result = 0;                       // eax
	int v7;                               // edi
	int v8;                               // esi
	void (*v9)(uint32_t, int, int, int);  // eax
	void (*v10)(uint32_t, int, int, int); // eax
	void (*v11)(uint32_t, int, int, int); // eax
	void* v12;                            // ecx
	void (*v13)(void*, int, int, int);    // eax
	void (*v14)(uint32_t, int, int, int); // eax
	int v15;                              // ecx
	void (*v16)(int, int, int, int);      // eax
	int v17;                              // eax
	void (*v18)(uint32_t, int, int, int); // eax
	void (*v19)(uint32_t, int, int, int); // eax

	switch (a2) {
	case 16389:
		nox_xxx_clientPlaySoundSpecial_452D80(920, 100);
		result = 1;
		break;
	case 16391:
		v4 = nox_xxx_wndGetID_46B0A0(a3);
		if (v4 >= 380) {
			return nox_gui_menu_proc_ext(v4);
		}
		if (v4 > 341) {
			if (v4 > 363) {
				if (v4 == 2099) {
					sub_4AAA70();
				}
			} else if (v4 == 363) {
				if (sub_43DC30() == 1) {
					sub_43DC00();
				} else {
					sub_43DC10();
					sub_486320(*(uint32_t**)&dword_587000_93164, *(uint32_t*)((uint32_t)dword_587000_93164 + 4) >> 16);
				}
			} else {
				v5 = v4 - 361;
				if (v5) {
					if (v5 == 1) {
						if (sub_44D990() == 1) {
							sub_44D960();
						} else {
							sub_44D970();
						}
					}
				} else if (sub_453070() == 1) {
					sub_453050();
				} else {
					nox_xxx____setargv_9_453060();
				}
			}
		} else if (v4 == 341) {
			sub_4AA9C0();
			nox_wnd_xxx_1309740->fnc_done_out = sub_4AB0C0;
			nox_wnd_xxx_1309740->field_13 = sub_4CB880;
		} else {
			switch (v4) {
			case 311:
				nox_video_setCutSize_4766A0(65);
				break;
			case 312:
				nox_video_setCutSize_4766A0(75);
				break;
			case 313:
				nox_video_setCutSize_4766A0(85);
				break;
			case 314:
				nox_video_setCutSize_4766A0(100);
				break;
			case 332:
				nox_xxx_normalWndBits_587000_172880 = 16;
			case 333:
				nox_video_setFullScreen(1);
				break;
			case 331:
				nox_xxx_normalWndBits_587000_172880 = 8;
			case 334:
				nox_video_setFullScreen(0);
				break;
			default:
				break;
			}
		}
		nox_xxx_clientPlaySoundSpecial_452D80(921, 100);
		result = 1;
		break;
	case 16393:
		v7 = nox_xxx_wndGetID_46B0A0(a3);
		v8 = nox_client_getWin1064916_46C720();
		switch (v7) {
		case 351:
			sub_486320(*(uint32_t**)&dword_587000_127004, a4);
			if (v8 && *(int**)(v8 + 396) == a3) {
				return 0;
			}
			if (a4) {
				if (sub_453070()) {
					nox_xxx_clientPlaySoundSpecial_452D80(768, 100);
					return 0;
				}
				v15 = dword_5d4594_1309728;
				v16 = *(void (**)(int, int, int, int))(dword_5d4594_1309728 + 372);
				if (!v16) {
					nox_xxx_clientPlaySoundSpecial_452D80(768, 100);
					return 0;
				}
				v16(v15, 21, 28, 2);
				nox_xxx_clientPlaySoundSpecial_452D80(768, 100);
				return 0;
			}
			if (sub_453070() != 1) {
				return 0;
			}
			v14 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309728 + 372);
			if (!v14) {
				return 0;
			}
			v14(dword_5d4594_1309728, 21, 28, 2);
			result = 0;
			break;
		case 352:
			sub_486320(getMemAt(0x5D4594, 830980), a4);
			if (a4) {
				if (sub_44D990()) {
					sub_4AA650();
					return 0;
				}
				v12 = *(void**)&dword_5d4594_1309732;
				v13 = *(void (**)(void*, int, int, int))(dword_5d4594_1309732 + 372);
				if (!v13) {
					sub_4AA650();
					return 0;
				}
				v13(v12, 21, 28, 2);
				sub_4AA650();
				return 0;
			}
			if (sub_44D990() != 1) {
				return 0;
			}
			v11 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309732 + 372);
			if (!v11) {
				return 0;
			}
			v11(dword_5d4594_1309732, 21, 28, 2);
			result = 0;
			break;
		case 353:
			sub_486320(*(uint32_t**)&dword_587000_93164, a4);
			if (a4) {
				if (sub_43DC30()) {
					return 0;
				}
				v10 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309736 + 372);
				if (!v10) {
					return 0;
				}
				v10(dword_5d4594_1309736, 21, 28, 2);
				result = 0;
			} else {
				if (sub_43DC30() != 1) {
					return 0;
				}
				v9 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309736 + 372);
				if (!v9) {
					return 0;
				}
				v9(dword_5d4594_1309736, 21, 28, 2);
				result = 0;
			}
			break;
		case 316: // gamma slider
			nox_video_setGammaSlider(a4);
			break;
		case 318: // input sensitivity slider
			nox_input_setSensitivity(powf(10.0, a4 / 50.0f - 1.0));
			break;
		default:
			return 0;
		}
		break;
	case 16396:
		v17 = nox_xxx_wndGetID_46B0A0(a3) - 351;
		if (v17) {
			if (v17 != 1) {
				return 0;
			}
			sub_486320(getMemAt(0x5D4594, 830980), a4);
			if (a4) {
				if (!sub_44D990()) {
					v12 = *(void**)&dword_5d4594_1309732;
					v13 = *(void (**)(void*, int, int, int))(dword_5d4594_1309732 + 372);
					if (v13) {
						v13(v12, 21, 28, 2);
					}
				}
				sub_4AA650();
				return 0;
			} else {
				if (sub_44D990() != 1) {
					return 0;
				}
				v18 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309732 + 372);
				if (!v18) {
					return 0;
				}
				v18(dword_5d4594_1309732, 21, 28, 2);
				result = 0;
			}
		} else {
			sub_486320(*(uint32_t**)&dword_587000_127004, a4);
			if (a4) {
				if (!sub_453070()) {
					v15 = dword_5d4594_1309728;
					v16 = *(void (**)(int, int, int, int))(dword_5d4594_1309728 + 372);
					if (v16) {
						v16(v15, 21, 28, 2);
					}
				}
				nox_xxx_clientPlaySoundSpecial_452D80(768, 100);
				return 0;
			}
			if (sub_453070() != 1) {
				return 0;
			}
			v19 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309728 + 372);
			if (!v19) {
				return 0;
			}
			v19(dword_5d4594_1309728, 21, 28, 2);
			result = 0;
		}
		break;
	default:
		return 0;
	}
	return result;
}
// 4AAEC8: variable 'v12' is possibly undefined

//----- (004AB0C0) --------------------------------------------------------
int sub_4AB0C0() {
	int (*v0)(void); // esi

	v0 = nox_wnd_xxx_1309740->field_13;
	nox_gui_freeAnimation_43C570(nox_wnd_xxx_1309740);
	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1309720);
	v0();
	return 1;
}

//----- (004AB260) --------------------------------------------------------
int sub_4AB260() {
	*getMemU32Ptr(0x5D4594, 1309752) = nox_xxx_gLoadImg_42F970("DisconnectIcon");
	dword_5d4594_1309756 = nox_window_new(0, 136, nox_win_width - 50, nox_win_height / 2 + 3, 50, 50, 0);
	nox_xxx_wndSetIcon_46AE60(*(int*)&dword_5d4594_1309756, *getMemIntPtr(0x5D4594, 1309752));
	nox_window_set_all_funcs(*(uint32_t**)&dword_5d4594_1309756, 0, sub_4AB420, 0);
	dword_5d4594_1309748 = nox_new_window_from_file("discon.wnd", sub_4AB390);
	nox_xxx_wndSetWindowProc_46B300(*(int*)&dword_5d4594_1309748, sub_4AB340);
	sub_46B120(*(uint32_t**)&dword_5d4594_1309748, 0);
	nox_window_setPos_46A9B0(*(uint32_t**)&dword_5d4594_1309748,
							 nox_win_width / 2 - *(uint32_t*)(dword_5d4594_1309748 + 24) / 2,
							 nox_win_height / 2 - *(uint32_t*)(dword_5d4594_1309748 + 28) / 2);
	return 1;
}

//----- (004AB340) --------------------------------------------------------
int sub_4AB340(int a1, int a2, int a3, int a4) {
	if (a2 != 21) {
		return 0;
	}
	if (a3 == 1) {
		return 1;
	}
	if (a3 == 57) {
		nox_point mpos = nox_client_getMousePos_4309F0();
		nox_window_call_field_93(a1, 5, mpos.x | (mpos.y << 16), 0);
	}
	return 0;
}

//----- (004AB390) --------------------------------------------------------
int sub_4AB390(int a1, int a2, int* a3, int a4) {
	int v3;     // eax
	int result; // eax

	if (a2 == 23) {
		return 1;
	}
	if (a2 != 16391) {
		return 0;
	}
	v3 = nox_xxx_wndGetID_46B0A0(a3) - 576;
	if (!v3) {
		sub_43CF40();
		return 0;
	}
	if (v3 != 1) {
		return 0;
	}
	sub_446380();
	if (dword_5d4594_2650652 && sub_41E2F0() == 9) {
		sub_41F4B0();
		sub_41EC30();
		sub_446490(0);
		nox_xxx____setargv_4_44B000();
		sub_4AB4D0(0);
		result = 0;
	} else {
		sub_43B750();
		sub_4AB4D0(0);
		result = 0;
	}
	return result;
}

//----- (004AB420) --------------------------------------------------------
int sub_4AB420(int* a1) {
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

//----- (004AB470) --------------------------------------------------------
int sub_4AB470() {
	int result; // eax

	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1309748);
	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1309756);
	result = 0;
	dword_5d4594_1309756 = 0;
	dword_5d4594_1309748 = 0;
	return result;
}

//----- (004AB4A0) --------------------------------------------------------
int sub_4AB4A0(int a1) {
	int result; // eax

	if (a1) {
		result = nox_window_set_hidden(*(int*)&dword_5d4594_1309756, 0);
	} else {
		result = nox_window_set_hidden(*(int*)&dword_5d4594_1309756, 1);
	}
	return result;
}

//----- (004AB4D0) --------------------------------------------------------
int sub_4AB4D0(int a1) {
	int result; // eax

	if (a1) {
		nox_video_stopAllFades_44E040();
		nox_window_set_hidden(*(int*)&dword_5d4594_1309748, 0);
		nox_xxx_wndShowModalMB_46A8C0(*(int*)&dword_5d4594_1309748);
		sub_46C690(*(int*)&dword_5d4594_1309748);
		nox_xxx_windowFocus_46B500(*(int*)&dword_5d4594_1309748);
		result = nox_xxx_wnd_46ABB0(*(int*)&dword_5d4594_1309748, 1);
	} else {
		nox_window_set_hidden(*(int*)&dword_5d4594_1309748, 1);
		nox_xxx_wnd_46C6E0(*(int*)&dword_5d4594_1309748);
		nox_xxx_windowFocus_46B500(0);
		result = nox_xxx_wnd_46ABB0(*(int*)&dword_5d4594_1309748, 0);
	}
	return result;
}

//----- (004ABDA0) --------------------------------------------------------
int sub_4ABDA0(int a1, short a2, short a3, uint32_t* a4) {
	int v4;          // edi
	int result;      // eax
	int v6;          // esi
	unsigned int v7; // edx
	int v8;          // eax
	float v9;        // [esp-14h] [ebp-148h]
	char v10;        // [esp+Fh] [ebp-125h]
	float v11;       // [esp+10h] [ebp-124h]
	int v12;         // [esp+14h] [ebp-120h]
	int v13[2];      // [esp+18h] [ebp-11Ch]
	float v14;       // [esp+20h] [ebp-114h]
	int v15;         // [esp+24h] [ebp-110h]
	int v16;         // [esp+28h] [ebp-10Ch]
	float v17;       // [esp+2Ch] [ebp-108h]
	int v18;         // [esp+30h] [ebp-104h]
	char v19[256];   // [esp+34h] [ebp-100h]

	nox_xxx_fileReadWrite_426AC0_file3_fread(&v15, 4u);
	*a4 += 4;
	nox_xxx_fileReadWrite_426AC0_file3_fread(&v16, 4u);
	*a4 += 4;
	if (a3 < 40 || a2 < 4) {
		nox_xxx_fileReadWrite_426AC0_file3_fread(v13, 8u);
		*a4 += 8;
	} else {
		nox_xxx_fileReadWrite_426AC0_file3_fread(&v11, 4u);
		*a4 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread(&v17, 4u);
		v9 = v11;
		*a4 += 4;
		v13[0] = nox_float2int(v9);
		v13[1] = nox_float2int(v17);
	}
	if (a3 >= 10) {
		nox_xxx_fileReadWrite_426AC0_file3_fread(&v12, 1u);
		++*a4;
		nox_xxx_fileReadWrite_426AC0_file3_fread(v19, (unsigned char)v12);
		*a4 += (unsigned char)v12;
	}
	if (a3 >= 20) {
		nox_xxx_fileReadWrite_426AC0_file3_fread(&v10, 1u);
		++*a4;
	}
	if (a3 >= 30) {
		nox_xxx_fileReadWrite_426AC0_file3_fread(&v12, 1u);
		++*a4;
	}
	if (a3 >= 40) {
		nox_xxx_cryptSeekCur_40E0A0(4);
		*a4 += 4;
		if (a2 >= 2) {
			if (a2 < 5) {
				nox_xxx_fileReadWrite_426AC0_file3_fread(&v14, 4u);
				*a4 += 4;
				v11 = v14;
			} else {
				nox_xxx_fileReadWrite_426AC0_file3_fread(&v11, 2u);
				*a4 += 2;
			}
			v4 = (unsigned short)(4 * LOWORD(v11));
			nox_xxx_cryptSeekCur_40E0A0(v4);
			*a4 += v4;
		}
		if (a2 >= 3) {
			nox_xxx_fileReadWrite_426AC0_file3_fread(&v18, 4u);
			*a4 += 4;
		}
	}
	result = nox_xxx_spriteLoadAdd_45A360_drawable(a1, v13[0], v13[1]);
	v6 = result;
	if (result) {
		v7 = *(uint32_t*)(result + 120) & 0xEEBF7E9D;
		*(uint32_t*)(result + 120) = v7;
		*(uint32_t*)(result + 120) = v16 | v7;
		v8 = *(uint32_t*)(result + 280);
		LOBYTE(v8) = v8 & 0xA1;
		*(uint32_t*)(v6 + 280) = v8;
		*(uint32_t*)(v6 + 280) = v18 | v8;
		*(uint32_t*)(v6 + 128) = v15;
		*(uint8_t*)(v6 + 28) = v10;
		if (!nox_common_gameFlags_check_40A5C0(1) && *(uint8_t*)(v6 + 28) &&
			(!(*(uint32_t*)(v6 + 112) & 0x10000000) || !nox_common_gameFlags_check_40A5C0(128))) {
			nox_xxx_createAtImpl_4191D0(*(uint8_t*)(v6 + 28), v6 + 24, 0, *(uint32_t*)(v6 + 128), 0);
		}
		result = v6;
	}
	return result;
}

//----- (004AC020) --------------------------------------------------------
int nox_xxx_spriteLoadFromMap_4AC020(int thingInd, short a2, uint32_t* a3) {
	uint32_t* v3;      // esi
	short v4;          // di
	short v5;          // ax
	bool v6;           // cc
	int v7;            // eax
	bool v10;          // zf
	int v11;           // ecx
	int v12;           // ebx
	char v13;          // dl
	int v14;           // ebx
	int v15;           // ecx
	int v16;           // eax
	int v17;           // eax
	unsigned char v18; // cl
	int v19;           // [esp-30h] [ebp-60h]
	float v20;         // [esp-24h] [ebp-54h]
	int v21;           // [esp-14h] [ebp-44h]
	int v22;           // [esp-4h] [ebp-34h]
	int v23;           // [esp+Ch] [ebp-24h]
	int v24;           // [esp+10h] [ebp-20h]
	int v25;           // [esp+14h] [ebp-1Ch]
	int v26;           // [esp+18h] [ebp-18h]
	int v27;           // [esp+1Ch] [ebp-14h]
	int v28;           // [esp+20h] [ebp-10h]
	int v29;           // [esp+24h] [ebp-Ch]
	float v30;         // [esp+28h] [ebp-8h]
	float v31;         // [esp+2Ch] [ebp-4h]

	v3 = a3;
	v4 = a2;
	v5 = 1;
	v23 = 1;
	if (a2 < 40) {
		return sub_4ABDA0(thingInd, v5, v4, v3);
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread(&v23, 2u);
	v5 = v23;
	v6 = (short)v23 < 61;
	*v3 += 2;
	if (v6) {
		return sub_4ABDA0(thingInd, v5, v4, v3);
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread(&v25, 4u);
	*v3 += 4;
	nox_xxx_fileReadWrite_426AC0_file3_fread(&v29, 4u);
	*v3 += 4;
	nox_xxx_fileReadWrite_426AC0_file3_fread(&v30, 4u);
	*v3 += 4;
	nox_xxx_fileReadWrite_426AC0_file3_fread(&v31, 4u);
	v20 = v31;
	*v3 += 4;
	v22 = nox_float2int(v20);
	v7 = nox_float2int(v30);
	int v9 = nox_xxx_spriteLoadAdd_45A360_drawable(thingInd, v7, v22);
	if (!v9) {
		return 0;
	}

	*(uint32_t*)(v9 + 128) = v25;
	nox_xxx_fileReadWrite_426AC0_file3_fread(&a2, 1u);
	v10 = (uint8_t)a2 == 0;
	++*v3;
	if (!v10) {
		nox_xxx_fileReadWrite_426AC0_file3_fread(&v26, 4u);
		v11 = v26;
		*v3 += 4;
		*(uint32_t*)(v9 + 120) = v11 | *(uint32_t*)(v9 + 120) & 0xEEBF7E9D;
		nox_xxx_fileReadWrite_426AC0_file3_fread(&a3, 1u);
		v12 = (unsigned char)a3;
		v21 = (unsigned char)a3;
		++*v3;
		nox_xxx_cryptSeekCur_40E0A0(v21);
		*v3 += v12;
		nox_xxx_fileReadWrite_426AC0_file3_fread(&thingInd, 1u);
		v13 = thingInd;
		++*v3;
		*(uint8_t*)(v9 + 28) = v13;
		nox_xxx_fileReadWrite_426AC0_file3_fread(&a3, 1u);
		++*v3;
		nox_xxx_fileReadWrite_426AC0_file3_fread(&v27, 2u);
		v14 = 4 * (unsigned short)v27;
		v19 = 4 * (unsigned short)v27;
		*v3 += 2;
		nox_xxx_cryptSeekCur_40E0A0(v19);
		*v3 += v14;
		nox_xxx_fileReadWrite_426AC0_file3_fread(&v28, 4u);
		v15 = v28;
		*v3 += 4;
		v16 = *(uint32_t*)(v9 + 280);
		LOBYTE(v16) = v16 & 0xA1;
		*(uint32_t*)(v9 + 280) = v15 | v16;
		if ((short)v23 >= 63) {
			nox_xxx_cryptSeekCur_40E0A0(2);
			*v3 += 2;
			nox_xxx_fileReadWrite_426AC0_file3_fread(&v24, 4u);
			*v3 += 4;
			nox_xxx_cryptSeekCur_40E0A0(v24);
			*v3 += v24;
			nox_xxx_cryptSeekCur_40E0A0(4);
			*v3 += 4;
		}
		if ((short)v23 >= 64) {
			nox_xxx_cryptSeekCur_40E0A0(4);
			*v3 += 4;
		}
	}
	if (!nox_common_gameFlags_check_40A5C0(1)) {
		v17 = *getMemU32Ptr(0x5D4594, 1309788);
		if (!*getMemU32Ptr(0x5D4594, 1309788)) {
			v17 = nox_xxx_getTTByNameSpriteMB_44CFC0("FlagMarker");
			*getMemU32Ptr(0x5D4594, 1309788) = v17;
		}
		v18 = *(uint8_t*)(v9 + 28);
		if (v18) {
			if (*(uint32_t*)(v9 + 108) != v17) {
				nox_xxx_createAtImpl_4191D0(v18, v9 + 24, 0, *(uint32_t*)(v9 + 128), 0);
			}
		}
	}
	return v9;
}

//----- (004AC610) --------------------------------------------------------
int nox_client_mapSpecialRWObjectData_4AC610() {
	unsigned short v1; // ax
	unsigned short v2; // si
	int v3;            // eax
	int v4;            // [esp+4h] [ebp-Ch]
	int v5;            // [esp+8h] [ebp-8h]
	int v6;            // [esp+Ch] [ebp-4h]

	v6 = 1;
	nox_xxx_fileReadWrite_426AC0_file3_fread(&v6, 2u);
	if ((short)v6 > 1) {
		return 0;
	}
	if (!nox_crypt_IsReadOnly()) {
		return 0;
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread(&v4, 2u);
	if ((uint16_t)v4) {
		while (1) {
			nox_xxx_fileCryptReadCrcMB_426C20(&v5, 4u);
			v1 = nox_xxx_objectTOCgetTT_42C2B0(v4);
			v2 = v1;
			if (!v1) {
				break;
			}
			if (sub_44D090(v1)) {
				v3 = nox_xxx_clientLoadSomeObject_4AC6E0(v2);
				v5 -= v3;
			}
			if (v5 > 0) {
				nox_xxx_cryptSeekCur_40E0A0(v5);
			}
			nox_xxx_fileReadWrite_426AC0_file3_fread(&v4, 2u);
			if (!(uint16_t)v4) {
				return 1;
			}
		}
		return 0;
	}
	return 1;
}

//----- (004AC6E0) --------------------------------------------------------
int nox_xxx_clientLoadSomeObject_4AC6E0(unsigned short a1) {
	if (!*getMemU32Ptr(0x5D4594, 1309792)) {
		*getMemU32Ptr(0x5D4594, 1309792) = nox_xxx_getTTByNameSpriteMB_44CFC0("ColorLight");
		*getMemU32Ptr(0x5D4594, 1309796) = nox_xxx_getTTByNameSpriteMB_44CFC0("ColorLightMovable");
		*getMemU32Ptr(0x5D4594, 1309800) = nox_xxx_getTTByNameSpriteMB_44CFC0("TeamBase");
		*getMemU32Ptr(0x5D4594, 1309804) = nox_xxx_getTTByNameSpriteMB_44CFC0("PressurePlate");
	}
	if (a1 == *getMemU32Ptr(0x5D4594, 1309792) || a1 == *getMemU32Ptr(0x5D4594, 1309796)) {
		return nox_xxx_colorLightClientLoad_4AC980(a1);
	}
	if (a1 == *getMemU32Ptr(0x5D4594, 1309800)) {
		return nox_xxx_cliLoadTeamBase_4ACE00(a1);
	}
	if (a1 == *getMemU32Ptr(0x5D4594, 1309804)) {
		return sub_4ACEF0(a1);
	}
	if (sub_44D060(a1)) {
		return sub_4AD040(a1);
	}
	if (sub_44D040(a1)) {
		return sub_4AC7B0(a1);
	}
	return 0;
}

//----- (004AC7B0) --------------------------------------------------------
int sub_4AC7B0(int a1) {
	int v1;     // eax
	int v2;     // nox_drawable* ?
	int result; // eax
	int v4;     // eax
	char v5;    // al
	float v6;   // edi
	int v7;     // ecx
	double v9;  // st7
	double v10; // st7
	int v11;    // [esp+4h] [ebp-10h]
	int v12;    // [esp+8h] [ebp-Ch]
	float v13;  // [esp+Ch] [ebp-8h]
	int v14;    // [esp+10h] [ebp-4h]

	nox_xxx_fileReadWrite_426AC0_file3_fread(&v12, 2u);
	v11 = 2;
	v1 = nox_xxx_spriteLoadFromMap_4AC020(a1, v12, &v11);
	v2 = v1;
	if (v1) {
		v4 = *(uint32_t*)(v1 + 112);
		if ((v4 & 0x80u) == 0) {
			if (v4 & 0x200) {
				nox_xxx_fileReadWrite_426AC0_file3_fread(&a1, 4u);
				nox_xxx_fileReadWrite_426AC0_file3_fread(&v14, 4u);
				v9 = (double)a1;
				v11 += 8;
				v13 = v9;
				*(float*)(v2 + 56) = v9;
				v10 = (double)v14;
				*(float*)(v2 + 60) = v10;
				if (v13 > 60.0) {
					*(uint32_t*)(v2 + 56) = 1114636288;
				}
				if (v10 > 60.0) {
					*(uint32_t*)(v2 + 60) = 1114636288;
				}
				nox_shape_box_calc((nox_shape*)(v2 + 44));
			}
		} else {
			nox_xxx_fileReadWrite_426AC0_file3_fread(&v13, 4u);
			v11 += 4;
			nox_xxx_fileReadWrite_426AC0_file3_fread(&v14, 4u);
			v5 = v14;
			v11 += 4;
			*(uint8_t*)(v2 + 433) = v14;
			*(uint8_t*)(v2 + 432) = 0;
			if (v5) {
				*(uint8_t*)(v2 + 432) = 1;
			}
			v6 = v13;
			a1 = LODWORD(v13);
			if ((short)v12 >= 41) {
				nox_xxx_fileReadWrite_426AC0_file3_fread(&a1, 4u);
				v6 = *(float*)&a1;
				v11 += 4;
			}
			*(uint8_t*)(v2 + 299) = LOBYTE(v13);
			v7 = *(uint32_t*)(v2 + 12) + *getMemIntPtr(0x587000, 196184 + 8 * LODWORD(v6)) / 2;
			sub_410390(v2, v7 / 23,
					   (*(uint32_t*)(v2 + 16) + *getMemIntPtr(0x587000, 196188 + 8 * LODWORD(v6)) / 2) / 23);
		}
		*(uint32_t*)(v2 + 288) = 0;
		nox_xxx_spriteSetActiveMB_45A990_drawable(v2);
		result = v11;
	} else {
		nox_xxx_spriteLoadError_4356E0();
		result = 0;
	}
	return result;
}

//----- (004AC980) --------------------------------------------------------
int nox_xxx_colorLightClientLoad_4AC980(int a1) {
	int v1;             // eax
	int v2;             // esi
	unsigned short* v4; // edi
	uint8_t* v5;        // eax
	uint8_t* v6;        // ecx
	int v7;             // [esp+Ch] [ebp-8h]
	int v8;             // [esp+10h] [ebp-4h]

	nox_xxx_fileReadWrite_426AC0_file3_fread(&v8, 2u);
	v7 = 2;
	v1 = nox_xxx_spriteLoadFromMap_4AC020(a1, v8, &v7);
	v2 = v1;
	if (!v1) {
		nox_xxx_spriteLoadError_4356E0();
		return 0;
	}
	if ((short)v8 >= 2) {
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v1 + 136), 4u);
		v7 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 140), 4u);
		v7 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 144), 4u);
		v7 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 148), 4u);
		v7 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 152), 0xCu);
		v4 = (unsigned short*)(v2 + 164);
		v7 += 12;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 164), 2u);
		v7 += 2;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 166), 2u);
		v7 += 2;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 168), 4u);
		v7 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 176), 2u);
		v7 += 2;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 178), 0x30u);
		v7 += 48;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 226), 0x10u);
		v7 += 16;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 242), 0x10u);
		v7 += 16;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 258), 2u);
		v7 += 2;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 260), 2u);
		v7 += 2;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 262), 2u);
		v7 += 2;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 264), 4u);
		v7 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 270), 2u);
		v7 += 2;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 272), 2u);
		v7 += 2;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 274), 1u);
		++v7;
		if ((short)v8 > 40) {
			if ((short)v8 >= 42) {
				nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 172), 4u);
				v7 += 4;
			} else {
				nox_xxx_fileReadWrite_426AC0_file3_fread(&a1, 1u);
				++v7;
				*(uint32_t*)(v2 + 172) = (unsigned char)a1;
			}
		}
	} else {
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v1 + 136), 4u);
		v7 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 140), 4u);
		v7 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 144), 4u);
		v7 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 148), 4u);
		v7 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 152), 0xCu);
		v4 = (unsigned short*)(v2 + 164);
		v7 += 12;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 164), 2u);
		v7 += 2;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 166), 2u);
		v7 += 2;
		nox_xxx_fileReadWrite_426AC0_file3_fread((uint8_t*)(v2 + 168), 4u);
		v7 += 4;
		*(uint16_t*)(v2 + 176) = 0;
		*(uint16_t*)(v2 + 258) = 0;
		*(uint16_t*)(v2 + 260) = 0;
		*(uint16_t*)(v2 + 262) = 0;
		*(uint32_t*)(v2 + 264) = 0;
		*(uint16_t*)(v2 + 270) = 0;
		*(uint8_t*)(v2 + 274) = -128;
		if (nox_crypt_IsReadOnly() != 1) {
			goto LABEL_24;
		}
		if (*(float*)(v2 + 140) <= 63.0 &&
			(double)*(int*)(v2 + 148) * *getMemDoublePtr(0x581450, 9752) <= *getMemDoublePtr(0x581450, 9744)) {
			goto LABEL_13;
		}
		sub_484CE0(v2 + 136, 63.0);
	}
	if (nox_crypt_IsReadOnly() == 1) {
		goto LABEL_13;
	} else {
		goto LABEL_24;
	}
LABEL_13:
	*(uint8_t*)(v2 + 432) = 0;
	*(uint8_t*)(v2 + 433) = 0;
	*(uint8_t*)(v2 + 434) = 0;
	v5 = (uint8_t*)(v2 + 242);
	v6 = (uint8_t*)(v2 + 179);
	do {
		if (*(v6 - 1) || *v6 || v6[1]) {
			++*(uint8_t*)(v2 + 432);
		}
		if (*(v5 - 16)) {
			++*(uint8_t*)(v2 + 433);
		}
		if (*v5) {
			++*(uint8_t*)(v2 + 434);
		}
		v6 += 3;
		++v5;
	} while ((int)&v5[-242 - v2] < 16);
	a1 = *v4;
	*(uint16_t*)(v2 + 268) = (long long)((double)a1 * *getMemDoublePtr(0x581450, 9752) * *getMemDoublePtr(0x581450, 9736));
LABEL_24:
	*(uint32_t*)(v2 + 288) = 0;
	nox_xxx_spriteSetActiveMB_45A990_drawable(v2);
	return v7;
}

//----- (004ACE00) --------------------------------------------------------
int nox_xxx_cliLoadTeamBase_4ACE00(int a1) {
	int v1;       // eax
	int v2;       // esi
	int result;   // eax
	int i;        // edi
	int v5;       // eax
	int v6;       // [esp+Ch] [ebp-10Ch]
	int v7;       // [esp+10h] [ebp-108h]
	int v8;       // [esp+14h] [ebp-104h]
	char v9[256]; // [esp+18h] [ebp-100h]

	nox_xxx_fileReadWrite_426AC0_file3_fread(&v8, 2u);
	v6 = 2;
	v1 = nox_xxx_spriteLoadFromMap_4AC020(a1, v8, &v6);
	v2 = v1;
	if (v1) {
		*(uint32_t*)(v1 + 288) = 0;
		nox_xxx_spriteSetActiveMB_45A990_drawable(v1);
		for (i = 0; i < 16; i += 4) {
			nox_xxx_fileReadWrite_426AC0_file3_fread(&v7, 1u);
			++v6;
			nox_xxx_fileReadWrite_426AC0_file3_fread(v9, (unsigned char)v7);
			v6 += (unsigned char)v7;
			v9[(unsigned char)v7] = 0;
			v5 = nox_xxx_modifGetIdByName_413290(v9);
			*(uint32_t*)(i + v2 + 432) = nox_xxx_modifGetDescById_413330(v5);
			*(uint16_t*)(v2 + 448) = -1;
			*(uint16_t*)(v2 + 450) = -1;
		}
		result = v6;
	} else {
		nox_xxx_spriteLoadError_4356E0();
		result = 0;
	}
	return result;
}
// 4ACE00: using guessed type char var_100[256];

//----- (004ACEF0) --------------------------------------------------------
int sub_4ACEF0(int a1) {
	int v1;      // eax
	int v2;      // edi
	int result;  // eax
	uint8_t* v4; // esi
	int v5;      // [esp+Ch] [ebp-10h]
	int v6;      // [esp+10h] [ebp-Ch]
	int v7;      // [esp+14h] [ebp-8h]
	uint8_t* v8; // [esp+18h] [ebp-4h]

	nox_xxx_fileReadWrite_426AC0_file3_fread(&v7, 2u);
	v5 = 2;
	v1 = nox_xxx_spriteLoadFromMap_4AC020(a1, v7, &v5);
	v2 = v1;
	if (v1) {
		v4 = (uint8_t*)(v1 + 432);
		*(uint32_t*)(v1 + 288) = 0;
		nox_xxx_spriteSetActiveMB_45A990_drawable(v1);
		nox_xxx_fileReadWrite_426AC0_file3_fread(&v6, 4u);
		*(float*)(v2 + 56) = (double)v6;
		v5 += 4;
		nox_xxx_fileReadWrite_426AC0_file3_fread(&v6, 4u);
		*(float*)(v2 + 60) = (double)v6;
		v5 += 4;
		nox_shape_box_calc((nox_shape*)(v2 + 44));
		v4[4] = 10;
		*v4 = 90;
		v4[1] = 90;
		v4[2] = 90;
		v4[3] = 10;
		v4[5] = 10;
		v8 = v4 + 5;
		if ((short)v7 >= 41) {
			nox_xxx_fileReadWrite_426AC0_file3_fread(v4, 1u);
			++v5;
			nox_xxx_fileReadWrite_426AC0_file3_fread(v4 + 1, 1u);
			++v5;
			nox_xxx_fileReadWrite_426AC0_file3_fread(v4 + 2, 1u);
			++v5;
			nox_xxx_fileReadWrite_426AC0_file3_fread(v4 + 3, 1u);
			++v5;
			nox_xxx_fileReadWrite_426AC0_file3_fread(v4 + 4, 1u);
			++v5;
			nox_xxx_fileReadWrite_426AC0_file3_fread(v8, 1u);
			++v5;
		}
		result = v5;
	} else {
		nox_xxx_spriteLoadError_4356E0();
		result = 0;
	}
	return result;
}

//----- (004AD040) --------------------------------------------------------
int sub_4AD040(int a1) {
	int v1;     // esi
	int result; // eax
	int v3;     // [esp+4h] [ebp-Ch]
	int v4;     // [esp+8h] [ebp-8h]
	int v5;     // [esp+Ch] [ebp-4h]

	nox_xxx_fileReadWrite_426AC0_file3_fread(&v4, 2u);
	v3 = 2;
	v1 = nox_xxx_spriteLoadFromMap_4AC020(a1, v4, &v3);
	if (v1) {
		if ((short)v4 >= 61) {
			LOBYTE(a1) = 0;
			nox_xxx_fileReadWrite_426AC0_file3_fread(&v5, 4u);
			v3 += 4;
			nox_xxx_fileReadWrite_426AC0_file3_fread(&a1, 1u);
			++v3;
			if ((uint8_t)a1 == 1) {
				sub_459DD0(v1, 1);
			}
		}
		*(uint32_t*)(v1 + 288) = 0;
		nox_xxx_spriteSetActiveMB_45A990_drawable(v1);
		result = v3;
	} else {
		nox_xxx_spriteLoadError_4356E0();
		result = 0;
	}
	return result;
}

//----- (004AD570) --------------------------------------------------------
int sub_4AD570() {
	uint32_t* v1; // esi

	nox_point mpos = nox_client_getMousePos_4309F0();
	if (0) {
		v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309812, 10317);
		if (!wndIsShown_nox_xxx_wndIsShown_46ACC0((int)v1) && !nox_xxx_wndPointInWnd_46AAB0(v1, mpos.x, mpos.y)) {
			nox_xxx_wndClearCaptureMain_46ADE0((int)v1);
			nox_window_set_hidden((int)v1, 1);
		}
	}
	return 1;
}

//----- (004AD5D0) --------------------------------------------------------
int nox_xxx_windowServerOptionsGeneralProc_4AD5D0(int a1, int a2, int* a3, int a4) {
	uint32_t* v4;  // ebx
	int v5;        // eax
	int v6;        // esi
	int v7;        // eax
	int v8;        // eax
	int v9;        // esi
	uint32_t* v10; // eax
	int v12;       // esi
	uint32_t* v13; // esi

	if (a2 == 16391) {
		v12 = nox_xxx_wndGetID_46B0A0(a3);
		nox_xxx_clientPlaySoundSpecial_452D80(766, 100);
		switch (v12) {
		case 10301:
			if (nox_server_doPlayersAutoRespawn_40A5F0()) {
				nox_xxx_ruleSetNoRespawn_40A5E0(0);
			} else {
				nox_xxx_ruleSetNoRespawn_40A5E0(1);
			}
			return 0;
		case 10302:
			nox_server_sendMotd_108752 ^= 1u;
			return 0;
		case 10304:
			if (sub_4D0D70()) {
				sub_4D0D90(0);
			} else {
				sub_4D0D90(1);
			}
			return 0;
		case 10305:
			sub_409EF0(2);
			return 0;
		case 10306:
			sub_409EF0(0x2000);
			if (sub_409F40(0x2000)) {
				return 0;
			}
			sub_4D7EA0();
			return 0;
		case 10316:
			v13 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309812, 10317);
			nox_window_set_hidden((int)v13, 0);
			nox_xxx_wndShowModalMB_46A8C0((int)v13);
			nox_xxx_wndSetCaptureMain_46ADC0((int)v13);
			return 0;
		case 10319:
			sub_4BDFD0();
			return 0;
		default:
			return 0;
		}
	} else if (a2 == 16393) {
		nox_xxx_rateUpdate_40A6D0(4 - a4);
		return 0;
	} else if (a2 != 16400 || nox_xxx_wndGetID_46B0A0(a3) != 10317 ||
			   (v4 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309812, 10316),
				v5 = nox_window_call_field_94((int)a3, 16404, 0, 0), v6 = v5, v5 < 0) ||
			   v5 >= *(short*)(a3[8] + 44)) {
		return 0;
	} else {
		v7 = nox_window_call_field_94((int)a3, 16406, a4, 0);
		nox_window_call_field_94((int)v4, 16385, v7, -1);
		nox_server_connectionType_3596 = v6 + 1;
		v8 = sub_40A710(v6 + 1);
		nox_xxx_rateUpdate_40A6D0(v8);
		v9 = 4 - nox_xxx_rateGet_40A6C0();
		v10 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309812, 10312);
		nox_window_call_field_94((int)v10, 16394, v9, 0);
		nox_window_set_hidden((int)a3, 1);
		nox_xxx_wndClearCaptureMain_46ADE0((int)a3);
		return 0;
	}
}

//----- (004AD820) --------------------------------------------------------
int sub_4AD820() {
	nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1309812);
	dword_5d4594_1309812 = 0;
	return sub_4BE610();
}

//----- (004AD9B0) --------------------------------------------------------
void nox_common_writecfgfile(char* a1);
int sub_4AD9B0(int a1) {
	if (wndIsShown_nox_xxx_wndIsShown_46ACC0(*(int*)&dword_5d4594_1309820)) {
		return 0;
	}
	sub_413A00(0);
	sub_44D8F0();
	if (a1) {
		nox_window_set_hidden(*(int*)&dword_5d4594_1309820, 1);
		nox_window_set_hidden(*(int*)&dword_5d4594_1309824, 1);
	} else {
		nox_common_writecfgfile("nox.cfg");
		nox_window_set_hidden(*(int*)&dword_5d4594_1309820, 1);
		nox_window_set_hidden(*(int*)&dword_5d4594_1309824, 1);
		sub_445C40();
	}
	return 1;
}

//----- (004ADA40) --------------------------------------------------------
int sub_4ADA40() {
	int v0;       // eax
	uint32_t* v1; // eax
	int v3;       // [esp-4h] [ebp-4h]

	sub_413A00(1);
	nox_xxx_wndShowModalMB_46A8C0(*(int*)&dword_5d4594_1309820);
	nox_xxx_wndShowModalMB_46A8C0(*(int*)&dword_5d4594_1309824);
	nox_client_advVideoOptsLoad_4CB330();
	v0 = nox_video_getCutSize_4766D0();
	if (v0 > 69) {
		if (v0 <= 79) {
			v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, 312);
			goto LABEL_9;
		}
		if (v0 <= 89) {
			v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, 313);
			goto LABEL_9;
		}
		v3 = 314;
	} else {
		v3 = 311;
	}
	v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, v3);
LABEL_9:
	if (v1) {
		nox_window_call_field_94((int)v1, 16392, 0, 0);
	}
	return nox_draw_setTabWidth_43FE20(15);
}

//----- (004ADAD0) --------------------------------------------------------
int nox_game_initOptionsInGame_4ADAD0() {
	uint32_t* v1;     // esi
	char* v2;         // eax
	int v3;           // eax
	int v4;           // ecx
	unsigned int v5;  // ecx
	uint32_t* v6;     // esi
	char* v7;         // eax
	int v8;           // eax
	int v9;           // ecx
	unsigned int v10; // ecx
	uint32_t* v11;    // esi
	char* v12;        // eax
	int v13;          // eax
	int v14;          // ecx
	unsigned int v15; // ecx
	uint32_t* v16;    // eax
	uint32_t* v17;    // eax
	int i;            // esi
	uint32_t* v19;    // eax
	int v20;          // eax
	int* v21;         // esi
	uint32_t* v22;    // eax
	int v23;          // kr00_4
	uint32_t* v24;    // eax
	char* v25;        // [esp-Ch] [ebp-18h]
	char* v26;        // [esp-Ch] [ebp-18h]
	char* v27;        // [esp-Ch] [ebp-18h]
	char* v28;        // [esp-8h] [ebp-14h]
	char* v29;        // [esp-8h] [ebp-14h]
	char* v30;        // [esp-8h] [ebp-14h]
	int v31;          // [esp-8h] [ebp-14h]

	dword_5d4594_1309820 = nox_new_window_from_file("Options.wnd", nox_xxx_windowOptionsProc_4ADF30);
	if (!dword_5d4594_1309820) {
		return 0;
	}
	if (!nox_client_advVideoOpts_New_4CB590(dword_5d4594_1309820)) {
		return 0;
	}
	v1 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, 351);
	*(uint32_t*)(v1[100] + 8) = 24;
	*(uint32_t*)(v1[100] + 12) = 20;
	v28 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v25 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v2 = nox_xxx_gLoadImg_42F970("OptionsVolumeSlider");
	sub_4B5700((int)v1, 0, 0, (int)v2, (int)v25, (int)v28);
	nox_window_call_field_94((int)v1, 16395, 0, 0x4000);
	nox_window_call_field_94((int)v1, 16394, *(uint32_t*)((uint32_t)dword_587000_127004 + 4) >> 16, 0);
	dword_5d4594_1309828 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, 361);
	v3 = sub_453070();
	v4 = *(uint32_t*)(dword_5d4594_1309828 + 36);
	if (v3 == 1) {
		v5 = v4 | 4;
	} else {
		v5 = v4 & 0xFFFFFFFB;
	}
	*(uint32_t*)(dword_5d4594_1309828 + 36) = v5;
	v6 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, 352);
	*(uint32_t*)(v6[100] + 8) = 24;
	*(uint32_t*)(v6[100] + 12) = 20;
	v29 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v26 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v7 = nox_xxx_gLoadImg_42F970("OptionsVolumeSlider");
	sub_4B5700((int)v6, 0, 0, (int)v7, (int)v26, (int)v29);
	nox_window_call_field_94((int)v6, 16395, 0, 0x4000);
	nox_window_call_field_94((int)v6, 16394, *(uint32_t*)((uint32_t)getMemAt(0x5D4594, 830980) + 4) >> 16, 0);
	dword_5d4594_1309836 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, 362);
	v8 = sub_44D990();
	v9 = *(uint32_t*)(dword_5d4594_1309836 + 36);
	if (v8 == 1) {
		v10 = v9 | 4;
	} else {
		v10 = v9 & 0xFFFFFFFB;
	}
	*(uint32_t*)(dword_5d4594_1309836 + 36) = v10;
	v11 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, 353);
	*(uint32_t*)(v11[100] + 8) = 24;
	*(uint32_t*)(v11[100] + 12) = 20;
	v30 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v27 = nox_xxx_gLoadImg_42F970("OptionsVolumeSliderLit");
	v12 = nox_xxx_gLoadImg_42F970("OptionsVolumeSlider");
	sub_4B5700((int)v11, 0, 0, (int)v12, (int)v27, (int)v30);
	nox_window_call_field_94((int)v11, 16395, 0, 0x4000);
	nox_window_call_field_94((int)v11, 16394, *(uint32_t*)((uint32_t)dword_587000_93164 + 4) >> 16, 0);
	dword_5d4594_1309832 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, 363);
	v13 = sub_43DC30();
	v14 = *(uint32_t*)(dword_5d4594_1309832 + 36);
	if (v13 == 1) {
		v15 = v14 | 4;
	} else {
		v15 = v14 & 0xFFFFFFFB;
	}
	*(uint32_t*)(dword_5d4594_1309832 + 36) = v15;

	nox_video_setMenuOptions(dword_5d4594_1309820);
	if (nox_video_getFullScreen()) {
		v17 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, 333);
	} else {
		v17 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, 334);
	}
	nox_window_call_field_94((int)v17, 16392, 1, 0);
	for (i = 320; i <= 332; ++i) {
		v19 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, i);
		if (v19 && !(v19[9] & 4)) {
			nox_xxx_wndClearFlag_46AD80((int)v19, 8);
		}
	}
	dword_5d4594_1309824 = nox_window_new(0, 32, 0, 0, 1, 1, 0);
	v20 = *getMemU32Ptr(0x587000, 174072);
	if (*getMemIntPtr(0x587000, 174072) != -1) {
		v21 = getMemIntPtr(0x587000, 174080);
		do {
			v22 = nox_window_new(*(int*)&dword_5d4594_1309824, 0, v20, *(v21 - 1), *v21, v21[1], 0);
			nox_xxx_wndSetDrawFn_46B340((int)v22, sub_4ADEF0);
			v20 = v21[2];
			v21 += 4;
		} while (v20 != -1);
	}
	v23 = nox_win_width - *(uint32_t*)(dword_5d4594_1309820 + 8);
	nox_window_setPos_46A9B0(*(uint32_t**)&dword_5d4594_1309820, v23 / 2, 0);
	nox_window_setPos_46A9B0(*(uint32_t**)&dword_5d4594_1309824, v23 / 2, 0);
	v24 = nox_xxx_wndGetChildByID_46B0C0(*(uint32_t**)&dword_5d4594_1309820, 371);
	nox_window_set_hidden((int)v24, 0);
	nox_window_set_hidden(*(int*)&dword_5d4594_1309820, 1);
	nox_window_set_hidden(*(int*)&dword_5d4594_1309824, 1);
	return 1;
}

//----- (004ADEF0) --------------------------------------------------------
int sub_4ADEF0(uint32_t* a1, int a2) {
	int xLeft; // [esp+4h] [ebp-8h]
	int yTop;  // [esp+8h] [ebp-4h]

	nox_client_wndGetPosition_46AA60(a1, &xLeft, &yTop);
	nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, a1[2], a1[3]);
	return 1;
}

//----- (004ADF30) --------------------------------------------------------
int nox_xxx_windowOptionsProc_4ADF30(int a1, int a2, int* a3, int a4) {
	int result = 0;                       // eax
	int v5;                               // edi
	int v6;                               // esi
	void (*v7)(uint32_t, int, int, int);  // eax
	void (*v8)(uint32_t, int, int, int);  // eax
	void (*v9)(uint32_t, int, int, int);  // eax
	void* v10;                            // ecx
	void (*v11)(void*, int, int, int);    // eax
	void (*v12)(uint32_t, int, int, int); // eax
	int v13;                              // ecx
	void (*v14)(int, int, int, int);      // eax
	int v15;                              // eax
	void (*v16)(uint32_t, int, int, int); // eax
	void (*v17)(uint32_t, int, int, int); // eax

	switch (a2) {
	case 16389:
		nox_xxx_clientPlaySoundSpecial_452D80(920, 100);
		return 1;
	case 16391:
		switch (nox_xxx_wndGetID_46B0A0(a3)) {
		case 311:
			nox_draw_setCutSize_476700(65, 0);
			break;
		case 312:
			nox_draw_setCutSize_476700(75, 0);
			break;
		case 313:
			nox_draw_setCutSize_476700(85, 0);
			break;
		case 314:
			nox_draw_setCutSize_476700(100, 0);
			break;
		case 341:
			sub_4AD9B0(0);
			sub_445C40();
			sub_4C4260();
			break;
		case 361:
			if (sub_453070() == 1) {
				sub_453050();
			} else {
				nox_xxx____setargv_9_453060();
			}
			break;
		case 362:
			if (sub_44D990() == 1) {
				sub_44D960();
			} else {
				sub_44D970();
			}
			break;
		case 363:
			if (sub_43DC30() == 1) {
				sub_43DC00();
			} else {
				sub_43DC10();
				sub_486320(*(uint32_t**)&dword_587000_93164, *(uint32_t*)((uint32_t)dword_587000_93164 + 4) >> 16);
			}
			break;
		case 371:
			sub_4AD9B0(0);
			break;
		default:
			break;
		}
		nox_xxx_clientPlaySoundSpecial_452D80(921, 100);
		return 1;
	case 16393:
		v5 = nox_xxx_wndGetID_46B0A0(a3);
		v6 = nox_client_getWin1064916_46C720();
		switch (v5) {
		case 351:
			sub_486320(*(uint32_t**)&dword_587000_127004, a4);
			if (v6 && *(int**)(v6 + 396) == a3) {
				return 0;
			}
			if (a4) {
				if (sub_453070()) {
					nox_xxx_clientPlaySoundSpecial_452D80(768, 100);
					return 0;
				}
				v13 = dword_5d4594_1309828;
				v14 = *(void (**)(int, int, int, int))(dword_5d4594_1309828 + 372);
				if (!v14) {
					nox_xxx_clientPlaySoundSpecial_452D80(768, 100);
					return 0;
				}
				v14(v13, 21, 28, 2);
				nox_xxx_clientPlaySoundSpecial_452D80(768, 100);
				return 0;
			}
			if (sub_453070() != 1) {
				return 0;
			}
			v12 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309828 + 372);
			if (!v12) {
				return 0;
			}
			v12(dword_5d4594_1309828, 21, 28, 2);
			result = 0;
			break;
		case 352:
			sub_486320(getMemAt(0x5D4594, 830980), a4);
			if (a4) {
				if (sub_44D990()) {
					sub_4AA650();
					return 0;
				}
				v10 = *(void**)&dword_5d4594_1309836;
				v11 = *(void (**)(void*, int, int, int))(dword_5d4594_1309836 + 372);
				if (!v11) {
					sub_4AA650();
					return 0;
				}
				v11(v10, 21, 28, 2);
				sub_4AA650();
				return 0;
			}
			if (sub_44D990() != 1) {
				return 0;
			}
			v9 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309836 + 372);
			if (!v9) {
				return 0;
			}
			v9(dword_5d4594_1309836, 21, 28, 2);
			return 0;
		case 353:
			sub_486320(*(uint32_t**)&dword_587000_93164, a4);
			if (a4) {
				if (sub_43DC30()) {
					return 0;
				}
				v8 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309832 + 372);
				if (!v8) {
					return 0;
				}
				v8(dword_5d4594_1309832, 21, 28, 2);
				result = 0;
			} else {
				if (sub_43DC30() != 1) {
					return 0;
				}
				v7 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309832 + 372);
				if (!v7) {
					return 0;
				}
				v7(dword_5d4594_1309832, 21, 28, 2);
				result = 0;
			}
			break;
		case 316: // gamma slider
			nox_video_setGammaSlider(a4);
			break;
		case 318: // input sensitivity slider
			nox_input_setSensitivity(powf(10.0, a4 / 50.0f - 1.0));
			break;
		default:
			return 0;
		}
		break;
	case 16396:
		v15 = nox_xxx_wndGetID_46B0A0(a3) - 351;
		if (v15) {
			if (v15 != 1) {
				return 0;
			}
			sub_486320(getMemAt(0x5D4594, 830980), a4);
			if (a4) {
				if (!sub_44D990()) {
					v10 = *(void**)&dword_5d4594_1309836;
					v11 = *(void (**)(void*, int, int, int))(dword_5d4594_1309836 + 372);
					if (v11) {
						v11(v10, 21, 28, 2);
					}
				}
				sub_4AA650();
				return 0;
			} else {
				if (sub_44D990() != 1) {
					return 0;
				}
				v16 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309836 + 372);
				if (!v16) {
					return 0;
				}
				v16(dword_5d4594_1309836, 21, 28, 2);
				return 0;
			}
		} else {
			sub_486320(*(uint32_t**)&dword_587000_127004, a4);
			if (a4) {
				if (!sub_453070()) {
					v13 = dword_5d4594_1309828;
					v14 = *(void (**)(int, int, int, int))(dword_5d4594_1309828 + 372);
					if (v14) {
						v14(v13, 21, 28, 2);
					}
				}
				nox_xxx_clientPlaySoundSpecial_452D80(768, 100);
				return 0;
			}
			if (sub_453070() != 1) {
				return 0;
			}
			v17 = *(void (**)(uint32_t, int, int, int))(dword_5d4594_1309828 + 372);
			if (!v17) {
				return 0;
			}
			v17(dword_5d4594_1309828, 21, 28, 2);
			return 0;
		}
		break;
	default:
		return 0;
	}
	return result;
}
// 4AE184: variable 'v10' is possibly undefined

//----- (004AE3B0) --------------------------------------------------------
int sub_4AE3B0() {
	int result; // eax

	result = nox_xxx_windowDestroyMB_46C4E0(*(uint32_t**)&dword_5d4594_1309820);
	dword_5d4594_1309820 = 0;
	return result;
}

//----- (004AE3D0) --------------------------------------------------------
int sub_4AE3D0() {
	int result; // eax

	if (wndIsShown_nox_xxx_wndIsShown_46ACC0(*(int*)&dword_5d4594_1309820)) {
		result = sub_4C4280();
	} else {
		result = 1;
	}
	return result;
}

//----- (004AEE30) --------------------------------------------------------
long long sub_4AEE30() {
	int v0;            // edi
	unsigned char* v1; // esi
	long long result;  // rax

	v0 = 0;
	v1 = getMemAt(0x5D4594, 1309840);
	do {
		result =
			(long long)(sin((double)(v0 + 192) * *getMemDoublePtr(0x581450, 9768) * *getMemDoublePtr(0x581450, 9760)) *
						*(double*)&qword_581450_9552);
		*(uint32_t*)v1 = result;
		v1 += 4;
		++v0;
	} while ((int)v1 < (int)getMemAt(0x5D4594, 1311120));
	return result;
}

//----- (004B4860) --------------------------------------------------------
int sub_4B4860(int a1, int a2, int a3, int a4) {
	int v4;     // edi
	int v5;     // eax
	int v7;     // eax
	int v8;     // eax
	int v9;     // ebx
	int v10;    // edi
	int v11;    // eax
	int v12;    // eax
	int v13;    // eax
	int v14;    // eax

	v4 = *(uint32_t*)(a1 + 32);
	switch (a2) {
	case 5:
		return 1;
	case 6:
	case 7:
		v9 = a3;
		v10 = (unsigned short)a3;
		nox_client_wndGetPosition_46AA60((uint32_t*)a1, &a4, &a3);
		if (v10 > *(int*)(a1 + 8) + a4 - 10) {
			v10 = *(uint32_t*)(a1 + 8) + a4 - 10;
		}
		nox_window_setPos_46A9B0(*(uint32_t**)(a1 + 400), v10 - a4 - 5, 0);
		nox_window_call_field_94(a1, 0x4000, 0, v9);
		return 1;
	case 8:
		v8 = *(uint32_t*)(a1 + 44);
		if (!(v8 & 0x100)) {
			return 1;
		}
		nox_window_call_field_94(*(uint32_t*)(a1 + 52), 0x4000, a1, 0);
		return 1;
	case 17:
		v5 = *(uint32_t*)(a1 + 44);
		if (!(v5 & 0x100)) {
			return 1;
		}
		nox_xxx_wndSetRectColor2MB_46AFE0(a1, *(uint32_t*)(a1 + 72));
		nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16389, a1, 0);
		nox_xxx_windowFocus_46B500(a1);
		return 1;
	case 18:
		v7 = *(uint32_t*)(a1 + 44);
		if (v7 & 0x100) {
			nox_xxx_wndSetRectColor2MB_46AFE0(a1, *(uint32_t*)(a1 + 64));
			nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16390, a1, 0);
		}
		return 1;
	case 21:
		switch (a3) {
		case 15:
		case 208:
			if (a4 != 2) {
				return 1;
			}
			nox_xxx_wndRetNULL_46A8A0();
			return 1;
		case 200:
			if (a4 == 2) {
				nox_xxx_wndRetNULL_0_46A8B0();
			}
			return 1;
		case 203:
			if (a4 != 2) {
				return 1;
			}
			v13 = *(uint32_t*)(v4 + 12);
			if (v13 >= *(int*)(v4 + 4) - 1) {
				return 1;
			}
			v14 = v13 + 2;
			*(uint32_t*)(v4 + 12) = v14;
			nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16393, a1, v14);
			nox_window_setPos_46A9B0(
				*(uint32_t**)(a1 + 400),
				(long long)((double)(int)(*(uint32_t*)(v4 + 12) - *(uint32_t*)v4) * *(float*)(v4 + 8)), 0);
			return 1;
		case 205:
			if (a4 != 2) {
				return 1;
			}
			v11 = *(uint32_t*)(v4 + 12);
			if (v11 <= *(int*)v4 + 1) {
				return 1;
			}
			v12 = v11 - 2;
			*(uint32_t*)(v4 + 12) = v12;
			nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16393, a1, v12);
			nox_window_setPos_46A9B0(
				*(uint32_t**)(a1 + 400),
				(long long)((double)(int)(*(uint32_t*)(v4 + 12) - *(uint32_t*)v4) * *(float*)(v4 + 8)), 0);
			return 1;
		default:
			return 0;
		}
	default:
		return 0;
	}
}

//----- (004B4BA0) --------------------------------------------------------
int nox_xxx_wndScrollBoxDraw_4B4BA0(int a1, int a2, unsigned int a3, int a4) {
	int v4;         // edi
	int v5;         // eax
	int v7;         // eax
	int v8;         // eax
	int v9;         // ebx
	signed int v10; // edi
	int v11;        // eax
	int v12;        // eax
	int v13;        // eax
	int v14;        // eax

	v4 = *(uint32_t*)(a1 + 32);
	switch (a2) {
	case 5:
		return 1;
	case 6:
	case 7:
		v9 = a3;
		v10 = a3 >> 16;
		nox_client_wndGetPosition_46AA60((uint32_t*)a1, &a3, &a4);
		if (v10 > *(int*)(a1 + 12) + a4 - 10) {
			v10 = *(uint32_t*)(a1 + 12) + a4 - 10;
		}
		nox_window_setPos_46A9B0(*(uint32_t**)(a1 + 400), 0, v10 - a4 - 5);
		nox_window_call_field_94(a1, 0x4000, 0, v9);
		return 1;
	case 8:
		v8 = *(uint32_t*)(a1 + 44);
		if (!(v8 & 0x100)) {
			return 1;
		}
		nox_window_call_field_94(*(uint32_t*)(a1 + 52), 0x4000, a1, 0);
		return 1;
	case 17:
		v5 = *(uint32_t*)(a1 + 44);
		if (!(v5 & 0x100)) {
			return 1;
		}
		nox_xxx_wndSetRectColor2MB_46AFE0(a1, *(uint32_t*)(a1 + 72));
		nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16389, a1, 0);
		nox_xxx_windowFocus_46B500(a1);
		return 1;
	case 18:
		v7 = *(uint32_t*)(a1 + 44);
		if (v7 & 0x100) {
			nox_xxx_wndSetRectColor2MB_46AFE0(a1, *(uint32_t*)(a1 + 64));
			nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16390, a1, 0);
		}
		return 1;
	case 21:
		switch (a3) {
		case 0xFu:
		case 0xCDu:
			if (a4 != 2) {
				return 1;
			}
			nox_xxx_wndRetNULL_46A8A0();
			return 1;
		case 0xC8u:
			if (a4 != 2) {
				return 1;
			}
			v11 = *(uint32_t*)(v4 + 12);
			if (v11 >= *(int*)(v4 + 4) - 1) {
				return 1;
			}
			v12 = v11 + 2;
			*(uint32_t*)(v4 + 12) = v12;
			nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16393, a1, v12);
			nox_window_setPos_46A9B0(
				*(uint32_t**)(a1 + 400), 0,
				(long long)((double)(int)(*(uint32_t*)(v4 + 4) - *(uint32_t*)(v4 + 12)) * *(float*)(v4 + 8)));
			return 1;
		case 0xCBu:
			if (a4 == 2) {
				nox_xxx_wndRetNULL_0_46A8B0();
			}
			return 1;
		case 0xD0u:
			if (a4 != 2) {
				return 1;
			}
			v13 = *(uint32_t*)(v4 + 12);
			if (v13 <= *(int*)v4 + 1) {
				return 1;
			}
			v14 = v13 - 2;
			*(uint32_t*)(v4 + 12) = v14;
			nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16393, a1, v14);
			nox_window_setPos_46A9B0(
				*(uint32_t**)(a1 + 400), 0,
				(long long)((double)(int)(*(uint32_t*)(v4 + 4) - *(uint32_t*)(v4 + 12)) * *(float*)(v4 + 8)));
			return 1;
		default:
			return 0;
		}
	default:
		return 0;
	}
}

//----- (004B4EE0) --------------------------------------------------------
nox_window* nox_gui_newSlider_4B4EE0(int a1, int a2, int a3, int a4, int a5, int a6, uint32_t* a7, float* a8) {
	int v8;        // eax
	int v9;        // ebp
	uint32_t* v10; // edi
	int v11;       // eax
	double v12;    // st7
	float* v13;    // edx

	v8 = a7[2];
	if (v8 & 0x10) {
		v9 = a6;
		v10 = nox_window_new(a1, a2 | 0x100, a3, a4, a5, a6, sub_4B5010);
		sub_4B51A0((int)v10);
	} else {
		if (!(v8 & 8)) {
			return 0;
		}
		v9 = a6;
		v10 = nox_window_new(a1, a2 | 0x100, a3, a4, a5, a6, nox_xxx_wndScrollBoxProc_4B5320);
		nox_xxx_wndScrollBoxSetAllFn_4B5500((int)v10);
	}
	if (v10) {
		if (!a7[4]) {
			a7[4] = v10;
		}
		nox_gui_windowCopyDrawData_46AF80((int)v10, a7);
		nox_xxx_wndScrollBoxButtonCreate_4B5640((int)v10, a2 | 0x100, (int)a7);
		v11 = *(uint32_t*)a8;
		if (*((uint32_t*)a8 + 1) == *(uint32_t*)a8) {
			*((uint32_t*)a8 + 1) = v11 + 1;
		}
		if (a7[2] & 0x10) {
			v12 = (double)(a5 - 10) / (double)(int)(*((uint32_t*)a8 + 1) - v11);
		} else {
			v12 = (double)(v9 - 10) / (double)(int)(*((uint32_t*)a8 + 1) - v11);
		}
		a8[2] = v12;
		v13 = (float*)calloc(1, 0x10u);
		*v13 = *a8;
		v13[1] = a8[1];
		v13[2] = a8[2];
		v13[3] = a8[3];
		v10[8] = v13;
	}
	return v10;
}

//----- (004B5010) --------------------------------------------------------
int sub_4B5010(int a1, unsigned int a2, int a3, int a4) {
	int v4;        // edi
	int v5;        // esi
	int v6;        // esi
	int v7;        // eax
	int v8;        // eax
	int v10;       // ebx
	int v11;       // ecx
	int v12;       // eax
	long long v13; // rax
	int v14;       // ecx
	int v15;       // eax
	int v16;       // edx
	double v17;    // st7
	int* v18;      // [esp-4h] [ebp-Ch]

	v4 = a1;
	v5 = *(uint32_t*)(a1 + 32);
	if (a2 > 0x4000) {
		if (a2 == 16394) {
			if (a3 >= *(int*)v5 && a3 <= *(int*)(v5 + 4)) {
				v17 = (double)a3;
				*(uint32_t*)(v5 + 12) = a3;
				nox_window_setPos_46A9B0(*(uint32_t**)(v4 + 400), (long long)(v17 * *(float*)(v5 + 8)), 0);
			}
		} else if (a2 == 16395) {
			v14 = a3;
			v15 = a4;
			*(uint32_t*)v5 = a3;
			*(uint32_t*)(v5 + 4) = v15;
			v16 = *(uint32_t*)(v4 + 8) - 10;
			*(uint32_t*)(v5 + 12) = v14;
			a3 = v15 - v14;
			*(float*)(v5 + 8) = (double)v16 / (double)(v15 - v14);
			nox_window_setPos_46A9B0(*(uint32_t**)(v4 + 400), 0, 0);
			return 0;
		}
		return 0;
	}
	if (a2 == 0x4000) {
		v10 = (unsigned short)a4;
		nox_client_wndGetPosition_46AA60((uint32_t*)a1, &a3, &a1);
		if (v10 < a3) {
			nox_window_setPos_46A9B0(*(uint32_t**)(v4 + 400), 0, 0);
			v11 = *(uint32_t*)v5;
			*(uint32_t*)(v5 + 12) = v11;
			nox_window_call_field_94(*(uint32_t*)(v4 + 52), 16393, v4, *(uint32_t*)(v5 + 12));
			return 0;
		}
		v12 = *(uint32_t*)(v4 + 8);
		if (v10 < v12 + a3) {
			a4 = v10 - a3;
			v13 = (long long)((double)(v10 - a3) / *(float*)(v5 + 8));
			v11 = *(uint32_t*)(v5 + 4);
			*(uint32_t*)(v5 + 12) = v13;
			if ((int)v13 > v11) {
				*(uint32_t*)(v5 + 12) = v11;
				nox_window_call_field_94(*(uint32_t*)(v4 + 52), 16393, v4, *(uint32_t*)(v5 + 12));
				return 0;
			}
		} else {
			nox_window_setPos_46A9B0(*(uint32_t**)(v4 + 400), v12 - *(uint32_t*)(*(uint32_t*)(v4 + 400) + 8), 0);
			*(uint32_t*)(v5 + 12) = *(uint32_t*)(v5 + 4);
		}
		nox_window_call_field_94(*(uint32_t*)(v4 + 52), 16393, v4, *(uint32_t*)(v5 + 12));
		return 0;
	}
	if (a2 == 2) {
		free(*(void**)(a1 + 32));
		return 0;
	}
	if (a2 != 23) {
		return 0;
	}
	v6 = a3;
	v7 = *(uint32_t*)(a1 + 36);
	if (a3) {
		LOBYTE(v7) = v7 | 2;
	} else {
		LOBYTE(v7) = v7 & 0xFD;
	}
	v18 = (int*)a1;
	*(uint32_t*)(a1 + 36) = v7;
	v8 = nox_xxx_wndGetID_46B0A0(v18);
	nox_window_call_field_94(*(uint32_t*)(v4 + 52), 16387, v6, v8);
	return 1;
}

//----- (004B51A0) --------------------------------------------------------
int sub_4B51A0(int a1) {
	int result; // eax

	result = a1;
	if (a1) {
		if ((signed char)*(uint8_t*)(a1 + 4) >= 0) {
			result = nox_window_set_all_funcs((uint32_t*)a1, sub_4B4860, sub_4B51E0, 0);
		} else {
			result = nox_window_set_all_funcs((uint32_t*)a1, sub_4B4860, sub_4B52C0, 0);
		}
	}
	return result;
}

//----- (004B51E0) --------------------------------------------------------
int sub_4B51E0(int a1, int a2) {
	int v2;    // ebp
	int v3;    // ebx
	int v4;    // edi
	int xLeft; // [esp+10h] [ebp-8h]
	int yTop;  // [esp+14h] [ebp-4h]

	v2 = *(uint32_t*)(a2 + 28);
	v3 = *(uint32_t*)(a2 + 20);
	nox_client_wndGetPosition_46AA60((uint32_t*)a1, &xLeft, &yTop);
	if (*(uint8_t*)(a1 + 4) & 8) {
		if (*(uint8_t*)a2 & 2 && *(uint32_t*)(a2 + 36) != 0x80000000) {
			nox_client_drawSetColor_434460(*(uint32_t*)(a2 + 36));
			nox_client_drawBorderLines_49CC70(xLeft, yTop, *(uint32_t*)(a1 + 8), *(uint32_t*)(a1 + 12));
		}
	} else {
		v3 = *(uint32_t*)(a2 + 44);
	}
	if (v3 != 0x80000000) {
		nox_client_drawSetColor_434460(v3);
		nox_client_drawRectFilledOpaque_49CE30(xLeft + 1, yTop + 1, *(uint32_t*)(a1 + 8) - 2,
											   *(uint32_t*)(a1 + 12) - 2);
	}
	v4 = yTop + *(uint32_t*)(a1 + 12) / 2;
	if (v2 != 0x80000000) {
		nox_client_drawSetColor_434460(v2);
		nox_client_drawRectFilledOpaque_49CE30(xLeft, v4 - 1, *(uint32_t*)(a1 + 8), 3);
	}
	return 1;
}

//----- (004B52C0) --------------------------------------------------------
int sub_4B52C0(int a1, int a2) {
	int v2; // esi
	int v4; // [esp+Ch] [ebp-8h]
	int v5; // [esp+10h] [ebp-4h]

	v2 = *(uint32_t*)(a2 + 24);
	nox_client_wndGetPosition_46AA60((uint32_t*)a1, &v4, &v5);
	if (!(*(uint8_t*)(a1 + 4) & 8)) {
		v2 = *(uint32_t*)(a2 + 48);
	}
	if (v2) {
		nox_client_drawImageAt_47D2C0(v2, v4, v5);
	}
	return 1;
}

//----- (004B5320) --------------------------------------------------------
int nox_xxx_wndScrollBoxProc_4B5320(int a1, unsigned int a2, int a3, unsigned int a4) {
	int v4;           // edi
	int v5;           // esi
	signed int v6;    // ebx
	int v7;           // eax
	long long v8;     // rax
	int v9;           // ecx
	int v11;          // esi
	int v12;          // eax
	int v13;          // eax
	unsigned int v14; // eax
	int v15;          // ecx
	int v16;          // eax
	double v17;       // st7
	int v18;          // eax
	int v19;          // eax
	int* v20;         // [esp-4h] [ebp-Ch]

	v4 = a1;
	v5 = *(uint32_t*)(a1 + 32);
	if (a2 > 0x4007) {
		if (a2 == 16394) {
			if (a3 >= *(int*)v5) {
				v18 = *(uint32_t*)(v5 + 4);
				if (a3 <= v18) {
					v19 = v18 - a3;
					*(uint32_t*)(v5 + 12) = a3;
					a3 = v19;
					nox_window_setPos_46A9B0(*(uint32_t**)(v4 + 400), 0, (long long)((double)v19 * *(float*)(v5 + 8)));
				}
			}
		} else if (a2 == 16395) {
			v14 = a4;
			v15 = a3;
			*(uint32_t*)(v5 + 4) = a4;
			*(uint32_t*)v5 = v15;
			a3 = v14 - v15;
			v16 = *(uint32_t*)(v4 + 12);
			*(uint32_t*)(v5 + 12) = v15;
			v17 = (double)a3;
			a3 = v16 - 10;
			*(float*)(v5 + 8) = (double)(v16 - 10) / v17;
			nox_window_setPos_46A9B0(*(uint32_t**)(v4 + 400), 0, (long long)(v17 * *(float*)(v5 + 8)));
			return 0;
		}
		return 0;
	}
	switch (a2) {
	case 0x4007u:
		nox_window_call_field_94(*(uint32_t*)(a1 + 52), 16396, a1, *(uint32_t*)(v5 + 12));
		return 0;
	case 2u:
		free(*(void**)(a1 + 32));
		return 0;
	case 0x17u:
		v11 = a3;
		v12 = *(uint32_t*)(a1 + 36);
		if (a3) {
			LOBYTE(v12) = v12 | 2;
		} else {
			LOBYTE(v12) = v12 & 0xFD;
		}
		v20 = (int*)a1;
		*(uint32_t*)(a1 + 36) = v12;
		v13 = nox_xxx_wndGetID_46B0A0(v20);
		nox_window_call_field_94(*(uint32_t*)(v4 + 52), 16387, v11, v13);
		return 1;
	}
	if (a2 != 0x4000) {
		return 0;
	}
	v6 = a4 >> 16;
	nox_client_wndGetPosition_46AA60((uint32_t*)a1, &a1, &a3);
	if (v6 >= a3) {
		v7 = *(uint32_t*)(v4 + 12);
		if (v6 < v7 + a3) {
			a4 = v6 - a3;
			v8 = (long long)((double)(v6 - a3) / *(float*)(v5 + 8));
			v9 = *(uint32_t*)(v5 + 4);
			*(uint32_t*)(v5 + 12) = v8;
			if ((int)v8 > v9) {
				*(uint32_t*)(v5 + 12) = v9;
			}
			*(uint32_t*)(v5 + 12) = v9 - *(uint32_t*)(v5 + 12);
		} else {
			nox_window_setPos_46A9B0(*(uint32_t**)(v4 + 400), 0, v7 - *(uint32_t*)(*(uint32_t*)(v4 + 400) + 12));
			*(uint32_t*)(v5 + 12) = *(uint32_t*)v5;
		}
	} else {
		nox_window_setPos_46A9B0(*(uint32_t**)(v4 + 400), 0, 0);
		*(uint32_t*)(v5 + 12) = *(uint32_t*)(v5 + 4);
	}
	nox_window_call_field_94(*(uint32_t*)(v4 + 52), 16393, v4, *(uint32_t*)(v5 + 12));
	return 0;
}

//----- (004B5500) --------------------------------------------------------
int nox_xxx_wndScrollBoxSetAllFn_4B5500(int a1) {
	int result; // eax

	result = a1;
	if (a1) {
		if ((signed char)*(uint8_t*)(a1 + 4) >= 0) {
			result = nox_window_set_all_funcs((uint32_t*)a1, nox_xxx_wndScrollBoxDraw_4B4BA0,
											  nox_xxx_wndScrollBoxDraw_4B5540, 0);
		} else {
			result = nox_window_set_all_funcs((uint32_t*)a1, nox_xxx_wndScrollBoxDraw_4B4BA0,
											  nox_xxx_wndScrollBoxDraw_4B5620, 0);
		}
	}
	return result;
}

//----- (004B5540) --------------------------------------------------------
int nox_xxx_wndScrollBoxDraw_4B5540(int a1, int a2) {
	int v2;    // ebx
	int v3;    // ebp
	int v4;    // edi
	int xLeft; // [esp+10h] [ebp-8h]
	int yTop;  // [esp+14h] [ebp-4h]

	v2 = *(uint32_t*)(a2 + 28);
	v3 = *(uint32_t*)(a2 + 20);
	nox_client_wndGetPosition_46AA60((uint32_t*)a1, &xLeft, &yTop);
	if (*(uint8_t*)(a1 + 4) & 8) {
		if (*(uint8_t*)a2 & 2) {
			v2 = *(uint32_t*)(a2 + 52);
			if (*(uint32_t*)(a2 + 36) != 0x80000000) {
				nox_client_drawSetColor_434460(*(uint32_t*)(a2 + 36));
				nox_client_drawBorderLines_49CC70(xLeft, yTop, *(uint32_t*)(a1 + 8), *(uint32_t*)(a1 + 12));
			}
		}
	} else {
		v3 = *(uint32_t*)(a2 + 44);
	}
	if (v3 != 0x80000000) {
		nox_client_drawSetColor_434460(v3);
		nox_client_drawRectFilledOpaque_49CE30(xLeft + 1, yTop + 1, *(uint32_t*)(a1 + 8) - 2,
											   *(uint32_t*)(a1 + 12) - 2);
	}
	v4 = xLeft + *(uint32_t*)(a1 + 8) / 2;
	if (v2 != 0x80000000) {
		nox_client_drawSetColor_434460(v2);
		nox_client_drawRectFilledOpaque_49CE30(v4 - 1, yTop + 4, 3, *(uint32_t*)(a1 + 12) - 8);
	}
	return 1;
}

//----- (004B5620) --------------------------------------------------------
int nox_xxx_wndScrollBoxDraw_4B5620(uint32_t* a1, int a2) {
	int v2; // [esp+0h] [ebp-8h]
	int v3; // [esp+4h] [ebp-4h]

	nox_client_wndGetPosition_46AA60(a1, &v2, &v3);
	return 1;
}

//----- (004B5640) --------------------------------------------------------
int nox_xxx_wndScrollBoxButtonCreate_4B5640(int a1, int a2, int a3) {
	int v3;       // esi
	int v4;       // eax
	int result;   // eax
	char v6[332]; // [esp+Ch] [ebp-14Ch]

	v3 = *(uint32_t*)(a3 + 68);
	*(uint32_t*)v6 = 0;
	memset(&v6[4], 0, 0x148u);
	v4 = a2;
	*(uint32_t*)&v6[16] = a1;
	LOBYTE(v4) = a2 & 0xEF | 0xC;
	*(uint32_t*)&v6[8] = 1;
	*(uint32_t*)&v6[68] = v3;
	if ((v4 & 0x80u) == 0) {
		*(uint32_t*)&v6[52] = *(uint32_t*)(a3 + 52);
		*(uint32_t*)&v6[28] = *(uint32_t*)(a3 + 20);
		*(uint32_t*)&v6[20] = *(uint32_t*)(a3 + 28);
	} else {
		*(uint32_t*)&v6[32] = *(uint32_t*)(a3 + 32);
		*(uint32_t*)&v6[48] = *(uint32_t*)(a3 + 48);
		*(uint32_t*)&v6[56] = *(uint32_t*)(a3 + 56);
		*(uint32_t*)&v6[24] = *(uint32_t*)(a3 + 24);
		*(uint32_t*)&v6[40] = *(uint32_t*)(a3 + 40);
	}
	if (*(uint8_t*)(a3 + 8) & 0x10) {
		result = nox_gui_newButtonOrCheckbox_4A91A0(a1, v4, 0, 0, 10, *(uint32_t*)(a1 + 12), v6);
	} else {
		result = nox_gui_newButtonOrCheckbox_4A91A0(a1, v4, 0, 0, *(uint32_t*)(a1 + 8), 10, v6);
	}
	return result;
}

//----- (004B5CD0) --------------------------------------------------------
int sub_4B5CD0() {
	char v0;           // si
	unsigned char* v1; // edi
	int v2;            // edi
	unsigned char* v3; // esi
	int v4;            // eax
	int v5;            // edx
	int v6;            // esi
	unsigned char* v7; // edi
	int v8;            // esi
	unsigned char* v9; // edi
	int result;        // eax
	int v12;           // [esp+8h] [ebp-4h]

	v0 = 0;
	v1 = getMemAt(0x5D4594, 1312500);
	do {
		LOBYTE(v12) = v0 / 63;
		*(uint32_t*)v1 = nox_color_rgb_4344A0((unsigned char)(v0 / 63) / 3, 3 * (unsigned char)(v0 / 63) / 5, v12);
		v1 += 4;
		--v0;
	} while ((int)v1 < (int)getMemAt(0x5D4594, 1312756));
	v2 = 0;
	v3 = getMemAt(0x5D4594, 1312756);
	do {
		LOBYTE(v12) = v2 / 32;
		v4 = nox_color_rgb_4344A0(v12, v12, 0);
		LOBYTE(v5) = -1 - v12;
		*(uint32_t*)v3 = v4;
		*((uint32_t*)v3 + 32) = nox_color_rgb_4344A0(255, v5, 0);
		v3 += 4;
		v2 += 255;
	} while ((int)v3 < (int)getMemAt(0x5D4594, 1312884));
	v6 = 0;
	v7 = getMemAt(0x5D4594, 1313012);
	do {
		LOBYTE(v12) = v6 / 63;
		*(uint32_t*)v7 = nox_color_rgb_4344A0(v12, v12, v12);
		v7 += 4;
		v6 += 255;
	} while ((int)v7 < (int)getMemAt(0x5D4594, 1313268));
	v8 = 0;
	v9 = getMemAt(0x5D4594, 1313268);
	do {
		result = nox_color_rgb_4344A0(v8 / 63, 50, 50);
		*(uint32_t*)v9 = result;
		v9 += 4;
		v8 += 255;
	} while ((int)v9 < (int)getMemAt(0x5D4594, 1313524));
	return result;
}
// 4B5D21: variable 'v12' is possibly undefined
// 4B5D71: variable 'v5' is possibly undefined

//----- (004B63B0) --------------------------------------------------------
int sub_4B63B0(int2* a1, int2* a2) {
	int v2; // ebx
	int v3; // ebp

	nox_client_drawSetColor_434460(*getMemIntPtr(0x5D4594, 1312492));
	nox_client_drawAddPoint_49F500(a1->field_0, a1->field_4);
	nox_client_drawAddPoint_49F500(a2->field_0, a2->field_4);
	nox_client_drawLineFromPoints_49E4B0();
	v2 = a2->field_0 - a1->field_0;
	v3 = a2->field_4 - a1->field_4;
	nox_client_drawSetColor_434460(*getMemIntPtr(0x5D4594, 1312496));
	if (v2 < 0) {
		v2 = -v2;
	}
	if (v3 < 0) {
		v3 = -v3;
	}
	if (v2 <= v3) {
		nox_client_drawAddPoint_49F500(a1->field_0 - 1, a1->field_4);
		nox_client_drawAddPoint_49F500(a2->field_0 - 1, a2->field_4);
		nox_client_drawLineFromPoints_49E4B0();
		nox_client_drawAddPoint_49F500(a1->field_0 + 1, a1->field_4);
		nox_client_drawAddPoint_49F500(a2->field_0 + 1, a2->field_4);
	} else {
		nox_client_drawAddPoint_49F500(a1->field_0, a1->field_4 - 1);
		nox_client_drawAddPoint_49F500(a2->field_0, a2->field_4 - 1);
		nox_client_drawLineFromPoints_49E4B0();
		nox_client_drawAddPoint_49F500(a1->field_0, a1->field_4 + 1);
		nox_client_drawAddPoint_49F500(a2->field_0, a2->field_4 + 1);
	}
	return nox_client_drawLineFromPoints_49E4B0();
}

//----- (004B64C0) --------------------------------------------------------
int sub_4B64C0() {
	unsigned char v0;  // bl
	unsigned char* v1; // esi
	int v2;            // ebp
	int v3;            // edi
	int v5;            // [esp+10h] [ebp-10h]
	int v6;            // [esp+14h] [ebp-Ch]
	int v7;            // [esp+18h] [ebp-8h]
	int v8;            // [esp+1Ch] [ebp-4h]

	*getMemU32Ptr(0x5D4594, 1313524) = nox_color_rgb_4344A0(255, 255, 0);
	*getMemU32Ptr(0x5D4594, 1313528) = nox_color_rgb_4344A0(255, 100, 0);
	dword_5d4594_1313532 = nox_color_rgb_4344A0(255, 255, 0);
	dword_5d4594_1313536 = nox_color_rgb_4344A0(0, 0, 255);
	dword_5d4594_1313540 = nox_color_rgb_4344A0(0, 200, 255);
	*getMemU32Ptr(0x5D4594, 1313544) = nox_color_rgb_4344A0(0, 200, 200);
	*getMemU32Ptr(0x5D4594, 1313548) = nox_color_rgb_4344A0(50, 255, 255);
	*getMemU32Ptr(0x5D4594, 1313552) = nox_color_rgb_4344A0(255, 0, 255);
	*getMemU32Ptr(0x5D4594, 1313556) = nox_color_rgb_4344A0(255, 200, 255);
	*getMemU32Ptr(0x5D4594, 1313560) = nox_color_rgb_4344A0(255, 200, 0);
	dword_5d4594_1313564 = nox_color_rgb_4344A0(255, 255, 100);
	*getMemU32Ptr(0x5D4594, 1313568) = nox_color_rgb_4344A0(100, 255, 50);
	*getMemU32Ptr(0x5D4594, 1313572) = nox_color_rgb_4344A0(150, 255, 150);
	*getMemU32Ptr(0x5D4594, 1313576) = nox_color_rgb_4344A0(255, 255, 0);
	*getMemU32Ptr(0x5D4594, 1313580) = nox_color_rgb_4344A0(0, 220, 0);
	*getMemU32Ptr(0x5D4594, 1313584) = nox_color_rgb_4344A0(150, 255, 150);
	*getMemU32Ptr(0x5D4594, 1313588) = nox_color_rgb_4344A0(200, 200, 200);
	*getMemU32Ptr(0x5D4594, 1313592) = nox_color_rgb_4344A0(255, 255, 255);
	LOBYTE(v8) = -1;
	LOBYTE(v7) = -1;
	LOBYTE(v6) = -1;
	v0 = 0;
	v1 = getMemAt(0x5D4594, 1313656);
	v2 = 1085;
	v5 = 600;
	v3 = 765;
	do {
		if (v0 > 3u) {
			if (v0 > 7u) {
				LOBYTE(v8) = v2 / 9 - 1;
			} else {
				LOBYTE(v7) = v5 / 4 - 1;
			}
		} else {
			LOBYTE(v6) = v3 / 3;
		}
		*(uint32_t*)v1 = nox_color_rgb_4344A0(v8, v7, v6);
		++v0;
		v3 -= 255;
		v2 -= 155;
		v1 -= 4;
		v5 -= 200;
	} while (v0 < 0x10u);
	return sub_4B5CD0();
}
// 4B66DF: variable 'v8' is possibly undefined
// 4B66DF: variable 'v7' is possibly undefined
// 4B66DF: variable 'v6' is possibly undefined

//----- (004B6880) --------------------------------------------------------
int sub_4B6880(uint32_t* a1, int a2, int a3, int a4) {
	int v4;     // edx
	int v5;     // ebx
	int v6;     // eax
	int result; // eax
	int v8;     // ecx
	int v9;     // eax
	int v10;    // edx
	int v11;    // eax
	bool v12;   // cc
	int v13;    // esi
	int2 xLeft; // [esp+Ch] [ebp-8h]
	int v15;    // [esp+1Ch] [ebp+8h]

	v4 = a2;
	v5 = *(uint32_t*)(a2 + 448) - *(uint32_t*)(a2 + 444);
	v15 = *(uint32_t*)(a2 + 448) - gameFrame();
	v6 = v15;
	if (v15 == v5) {
		v6 = --v15;
	}
	if (v6 > 0) {
		v8 = *a1 + *(uint32_t*)(v4 + 12) - a1[4];
		v9 = *(uint32_t*)(v4 + 16) - *(short*)(v4 + 106) - *(short*)(v4 + 104) - a1[5];
		v10 = a1[1];
		v11 = v10 + v9;
		v12 = v8 - 10 < *a1;
		xLeft.field_0 = v8;
		xLeft.field_4 = v11;
		if (!v12 && v11 - 10 >= v10 && v8 + 10 < a1[2] && v11 + 10 < a1[3]) {
			v13 = 4 * v15 / v5;
			sub_4B6720(&xLeft, a4, 2 * v13 + 1, 5 * v15 / v5);
			nox_client_drawSetColor_434460(a3);
			nox_xxx_drawPointMB_499B70(xLeft.field_0, xLeft.field_4, v13);
		}
		result = 1;
	} else {
		nox_xxx_spriteDeleteStatic_45A4E0_drawable(v4);
		result = 0;
	}
	return result;
}

//----- (004B6970) --------------------------------------------------------
int sub_4B6970(uint32_t* a1, nox_drawable* dr, int a3, int a4) {
	int v4; // eax
	int v5; // eax

	int a2 = dr;

	v4 = *(uint32_t*)(a2 + 440);
	*(uint32_t*)(a2 + 432) += v4 * *getMemIntPtr(0x587000, 192088 + 8 * *(unsigned char*)(a2 + 299));
	v5 = v4 * *getMemIntPtr(0x587000, 192092 + 8 * *(unsigned char*)(a2 + 299)) + *(uint32_t*)(a2 + 436);
	*(uint32_t*)(a2 + 436) = v5;
	nox_xxx_updateSpritePosition_49AA90((uint32_t*)a2, *(uint32_t*)(a2 + 432) >> 12, v5 >> 12);
	sub_4B69F0(a2);
	return sub_4B6880(a1, a2, a3, a4);
}

//----- (004B69F0) --------------------------------------------------------
short sub_4B69F0(int a1) {
	char v1;      // cl
	short result; // ax
	int v3;       // edx

	v1 = *(uint8_t*)(a1 + 296);
	*(uint16_t*)(a1 + 104) += v1;
	result = *(uint16_t*)(a1 + 104);
	if (result >= 0) {
		if ((unsigned char)gameFrame() & 1) {
			*(uint8_t*)(a1 + 296) = v1 - 1;
		}
	} else {
		*(uint16_t*)(a1 + 104) = -result;
		result = 26209 * v1;
		v3 = -9 * v1 / 10;
		*(uint8_t*)(a1 + 296) = v3;
		if ((char)v3 < 2) {
			*(uint16_t*)(a1 + 104) = 0;
			*(uint8_t*)(a1 + 296) = 0;
		}
	}
	return result;
}

//----- (004B6B80) --------------------------------------------------------
int sub_4B6B80(int* a1, nox_drawable* dr, int a3) {
	int v3;     // esi
	int v4;     // eax
	int v5;     // ecx
	int v6;     // edi
	int v7;     // ebx
	int v8;     // eax
	int v10;    // edi
	int v11;    // ebp
	int v12;    // edx
	int v13;    // edx
	int v14;    // edi
	int v15;    // ebx
	char v16;   // al
	char v17;   // cl
	char v18;   // al
	int v19;    // [esp+10h] [ebp-Ch]
	int2 xLeft; // [esp+14h] [ebp-8h]
	int v21;    // [esp+24h] [ebp+8h]

	int a2 = dr;

	v3 = a2;
	if (!*getMemU32Ptr(0x5D4594, 1313660)) {
		*getMemU32Ptr(0x5D4594, 1313660) = nox_xxx_getTTByNameSpriteMB_44CFC0("DrainManaOrb");
		*getMemU32Ptr(0x5D4594, 1313664) = nox_xxx_getTTByNameSpriteMB_44CFC0("HealOrb");
		*getMemU32Ptr(0x5D4594, 1313668) = nox_xxx_getTTByNameSpriteMB_44CFC0("CharmOrb");
		*getMemU32Ptr(0x5D4594, 1313672) = nox_xxx_getTTByNameSpriteMB_44CFC0("WhiteOrb");
		*getMemU32Ptr(0x5D4594, 1313676) = nox_xxx_getTTByNameSpriteMB_44CFC0("ManaBombOrb");
		*getMemU32Ptr(0x5D4594, 1313680) = nox_xxx_getTTByNameSpriteMB_44CFC0("WhiteMoveOrb");
		*getMemU32Ptr(0x5D4594, 1313684) = nox_xxx_getTTByNameSpriteMB_44CFC0("BlueMoveOrb");
	}
	v4 = *(uint32_t*)(a2 + 108);
	if (v4 == *getMemU32Ptr(0x5D4594, 1313660) || v4 == *getMemU32Ptr(0x5D4594, 1313684)) {
		v5 = dword_5d4594_1313540;
		v21 = dword_5d4594_1313536;
		v19 = v5;
		goto LABEL_14;
	}
	if (v4 == *getMemU32Ptr(0x5D4594, 1313668)) {
		v5 = *getMemU32Ptr(0x5D4594, 1313584);
		v21 = *getMemU32Ptr(0x5D4594, 1313580);
		v19 = v5;
		goto LABEL_14;
	}
	if (v4 == *getMemU32Ptr(0x5D4594, 1313672) || v4 == *getMemU32Ptr(0x5D4594, 1313676) ||
		v4 == *getMemU32Ptr(0x5D4594, 1313680)) {
		v21 = *getMemU32Ptr(0x5D4594, 1313588);
		v19 = *getMemU32Ptr(0x5D4594, 1313592);
	} else {
		v21 = *getMemU32Ptr(0x5D4594, 1313528);
		v19 = dword_5d4594_1313532;
	}
LABEL_14:
	if (a3) {
		v6 = *(unsigned short*)(v3 + 432) - *(uint32_t*)(v3 + 12);
		v7 = *(unsigned short*)(v3 + 434) - *(uint32_t*)(v3 + 16);
		v8 = sub_48C6B0(v6, v7);
		if (v8 + 1 <= 10) {
			nox_xxx_spriteDeleteStatic_45A4E0_drawable(v3);
			return 0;
		}
		nox_xxx_updateSpritePosition_49AA90((uint32_t*)v3,
											*(uint32_t*)(v3 + 12) + v6 * *(unsigned char*)(v3 + 443) / (v8 + 1),
											*(uint32_t*)(v3 + 16) + v7 * *(unsigned char*)(v3 + 443) / (v8 + 1));
	}
	v10 = *a1;
	v11 = a1[1];
	v12 = *(uint32_t*)(v3 + 16) - a1[5];
	xLeft.field_0 = *(uint32_t*)(v3 + 12) + *a1 - a1[4];
	xLeft.field_4 = v12 + v11 - 22;
	v13 = *(unsigned char*)(v3 + 444);
	if (xLeft.field_0 - v13 < v10) {
		return 1;
	}
	if (xLeft.field_4 - v13 < v11) {
		return 1;
	}
	if (v13 + xLeft.field_0 >= a1[2]) {
		return 1;
	}
	if (v13 + xLeft.field_4 >= a1[3]) {
		return 1;
	}
	sub_4B6720(&xLeft, v21, v13, 5);
	nox_client_drawSetColor_434460(v19);
	nox_xxx_drawPointMB_499B70(xLeft.field_0, xLeft.field_4, *(unsigned char*)(v3 + 444) >> 1);
	v14 = *(uint32_t*)(v3 + 32) - *(uint32_t*)(v3 + 12);
	v15 = *(uint32_t*)(v3 + 36) - *(uint32_t*)(v3 + 16);
	nox_client_drawSetColor_434460(v19);
	nox_client_drawAddPoint_49F500(xLeft.field_0, xLeft.field_4);
	nox_xxx_rasterPointRel_49F570(v14, v15);
	nox_client_drawLineFromPoints_49E4B0();
	v16 = *(uint8_t*)(v3 + 445);
	if (!v16) {
		return 1;
	}
	v17 = *(uint8_t*)(v3 + 446) - 1;
	*(uint8_t*)(v3 + 446) = v17;
	if (v17) {
		return 1;
	}
	*(uint8_t*)(v3 + 446) = v16;
	v18 = *(uint8_t*)(v3 + 444) - 1;
	*(uint8_t*)(v3 + 444) = v18;
	if (v18) {
		return 1;
	}
	nox_xxx_spriteDeleteStatic_45A4E0_drawable(v3);
	return 0;
}
// 4B6CB4: variable 'v8' is possibly undefined

//----- (004B7C40) --------------------------------------------------------
uint32_t* nox_xxx_netHandleSummonPacket_4B7C40(short a1, unsigned short* a2, unsigned short a3, unsigned char a4,
											   short a5) {
	int v5;           // eax
	uint32_t* result; // eax
	uint32_t* v7;     // edi
	uint32_t* v8;     // esi
	int v9;           // [esp-8h] [ebp-1Ch]
	int v10;          // [esp-4h] [ebp-18h]
	int v11;          // [esp+10h] [ebp-4h]

	v10 = a2[1];
	v9 = *a2;
	v5 = nox_xxx_getTTByNameSpriteMB_44CFC0("SummonEffect");
	result = (uint32_t*)nox_xxx_spriteLoadAdd_45A360_drawable(v5, v9, v10);
	v7 = result;
	if (result) {
		result = nox_new_drawable_for_thing(a3);
		v8 = result;
		if (result) {
			result[3] = *a2;
			result[4] = a2[1];
			*((uint8_t*)result + 297) = nox_xxx_math_509EA0(a4);
			HIWORD(v11) = a1;
			LOWORD(v11) = a5;
			v8[69] = 8;
			v7[108] = v8;
			v7[109] = v11;
			result = gameFrame();
			v7[79] = gameFrame();
		}
	}
	return result;
}

//----- (004B7EE0) --------------------------------------------------------
void sub_4B7EE0(short a1) {
	if (!*getMemU32Ptr(0x5D4594, 1313744)) {
		*getMemU32Ptr(0x5D4594, 1313744) = nox_xxx_getTTByNameSpriteMB_44CFC0("SummonEffect");
	}
	if (!dword_5d4594_1313740) {
		dword_5d4594_1313740 = nox_xxx_getTTByNameSpriteMB_44CFC0("BlueSpark");
	}

	int v2 = sub_45A060();
	if (!v2) {
		return;
	}

	while (*(uint32_t*)(v2 + 108) != *getMemU32Ptr(0x5D4594, 1313744) || *(uint16_t*)(v2 + 438) != a1) {
		v2 = nox_drawable_next_45A070(v2);
		if (!v2) {
			return;
		}
	}
	nox_xxx_makePointFxCli_499610(*(int*)&dword_5d4594_1313740, 50, 1000, 30, *(uint32_t*)(v2 + 12),
								  *(uint32_t*)(v2 + 16));
	nox_xxx_spriteDelete_45A4B0(*(uint64_t**)(v2 + 432));
	nox_xxx_spriteDeleteStatic_45A4E0_drawable(v2);
}

//----- (004B7F90) --------------------------------------------------------
int nox_xxx_spriteShieldLoad_4B7F90() {
	int result; // eax

	*getMemU32Ptr(0x5D4594, 1313748) = nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldNW");
	*getMemU32Ptr(0x5D4594, 1313752) = nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldN");
	*getMemU32Ptr(0x5D4594, 1313756) = nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldNE");
	*getMemU32Ptr(0x5D4594, 1313760) = nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldW");
	*getMemU32Ptr(0x5D4594, 1313768) = nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldE");
	*getMemU32Ptr(0x5D4594, 1313772) = nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldSW");
	*getMemU32Ptr(0x5D4594, 1313776) = nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldS");
	result = nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldSE");
	*getMemU32Ptr(0x5D4594, 1313780) = result;
	*getMemU32Ptr(0x5D4594, 1313764) = 0;
	*getMemU32Ptr(0x5D4594, 1313784) = 1;
	return result;
}
// 4B8040: variable 'v3' is possibly undefined
// 4B804D: variable 'v5' is possibly undefined

//----- (004B8090) --------------------------------------------------------
uint32_t* nox_xxx_fxShield_4B8090(unsigned int a1, int a2) {
	int v2;           // edi
	int v3;           // eax
	uint32_t* result; // eax
	int v5;           // eax
	uint32_t* v6;     // esi
	int4 v7;          // [esp+0h] [ebp-10h]

	if (!*getMemU32Ptr(0x5D4594, 1313784)) {
		nox_xxx_spriteShieldLoad_4B7F90();
	}
	v2 = a2;
	switch (a2) {
	case 0:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldNW");
		break;
	case 1:
		nox_xxx_getTTByNameSpriteMB_44CFC0("ShpericalShieldN");
		break;
	case 2:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldNW");
		break;
	case 3:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldW");
		break;
	case 5:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldE");
		break;
	case 6:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldSW");
		break;
	case 7:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldS");
		break;
	case 8:
		nox_xxx_getTTByNameSpriteMB_44CFC0("SphericalShieldSE");
		break;
	default:
		break;
	}
	if (nox_xxx_netTestHighBit_578B70(a1)) {
		v3 = nox_xxx_netClearHighBit_578B30(a1);
		result = nox_xxx_netSpriteByCodeStatic_45A720(v3);
	} else {
		v5 = nox_xxx_netClearHighBit_578B30(a1);
		result = nox_xxx_netSpriteByCodeDynamic_45A6F0(v5);
	}
	v6 = result;
	if (result) {
		v7.field_0 = result[3] - 10;
		v7.field_4 = result[4] - 10;
		v7.field_8 = result[3] + 10;
		v7.field_C = result[4] + 10;
		dword_5d4594_1313788 = 0;
		nox_xxx_forEachSprite_49AB00(&v7, nox_xxx_spriteScanForShield_4B81E0, (int)&a1);
		result = *(uint32_t**)&dword_5d4594_1313788;
		if (dword_5d4594_1313788 != 1) {
			result = (uint32_t*)nox_xxx_spriteLoadAdd_45A360_drawable(*getMemU32Ptr(0x5D4594, 1313748 + 4 * v2), v6[3],
																	  v6[4] + 3);
			if (result) {
				result[108] = a1;
			}
		}
	}
	return result;
}
// 4B810D: variable 'v3' is possibly undefined
// 4B811F: variable 'v5' is possibly undefined

//----- (004B81E0) --------------------------------------------------------
void nox_xxx_spriteScanForShield_4B81E0(int a1, int a2) {
	unsigned char* v2; // eax

	v2 = getMemAt(0x5D4594, 1313748);
	do {
		if (*(uint32_t*)(a1 + 108) == *(uint32_t*)v2 && *(uint32_t*)(a1 + 432) == *(uint32_t*)a2) {
			dword_5d4594_1313788 = 1;
		}
		v2 += 4;
	} while ((int)v2 < (int)getMemAt(0x5D4594, 1313784));
}

//----- (004B8E10) --------------------------------------------------------
uint32_t* sub_4B8E10(uint32_t* a1, char* a2) {
	uint32_t* result; // eax
	int v3;           // ebx
	int v4;           // eax
	int v5;           // edx
	int v6;           // ecx
	uint32_t* v7;     // ebp
	int v8;           // edi
	uint8_t* v9;      // esi
	uint32_t* v10;    // ecx
	int v11;          // eax
	int* v12;         // edi
	int v13;          // ebx
	uint32_t** v14;   // esi
	int v15;          // [esp-Ch] [ebp-14h]

	result = a1;
	v3 = 0;
	while ((char*)*result != a2) {
		++v3;
		result += 6;
		if (v3 >= 27) {
			return result;
		}
	}
	v4 = sub_415840(a2);
	result = nox_xxx_getProjectileClassById_413250(v4);
	v7 = result;
	if (result) {
		v8 = 1;
		v9 = result + 4;
		do {
			LOBYTE(result) = v9[1];
			LOBYTE(v6) = *v9;
			LOBYTE(v5) = *(v9 - 1);
			nox_draw_setMaterial_4340A0(v8++, v5, v6, (int)result);
			v9 += 3;
		} while (v8 < 7);
		v10 = a1;
		v11 = 3 * v3;
		v12 = v7 + 9;
		v13 = 4;
		v14 = (uint32_t**)&a1[2 * v11 + 1];
		do {
			result = *v14;
			if (*v14) {
				LOBYTE(v5) = *((uint8_t*)result + 26);
				LOBYTE(v10) = *((uint8_t*)result + 25);
				v15 = v5;
				LOBYTE(v5) = *((uint8_t*)result + 24);
				nox_draw_setMaterial_4340A0(*v12, v5, (int)v10, v15);
			}
			++v14;
			++v12;
			--v13;
		} while (v13);
	}
	return result;
}
// 4B8E57: variable 'v5' is possibly undefined
// 4B8E57: variable 'v6' is possibly undefined
// 4B8E90: variable 'v10' is possibly undefined
