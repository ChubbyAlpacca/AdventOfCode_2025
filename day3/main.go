package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "data/test_input.txt"

func formatInputData(data []byte) []string {
	s := strings.TrimSpace(string(data))
	if s == "" {
		return []string{}
	}
	return strings.Split(s, "\n")
}

// parseToGridStrings converts raw file bytes into a slice-of-slices of strings.
// Behavior:
// - If a line contains spaces, it is split on whitespace into tokens.
// - Else if a line contains commas, it is split on commas.
// - Otherwise the line is split into single-character strings (useful for digit grids).
func parseToGridStrings(data []byte) [][]string {
	s := strings.TrimSpace(string(data))
	if s == "" {
		return [][]string{}
	}
	lines := strings.Split(s, "\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			grid[i] = []string{}
			continue
		}
		if strings.Contains(line, " ") {
			grid[i] = strings.Fields(line)
			continue
		}
		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			for j := range parts {
				parts[j] = strings.TrimSpace(parts[j])
			}
			grid[i] = parts
			continue
		}
		// otherwise split into single-character strings
		row := make([]string, 0, len(line))
		for _, r := range line {
			row = append(row, string(r))
		}
		grid[i] = row
	}
	return grid
}

// parseToIntGrid converts raw file bytes into a [][]int.
// It supports:
// - space-separated or comma-separated integers on a line
// - digit grids where each character is a single digit (0-9)
// Returns an error if any token can't be parsed as an int.
func parseToIntGrid(data []byte) ([][]int, error) {
	s := strings.TrimSpace(string(data))
	if s == "" {
		return [][]int{}, nil
	}
	lines := strings.Split(s, "\n")
	grid := make([][]int, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			grid[i] = []int{}
			continue
		}
		// support comma or space separated numbers
		if strings.Contains(line, " ") || strings.Contains(line, ",") {
			cleaned := strings.ReplaceAll(line, ",", " ")
			parts := strings.Fields(cleaned)
			row := make([]int, len(parts))
			for j, p := range parts {
				v, err := strconv.Atoi(p)
				if err != nil {
					return nil, fmt.Errorf("parse error line %d col %d: %w", i, j, err)
				}
				row[j] = v
			}
			grid[i] = row
			continue
		}
		// otherwise assume each rune is a digit
		row := make([]int, 0, len(line))
		for j, r := range line {
			if r < '0' || r > '9' {
				return nil, fmt.Errorf("non-digit rune %q at line %d col %d", r, i, j)
			}
			row = append(row, int(r-'0'))
		}
		grid[i] = row
	}
	return grid, nil
}

func calcBatteries(num int64) {
	numStr := strconv.FormatInt(num, 2)
	fmt.Println(numStr)
}

func solveProblem(data []string) {
	var intArr = make([]int64, len(data))
	for i := 0; i < len(data); i++ {
		number, err := strconv.ParseInt(data[i], 0, 64)
		if err != nil {
			panic(err)
		}
		intArr[i] = number
	}
	for _, num := range intArr {
		calcBatteries(num)
	}
}

func main() {
	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	splitData := formatInputData(data)
	solveProblem(splitData)
}