package requests

import (
	"fmt"
	"net/http"
	"time"
)

func Trace(url string) (response *http.Response, err error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	request, err := http.NewRequest("TRACE", url, nil)
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

func Options(url string) (response *http.Response, err error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	request, err := http.NewRequest("OPTIONS", url, nil)
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

func Connect(url string) (response *http.Response, err error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	request, err := http.NewRequest("CONNECT", url, nil)
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
