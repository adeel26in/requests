// Package requests allows you to make HTTP requests easily.
package requests

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(url string) (response *http.Response, err error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// request.Header.Add("DNT", "1") // Do Not Track Header for Privacy (DEPRECIATED)
	request.Header.Add("Sec-GPC", "1")                // Global Privacy Control Header
	request.Header.Add("Sec-CH-UA", "Chromium;v=143") // User agent for Chromium

	response, err = client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	return response, nil
}

func Head(url string) (response *http.Response, err error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	request, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// request.Header.Add("DNT", "1") // Do Not Track Header for Privacy (DEPRECIATED)
	request.Header.Add("Sec-GPC", "1")                // Global Privacy Control Header
	request.Header.Add("Sec-CH-UA", "Chromium;v=143") // User agent for Chromium

	response, err = client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	return response, nil
}

func POST(url string, data io.Reader) (response *http.Response, err error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	request, err := http.NewRequest("POST", url, data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response, err = client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	return response, nil
}
