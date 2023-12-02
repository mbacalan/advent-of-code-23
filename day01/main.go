package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// regex doesn't work with overlapping digits: eightwo
var digits = map[string]string{
	"one":   "1",
	"1":     "1",
	"two":   "2",
	"2":     "2",
	"three": "3",
	"3":     "3",
	"four":  "4",
	"4":     "4",
	"five":  "5",
	"5":     "5",
	"six":   "6",
	"6":     "6",
	"seven": "7",
	"7":     "7",
	"eight": "8",
	"8":     "8",
	"nine":  "9",
	"9":     "9",
}

func main() {
	f, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0

	for scanner.Scan() {
		digitBuffer := ""
		line := scanner.Text()

		// loop through characters in the line,
		// creating a new string without the first character each time
		// to check for prefixes that exist in digits
		for line != "" {
			for k, v := range digits {
				if strings.HasPrefix(line, k) {
					digitBuffer += v
					break
				}
			}

			line = line[1:]
		}

		if digitBuffer == "" {
			continue
		}

		// convert the first and last characters in digitBuffer to integers
		// and put them together. "1234" -> 14
		// buf[0] = '5', buf[0] - '0' = 5
		num := int(digitBuffer[0] - '0') * 10 + int(digitBuffer[len(digitBuffer) - 1] - '0')
		sum += num
	}

	fmt.Printf("Calibration sum is %d \n", sum)
}
