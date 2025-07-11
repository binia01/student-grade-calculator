package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func average(grades map[string]int) float64 {
	total := 0
	for _, grade := range grades {
		total += grade
	}
	return float64(total) / float64(len(grades))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Please enter the number of subjects: ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	num, err := strconv.Atoi(numStr)
	if err != nil || num <= 0 {
		fmt.Println("Invalid number of subjects. Exiting...")
		return
	}

	scores := make(map[string]int)

	for i := 0; i < num; i++ {
		fmt.Printf("\n Enter subject #%d name: ", i+1)
		subject, _ := reader.ReadString('\n')
		subject = strings.TrimSpace(subject)

		var score int
		for {
			fmt.Printf("Enter score for %s (0–100): ", subject)
			scoreStr, _ := reader.ReadString('\n')
			scoreStr = strings.TrimSpace(scoreStr)

			s, err := strconv.Atoi(scoreStr)
			if err != nil || s < 0 || s > 100 {
				fmt.Println("Invalid score. Please enter a number between 0 and 100.")
				continue
			}
			score = s
			break
		}

		scores[subject] = score
		fmt.Println("Grade added!")
	}

	gradeAverage := average(scores)

	fmt.Printf("Student Name: %s\n", name)
	fmt.Println("Subject Grades:")
	for subject, score := range scores {
		fmt.Printf("  - %s: %d\n", subject, score)
	}
	fmt.Printf("Total Average: %.2f\n", gradeAverage)
}
