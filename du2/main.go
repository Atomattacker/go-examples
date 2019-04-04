package main

import (
	"flag"
	"fmt"
	"net/http"
	"reflect"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	x := 123
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t.String())
	fmt.Println(v.String())

	http.ListenAndServe
}
