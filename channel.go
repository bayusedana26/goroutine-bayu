package main

import (
	"fmt"
	"time"
)

var channel = make(chan string)
var storage = make(chan string)

func getItem(items [7]string) {
	for _, item := range items {
		if item == "Treasure" {
			fmt.Println("Sucess get", item)
			channel <- item
		}
	}
}

func cleanItem() {
	for rawItem := range channel {
		fmt.Println("Success clean", rawItem)
		storage <- "Clean Treasure is been moved"
	}
}

func savedItem() {
	for save := range storage {
		fmt.Println("Success keep the treasure", save)
	}
}

func main() {
	items := [7]string{"Rock", "Treasure", "Coral", "Treasure", "Rock", "Treasure", "Coral"}

	go getItem(items)
	go cleanItem()
	go savedItem()

	time.Sleep(5 * time.Second)
}
