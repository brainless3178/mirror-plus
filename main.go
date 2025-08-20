package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	method := flag.String("method", "GET", "HTTP method to use (GET or POST)")
	data := flag.String("data", "", "POST data to send")
	minLen := flag.Int("minlen", 3, "Minimum length of reflected value to check")
	flag.Parse()

	var urls []string

	if flag.NArg() > 0 {
		urls = flag.Args()
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
	}

	for _, url := range urls {
		TestReflections(url, *method, *data, *minLen)
	}
}
