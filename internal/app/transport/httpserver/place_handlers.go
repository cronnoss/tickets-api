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

// @Summary GetPlaces
// @Tags place
// @Description get places by event ID
// @ID get-places
// @Accept  json
// @Produce  json
// @Param id path int true "event ID"
// @Success 200 {object} PlaceListResponse
// @Failure 400,404 {object} server.ErrorResponse
// @Failure 500 {object} server.ErrorResponse
// @Router /events/{id}/places [get]
func (h *HTTPServer) GetPlaces(w http.ResponseWriter, r *http.Request) {
	// Step 1: Make a GET request to the remote API
	vars := mux.Vars(r)
	id := vars["id"]
	remoteURL := "https://leadbook.ru/test-task-api/events/" + id + "/places"
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, remoteURL, nil)
	if err != nil {
		server.RespondWithError(fmt.Errorf("failed to create request: %w", err), w, r)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		server.RespondWithError(fmt.Errorf("failed to fetch places from remote: %w", err), w, r)
		return
	}
	defer resp.Body.Close()

	// Step 2: Read and decode the remote API response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		server.RespondWithError(fmt.Errorf("failed to read response body: %w", err), w, r)
		return
	}

	var placeListResponse PlaceListResponse
	if err := json.Unmarshal(body, &placeListResponse); err != nil {
		server.RespondWithError(fmt.Errorf("failed to decode response: %w", err), w, r)
		return
	}

	// Step 3: Iterate over places and store them in the local service
	for _, place := range placeListResponse.Response {
		_, err := h.placeService.CreatePlace(r.Context(), domain.Place{
			ID:          int64(place.ID),
			X:           float32(place.X),
			Y:           float32(place.Y),
			Width:       float32(place.Width),
			Height:      float32(place.Height),
			IsAvailable: place.IsAvailable,
		})
		if err != nil {
			server.RespondWithError(fmt.Errorf("failed to create place: %w", err), w, r)
			return
		}
	}

	server.RespondOK(placeListResponse, w, r)
}
