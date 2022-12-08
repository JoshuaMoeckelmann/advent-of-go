package common

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/hyrmn/lc/pkg/lc"
)

func PrepareFile(token string, day int) (*bufio.Scanner, int, *os.File) {
	input := inputData{
		token: token,
		day:   day,
	}
	file := scrapeInput(input)

	resetFile(file)
	lineCount, _ := lc.CountLines(bufio.NewReader(file))
	resetFile(file)
	scanner := bufio.NewScanner(file)

	return scanner, lineCount, file
}

func LoadLocalFile() (*bufio.Scanner, int, *os.File) {
	file, err := os.Open("input.txt")
	handleError(err)

	resetFile(file)
	lineCount, _ := lc.CountLines(bufio.NewReader(file))
	return CreateNewScanner(file), lineCount, file
}

func CreateNewScanner(file *os.File) *bufio.Scanner {
	resetFile(file)
	return bufio.NewScanner(file)
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
