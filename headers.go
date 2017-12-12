package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: headers <url>")
		return
	}
	response, _ := http.Get(os.Args[1])
	for k, v := range response.Header {
		fmt.Printf("%s: ", k)
		for _, val := range v {
			fmt.Print(val)
		}
		fmt.Println()
	}
}
