package script

type Breaker interface {
	Break()
}

type Wall interface {
	Positioner
	GridPositioner
	Enabler // opens and closes the wall
	Breaker
}

var (
	_ Identifiable = &WallGroup{}
	_ EnableSetter = &WallGroup{}
	_ Toggler      = &WallGroup{}
	_ Breaker      = &WallGroup{}
)

func NewWallGroup(id string, list ...Wall) *WallGroup {
	return &WallGroup{id: id, list: list}
}

type WallGroup struct {
	id   string
	list []Wall
}

func (g *WallGroup) ID() string {
	if g == nil {
		return ""
	}
	return g.id
}

func (g *WallGroup) Walls() []Wall {
	if g == nil {
		return nil
	}
	return g.list
}

func (g *WallGroup) Enable(enable bool) {
	if g == nil {
		return
	}
	for _, v := range g.list {
		v.Enable(enable)
	}
}

func (g *WallGroup) Toggle() bool {
	if g == nil {
		return false
	}
	var st bool
	for i, v := range g.list {
		ns := Toggle(v)
		if i == 0 {
			st = ns
		}
	}
	return st
}

func (g *WallGroup) Break() {
	if g == nil {
		return
	}
	for _, v := range g.list {
		v.Break()
	}
}
