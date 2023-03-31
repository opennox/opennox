#ifndef NOX_PORT_GAME4_3
#define NOX_PORT_GAME4_3

#include "defs.h"

int nox_xxx_onFrameLightning_52F8A0(float a1);
void nox_xxx_lightningCanAttackCheck_52FF10(int a1, int a2);
void nox_xxx_lightningSpellTrapEffect_530020(int a1, int a2);
char sub_530100(uint32_t* a1);
int nox_xxx_spellTagCreature_530160(uint32_t* a1);
unsigned int sub_530250(int a1);
int sub_530270(int a1);
int nox_xxx_spellBlink2_530310(uint32_t* a1);
int nox_xxx_spellBlink1_530380(int* a1);
uint32_t* nox_xxx_spellTeleportCreateWake_530560(int a1, int* a2, uint32_t* a3);
int sub_5305D0(uint32_t* a1);
int sub_530650(int* a1);
int nox_xxx_castTele_530820(int a1);
int sub_530880(int* a1);
int nox_xxx_castTTT_530B70(int* a1);
int sub_530CA0(int a1);
int sub_530D30(int* a1);
int nox_xxx_manaBomb_530F90(uint32_t* a1);
int nox_xxx_manaBombBoom_5310C0(int* a1);
int sub_531290(int a1);
int nox_xxx_spellTurnUndeadCreate_531310(uint32_t* a1);
int nox_xxx_spellTurnUndeadUpdate_531410();
int nox_xxx_spellTurnUndeadDelete_531420(int a1);
int sub_531490(uint32_t* a1);
int sub_5314F0(int a1);
int sub_531560(int a1);
int nox_xxx_plasmaSmth_531580(int a1);
int nox_xxx_plasmaShot_531600(int a1);
void sub_531920(int a1, int a2);
int sub_5319E0(int a1);
int nox_xxx_spellCreateMoonglow_531A00(uint32_t* a1);
int sub_531AF0(int a1);
int* nox_xxx_TODOsomeCallingMeleeAttack_531B40(int a1, int a2);
int* sub_531C10(int a1, int a2);
int* nox_xxx_monsterAction_531C60(int a1, int a2);
int* sub_531D50(int a1, int a2);
int nox_xxx_mobActionFightStart_531E20(uint32_t a1);
int sub_531E90(int a1);
char nox_xxx_mobActionFight_531EC0(int a1);
int sub_532040(int a1, int a2);
char nox_xxx_monsterShieldBlockStart_532070(int a1);
char nox_xxx_monsterShieldBlockStop_5320E0(int a1);
char nox_ai_action_pop_532100(int a1);
char sub_532110(int a1);
int* nox_xxx_mobActionMelee1_532130(int a1);
void sub_532390(int a1, int a2);
char nox_xxx_mobActionMeleeAtt_532440(int a1);
uint32_t* sub_532540(int a1);
char nox_xxx_mobActionMissileAtt_532610(int a1);
char nox_xxx_monsterPlayHurtSound_532800(nox_object_t* a1);
int sub_532880(int a1);
int nox_xxx_soundPlayerDamageSound_5328B0(int a1, int a2);
void sub_532930(int a1, unsigned short a2, unsigned short a3);
int nox_xxx_soundDefaultDamageSound_532E20(nox_object_t* a1, nox_object_t* a2);
void sub_532EC0(int a1, unsigned short a2);
int sub_532F70(int a1);
int sub_532FB0(short a1);
int sub_532FE0(unsigned short a1, int a2);
int sub_533010(unsigned short a1, int a2);
int nox_xxx_projAddVelocitySmth_533080(int a1, int a2, float a3, int a4);
int nox_xxx_unitIsEnemyTo_5330C0(nox_object_t* a1, nox_object_t* a2);
int sub_533360(nox_object_t* a1, nox_object_t* a2);
nox_object_t* nox_xxx_enemyAggro_5335D0(nox_object_t* a1, float a2);
double sub_5336D0(nox_object_t* a1);
int nox_xxx_mobActionToAnimation_533790(int a1);
void nox_xxx_orderUnit_533900(nox_object_t* owner, nox_object_t* creature, int orderType);
void nox_xxx_enactUnitOrder_5339A0(int source, int unit, int orderId);
void nox_xxx_mobCalcDir_533CC0(int a1, float* a2);
unsigned char* nox_xxx_unitNPCActionToAnim_533D00(int a1);
int nox_xxx_monsterTestBlockShield_533E70(nox_object_t* a1);
void sub_533EB0(int a1, int a2);
int sub_534020(int a1);
char nox_xxx_monsterMoveAudio_534030(int a1);
int sub_534120(int a1, float2* a2);
void nox_ai_debug_printf_5341A0(char* a1, ...);
int sub_5341D0(int a1);
int sub_5341F0(nox_object_t* a1p);
int nox_xxx_monsterCanMelee_534220(int a1);
int nox_xxx_monsterCanShoot_534280(int a1);
int nox_xxx_monsterHasShield_5342C0(int a1);
int nox_xxx_monsterCanCast_534300(nox_object_t* a1);
int nox_xxx_monsterIsMoveing_534320(int a1);
int sub_534340(int a1);
int nox_xxx_monsterCanAttackAtWill_534390(nox_object_t* a1);
int sub_5343C0(int a1);
int sub_534400(int a1);
int sub_534440(int a1);
double sub_534470(int a1);
char* sub_5345B0(int a1);
int nox_xxx_actionNByNameMB_5345F0(const char* a1);
char* sub_534650(int a1);
int nox_xxx_actionByName_534670(const char* a1);
int sub_5346D0(int a1);
int nox_xxx_monsterResetEnemy_5346F0(int a1);
int sub_534710(int a1);
int sub_534750(int a1);
int sub_534780(int a1);
int sub_5347A0(nox_object_t* a1);
int sub_5347C0(int a1);
int nox_xxx_isNotPoisoned_5347F0(int a1);
int nox_xxx_mobGetMoveAttemptTime_534810(nox_object_t* a1);
int nox_xxx_unitIsMimic_534840(int a1);
void nox_xxx_monsterMimicCheckMorph_534950(nox_object_t* a1);
int nox_xxx_unitIsPlant_534A10(int a1);
int nox_xxx_unitIsZombie_534A40(int a1);
char nox_xxx_mobActionGetUp_534A90(int a1);
unsigned int nox_xxx_mobRaiseZombie_534AB0(int a1);
int nox_xxx_damageToMap_534BC0(int a1, int a2, int a3, int a4, nox_object_t* a5);
int nox_xxx_wallPreDestroy_534DA0(int* a1);
bool nox_xxx_mapDamageToWalls_534FC0(int4* a1, void* a2, float a3, int a4, int a5, void* a6);
int nox_xxx_mapTraceRay_535250(float4* a1, float2* a2, int2* a3, char a4);
char* sub_536130(char* a1, int* a2);
char* sub_536180(char* a1, int* a2);
char* sub_5361B0(char* a1, int a2);
char* sub_536260(char* a1, int a2);
char* sub_536390(char* a1, int* a2);
char* sub_5363C0(char* a1, int* a2);
int sub_5364E0(char* a1, int a2);
int sub_536550(char* a1, uint32_t* a2);
int sub_536580(char* a1, int a2);
int sub_5365B0(char* a1, int a2);
int sub_536600(char* a1, int a2);
int sub_536B40(char* a1, int a2);
int sub_536D80(char* a1, int a2);
int sub_536DA0(char* a1, int* a2);
int sub_536DE0(char* a1, uint8_t* a2);
int nox_xxx_collideDamageLoad_536E10(char* a1, int a2);
int sub_536E50(char* a1, uint8_t* a2);
int sub_536E80(char* a1, int* a2);
int nox_xxx_unitCanSee_536FB0(int a1, int a2, char a3);
int nox_xxx_unitCanInteractWith_5370E0(nox_object_t* a1, nox_object_t* a2, char a3);
int nox_xxx_mapCheck_537110(nox_object_t* a1, nox_object_t* a2);
void nox_xxx_lineCollisionChk_537230(float* a1, int arg4);
int nox_xxx_traceRay_5374B0(float4* a1);
void nox_xxx_harpoonBreakForPlr_537520(nox_object_t* a1);
void sub_537540(int a1);
int sub_537580(int a1);
void sub_5375A0(int a1);
char nox_xxx_unitHasCollideOrUpdateFn_537610(nox_object_t* a1);
int sub_537740();
int sub_537750(int a1);
unsigned int sub_537760();
void sub_537770(nox_object_t* a1);
char nox_xxx_projectileTraceHit_537850(int a1, int* a2, float2* a3);
void nox_xxx_sMakeScorch_537AF0(float* a1, int a2);
int nox_xxx_scorchInit_537BD0();
char nox_xxx_trapBAH_537C10(int a1, int a2);
void sub_537DD0(float* a1, int a2);
int sub_537E60(int a1, int a2, int a3, int a4);
void sub_537F00(float* a1, int a2);
char nox_xxx___mkgmtime_538280(int a1);
int nox_xxx_playerPreAttackEffects_538290(int a1, int a2, int a3, int a4);
int nox_xxx_playerTraceAttack_538330(int a1, int a2);
void sub_538510(int a1, int a2);
void sub_5386A0(int a3, int a2);
int nox_xxx_itemApplyAttackEffect_538840(int a1, int a2, int a3);
int nox_xxx_playerAttack_538960(nox_object_t* a1);
short nox_xxx_warcryStunMonsters_539B90(int a1, int a2);
int nox_xxx_shootBowCrossbow1_539BD0(int a1, int a2);
uint32_t* nox_xxx_shootBowCrossbow2_539D80(int a1, int a2, int a3, char* a4);
int nox_xxx_shootApplyEffects_539F40(int a1, int a2, int a3);
int sub_539FB0(uint32_t* a1);
int nox_xxx_playerTryReloadQuiver_539FF0(uint32_t* a1);
int nox_xxx_equipWeaponNPC_53A030(int a1, int a2);
void sub_53A0F0(int a1, int a2, int a3);
int nox_xxx_playerDequipWeapon_53A140(uint32_t* a1, nox_object_t* item, int a3, int a4);
int nox_xxx_NPCEquipWeapon_53A2C0(int a1, nox_object_t* item);
void sub_53A3D0(uint32_t* a1);
int nox_xxx_playerEquipWeapon_53A420(uint32_t* a1, nox_object_t* item, int a3, int a4);
int sub_53A680(int a1);
void sub_53A6C0(int a1, nox_object_t* item);
int sub_53A720(int a1, nox_object_t* item, int a3, int a4);
int nox_xxx_sendMsgOblivionPickup_53A9C0(int a1, nox_object_t* item, int a3, int a4);
void sub_53AAB0(int a1);
int nox_xxx_dropWeapon_53AB10(int a1, uint32_t* a2, int* a3);
void sub_53AB90(int a1, int a2);
char nox_xxx_updateDoor_53AC50(int a1);
void nox_xxx_updateSpark_53ADC0(int a1);
float* nox_xxx_updateProjTrail_53AEC0(int a1);
void nox_xxx_updatePush_53B030(int a1);
char nox_xxx_updateToggle_53B060(uint32_t* a1);
char nox_xxx_updateTrigger_53B1B0(int a1);
char sub_53B300(int a1);
char nox_xxx_updateSwitch_53B320(uint32_t* a1);
char nox_xxx_updateElevatorShaft_53B380(int a1);
void nox_xxx_fnElevatorShaft_53B410(int a1, int a2);
void nox_xxx_elevatorAud_53B490(int a1, int a2);
void nox_xxx_updateElevator_53B5D0(uint32_t* a1);
void nox_xxx_elevatorFn_53B750(int a1, int a2);
void nox_xxx_updatePhantomPlayer_53B860(int a1);
void nox_xxx_updateLifetime_53B8F0(int unit);
void nox_xxx_spellFlyUpdate_53B940(int a1);
void nox_xxx_updateAntiSpellProj_53BB00(int a1);
void sub_53BD10(int a1, int a2);
int nox_xxx_updateMagicMissile_53BDA0(int a1);
int nox_xxx_updateTeleportPentagram_53BEF0(int a1);
void nox_xxx_fnPentagramTeleport_53C060(float* a1, int a2);
int nox_xxx_updateInvisiblePentagram_53C0C0(int a1);
void sub_53C140(float* a1, int a2);
void nox_xxx_updateBlow_53C160(int a3);
void sub_53C240(float* a1, int arg4);
int nox_xxx_rechargeItem_53C520(int a1, int a2);
signed int nox_xxx_updateObelisk_53C580(int a1);
int nox_xxx_getRechargeRate_53C940(uint32_t* a1);
void nox_xxx_updateBlackPowderBarrel_53C9A0(float* a1);
void nox_xxx_updateOneSecondDie_53CB60(int a1);
void nox_xxx_updateWaterBarrel_53CB90(int a1);
void nox_xxx_waterBarrel_53CC30(float* a1, int a2);
void nox_xxx_updateSelfDestruct_53CC90(int a1);
void nox_xxx_updateBlackPowderBurn_53CCB0(int a1);
void nox_xxx_updatePixie_53CD20(nox_object_t* a1);
void nox_xxx_updateDeathBall_53D080(int a1);
void sub_53D170(int a1, int a2);
void nox_xxx_updateDeathBallFragment_53D220(int a1);
void nox_xxx_updateMoonglow_53D270(int a1);
void nox_xxx_updateTelekinesis_53D330(int a1);
void nox_xxx_updateFist_53D400(int a1);
void nox_xxx_updateFlameCleanse_53D510(int a1);
void nox_xxx_updateMeteorShower_53D5A0(float* a2);
void nox_xxx_meteorExplode_53D6E0(int a6);
void nox_xxx_updateToxicCloud_53D850(int a1);
void sub_53D8C0(int a1, int a2);
void nox_xxx_updateSmallToxicCloud_53D960(int a1);
void nox_xxx_toxicCloudPoison_53D9D0(int a1, int a2);
void nox_xxx_updateArachnaphobia_53DA60(int* a1);
void nox_xxx_updateExpire_53DB00(int a1);
int* nox_xxx_updateBreak_53DB30(uint32_t* a1);
int* nox_xxx_updateOpen_53DBB0(uint32_t* a1);
void nox_xxx_updateBreakAndRemove_53DC30(uint32_t* a1);
void nox_xxx_updateChakramInMotion_53DCC0(int a1);
int nox_xxx_updateFlag_53DDF0(int a1);
int* nox_xxx_updateTrapDoor_53DE80(uint32_t* a1);
void nox_xxx_updateGameBall_53DF40(int a3);
void nox_xxx_updateUndeadKiller_53E190(int a1);
void nox_xxx_updateCrown_53E1D0(int a1);
int sub_53E2D0(int a1);
int nox_xxx_recalculateArmorVal_53E300(uint32_t* a1);
int sub_53E3A0(int a1, nox_object_t* object);
int sub_53E430(uint32_t* a1, nox_object_t* object, int a3, int a4);
int nox_xxx_NPCEquipArmor_53E520(int a1, uint32_t* a2);
void sub_53E600(uint32_t* a1);
int nox_xxx_playerEquipArmor_53E650(uint32_t* a1, nox_object_t* item, int a3, int a4);
uint32_t* nox_xxx_armorHaveSameSubclass_53E7B0(int a1, int a2);
int nox_xxx_pickupArmor_53E7F0(int a1, int a2, int a3, int a4);
void sub_53EAE0(int a1);
int nox_xxx_dropArmor_53EB70(int a1, uint32_t* a2, int* a3);
int nox_xxx_ItemIsDroppable_53EBF0(int a1);
char* sub_53EC40();
int sub_53EC80(int a1, int a2);
int nox_xxx_useMushroom_53ECE0(int a1, int a2);
int nox_xxx_useEnchant_53ED60(int a1, int a2);
int nox_xxx_useCast_53ED90(int a1, uint32_t* a2);
int nox_xxx_useConsume_53EE10(int a1, int a2);
int nox_xxx_useCiderConfuse_53EF00(int a1, int a2);
int nox_xxx_usePotion_53EF70(int a1, int a2);
int nox_xxx_useLesserFireballStaff_53F290(int a1, uint32_t* a2);
uint32_t* nox_xxx_wandShot_53F480(int a1, int a2, int* a3, uint32_t* a4);
int nox_xxx_useWandCastSpell_53F4F0(int a1, uint32_t* a2);
int nox_xxx_useFireWand_53F670(int a1, int a2);
int nox_xxx_useRead_53F7C0(int a1, int a2);
int sub_53F830(int a1, int a2);
int nox_xxx_useByNetCode_53F8E0(int a1, int a2);
int sub_53F930(int a1, int a2);
int nox_xxx_useSpellReward_53F9E0(int a1, int a2);
int nox_xxx_useAbilityReward_53FAE0(int a1, int a2);
uint32_t* nox_xxx_respawnPlayerImpl_53FBC0(float* a1, int a2);
void nox_xxx_createCorpse_53FCA0();
int nox_xxx_castPixies_540440(int a1, int a2, int a3, int a4, int a5, int a6);
nox_object_t* nox_xxx_spellFlySearchTarget_540610(float2* a1, nox_object_t* a2, int a3, float a4, int a5, nox_object_t* a6);
int sub_5408A0(int a1);
int nox_xxx_mobCastInversion_5408D0(int a1);
void nox_xxx_monsterCast_540A30(nox_object_t* a1, int a2, nox_object_t* a3);
int nox_xxx_unitIsMagicMissile_540B60(int a1, int a2);
int nox_xxx_monsterBuffSelf_540B90(int a1);
int sub_540CE0(int a1, int a2);
int sub_540D20(int a1);
int sub_540D40(int a1);
int nox_xxx_mobCastRelated2_540D90(int a1, int a2);
int nox_xxx_monsterCastOffensive_540F20(int a1, int a2);
int nox_xxx_mobCastRelated_541050(int a1);
int nox_xxx_mobHealSomeone_5411A0(nox_object_t* a3);
void nox_xxx_mobMayHealThis_5412A0(float* a1, int a2);
char nox_xxx_mobCast_541300(int a1, uint32_t* a2, int a3);
char nox_xxx_mobActionCast_5413B0(nox_object_t* a1, int a2);
void nox_xxx_mobCastRandomRecoil_541490(int a1, float* a2, float2* a3);
char sub_541630(int a1, int a2);
char* sub_542BF0(int a1, int a2, int a3);
char* sub_5435C0(int a1, int a2, int a3, int a4);
char* sub_543620(int a1, int a2);
int sub_543680(float* a1);
int sub_5437E0(int* a1, int a2, int a3);
void sub_543BC0(int a1, int a2, int a3, int a4, int a5, int a6);
int nox_xxx_tile_543C50(uint32_t* a1, int a2, int a3, int a4, int a5, int a6);
int sub_543E60(int a1, int a2);
int nox_xxx_mapGenEdge_543EB0(int a1, int a2);
int sub_543FB0(const char* a1);
int sub_544020(char* a1);
int nox_xxx_tileCheckByte3_544070(int a1);
int nox_xxx_tileCheckByte4_5440A0(int a1);
int nox_xxx_tileSubtile_544310(float2* a1);
char nox_xxx_mobActionMoveTo_5443F0(int a1);
char nox_xxx_mobActionMoveToFar_5445C0(int* a1);
void nox_xxx_mobActionDodge_544640(int a1);
int sub_544740(int a1);
int sub_544750(int a1);
char nox_xxx_mobActionFlee_544760(int a1);
int nox_xxx_mobActionReturnToHome_544920(int a1);
int sub_544930(int a1);
int sub_544940(int a1);
char sub_544950(int a1);
int* nox_xxx_mobActionHunt_5449D0(int a1);
int nox_xxx_mobSearchEdible_544A00(nox_object_t* a1, float a2);
void nox_xxx_mobSearchEdible2_544A40(int a1, int a2);
int sub_544AE0(int a1, float a2);
void sub_544B20(int a1, int a2);
char nox_xxx_mobActionPickupObject_544B90(int a1);
int nox_xxx_mobGenericDeath_544C40(int a1);
void nox_xxx_zombieBurnDeleteCheck_544CA0(uint32_t* a1);
void nox_xxx_zombieBurnDelete_544CE0(uint32_t* a1);
char sub_544D60(int a1);
char nox_xxx_mobActionDead1_544D80(uint32_t* a1);
void nox_xxx_createReleasedSoul_544E60(int a1);
int sub_544F70(int a1);
void nox_xxx_mobActionDead2_544EC0(int a1);

#endif // NOX_PORT_GAME4_3
