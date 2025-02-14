#ifndef NOX_PORT_GAME3
#define NOX_PORT_GAME3

#include "defs.h"
#include "common__savegame.h"

void sub_4A1A40(int a1);
int sub_4A1BE0(int a1);
int nox_client_guiXxxDestroy_4A24A0();
int sub_4A2560(uint32_t* a1, int a2);
int sub_4A25C0(uint32_t* a1, int* a2);
int sub_4A2610(int a1, uint32_t* a2, int* a3);
uint32_t* sub_4A2830(int a1, int a2, uint32_t* a3);
int sub_4A2890();
int sub_4A28B0();
int sub_4A28C0(int a1);
int nox_xxx_wndListboxProcWithoutData10_4A28E0(uint32_t* a1, int a2, unsigned int a3, int a4);
int nox_xxx_wndListBox_4A2D10(int a1, int a2, int a3);
int nox_xxx_wndListboxProcWithData10_4A2DE0(int a1, int a2, unsigned int a3, int a4);
short* sub_4A3090(short* a1, int a2);
int nox_xxx_wndListboxProcPre_4A30D0(nox_window* win, unsigned int a2, uint32_t a3, int a4);
int nox_xxx_wndListBox_4A3A70(int a1);
int nox_xxx_wndListBoxAddLine_4A3AC0(wchar2_t* a1, int a2, uint32_t* a3);
void nox_xxx_wndListboxInit_4A3C00(int a1, int a2);
int nox_xxx_wndListboxDrawNoImage_4A3C50(uint32_t* a1, int a2);
int nox_xxx_wndListboxDrawWithImage_4A3FC0(uint32_t* a1, int a2);
int sub_4A4800(int a1);
int nox_game_showSelClass_4A4840();
int sub_4A4970();
int sub_4A49A0();
int sub_4A49D0(int yTop, int a2);
int sub_4A50A0();
int sub_4A50D0();
int sub_4A5E90();
unsigned char* sub_4A61E0(uint32_t* a1, int a2, unsigned char* a3);
int sub_4A6890();
int sub_4A6B50(wchar2_t* a1);
int sub_4A6C90();
int sub_4A6D20(int a1, int a2);
int sub_4A6DC0(uint32_t* a1, int a2);
int sub_4A7270(int a1, int a2, unsigned int a3, int a4);
uint32_t* sub_4A72D0(unsigned short a1);
int sub_4A7330(int a1, int a2, int* a3, unsigned int a4);
uint32_t* sub_4A7530(unsigned short a1);
int sub_4A7580(int a1, int a2);
int sub_4A7A60(int a1);
int sub_4A7A70(int a1);
int sub_4A7A80(const char* a1);
int sub_4A7AC0(const char* a1);
int sub_4A7B00(const char* a1);
int sub_4A7B40(char* a1);
int sub_4A7BA0(char* a1);
int sub_4A7BC0(const char* a1);
int sub_4A7C00(const char* a1);
int sub_4A7C40(char* a1);
int sub_4A7C60(char* a1);
int sub_4A7CE0(char* a1);
int sub_4A7D00(const char* a1);
int sub_4A7D50(char* a1);
char* sub_4A7EF0();
int nox_xxx_wndRadioButtonProc_4A84E0(uint32_t* a1, int a2, int a3, int a4);
int nox_xxx_wndRadioButtonSetAllFn_4A87E0(int a1);
int nox_xxx_wndRadioButtonDrawNoImg_4A8820(int a1, int a2);
int nox_xxx_wndRadioButtonDraw_4A8A20(int a1, int a2);
nox_window* nox_gui_newButtonOrCheckbox_4A91A0(nox_window* parent, int a2, int a3, int a4, int a5, int a6, nox_window_data* draw);
int nox_xxx_wndRadioButtonProcPre_4A93C0(int a1, int a2, wchar2_t* a3, int a4);
int nox_xxx_compassGenStrings_4A9C80();
int nox_game_showOptions_4AA6B0();
int sub_4AA9C0();
int sub_4AAA10();
uint32_t* sub_4AAA70();
int sub_4AABE0(int a1, int a2, int* a3, int a4);
int sub_4AB0C0();
int sub_4AB260();
int sub_4AB340(int a1, int a2, int a3, int a4);
int sub_4AB390(int a1, int a2, int* a3, int a4);
int sub_4AB420(int* a1);
int sub_4AB470();
int sub_4AB4A0(int a1);
int sub_4AB4D0(int a1);
int sub_4ABDA0(int a1, short a2, short a3, uint32_t* a4);
int nox_xxx_spriteLoadFromMap_4AC020(int thingInd, short a2, uint32_t* a3);
int nox_client_mapSpecialRWObjectData_4AC610();
int nox_xxx_clientLoadSomeObject_4AC6E0(unsigned short a1);
int sub_4AC7B0(int a1);
int nox_xxx_colorLightClientLoad_4AC980(int a1);
int nox_xxx_cliLoadTeamBase_4ACE00(int a1);
int sub_4ACEF0(int a1);
int sub_4AD040(int a1);
int sub_4AD570();
int nox_xxx_windowServerOptionsGeneralProc_4AD5D0(int a1, int a2, int* a3, int a4);
int sub_4AD820();
int sub_4AD9B0(int a1);
int sub_4ADA40();
int nox_game_initOptionsInGame_4ADAD0();
int sub_4ADEF0(uint32_t* a1, int a2);
int nox_xxx_windowOptionsProc_4ADF30(int a1, int a2, int* a3, int a4);
int sub_4AE3B0();
int sub_4AE3D0();
void sub_4AE6F0(int a1, int a2, int a3, int a4, int a5);
long long sub_4AEE30();
void nox_client_drawPoint_4B0BC0(int a1, int a2, int a3);
int sub_4B4860(int a1, int a2, int a3, int a4);
int nox_xxx_wndScrollBoxDraw_4B4BA0(int a1, int a2, unsigned int a3, int a4);
nox_window* nox_gui_newSlider_4B4EE0(int a1, int a2, int a3, int a4, int a5, int a6, uint32_t* a7, float* a8);
int sub_4B5010(int a1, unsigned int a2, int a3, int a4);
int sub_4B51A0(int a1);
int sub_4B51E0(int a1, int a2);
int sub_4B52C0(int a1, int a2);
int nox_xxx_wndScrollBoxProc_4B5320(int a1, unsigned int a2, int a3, unsigned int a4);
int nox_xxx_wndScrollBoxSetAllFn_4B5500(int a1);
int nox_xxx_wndScrollBoxDraw_4B5540(int a1, int a2);
int nox_xxx_wndScrollBoxDraw_4B5620(uint32_t* a1, int a2);
int nox_xxx_wndScrollBoxButtonCreate_4B5640(int a1, int a2, int a3);
void sub_4B5700(nox_window* a1, void* a2, void* a3, void* a4, void* a5, void* a6);
int sub_4B5CD0();
int sub_4B63B0(int2* a1, int2* a2);
int sub_4B64C0();
void sub_4B6720(int2* a1, int a2, int a3, char a4);
int sub_4B6880(uint32_t* a1, int a2, int a3, int a4);
int sub_4B6970(uint32_t* a1, nox_drawable* dr, int a3, int a4);
short sub_4B69F0(int a1);
int sub_4B6B80(int* a1, nox_drawable* dr, int a3);
uint32_t* nox_xxx_netHandleSummonPacket_4B7C40(short a1, unsigned short* a2, unsigned short a3, unsigned char a4,
											   short a5);
void sub_4B7EE0(short a1);
int nox_xxx_spriteShieldLoad_4B7F90();
uint32_t* nox_xxx_fxShield_4B8090(unsigned int a1, int a2);
void nox_xxx_spriteScanForShield_4B81E0(int a1, int a2);
uint32_t* sub_4B8E10(uint32_t* a1, char* a2);

#endif // NOX_PORT_GAME3
