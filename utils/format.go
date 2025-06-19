package utils

import (
	"bytes"
	"image"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func ImageFormat(data []byte) (string, error) {
	reader := bytes.NewReader(data)

	_, format, err := image.DecodeConfig(reader)
	if err != nil {
		return "", err
	}

	return format, nil
}
