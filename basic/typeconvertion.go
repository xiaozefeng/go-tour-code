package basic

import (
	"fmt"
	"strconv"
)

func number2number() {
	var x int = 1
	var y = float64(x)
	var z = int(x)
	fmt.Println(x, y, z)
}

func number2string() {
	var x string = "1"
	var y, err = strconv.Atoi(x)
	if err != nil {
		fmt.Println("int to string error", err)
	} else {
		fmt.Println("convert string to int, ", y)
	}

	var a = 10
	var b = strconv.Itoa(a)
	fmt.Println("convert int to string,", b)
}

func byteArray2string() {
	var s = "hello"
	var bs = []byte(s)
	fmt.Println("convert string to byte array: ", bs)

	s = string(bs)
	fmt.Println("convert byte array to string: ", s)
}

func byteNumberToString() {
	var b byte = 127
	var a = int(b)
	var s = strconv.Itoa(a)
	fmt.Println("convert byte to string: ", s)
}
