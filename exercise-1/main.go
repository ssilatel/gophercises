package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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
	totalTime := flag.Int("t", 30, "Total time for the quiz")
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

	timer := time.NewTimer(time.Duration(*totalTime) * time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	var correct int
	var input string
	for i, p := range problems {
		fmt.Printf("Problem %d: %s\n", i+1, p.q)
		inputCh := make(chan string)
		go func() {
			scanner.Scan()
			input = scanner.Text()
			inputCh <- input
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou answered %d out of %d questions correctly.\n", correct, len(problems))
			return
		case input := <-inputCh:
			if input == p.a {
				correct++
			}
		}
	}
}
