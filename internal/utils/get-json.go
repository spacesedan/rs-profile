package utils

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

// getJson send HTTP request and return JSON response
func GetJson(url string, target interface{}) error {
	apiKey := os.Getenv("OS_KEY")

	req := fasthttp.AcquireRequest()
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.SetRequestURI(url)
	defer fasthttp.ReleaseRequest(req)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	if err := fasthttp.Do(req, res); err != nil {
		log.Fatalln(err)
	}

	body := res.Body()

	return json.Unmarshal(body, &target)

}
