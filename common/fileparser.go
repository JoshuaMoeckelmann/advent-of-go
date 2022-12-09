package common

import (
	"bufio"
	"io"
	"log"
	"os"
)

func PrepareFile(token string, day int) []string {
	input := inputData{
		token: token,
		day:   day,
	}
	file := scrapeInput(input)
	return readWholeFile(createNewScanner(file))
}

func LoadLocalFile() []string {
	file, err := os.Open("input.txt")
	handleError(err)
	return readWholeFile(createNewScanner(file))
}

func createNewScanner(file *os.File) *bufio.Scanner {
	resetFile(file)
	return bufio.NewScanner(file)
}

func readWholeFile(scanner *bufio.Scanner) []string {
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
