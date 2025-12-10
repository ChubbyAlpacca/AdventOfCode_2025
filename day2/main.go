package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFilePath = "data/full_input.txt"


func formatInputData(data []byte) []string {
	s := strings.TrimSpace(string(data))
	if s == "" {
		return []string{}
	}
	return strings.Split(s, ",")
}


func main() {
	data, err := os.ReadFile(inputFilePath)
	if err!= nil {
		panic(err)
	}
	fmt.Print("Input data recieved => \n" + string(data) + "\n\n\n\n")
	formattedData := formatInputData(data)
	var totalRepeatedSum int64 = 0
	totalRepeatedCount := 0

	for i := 0; i < len(formattedData); i++ {
		entry := strings.TrimSpace(formattedData[i])
		if entry == "" {
			continue
		}
		fmt.Printf("Entry %d: %s\n", i, entry)

		if strings.Contains(entry, "-") {
			parts := strings.SplitN(entry, "-", 2)
			if len(parts) != 2 {
				fmt.Println("  invalid range format:", entry)
				continue
			}
			start64, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
			end64, err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
			if err1 != nil || err2 != nil {
				fmt.Println("  invalid range numbers:", entry)
				continue
			}
			if start64 > end64 {
				start64, end64 = end64, start64
			}

			entrySum := int64(0)
			entryCount := 0
			for n := start64; n <= end64; n++ {
				s := strconv.FormatInt(n, 10)
				if isRepeatedSequence(s) {
					entrySum += n
					entryCount++
					fmt.Printf("    invalid ID: %s\n", s)
				}
			}
			fmt.Printf("  found %d invalid IDs, sum=%d\n", entryCount, entrySum)
			totalRepeatedCount += entryCount
			totalRepeatedSum += entrySum
		} else {
			s := entry
			if isRepeatedSequence(s) {
				v, err := strconv.ParseInt(s, 10, 64)
				if err == nil {
					fmt.Printf("  invalid ID: %s\n", s)
					totalRepeatedCount++
					totalRepeatedSum += v
				}
			} else {
				fmt.Println("  no invalid pattern")
			}
		}
	}

	fmt.Printf("\nTotals across all entries:\n")
	fmt.Printf("  total repeated-half IDs: %d\n", totalRepeatedCount)
	fmt.Printf("  total sum of repeated-half numeric IDs: %d\n", totalRepeatedSum)
}


func isRepeatedSequence(s string) bool {
	L := len(s)
	if L < 2 {
		return false
	}
	for p := 1; p <= L/2; p++ {
		if L%p != 0 {
			continue
		}
		times := L / p
		if times < 2 {
			continue
		}
		pattern := s[:p]
		ok := true
		for i := 1; i < times; i++ {
			if s[i*p:(i+1)*p] != pattern {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}