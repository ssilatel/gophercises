package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("f", "problems.csv", "Path to the CSV file for the quiz")
	flag.Parse()

	in, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	reader := csv.NewReader(in)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var correct int
	for _, line := range records {
		fmt.Println(line[0])
		scanner.Scan()
		if scanner.Text() == line[1] {
			correct++
		}
	}

	fmt.Printf("\nYou answered %d out of %d questions correctly\n", correct, len(records))
}
