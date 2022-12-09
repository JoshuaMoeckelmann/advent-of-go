package common

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type inputData struct {
	token           string
	day             int
	purgeCachedFile bool
}

func scrapeInput(inputs inputData) *os.File {
	// Create the file
	file, err, fileAlreadyExists := getOrCreateAocFileCurrentDay(inputs)
	if err != nil {
		log.Fatal(err)
	}
	if fileAlreadyExists {
		log.Default().Printf("Input exists for day %d, opening...\n", inputs.day)
		return file
	}
	log.Default().Printf("Downloading input for day %d...\n", inputs.day)

	// Get the data
	url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", inputs.day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("client: could not create request: %s\n", err)
	}
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", inputs.token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("client: could not create request: %s\n", err)
	}
	defer resp.Body.Close()

	// Writer the body to file
	if _, err := io.Copy(file, resp.Body); err != nil {
		log.Fatal(err)
	}
	log.Default().Println("Successfully download!")

	return file
}

func getOrCreateAocFileCurrentDay(inputs inputData) (*os.File, error, bool) {
	dir := filepath.Join(os.TempDir(), "AOC")
	err := os.MkdirAll(dir, os.ModePerm)
	handleError(err)

	fileName := filepath.Join(dir, fmt.Sprintf("AOC-day-%d-input", inputs.day))
	if _, err := os.Stat(fileName); err == nil || inputs.purgeCachedFile {
		if inputs.purgeCachedFile {
			file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, 0666)
			return file, err, false
		}

		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
		contentOfFile := readWholeFile(file)
		return file, err, len(contentOfFile) > 0
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	return file, err, false
}
