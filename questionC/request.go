package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Handler struct {
	URLs       []*url.URL
	waiteGroup *sync.WaitGroup
	lock       *sync.Mutex
	response   []*Response
}

type Response struct {
	URL  *url.URL
	Body string
}

func NewHandler() *Handler {
	return &Handler{
		URLs:       make([]*url.URL, 0),
		waiteGroup: new(sync.WaitGroup),
		lock:       new(sync.Mutex),
		response:   make([]*Response, 0),
	}
}

func (h *Handler) HandleURLs(urls []*url.URL) {
	for _, u := range urls {
		go h.doNetworkRequest(u)
	}
	h.printResponses(len(urls))
}

func (h *Handler) doNetworkRequest(address *url.URL) {
	h.waiteGroup.Add(1)
	defer h.waiteGroup.Done()
	request, err := http.NewRequest(http.MethodGet, address.Path, nil)
	if err != nil {
		log.WithError(err).Errorf("failed to create request")
		return
	}

	client := &http.Client{
		Timeout: time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		log.WithError(err).Errorf("failed to make request")
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Errorf("request failed with status %v", response.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.WithError(err).Errorf("failed to handle body")
		return
	}

	h.lock.Lock()
	defer h.lock.Unlock()

	h.response = append(h.response, &Response{
		URL:  address,
		Body: string(body),
	})
}

func (h *Handler) printResponses(totalRequests int) {
	h.waiteGroup.Wait()
	if len(h.response) == totalRequests {
		for _, s := range h.response {
			log.Infof("respons eof url %v is:\n%v", s.URL, s.Body)
		}
	} else {
		fmt.Println("some requests failed")
	}
}
