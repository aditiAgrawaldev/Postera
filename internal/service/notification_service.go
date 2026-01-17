package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	"github.com/aditiAgrawaldev/postera/internal/channels"
	"github.com/aditiAgrawaldev/postera/internal/models"
)

type NotificationService struct {
	channels []channels.Channel
}

//We will send to multiple channels

func NewNotificationService(channelList ...channels.Channel) *NotificationService {
	return &NotificationService{
		channels: channelList,
	}
}

// Refactor karenge, same producer.go code
func (n *NotificationService) LoadRecipient(filePath string) ([]models.Recipient, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var recipients []models.Recipient
	for _, record := range records[1:] {
		recipients = append(recipients, models.Recipient{
			Name:  record[0],
			Email: record[1],
		})
	}

	return recipients, nil
}

func (n *NotificationService) SendNotifications(notifications []models.Notification, consumerCount int) error {
	notificationch := make(chan models.Notification)
	var wg sync.WaitGroup

	for i := 1; i <= consumerCount; i++ {
		wg.Add(1)
		go n.consumer(i, notificationch, &wg)
	}

	for _, notification := range notifications {
		notificationch <- notification
	}

	close(notificationch)
	wg.Wait()

	return nil
}

func (n *NotificationService) consumer(id int, notificationch chan models.Notification, wg *sync.WaitGroup) {
	defer wg.Done()
	for notification := range notificationch {
		fmt.Printf("Worker %d sending notification to %s\n", id, notification.Recipient.Email)

		for _, channel := range n.channels {

			err := channel.Send(notification)
			if err != nil {
				fmt.Printf("Worker %d failed to send notification to %s: %v\n", id, notification.Recipient.Email, err)
			} else {
				fmt.Printf("Worker %d sent notification to %s\n", id, notification.Recipient.Email)
			}
		}
	}
}
