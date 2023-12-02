package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var totalCubes = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
}

func getGameID(s string) int {
	colonSplit := strings.Split(s, ":")
	spaceSplit := strings.Split(colonSplit[0], " ")
	number, err := strconv.Atoi(spaceSplit[1])

	if err != nil {
		panic (err)
	}

	return number
}

func getGameSets(s string) []string {
	colonSplit := strings.Split(s, ":")
	semiSplit := strings.Split(colonSplit[1], ";")
	var sets []string

	for _, split := range semiSplit {
		sets = append(sets, strings.TrimSpace(split))
	}

	return sets
}

func getCubes(s string) map[string]int {
	var red int
	var green int
	var blue int

	commaSplit := strings.Split(s, ",")

	for _, cubeSet := range commaSplit {
		set := strings.Split(strings.TrimSpace(cubeSet), " ")

		if set[1] == "red" {
			num, err := strconv.Atoi(set[0])

			if err != nil {
				panic (err)
			}

			red = num
		}

		if set[1] == "green" {
			num, err := strconv.Atoi(set[0])

			if err != nil {
				panic (err)
			}

			green = num
		}

		if set[1] == "blue" {
			num, err := strconv.Atoi(set[0])

			if err != nil {
				panic (err)
			}

			blue = num
		}
	}

	return map[string]int{"red": red, "green": green, "blue": blue}
}

func main() {
	f, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var validGames []int

	for scanner.Scan() {
		line := scanner.Text()
		gameID := getGameID(line)
		sets := getGameSets(line)

		if err != nil {
			panic (err)
		}

		var setValidity []bool
		for _, set := range sets {
			setCubes := getCubes(set)

			if setCubes["red"] > totalCubes["red"] || setCubes["green"] > totalCubes["green"] || setCubes["blue"] > totalCubes["blue"] {
				setValidity = append(setValidity, false)
				continue
			}

			setValidity = append(setValidity, true)
		}

		if !slices.Contains(setValidity, false) {
			validGames = append(validGames, gameID)
		}
	}

	sum := 0
	for _, game := range validGames {
		sum += game
	}

	fmt.Printf("valid games are %v. their sum is %d \n", validGames, sum)
}
