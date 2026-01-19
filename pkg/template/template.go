package template

import (
	"fmt"
	"os"
	"strings"
)

func LoadTemplate(filePath string, data map[string]string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %w", err)
	}

	text := string(content)

	for key, value := range data {
		placeholder := "{{" + key + "}}"
		text = strings.ReplaceAll(text, placeholder, value)
	}

	return text, nil

}
