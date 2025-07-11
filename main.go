package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func average(x map[string]int, y int) float64 {
	ttl := 0
	for _, score := range x {
		ttl += score
	}
	return float64(ttl / y)

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Please enter number of grades: ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	var num int
	fmt.Sscanf(numStr, "%d", &num)

	scores := make(map[string]int)
	for i := 0; i < num; i++ {
		fmt.Print("Please enter a subject: ")
		subject, _ := reader.ReadString('\n')
		subject = strings.TrimSpace(subject)

		fmt.Print("Enter corresponding score: ")
		scoreStr, _ := reader.ReadString('\n')
		scoreStr = strings.TrimSpace(scoreStr)
		var score int
		fmt.Sscanf(scoreStr, "%d", &score)
		fmt.Println("Added!")
		scores[subject] = score
	}

	grade_average := average(scores, num)

	fmt.Println("Student Grade")
	fmt.Println("Name: ", name)
	fmt.Println("These are the subjects and grades you have:")
	for subject, score := range scores {
		fmt.Printf("%s: %d\n", subject, score)
	}
	fmt.Println("Total Grade Average: ", grade_average)

}
