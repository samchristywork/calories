package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

	food = strings.Trim(food, " \n")
	calories = strings.Trim(calories, " \n")
	protein = strings.Trim(protein, " \n")
	servings = strings.Trim(servings, " \n")

	if !isNumeric(calories) {
		fmt.Println("Calories must be a number")
		return
	}

	if !isNumeric(protein) {
		fmt.Println("Protein must be a number")
		return
	}

	if !isNumeric(servings) {
		fmt.Println("Servings must be a number")
		return
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

	fmt.Println("Food added")
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

func duration(seconds int) string {
	if seconds < 0 {
		seconds = -seconds
	}
	h := seconds / 3600
	m := (seconds % 3600) / 60
	s := seconds % 60

	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	} else if m > 0 {
		return fmt.Sprintf("%d:%02d", m, s)
	} else {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}
}

func summary(filename string) {
	entries := readFoodLog(filename)

	startDate, err := time.Parse("2006-01-02", entries[0].date)
	if err != nil {
		fmt.Println(err)
		return
	}

	now := time.Now()

	days := now.Sub(startDate).Hours() / 24

	calories := 0.0
	protein := 0.0

	for _, entry := range entries {
		servings := entry.servings
		calories += entry.calories * servings
		protein += entry.protein * servings
	}

	predictedCalories := days * 1800
	predictedProtein := days * 100

	caloriesHours := (predictedCalories - calories) / 1800 * 24
	proteinHours := (predictedProtein - protein) / 100 * 24

	reset := "\033[0m"
	red := "\033[31m"
	green := "\033[32m"
	if caloriesHours < 0 {
		fmt.Printf("%s", red)
		fmt.Printf("Cal: % 4.0f (%s)\n", calories-predictedCalories, duration(int(caloriesHours*3600)))
		fmt.Printf("%s", reset)
	} else {
		fmt.Printf("%s", green)
		fmt.Printf("Cal: % 4.0f (%s)\n", calories-predictedCalories, duration(int(caloriesHours*3600)))
		fmt.Printf("%s", reset)
	}
	if proteinHours < 0 {
		fmt.Printf("%s", green)
		fmt.Printf("Pro: % 4.0f (%s)\n", protein-predictedProtein, duration(int(proteinHours*3600)))
		fmt.Printf("%s", reset)
	} else {
		fmt.Printf("%s", red)
		fmt.Printf("Pro: % 4.0f (%s)\n", protein-predictedProtein, duration(int(proteinHours*3600)))
		fmt.Printf("%s", reset)
	}
}

func main() {
}
