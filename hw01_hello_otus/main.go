package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	strToReverse := "Hello, OTUS!"
	reversedStr := stringutil.Reverse(strToReverse)

	fmt.Println(reversedStr)
}
