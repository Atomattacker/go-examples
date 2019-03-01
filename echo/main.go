package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now().UnixNano()

	for i := 0; i < 1000; i++ {
		s, sep := "", ""
		for _, v := range os.Args[1:] {
			s += sep + v
			sep = " "
		}
	}
	end := time.Now().UnixNano()
	fmt.Printf("1:%v\n", end-start)

	start = time.Now().UnixNano()
	for i := 0; i < 1000; i++ {
		strings.Join(os.Args[1:], " ")
	}
	end = time.Now().UnixNano()
	fmt.Printf("2:%v", end-start)

	//fmt.Println(s)
}
