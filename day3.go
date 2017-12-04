package main

import (
	"fmt"
	"sort"
)

var input3 = 361527

func main() {
	goal := input3
	currentLimit := 0
	circles := 1
	sideCenters := make([]int, 4)
	for i := 1; currentLimit < goal; i++ {
		quarterDistance := 2 * i
		base := quarterDistance + 1
		//bottom right corner of last completed circle
		currentLimit = base * base
		circles = i
		updateCenters(sideCenters, currentLimit, quarterDistance)
	}
	fmt.Printf("Centers: %d; Circles: %d\n", sideCenters, circles)
	fmt.Printf("Min distance to center: %d\n", minDistance(sideCenters, goal))
	fmt.Printf("Manhattan Distance: %d", circles+minDistance(sideCenters, goal))
}

func updateCenters(centers []int, farthestCorner, quarterDistance int) {
	centers[3] = farthestCorner - quarterDistance/2
	for i := 2; i > -1; i-- {
		centers[i] = centers[i+1] - quarterDistance
	}
}

func minDistance(array []int, value int) int {
	distArray := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		//dont wanna use math Abs with float conversion here
		if array[i] > value {
			distArray[i] = array[i] - value
		} else {
			distArray[i] = value - array[i]
		}
	}
	sort.Ints(distArray)
	return distArray[0]
}

func lastIntOfModuloArray(goal int) int {
	start := []int{1, 1, 2, 4, 5, 10, 11, 23, 25}
	lastNum := 25
	ring := 2
	side := 0
	for i := 10; lastNum < goal; i++ {
		//now i counts the actual array index, so we need to defer the current "ring" to determine the position relative to corners
		//for each ring (beyond the zero ring with only 1 number), 8 more numbers are added
		if i%8 == 0 {
			ring++
			side := 0
		}
		//for each side 2 more numbers are added
		currentSideLength := ring + 2
		if i%currentSideLength == 0 {
			side++
			side = side % 4
			//case: first number of a side
			if side == 0 {
				start = append(start, start[i-1])
			} else {
				//wtf...
				sum := start[i-1] + start[i-2] + start[i-currentSideLength-ring*8+1]
			}
		}
		//case: middle number of side
		//case: last number of a side
	}
}
