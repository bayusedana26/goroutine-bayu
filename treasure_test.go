package goroutinebayu

import (
	"fmt"
	"testing"
	"time"
)

var channel = make(chan string)
var storage = make(chan string)

var items = [7]string{"Rock", "Treasure", "Coral", "Treasure", "Rock", "Treasure", "Coral"}

func TestTreasure(t *testing.T) {

	go getItem(items)
	go cleanItem()
	go savedItem()

	time.Sleep(5 * time.Second)
}

// Running 3 goroutines, use channel to interact

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
		storage <- "is safe now"
	}
}

func savedItem() {
	for save := range storage {
		fmt.Println("Success keep the treasure", save)
	}
}
