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

func TestClient_PublishPromptTemplate(t *testing.T) {
	client := NewClient("your-api-key")

	t.Run("Successful Publish", func(t *testing.T) {
		// Mock the input data
		input := &PublishPromptTemplateInput{
			PromptName: "Test Prompt",
			PromptTemplate: PromptTemplate{
				Template:       "This is a test prompt",
				InputVariables: []string{"variable1", "variable2"},
			},
			Tags:   []string{"tag1", "tag2"},
			APIKey: "test-api-key",
		}

		// Mock the HTTP request and response
		mockHTTPClient := &mockHTTPClient{
			doFunc: func(req *http.Request) (*http.Response, error) {
				// Assert the request details
				assert.Equal(t, http.MethodPost, req.Method)
				assert.Equal(t, "/rest/publish-prompt-template", req.URL.Path)

				publishPromptTemplateOutput := map[string]any{
					"id":      123,
					"success": true,
				}

				responseBody, _ := json.Marshal(publishPromptTemplateOutput)

				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader(responseBody)),
				}, nil
			},
		}

		client.httpClient = mockHTTPClient

		// Perform the publish prompt template
		output, err := client.PublishPromptTemplate(context.Background(), input)

		// Assert that no error occurred
		assert.NoError(t, err)

		// Assert the output values
		expectedOutput := &PublishPromptTemplateOutput{
			ID:      "123",
			Success: true,
		}
		assert.Equal(t, expectedOutput, output)
	})

	t.Run("Failed Publish", func(t *testing.T) {
		// Mock the input data
		input := &PublishPromptTemplateInput{
			PromptName: "Test Prompt",
			PromptTemplate: PromptTemplate{
				Template:       "This is a test prompt",
				InputVariables: []string{"variable1", "variable2"},
			},
			Tags:   []string{"tag1", "tag2"},
			APIKey: "invalid-api-key", // Invalid API key to trigger an error
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

		// Perform the publish prompt template
		output, err := client.PublishPromptTemplate(context.Background(), input)

		// Assert the error and output
		assert.Error(t, err)
		assert.Nil(t, output)
	})
}
