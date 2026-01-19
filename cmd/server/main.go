package main

import (
	"fmt"
	"log"

	"github.com/aditiAgrawaldev/postera/internal/channels"
	"github.com/aditiAgrawaldev/postera/internal/config"
	"github.com/aditiAgrawaldev/postera/internal/models"
	"github.com/aditiAgrawaldev/postera/internal/service"
	"github.com/aditiAgrawaldev/postera/pkg/template"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	emailCh := channels.NewEmailChannel(config.SMTP.Host, config.SMTP.Port, config.SMTP.From)
	slackCh := channels.NewSlackChannel(config.Slack.WebhookURL)

	notificationService := service.NewNotificationService(emailCh, slackCh)
	recipients, err := notificationService.LoadRecipient("data.csv")
	if err != nil {
		log.Fatalf("Failed to load recipients: %v", err)
	}

	fmt.Println("Loaded recipients")

	var notifications []models.Notification
	for _, recipient := range recipients {

		data := map[string]string{
			"Name":  recipient.Name,
			"Email": recipient.Email,
		}

		subject, err := template.LoadTemplate("templates/subject.txt", data)
		if err != nil {
			log.Fatalf("Failed to load subject template: %v", err)
		}

		body, err := template.LoadTemplate("templates/body.txt", data)
		if err != nil {
			log.Fatalf("Failed to load body template: %v", err)
		}

		notifications = append(notifications, models.Notification{
			Recipient: recipient,
			Subject:   subject,
			Body:      body,
		})
	}

	fmt.Println("Sending notifications")
	notificationService.SendNotifications(notifications, 5)
	fmt.Println("Notifications sent successfully")
}
