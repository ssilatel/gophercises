package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func main() {
	filename := flag.String("f", "problems.csv", "Path to the CSV file for the quiz")
	flag.Parse()

	in, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	reader := csv.NewReader(in)

	lines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	problems := parseLines(lines)

	scanner := bufio.NewScanner(os.Stdin)
	var correct int
	for i, p := range problems {
		fmt.Printf("Problem %d: %s\n", i+1, p.q)
		scanner.Scan()
		if scanner.Text() == p.a {
			correct++
		}
	}

	fmt.Printf("\nYou answered %d out of %d questions correctly.\n", correct, len(problems))
}
