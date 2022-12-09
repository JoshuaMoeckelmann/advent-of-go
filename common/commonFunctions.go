package common

import (
	"bufio"
	"log"
	"math"
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SignOrNull(num int) int {
	if num == 0 {
		return 0
	}

	return int(math.Copysign(1, float64(num)))
}
