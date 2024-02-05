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

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func addFood(filename string, food string, calories string, protein string, servings string) {
	entries := readFoodLog(filename)

	reader := bufio.NewReader(os.Stdin)
	if food == "" {
		fmt.Print("Enter food: ")
		food, _ = reader.ReadString('\n')
	}
	if calories == "" {
		for _, entry := range entries {
			if entry.food == food {
				calories = strconv.FormatFloat(entry.calories, 'f', -1, 64)
			}
		}
		if calories == "" {
			fmt.Print("Enter calories: ")
			calories, _ = reader.ReadString('\n')
		}
	}
	if protein == "" {
		for _, entry := range entries {
			if entry.food == food {
				protein = strconv.FormatFloat(entry.protein, 'f', -1, 64)
			}
		}
		if protein == "" {
			fmt.Print("Enter protein: ")
			protein, _ = reader.ReadString('\n')
		}
	}
	if servings == "" {
		fmt.Print("Enter servings: ")
		servings, _ = reader.ReadString('\n')
	}

	file, err := os.OpenFile(
		filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	date := time.Now().Format("2006-01-02")
	_, err = file.WriteString(
		fmt.Sprintf(
			"%s	%s	%s	%s	%s\n",
			date,
			food,
			calories,
			protein,
			servings,
		),
	)
	if err != nil {
		fmt.Println(err)
	}
}

func showFood(filename string) {
	entries := readFoodLog(filename)

	fmt.Println("Date        Cal  Pro  Srv Food")
	for _, entry := range entries {
		fmt.Printf(
			"%s % 4.0f % 4.0f % 4.0f %s\n",
			entry.date,
			entry.calories,
			entry.protein,
			entry.servings,
			entry.food,
		)
	}
}

func main() {
}
