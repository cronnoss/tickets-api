package httpserver

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/cronnoss/tickets-api/internal/app/common/server"
	"github.com/cronnoss/tickets-api/internal/app/domain"
	"github.com/cronnoss/tickets-api/internal/app/transport/httpserver/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetShows(t *testing.T) {
	showServiceMock := mocks.NewShowService(t)
	showServiceMock.On("CreateShow", mock.Anything, mock.Anything).Return(domain.Show{}, nil)

	_, err := showServiceMock.CreateShow(context.Background(), domain.Show{})
	require.NoError(t, err)

	// Step 1: Mock the remote API
	mockResponse := `{
    "response": [
        {"id": 1, "name": "Show #1"},
        {"id": 2, "name": "Show #2"}
    ]
  }`
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, mockResponse)
	},
	))

	// Step 2: Make a GET request to the remote API
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, mockServer.URL, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()

	// Step 3: Read and decode the remote API response
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var showListResponse ShowListResponse
	err = json.Unmarshal(body, &showListResponse)
	require.NoError(t, err)

	assert.Equal(t, 2, len(showListResponse.Response))
	assert.Equal(t, 1, showListResponse.Response[0].ID)
	assert.Equal(t, "Show #1", showListResponse.Response[0].Name)
	assert.Equal(t, 2, showListResponse.Response[1].ID)
	assert.Equal(t, "Show #2", showListResponse.Response[1].Name)
}
