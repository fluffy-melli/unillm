package genai

import (
	"fmt"

	"github.com/fluffy-melli/picrocess"
	"github.com/fluffy-melli/unillm/utils"
	"github.com/google/generative-ai-go/genai"
)

func (c *MD) Text(system, user string) (string, error) {
	c.model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(system),
		},
	}

	var generate []genai.Part

	generate = append(generate, genai.Text(user))

	resp, err := c.model.GenerateContent(c.ctx, generate...)

	if err != nil {
		return "", err
	}

	var respond string

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				respond += fmt.Sprint(part)
			}
		}
	}

	return respond, nil
}

func (c *MD) TextWithImage(system, user string, images ...[]byte) (string, error) {
	c.model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(system),
		},
	}

	var generate []genai.Part

	for _, i := range images {
		format, err := utils.ImageFormat(i)

		if err != nil {
			return "", err
		}

		generate = append(generate, genai.ImageData(format, i))
	}

	generate = append(generate, genai.Text(user))

	resp, err := c.model.GenerateContent(c.ctx, generate...)

	if err != nil {
		return "", err
	}

	var respond string

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				respond += fmt.Sprint(part)
			}
		}
	}

	return respond, nil
}

func (c *MD) TextWithImageURL(system, user string, images ...string) (string, error) {
	c.model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(system),
		},
	}

	var generate []genai.Part

	for _, url := range images {
		i, err := picrocess.ImageURL(url)

		if err != nil {
			return "", err
		}

		buf, err := i.ToPNGByte()

		if err != nil {
			return "", err
		}

		format, err := utils.ImageFormat(buf)

		if err != nil {
			return "", err
		}

		generate = append(generate, genai.ImageData(format, buf))
	}

	generate = append(generate, genai.Text(user))

	resp, err := c.model.GenerateContent(c.ctx, generate...)

	if err != nil {
		return "", err
	}

	var respond string

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				respond += fmt.Sprint(part)
			}
		}
	}

	return respond, nil
}
