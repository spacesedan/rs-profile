package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func MakeMoralisGetRequest(url string, target interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Could not make a new request: %v\n", err)
		return err
	}

	req.Header.Set("X-API-KEY", os.Getenv("MORALIS_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Could not make get response: %v\n", err)
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Could not read content of reponse body: %v\n", err)
		return err
	}

	return json.Unmarshal(body, &target)

}
