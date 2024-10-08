package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cronnoss/tickets-api/internal/app/common/server"
	"github.com/cronnoss/tickets-api/internal/app/domain"
	"github.com/gorilla/mux"
)

// @Summary GetEvents
// @Tags event
// @Description get events by show ID
// @ID get-events
// @Accept  json
// @Produce  json
// @Param id path int true "show ID"
// @Success 200 {object} EventListResponse
// @Failure 400,404 {object} server.ErrorResponse
// @Failure 500 {object} server.ErrorResponse
// @Router /shows/{id}/events [get]
func (h *HTTPServer) GetEvents(w http.ResponseWriter, r *http.Request) {
	// Step 1: Make a GET request to the remote API
	vars := mux.Vars(r)
	id := vars["id"]
	remoteURL := "https://leadbook.ru/test-task-api/shows/" + id + "/events"
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, remoteURL, nil)
	if err != nil {
		server.RespondWithError(fmt.Errorf("failed to create request: %w", err), w, r)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		server.RespondWithError(fmt.Errorf("failed to fetch events from remote: %w", err), w, r)
		return
	}
	defer resp.Body.Close()

	// Step 2: Read and decode the remote API response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		server.RespondWithError(fmt.Errorf("failed to read response body: %w", err), w, r)
		return
	}

	var eventListResponse EventListResponse
	if err := json.Unmarshal(body, &eventListResponse); err != nil {
		server.RespondWithError(fmt.Errorf("failed to decode response: %w", err), w, r)
		return
	}

	// Step 3: Iterate over events and store them in the local service
	for _, event := range eventListResponse.Response {
		_, err := h.eventService.CreateEvent(r.Context(), domain.Event{
			ID:     int64(event.ID),
			ShowID: int64(event.ShowID),
			Date:   event.Date,
		})
		if err != nil {
			server.RespondWithError(fmt.Errorf("failed to create event: %w", err), w, r)
			return
		}
	}

	server.RespondOK(eventListResponse, w, r)
}
