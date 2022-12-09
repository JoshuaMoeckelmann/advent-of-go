package common

import (
	"bufio"
	"io"
	"log"
	"os"
)

func PrepareFile(token string, day int, purgeCachedFile bool) []string {
	input := inputData{
		token:           token,
		day:             day,
		purgeCachedFile: purgeCachedFile,
	}
	file := scrapeInput(input)
	return readWholeFile(file)
}

func LoadLocalFile() []string {
	file, err := os.Open("input.txt")
	handleError(err)
	return readWholeFile(file)
}

func createNewScanner(file *os.File) *bufio.Scanner {
	resetFile(file)
	return bufio.NewScanner(file)
}

func readWholeFile(file *os.File) []string {
	scanner := createNewScanner(file)
	allLines := make([]string, 0)
	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}

	CheckScannerForError(scanner)
	return allLines
}

func resetFile(file *os.File) {
	_, err := file.Seek(0, io.SeekStart)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
