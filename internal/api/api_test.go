package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/0xN0x/go-artifactsmmo/internal/api"
	"github.com/0xN0x/go-artifactsmmo/internal/client"
	"github.com/0xN0x/go-artifactsmmo/models"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a test server and config
func setupTestServer(handler http.HandlerFunc) (*client.ArtifactsConfig, func()) {
	server := httptest.NewServer(handler)

	config := client.NewConfig(
		server.Client(),
		server.URL,
		"test-token",
		"test-user",
	)

	cleanup := func() {
		server.Close()
	}

	return config, cleanup
}

func TestNewRequest(t *testing.T) {
	config := client.NewConfig(
		&http.Client{},
		"https://api.example.com",
		"test-token",
		"test-user",
	)

	req := api.NewRequest(config)
	assert.NotNil(t, req, "NewRequest should return a non-nil request")
	assert.Equal(t, config, req.Config, "Config should be set correctly")
}

func TestSetMethod(t *testing.T) {
	config := client.NewConfig(
		&http.Client{},
		"https://api.example.com",
		"test-token",
		"test-user",
	)

	req := api.NewRequest(config)
	req.SetMethod("GET")

	httpReq := req.GetRequest()
	assert.Equal(t, "GET", httpReq.Method, "Method should be set to GET")
}

func TestSetURL(t *testing.T) {
	config := client.NewConfig(
		&http.Client{},
		"https://api.example.com",
		"test-token",
		"test-user",
	)

	req := api.NewRequest(config)
	req.SetURL("/test")

	httpReq := req.GetRequest()
	assert.Equal(t, "https://api.example.com/test", httpReq.URL.String(), "URL should be correctly formed")
}

func TestSetParam(t *testing.T) {
	config := client.NewConfig(
		&http.Client{},
		"https://api.example.com",
		"test-token",
		"test-user",
	)

	req := api.NewRequest(config)
	req.SetURL("/test").SetParam("key", "value")

	httpReq := req.GetRequest()
	assert.Equal(t, "key=value", httpReq.URL.RawQuery, "Query parameter should be correctly added")
}

func TestSetBody(t *testing.T) {
	config := client.NewConfig(
		&http.Client{},
		"https://api.example.com",
		"test-token",
		"test-user",
	)

	type TestBody struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	body := TestBody{
		Name: "Test User",
		Age:  30,
	}

	req := api.NewRequest(config)
	req.SetMethod("POST").SetURL("/test").SetBody(body)

	// Verify the body is set
	assert.Equal(t, body, req.Body, "Body should be set correctly")
}

func TestSetResultStruct(t *testing.T) {
	config := client.NewConfig(
		&http.Client{},
		"https://api.example.com",
		"test-token",
		"test-user",
	)

	type TestResponse struct {
		Success bool `json:"success"`
	}

	responseStruct := &TestResponse{}

	req := api.NewRequest(config)
	req.SetResultStruct(responseStruct)

	assert.Equal(t, responseStruct, req.GetCustomStruct(), "Result struct should be set correctly")
}

func TestRun_Success(t *testing.T) {
	// Mock server that returns a success response
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Check request headers
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"), "Authorization header should be set")
		assert.Equal(t, "application/json", r.Header.Get("Accept"), "Accept header should be set")

		// Check request path
		assert.Equal(t, "/character", r.URL.Path, "Path should be set correctly")

		// Return a mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		responseData := struct {
			Data struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"data"`
		}{
			Data: struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			}{
				ID:   "char-123",
				Name: "Test Character",
			},
		}

		json.NewEncoder(w).Encode(responseData)
	}

	config, cleanup := setupTestServer(handler)
	defer cleanup()

	// Define a struct to receive the response
	character := &struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{}

	// Create and execute the request
	req := api.NewRequest(config)
	req.SetMethod("GET").SetURL("/character").SetResultStruct(character)

	res, err := req.Run()

	// Assert success
	assert.NoError(t, err, "Request should succeed")
	assert.Equal(t, http.StatusOK, res.StatusCode, "Status code should be 200 OK")
	assert.Equal(t, "char-123", character.ID, "Character ID should be correct")
	assert.Equal(t, "Test Character", character.Name, "Character name should be correct")
}

func TestRun_ErrorHandling(t *testing.T) {
	// Test cases for different error codes
	testCases := []struct {
		statusCode  int
		expectedErr error
		name        string
	}{
		{452, models.ErrBadToken, "Bad Token Error"},
		{474, models.ErrTaskNotOwned, "Task Not Owned Error"},
		{475, models.ErrTaskAlreadyCompleted, "Task Already Completed Error"},
		{478, models.ErrInsufficientQuantity, "Insufficient Quantity Error"},
		{486, models.ErrActionInProgress, "Action In Progress Error"},
		{487, models.ErrCharacterHasNoTask, "Character Has No Task Error"},
		{488, models.ErrTaskNotCompleted, "Task Not Completed Error"},
		{489, models.ErrCharacterAlreadyHasTask, "Character Already Has Task Error"},
		{497, models.ErrCharacterFullInventory, "Character Full Inventory Error"},
		{498, models.ErrCharacterNotFound, "Character Not Found Error"},
		{499, models.ErrCharacterInCooldown, "Character In Cooldown Error"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock server that returns the specified error status code
			handler := func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				w.Write([]byte("{}")) // Empty response body
			}

			config, cleanup := setupTestServer(handler)
			defer cleanup()

			// Create and execute the request
			req := api.NewRequest(config)
			req.SetMethod("GET").SetURL("/test")

			_, err := req.Run()

			// Assert that the correct error is returned
			assert.Equal(t, tc.expectedErr, err, "Request should return the expected error")
		})
	}
}

func TestRun_HTTPError(t *testing.T) {
	// Create a server that will be immediately closed to simulate a connection error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	server.Close() // Close immediately to force connection error

	config := client.NewConfig(
		&http.Client{},
		server.URL,
		"test-token",
		"test-user",
	)

	// Create and execute the request
	req := api.NewRequest(config)
	req.SetMethod("GET").SetURL("/test")

	_, err := req.Run()

	// Assert that an error is returned
	assert.Error(t, err, "Request should return an error for connection issues")
}

func TestSet_Headers(t *testing.T) {
	// Mock server that checks headers
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Check request headers
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"), "Authorization header should be set")
		assert.Equal(t, "application/json", r.Header.Get("Accept"), "Accept header should be set")

		if r.Method == "POST" {
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"), "Content-Type header should be set for POST")
		} else {
			assert.Empty(t, r.Header.Get("Content-Type"), "Content-Type should not be set for non-POST")
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{}}`))
	}

	config, cleanup := setupTestServer(handler)
	defer cleanup()

	// Test GET request
	t.Run("GET Headers", func(t *testing.T) {
		req := api.NewRequest(config)
		req.SetMethod("GET").SetURL("/test")
		_, err := req.Run() // This will call Set() internally
		assert.NoError(t, err)
	})

	// Test POST request
	t.Run("POST Headers", func(t *testing.T) {
		req := api.NewRequest(config)
		req.SetMethod("POST").SetURL("/test").SetBody(map[string]string{"key": "value"})
		_, err := req.Run() // This will call Set() internally
		assert.NoError(t, err)
	})
}
