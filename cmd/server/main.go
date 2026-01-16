package main

import (
	"fmt"
	"log"

	"github.com/aditiAgrawaldev/postera/internal/channels"
	"github.com/aditiAgrawaldev/postera/internal/service"
)

func main() {
	emailCh := channels.NewEmailChannel("localhost", "1025", "aditi@testing.com")
	notificationService := service.NewNotificationService(emailCh)
	recipients, err := notificationService.loadRecipient("email.csv")
	if err != nil {
		log.Fatalf("Failed to load recipients: %v", err)
	}

	fmt.Println("Loaded recipients:", recipients)

	var notifications []models.Notification
	for _, recipient := range recipients {
		notifications = append(notifications, models.Notification{
			Recipient: recipient,
			Subject: "Test Notification",
			Body: "This is a test notification",
		})
	}

	fmt.Println("Sending notifications:", notifications)
	notificationService.SendNotifications(notifications, 5)
	fmt.Println("Notifications sent successfully")
}
