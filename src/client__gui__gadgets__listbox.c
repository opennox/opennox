#include "client__gui__gadgets__listbox.h"

#include "proto.h"

//----- (004A4310) --------------------------------------------------------
nox_window*  nox_gui_newScrollListBox_4A4310(int a1, int a2, int a3, int a4, int a5, int a6, int a7, __int16* a8) {
	_DWORD* v8;       // ebp
	void* v9;         // edi
	int v10;          // edi
	int v11;          // ecx
	void* v12;        // eax
	int v14;          // ecx
	int v15;          // eax
	wchar_t* v16;     // eax
	int v17;          // ebx
	wchar_t* v18;     // eax
	int v19;          // eax
	__int16* v20;     // eax
	int v21;          // [esp+10h] [ebp-168h]
	int v22;          // [esp+10h] [ebp-168h]
	int v23;          // [esp+14h] [ebp-164h]
	int v24[4];       // [esp+1Ch] [ebp-15Ch]
	char v25[332];    // [esp+2Ch] [ebp-14Ch]
	unsigned int v26; // [esp+180h] [ebp+8h]

	v21 = 0;
	if (a8[1] < nox_xxx_guiFontHeightMB_43F320(*(_DWORD*)(a7 + 200)))
		a8[1] = nox_xxx_guiFontHeightMB_43F320(*(_DWORD*)(a7 + 200));
	if (*(_WORD*)(a7 + 72))
		v21 = 1;
	if (!(*(_BYTE*)(a7 + 8) & 0x20))
		return 0;
	v8 = nox_window_new(a1, a2, a3, a4, a5, a6, nox_xxx_wndListboxProcPre_4A30D0);
	nox_xxx_wndListboxInit_4A3C00((int)v8, (int)a8);
	if (v8) {
		if (!*(_DWORD*)(a7 + 16))
			*(_DWORD*)(a7 + 16) = v8;
		nox_gui_windowCopyDrawData_46AF80((int)v8, (const void*)a7);
		v9 = malloc(524 * *a8);
		*((_DWORD*)a8 + 6) = v9;
		if (!v9)
			return 0;
		memset(v9, 0, 524 * *a8);
		v10 = a6;
		a8[26] = a6;
		if (v21)
			a8[26] -= nox_xxx_guiFontHeightMB_43F320(*(_DWORD*)(a7 + 200));
		v11 = *((_DWORD*)a8 + 4);
		a8[27] = 0;
		*((_DWORD*)a8 + 12) = -1;
		a8[23] = 0;
		a8[22] = 0;
		*((_DWORD*)a8 + 10) = 0;
		if (v11) {
			v12 = malloc(4 * *a8);
			*((_DWORD*)a8 + 12) = v12;
			if (!v12) {
				free(*((LPVOID*)a8 + 6));
				return 0;
			}
			memset(v12, 0xFFu, 4 * *a8);
			v10 = a6;
		}
		if (*((_DWORD*)a8 + 3)) {
			v24[0] = 0;
			v14 = *(_DWORD*)(a7 + 200);
			v24[1] = 0;
			v24[2] = 0;
			v24[3] = 0;
			v26 = a2 & 0xFFFFEFEF;
			v15 = nox_xxx_guiFontHeightMB_43F320(v14);
			if (v21) {
				v22 = v15 + 1;
				v10 = v10 - v15 - 1;
			} else {
				v22 = 0;
			}
			memset(v25, 0, sizeof(v25));
			if ((signed char)*((_BYTE*)v8 + 4) >= 0) {
				*(_DWORD*)&v25[20] = *getMemU32Ptr(0x5D4594, 2650656);
				*(_DWORD*)&v25[44] = *getMemU32Ptr(0x5D4594, 2650656);
				*(_DWORD*)&v25[28] = *getMemU32Ptr(0x5D4594, 2614256);
				*(_DWORD*)&v25[36] = *getMemU32Ptr(0x5D4594, 2523948);
				*(_DWORD*)&v25[52] = *getMemU32Ptr(0x5D4594, 2589772);
				*(_DWORD*)&v25[68] = *getMemU32Ptr(0x5D4594, 2614256);
				v16 = nox_strman_loadString_40F1D0("WindowDir:Up", 0,
											"C:\\NoxPost\\src\\Client\\Gui\\Gadgets\\listbox.c", 1483);
				nox_wcscpy((wchar_t*)&v25[72], v16);
				v23 = 10;
			} else {
				*(_DWORD*)&v25[24] = nox_xxx_gLoadImg_42F970("DefaultLBUpButton");
				*(_DWORD*)&v25[40] = nox_xxx_gLoadImg_42F970("DefaultLBUpButtonLit");
				*(_DWORD*)&v25[48] = nox_xxx_gLoadImg_42F970("DefaultLBUpButtonDis");
				*(_DWORD*)&v25[56] = nox_xxx_gLoadImg_42F970("DefaultLBUpButtonLit");
				*(_DWORD*)&v25[32] = 0;
				v23 = 13;
			}
			v17 = v26 | 9;
			*(_DWORD*)&v25[16] = v8;
			*(_DWORD*)&v25[8] = 1;
			*((_DWORD*)a8 + 7) = nox_gui_newButtonOrCheckbox_4A91A0((int)v8, v26 | 9, a5 - 10, v22, 10, v23, v25);
			memset(v25, 0, sizeof(v25));
			if ((signed char)*((_BYTE*)v8 + 4) >= 0) {
				*(_DWORD*)&v25[20] = *getMemU32Ptr(0x5D4594, 2650656);
				*(_DWORD*)&v25[44] = *getMemU32Ptr(0x5D4594, 2650656);
				*(_DWORD*)&v25[28] = *getMemU32Ptr(0x5D4594, 2614256);
				*(_DWORD*)&v25[36] = *getMemU32Ptr(0x5D4594, 2523948);
				*(_DWORD*)&v25[52] = *getMemU32Ptr(0x5D4594, 2589772);
				*(_DWORD*)&v25[68] = *getMemU32Ptr(0x5D4594, 2614256);
				v18 = nox_strman_loadString_40F1D0("WindowDir:Down", 0,
											"C:\\NoxPost\\src\\Client\\Gui\\Gadgets\\listbox.c", 1519);
				nox_wcscpy((wchar_t*)&v25[72], v18);
			} else {
				*(_DWORD*)&v25[24] = nox_xxx_gLoadImg_42F970("DefaultLBDownButton");
				*(_DWORD*)&v25[40] = nox_xxx_gLoadImg_42F970("DefaultLBDownButtonLit");
				*(_DWORD*)&v25[48] = nox_xxx_gLoadImg_42F970("DefaultLBDownButtonDis");
				*(_DWORD*)&v25[56] = nox_xxx_gLoadImg_42F970("DefaultLBDownButtonLit");
				*(_DWORD*)&v25[32] = 0;
			}
			*(_DWORD*)&v25[8] = 1;
			*(_DWORD*)&v25[16] = v8;
			*((_DWORD*)a8 + 8) = nox_gui_newButtonOrCheckbox_4A91A0((int)v8, v17, a5 - 10, v22 + v10 - v23, 10, v23, v25);
			memset(v25, 0, sizeof(v25));
			if ((signed char)*((_BYTE*)v8 + 4) >= 0) {
				*(_DWORD*)&v25[20] = *getMemU32Ptr(0x5D4594, 2650656);
				*(_DWORD*)&v25[44] = *getMemU32Ptr(0x5D4594, 2650656);
				*(_DWORD*)&v25[36] = *getMemU32Ptr(0x5D4594, 2650656);
				*(_DWORD*)&v25[28] = *getMemU32Ptr(0x5D4594, 2614256);
				*(_DWORD*)&v25[52] = *getMemU32Ptr(0x5D4594, 2614256);
				v19 = 10;
			} else {
				*(_DWORD*)&v25[24] = nox_xxx_gLoadImg_42F970("DefaultSliderThumb");
				*(_DWORD*)&v25[40] = nox_xxx_gLoadImg_42F970("DefaultSliderThumbLit");
				*(_DWORD*)&v25[48] = nox_xxx_gLoadImg_42F970("DefaultSliderThumbDis");
				*(_DWORD*)&v25[56] = nox_xxx_gLoadImg_42F970("DefaultSliderThumbLit");
				*(_DWORD*)&v25[32] = 0;
				v19 = 9;
			}
			v24[0] = 0;
			v24[1] = 0;
			v24[2] = 0;
			*(_DWORD*)&v25[8] = 8;
			v24[3] = 0;
			*(_DWORD*)&v25[16] = v8;
			*((_DWORD*)a8 + 9) = nox_gui_newSlider_4B4EE0((int)v8, v17, a5 - v19, v22 + v23, v19, v10 - 2 * v23, v25, (float*)v24);
		}
		v20 = (__int16*)malloc(0x38u);
		memcpy(v20, a8, 0x38u);
		v8[8] = v20;
	}
	return v8;
}
