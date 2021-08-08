package basic

import "fmt"

func traverseSliceWithIndex() {
	var s = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d,", s[i])
	}
	fmt.Println()
}

func traverseSliceWithRange() {
	var s = []int{10, 11, 12, 13, 14, 15}
	for index, val := range s {
		fmt.Printf("index=%d, val=%d", index, val)
		fmt.Println()
	}
	fmt.Println()
}

func traverseMap() {
	var m = map[string]int{
		"jack": 100,
		"rose": 89,
		"lucy": 93,
	}
	for k, v := range m {
		fmt.Printf("key=%s, value=%d", k, v)
		fmt.Println()
	}
	fmt.Println()
}

func traverseChannel() {
	var c = make(chan int)
	go fillChannel(c)

	for i := range c {
		fmt.Printf("get value from channel: %d", i)
		fmt.Println()
	}
	fmt.Println()
}

func fillChannel(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
