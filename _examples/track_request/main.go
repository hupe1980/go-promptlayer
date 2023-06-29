package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hupe1980/go-promptlayer"
)

func main() {
	client := promptlayer.NewClient(os.Getenv("PROMPTLAYER_API_KEY"))

	startTime := time.Now()
	endTime := startTime.Add(3 * time.Second)

	output, err := client.TrackRequest(context.Background(), &promptlayer.TrackRequestInput{
		FunctionName: "openai.Completion.create",
		// kwargs will need messages if using chat-based completion
		Kwargs: map[string]any{
			"engine": "text-ada-001",
			"prompt": "My name is",
		},
		Tags: []string{"hello", "world"},
		RequestResponse: map[string]any{
			"id":      "cmpl-6TEeJCRVlqQSQqhD8CYKd1HdCcFxM",
			"object":  "text_completion",
			"created": 1672425843,
			"model":   "text-ada-001",
			"choices": []map[string]any{
				{
					"text":          " advocacy\"\n\nMy name is advocacy.",
					"index":         0,
					"logprobs":      nil,
					"finish_reason": "stop",
				},
			},
		},
		RequestStartTime: startTime,
		RequestEndTime:   endTime,
		Metadata: map[string]string{
			"Hello": "World",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID:", output.RequestID)
}
