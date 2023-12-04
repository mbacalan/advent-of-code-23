package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func calculateSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	sum := 1
	for i := 1; i < len(arr); i++ {
		sum *= 2
	}

	return sum
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
		line := scanner.Text()
		colonSplit := strings.Split(line, ":")
		numberSplit := strings.Split(colonSplit[1], "|")
		winningNumbers := strings.Split(strings.TrimSpace(numberSplit[0]), " ")
		ourNumbers := strings.Split(strings.TrimSpace(numberSplit[1]), " ")

		var overlappingNumbers []int
		for _, ourNum := range ourNumbers {
			if ourNum == "" {
				continue
			}

			num, err := strconv.Atoi(ourNum)
			if err != nil {
				panic (err)
			}

			if slices.Contains(winningNumbers, ourNum) {
				overlappingNumbers = append(overlappingNumbers, num)
			}
		}

		sum += calculateSum(overlappingNumbers)
	}

	fmt.Printf("sum is %d \n",sum)
}
