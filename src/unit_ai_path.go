package opennox

import (
	"image"
	"math"
	"unsafe"

	"github.com/opennox/libs/types"
	"github.com/opennox/libs/wall"

	noxflags "github.com/opennox/opennox/v1/common/flags"
	"github.com/opennox/opennox/v1/common/memmap"
	"github.com/opennox/opennox/v1/common/ntype"
	"github.com/opennox/opennox/v1/common/unit/ai"
	"github.com/opennox/opennox/v1/legacy"
	"github.com/opennox/opennox/v1/server"
)

func (s *Server) nox_xxx_pathFind_50BA00(far bool, obj *server.Object, a3 *types.Pointf, a4 *types.Pointf, fnc func(obj *server.Object, x, y int) bool, a6 int32) {
	var (
		v30 *server.Object
		v33 *server.AIVisitNode
		v36 int32
		v37 int32
		v38 bool
		v39 int32
		v40 int32
		v41 int32
		v42 int32
		v64 int32
		v65 int32
		v67 *server.AIVisitNode
		v68 [2]uint16
		v77 *server.AIVisitNode
	)
	v67 = nil
	v64 = 0
	v65 = 0
	var cur int32
	if far {
		cur = 0
	} else {
		cur = 999999
	}
	s.AI.Paths.PathStatus = 0
	s.AI.Paths.MapIndexLast++
	s.AI.Paths.MaybeIndexObjects()
	var a2a ntype.Point32
	a2a.X = int32(float32(float64(a3.X) / 23))
	a2a.Y = int32(float32(float64(a3.Y) / 23))
	s.nox_xxx_pathfind_preCheckWalls_50C8D0(obj, &a2a)
	v63 := bool2int(!s.AI.Paths.Nox_xxx_pathfind_preCheckWalls2_50B8A0(obj, int(a2a.X), int(a2a.Y)))
	v61 := false
	if fnc != nil && !fnc(obj, int(a2a.X), int(a2a.Y)) {
		v61 = true
	}
	x00 := int32(float64(a4.X) / 23)
	y00 := int32(float64(a4.Y) / 23)
	if !s.AI.Paths.Valid() {
		s.AI.Paths.ResetPoints()
		s.AI.Paths.PathStatus = 2
		return
	}
	s.AI.Paths.ResetVisitNodes()
	vn := s.AI.Paths.NewVisitNode()
	vn.X0 = uint16(int16(a2a.X))
	vn.Y2 = uint16(int16(a2a.Y))
	x0 := vn.X0
	y0 := vn.Y2
	vn.Field4 = nil
	vn.Field8 = nil
	vn2 := vn

	s.AI.Paths.MapIndex(int(x0), int(y0)).Index0 = s.AI.Paths.MapIndexLast
	for v66 := int32(0); ; {
		vn3 := vn2
		vn2 = nil
		v77 = nil
		if vn3 == nil {
			goto LABEL_40
		}
		for {
			x1 := int32(vn3.X0)
			y1 := int32(vn3.Y2)
			dd := (x1-x00)*(x1-x00) + (y1-y00)*(y1-y00)
			if far {
				if dd > cur {
					v67 = vn3
					cur = dd
				}
			} else {
				if dd < cur {
					v67 = vn3
					cur = dd
				}
				if v63 == 0 && !v61 && x1 == x00 && y1 == y00 {
					s.AI.Paths.PathStatus = 0
					s.AI.Paths.Sub_50C320(obj, vn3, a4)
					return
				}
			}
			ioff := s.Rand.Logic.IntClamp(0, 7)
			parr := memmap.PtrT[[8]ntype.Point32](0x587000, 234216)
			for i := 0; i < 8; i++ {
				ii := (i + ioff) % 8
				pp := parr[ii]
				x2 := pp.X + int32(vn3.X0)
				y2 := pp.Y + int32(vn3.Y2)
				ip := s.AI.Paths.MapIndex(int(x2), int(y2))
				if ip == nil || ip.Index0 == s.AI.Paths.MapIndexLast {
					continue
				}
				ip.Index0 = s.AI.Paths.MapIndexLast
				if x2 == x00 && y2 == y00 {
					var v76a, v76b types.Pointf
					v76a.X = float32(int32(vn3.X0)*23 + 11)
					v76a.Y = float32(int32(vn3.Y2)*23 + 11)
					v76b.X = a4.X
					v76b.Y = a4.Y
					if s.MapTraceObstacles(obj, v76a, v76b) {
						if obj.ObjFlags&0x4000 != 0 {
							if s.MapTraceRayAt(v76a, v76b, nil, nil, 5) {
								v30 = obj
								goto LABEL_32
							}
						} else {
							if s.MapTraceRayAt(v76a, v76b, nil, nil, 1) {
								v30 = obj
								goto LABEL_32
							}
						}
					}
				}
				if ii < 4 {
					v30 = obj
				} else {
					v60 := byte(int8(int32(uint16(int16(^(int32(*(*uint16)(unsafe.Add(unsafe.Pointer(obj), 16)))>>8))))&0xD8 | 0x98))
					s.Doors.SetKeyHolder(obj)
					switch ii {
					case 4:
						v36 = int32(uint8(s.Sub_57B500(image.Pt(int(vn3.X0), int(vn3.Y2)+1), v60)))
						if v36 != 1 && v36 != 6 && v36 != 10 && v36 != 9 && v36 != math.MaxUint8 {
							continue
						}
						*(*uint8)(unsafe.Pointer(&v37)) = uint8(s.Sub_57B500(image.Pt(int(vn3.X0)+1, int(vn3.Y2)), v60))
						v37 = int32(uint8(int8(v37)))
						if int32(uint8(int8(v37))) == 1 || v37 == 4 || v37 == 7 {
							// nop
						} else {
							v38 = v37 == 8
							if !v38 && v37 != math.MaxUint8 {
								continue
							}
						}
					case 5:
						v39 = int32(uint8(s.Sub_57B500(image.Pt(int(vn3.X0)-1, int(vn3.Y2)), v60)))
						if v39 != 1 && v39 != 6 && v39 != 10 && v39 != 9 && v39 != math.MaxUint8 {
							continue
						}
						*(*uint8)(unsafe.Pointer(&v37)) = uint8(s.Sub_57B500(image.Pt(int(vn3.X0), int(vn3.Y2)-1), v60))
						v37 = int32(uint8(int8(v37)))
						if int32(uint8(int8(v37))) == 1 || v37 == 4 || v37 == 7 {
							// nop
						} else {
							v38 = v37 == 8
							if !v38 && v37 != math.MaxUint8 {
								continue
							}
						}
					case 6:
						v40 = int32(uint8(s.Sub_57B500(image.Pt(int(vn3.X0), int(vn3.Y2)+1), v60)))
						if v40 != 0 && v40 != 5 && v40 != 9 && v40 != 8 && v40 != math.MaxUint8 {
							continue
						}
						v37 = int32(uint8(s.Sub_57B500(image.Pt(int(vn3.X0)-1, int(vn3.Y2)), v60)))
						if v37 == 0 || v37 == 3 || v37 == 10 {
							// nop
						} else {
							v38 = v37 == 7
							if !v38 && v37 != math.MaxUint8 {
								continue
							}
						}
					case 7:
						v41 = int32(uint8(s.Sub_57B500(image.Pt(int(vn3.X0)+1, int(vn3.Y2)), v60)))
						if v41 != 0 && v41 != 5 && v41 != 9 && v41 != 8 && v41 != math.MaxUint8 {
							continue
						}
						v42 = int32(uint8(s.Sub_57B500(image.Pt(int(vn3.X0), int(vn3.Y2)-1), v60)))
						if v42 != 0 && v42 != 3 && v42 != 10 && v42 != 7 && v42 != math.MaxUint8 {
							continue
						}
					}
					v30 = obj
					if s.AI.Paths.CheckIndexFlags(obj, int(vn3.X0), int(pp.Y)+int(vn3.Y2)) || s.AI.Paths.CheckIndexFlags(obj, int(pp.X)+int(vn3.X0), int(vn3.Y2)) {
						continue
					}
				}
				if s.sub_50B870(v30, int(x2), int(y2)) {
					continue
				}
				if x00 != x2 || y00 != y2 {
					if sub_50C830(v30, x2, y2) == 0 {
						continue
					}
				} else if sub_50C830(v30, x2, y2) == 0 {
					continue
				}
				if v63 != 0 {
					v43 := s.AI.Paths.Nox_xxx_pathfind_preCheckWalls2_50B8A0(v30, int(x2), int(y2))
					if v43 {
						v64 = 1
					} else if !s.AI.Paths.CheckIndexFlags(v30, int(vn3.X0), int(vn3.Y2)) && s.AI.Paths.CheckIndexFlags(v30, int(x2), int(y2)) { // TODO: why condition is asymmetric?
						continue
					}
				} else {
					v43 := s.AI.Paths.Nox_xxx_pathfind_preCheckWalls2_50B8A0(v30, int(x2), int(y2))
					if !v43 {
						continue
					}
					if fnc != nil {
						v44 := fnc(v30, int(x2), int(y2))
						if v61 {
							if v44 {
								v65 = 1
							}
						} else if !v44 {
							continue
						}
					}
				}
			LABEL_32:
				v31 := s.AI.Paths.NewVisitNode()
				if v31 == nil {
					if noxflags.HasEngine(noxflags.EngineShowAI) {
						ai.Log.Printf("%d: %s(#%d), buildPath: Exhausted search storage\n", s.Frame(), v30, v30.NetCode)
					}
					s.AI.Paths.PathStatus = 1
					s.AI.Paths.Sub_50C320(v30, v67, nil)
					return
				}
				v31.X0 = uint16(int16(x2))
				v31.Y2 = uint16(int16(y2))
				v31.Field4 = vn3
				v31.Field8 = v77
				v77 = v31
			}
			if s.AI.Paths.Sub_50AC20(vn3, &v68) != 0 {
				vn3.Flags12 |= 2
				v32 := s.AI.Paths.NewVisitNode()
				if v32 == nil {
					if noxflags.HasEngine(noxflags.EngineShowAI) {
						ai.Log.Printf("%d: %s(#%d), buildPath: Exhausted search storage\n", s.Frame(), obj, obj.NetCode)
					}
					s.AI.Paths.PathStatus = 1
					s.AI.Paths.Sub_50C320(obj, v67, nil)
					return
				}
				v32.X0 = v68[0]
				v32.Y2 = v68[1]
				v32.Field4 = vn3
				v32.Field8 = v77
				v77 = v32
			}
			vn3 = vn3.Field8
			if vn3 == nil {
				break
			}
		}
		vn2 = v77
	LABEL_40:
		v33 = nil
		if v64 != 0 {
			v34 := vn2
			if vn2 != nil {
				for {
					v35 := v34.Field8
					if s.AI.Paths.Nox_xxx_pathfind_preCheckWalls2_50B8A0(obj, int(v34.X0), int(v34.Y2)) {
						v33 = v34
					} else if v33 != nil {
						v33.Field8 = v34.Field8
					} else {
						vn2 = v34.Field8
					}
					v34 = v35
					if v35 == nil {
						break
					}
				}
				v33 = nil
			}
			v63 = 0
			v64 = 0
		}
		if v65 != 0 {
			v45 := vn2
			if vn2 != nil {
				for {
					v46 := v45.Field8
					if fnc(obj, int(v45.X0), int(v45.Y2)) {
						v33 = v45
					} else if v33 != nil {
						v33.Field8 = v45.Field8
					} else {
						vn2 = v45.Field8
					}
					v45 = v46
					if v46 == nil {
						break
					}
				}
			}
			v61 = false
			v65 = 0
		}
		if vn2 == nil {
			break
		}
		if a6 != 0 && v66 >= a6 {
			if noxflags.HasEngine(noxflags.EngineShowAI) {
				ai.Log.Printf("%d: %s(#%d), buildPath: Reached search depth limit\n", s.Frame(), obj, obj.NetCode)
			}
			s.AI.Paths.PathStatus = 2
			s.AI.Paths.Sub_50C320(obj, v67, nil)
			return
		}
		v66++
	}
	if noxflags.HasEngine(noxflags.EngineShowAI) {
		ai.Log.Printf("%d: %s(#%d), buildPath: Exhausted search space\n", s.Frame(), obj, obj.NetCode)
	}
	s.AI.Paths.PathStatus = 2
	s.AI.Paths.Sub_50C320(obj, v67, nil)
}

func (s *Server) Sub_50CB20(a1 *server.Object, a2 *types.Pointf) *server.Waypoint {
	s.AI.Paths.MapIndexLast++
	var a2a ntype.Point32
	a2a.X = int32(float64(a2.X) / 23)
	a2a.Y = int32(float64(a2.Y) / 23)
	s.nox_xxx_pathfind_preCheckWalls_50C8D0(a1, &a2a)
	s.AI.Paths.ResetVisitNodes()
	v4 := s.AI.Paths.NewVisitNode()
	v4.X0 = uint16(int16(a2a.X))
	v4.Y2 = uint16(int16(a2a.Y))
	v4.Field4 = nil
	v4.Field8 = nil
	s.AI.Paths.MapIndex(int(v4.X0), int(v4.Y2)).Index0 = s.AI.Paths.MapIndexLast
	for {
		var v19 *server.AIVisitNode
		if v4 == nil {
			break
		}
		for v6 := v4; v6 != nil; v6 = v6.Field8 {
			v7 := int32(v6.X0)
			v8 := int32(v6.Y2)
			if int32(s.AI.Paths.MapIndex(int(v7), int(v8)).Flags8)&0x40 != 0 && !s.sub_50B870(a1, int(v7), int(v8)) {
				var v18 types.Pointf
				v18.X = float32(float64(v7*23 + 11))
				v18.Y = float32(float64(v8*23 + 11))
				return s.Sub_518460(v18, 0x80, true)
			}
			v9 := memmap.PtrOff(0x587000, 234284)
			for {
				x2 := int32(uint32(v6.X0) + *(*uint32)(unsafe.Add(v9, -int(4*1))))
				y2 := int32(*(*uint32)(unsafe.Pointer(v9)) + uint32(v6.Y2))
				if p := s.AI.Paths.MapIndex(int(x2), int(y2)); p != nil && p.Index0 != s.AI.Paths.MapIndexLast {
					v16 := int32(*(*uint32)(unsafe.Pointer(v9)) + uint32(v6.Y2))
					p.Index0 = s.AI.Paths.MapIndexLast
					if !s.sub_50B870(a1, int(x2), int(v16)) {
						if sub_50C830(a1, x2, y2) != 0 {
							v12 := s.AI.Paths.NewVisitNode()
							if v12 != nil {
								v12.X0 = uint16(int16(x2))
								v12.Y2 = uint16(int16(y2))
								v12.Field4 = v6
								v12.Field8 = v19
								v19 = v12
							}
						}
					}
				}
				v9 = unsafe.Add(v9, 8)
				if uintptr(v9) >= uintptr(memmap.PtrOff(0x587000, 234316)) {
					break
				}
			}
		}
		v4 = v19
		if v19 == nil {
			break
		}
	}
	return nil
}

func (s *Server) sub_50B870(a1 *server.Object, x, y int) bool {
	return s.sub_57B630(a1, x, y) != -1
}
func (s *Server) sub_57B630(obj *server.Object, x, y int) int8 {
	if x < 0 || x >= 256 || y < 0 || y >= 256 {
		return -1
	}
	wl := s.Walls.GetWallAtGrid2(image.Pt(x, y))
	if wl == nil {
		return -1
	}
	if wl.Flags4.Has(wall.FlagDoor) {
		door := asObject(wl.Data).SObj()
		if door != nil {
			ud := door.UpdateData
			v7 := *(*uint32)(unsafe.Add(ud, 12))
			if v7 == *(*uint32)(unsafe.Add(ud, 4)) {
				v8x := memmap.Int32(0x587000, uintptr(v7)*8+196184)
				v8y := memmap.Int32(0x587000, uintptr(v7)*8+196188)
				if v8x > 0 {
					if v8y > 0 {
						if door.ObjOwner != nil {
							if s.Rand.Logic.IntClamp(0, 100) >= 50 {
								return 1
							}
						} else {
							if int32(*(*uint8)(unsafe.Add(ud, 1))) != 0 && s.DoorCheckKey(obj, door) == nil {
								return 1
							}
						}
						return -1
					} else if v8y < 0 {
						if door.ObjOwner != nil {
							if s.Rand.Logic.IntClamp(0, 100) >= 50 {
								return 0
							}
						} else {
							if int32(*(*uint8)(unsafe.Add(ud, 1))) != 0 && s.DoorCheckKey(obj, door) == nil {
								return 0
							}
						}
						return -1
					}
				} else if v8x < 0 {
					if v8y < 0 {
						if door.ObjOwner != nil {
							if s.Rand.Logic.IntClamp(0, 100) >= 50 {
								return 1
							}
						} else {
							if int32(*(*uint8)(unsafe.Add(ud, 1))) != 0 && s.DoorCheckKey(obj, door) == nil {
								return 1
							}
						}
						return -1
					} else if v8y > 0 {
						if door.ObjOwner != nil {
							if s.Rand.Logic.IntClamp(0, 100) >= 50 {
								return 0
							}
						} else {
							if int32(*(*uint8)(unsafe.Add(ud, 1))) != 0 && s.DoorCheckKey(obj, door) == nil {
								return 0
							}
						}
						return -1
					}
				}
			}
		}
	} else if (obj.ObjFlags&0x4000) == 0 || !wl.Flags4.Has(wall.FlagWindow) {
		if !wl.Flags4.Has(wall.FlagSecret) {
			return int8(wl.Dir0)
		}
		v13 := wl.Data
		if (int32(*(*uint8)(unsafe.Add(v13, 20)))&2) == 0 && int32(*(*uint8)(unsafe.Add(v13, 22))) <= 11 {
			return int8(wl.Dir0)
		}
	}
	return -1
}
func (s *Server) nox_xxx_pathfind_preCheckWalls_50C8D0(obj *server.Object, gpos *ntype.Point32) {
	if s.sub_50B870(obj, int(gpos.X), int(gpos.Y)) {
		dx := float64(obj.PosVec.X) - (float64(gpos.X)*23.0 + 11.5)
		dy := float64(obj.PosVec.Y) - (float64(gpos.Y)*23.0 + 11.5)
		if math.Abs(dy) >= math.Abs(dx) {
			if dy <= 0.0 {
				gpos.Y--
			} else {
				gpos.Y++
			}
		} else {
			if dx <= 0.0 {
				gpos.X--
			} else {
				gpos.X++
			}
		}
	}
}

func sub_50C830(a1 *server.Object, x, y int32) int32 {
	if a1.ObjFlags&0x4000 != 0 || sub_534020(a1) != 0 {
		return 1
	}
	v4 := memmap.PtrFloat32(0x587000, 234188)
	sx := float32(x * 23)
	sy := float32(y * 23)
	for {
		var v6 types.Pointf
		v6.X = sx + *(*float32)(unsafe.Add(unsafe.Pointer(v4), -int(unsafe.Sizeof(float32(0))*1)))
		v6.Y = sy + *v4
		if legacy.Nox_xxx_tileNFromPoint_411160(v6) == 6 {
			break
		}
		v4 = (*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*2))
		if int32(uintptr(unsafe.Pointer(v4))) >= int32(uintptr(memmap.PtrOff(0x587000, 234220))) {
			return 1
		}
	}
	return 0
}

func (s *Server) nox_xxx_genPathToPoint_50B9A0(path []types.Pointf, u *server.Object, a4 *types.Pointf) int {
	u.UpdateDataMonster().Field543_0 = 0
	s.nox_xxx_pathFind_50BA00(false, u, &u.PosVec, a4, nil, 0)
	return s.AI.Paths.MaybeAppendWorkPath(path)
}

func (s *Server) Nox_xxx_generateRetreatPath_50CA00(path []types.Pointf, u *server.Object, a4 *types.Pointf) int {
	u.UpdateDataMonster().Field543_0 = 0
	s.nox_xxx_pathFind_50BA00(true, u, &u.PosVec, a4, s.AI.Paths.HasNoEnemiesAround, 6)
	return s.AI.Paths.MaybeAppendWorkPath(path)
}

func (s *Server) Nox_xxx_creatureSetDetailedPath_50D220(obj *server.Object, a2 *types.Pointf) {
	ud := obj.UpdateDataMonster()
	ud.Field67 = 0
	if (s.Frame() - ud.Field70) >= 10 {
		ud.Field68 = *a2
		ud.Field2 = uint32(s.nox_xxx_genPathToPoint_50B9A0(ud.Path[:], obj, a2))
		*(*uint8)(unsafe.Add(unsafe.Pointer(ud), 284)) = uint8(int8(s.AI.Paths.PathFindStatus()))
		ud.Field70 = s.Frame()
	} else {
		ud.Field2 = 0
		*(*uint8)(unsafe.Add(unsafe.Pointer(ud), 284)) = 1
	}
}

func (s *Server) Sub_50B810(obj *server.Object, pos *types.Pointf) bool {
	x := int(float64(pos.X) / 23)
	y := int(float64(pos.Y) / 23)
	if s.sub_50B870(obj, x, y) {
		return false
	}
	return s.AI.Paths.Nox_xxx_pathfind_preCheckWalls2_50B8A0(obj, x, y)
}
