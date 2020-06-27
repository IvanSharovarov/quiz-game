package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fname := flag.String("csv", "problems.csv", "a csv file in the format  of 'question,answer'")
	timelimit := flag.Int("limit", 30, "the limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*fname)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *fname))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the provided CSV file."))
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) (problems []problem) {
	problems = make([]problem, len(lines))
	for i, p := range lines {
		problems[i] = problem{
			q: p[0],
			a: strings.TrimSpace(p[1]),
		}
	}
	return
}

type problem struct {
	q string
	a string
}

func exit(s string) {
	fmt.Println(s)
	os.Exit(1)
}
