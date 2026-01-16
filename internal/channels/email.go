package channels

import (
	"fmt"
	"net/smtp"

	"github.com/aditiAgrawaldev/postera/internal/models"
)

type EmailChannel struct {
	SMTPHost string
	SMTPPort string
	From     string
}

//ismei only vo fields rakhenge jo change nahi hongi, otherwise it will be a bad code

func NewEmailChannel(smtpHost, smtpPort, from string) *EmailChannel {
	return &EmailChannel{
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
		From:     from,
	}
}

func (e *EmailChannel) Send(notification models.Notification) error {

	formattedMessage := fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\n\r\n%s\r\n",
		notification.Recipient.Email,
		notification.Subject,
		notification.Body,
	)

	address := e.SMTPHost + ":" + e.SMTPPort
	err := smtp.SendMail(address, nil, e.From, []string{notification.Recipient.Email}, []byte(formattedMessage))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil

	//Abhi we are refactoring the code to use the EmailChannel struct, so that we can use the EmailChannel struct to send emails
	//Some code of consumer.go is refactored
}
