package promptlayer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TrackRequestInput struct {
	FunctionName         string            `json:"function_name,omitempty"`
	Args                 []string          `json:"args,omitempty"`
	Kwargs               map[string]any    `json:"kwargs,omitempty"`
	Tags                 []string          `json:"tags,omitempty"`
	RequestResponse      map[string]any    `json:"request_response,omitempty"`
	RequestStartTime     time.Time         `json:"request_start_time,omitempty"`
	RequestEndTime       time.Time         `json:"request_end_time,omitempty"`
	PromptID             string            `json:"prompt_id,omitempty"`
	PromptInputVariables map[string]string `json:"prompt_input_variables,omitempty"`
	PromptVersion        int               `json:"prompt_version,omitempty"`
	Metadata             map[string]string `json:"metadata,omitempty"`
	APIKey               string            `json:"api_key,omitempty"`
}

func (i *TrackRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		FunctionName         string            `json:"function_name,omitempty"`
		Args                 []string          `json:"args,omitempty"`
		Kwargs               map[string]any    `json:"kwargs,omitempty"`
		Tags                 []string          `json:"tags,omitempty"`
		RequestResponse      map[string]any    `json:"request_response,omitempty"`
		RequestStartTime     int64             `json:"request_start_time,omitempty"`
		RequestEndTime       int64             `json:"request_end_time,omitempty"`
		PromptID             string            `json:"prompt_id,omitempty"`
		PromptInputVariables map[string]string `json:"prompt_input_variables,omitempty"`
		PromptVersion        int               `json:"prompt_version,omitempty"`
		Metadata             map[string]string `json:"metadata,omitempty"`
		APIKey               string            `json:"api_key,omitempty"`
	}{
		FunctionName:         i.FunctionName,
		Args:                 i.Args,
		Kwargs:               i.Kwargs,
		Tags:                 i.Tags,
		RequestResponse:      i.RequestResponse,
		RequestStartTime:     i.RequestStartTime.Unix(),
		RequestEndTime:       i.RequestEndTime.Unix(),
		PromptID:             i.PromptID,
		PromptInputVariables: i.PromptInputVariables,
		Metadata:             i.Metadata,
		APIKey:               i.APIKey,
	})
}

type TrackRequestOutput struct {
	RequestID uint64 `json:"request_id"`
	Success   bool   `json:"success"`
}

func (c *Client) TrackRequest(ctx context.Context, input *TrackRequestInput) (*TrackRequestOutput, error) {
	url := fmt.Sprintf("%s/rest/track-request", c.baseURL)

	input.Tags = append(input.Tags, c.tags...)
	input.APIKey = c.apiKey

	body, err := c.doRequest(ctx, http.MethodPost, url, input)
	if err != nil {
		return nil, err
	}

	output := &TrackRequestOutput{}
	if err := json.Unmarshal(body, output); err != nil {
		return nil, err
	}

	return output, nil
}
