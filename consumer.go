package main

import "fmt"

func consumer(id int, ch chan Recipient) {
	for recipient := range ch {
		fmt.Println(id, recipient)
	}
}
