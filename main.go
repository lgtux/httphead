package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	Tick = "\u2713"
)

func main() {
	var url string

	if len(os.Args) == 1 {
		url = "https://golang.org"
	} else {
		url = os.Args[1]
		if !strings.Contains(url, "http") {
			url = "http://" + url
		}
	}
	res, err := http.Head(url)
	if err != nil {
		panic(err)
	}

	// Determine output widths
	var maxKeyWidth int

	for k := range res.Header {
		if len(k) > maxKeyWidth {
			maxKeyWidth = len(k)
		}
	}

	if res.Status == "200 OK" {
		fmt.Printf("%[1]*v: %v %v\n", maxKeyWidth, "HTTP Status", res.Status, Tick)
	} else {
		fmt.Printf("%[1]*v: %v\n", maxKeyWidth, "HTTP Status", res.Status)
	}
	fmt.Printf("%[1]*v: %v\n", maxKeyWidth, "Protocol", res.Proto)

	for k, v := range res.Header {
		fmt.Printf("%[1]*v: %v\n", maxKeyWidth, k, v)
	}

}
