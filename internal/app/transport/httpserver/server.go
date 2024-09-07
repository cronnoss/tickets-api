package httpserver

// HTTPServer is a HTTP server for ports.
type HTTPServer struct {
	showService  ShowService
	eventService EventService
	placeService PlaceService
}

// NewHTTPServer creates a new HTTP server for ports.
func NewHTTPServer(showService ShowService, eventService EventService, placeService PlaceService) HTTPServer {
	return HTTPServer{
		showService:  showService,
		eventService: eventService,
		placeService: placeService,
	}
}
