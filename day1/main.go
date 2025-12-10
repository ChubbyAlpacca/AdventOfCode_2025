package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// const inputFilePath = "data/test_input.txt"
const inputFilePath = "data/full_input.txt"
var counter int = 0

func createLockDial(min, max int) []int {
	length := max - min + 1
	if length <= 0 {
		return []int{}
	}
	arr := make([]int, length)
	for i := range arr {
		arr[i] = min + i
	}
	return arr
}

func formatInputData(data []byte) []string {
	s := strings.TrimSpace(string(data))
	if s == "" {
		return []string{}
	}
	return strings.Split(s, "\n")
}


func applyRotation(pos int, instruction string) (int, bool, int) {
	const N = 100
	instruction = strings.TrimSpace(instruction)
	if instruction == "" {
		return pos, false, 0
	}
	cmd := strings.ToUpper(string(instruction[0]))
	amt, err := strconv.Atoi(strings.TrimSpace(instruction[1:]))
	if err != nil {
		fmt.Println("invalid instruction (cannot parse amount):", instruction)
		return pos, false, 0
	}

	var newPos int
	var passes int
	switch cmd {
	case "L":
		newPos = ((pos-amt)%N + N) % N
		k0 := pos
		if k0 == 0 {
			k0 = N
		}
		if amt >= k0 {
			passes = 1 + (amt-k0)/N
		} else {
			passes = 0
		}
	case "R":
		newPos = (pos + amt) % N
		k0 := N - pos
		if k0 == 0 {
			k0 = N
		}
		if amt >= k0 {
			passes = 1 + (amt-k0)/N
		} else {
			passes = 0
		}
	default:
		fmt.Println("invalid command:", cmd)
		return pos, false, 0
	}
	return newPos, newPos == 0, passes
}

func solveProblem(data []byte) string {
	formattedData := formatInputData(data)
	pos := 50
	counter = 0
	passCount := 0
	for _, instr := range formattedData {
		newPos, hitZero, passes := applyRotation(pos, instr)
		pos = newPos
		if hitZero {
			counter++
		}
		passCount += passes
	}
	return fmt.Sprintf("final pos=%d", pos)
}

func main() {
	data, err := os.ReadFile(inputFilePath)
	if err!= nil {
		panic(err)
	}
	fmt.Print("Input data recieved => \n" + string(data) + "\n\n\n\n")
	result := solveProblem(data)
	fmt.Printf("%s\n", result)
	formatted := formatInputData(data)
	pos := 50
	totalPasses := 0
	landings := 0
	for _, instr := range formatted {
		newPos, hitZero, passes := applyRotation(pos, instr)
		pos = newPos
		if hitZero {
			landings++
		}
		totalPasses += passes
	}
	fmt.Printf("Total times dial landed on zero: %d\n", landings)
	fmt.Printf("Total times dial passed zero while rotating: %d\n", totalPasses)
}