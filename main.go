package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fname := flag.String("csv", "problems.csv", "a csv file in the format  of 'question,answer'")
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
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
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
