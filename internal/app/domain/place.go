package domain

// Place is a domain place.
type Place struct {
	ID          int64
	X           float32
	Y           float32
	Width       float32
	Height      float32
	IsAvailable bool
}
