package main

import (
	"errors"
	"fmt"
	"net/http"
)

// error
var errRequestFailed error = errors.New("Request failed")

// hit url
func hitURL(url string, channel chan<- result) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	channel <- result{url: url, status: status}
}

// set results
func setResultsFromUrl(results map[string]string, url []string, channel chan result) map[string]string {
	for _, url := range url {
		go hitURL(url, channel)
	}

	for i := 0; i < len(url); i++ {
		result := <-channel
		results[result.url] = result.status
	}

	return results
}

// print results
func printResults(results map[string]string) {
	for url, status := range results {
		fmt.Println(url, status)
	}
}

type result struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	channel := make(chan result)
	url := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	// check url
	results = setResultsFromUrl(results, url, channel)

	// print results
	printResults(results)
}
