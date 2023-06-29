package promptlayer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PromptTemplate struct {
	Template       string   `json:"template,omitempty"`
	InputVariables []string `json:"input_variables,omitempty"`
}

func (pt *PromptTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type           string   `json:"_type"`
		Template       string   `json:"template,omitempty"`
		InputVariables []string `json:"input_variables,omitempty"`
	}{
		Type:           "prompt",
		Template:       pt.Template,
		InputVariables: pt.InputVariables,
	})
}

type PublishPromptTemplateInput struct {
	PromptName     string         `json:"prompt_name,omitempty"`
	PromptTemplate PromptTemplate `json:"prompt_template,omitempty"`
	Tags           []string       `json:"tags,omitempty"`
	APIKey         string         `json:"api_key,omitempty"`
}

type PublishPromptTemplateOutput struct {
	ID      uint64 `json:"id"`
	Success bool   `json:"success"`
}

func (c *Client) PublishPromptTemplate(ctx context.Context, input *PublishPromptTemplateInput) (*PublishPromptTemplateOutput, error) {
	url := fmt.Sprintf("%s/rest/publish-prompt-template", c.baseURL)

	input.Tags = append(input.Tags, c.tags...)
	input.APIKey = c.apiKey

	body, err := c.doRequest(ctx, http.MethodPost, url, input)
	if err != nil {
		return nil, err
	}

	output := &PublishPromptTemplateOutput{}
	if err := json.Unmarshal(body, output); err != nil {
		return nil, err
	}

	return output, nil
}
