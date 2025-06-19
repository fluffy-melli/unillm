package test

import (
	"fmt"
	"os"
	"testing"

	genai_api "github.com/fluffy-melli/unillm/genai"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
)

func TestGenAI(t *testing.T) {
	err := godotenv.Load()

	if err != nil {
		t.Fatal(err)
	}

	md, err := genai_api.NewModel("gemini-2.0-flash-lite", os.Getenv("API_KEY"), &genai_api.Option{
		MaxTokens: 100,
		Safety: []*genai.SafetySetting{
			{
				Category:  genai.HarmCategorySexuallyExplicit,
				Threshold: genai.HarmBlockLowAndAbove,
			},
			{
				Category:  genai.HarmCategoryDangerousContent,
				Threshold: genai.HarmBlockLowAndAbove,
			},
			{
				Category:  genai.HarmCategoryHateSpeech,
				Threshold: genai.HarmBlockLowAndAbove,
			},
			{
				Category:  genai.HarmCategoryHarassment,
				Threshold: genai.HarmBlockMediumAndAbove,
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	output, err := md.Text("정직하게 답변하세요.", "안녕?")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(output)
}
