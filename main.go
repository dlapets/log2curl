package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/moul/http2curl"
)

type requestLike struct {
	Form   url.Values
	Header http.Header
	Method string
	Path   string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give a base url!")
		os.Exit(1)
	}
	target := os.Args[1]
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		b := scanner.Bytes()

		rl := &requestLike{}
		if err := json.Unmarshal(b, rl); err != nil {
			log.Printf("unmarshal failed: %s from %s", err, b)
			continue
		}

		req, err := http.NewRequest(
			rl.Method,
			target+rl.Path,
			strings.NewReader(rl.Form.Encode()), // TODO support non-form body
		)
		if err != nil {
			log.Printf("failed to get request %s from: %v", rl, err)
			continue
		}
		req.Header = rl.Header

		command, err := http2curl.GetCurlCommand(req)
		if err != nil {
			log.Printf("failed to get command: %s from: %v", err, rl)
			continue
		}
		fmt.Println(command)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}
