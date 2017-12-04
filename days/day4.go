package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Day4() {
	scanner := getScanner("inputs/input4.csv")
	var validCountStar7, validCountStar8 int
	for {
		if scanner.Scan() {
			line := scanner.Text()
			words := strings.Split(line, " ")
			validCountStar7++
			validCountStar8++

			if !validLine(stringsEqual, words) {
				validCountStar7--
			}
			if !validLine(charsetsEqual, words) {
				validCountStar8--
			}
		} else {
			break
		}
	}
	fmt.Printf("Valid passwords by equal strings: %d\n", validCountStar7)
	fmt.Printf("Valid passwords by equal character sets: %d", validCountStar8)
}

func validLine(comparePairs func(string, string) bool, words []string) bool {
	lineValid := true
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j {
				if comparePairs(words[i], words[j]) {
					lineValid = false
				}
			}
		}
	}
	return lineValid
}

func charsetsEqual(a, b string) bool {
	chars := make(map[rune]int)
	for _, character := range a {
		chars[character]++
	}
	for _, character := range b {
		_, exists := chars[character]
		if !exists {
			return false
		}
		chars[character]--
	}
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}
	return true
}

func stringsEqual(a, b string) bool {
	return strings.Compare(a, b) == 0
}

func getScanner(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Couldn't open file %s, error was: %v", filename, err)
		os.Exit(1)
	}
	s := bufio.NewScanner(file)
	return s
}
