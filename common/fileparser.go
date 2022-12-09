package common

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/hyrmn/lc/pkg/lc"
)

func PrepareFile(token string, day int) (*bufio.Scanner, int, *os.File, []string) {
	input := inputData{
		token: token,
		day:   day,
	}
	file := scrapeInput(input)

	// TODO remove
	resetFile(file)
	lineCount, _ := lc.CountLines(bufio.NewReader(file))
	resetFile(file)
	scanner := bufio.NewScanner(file)

	return scanner, lineCount, file, readWholeFile(CreateNewScanner(file))
}

func LoadLocalFile() (*bufio.Scanner, int, *os.File, []string) {
	file, err := os.Open("input.txt")
	handleError(err)

	resetFile(file)
	lineCount, _ := lc.CountLines(bufio.NewReader(file))
	return CreateNewScanner(file), lineCount, file, readWholeFile(CreateNewScanner(file))
}

func CreateNewScanner(file *os.File) *bufio.Scanner {
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
