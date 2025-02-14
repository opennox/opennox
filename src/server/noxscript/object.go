package noxscript

import (
	"github.com/opennox/noxscript/ns/asm"
	ns4 "github.com/opennox/noxscript/ns/v4"
	"github.com/opennox/noxscript/ns/v4/audio"
	"github.com/opennox/noxscript/ns/v4/class"
	"github.com/opennox/noxscript/ns/v4/damage"
	"github.com/opennox/noxscript/ns/v4/subclass"
)

func init() {
	Register(asm.BuiltinGetTrigger, nsGetTrigger)
	Register(asm.BuiltinGetCaller, nsGetCaller)
	Register(asm.BuiltinIsTrigger, nsIsTrigger)
	Register(asm.BuiltinIsCaller, nsIsCaller)
	Register(asm.BuiltinObject, nsObject)
	Register(asm.BuiltinCreateObject, nsCreateObject)
	Register(asm.BuiltinGetObjectX, nsObjectX)
	Register(asm.BuiltinGetObjectY, nsObjectY)
	Register(asm.BuiltinMoveObject, nsSetPos)
	Register(asm.BuiltinGetObjectZ, nsZ)
	Register(asm.BuiltinRaise, nsSetZ)
	Register(asm.BuiltinIsObjectOn, nsIsEnabled)
	Register(asm.BuiltinObjectOn, nsObjectOn)
	Register(asm.BuiltinObjectOff, nsObjectOff)
	Register(asm.BuiltinObjectToggle, nsObjectToggle)
	Register(asm.BuiltinHasClass, nsHasClass)
	Register(asm.BuiltinHasSubclass, nsHasSubclass)
	Register(asm.BuiltinCurrentHealth, nsCurrentHealth)
	Register(asm.BuiltinMaxHealth, nsMaxHealth)
	Register(asm.BuiltinRestoreHealth, nsRestoreHealth)
	Register(asm.BuiltinGetDirection, nsGetDirection)
	Register(asm.BuiltinLookWithAngle, nsLookWithAngle)
	Register(asm.BuiltinLookAtObject, nsLookAtObject)
	Register(asm.BuiltinLookAtDirection, nsLookAtDirection)
	Register(asm.BuiltinDelete, nsDelete)
	Register(asm.BuiltinDeleteObjectTimer, nsDeleteObjectTimer)
	Register(asm.BuiltinPushObjectTo, nsPushTo)
	Register(asm.BuiltinPushObject, nsPush)
	Register(asm.BuiltinFrozen, nsFreeze)
	Register(asm.BuiltinPickup, nsPickup)
	Register(asm.BuiltinDrop, nsDrop)
	Register(asm.BuiltinIsOwnedBy, nsIsOwnedBy)
	Register(asm.BuiltinClearOwner, nsClearOwner)
	Register(asm.BuiltinSetOwner, nsSetOwner)
	Register(asm.BuiltinSetOwners, nsSetOwners)
	Register(asm.BuiltinIsOwnedByAny, nsIsOwnedByAny)
	Register(asm.BuiltinGoBackHome, nsReturn)
	Register(asm.BuiltinCreatureIdle, nsIdle)
	Register(asm.BuiltinWander, nsWander)
	Register(asm.BuiltinCreatureFollow, nsFollow)
	Register(asm.BuiltinCreatureHunt, nsHunt)
	Register(asm.BuiltinWalk, nsWalkTo)
	Register(asm.BuiltinCreateMover, nsCreateMover)
	Register(asm.BuiltinIsLocked, nsIsLocked)
	Register(asm.BuiltinUnlockDoor, nsUnlock)
	Register(asm.BuiltinLockDoor, nsLock)
	Register(asm.BuiltinGetLastItem, nsFirstItem)
	Register(asm.BuiltinGetPreviousItem, nsNextItem)
	Register(asm.BuiltinHasItem, hasItem)
	Register(asm.BuiltinGetHolder, nsGetHolder)
	Register(asm.BuiltinRaiseZombie, nsRaiseZombie)
	Register(asm.BuiltinZombieStayDown, nsZombieStayDown)
	Register(asm.BuiltinGetElevatorStatus, nsGetElevatorStat)
	Register(asm.BuiltinMove, nsMove)
	Register(asm.BuiltinDamage, nsDamage)
	Register(asm.BuiltinGetHost, nsGetHost)
	Register(asm.BuiltinClearMessages, nsClearMessages)
	Register(asm.BuiltinIsAttackedBy, nsIsAttackedBy)
	Register(asm.BuiltinChat, nsChat)
	Register(asm.BuiltinChatTimer, nsChatTimerFrames)
	Register(asm.BuiltinChatTimerSeconds, nsChatTimerSeconds)
	Register(asm.BuiltinDestroyChat, nsDestroyChat)
	Register(asm.BuiltinDestroyEveryChat, nsDestroyEveryChat)
	Register(asm.BuiltinIsVisibleTo, nsIsVisibleTo)
	Register(asm.BuiltinAggressionLevel, nsSetAggressionLevel)
	Register(asm.BuiltinRetreatLevel, nsSetRetreatLevel)
	Register(asm.BuiltinResumeLevel, nsSetResumeLevel)
	Register(asm.BuiltinIsGameBall, nsIsGameBall)
	Register(asm.BuiltinIsCrown, nsIsCrown)
	Register(asm.BuiltinIsSummoned, nsIsSummoned)
	Register(asm.BuiltinGetGold, nsGetGold)
	Register(asm.BuiltinChangeGold, nsChangeGold)
	Register(asm.BuiltinSetDialog, nsSetDialog)
	Register(asm.BuiltinStartDialog, nsStartDialog)
	Register(asm.BuiltinGetAnswer, nsGetAnswer)
	Register(asm.BuiltinCancelDialog, nsCancelDialog)
	Register(asm.BuiltinStoryPic, nsStoryPic)
	Register(asm.BuiltinTellStory, nsTellStory)
	Register(asm.BuiltinAttack, nsAttack)
	Register(asm.BuiltinRunAway, nsRunAway)
	Register(asm.BuiltinCreatureGuard, nsGuard)
	Register(asm.BuiltinPauseObject, nsPause)
	Register(asm.BuiltinSetCallback, nsSetCallback)
	Register(asm.BuiltinHitLocation, nsHitLocation)
	Register(asm.BuiltinHitFarLocation, nsHitFarLocation)
	Register(asm.BuiltinSetShopkeeperText, nsSetShopkeeperText)
}

func nsGetTrigger(vm VM) int {
	vm.PushHandleNS(vm.NoxScript().GetTrigger())
	return 0
}

func nsGetCaller(vm VM) int {
	vm.PushHandleNS(vm.NoxScript().GetCaller())
	return 0
}

func nsIsTrigger(vm VM) int {
	obj := vm.PopObjectNS()
	if obj == nil {
		vm.PushBool(false)
		return 0
	}
	vm.PushBool(vm.NoxScript().IsTrigger(obj))
	return 0
}

func nsIsCaller(vm VM) int {
	obj := vm.PopObjectNS()
	if obj == nil {
		vm.PushBool(false)
		return 0
	}
	vm.PushBool(vm.NoxScript().IsCaller(obj))
	return 0
}

func nsObject(vm VM) int {
	str := vm.PopString()
	str += vm.NameSuff()
	vm.PushHandleNS(vm.NoxScript().Object(str))
	return 0
}

func nsCreateObject(s VM) int {
	wp := s.PopWaypointNS()
	typ := s.PopString()
	if wp == nil {
		s.PushI32(0)
		return 0
	}
	s.PushHandleNS(s.NoxScript().CreateObject(typ, wp))
	return 0
}

func nsObjectX(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		s.PushF32(obj.Pos().X)
	} else {
		s.PushF32(0)
	}
	return 0
}

func nsObjectY(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		s.PushF32(obj.Pos().Y)
	} else {
		s.PushF32(0)
	}
	return 0
}

func nsSetPos(s VM) int {
	p := s.PopPointf()
	obj := s.PopObjectNS()
	dpos := s.DPosf()
	p = p.Add(dpos)
	if obj != nil {
		obj.SetPos(p)
	}
	return 0
}

func nsZ(s VM) int {
	obj := s.PopObjectNS()
	if obj != nil {
		s.PushF32(obj.Z())
	} else {
		s.PushI32(0)
	}
	return 0
}

func nsSetZ(s VM) int {
	z := s.PopF32()
	obj := s.PopObjectNS()
	if obj != nil {
		obj.SetZ(z)
	}
	return 0
}

func nsIsEnabled(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		s.PushBool(obj.IsEnabled())
	} else {
		s.PushBool(false)
	}
	return 0
}

func nsObjectOn(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.Enable(true)
	}
	return 0
}

func nsObjectOff(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.Enable(false)
	}
	return 0
}

func nsObjectToggle(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.Toggle()
	}
	return 0
}

func nsHasClass(s VM) int {
	name := class.Class(s.PopString())
	obj := s.PopObjectNS()
	s.PushBool(obj != nil && obj.HasClass(name))
	return 0
}

func nsHasSubclass(s VM) int {
	name := subclass.SubClass(s.PopString())
	obj := s.PopObjectNS()
	s.PushBool(obj != nil && obj.HasSubclass(name))
	return 0
}

func nsCurrentHealth(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		s.PushI32(int32(obj.CurrentHealth()))
	} else {
		s.PushI32(0)
	}
	return 0
}

func nsMaxHealth(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		s.PushI32(int32(obj.MaxHealth()))
	} else {
		s.PushI32(0)
	}
	return 0
}

func nsRestoreHealth(s VM) int {
	val := int(s.PopI32())
	obj := s.PopObjectNS()
	if obj != nil {
		obj.RestoreHealth(val)
	}
	return 0
}

func nsGetDirection(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		s.PushI32(int32(obj.Direction()))
	} else {
		s.PushI32(0)
	}
	return 0
}

func nsLookWithAngle(s VM) int {
	angle := int(s.PopI32())
	obj := s.PopObjectNS()
	if obj != nil {
		obj.LookWithAngle(angle)
	}
	return 0
}

func nsLookAtObject(s VM) int {
	targ := s.PopObjectNS()
	obj := s.PopObjectNS()
	if obj != nil {
		obj.LookAtObject(targ)
	}
	return 0
}

func nsLookAtDirection(s VM) int {
	dir := s.PopI32()
	obj := s.PopObjectNS()
	if obj != nil {
		obj.LookAtDirection(ns4.Direction(dir))
	}
	return 0
}

func nsDelete(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.Delete()
	}
	return 0
}

func nsDeleteObjectTimer(s VM) int {
	frames := s.PopU32()
	obj := s.PopObjectNS()
	if obj != nil {
		obj.DeleteAfter(ns4.Frames(int(frames)))
	}
	return 0
}

func nsPushTo(s VM) int {
	vec := s.PopPointf()
	obj := s.PopObjectNS()
	if obj != nil {
		obj.ApplyForce(vec)
	}
	return 0
}

func nsPush(s VM) int {
	p := s.PopPointf()
	force := s.PopF32()
	obj := s.PopObjectNS()
	dpos := s.DPosf()
	p = p.Add(dpos)
	if obj != nil {
		obj.PushTo(p, force)
	}
	return 0
}

func nsFreeze(s VM) int {
	val := s.PopBool()
	if obj := s.PopObjectNS(); obj != nil {
		obj.Freeze(val)
	}
	return 0
}

func nsPickup(s VM) int {
	item := s.PopObjectNS()
	picker := s.PopObjectNS()
	if picker != nil {
		s.PushBool(picker.Pickup(item))
	} else {
		s.PushBool(false)
	}
	return 0
}

func nsDrop(s VM) int {
	item := s.PopObjectNS()
	holder := s.PopObjectNS()
	if holder != nil {
		s.PushBool(holder.Drop(item))
	} else {
		s.PushBool(false)
	}
	return 0
}

func nsIsOwnedBy(s VM) int {
	owner := s.PopObjectNS()
	obj := s.PopObjectNS()
	if obj != nil {
		s.PushBool(obj.HasOwner(owner))
	} else {
		s.PushBool(false)
	}
	return 0
}

func nsClearOwner(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.SetOwner(nil)
	}
	return 0
}

func nsSetOwner(s VM) int {
	obj := s.PopObjectNS()
	owner := s.PopObjectNS()
	if obj != nil {
		obj.SetOwner(owner)
	}
	return 0
}

func nsSetOwners(s VM) int {
	obj := s.PopObjectNS()
	g := s.PopObjGroupNS()
	if obj != nil {
		obj.SetOwners(g)
	}
	return 0
}

func nsIsOwnedByAny(s VM) int {
	g := s.PopObjGroupNS()
	obj := s.PopObjectNS()
	if obj != nil {
		s.PushBool(obj.HasOwnerIn(g))
	} else {
		s.PushBool(false)
	}
	return 0
}

func nsReturn(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.Return()
	}
	return 0
}

func nsIdle(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.Idle()
	}
	return 0
}

func nsWander(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.Wander()
	}
	return 0
}

func nsFollow(s VM) int {
	targ := s.PopObjectNS()
	obj := s.PopObjectNS()
	if obj != nil && targ != nil {
		obj.Follow(targ)
	}
	return 0
}

func nsHunt(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.Hunt()
	}
	return 0
}

func nsWalkTo(s VM) int {
	p := s.PopPointf()
	obj := s.PopObjectNS()
	dpos := s.DPosf()
	p = p.Add(dpos)
	if obj != nil {
		obj.WalkTo(p)
	}
	return 0
}

func nsCreateMover(s VM) int {
	speed := s.PopF32()
	wp := s.PopWaypointNS()
	obj := s.PopObjectNS()
	if obj != nil {
		s.PushHandleNS(obj.CreateMover(wp, speed))
	} else {
		s.PushHandleNS(nil)
	}
	return 0
}

func nsUnlock(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.Lock(false)
	}
	return 0
}

func nsLock(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.Lock(true)
	}
	return 0
}

func nsFirstItem(s VM) int {
	obj := s.PopObjectNS()
	if obj == nil {
		s.PushHandleNS(nil)
		return 0
	}
	s.PushHandleNS(obj.GetLastItem())
	return 0
}

func nsNextItem(s VM) int {
	obj := s.PopObjectNS()
	if obj == nil {
		s.PushHandleNS(nil)
		return 0
	}
	s.PushHandleNS(obj.GetPreviousItem())
	return 0
}

func hasItem(s VM) int {
	item := s.PopObjectNS()
	obj := s.PopObjectNS()
	if obj != nil {
		s.PushBool(obj.HasItem(item))
	} else {
		s.PushBool(false)
	}
	return 0
}

func nsGetHolder(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		s.PushHandleNS(obj.GetHolder())
	} else {
		s.PushHandleNS(nil)
	}
	return 0
}

func nsRaiseZombie(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.RaiseZombie()
	}
	return 0
}

func nsZombieStayDown(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.ZombieStayDown()
	}
	return 0
}

func nsGetElevatorStat(s VM) int {
	obj := s.PopObjectNS()
	if obj != nil {
		s.PushI32(int32(obj.GetElevatorStatus()))
	} else {
		s.PushI32(-1)
	}
	return 0
}

func nsMove(s VM) int {
	wp := s.PopWaypointNS()
	obj := s.PopObjectNS()
	if obj != nil {
		obj.Move(wp)
	}
	return 0
}

func nsIsLocked(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		s.PushBool(obj.IsLocked())
	} else {
		s.PushBool(false)
	}
	return 0
}

func nsDamage(s VM) int {
	typ := damage.Type(s.PopI32())
	dmg := int(s.PopI32())
	src := s.PopObjectNS()
	dest := s.PopObjectNS()
	if dest != nil {
		dest.Damage(src, dmg, typ)
	}
	return 0
}

func nsGetHost(s VM) int {
	if u := s.NoxScript().GetHost(); u != nil {
		s.PushI32(int32(u.ScriptID()))
	} else {
		s.PushI32(0)
	}
	return 0
}

func nsClearMessages(s VM) int {
	if pl := s.PopObjectNS(); pl != nil {
		s.NoxScript().ClearMessages(pl)
	}
	return 0
}

func nsIsAttackedBy(s VM) int {
	obj2 := s.PopObjectNS()
	obj1 := s.PopObjectNS()
	val := obj1 != nil && obj1.IsAttackedBy(obj2)
	s.PushBool(val)
	return 0
}

func nsChat(s VM) int {
	msg := s.PopString()
	obj := s.PopObjectNS()
	if obj != nil {
		obj.Chat(msg)
	}
	return 0
}

func nsChatTimerSeconds(s VM) int {
	sec := s.PopU32()
	msg := s.PopString()
	obj := s.PopObjectNS()
	if obj != nil {
		obj.ChatTimer(msg, ns4.Seconds(float64(sec)))
	}
	return 0
}

func nsChatTimerFrames(s VM) int {
	frames := int(s.PopU32())
	msg := s.PopString()
	obj := s.PopObjectNS()
	if obj != nil {
		obj.ChatTimer(msg, ns4.Frames(frames))
	}
	return 0
}

func nsDestroyChat(s VM) int {
	if obj := s.PopObjectNS(); obj != nil {
		obj.DestroyChat()
	}
	return 0
}

func nsDestroyEveryChat(s VM) int {
	s.NoxScript().DestroyEveryChat()
	return 0
}

func nsIsVisibleTo(s VM) int {
	obj2 := s.PopObjectNS()
	obj := s.PopObjectNS()
	ok := obj != nil && obj.CanSee(obj2)
	s.PushBool(ok)
	return 0
}

func nsSetAggressionLevel(vm VM) int {
	v := vm.PopF32()
	obj := vm.PopObjectNS()
	if obj != nil {
		obj.AggressionLevel(v)
	}
	return 0
}

func nsSetRetreatLevel(vm VM) int {
	v := vm.PopF32()
	obj := vm.PopObjectNS()
	if obj != nil {
		obj.RetreatLevel(v)
	}
	return 0
}

func nsSetResumeLevel(vm VM) int {
	v := vm.PopF32()
	obj := vm.PopObjectNS()
	if obj != nil {
		obj.ResumeLevel(v)
	}
	return 0
}

func nsIsGameBall(vm VM) int {
	obj := vm.PopObjectNS()
	ok := obj != nil && vm.NoxScript().IsGameBall(obj)
	vm.PushBool(ok)
	return 0
}

func nsIsCrown(vm VM) int {
	obj := vm.PopObjectNS()
	ok := obj != nil && vm.NoxScript().IsCrown(obj)
	vm.PushBool(ok)
	return 0
}

func nsIsSummoned(vm VM) int {
	obj := vm.PopObjectNS()
	ok := obj != nil && vm.NoxScript().IsSummoned(obj)
	vm.PushBool(ok)
	return 0
}

func nsGetGold(vm VM) int {
	obj := vm.PopObjectNS()
	var v int
	if obj != nil {
		v = obj.GetGold()
	}
	vm.PushI32(int32(v))
	return 0
}

func nsChangeGold(vm VM) int {
	v := int(vm.PopI32())
	obj := vm.PopObjectNS()
	if obj != nil {
		obj.ChangeGold(v)
	}
	return 0
}

func nsSetDialog(vm VM) int {
	end := vm.PopI32()
	start := vm.PopI32()
	name := vm.PopString()
	obj := vm.PopObjectNS()
	if obj != nil {
		vm.NoxScript().SetDialog(obj, ns4.DialogType(name), int(start), int(end))
	}
	return 0
}

func nsStartDialog(vm VM) int {
	obj2 := vm.PopObjectNS()
	obj1 := vm.PopObjectNS()
	if obj1 != nil && obj2 != nil {
		vm.NoxScript().StartDialog(obj1, obj2)
	}
	return 0
}

func nsGetAnswer(vm VM) int {
	obj := vm.PopObjectNS()
	var res int32
	if obj != nil {
		res = int32(vm.NoxScript().GetAnswer(obj))
	}
	vm.PushI32(res)
	return 0
}

func nsCancelDialog(vm VM) int {
	obj := vm.PopObjectNS()
	if obj != nil {
		vm.NoxScript().CancelDialog(obj)
	}
	return 0
}

func nsStoryPic(vm VM) int {
	name := vm.PopString()
	obj := vm.PopObjectNS()
	if obj != nil {
		vm.NoxScript().StoryPic(obj, name)
	}
	return 0
}

func nsTellStory(vm VM) int {
	v0 := vm.PopString()
	v1 := vm.PopString()
	vm.NoxScript().TellStory(audio.Name(v1), ns4.StringID(v0))
	return 0
}

func nsAttack(vm VM) int {
	targ := vm.PopObjectNS()
	obj := vm.PopObjectNS()
	if obj != nil {
		obj.Attack(targ)
	}
	return 0
}

func nsRunAway(vm VM) int {
	dt := vm.PopU32()
	targ := vm.PopObjectNS()
	obj := vm.PopObjectNS()
	if targ != nil && obj != nil {
		obj.Flee(targ, ns4.Frames(int(dt)))
	}
	return 0
}

func nsGuard(vm VM) int {
	dist := vm.PopF32()
	p2 := vm.PopPointf()
	p1 := vm.PopPointf()
	obj := vm.PopObjectNS()
	if obj != nil {
		obj.Guard(p1, p2, dist)
	}
	return 0
}

func nsPause(vm VM) int {
	dt := vm.PopU32()
	obj := vm.PopObjectNS()
	if obj != nil {
		obj.Pause(ns4.Frames(int(dt)))
	}
	return 0
}

func nsSetCallback(vm VM) int {
	fnc := int32(vm.PopU32())
	ev := ns4.ObjectEvent(vm.PopU32())
	obj := vm.PopObjectNS()
	if obj != nil {
		obj.OnEvent(ev, int(fnc))
	}
	return 0
}

func nsHitLocation(vm VM) int {
	pos := vm.PopPointf()
	obj := vm.PopObjectNS()
	if obj != nil {
		obj.HitMelee(pos)
	}
	return 0
}

func nsHitFarLocation(vm VM) int {
	pos := vm.PopPointf()
	obj := vm.PopObjectNS()
	if obj != nil {
		obj.HitRanged(pos)
	}
	return 0
}

func nsSetShopkeeperText(vm VM) int {
	str := vm.PopString()
	obj := vm.PopObjectNS()
	if obj != nil {
		obj.SetShopText(ns4.StringID(str))
	}
	return 0
}
