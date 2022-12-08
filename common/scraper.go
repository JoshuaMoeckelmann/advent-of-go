package common

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type inputData struct {
	token string
	day   int
}

func scrapeInput(inputs inputData) *os.File {
	// Create the file
	file, err := ioutil.TempFile("", fmt.Sprintf("AOC-day-%d-", inputs.day))
	if err != nil {
		log.Fatal(err)
	}

	// Get the data
	url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", inputs.day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", inputs.token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Writer the body to file
	if _, err := io.Copy(file, resp.Body); err != nil {
		log.Fatal(err)
	}

	return file
}
