package channels

import "github.com/aditiAgrawaldev/postera/internal/models"

type Channel interface {
	Send(notification models.Notification) error
}