package main

import "fmt"

func main() {
	str := "Hello, 世界"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v\t", str[i])
	}

}
