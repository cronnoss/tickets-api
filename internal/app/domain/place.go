package domain

// Place is a domain place.
type Place struct {
	id          int64
	x           float32
	y           float32
	width       float32
	height      float32
	isAvailable bool
}

type NewPlaceData struct {
	ID          int64
	X           float32
	Y           float32
	Width       float32
	Height      float32
	IsAvailable bool
}

// NewPlace creates a new place.
func NewPlace(data NewPlaceData) (Place, error) {
	return Place{
		id:          data.ID,
		x:           data.X,
		y:           data.Y,
		width:       data.Width,
		height:      data.Height,
		isAvailable: data.IsAvailable,
	}, nil
}

// ID returns the place ID.
func (p Place) ID() int64 {
	return p.id
}

// X returns the place x.
func (p Place) X() float32 {
	return p.x
}

// Y returns the place y.
func (p Place) Y() float32 {
	return p.y
}

// Width returns the place width.
func (p Place) Width() float32 {
	return p.width
}

// Height returns the place height.
func (p Place) Height() float32 {
	return p.height
}

// IsAvailable returns the place isAvailable.
func (p Place) IsAvailable() bool {
	return p.isAvailable
}
