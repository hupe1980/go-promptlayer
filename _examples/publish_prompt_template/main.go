package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hupe1980/go-promptlayer"
)

func main() {
	client := promptlayer.NewClient(os.Getenv("PROMPTLAYER_API_KEY"))

	output, err := client.PublishPromptTemplate(context.Background(), &promptlayer.PublishPromptTemplateInput{
		PromptName: "joke",
		PromptTemplate: promptlayer.PromptTemplate{
			Template:       "Tell me a {adjective} joke about {content}.",
			InputVariables: []string{"adjective", "content"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID:", output.ID)
}
