package main

import "time"

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)
	go loadRecipient("email.csv", recipientChannel)

	consumerCount := 5

	for i := 1; i <= consumerCount; i++ {
		go consumer(i, recipientChannel)
	}

	time.Sleep(2 * time.Second)
}
