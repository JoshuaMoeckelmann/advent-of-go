package common

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func ConvertStringToInt(stringToConvert string) int {
	i, err := strconv.Atoi(strings.TrimSpace(stringToConvert))

	if err != nil {
		panic(err)
	}

	return i
}

func CheckScannerForError(scanner *bufio.Scanner) {
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
