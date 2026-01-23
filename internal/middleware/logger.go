package middleware

import (
	"fmt"
	"time"
)

func LogInfo(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[INFO] [%s] %s\n", timestamp, message)
}

func LogSuccess(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[SUCCESS] [%s] %s\n", timestamp, message)
}

func LogError(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[ERROR] [%s] %s\n", timestamp, message)
}

func LogRetry(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[RETRY] [%s] %s\n", timestamp, message)
}

