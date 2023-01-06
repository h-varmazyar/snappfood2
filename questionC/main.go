package main

import (
	"net/url"
)

func main() {
	urls := make([]*url.URL, 0)
	handler := NewHandler()
	handler.HandleURLs(urls)
}
