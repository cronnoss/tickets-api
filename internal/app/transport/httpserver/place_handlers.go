package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cronnoss/tickets-api/internal/app/common/server"
	"github.com/cronnoss/tickets-api/internal/app/repository/models"
	"github.com/gorilla/mux"
)

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
		_, err := h.placeService.CreatePlace(r.Context(), models.Place{
			ID:          place.ID,
			X:           place.X,
			Y:           place.Y,
			Width:       place.Width,
			Height:      place.Height,
			IsAvailable: place.IsAvailable,
		})
		if err != nil {
			server.RespondWithError(fmt.Errorf("failed to create place: %w", err), w, r)
			return
		}
	}

	server.RespondOK(placeListResponse, w, r)
}
