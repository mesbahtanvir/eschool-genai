package services

import (
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIService struct {
	openAIClient *openai.Client
}

func NewOpenAIService() OpenAIService {
	client := openai.NewClient(option.WithAPIKey(os.Getenv("OPENAI_API_KEY")))
	return OpenAIService{
		openAIClient: client,
	}
}

func (oas OpenAIService) GenerateCourseBlueprint(courseHint string, userKnowledge string) (*models.CourseBlueprint, error) {

	// Generate prompt
	prompt := fmt.Sprintf(
		"Generate a course for: %s, where the user have the knowledge %s. "+
			"The response should be in JSON format like this: %s, just raw json text",
		courseHint,
		userKnowledge,
		GetStructMetadata(models.CourseBlueprint{}),
	)

	// Log the prompt that will be sent to OpenAI
	log.Printf("OpenAI Request Prompt: %s", prompt)
	chatCompletion, err := oas.openAIClient.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		panic(err.Error())
	}
	content := chatCompletion.Choices[0].Message.Content

	// Assuming the response text is JSON formatted with title, description, and modules
	var blueprint models.CourseBlueprint
	cleanedResponse := strings.TrimPrefix(content, "```json")
	cleanedResponse = strings.TrimSuffix(cleanedResponse, "```")
	cleanedResponse = strings.TrimSpace(cleanedResponse)
	log.Printf("OpenAI Response: %s", cleanedResponse)
	if err := json.Unmarshal([]byte(cleanedResponse), &blueprint); err != nil {
		log.Printf("JSON Unmarshal Error: %v", err)
		return nil, err
	}

	// Log the parsed CourseBlueprint struct
	log.Printf("Parsed Course Blueprint: %+v", blueprint)

	return &blueprint, nil
}
