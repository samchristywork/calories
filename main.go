package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	date     string
	food     string
	calories float64
	protein  float64
	servings float64
}

func readFoodLog(filename string) []Entry {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	entries := make([]Entry, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "	")

		date := splitLine[0]
		name := splitLine[1]
		calories := splitLine[2]
		protein := splitLine[3]
		servings := splitLine[4]

		caloriesFloat, _ := strconv.ParseFloat(calories, 64)
		proteinFloat, _ := strconv.ParseFloat(protein, 64)
		servingsFloat, _ := strconv.ParseFloat(servings, 64)

		entries = append(entries, Entry{
			date:     date,
			food:     name,
			calories: caloriesFloat,
			protein:  proteinFloat,
			servings: servingsFloat,
		})
	}

	return entries
}

func main() {
}
