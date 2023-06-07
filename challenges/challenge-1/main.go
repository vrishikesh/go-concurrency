package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	updateMessage("Hello, cosmos!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	updateMessage("Hello, world!", &wg)
	wg.Wait()
	printMessage()
}
