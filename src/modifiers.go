package opennox

import (
	"sync"
	"unsafe"

	"github.com/opennox/libs/spell"

	"github.com/opennox/opennox/v1/common/memmap"
	"github.com/opennox/opennox/v1/legacy"
	"github.com/opennox/opennox/v1/legacy/common/alloc"
	"github.com/opennox/opennox/v1/server"
)

var modifierColorsOnce sync.Once

func sub_4A5E90_A() {
	modifierColorsOnce.Do(sub_4A62B0)
	for it := noxServer.Modif.Dword_5d4594_251608; it != nil; it = it.Next80 {
		switch it.Name() {
		case "StreetPants":
			legacy.Set_dword_5d4594_1308156(it.C())
		case "StreetShirt":
			legacy.Set_dword_5d4594_1308160(it.C())
		case "StreetSneakers":
			legacy.Set_dword_5d4594_1308164(it.C())
		}
	}
}

func sub_4A62B0() {
	s := noxServer
	arr := memmap.PtrT[[3][32]server.ModColor](0x5D4594, 1307796)
	copy(arr[:], s.Modif.Colors[:])
}

func nox_xxx_fireRingEffect_4E05B0(a1 unsafe.Pointer, a2p, src, a4p *server.Object) {
	if src != nil {
		sa, free := alloc.New(server.SpellAcceptArg{})
		defer free()
		sa.Obj = nil
		sa.Pos = src.PosVec
		legacy.Nox_xxx_spellCastCleansingFlame_52D5C0(spell.SPELL_CLEANSING_FLAME, src, src, src, sa, int(*(*uint32)(unsafe.Add(a1, 48))))
	}
}

func nox_xxx_blueFREffect_4E05F0(a1 unsafe.Pointer, a2p, src, a4p *server.Object) {
	if src != nil {
		sa, free := alloc.New(server.SpellAcceptArg{})
		defer free()
		sa.Obj = nil
		sa.Pos = src.PosVec
		legacy.Nox_xxx_spellCastCleansingFlame_52D5C0(spell.SPELL_CLEANSING_MANA_FLAME, src, src, src, sa, int(*(*uint32)(unsafe.Add(a1, 48))))
	}
}
