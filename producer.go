package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadRecipient(filePath string, ch chan Recipient) error {
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

		//Now Next, this record will be consumed by the consumers, so this record need to be sent to the consumers via a path --> channel
		// producer -> channel -> consumers

		ch <- Recipient{
			Name: record[0],
			Email: record[1],
		}

		//Ab jo data hai, vo tabhi consume hoga jab consumer bhi ready hoga, because hum use kar rahe hai unbuffered channel
		// Iski capacity 0 hoti hai, isliye jab producer data send karega, toh consumer ready nahi hai, so data ko store nahi karega

	}

	return nil
}
