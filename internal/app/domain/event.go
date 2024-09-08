package domain

// Event is a domain event.
type Event struct {
	id     int64
	showID int64
	date   string
}

type NewEventData struct {
	ID     int64
	ShowID int64
	Date   string
}

// NewEvent creates a new event.
func NewEvent(data NewEventData) (Event, error) {
	return Event{
		id:     data.ID,
		showID: data.ShowID,
		date:   data.Date,
	}, nil
}

// ID returns the event ID.
func (e Event) ID() int64 {
	return e.id
}

// ShowID returns the event show ID.
func (e Event) ShowID() int64 {
	return e.showID
}

// Date returns the event date.
func (e Event) Date() string {
	return e.date
}
