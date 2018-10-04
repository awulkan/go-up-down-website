package main

import (
	"net/http"
	"regexp"
	"strings"
	"time"
)

var urlRegx = regexp.MustCompile(`(?mi)^(http://|https://)?[a-z0-9_.-]+\.[a-z]+$`)
var httpClient = http.Client{
	Timeout: 5 * time.Second,
}

// checkURL performs a GET request tothe provided URL and returns
// a boolean representing if the GET was successful or not.
func checkURL(url string) bool {
	if _, err := httpClient.Get(url); err != nil {
		return false
	}
	return true
}

// checkDifferentURLs (in lack of a better name) sequentially checks
// a domain with both the secure and non-secure http protocols.
func checkDifferentURLs(domain string) bool {
	if !checkURL("https://" + domain) {
		if !checkURL("http://" + domain) {
			return false
		}
	}
	return true
}

// IsSiteUp checks if a site is up or not. It returns a boolean value
// of 'true' if it's reachable, and 'false' if it's unreachable.
func IsSiteUp(domain string) bool {
	if strings.HasPrefix(domain, "https://") {
		return checkURL(domain)
	}
	if strings.HasPrefix(domain, "http://") {
		return checkURL(domain)
	}
	return checkDifferentURLs(domain)
}

// IsValidDomain is a simple naive approach to check if the URL provided is valid.
// It's definitely not enough to guarantee that the string is 100% an URL, but for
// this simple program it will do.
func IsValidDomain(url string) bool {
	if l := len(url); l == 0 || l > 300 {
		return false
	}
	if urlRegx.MatchString(url) {
		return true
	}
	return false
}
