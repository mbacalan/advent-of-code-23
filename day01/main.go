package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func findDigits(input string) []string {
	re := regexp.MustCompile(`\d`)
	matches := re.FindAllString(input, -1)

	var output []string

	for _, match := range matches {
		output = append(output, match)
	}

	return output
}

func main() {
	f, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	calibration_sum := 0

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Printf("Calibration sum is %d \n", calibration_sum)
				return
			}

			panic(err)
		}

		calibration := findDigits(line)
		line_sum := calibration[0] + calibration[len(calibration)-1]
		line_sum_int, err := strconv.Atoi(line_sum)
		calibration_sum += line_sum_int
	}
}
