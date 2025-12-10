package grid

import (
	L "github.com/wthys/advent-of-code-2025/location"
	I "github.com/wthys/advent-of-code-2025/util/interval"
)

type (
	Bounds struct {
		Xmin, Xmax, Ymin, Ymax int
	}
)

func (b *Bounds) Has(loc L.Location) bool {
	return loc.X >= b.Xmin && loc.X <= b.Xmax && loc.Y >= b.Ymin && loc.Y <= b.Ymax
}

func (b Bounds) Width() int {
	return b.Xmax - b.Xmin + 1
}

func (b Bounds) Height() int {
	return b.Ymax - b.Ymin + 1
}

func (b Bounds) Accomodate(loc L.Location) Bounds {
	newb := b
	newb.Xmin = min(b.Xmin, loc.X)
	newb.Xmax = max(b.Xmax, loc.X)
	newb.Ymin = min(b.Ymin, loc.Y)
	newb.Ymax = max(b.Ymax, loc.Y)
	return newb
}

func (b Bounds) TopLeft() L.Location {
	return L.New(b.Xmin, b.Ymin)
}

func (b Bounds) TopRight() L.Location {
	return L.New(b.Xmax, b.Ymin)
}

func (b Bounds) BottomLeft() L.Location {
	return L.New(b.Xmin, b.Ymax)
}

func (b Bounds) BottomRight() L.Location {
	return L.New(b.Xmax, b.Ymax)
}

func (b Bounds) Intersects(o Bounds) bool {
	bx := I.New(b.Xmin, b.Xmax)
	by := I.New(b.Ymin, b.Ymax)
	ox := I.New(o.Xmin, o.Xmax)
	oy := I.New(o.Ymin, o.Ymax)
	return bx.Intersects(ox) && by.Intersects(oy)
} 

func BoundsFromLocation(loc L.Location) Bounds {
	b := Bounds{}
	b.Xmin = loc.X
	b.Xmax = loc.X
	b.Ymin = loc.Y
	b.Ymax = loc.Y
	return b
}

func BoundsFromSlice(locations L.Locations) Bounds {
	if len(locations) == 0 {
		return Bounds{}
	}
	b := BoundsFromLocation(locations[0])
	for _, loc := range locations {
		b = b.Accomodate(loc)
	}
	return b
}

func BoundsFromLocations(locs ...L.Location) Bounds {
	return BoundsFromSlice(locs)
}

func (b Bounds) ForEach(forEach func(loc L.Location)) {
	for y := b.Ymin; y <= b.Ymax; y++ {
		for x := b.Xmin; x <= b.Xmax; x++ {
			loc := L.New(x, y)
			forEach(loc)
		}
	}
}
