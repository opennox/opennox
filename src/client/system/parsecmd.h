#ifndef NOX_PORT_CLIENT_SYSTEM_PARSECMD
#define NOX_PORT_CLIENT_SYSTEM_PARSECMD

#include "../../defs.h"

int nox_cmd_racoiaws();

void sub_440A20(wchar_t* a1, ...);

int __cdecl nox_cmd_lock(int a1, char a2, int a3);
int __cdecl nox_cmd_unlock(int a1, char a2);
int __cdecl nox_cmd_set_sysop(int a1, char a2, int a3);
int __cdecl nox_cmd_telnet_off(int a1, char a2);
int __cdecl nox_cmd_telnet_on(int a1, unsigned __int8 a2, int a3);
int __cdecl nox_cmd_macros_on(int a1, char a2);
int __cdecl nox_cmd_macros_off(int a1, char a2);
int nox_cmd_list_weapons();
int nox_cmd_list_armor();
int nox_cmd_list_spells();
int nox_cmd_list_staffs();
int __cdecl nox_cmd_show_bindings(int a1, char a2);
int __cdecl nox_cmd_show_game(int a1, char a2);
int nox_cmd_show_mmx();
int __cdecl nox_cmd_load(int a1, char a2, int a3);
int sub_4444F0();
int __cdecl sub_443E90(int a1, char a2, wchar_t* a3);
int nox_cmd_set_obs();
int nox_cmd_set_save_debug();
int nox_cmd_set_god();
int nox_cmd_unset_god();
int nox_cmd_set_sage();
int nox_cmd_unset_sage();
int __cdecl nox_cmd_set_cycle(int a1, char a2, int a3);
int __cdecl nox_cmd_set_weapons(int a1, char a2, int a3);
int __cdecl nox_cmd_set_staffs(int a1, char a2, int a3);
int __cdecl nox_cmd_set_name(int a1, unsigned __int8 a2, int a3);
int __cdecl nox_cmd_set_mnstrs(int a1, char a2, int a3);
int __cdecl nox_cmd_set_spell(int a1, char a2, int a3);
int __cdecl nox_cmd_set_weapon(int a1, char a2, int a3);
int __cdecl nox_cmd_set_armor(int a1, char a2, int a3);
int __cdecl nox_cmd_set_staff(int a1, char a2, int a3);
int __cdecl nox_cmd_ban(int a1, char a2, int a3);
int nox_cmd_allow_user();
int nox_cmd_allow_ip();
int __cdecl nox_cmd_kick(int a1, char a2, int a3);
int __cdecl nox_cmd_set_players(int a1, char a2, int a3);
int nox_cmd_set_spellpts();
int nox_cmd_list_users();
int __cdecl nox_cmd_unmute(int a1, unsigned __int8 a2, int a3);
int __cdecl nox_cmd_mute(int a1, unsigned __int8 a2, int a3);
int __cdecl nox_cmd_exec(int a1, char a2);
int __cdecl nox_cmd_exec_rul(int a1, char a2, int a3);
int __cdecl sub_4439B0(int a1, unsigned __int8 a2);
int __cdecl sub_443C80(wchar_t* a1, int a2);
int __cdecl nox_cmd_unbind(int a1, char a2, int a3);
int __cdecl nox_cmd_broadcast(int a1, unsigned __int8 a2, int a3);
int __cdecl nox_cmd_say(int a1, char a2);
int __cdecl nox_cmd_offonly1(int a1, char a2);
int __cdecl nox_cmd_offonly2(int a1, char a2, int a3);
int __cdecl nox_cmd_set_fr(int a1, char a2);
int __cdecl nox_cmd_unset_fr(int a1, char a2);
int __cdecl nox_cmd_set_net_debug(int a1, char a2);
int __cdecl nox_cmd_unset_net_debug(int a1, char a2);
int nox_cmd_show_ai();
int nox_cmd_show_gui();
int nox_cmd_show_extents();
int nox_cmd_show_perfmon();
int nox_cmd_show_netstat();
int nox_cmd_show_info();
int nox_cmd_show_mem();
int nox_cmd_show_rank();
int __cdecl nox_cmd_show_motd(int a1, char a2);
int __cdecl nox_cmd_show_seq(int a1, char a2);
int nox_cmd_list_maps();
int __cdecl nox_cmd_log_file(int a1, char a2, int a3);
int __cdecl nox_cmd_log_console(int a1, char a2);
int __cdecl nox_cmd_log_stop(int a1, char a2);
int nox_cmd_set();
int nox_cmd_cheat_ability();
int __cdecl nox_cmd_cheat_goto(int a1, unsigned __int8 a2, int a3);
int nox_cmd_cheat_health();
int nox_cmd_cheat_mana();
int __cdecl nox_cmd_cheat_level(int a1, unsigned __int8 a2, int a3);
int __cdecl nox_cmd_cheat_spells(int a1, unsigned __int8 a2, int a3);
int __cdecl nox_cmd_cheat_gold(int a1, unsigned __int8 a2, int a3);
int nox_cmd_image();
int nox_cmd_quit();
int nox_cmd_exit();
int __cdecl nox_cmd_watch(int a1, unsigned __int8 a2, int a3);
int __cdecl nox_cmd_gamma(int a1, unsigned __int8 a2, int a3);
int __cdecl nox_cmd_window(int a1, unsigned __int8 a2, int a3);
int nox_cmd_set_qual_modem();
int nox_cmd_set_qual_isdn();
int nox_cmd_set_qual_cable();
int nox_cmd_set_qual_t1();
int nox_cmd_set_qual_lan();
int __cdecl nox_cmd_set_time(int a1, char a2, int a3);
int __cdecl nox_cmd_set_lessons(int a1, char a2, int a3);
int nox_cmd_clear();
int nox_cmd_menu_options();
int nox_cmd_menu_vidopt();
int __cdecl nox_cmd_help(int a1, int a2, int a3);
int __cdecl nox_cmd_bind(int a1, char a2, int a3);

#endif //NOX_PORT_CLIENT_SYSTEM_PARSECMD
