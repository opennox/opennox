#ifndef NOX_PORT_GAME2_1
#define NOX_PORT_GAME2_1

#include "common__savegame.h"
#include "defs.h"

int sub_460D40();
int sub_460D50();
int nox_xxx_cliPrepareGameplay1_460E60();
int sub_460EA0(int a1);
void sub_460EB0(int a1, char a2);
void sub_461010();
void sub_461060();
char* sub_461090(int a1, int a2);
char* sub_4610D0(unsigned char a1);
char* sub_461120(int a1, int a2);
int sub_461160(int a1);
int sub_4611A0();
int sub_4611B0();
void nox_xxx_netAbilityRewardCli_4611E0(int a1, int a2, char* a3);
int nox_xxx_buttonFindFirstEmptySlot_461250();
int sub_4612A0();
int nox_xxx_buttonHaveSpellInBarMB_4612D0(int a1);
void nox_xxx_buttonSetImgMB_461320(int a1, uint32_t* a2);
int sub_461360(int a1);
int sub_461400();
int sub_461440(int a1);
int sub_461450();
void nox_xxx_playerInitColors_461460(nox_playerInfo* pl);
char* sub_461520();
int nox_xxx_clientSetAltWeapon_461550(int a1);
int sub_4615C0();
int sub_461600(int a1);
int nox_xxx_send2ServInvenFail_461630(short a1);
int sub_461930();
int* sub_461970(int a1, int a2);
char* sub_4619F0();
unsigned char* sub_461B50();
uint64_t** sub_461E60(uint64_t*** a1);
char* sub_461EF0(int a1);
int sub_461F90(int a1);
int sub_4622E0(int a1);
int nox_xxx_clientEquip_4623B0(int a1);
uint32_t* sub_4623E0(uint32_t* a1, int a2);
int sub_4624D0(int a1);
int sub_4625D0(uint32_t* a1);
double sub_4626C0(int a1);
double sub_462700(int a1);
int sub_463370(uint32_t* a1, nox_point* pos, uint32_t* a3);
int sub_4633B0(int a1, float* a2, float* a3);
int sub_463420(int a1);
int nox_xxx_inventoryDrawAllMB_463430(int a1);
int nox_xxx_guiDrawInventoryTray_4643B0(int a1, int a2);
int sub_464770(int a1, int a2, unsigned int a3);
int sub_464B40(int a1, int a2);
int nox_xxx_clientDequip_464B70(int a1);
int nox_xxx_XorEaxEaxSub_464BA0();
int nox_xxx_inventoryWndProc_464BB0(int a1, int a2);
int nox_xxx_trade_4657B0(short a1);
char nox_xxx_clientTradeMB_4657E0(uint32_t* a1);
int nox_xxx_clientTrade_465870(short a1);
void sub_4658A0(int a1, int2* a2);
int sub_465990(uint32_t* a1);
int nox_xxx_clientDrop_465BE0(int2* a1);
int nox_xxx_clientKeyEquip_465C30(int a1, int a2);
void nox_xxx_clientUse_465C70(int a1);
int sub_465CA0();
void sub_465CD0(uint32_t* a1, int a2, int a3, int a4);
int sub_465D50_draw(int a1);
int sub_465DE0(int a1);
int nox_xxx_wndCreateInventoryMB_465E00();
int nox_xxx_movEax1Sub_4661C0();
int sub_466220(int a1, int a2, int* a3, int a4);
int sub_466550(int a1, unsigned int a2);
int nox_xxx_inventoryDrawProc_466580(uint32_t* a1);
int sub_466620(int a1, int a2, unsigned int a3);
int sub_466950(int a1);
int sub_466BA0(uint32_t* a1, int a2, unsigned int a3, int a4);
int sub_466BF0(int a1, int a2, unsigned int a3, int a4);
int sub_466C40(int a1);
int sub_466ED0(int a1);
int sub_466F50(uint32_t* a1, int* a2);
char* nox_xxx_inventoryLoadImages_467050();
void nox_client_invAlterWeapon_4672C0();
int sub_4673F0(int a1, int a2);
int sub_467410(int a1);
char sub_467420(char a1);
unsigned char sub_467430();
int sub_467440(int a1);
int sub_467450(int a1);
int sub_467470(int a1, float a2);
int sub_467490(int a1);
int sub_4674A0();
void nox_window_set_visible_unk5(int visible);
void nox_xxx_cliUseCurePoison_4674E0(int a1);
char* nox_xxx_cliInventoryFirstItemByTT_467520(int a1);
int sub_467590();
int sub_4675B0();
short sub_4675E0(int a1, short a2, short a3);
int sub_467650();
void sub_467680();
nox_window* nox_xxx_wndGetHandle_4676A0();
int sub_4676D0(int a1);
int sub_467700(int a1);
int sub_467740(int a1);
int sub_467810(int a1, int a2);
int sub_467850(int a1);
char* sub_467870(int a1, int a2);
int sub_4678B0();
int sub_4678C0();
int sub_4678D0();
char* sub_467930(int a1, int a2, int a3);
int sub_467980();
int sub_467B00(int a1, int a2);
int sub_467BB0();
int sub_467C10();
int nox_client_toggleInventory_467C60();
int sub_467C80();
int sub_467CD0();
int nox_xxx_gameClearAll_467DF0(int a1);
char* sub_469920(nox_point* a1);
int sub_469B90(int* a1);
char* nox_xxx_getAmbientColor_469BB0();
int sub_469FA0();
void* nox_xxx_getWallSprite_46A3B0(int a1, int a2, int a3, int a4);
void nox_xxx_getWallDrawOffset_46A3F0(int a1, int a2, int a3, int a4, int* px, int* py);
void nox_client_chatStart_46A430(int a1);
int sub_46A4A0();
size_t nox_xxx_cmdSayDo_46A4B0(wchar2_t* a1, int a2);
int sub_46A5D0(uint32_t* a1, int a2);
int sub_46A6A0();
uint32_t* sub_46A730();
int sub_46A7E0(uint32_t* a1, int a2, int a3, int a4);
int sub_46A820(int a1, int a2, int a3, int a4);
int sub_46A860();
int nox_xxx_wndRetNULL_46A8A0();
int nox_xxx_wndRetNULL_0_46A8B0();
int sub_46AE10(int a1, int a2);
int nox_xxx_wndSetOffsetMB_46AE40(int a1, int a2, int a3);
int nox_xxx_wndSetIcon_46AE60(int a1, int a2);
int nox_xxx_wndSetIconLit_46AEA0(int a1, int a2);
int sub_46AEC0(int a1, int a2);
int sub_46AEE0(int a1, int a2);
wchar2_t* sub_46AF00(void* a1);
void* sub_46AF40(void* a1);
int nox_gui_windowCopyDrawData_46AF80(nox_window* win, const void* p);
nox_window* sub_46B630(nox_window* a1p, int a2, int a3);
int nox_xxx_wnd_46C2A0(nox_window* a1p);
nox_window* nox_client_getWin1064916_46C720();
int sub_46D6F0();
int sub_46DB80();
int sub_46DC00(int a1, unsigned char a2, int a3);
int sub_46DC30(int a1, unsigned char a2, wchar2_t* a3, ...);
char* sub_46DCC0();
int sub_46E080(int a1);
int sub_46E130(int a1);
unsigned short* sub_46E170(wchar2_t* a1);
int sub_46E1E0(int a1);
char* sub_46E4E0();
int sub_46F060();
int nox_xxx_Proc_46F070();
void sub_46FAE0();
unsigned char sub_46FE60(int a1);
unsigned char sub_46FEB0(unsigned char a1);
char sub_46FEE0();
char sub_46FF70(int a1);
unsigned char sub_46FFD0();
void sub_4703F0();
void sub_470510();
int sub_470580();
void sub_4705B0();
char sub_4705F0(char a1, char a2, short a3);
char sub_470650(char a1, short a2);
int sub_470680();
int sub_4706A0();
int nox_xxx_playerGet_470A90();
void nox_xxx_cliShowHideTubes_470AA0(int a1);
unsigned char* nox_xxx_guiHealthManaColorInit_470B00();
int sub_470C40(int a1);
int nox_xxx_cliSetTotalHealth_470C80(int a1, int a2);
int sub_470CB0(int a1);
int sub_470CC0();
int sub_470CD0();
int nox_xxx_cliSetManaAndMax_470CE0(int a1, int a2);
int nox_xxx_cliSetMana_470D10(int a1);
int sub_470D20(int a1, int a2);
void sub_470D70();
int sub_470D90(int a1, int a2);
int nox_xxx_cliGetMana_470DD0();
int sub_470DE0();
int sub_470E90(int a1, int a2);
void nox_win_init_cur_weapon(nox_window* a1, int a2, int a3, int w, int h);
int sub_470F40_draw(nox_window* win);
int sub_471250(uint32_t* a1);
int sub_471450(uint32_t* a1);
int nox_xxx_guiBottleSlotDrawFn_471A80(uint32_t* a1);
int nox_xxx_guiBottleSlotProc_471B90(int a1, int a2);
int nox_xxx_drawHealthManaBar_471C00(int a1);
int sub_472080();
int sub_4720C0(int xLeft, int a2);
int nox_xxx_guiHealthManaTubeProc_472100(int a1, int a2);
int sub_4721A0(int a1);
int nox_xxx_cliPrepareGameplay2_4721D0();
void nox_client_quickHealthPotion_472220();
void nox_client_quickManaPotion_472240();
void nox_client_quickCurePoisonPotion_472260();
wchar2_t* sub_472280();
unsigned char* sub_472310();
void nox_client_mapZoomIn_4724E0();
void nox_client_mapZoomOut_472500();
int nox_xxx_cliSetMinimapZoom_472520(int a1);
int sub_472540(int a1);
void nox_xxx_drawMinimap4Sprite_4725C0(int a1);
int nox_xxx_cliDrawMinimap_472600(int a1, int a2);
int sub_4730D0(int2* a1, unsigned char a2, int a3);
int sub_473380(int a1, int a2, int a3, int a4, int a5);
int sub_4733B0(uint32_t* a1);
int sub_473420(uint32_t* a1);
void nox_video_drawCircleRad3_4734F0(int* a1);
int nox_client_drawRectLines_473510(int a1, int a2, int a3, int a4);
void nox_xxx_minimapDrawPoint_473570(int xLeft, int yTop);
void sub_4735C0(int xLeft, int yTop);
char nox_client_toggleMap_473610();
int sub_473670();
int nox_xxx_drawMinimapAndLines_4738E0();
void nox_xxx____setargv_11_473920();
char* sub_473930();
int sub_473960();
void sub_473970(int2* a1, int2* a2);
int sub_4739E0(uint32_t* a1, int2* a2, int2* a3);
int sub_473A10(uint32_t* a1, int2* a2, uint32_t* a3);
void nox_xxx_drawWalls_473C10(nox_draw_viewport_t* vp, void* data);
int sub_474B40(nox_drawable* dr);
void nox_xxx_drawList1096512_Append_4754C0(void* a1);
int nox_xxx_sprite_4756E0_drawable(nox_drawable* dr);
int nox_xxx_sprite_475740_drawable(nox_drawable* dr);
int nox_xxx_sprite_4757A0_drawable(nox_drawable* dr);
int sub_4757D0_drawable(nox_drawable* dr);

#endif // NOX_PORT_GAME2_1
