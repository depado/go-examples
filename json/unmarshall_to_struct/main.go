package main

import (
	"encoding/json"
	"net/http"
)

// Fetch the url given in parameter and unmarshall the JSON to the given out struct.
// Usage of a client for example for the Github API, request the v3 automatically
// with the Accept header.
func fetchURLwithClient(url string, out interface{}) (err error) {
	c := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	// Add the API version in Headers
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(out)
	return
}

// Fetch the url given in parameter and unmarshall the JSON to the given out struct
func fetchURL(url string, out interface{}) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(out)
	return
}

func main() {

}
