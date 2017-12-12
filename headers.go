package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: headers <url>")
		return
	}
	url := os.Args[1]
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
		fmt.Printf("prepended http:// to %s...\n", os.Args[1])
	}
	response, _ := http.Get(url)
	for k, v := range response.Header {
		fmt.Printf("%s: ", k)
		for _, val := range v {
			fmt.Print(val)
		}
		fmt.Println()
	}
}
