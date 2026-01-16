package main

import "time"

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)
	go loadRecipient("email.csv", recipientChannel)
	go consumer(1, recipientChannel)

	time.Sleep(2 * time.Second)
}
