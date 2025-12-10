package grid

import (
	L "github.com/wthys/advent-of-code-2025/location"
	I "github.com/wthys/advent-of-code-2025/util/interval"
)

type (
	Bounds interface {
		Has(L.Location) bool
		Width() int
		Height() int
		IsEmpty() bool

		TopLeft() L.Location
		TopRight() L.Location
		BottomLeft() L.Location
		BottomRight() L.Location

		Intersects(Bounds) bool
		Accomodate(L.Location) Bounds
		ForEach(func(L.Location))
	}

	bounds struct {
		X, Y I.Interval
	}
)

func (b bounds) Has(loc L.Location) bool {
	return b.X.Contains(loc.X) && b.Y.Contains(loc.Y)
}

func (b bounds) Width() int {
	return b.X.Len()
}

func (b bounds) Height() int {
	return b.Y.Len()
}

func (b bounds) IsEmpty() bool {
	return b.X.IsEmpty() || b.Y.IsEmpty()
}

func (b bounds) Accomodate(loc L.Location) Bounds {
	if b.IsEmpty() {
		return BoundsFromLocation(loc)
	}

	return bounds{
		I.New(min(b.X.Lower(), loc.X), max(b.X.Upper(), loc.X)),
		I.New(min(b.Y.Lower(), loc.Y), max(b.Y.Upper(), loc.Y)),
	}
}

func (b bounds) TopLeft() L.Location {
	return L.New(b.X.Lower(), b.Y.Lower())
}

func (b bounds) TopRight() L.Location {
	return L.New(b.X.Upper(), b.Y.Lower())
}

func (b bounds) BottomLeft() L.Location {
	return L.New(b.X.Lower(), b.Y.Upper())
}

func (b bounds) BottomRight() L.Location {
	return L.New(b.X.Upper(), b.Y.Upper())
}

func (b bounds) Intersects(o Bounds) bool {
	ob, ok := o.(bounds)
	if !ok {
		tl := o.TopLeft()
		br := o.BottomRight()
		ob = bounds{I.New(tl.X, br.X), I.New(tl.Y, br.Y)}
	}
	return b.X.Intersects(ob.X) && b.Y.Intersects(ob.Y)
} 

func BoundsFromLocation(loc L.Location) Bounds {
	return bounds{I.New(loc.X, loc.X), I.New(loc.Y, loc.Y)}
}

func BoundsFromSlice(locations L.Locations) Bounds {
	if len(locations) == 0 {
		return BoundsEmpty()
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

func BoundsEmpty() Bounds {
	return bounds{I.Empty(), I.Empty()}
}

func boundsFromLimits(xmin, xmax, ymin, ymax int) Bounds {
	return bounds{I.New(xmin, xmax), I.New(ymin, ymax)}
}

func (b bounds) ForEach(forEach func(loc L.Location)) {
	for y := b.Y.Lower(); y <= b.Y.Upper(); y++ {
		for x := b.X.Lower(); x <= b.X.Upper(); x++ {
			loc := L.New(x, y)
			forEach(loc)
		}
	}
}
