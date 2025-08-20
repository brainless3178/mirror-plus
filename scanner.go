package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func TestReflections(rawurl, method, postData string, minLen int) {
	parsed, err := url.Parse(rawurl)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		fmt.Printf("[-] Invalid URL: %s\n", rawurl)
		return
	}

	var resp *http.Response

	if strings.ToUpper(method) == "POST" {
		resp, err = http.Post(rawurl, "application/x-www-form-urlencoded", strings.NewReader(postData))
	} else {
		resp, err = http.Get(rawurl)
	}

	if err != nil {
		fmt.Printf("[-] Request failed: %s\n", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	body := string(bodyBytes)

	// Check query parameters
	for key, values := range parsed.Query() {
		for _, val := range values {
			if len(val) < minLen {
				continue
			}
			checkReflections(key, val, body, resp)
		}
	}

	// Check POST data
	if method == "POST" && postData != "" {
		for _, param := range strings.Split(postData, "&") {
			parts := strings.SplitN(param, "=", 2)
			if len(parts) != 2 || len(parts[1]) < minLen {
				continue
			}
			checkReflections(parts[0], parts[1], body, resp)
		}
	}

	// Check path segments
	for _, segment := range strings.Split(parsed.Path, "/") {
		if len(segment) >= minLen && strings.Contains(body, segment) {
			fmt.Printf("[/] Path segment reflected: '%s' in body of %s\n", segment, rawurl)
		}
	}
}

func checkReflections(key, val, body string, resp *http.Response) {
	if strings.Contains(body, val) {
		fmt.Printf("[+] Reflected in body: %s=%s\n", key, val)
	}
	encodedVal := url.QueryEscape(val)
	if strings.Contains(body, encodedVal) {
		fmt.Printf("[+] URL-encoded reflection in body: %s=%s\n", key, encodedVal)
	}
	// Check headers
	for hKey, hVals := range resp.Header {
		for _, h := range hVals {
			if strings.Contains(h, val) {
				fmt.Printf("[!] Reflected in header '%s': %s=%s\n", hKey, key, val)
			}
		}
	}
}
