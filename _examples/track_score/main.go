package main

import (
	"context"
	"log"
	"os"

	"github.com/hupe1980/go-promptlayer"
)

func main() {
	client := promptlayer.NewClient(os.Getenv("PROMPTLAYER_API_KEY"))

	_, err := client.TrackScore(context.Background(), &promptlayer.TrackScoreInput{
		RequestID: 6370286,
		Score:     10,
	})
	if err != nil {
		log.Fatal(err)
	}
}
