package main

import "sync"

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)
	go loadRecipient("email.csv", recipientChannel)

	consumerCount := 5
	var wg sync.WaitGroup

	for i := 1; i <= consumerCount; i++ {
		wg.Add(1)
		go consumer(i, recipientChannel, &wg)
	}

	wg.Wait() // check ki counter 0 hua ya nahi, agar 0 hua toh main thread ko notify karega ki kaam khatam ho gaya, and main thread then exit kar dega
	//time.Sleep(2 * time.Second)
	//I will set wait group to synchronize the main thread with the goroutines, so that we need not to set hardcoded time
	//Consumer ko set karenge wait group ke through, so that main thread wait for all the consumers to finish their work
}
