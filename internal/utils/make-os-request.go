package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func MakeOpenSeaRequest(url string, target interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Could not create request: %v\n", err)
		return err
	}
	req.Header.Set("X-API-KEY", os.Getenv("OS_KEY"))

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Could not get a response: %v\n", err)
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Could not read response body: %v\n", err)
		return err
	}

	if res.StatusCode == http.StatusTooManyRequests {
		fmt.Printf("Got 429, Retry-After value: %v\n", res.Header.Get("Retry-After"))
		ta, _ := strconv.Atoi(res.Header.Get("Retry-After"))
		duration := time.Duration(ta)
		fmt.Printf("Experiencing throttling from OS retrying request in: %v seconds\n", duration*time.Second)
		time.Sleep(duration * time.Second)
		client.Do(req)

		res, err := client.Do(req)
		if err != nil {
			log.Printf("Could not get a response: %v\n", err)
			return err
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("Could not read response body: %v\n", err)
			return err
		}

		return json.Unmarshal(body, &target)
	}
	if res.StatusCode != http.StatusOK {
		log.Printf("STATUS CODE: %v\n", res.StatusCode)
		log.Printf("RES HEADERS: %v\n", res.Header.Clone())
		log.Printf("BODY: %v\n", string(body))
	}
	return json.Unmarshal(body, &target)

}
