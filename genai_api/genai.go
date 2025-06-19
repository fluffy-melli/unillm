package genai_api

import (
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Option struct {
	Safety        []*genai.SafetySetting
	Temperature   float32
	TopP          float32
	TopK          int32
	MaxTokens     int32
	StopSequences []string
	Tools         []*genai.Tool
	ToolConfig    *genai.ToolConfig
	MimeType      string
}

type MD struct {
	ctx   context.Context
	model *genai.GenerativeModel
}

func NewModel(name, apikey string, opts *Option) (*MD, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apikey))

	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel(name)

	if opts == nil {
		return &MD{
			model: model,
			ctx:   ctx,
		}, nil
	}

	config := genai.GenerationConfig{}

	if opts.Temperature > 0 {
		config.Temperature = &opts.Temperature
	}

	if opts.TopP > 0 {
		config.TopP = &opts.TopP
	}

	if opts.TopK > 0 {
		config.TopK = &opts.TopK
	}

	if opts.MaxTokens > 0 {
		config.MaxOutputTokens = &opts.MaxTokens
	}

	if opts.StopSequences != nil {
		config.StopSequences = opts.StopSequences
	}

	if opts.MimeType != "" {
		config.ResponseMIMEType = opts.MimeType
	}

	model.GenerationConfig = config

	if opts.Safety != nil {
		model.SafetySettings = opts.Safety
	}

	if opts.Tools != nil {
		model.Tools = opts.Tools
	}

	if opts.ToolConfig != nil {
		model.ToolConfig = opts.ToolConfig
	}

	return &MD{
		model: model,
		ctx:   ctx,
	}, nil
}
