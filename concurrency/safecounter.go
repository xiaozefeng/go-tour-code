package concurrency

import (
	"errors"
	"sync"
)

type SafeCounter struct {
	mutex sync.Mutex
	v map[string]int
}

func (c *SafeCounter) Increase(key string)  {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.v[key] ++
}

func (c *SafeCounter) Value(key string) (int,error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if val, ok := c.v[key]; ok  {
		return val,nil
	}
	return 0, errors.New("key not exist")
}
