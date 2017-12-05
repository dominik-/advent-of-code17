package days

import (
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
)

func Day5() {
	scanner := GetScanner("inputs/input5")
	inputAsArray := make([]int, 0)
	for {
		if !scanner.Scan() {
			break
		}
		intVal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		inputAsArray = append(inputAsArray, intVal)
	}
	inputClone := make([]int, len(inputAsArray))
	copy(inputClone, inputAsArray)

	fmt.Printf("Jumps Star 1: %d\n", jumpAndUpdateUntilOOB(inputAsArray, 0, 0))
	//we will spawn a lot of goroutines with the recursion, so we need this
	debug.SetMaxStack(4000000000)
	fmt.Printf("Jumps Star 2: %d\n", jumpAndUpdateUntilOOBWithTwist(inputClone, 0, 0))
}

func jumpAndUpdateUntilOOB(array []int, pos, counter int) int {
	counter++
	delta := array[pos]
	array[pos]++
	if pos+delta > len(array)-1 {
		return counter
	}
	return jumpAndUpdateUntilOOB(array, pos+delta, counter)
}

func jumpAndUpdateUntilOOBWithTwist(array []int, pos, counter int) int {
	counter++
	delta := array[pos]
	if delta >= 3 {
		array[pos]--
	} else {
		array[pos]++
	}
	if pos+delta > len(array)-1 {
		return counter
	}
	return jumpAndUpdateUntilOOBWithTwist(array, pos+delta, counter)
}
