package main

import (
	"fmt"
	"log"

	"github.com/aditiAgrawaldev/postera/internal/channels"
	"github.com/aditiAgrawaldev/postera/internal/config"
	"github.com/aditiAgrawaldev/postera/internal/models"
	"github.com/aditiAgrawaldev/postera/internal/service"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	emailCh := channels.NewEmailChannel(config.SMTP.Host, config.SMTP.Port, config.SMTP.From)
	notificationService := service.NewNotificationService(emailCh)
	recipients, err := notificationService.LoadRecipient("email.csv")
	if err != nil {
		log.Fatalf("Failed to load recipients: %v", err)
	}

	fmt.Println("Loaded recipients:", recipients)

	var notifications []models.Notification
	for _, recipient := range recipients {
		notifications = append(notifications, models.Notification{
			Recipient: recipient,
			Subject:   "Test Notification",
			Body:      "This is a test notification",
		})
	}

	fmt.Println("Sending notifications:", notifications)
	notificationService.SendNotifications(notifications, 5)
	fmt.Println("Notifications sent successfully")
}
