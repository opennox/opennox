package server

import "github.com/opennox/noxscript/ns/asm"

func (s *NoxScriptVM) callBuiltin(sc *ScriptFunc, fi asm.Builtin) error {
	if fi < 0 || fi > asm.BuiltinGetScore {
		if s.panicCompilerCheck(fi) {
			return nil
		}
	}
	s.vm.nameSuff = sc.NamePref
	if s.builtinNeedsDPos(fi) {
		s.vm.dpos = sc.PosOff
	}
	err := s.CallBuiltinNative(fi)
	s.ResetBuiltin()
	return err
}

func (s *NoxScriptVM) builtinNeedsDPos(fi asm.Builtin) bool {
	// TODO: 7 items in the array, but the count is set to 5; why?
	var check = []asm.Builtin{
		asm.BuiltinWall,
		asm.BuiltinMoveObject, asm.BuiltinMoveWaypoint,
		asm.BuiltinPushObject,
		asm.BuiltinWalk, asm.BuiltinGroupWalk,
		asm.BuiltinEffect,
	}
	for _, ind := range check[:5] {
		if fi == ind {
			return true
		}
	}
	return false
}
