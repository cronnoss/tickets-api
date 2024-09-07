package domain

// Show is a domain show.

type Show struct {
	id   int
	name string
}

type NewShowData struct {
	ID   int
	Name string
}

// NewShow creates a new show.
func NewShow(data NewShowData) (Show, error) {
	return Show{
		id:   data.ID,
		name: data.Name,
	}, nil
}

// ID returns the show ID.
func (s Show) ID() int {
	return s.id
}

// Name returns the show name.
func (s Show) Name() string {
	return s.name
}
