package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var url = []string{

	`https://pkg.go.dev/`,

	`https://github.com/`,

	`abc.com/1234`,
}

type response struct {
	url  string
	resp *http.Response
	err  error
}

// var wg = &sync.WaitGroup{}

var wgWorker = sync.WaitGroup{}

func main() {

	//create a struct which stores the response, error and the url name

	//
	respChan := make(chan response)
	go doGetRequest(url, respChan)
	fetchResults(respChan)

}

func doGetRequest(urls []string, respChan chan response) {

	for _, url := range urls {

		wgWorker.Add(1)
		go func(url string) {

			defer wgWorker.Done()
			resp, err := http.Get(url)

			r := response{
				url:  url,
				resp: resp,
				err:  err,
			}

			respChan <- r //sending the resp struct to respCh

		}(url)

	}

	//waiting for get request to finish, then we would close the
	// channel to signal no more signals would be sent
	wgWorker.Wait()
	close(respChan)

	fetchResults(respChan)

}

func fetchResults(respChan chan response) {

	// Loop through every element in respChan

	for r := range respChan {

		// If there is an error with the response object, log it and continue with the next iteration

		if r.err != nil {

			log.Println(r.err)

			continue

		}

		// Attempt to read the entire response body

		bytes, err := io.ReadAll(r.resp.Body)

		// If an error occurred during reading, log the error and continue with the next iteration

		if err != nil {

			log.Println(err)

			continue

		}

		// Ensuring that the response body is eventually closed, to prevent potential memory leak

		func() {

			defer r.resp.Body.Close()

		}()

		// If the status code indicates an unsuccessful response (i.e., is greater than 299)

		// log the status code and the body content, then move on to the next iteration

		if r.resp.StatusCode > 299 {

			log.Printf("Response failed with status code: %d and\nbody: %s\n", r.resp.StatusCode, bytes)

			continue

		}

		// If the response is successful, print out the URL associated with it and its status

		fmt.Println(r.url, r.resp.Status)

	}

	//some more work

}
