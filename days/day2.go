package days

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

var inputFilename = "inputs/input2.csv"

func Day2() {
	file, err := os.Open(inputFilename)
	if err != nil {
		log.Printf("Couldn't open file %s, error was: %v", inputFilename, err)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	r.Comma = rune("\t"[0])

	lineChecksumsStar1 := make([]int, 0)
	lineChecksumsStar2 := make([]int, 0)

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lineAsInt := make([]int, len(row))
		for i, val := range row {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				fmt.Printf("Error parsing value %v", val)
			} else {
				lineAsInt[i] = intVal
			}
		}
		lineChecksumsStar1 = append(lineChecksumsStar1, minMaxDiff(lineAsInt))
		lineChecksumsStar2 = append(lineChecksumsStar2, evenlyDivisible(lineAsInt))
	}
	var totalChecksumStar1, totalChecksumStar2 int
	for _, lc := range lineChecksumsStar1 {
		totalChecksumStar1 += lc
	}

	for _, lc := range lineChecksumsStar2 {
		totalChecksumStar2 += lc
	}

	fmt.Printf("Checksum Star 1: %d\n", totalChecksumStar1)
	fmt.Printf("Checksum Star 2: %d\n", totalChecksumStar2)

}

func minMaxDiff(array []int) int {
	sort.Ints(array)
	return array[len(array)-1] - array[0]
}

func evenlyDivisible(array []int) int {
	sort.Ints(array)
	for i := len(array) - 1; i > -1; i-- {
		for j := 0; j < i; j++ {
			if array[i]%array[j] == 0 {
				//fmt.Printf("Found divisible pair: %d, %d\n", array[i], array[j])
				return array[i] / array[j]
			}
		}
	}
	//we should never get here
	return 0
}
