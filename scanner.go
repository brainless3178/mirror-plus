package main

import (
	"net/http"
	"io"
	"net/url"
	"strings"
)

func TestReflections(rawurl string) {
	parsed, err := url.Parse(rawurl)
	if err != nil || parsed.RawQuery == "" {
		return
	}

	resp, err := http.Get(rawurl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	bodyStr := string(body)

	queryParams := parsed.Query()
	for key, values := range queryParams {
		for _, val := range values {
			if strings.Contains(bodyStr, val) {
				fmt.Printf("[+] Reflected: %s=%s in body of %s\n", key, val, rawurl)
			}
			for name, header := range resp.Header {
				for _, h := range header {
					if strings.Contains(h, val) {
						fmt.Printf("[!] Reflected: %s=%s in header %s\n", key, val, name)
					}
				}
			}
		}
	}
}
