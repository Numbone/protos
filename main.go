package main

import (
	"fmt"
	"time"
)

func main() {
	send := make(chan string)
	receive := make(chan string)

	// Горутину для отправки "hello"
	go func() {
		for {
			send <- "hello"
			time.Sleep(500 * time.Millisecond) // не спамим процессор
		}
	}()

	// Горутину для отправки "world"
	go func() {
		for {
			receive <- "world"
			time.Sleep(1500 * time.Millisecond) // другое время, чтобы чередовалось
		}
	}()

	// Главный цикл обработки сообщений
	for {
		select {
		case msg := <-send:
			fmt.Println("Send:", msg)
		case msg := <-receive:
			fmt.Println("Receive:", msg)
		}
	}
}
