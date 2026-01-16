package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
)

func consumer(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		formattedMessage := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email, "Testing Email campaign")
		msg := []byte(formattedMessage)

		fmt.Printf("Consumer %d sending email to %s\n", id, recipient.Email)

		err := smtp.SendMail(smtpHost+":"+smtpPort, nil, "aditi@testing.com", []string{recipient.Email}, msg)
		if err != nil {
			log.Fatalf("Error sending email to %s: %v\n", recipient.Email, err)
		}
		
		fmt.Printf("Consumer %d sent email to %s\n", id, recipient.Email)
	}

	//Waitgroup chalta kaise hai ?
	//Waitgroup is a counter, agar consumer kaam kar raha hai to counter increase kardo, or jaise hi kaam khatam ho jaye to decrease kardo
}
