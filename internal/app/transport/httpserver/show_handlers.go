package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cronnoss/tickets-api/internal/app/common/server"
	"github.com/cronnoss/tickets-api/internal/app/repository/models"
)

func (h *HTTPServer) GetShows(w http.ResponseWriter, r *http.Request) {
	// Step 1: Make a GET request to the remote API
	remoteURL := "https://leadbook.ru/test-task-api/shows"
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, remoteURL, nil)
	if err != nil {
		server.RespondWithError(fmt.Errorf("failed to create request: %w", err), w, r)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		server.RespondWithError(fmt.Errorf("failed to fetch shows from remote: %w", err), w, r)
		return
	}
	defer resp.Body.Close()

	// Step 2: Read and decode the remote API response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		server.RespondWithError(fmt.Errorf("failed to read response body: %w", err), w, r)
		return
	}

	var showListResponse ShowListResponse
	if err := json.Unmarshal(body, &showListResponse); err != nil {
		server.RespondWithError(fmt.Errorf("failed to decode response: %w", err), w, r)
		return
	}

	// Step 3: Iterate over shows and store them in the local service
	for _, show := range showListResponse.Response {
		_, err := h.showService.CreateShow(r.Context(), models.Show{
			ID:   show.ID,
			Name: show.Name,
		})
		if err != nil {
			server.RespondWithError(fmt.Errorf("failed to create show: %w", err), w, r)
			return
		}
	}

	server.RespondOK(showListResponse, w, r)
}
