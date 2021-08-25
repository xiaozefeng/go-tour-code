package main

import (
	"sync"
	"time"
)

func sendValueOnCloseChannel() {
	ch := make(chan int)
	close(ch)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ch <- 1
		time.Sleep(time.Second * 2)
		wg.Done()
	}()
	wg.Wait()
}
