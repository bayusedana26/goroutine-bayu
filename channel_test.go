package goroutinebayu

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done sent data to Channel")

	}()
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
	close(channel) // Also can use defer close
}
