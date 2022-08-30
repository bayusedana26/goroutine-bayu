package goroutinebayu

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	pool.Put("Bayu")
	pool.Put("Sedana")
	pool.Put("Gale")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(2 * time.Second)
			pool.Put(data)

		}()

	}

	time.Sleep(10 * time.Second)
	fmt.Println("Selesai")
}
