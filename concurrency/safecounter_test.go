package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestSafeCounter(t *testing.T) {
	var c = SafeCounter{
		v: map[string]int{},
	}

	for i := 0; i < 1000; i++ {
		go c.Increase("some key")
	}

	time.Sleep(time.Second * 1 )
	fmt.Println(c.Value("some key"))
}

