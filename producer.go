package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadRecipient(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	reader := csv.NewReader(file)
	records, err :=reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	for _, record := range records[1:] {
		fmt.Println(record)
	}

	return nil
}
