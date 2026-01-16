package main

import (
	"fmt"
	"sync"
)

func consumer(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()

	for recipient := range ch {
		fmt.Println(id, recipient)
	}

	//Waitgroup chalta kaise hai ?
	//Waitgroup is a counter, agar consumer kaam kar raha hai to counter increase kardo, or jaise hi kaam khatam ho jaye to decrease kardo
}
