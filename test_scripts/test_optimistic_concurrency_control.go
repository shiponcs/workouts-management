package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const (
	url   = "http://localhost:8080/workouts/11"
	token = "NZU2NF2PO7VZHPXC3FDA2YPXLUQM4PUSCMFPGJJT5CPSUCHP5YRQ"
)

var payloadTemplate = `{
  "title": "%s",
  "description": "Test for concurrent modification conflict",
  "duration_minutes": 45,
  "calories_burned": 250,
  "version": 4,
  "entries": [
    {
      "exercise_name": "Walking",
      "sets": 1,
      "duration_seconds": 2700,
      "weight": 0,
      "notes": "Keep a steady pace",
      "order_index": 1
    }
  ]
}`

func makeRequest(title string, wg *sync.WaitGroup) {
	defer wg.Done()

	payload := fmt.Sprintf(payloadTemplate, title)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Printf("[%s] Error creating request: %v\n", title, err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[%s] Error making request: %v\n", title, err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("[%s] Status: %d\n", title, resp.StatusCode)
	fmt.Printf("[%s] Body: %s\n\n", title, body)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go makeRequest("First Update", &wg)
	go makeRequest("Second Update", &wg)

	wg.Wait()
}
