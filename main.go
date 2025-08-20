package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var urls []string

	// Read from stdin or args
	if len(os.Args) > 1 {
		urls = os.Args[1:]
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
	}

	for _, url := range urls {
		TestReflections(url)
	}
}
