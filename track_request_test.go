package promptlayer

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_TrackRequest(t *testing.T) {
	client := NewClient("your-api-key")

	t.Run("Successful request", func(t *testing.T) {
		// Prepare the input data
		input := &TrackRequestInput{
			FunctionName: "your-function",
			Args:         []string{"arg1", "arg2"},
			// Set other input fields as needed
		}

		// Mock the HTTP request and response
		mockHTTPClient := &mockHTTPClient{
			doFunc: func(req *http.Request) (*http.Response, error) {
				// Assert the request details
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "/rest/track-request", req.URL.Path)

				// Process the input and return a mock response
				trackRequestOutput := map[string]any{
					"request_id": 123,
					"success":    true,
				}

				responseBody, _ := json.Marshal(trackRequestOutput)

				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader(responseBody)),
				}, nil
			},
		}

		client.httpClient = mockHTTPClient

		// Make the request
		ctx := context.Background()
		output, err := client.TrackRequest(ctx, input)

		// Assert the output and error
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, "123", output.RequestID)
		assert.True(t, output.Success)
	})

	t.Run("Failed request", func(t *testing.T) {
		// Prepare the input data
		input := &TrackRequestInput{
			FunctionName: "your-function",
			Args:         []string{"arg1", "arg2"},
			// Set other input fields as needed
		}

		// Mock the HTTP request and response
		mockHTTPClient := &mockHTTPClient{
			doFunc: func(req *http.Request) (*http.Response, error) {
				// Return a mock error response
				responseBody := []byte(`{"message": "Request failed"}`)
				return &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       io.NopCloser(bytes.NewReader(responseBody)),
				}, nil
			},
		}

		client.httpClient = mockHTTPClient

		// Make the request
		ctx := context.Background()
		output, err := client.TrackRequest(ctx, input)

		// Assert the error and output
		assert.Error(t, err)
		assert.Nil(t, output)
	})
}
