package promptlayer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TrackScoreInput struct {
	RequestID uint64 `json:"request_id,omitempty"`
	Score     uint   `json:"score,omitempty"`
	APIKey    string `json:"api_key,omitempty"`
}

type TrackScoreOutput struct {
	Success bool `json:"success"`
}

func (c *Client) TrackScore(ctx context.Context, input *TrackScoreInput) (*TrackScoreOutput, error) {
	if input.Score > 100 {
		return nil, fmt.Errorf("score must be between 0 and 100, got %d", input.Score)
	}

	url := fmt.Sprintf("%s/rest/track-score", c.baseURL)

	input.APIKey = c.apiKey

	body, err := c.doRequest(ctx, http.MethodPost, url, input)
	if err != nil {
		return nil, err
	}

	output := &TrackScoreOutput{}
	if err := json.Unmarshal(body, output); err != nil {
		return nil, err
	}

	return output, nil
}
