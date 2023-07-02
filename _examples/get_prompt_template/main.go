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

	output, err := client.GetPromptTemplate(context.Background(), &promptlayer.GetPromptTemplateInput{
		PromptName: "joke",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID:", output.ID)
	fmt.Println("Template:", output.PromptTemplate)
}
