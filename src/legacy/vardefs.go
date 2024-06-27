package legacy

/*
#include "defs.h"
#include "server__script__internal.h"
extern uint32_t dword_5d4594_1046492;
extern void* nox_xxx_aClosewoodengat_587000_133480;
extern nox_window* nox_wol_wnd_world_814980;
extern uint32_t dword_5d4594_1045516;
extern uint32_t dword_5d4594_1045684;
extern uint32_t dword_5d4594_1307784;
extern uint32_t dword_5d4594_2650652;
extern uint32_t dword_5d4594_1320940;
extern nox_window* dword_5d4594_815004;
extern uint32_t nox_wol_wnd_gameList_815012;
extern uint32_t dword_5d4594_1062480;
extern uint32_t dword_5d4594_832484;
extern uint32_t dword_5d4594_2516348;
extern uint32_t dword_5d4594_1063116;
extern uint32_t dword_5d4594_1090120;
extern uint32_t dword_5d4594_2516328;
extern uint32_t dword_5d4594_1049504;
extern uint32_t dword_5d4594_1046868;
extern uint32_t dword_5d4594_1193380;
extern uint32_t dword_5d4594_1316972;
extern void* dword_5d4594_1123524;
extern uint32_t dword_5d4594_1045604;
extern uint32_t dword_5d4594_529316;
extern uint32_t dword_5d4594_1309820;
extern uint32_t dword_5d4594_1062508;
extern uint32_t dword_5d4594_1049864;
extern uint32_t dword_5d4594_1062512;
extern uint32_t dword_5d4594_1599656;
extern uint32_t dword_5d4594_1049500;
extern uint32_t dword_5d4594_3798840;
extern uint32_t dword_5d4594_1309812;
extern void* dword_587000_155144;
extern uint32_t dword_5d4594_1047516;
extern uint32_t dword_5d4594_1308084;
extern uint32_t dword_5d4594_3798824;
extern uint32_t dword_5d4594_3798820;
extern uint32_t dword_5d4594_1049508;
extern uint32_t dword_5d4594_1049796_inventory_click_column_index;
extern uint32_t dword_5d4594_1098576;
extern uint32_t dword_5d4594_1321040;
extern uint32_t dword_5d4594_1548476;
extern uint32_t dword_5d4594_2660652;
extern uint32_t dword_5d4594_1045636;
extern uint32_t dword_5d4594_3798804;
extern uint32_t dword_5d4594_832476;
extern uint32_t dword_5d4594_1570272;
extern uint32_t dword_5d4594_1305684;
extern uint32_t dword_5d4594_1321228;
extern uint32_t nox_game_createOrJoin_815048;
extern uint32_t dword_5d4594_1063636;
extern uint32_t nox_xxx_polygonNextIdx_587000_60352;
extern uint32_t dword_5d4594_1062456;
extern uint32_t dword_5d4594_816348;
extern uint32_t dword_5d4594_1098628;
extern uint32_t dword_5d4594_529340;
extern uint32_t nox_xxx_minimap_587000_149232;
extern uint32_t dword_5d4594_1046872;
extern uint32_t dword_5d4594_1046496;
extern uint32_t dword_5d4594_1523028;
extern uint32_t dword_5d4594_1316708;
extern uint32_t dword_5d4594_251572;
extern uint32_t dword_5d4594_1303452;
extern uint32_t dword_5d4594_2386224;
extern uint32_t dword_5d4594_815016;
extern uint32_t dword_5d4594_1197312;
extern uint32_t dword_5d4594_831236;
extern uint32_t dword_5d4594_1047540;
extern uint32_t dword_5d4594_1320964;
extern uint32_t dword_5d4594_1049800_inventory_click_row_index;
extern uint32_t dword_5d4594_1309720;
extern uint32_t dword_5d4594_1046936;
extern uint32_t nox_xxx_lightningTarget_5d4594_2487908;
extern uint32_t dword_5d4594_1107036;
extern uint32_t dword_5d4594_3798836;
extern uint32_t dword_5d4594_1565512;
extern uint32_t dword_5d4594_815028;
extern uint32_t dword_5d4594_1301848;
extern uint32_t dword_5d4594_1316712;
extern uint32_t dword_5d4594_1321232;
extern uint32_t dword_5d4594_815024;
extern uint32_t dword_5d4594_1045528;
extern uint32_t dword_5d4594_815020;
extern uint32_t dword_5d4594_1305680;
extern uint32_t dword_5d4594_815032;
extern uint32_t dword_5d4594_1193360;
extern uint32_t dword_5d4594_1049520;
extern uint32_t dword_5d4594_1309748;
extern uint32_t dword_5d4594_1045464;
extern uint32_t dword_5d4594_1045484;
extern uint32_t dword_5d4594_1049856;
extern uint32_t dword_5d4594_1064856;
extern uint32_t nox_xxx_aNox_cfg_0_587000_132132;
extern uint32_t dword_5d4594_1193348;
extern uint32_t dword_5d4594_1047520;
extern uint32_t dword_5d4594_534808;
extern uint32_t dword_5d4594_1197324;
extern uint32_t dword_5d4594_1599576;
extern uint32_t dword_5d4594_1599596;
extern uint32_t dword_5d4594_1045480;
extern uint32_t dword_5d4594_2487556;
extern uint32_t dword_5d4594_1523024;
extern uint32_t dword_5d4594_2516344;
extern uint32_t dword_5d4594_1046508;
extern uint32_t dword_5d4594_1193384;
extern uint32_t dword_5d4594_1045536;
extern uint32_t nox_wol_servers_sorting_166704;
extern uint32_t dword_5d4594_371692;
extern uint32_t dword_5d4594_3798832;
extern uint32_t dword_5d4594_1197320;
extern void* nox_alloc_pixelSpan_1301844;
extern uint32_t dword_5d4594_3798808;
extern uint32_t dword_5d4594_1062524;
extern uint32_t dword_5d4594_3835396;
extern uint32_t dword_5d4594_1062528;
extern uint32_t dword_5d4594_1522604;
extern uint32_t dword_5d4594_3798816;
extern uint32_t dword_5d4594_1045532;
extern uint32_t dword_5d4594_528256;
extern uint32_t dword_5d4594_1568024;
extern uint32_t dword_5d4594_1308088;
extern uint32_t dword_5d4594_826028;
extern uint32_t dword_5d4594_2487524;
extern uint32_t dword_5d4594_1522612;
extern uint32_t dword_5d4594_1045468;
extern uint32_t nox_xxx_lightningTargetArrayIndex_5d4594_2487904;
extern void* dword_587000_127004;
extern uint32_t nox_wol_server_result_cnt_815088;
extern uint32_t dword_5d4594_1046924;
extern uint32_t nox_xxx_polygonNextAngle_587000_60356;
extern uint32_t dword_5d4594_1307716;
extern uint32_t dword_587000_93156;
extern uint32_t dword_587000_93160;
extern uint32_t dword_5d4594_2523780;
extern uint32_t dword_5d4594_2650676;
extern uint32_t dword_5d4594_1321252;
extern uint32_t dword_5d4594_1522632;
extern uint32_t dword_5d4594_531648;
extern uint32_t dword_5d4594_1045432;
extern uint32_t dword_5d4594_815044;
extern uint32_t nox_wnd_briefing_831232;
extern uint32_t dword_5d4594_2487712;
extern void* nox_alloc_tradeItems_2386496;
extern uint32_t dword_5d4594_2386212;
extern uint32_t dword_5d4594_1319060;
extern uint32_t dword_5d4594_1047548;
extern uint32_t dword_5d4594_816340;
extern uint32_t dword_5d4594_1049512;
extern uint32_t dword_5d4594_1047932;
extern uint32_t dword_5d4594_816368;
extern uint32_t dword_5d4594_1062476;
extern uint32_t dword_5d4594_1090276;
extern uint32_t dword_5d4594_1049496;
extern void* dword_587000_93164;
extern uint32_t dword_5d4594_1098624;
extern uint32_t dword_5d4594_1569672;
extern uint32_t dword_5d4594_1197316;
extern uint32_t dword_5d4594_2386576;
extern uint32_t dword_5d4594_1320968;
extern uint32_t dword_5d4594_2488604;
extern uint64_t qword_581450_9544;
extern uint64_t qword_5d4594_1567940;
extern uint64_t qword_581450_10256;
extern uint64_t qword_581450_9552;
extern uint64_t qword_581450_9512;
extern uint64_t qword_5d4594_815068;
extern uint64_t qword_5d4594_814956;
extern uint64_t qword_581450_10176;
extern uint32_t dword_5d4594_1316408;
extern uint32_t dword_5d4594_2488660;
extern uint32_t dword_587000_292492;
extern uint32_t dword_5d4594_1309732;
extern uint32_t dword_5d4594_1321204;
extern uint32_t dword_5d4594_2386500;
extern uint32_t dword_5d4594_2649712;
extern uint32_t dword_5d4594_1046532;
extern uint32_t dword_5d4594_3798828;
extern uint32_t dword_5d4594_1309728;
extern uint32_t dword_5d4594_1309836;
extern uint32_t dword_5d4594_1320992;
extern uint32_t dword_5d4594_1047552;
extern uint32_t dword_587000_60044;
extern void* dword_5d4594_251560;
extern uint32_t dword_5d4594_3484;
extern void* nox_alloc_screenParticles_806044;
extern uint32_t dword_5d4594_528252;
extern uint32_t dword_5d4594_1309828;
extern void* dword_587000_122852;
extern uint32_t dword_5d4594_1301776;
extern uint32_t nox_xxx_screenWidth_587000_184452;
extern uint32_t dword_5d4594_1550916;
extern uint32_t dword_5d4594_2487540;
extern void* dword_587000_81128;
extern uint32_t dword_5d4594_3798800;
extern uint32_t dword_5d4594_3798812;
extern uint32_t dword_5d4594_1046952;
extern uint32_t dword_5d4594_832536;
extern uint32_t dword_587000_87408;
extern uint32_t dword_5d4594_814988;
extern uint32_t dword_5d4594_2487560;
extern uint32_t dword_5d4594_1062516;
extern uint32_t dword_5d4594_1316492;
extern uint32_t dword_5d4594_2491616;
extern uint32_t dword_5d4594_251744;
extern uint32_t dword_5d4594_2487248;
extern uint32_t dword_5d4594_1046516;

extern uint32_t dword_5d4594_2487532;
extern uint32_t nox_server_connectionType_3596;
extern void* nox_monsterBin_head_2386924;
extern uint32_t dword_5d4594_1200804;
extern uint32_t dword_5d4594_1049484;
extern uint32_t dword_5d4594_1321216;
extern uint32_t dword_5d4594_1193712;
extern uint32_t dword_5d4594_1319232;
extern uint32_t dword_5d4594_1046656;
extern uint32_t dword_5d4594_832532;
extern uint32_t dword_5d4594_2650680;
extern uint32_t dword_5d4594_588120;
extern uint32_t dword_5d4594_1047512;
extern uint32_t dword_5d4594_588084;
extern uint32_t dword_5d4594_1321044;
extern uint32_t dword_5d4594_816372;
extern uint32_t dword_5d4594_832492;
extern uint32_t dword_5d4594_1046504;
extern uint32_t dword_5d4594_1045508;
extern uint32_t dword_5d4594_825744;
extern uint32_t dword_5d4594_1319264;
extern uint32_t dword_5d4594_1046928;
extern uint32_t dword_5d4594_1098580;
extern uint32_t dword_5d4594_831224;
extern uint32_t nox_xxx_energyBoltTarget_5d4594_2487880;
extern uint32_t dword_5d4594_832504;
extern uint32_t dword_5d4594_1203864;
extern uint32_t dword_5d4594_832508;
extern uint32_t dword_5d4594_3835356;
extern uint32_t nox_client_renderBubbles_80844;
extern void* nox_alloc_chat_1197364;
extern uint32_t dword_5d4594_1046520;
extern uint32_t dword_5d4594_832516;
extern uint32_t nox_xxx_aNox_cfg_0_587000_132136;
extern uint32_t dword_5d4594_832496;
extern uint32_t dword_5d4594_832512;
extern uint32_t dword_5d4594_832524;
extern uint32_t dword_5d4594_823776;
extern uint32_t dword_5d4594_1045688;
extern uint32_t dword_5d4594_832528;
extern uint32_t dword_5d4594_832500;
extern uint32_t dword_5d4594_832520;
extern uint32_t dword_5d4594_1307736;
extern uint32_t dword_5d4594_1046636;
extern uint32_t dword_5d4594_1046500;
extern uint32_t dword_587000_136184;
extern uint32_t dword_5d4594_1098592;
extern uint32_t dword_5d4594_1321032;
extern uint32_t dword_5d4594_251720;
extern uint32_t dword_5d4594_1320932;
extern uint32_t dword_5d4594_1197332;
extern uint32_t dword_5d4594_1319056;
extern uint32_t dword_5d4594_1047528;
extern uint32_t dword_5d4594_608316;
extern uint32_t dword_5d4594_1046948;
extern uint32_t dword_5d4594_1046528;
extern uint32_t dword_5d4594_741364;
extern uint32_t dword_5d4594_1046932;
extern uint32_t dword_5d4594_1321224;
extern uint32_t dword_5d4594_588068;
extern uint32_t dword_587000_122848;
extern uint32_t dword_5d4594_1045460;
extern uint32_t dword_5d4594_831220;
extern uint32_t dword_5d4594_1203840;
extern uint32_t dword_5d4594_826032;
extern uint32_t dword_5d4594_2487708;
extern uint32_t nox_xxx_waypointCounterMB_587000_154948;
extern uint32_t dword_5d4594_1548700;
extern uint32_t dword_5d4594_741332;
extern uint32_t dword_5d4594_1565616;
extern uint32_t dword_5d4594_816356;
extern uint32_t dword_5d4594_1203836;
extern uint32_t dword_5d4594_1301780;
extern uint32_t dword_5d4594_1064864;
extern uint32_t dword_5d4594_1096252;
extern uint32_t dword_5d4594_741356;
extern uint32_t dword_5d4594_1197356;
extern uint32_t dword_5d4594_1045436;
extern uint32_t dword_5d4594_1321800;
extern uint32_t dword_5d4594_1197336;
extern uint32_t dword_5d4594_1197352;
extern void* dword_5d4594_826036;
extern uint32_t dword_5d4594_1064860;
extern uint32_t dword_5d4594_1046536;
extern uint32_t dword_5d4594_1062468;
extern uint32_t dword_5d4594_1523036;
extern uint32_t dword_5d4594_1197328;
extern uint32_t dword_5d4594_1090044;
extern uint32_t dword_5d4594_1305748;
extern void* nox_alloc_respawn_1568020;
extern void* nox_alloc_monsterList_2386220;
extern uint32_t nox_xxx_lightningClosestTargetDistance_5d4594_2487912;
extern uint32_t dword_5d4594_1567928;
extern uint32_t dword_5d4594_1307720;
extern uint32_t dword_5d4594_2386920;
extern uint32_t dword_5d4594_2386564;
extern uint32_t dword_5d4594_1550912;
extern uint32_t dword_5d4594_1062488;
extern uint32_t dword_5d4594_2516356;
extern uint32_t dword_5d4594_1047936;
extern uint32_t dword_5d4594_1321208;
extern uint32_t dword_5d4594_2487620;
extern uint32_t dword_5d4594_1309824;
extern uint32_t dword_5d4594_531656;
extern void* nox_alloc_tradeSession_2386492;
extern uint32_t dword_5d4594_1096288;
extern uint32_t dword_5d4594_3835352;
extern uint32_t dword_5d4594_1046956;
extern uint32_t dword_587000_145664;
extern uint32_t dword_5d4594_1316704;
extern uint32_t dword_5d4594_825736;
extern uint32_t dword_5d4594_1313540;
extern uint32_t dword_5d4594_1096284;
extern uint32_t dword_5d4594_2487624;
extern uint32_t dword_5d4594_1309832;
extern uint32_t dword_5d4594_2487576;
extern uint32_t dword_5d4594_1309756;
extern uint32_t dword_5d4594_1309736;
extern uint32_t nox_client_gui_flag_1556112;
extern uint32_t dword_5d4594_1301812;
extern void* nox_alloc_friendList_1203860;
extern uint32_t dword_5d4594_2660032;
extern uint32_t dword_5d4594_1090040;
extern uint32_t dword_5d4594_3835348;
extern uint32_t dword_5d4594_1045692;
extern uint32_t dword_5d4594_1046640;
extern uint32_t dword_5d4594_1320972;
extern uint32_t dword_5d4594_1565516;
extern uint32_t dword_5d4594_1096260;
extern uint32_t dword_5d4594_529336;
extern uint32_t dword_5d4594_1308124;
extern uint32_t dword_5d4594_2487652;
extern uint32_t dword_5d4594_2523776;
extern uint32_t dword_5d4594_3835388;
extern uint32_t dword_5d4594_1046648;
extern uint32_t dword_5d4594_831276;
extern uint32_t dword_5d4594_1301816;
extern uint32_t dword_5d4594_2491552;
extern uint32_t dword_5d4594_1045588;
extern uint32_t dword_5d4594_1599616;
extern uint32_t dword_5d4594_1319236;
extern uint32_t dword_5d4594_815056;
extern uint32_t dword_5d4594_1308128;
extern uint32_t nox_client_connError_814552;
extern uint32_t dword_5d4594_10984;
extern uint32_t dword_587000_145668;
extern uint32_t dword_587000_180480;
extern uint32_t dword_587000_230092;
extern uint32_t dword_5d4594_1197372;
extern uint32_t dword_5d4594_251716;
extern uint32_t dword_5d4594_251708;
extern uint32_t dword_5d4594_1063120;
extern uint32_t dword_5d4594_1046540;
extern uint32_t dword_5d4594_1062496;
extern void* nox_alloc_vote_1599652;
extern uint32_t dword_5d4594_1049804;
extern uint32_t dword_5d4594_1090112;
extern uint32_t dword_5d4594_3807140;
extern uint32_t dword_5d4594_2487676;
extern uint32_t dword_5d4594_1049696;
extern uint32_t dword_5d4594_1316476;
extern uint32_t dword_5d4594_2488652;
extern void* nox_alloc_spawn_2386216;
extern uint32_t dword_5d4594_2487672;
extern void* nox_alloc_magicEnt_1569668;
extern uint32_t dword_5d4594_814992;
extern uint32_t dword_587000_66116;
extern uint32_t dword_5d4594_2487564;
extern uint32_t dword_5d4594_3804684;
extern uint32_t dword_5d4594_1308120;
extern uint32_t dword_5d4594_2488608;
extern uint32_t dword_5d4594_1568868;
extern uint32_t dword_5d4594_1090108;
extern uint32_t dword_5d4594_1308108;
extern uint32_t dword_5d4594_3835312;
extern uint32_t dword_5d4594_1197308;
extern uint32_t dword_5d4594_1045584;
extern uint32_t dword_5d4594_814548;
extern uint32_t dword_5d4594_1308132;
extern uint32_t dword_5d4594_2386228;
extern uint32_t dword_5d4594_1308100;
extern uint32_t nox_server_sanctuaryHelp_54276;
extern uint32_t dword_5d4594_1045424;
extern void* nox_alloc_pendingOwn_2386916;
extern uint32_t dword_5d4594_1313740;
extern uint32_t dword_587000_237036;
extern uint32_t dword_5d4594_1062492;
extern uint32_t dword_5d4594_1098604;
extern uint32_t nox_server_needInitNetCodeCache;
extern uint32_t dword_5d4594_1308116;
extern uint32_t dword_5d4594_1319248;
extern uint32_t dword_5d4594_3807136;
extern uint32_t dword_5d4594_1046944;
extern uint32_t dword_5d4594_2487584;
extern uint32_t dword_5d4594_251712;
extern uint32_t dword_5d4594_531652;
extern uint32_t dword_5d4594_2487948;
extern void* nox_alloc_healthChange_1301772;
extern uint32_t dword_5d4594_1049536;
extern uint32_t dword_5d4594_1098616;
extern uint32_t dword_5d4594_3807152;
extern uint32_t dword_5d4594_1049996;
extern uint32_t dword_5d4594_1305788;
extern uint32_t dword_5d4594_1098600;
extern uint32_t dword_5d4594_1064868;
extern uint32_t dword_5d4594_1565520;
extern uint32_t dword_5d4594_1320988;
extern uint32_t dword_5d4594_1599696;
extern uint32_t dword_587000_261388;
extern uint32_t dword_5d4594_1200768;
extern uint32_t dword_5d4594_1523032;
extern uint32_t dword_5d4594_831256;
extern uint32_t dword_5d4594_1316436;
extern uint32_t dword_587000_145672;
extern uint32_t dword_5d4594_1046864;
extern uint32_t dword_5d4594_1523044;
extern uint32_t dword_5d4594_1046356;
extern uint32_t dword_5d4594_3799452;
extern uint32_t dword_5d4594_1096280;
extern uint32_t dword_5d4594_1316472;
extern uint32_t dword_5d4594_1062560;
extern uint32_t dword_5d4594_2491544;
extern uint32_t dword_587000_183456;
extern uint32_t nox_xxx_lightningOwner_5d4594_2487900;
extern uint32_t dword_5d4594_1096256;
extern uint32_t dword_5d4594_1046524;
extern uint32_t dword_5d4594_2523764;
extern uint32_t dword_5d4594_2487244;
extern uint32_t dword_5d4594_1090280;
extern uint32_t dword_5d4594_1200796;
extern uint32_t dword_5d4594_1313536;
extern uint32_t dword_5d4594_1308152;
extern uint32_t dword_5d4594_1098596;
extern uint32_t dword_5d4594_1047524;
extern uint32_t dword_5d4594_1563096;
extern uint32_t dword_5d4594_1062564;
extern uint32_t dword_5d4594_1308144;
extern uint32_t dword_5d4594_1049524;
extern uint32_t dword_5d4594_1316484;
extern uint32_t dword_5d4594_1320936;
extern uint32_t dword_5d4594_1045552;
extern uint32_t dword_587000_234176;
extern uint32_t dword_5d4594_1096264;
extern uint32_t dword_5d4594_831076;
extern uint32_t dword_5d4594_1321024;
extern uint32_t dword_5d4594_1523048;
extern uint32_t dword_5d4594_1308140;
extern uint32_t dword_5d4594_1203832;
extern uint32_t dword_5d4594_831244;
extern uint32_t dword_5d4594_2516372;
extern uint32_t dword_5d4594_1565632;
extern uint32_t dword_5d4594_1045520;
extern uint32_t dword_5d4594_1319268;
extern uint32_t dword_5d4594_1321520;
extern uint32_t dword_5d4594_526276;
extern uint32_t dword_5d4594_1062556;
extern uint32_t dword_5d4594_1049992;
extern uint32_t dword_587000_171388;
extern uint32_t dword_5d4594_2489460;
extern uint32_t dword_5d4594_816412;
extern uint32_t dword_5d4594_1308148;
extern uint32_t dword_5d4594_1062484;
extern uint32_t dword_5d4594_528260;
extern uint32_t dword_5d4594_1090284;
extern uint32_t dword_5d4594_1301796;
extern uint32_t dword_5d4594_2487932;
extern uint32_t dword_5d4594_2487628;
extern uint32_t dword_5d4594_1045544;
extern uint32_t dword_5d4594_1548480;
extern uint32_t nox_xxx_lightningSteps_587000_178216;
extern uint32_t dword_5d4594_1045428;
extern uint32_t dword_5d4594_2487536;
extern uint32_t dword_5d4594_2523804;
extern uint32_t dword_587000_126996;
extern uint32_t dword_5d4594_2487568;
extern uint32_t dword_5d4594_1045540;
extern uint32_t dword_5d4594_2487580;
extern uint32_t dword_5d4594_1556856;
extern uint32_t dword_5d4594_1046360;
extern uint32_t dword_5d4594_1049976;
extern uint32_t dword_5d4594_1193188;
extern uint32_t dword_5d4594_1565628;
extern uint32_t dword_5d4594_816444;
extern uint32_t nox_server_sendMotd_108752;
extern uint32_t dword_5d4594_1096276;
extern uint32_t dword_5d4594_3835392;
extern uint32_t dword_5d4594_1123520;
extern uint32_t dword_5d4594_2488720;
extern uint32_t dword_5d4594_1567988;
extern uint32_t dword_5d4594_1308112;
extern uint32_t dword_5d4594_832480;
extern uint32_t dword_5d4594_1568308;
extern uint32_t dword_587000_129656;
extern uint32_t dword_5d4594_1091364;
extern uint32_t dword_5d4594_599496;
extern uint32_t dword_5d4594_2516380;
extern uint32_t dword_5d4594_1599644;
extern uint32_t dword_5d4594_1563276;
extern uint32_t nox_server_resetQuestMinVotes_229988;
extern uint32_t dword_5d4594_1062520;
extern uint32_t dword_5d4594_831260;
extern uint32_t dword_5d4594_1049516;
extern uint32_t dword_5d4594_1090292;
extern uint32_t dword_5d4594_527660;
extern uint32_t dword_5d4594_2489160;
extern uint32_t dword_5d4594_1522968;
extern uint32_t dword_5d4594_2490508;
extern uint32_t dword_5d4594_1316456;
extern uint32_t dword_5d4594_1045640;
extern uint32_t dword_5d4594_528264;
extern uint32_t dword_5d4594_1098620;
extern uint32_t dword_5d4594_1200776;
extern uint32_t dword_5d4594_1045548;
extern uint32_t dword_5d4594_1316412;
extern uint32_t dword_5d4594_3807116;
extern uint32_t dword_5d4594_815100;
extern uint32_t dword_5d4594_1321036;
extern uint32_t dword_5d4594_1049808;
extern uint32_t dword_5d4594_1568300;
extern uint32_t dword_5d4594_1563320;
extern uint32_t dword_5d4594_1313692;
extern uint32_t dword_5d4594_1599688;
extern uint32_t dword_5d4594_1307724;
extern uint32_t nox_server_kickQuestPlayerMinVotes_229992;
extern uint32_t dword_5d4594_1523040;
extern uint32_t dword_5d4594_1096272;
extern uint32_t dword_5d4594_1096636;
extern uint32_t dword_5d4594_1522956;
extern uint32_t dword_587000_183460;
extern uint32_t dword_5d4594_1303508;
extern uint32_t dword_5d4594_1568288;
extern uint32_t dword_5d4594_2491704;
extern uint32_t dword_5d4594_1320948;
extern uint32_t dword_5d4594_1568280;
extern uint32_t dword_5d4594_1556144;
extern uint32_t dword_5d4594_1050008;
extern uint32_t dword_5d4594_3835372;
extern uint32_t dword_5d4594_1556136;
extern uint32_t dword_5d4594_527656;
extern uint32_t dword_5d4594_741652;
extern uint32_t dword_5d4594_1308096;
extern uint32_t dword_5d4594_600116;
extern uint32_t dword_5d4594_2487884;
extern uint32_t dword_5d4594_1313564;
extern uint32_t dword_5d4594_1568028;
extern uint32_t dword_5d4594_815052;
extern uint32_t dword_5d4594_1548524;
extern uint32_t dword_5d4594_2386560;
extern uint32_t dword_5d4594_741648;
extern uint32_t dword_5d4594_1556128;
extern uint32_t dword_5d4594_2386548;
extern uint32_t dword_5d4594_2491592;
extern uint32_t dword_5d4594_1090828;
extern uint32_t dword_5d4594_815096;
extern uint32_t dword_5d4594_529332;
extern uint32_t dword_5d4594_1313816;
extern uint32_t dword_5d4594_1193352;
extern uint32_t dword_5d4594_2488724;
extern void* dword_5d4594_1548532;
extern uint32_t dword_5d4594_1045580;
extern uint32_t dword_5d4594_1313532;
extern uint32_t dword_5d4594_805820;
extern uint32_t dword_5d4594_1308104;
extern uint32_t dword_5d4594_2489436;
extern uint32_t dword_5d4594_1320944;
extern uint32_t nox_xxx_normalWndBits_587000_172880;
extern uint32_t dword_5d4594_2491588;
extern uint32_t dword_5d4594_1316448;
extern uint32_t dword_5d4594_1313880;
extern uint32_t dword_5d4594_2491676;
extern uint32_t dword_5d4594_1308136;
extern uint32_t dword_5d4594_831240;
extern uint32_t dword_5d4594_1301808;
extern uint32_t dword_5d4594_1599708;
extern uint32_t dword_587000_292488;
extern uint32_t dword_5d4594_1321196;
extern uint32_t dword_5d4594_815748;
extern uint32_t dword_5d4594_1045556;
extern uint32_t dword_5d4594_1045420;
extern uint32_t dword_5d4594_739392;
extern uint32_t dword_5d4594_2488728;
extern uint32_t dword_5d4594_1567960;
extern uint32_t dword_5d4594_1049844;
extern uint32_t dword_5d4594_3835364;
extern uint32_t dword_5d4594_1313788;
extern uint32_t dword_5d4594_1304328;
extern uint32_t dword_5d4594_1047536;
extern uint32_t dword_5d4594_2516352;
extern uint32_t dword_5d4594_1549844;
extern uint32_t dword_5d4594_1313804;
extern uint32_t nox_xxx_respawnAllow_587000_205200;
extern uint32_t dword_5d4594_1046652;
extern uint32_t dword_5d4594_2487632;
extern uint32_t dword_5d4594_3835360;
extern uint32_t dword_5d4594_1098456;
extern uint32_t dword_5d4594_2487804;
extern uint32_t dword_5d4594_1062552;
extern uint32_t nox_xxx_aSpellphoneme_3_587000_123008;
extern uint32_t dword_5d4594_811904;
extern uint32_t dword_5d4594_2491580;
extern uint32_t dword_5d4594_1200832;
extern uint32_t dword_5d4594_1319260;
extern void* nox_alloc_hit_2491548;
extern uint32_t dword_5d4594_2488656;
extern uint32_t dword_5d4594_2386552;
extern uint32_t dword_5d4594_3835368;
extern uint32_t nox_xxx_useAudio_587000_80772;
extern uint32_t dword_5d4594_1307792;
extern uint32_t dword_5d4594_830272;
extern uint32_t dword_5d4594_2386836;
extern uint32_t dword_5d4594_2490504;
extern uint32_t dword_5d4594_1316452;
extern uint32_t dword_5d4594_1049692;
extern uint32_t dword_5d4594_1045596;
extern uint32_t dword_5d4594_2488620;
extern uint32_t dword_5d4594_1047532;
extern uint32_t dword_5d4594_1045576;
extern uint32_t dword_5d4594_2487656;
extern uint32_t dword_587000_122956;
extern uint32_t dword_5d4594_2491716;
extern uint32_t dword_587000_180476;
extern uint32_t dword_5d4594_1046852;
extern uint32_t array_5D4594_1049872[9];

extern nox_window* dword_5d4594_1321236;
extern nox_window* dword_5d4594_1321240;
extern nox_window* dword_5d4594_1321248;
extern nox_window* dword_5d4594_1321244;

extern nox_window* dword_5d4594_1522616;
extern nox_window* dword_5d4594_1522620;
extern nox_window* dword_5d4594_1522624;
extern nox_window* dword_5d4594_1522628;
extern nox_window* dword_5d4594_1062452;

extern unsigned int dword_5d4594_2660032;
extern void* dword_5d4594_814624;
extern unsigned int dword_5d4594_2649712;
extern unsigned int dword_587000_145664;
extern unsigned int dword_587000_145668;
extern unsigned int nox_gameDisableMapDraw_5d4594_2650672;
extern void* nox_alloc_chat_1197364;
extern uint32_t dword_5d4594_1090120;
extern nox_window* dword_5d4594_1090048;
extern nox_window* dword_5d4594_1090100;
extern unsigned int dword_5d4594_251744;
extern unsigned int dword_5d4594_1193156;
extern unsigned int dword_5d4594_1193188;
extern unsigned int dword_5d4594_2650676;
extern unsigned int dword_5d4594_2650680;
extern uint32_t dword_5d4594_3798820;
extern uint32_t dword_5d4594_3798824;
extern uint32_t dword_5d4594_3798836;
extern uint32_t dword_5d4594_3798840;
extern unsigned int nox_xxx_waypointCounterMB_587000_154948;
extern unsigned int nox_client_fadeObjects_80836;
extern unsigned int dword_5d4594_811904;
extern unsigned int nox_client_gui_flag_1556112;
extern unsigned int nox_client_highResFloors_154952;
extern unsigned int nox_client_highResFrontWalls_80820;
extern unsigned int nox_player_netCode_85319C;
extern unsigned int dword_5d4594_1599644;
extern uint32_t dword_5d4594_1599480;
extern uint32_t dword_5d4594_1599476;
extern void* dword_5d4594_1599540;
extern void* dword_5d4594_1599532;
extern void* dword_5d4594_1599556;
extern void* dword_5d4594_1599548;
extern void* dword_5d4594_1599588;
extern void* dword_5d4594_1599592;
extern unsigned int nox_game_createOrJoin_815048;
extern unsigned int nox_client_gui_flag_1556112;
extern unsigned int dword_5d4594_1548524;
extern unsigned int dword_5d4594_2650652;
extern void* dword_5d4594_1548532;
extern unsigned int dword_5d4594_1556144;
extern unsigned int dword_5d4594_251744;
extern unsigned int dword_5d4594_815052;
extern unsigned int dword_5d4594_1049508;
extern unsigned int nox_client_translucentFrontWalls_805844;

extern uint32_t dword_5d4594_10984;
extern unsigned int dword_5d4594_527988;
extern uint32_t dword_5d4594_528256;
extern void* dword_5d4594_830232;
extern void* dword_5d4594_830236;
extern unsigned int dword_5d4594_831220;
extern uint32_t dword_5d4594_831260;
extern uint32_t dword_5d4594_1563096;
extern uint32_t dword_5d4594_831224;
extern unsigned int dword_5d4594_251744;
extern unsigned int nox_game_createOrJoin_815048;
extern nox_screenParticle* dword_5d4594_806052;
extern nox_screenParticle* nox_screenParticles_head;
extern void* nox_alloc_screenParticles_806044;
extern nox_window* dword_5d4594_1082856;
extern void* dword_5d4594_1082864;
extern void* dword_5d4594_1082868;

extern uint32_t nox_color_white_2523948;
extern uint32_t nox_color_red_2589776;
extern uint32_t nox_color_blue_2650684;
extern uint32_t nox_color_green_2614268;
extern uint32_t nox_color_cyan_2649820;
extern uint32_t nox_color_yellow_2589772;
extern uint32_t nox_color_violet_2598268;
extern uint32_t nox_color_black_2650656;
extern uint32_t nox_color_orange_2614256;

extern unsigned int dword_5d4594_1064868;
extern unsigned int dword_5d4594_1316972;
extern unsigned int gameex_flags;

extern void* dword_5d4594_1308156;
extern void* dword_5d4594_1308160;
extern void* dword_5d4594_1308164;

extern int dword_5d4594_2386848;
extern unsigned int dword_5d4594_2386852;

extern int dword_587000_87412;
extern unsigned int dword_5d4594_527988;
extern unsigned int nox_client_connError_814552;
extern nox_window* dword_5d4594_815004;
extern unsigned int dword_5d4594_815044;
extern nox_window* nox_wol_wnd_world_814980;
extern unsigned int nox_wol_server_result_cnt_815088;
extern int dword_5d4594_815104;
extern unsigned int nox_game_createOrJoin_815048;
extern uint32_t dword_587000_87408;
extern void* dword_5d4594_814984;
extern void* dword_5d4594_814996;
extern void* dword_5d4594_815000;
extern uint32_t dword_5d4594_815056;

extern unsigned int nox_client_translucentFrontWalls_805844;
extern unsigned int nox_client_highResFrontWalls_80820;
extern unsigned int nox_client_highResFloors_154952;
extern unsigned int nox_client_fadeObjects_80836;
extern unsigned int nox_client_renderBubbles_80844;
extern unsigned int dword_5d4594_2523804;
extern unsigned int dword_5d4594_3799452;
extern unsigned int dword_5d4594_1193188;
extern uint32_t dword_5d4594_3798800;
extern uint32_t dword_5d4594_3798804;
extern uint32_t dword_5d4594_3798808;
extern uint32_t dword_5d4594_3798812;
extern uint32_t dword_5d4594_3798816;
extern uint32_t dword_5d4594_3798820;
extern uint32_t dword_5d4594_3798824;
extern uint32_t dword_5d4594_3798828;
extern uint32_t dword_5d4594_3798832;
extern uint32_t dword_5d4594_3798836;
extern uint32_t dword_5d4594_3798840;
extern uint32_t dword_8531A0_2576;
extern uint32_t dword_8531A0_2572;
extern unsigned int dword_5d4594_1599628;
extern void (*func_587000_154940)(int2*, uint32_t, uint32_t);
extern int (*func_587000_154944)(int, int);
extern void* nox_video_tileBuf_ptr_3798796;
extern void* nox_video_tileBuf_end_3798844;

extern unsigned int dword_5d4594_1193188;
extern unsigned int dword_5d4594_1305748;
extern int nox_win_width;
extern int nox_win_height;

extern uint8_t** nox_pixbuffer_rows_3798784;
extern uint32_t dword_5d4594_823776;

extern nox_render_data_t* nox_draw_curDrawData_3799572;
*/
import "C"
import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/client/gui"
	"github.com/noxworld-dev/opennox/v1/client/noxrender"
	"github.com/noxworld-dev/opennox/v1/server"
)

func Sub_460D40() bool {
	return C.dword_5d4594_1049508 != 0
}

func Sub_4D7140(a1 uint32) {
	C.dword_5d4594_1556144 = C.uint(a1)
}

func Sub_43AF30() uint32 {
	return uint32(C.dword_5d4594_815052)
}

func Get_nox_wol_server_result_cnt_815088() int {
	return int(C.nox_wol_server_result_cnt_815088)
}
func Get_dword_5d4594_815044() int {
	return int(C.dword_5d4594_815044)
}
func Get_dword_587000_87412() int {
	return int(C.dword_587000_87412)
}
func Get_dword_5d4594_815000() *gui.Window {
	return AsWindowP(C.dword_5d4594_815000)
}
func Get_dword_587000_87408() int {
	return int(C.dword_587000_87408)
}
func Get_nox_game_createOrJoin_815048() int {
	return int(C.nox_game_createOrJoin_815048)
}
func Get_dword_5d4594_814996() *gui.Window {
	return AsWindowP(C.dword_5d4594_814996)
}
func Get_dword_5d4594_814984() *gui.Window {
	return AsWindowP(C.dword_5d4594_814984)
}
func Get_dword_5d4594_2488604() int {
	return int(C.dword_5d4594_2488604)
}
func Get_dword_5d4594_528252() int {
	return int(C.dword_5d4594_528252)
}
func Get_dword_5d4594_528260() int {
	return int(C.dword_5d4594_528260)
}
func Get_dword_587000_93164() unsafe.Pointer {
	return C.dword_587000_93164
}
func Get_nox_client_translucentFrontWalls_805844() int {
	return int(C.nox_client_translucentFrontWalls_805844)
}
func Get_nox_client_highResFrontWalls_80820() int {
	return int(C.nox_client_highResFrontWalls_80820)
}
func Get_nox_client_highResFloors_154952() int {
	return int(C.nox_client_highResFloors_154952)
}
func Get_nox_client_fadeObjects_80836() int {
	return int(C.nox_client_fadeObjects_80836)
}
func Get_nox_client_renderBubbles_80844() int {
	return int(C.nox_client_renderBubbles_80844)
}
func Get_nox_server_sanctuaryHelp_54276() int {
	return int(C.nox_server_sanctuaryHelp_54276)
}
func Get_nox_server_sendMotd_108752() int {
	return int(C.nox_server_sendMotd_108752)
}
func Get_nox_server_connectionType_3596() int {
	return int(C.nox_server_connectionType_3596)
}
func Get_nox_server_resetQuestMinVotes_229988() int {
	return int(C.nox_server_resetQuestMinVotes_229988)
}
func Get_nox_server_kickQuestPlayerMinVotes_229992() int {
	return int(C.nox_server_kickQuestPlayerMinVotes_229992)
}
func Get_dword_5d4594_816340() int {
	return int(C.dword_5d4594_816340)
}
func Get_dword_587000_93156() int {
	return int(C.dword_587000_93156)
}
func Get_dword_587000_93160() int {
	return int(C.dword_587000_93160)
}
func Set_dword_587000_93160(v int) {
	C.dword_587000_93160 = C.uint(v)
}
func Get_dword_5d4594_816348() int {
	return int(C.dword_5d4594_816348)
}
func Get_dword_5d4594_811904() int {
	return int(C.dword_5d4594_811904)
}
func Get_dword_5d4594_1193156() int {
	return int(C.dword_5d4594_1193156)
}
func Get_dword_5d4594_1193188() int {
	return int(C.dword_5d4594_1193188)
}
func Get_dword_5d4594_3798840() int {
	return int(C.dword_5d4594_3798840)
}
func Get_dword_5d4594_3798824() int {
	return int(C.dword_5d4594_3798824)
}
func Get_dword_5d4594_3798836() int {
	return int(C.dword_5d4594_3798836)
}
func Get_dword_5d4594_3798820() int {
	return int(C.dword_5d4594_3798820)
}
func Get_nox_alloc_chat_1197364() unsafe.Pointer {
	return C.nox_alloc_chat_1197364
}
func Get_dword_5d4594_2660032() int {
	return int(C.dword_5d4594_2660032)
}
func Get_dword_5d4594_1090048() *gui.Window {
	return asWindow(C.dword_5d4594_1090048)
}
func Get_dword_5d4594_1090120() int {
	return int(C.dword_5d4594_1090120)
}
func Get_dword_5d4594_1090100() *gui.Window {
	return asWindow(C.dword_5d4594_1090100)
}
func Get_dword_5d4594_814624() unsafe.Pointer {
	return C.dword_5d4594_814624
}
func Get_dword_5d4594_251560() unsafe.Pointer {
	return C.dword_5d4594_251560
}
func Get_dword_5d4594_2489460() int {
	return int(C.dword_5d4594_2489460)
}
func Get_gameex_flags() uint32 {
	return uint32(C.gameex_flags)
}
func Get_gameex_flags_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.gameex_flags))
}
func Get_dword_5d4594_1064868() int {
	return int(C.dword_5d4594_1064868)
}
func Get_dword_5d4594_1316972() *gui.Window {
	return AsWindowP(unsafe.Pointer(uintptr(C.dword_5d4594_1316972)))
}
func Get_dword_5d4594_1599588() unsafe.Pointer {
	return C.dword_5d4594_1599588
}
func Get_dword_5d4594_1599592() unsafe.Pointer {
	return C.dword_5d4594_1599592
}
func Get_dword_5d4594_1599540() unsafe.Pointer {
	return C.dword_5d4594_1599540
}
func Get_dword_5d4594_1599476() int {
	return int(C.dword_5d4594_1599476)
}
func Get_dword_5d4594_1599548() unsafe.Pointer {
	return C.dword_5d4594_1599548
}
func Get_dword_5d4594_1599556() unsafe.Pointer {
	return C.dword_5d4594_1599556
}
func Get_dword_5d4594_1599532() unsafe.Pointer {
	return C.dword_5d4594_1599532
}
func Get_dword_5d4594_1200804() int {
	return int(C.dword_5d4594_1200804)
}
func Get_nox_gameDisableMapDraw_5d4594_2650672() int {
	return int(C.nox_gameDisableMapDraw_5d4594_2650672)
}
func Get_dword_5d4594_528256() int {
	return int(C.dword_5d4594_528256)
}
func Get_dword_5d4594_527988() int {
	return int(C.dword_5d4594_527988)
}
func Get_nox_alloc_screenParticles_806044() unsafe.Pointer {
	return C.nox_alloc_screenParticles_806044
}
func Get_dword_5d4594_251744() int {
	return int(C.dword_5d4594_251744)
}
func Get_dword_5d4594_830236() *gui.Window {
	return AsWindowP(C.dword_5d4594_830236)
}
func Get_dword_5d4594_830232() *gui.Window {
	return AsWindowP(C.dword_5d4594_830232)
}
func Get_dword_5d4594_831260() int {
	return int(C.dword_5d4594_831260)
}
func Get_dword_5d4594_831220() int {
	return int(C.dword_5d4594_831220)
}
func Get_nox_color_black_2650656() uint32 {
	return uint32(C.nox_color_black_2650656)
}
func Get_nox_color_white_2523948() uint32 {
	return uint32(C.nox_color_white_2523948)
}
func Get_nox_color_violet_2598268() uint32 {
	return uint32(C.nox_color_violet_2598268)
}
func Get_nox_color_red_2589776() uint32 {
	return uint32(C.nox_color_red_2589776)
}
func Get_nox_color_green_2614268() uint32 {
	return uint32(C.nox_color_green_2614268)
}
func Get_nox_color_cyan_2649820() uint32 {
	return uint32(C.nox_color_cyan_2649820)
}
func Get_nox_color_blue_2650684() uint32 {
	return uint32(C.nox_color_blue_2650684)
}
func Get_nox_color_orange_2614256() uint32 {
	return uint32(C.nox_color_orange_2614256)
}
func Get_nox_color_yellow_2589772() uint32 {
	return uint32(C.nox_color_yellow_2589772)
}
func Get_nox_color_black_2650656_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.nox_color_black_2650656))
}
func Get_nox_color_white_2523948_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.nox_color_white_2523948))
}
func Get_nox_color_violet_2598268_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.nox_color_violet_2598268))
}
func Get_nox_color_red_2589776_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.nox_color_red_2589776))
}
func Get_nox_color_green_2614268_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.nox_color_green_2614268))
}
func Get_nox_color_cyan_2649820_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.nox_color_cyan_2649820))
}
func Get_nox_color_blue_2650684_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.nox_color_blue_2650684))
}
func Get_nox_color_orange_2614256_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.nox_color_orange_2614256))
}
func Get_nox_color_yellow_2589772_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.nox_color_yellow_2589772))
}
func Get_dword_5d4594_823776() int {
	return int(C.dword_5d4594_823776)
}
func Get_dword_5d4594_831236() *gui.Window {
	return AsWindowP(unsafe.Pointer(uintptr(C.dword_5d4594_831236)))
}
func Get_dword_5d4594_1049844() int {
	return int(C.dword_5d4594_1049844)
}
func Get_dword_5d4594_3799452() int {
	return int(C.dword_5d4594_3799452)
}
func Get_dword_5d4594_1321236() *gui.Window {
	return asWindow(C.dword_5d4594_1321236)
}
func Get_dword_5d4594_1321240() *gui.Window {
	return asWindow(C.dword_5d4594_1321240)
}
func Get_dword_5d4594_1321244() *gui.Window {
	return asWindow(C.dword_5d4594_1321244)
}
func Get_dword_5d4594_1321248() *gui.Window {
	return asWindow(C.dword_5d4594_1321248)
}
func Get_dword_5d4594_1522616() *gui.Window {
	return asWindow(C.dword_5d4594_1522616)
}
func Get_dword_5d4594_1522620() *gui.Window {
	return asWindow(C.dword_5d4594_1522620)
}
func Get_dword_5d4594_1522624() *gui.Window {
	return asWindow(C.dword_5d4594_1522624)
}
func Get_dword_5d4594_1522628() *gui.Window {
	return asWindow(C.dword_5d4594_1522628)
}
func Get_dword_5d4594_1062452() *gui.Window {
	return asWindow(C.dword_5d4594_1062452)
}
func Get_dword_5d4594_2516356() int {
	return int(C.dword_5d4594_2516356)
}
func Get_dword_5d4594_1548524() int {
	return int(C.dword_5d4594_1548524)
}
func Get_dword_587000_155144() unsafe.Pointer {
	return C.dword_587000_155144
}
func Get_nox_alloc_magicEnt_1569668() unsafe.Pointer {
	return C.nox_alloc_magicEnt_1569668
}
func Get_nox_client_gui_flag_1556112() int {
	return int(C.nox_client_gui_flag_1556112)
}

func Set_dword_5d4594_1599628(v int) {
	C.dword_5d4594_1599628 = C.uint(v)
}

func Sub_41E2F0() int {
	return int(C.dword_5d4594_527988)
}

func Get_nox_wol_wnd_world_814980() *gui.Window {
	return asWindow(C.nox_wol_wnd_world_814980)
}

func Get_dword_5d4594_815004() *gui.Window {
	return asWindow(C.dword_5d4594_815004)
}

func Nox_xxx_get_57AF20() int {
	return int(C.dword_5d4594_2523804)
}

func Set_dword_5d4594_1599644(v int) {
	C.dword_5d4594_1599644 = C.uint(v)
}
func Set_dword_5d4594_1599480(v uint32) {
	C.dword_5d4594_1599480 = C.uint(v)
}
func Set_dword_5d4594_1599476(v int) {
	C.dword_5d4594_1599476 = C.uint(v)
}
func Set_dword_5d4594_1599540(v unsafe.Pointer) {
	C.dword_5d4594_1599540 = v
}
func Set_dword_5d4594_1599532(v unsafe.Pointer) {
	C.dword_5d4594_1599532 = v
}
func Set_dword_5d4594_1599556(v unsafe.Pointer) {
	C.dword_5d4594_1599556 = v
}
func Set_dword_5d4594_1599548(v unsafe.Pointer) {
	C.dword_5d4594_1599548 = v
}
func Get_dword_5d4594_2386836() int {
	return int(C.dword_5d4594_2386836)
}
func Set_dword_5d4594_2386836(v int) {
	C.dword_5d4594_2386836 = C.uint(v)
}
func Set_dword_5d4594_2386848(v int) {
	C.dword_5d4594_2386848 = C.int(v)
}
func Set_dword_5d4594_2386852(v int) {
	C.dword_5d4594_2386852 = C.uint(v)
}
func Set_dword_5d4594_815104(v int) {
	C.dword_5d4594_815104 = C.int(v)
}
func Set_nox_wol_server_result_cnt_815088(v int) {
	C.nox_wol_server_result_cnt_815088 = C.uint(v)
}
func Set_qword_5d4594_815068(v uint64) {
	C.qword_5d4594_815068 = C.ulonglong(v)
}
func Set_nox_client_connError_814552(v int) {
	C.nox_client_connError_814552 = C.uint(v)
}
func Set_dword_5d4594_815056(v int) {
	C.dword_5d4594_815056 = C.uint(v)
}
func Set_nox_game_createOrJoin_815048(v int) {
	C.nox_game_createOrJoin_815048 = C.uint(v)
}
func Set_dword_5d4594_527988(v int) {
	C.dword_5d4594_527988 = C.uint(v)
}
func Set_dword_5d4594_10984(v int) {
	C.dword_5d4594_10984 = C.uint(v)
}
func Set_nox_alloc_screenParticles_806044(v unsafe.Pointer) {
	C.nox_alloc_screenParticles_806044 = v
}
func Set_nox_screenParticles_head(v unsafe.Pointer) {
	C.nox_screenParticles_head = (*C.nox_screenParticle)(v)
}
func Set_dword_5d4594_806052(v unsafe.Pointer) {
	C.dword_5d4594_806052 = (*C.nox_screenParticle)(v)
}
func Set_dword_5d4594_830236(v unsafe.Pointer) {
	C.dword_5d4594_830236 = v
}
func Set_nox_gameDisableMapDraw_5d4594_2650672(v int) {
	C.nox_gameDisableMapDraw_5d4594_2650672 = C.uint(v)
}
func Set_dword_5d4594_831224(v int) {
	C.dword_5d4594_831224 = C.uint(v)
}
func Set_dword_5d4594_1563096(v int) {
	C.dword_5d4594_1563096 = C.uint(v)
}
func Set_nox_win_width(v int) {
	C.nox_win_width = C.int(v)
}
func Set_nox_win_height(v int) {
	C.nox_win_height = C.int(v)
}
func Set_dword_5d4594_823776(v int) {
	C.dword_5d4594_823776 = C.uint(v)
}
func Set_nox_draw_curDrawData_3799572(v *noxrender.RenderData) {
	C.nox_draw_curDrawData_3799572 = (*nox_render_data_t)(v.C())
}
func Set_nox_color_black_2650656(v uint32) {
	C.nox_color_black_2650656 = C.uint(v)
}
func Set_nox_color_white_2523948(v uint32) {
	C.nox_color_white_2523948 = C.uint(v)
}
func Set_nox_color_violet_2598268(v uint32) {
	C.nox_color_violet_2598268 = C.uint(v)
}
func Set_nox_color_red_2589776(v uint32) {
	C.nox_color_red_2589776 = C.uint(v)
}
func Set_nox_color_green_2614268(v uint32) {
	C.nox_color_green_2614268 = C.uint(v)
}
func Set_nox_color_cyan_2649820(v uint32) {
	C.nox_color_cyan_2649820 = C.uint(v)
}
func Set_nox_color_blue_2650684(v uint32) {
	C.nox_color_blue_2650684 = C.uint(v)
}
func Set_nox_color_orange_2614256(v uint32) {
	C.nox_color_orange_2614256 = C.uint(v)
}
func Set_nox_color_yellow_2589772(v uint32) {
	C.nox_color_yellow_2589772 = C.uint(v)
}
func Set_dword_5d4594_1305748(v int) {
	C.dword_5d4594_1305748 = C.uint(v)
}
func Set_nox_pixbuffer_rows_3798784(v unsafe.Pointer) {
	C.nox_pixbuffer_rows_3798784 = (**C.uchar)(v)
}
func Set_dword_5d4594_1193188(v int) {
	C.dword_5d4594_1193188 = C.uint(v)
}
func Set_dword_5d4594_2649712(v uint32) {
	C.dword_5d4594_2649712 = C.uint(v)
}
func Set_dword_5d4594_1548524(v int) {
	C.dword_5d4594_1548524 = C.uint(v)
}
func Set_nox_client_gui_flag_1556112(v int) {
	C.nox_client_gui_flag_1556112 = C.uint(v)
}
func Set_dword_5d4594_3798812(v int) {
	C.dword_5d4594_3798812 = C.uint(v)
}
func Set_dword_5d4594_3798800(v int) {
	C.dword_5d4594_3798800 = C.uint(v)
}
func Set_dword_5d4594_3798816(v int) {
	C.dword_5d4594_3798816 = C.uint(v)
}
func Set_dword_5d4594_3798808(v int) {
	C.dword_5d4594_3798808 = C.uint(v)
}
func Set_dword_5d4594_3798804(v int) {
	C.dword_5d4594_3798804 = C.uint(v)
}
func Set_dword_5d4594_3798820(v int) {
	C.dword_5d4594_3798820 = C.uint(v)
}
func Set_dword_5d4594_3798824(v int) {
	C.dword_5d4594_3798824 = C.uint(v)
}
func Set_dword_5d4594_3798828(v int) {
	C.dword_5d4594_3798828 = C.uint(v)
}
func Set_dword_5d4594_3798832(v int) {
	C.dword_5d4594_3798832 = C.uint(v)
}
func Set_dword_5d4594_3798836(v int) {
	C.dword_5d4594_3798836 = C.uint(v)
}
func Set_dword_5d4594_3798840(v int) {
	C.dword_5d4594_3798840 = C.uint(v)
}
func Set_nox_video_tileBuf_ptr_3798796(v unsafe.Pointer) {
	C.nox_video_tileBuf_ptr_3798796 = v
}
func Set_nox_video_tileBuf_end_3798844(v unsafe.Pointer) {
	C.nox_video_tileBuf_end_3798844 = v
}
func Set_func_587000_154940(v unsafe.Pointer) {
	C.func_587000_154940 = (*[0]byte)(v)
}
func Set_func_587000_154944(v unsafe.Pointer) {
	C.func_587000_154944 = (*[0]byte)(v)
}
func Set_dword_5d4594_3799452(v int) {
	C.dword_5d4594_3799452 = C.uint(v)
}
func Set_nox_alloc_chat_1197364(v unsafe.Pointer) {
	C.nox_alloc_chat_1197364 = v
}
func Set_dword_587000_145664(v int) {
	C.dword_587000_145664 = C.uint(v)
}
func Set_dword_587000_145668(v int) {
	C.dword_587000_145668 = C.uint(v)
}
func Set_dword_5d4594_1090120(v int) {
	C.dword_5d4594_1090120 = C.uint(v)
}
func Set_dword_5d4594_251744(v int) {
	C.dword_5d4594_251744 = C.uint(v)
}
func Set_dword_5d4594_816340(v int) {
	C.dword_5d4594_816340 = C.uint(v)
}
func Set_dword_5d4594_816348(v int) {
	C.dword_5d4594_816348 = C.uint(v)
}
func Set_dword_5d4594_2489460(v int) {
	C.dword_5d4594_2489460 = C.uint(v)
}
func Set_dword_5d4594_1548532(v unsafe.Pointer) {
	C.dword_5d4594_1548532 = v
}
func Set_dword_5d4594_831260(v int) {
	C.dword_5d4594_831260 = C.uint(v)
}
func Set_nox_client_translucentFrontWalls_805844(v int) {
	C.nox_client_translucentFrontWalls_805844 = C.uint(v)
}
func Set_nox_client_highResFrontWalls_80820(v int) {
	C.nox_client_highResFrontWalls_80820 = C.uint(v)
}
func Set_nox_client_highResFloors_154952(v int) {
	C.nox_client_highResFloors_154952 = C.uint(v)
}
func Set_nox_client_fadeObjects_80836(v int) {
	C.nox_client_fadeObjects_80836 = C.uint(v)
}
func Set_nox_client_renderBubbles_80844(v int) {
	C.nox_client_renderBubbles_80844 = C.uint(v)
}
func Set_nox_server_sanctuaryHelp_54276(v int) {
	C.nox_server_sanctuaryHelp_54276 = C.uint(v)
}
func Set_nox_server_sendMotd_108752(v int) {
	C.nox_server_sendMotd_108752 = C.uint(v)
}
func Set_nox_server_connectionType_3596(v int) {
	C.nox_server_connectionType_3596 = C.uint(v)
}
func Set_nox_server_resetQuestMinVotes_229988(v int) {
	C.nox_server_resetQuestMinVotes_229988 = C.uint(v)
}
func Set_nox_server_kickQuestPlayerMinVotes_229992(v int) {
	C.nox_server_kickQuestPlayerMinVotes_229992 = C.uint(v)
}
func Set_dword_5d4594_1308156(v unsafe.Pointer) {
	C.dword_5d4594_1308156 = v
}
func Set_dword_5d4594_1308160(v unsafe.Pointer) {
	C.dword_5d4594_1308160 = v
}
func Set_dword_5d4594_1308164(v unsafe.Pointer) {
	C.dword_5d4594_1308164 = v
}
func Set_dword_5d4594_2650676(v int) {
	C.dword_5d4594_2650676 = C.uint(v)
}
func Set_dword_5d4594_2650680(v int) {
	C.dword_5d4594_2650680 = C.uint(v)
}
func Set_qword_5d4594_814956(v uint64) {
	C.qword_5d4594_814956 = C.ulonglong(v)
}
func Set_dword_5d4594_1200804(v uint32) {
	C.dword_5d4594_1200804 = C.uint(v)
}
func Set_dword_5d4594_1200832(v int) {
	C.dword_5d4594_1200832 = C.uint(v)
}
func Set_nox_alloc_magicEnt_1569668(v unsafe.Pointer) {
	C.nox_alloc_magicEnt_1569668 = v
}
func Set_dword_5d4594_1569672(v int) {
	C.dword_5d4594_1569672 = C.uint(v)
}
func Inc_nox_xxx_waypointCounterMB_587000_154948() {
	C.nox_xxx_waypointCounterMB_587000_154948++
}
func Inc_dword_5d4594_2516356() {
	C.dword_5d4594_2516356++
}
func Inc_dword_5d4594_2386848() {
	C.dword_5d4594_2386848++
}
func Inc_nox_wol_server_result_cnt_815088() {
	C.nox_wol_server_result_cnt_815088++
}
func Set_dword_5d4594_1599588(v unsafe.Pointer) {
	C.dword_5d4594_1599588 = v
}
func Set_dword_5d4594_1599592(v unsafe.Pointer) {
	C.dword_5d4594_1599592 = v
}
func Get_dword_5d4594_1599592_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.dword_5d4594_2649712))
}
func Get_dword_5d4594_1123524() *gui.Window {
	return AsWindowP(C.dword_5d4594_1123524)
}
func Set_dword_5d4594_1123520(v int) {
	C.dword_5d4594_1123520 = C.uint(v)
}
func Get_dword_5d4594_826036() *byte {
	return (*byte)(C.dword_5d4594_826036)
}
func Get_dword_8531A0_2576() *server.Player {
	return AsPlayerP(unsafe.Pointer(uintptr(C.dword_8531A0_2576)))
}
func Set_dword_8531A0_2576(p *server.Player) {
	C.dword_8531A0_2576 = C.uint(uintptr(p.C()))
}
func Get_dword_8531A0_2572() uint32 {
	return uint32(C.dword_8531A0_2572)
}
func Get_dword_8531A0_2572_ptr() *uint32 {
	return (*uint32)(unsafe.Pointer(&C.dword_8531A0_2572))
}
func Set_dword_8531A0_2572(v uint32) {
	C.dword_8531A0_2572 = C.uint(v)
}
func Get_dword_5d4594_2488620() uint32 {
	return uint32(C.dword_5d4594_2488620)
}
func Nox_xxx_playerGet_470A90() int {
	return int(C.dword_5d4594_1096252)
}
