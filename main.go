package main

import "fmt"

func main() {
	str := "ababa"
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		fmt.Println("i:", i, ", j:", j)
	}
}
