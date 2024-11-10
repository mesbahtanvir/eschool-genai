package services

import (
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"log"

	openai "github.com/sashabaranov/go-openai"
)

func GenerateCourseBlueprint(courseHint string) (*models.CourseBlueprint, error) {
	client := openai.NewClient("OPENAI_API_KEY")

	// Generate prompt
	prompt := fmt.Sprintf(
		"Generate a course for: %s, where the user does not have any background knowledge on this. "+
			"The response should be in JSON format like this: %s",
		courseHint,
		GetStructMetadata(models.CourseBlueprint{}),
	)

	// Log the prompt that will be sent to OpenAI
	log.Printf("OpenAI Request Prompt: %s", prompt)

	// Create a request for a completion
	response, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
		Model:     "text-davinci-003", // Specify the model, e.g., text-davinci-003
		Prompt:    prompt,
		MaxTokens: 1000,
	})
	if err != nil {
		log.Printf("OpenAI API Error: %v", err)
		return nil, err
	}

	// Log the raw response text from OpenAI
	log.Printf("OpenAI Response Text: %s", response.Choices[0].Text)

	// Assuming the response text is JSON formatted with title, description, and modules
	var blueprint models.CourseBlueprint
	if err := json.Unmarshal([]byte(response.Choices[0].Text), &blueprint); err != nil {
		log.Printf("JSON Unmarshal Error: %v", err)
		return nil, err
	}

	// Log the parsed CourseBlueprint struct
	log.Printf("Parsed Course Blueprint: %+v", blueprint)

	return &blueprint, nil
}
