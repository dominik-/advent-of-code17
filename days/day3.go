package days

import (
	"fmt"
	"sort"
)

var input3 = 361527

func Day3() {
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
	fmt.Printf("Manhattan Distance: %d\n", circles+minDistance(sideCenters, goal))
	fmt.Printf("First Snake Head above %d: %d\n", input3, snake())
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

type Point struct {
	X, Y, Val int
}

type Coordinates struct {
	xAxis map[int][]*Point
}

func buildCoordinates() *Coordinates {
	xCoordMap := make(map[int][]*Point)
	xCoordMap[0] = make([]*Point, 0)
	coordinates := &Coordinates{
		xAxis: xCoordMap,
	}
	coordinates.addPoint(0, 0, 1)
	return coordinates
}

func (c *Coordinates) sumAdjacent(x, y int) int {
	newValue := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if !(i == 0 && j == 0) {
				newValue += c.getPointValue(x+i, y+j)
			}
		}
	}
	return newValue
}

func snake() int {
	coords := buildCoordinates()
	goal := input3
	currentHead := 1
	x := 0
	y := 0
	for i := 1; currentHead < goal; i++ {
		//distance to walk in each direction before turning left
		quadrantDistance := 2 * i
		//walk one step right
		x++
		currentHead = coords.addPointWithAdjacentValue(x, y)
		if currentHead > goal {
			return currentHead
		}
		//walk (quadrant-distance - 1) north
		for n := 0; n < quadrantDistance-1; n++ {
			y++
			currentHead = coords.addPointWithAdjacentValue(x, y)
			if currentHead > goal {
				return currentHead
			}
		}
		//walk 2x quadrant-distance west
		for w := 0; w < quadrantDistance; w++ {
			x--
			currentHead = coords.addPointWithAdjacentValue(x, y)
			if currentHead > goal {
				return currentHead
			}
		}
		//walk 2x quadrant-distance south
		for s := 0; s < quadrantDistance; s++ {
			y--
			currentHead = coords.addPointWithAdjacentValue(x, y)
			if currentHead > goal {
				return currentHead
			}
		}
		//walk 2x quadrant-distance east
		for e := 0; e < quadrantDistance; e++ {
			x++
			currentHead = coords.addPointWithAdjacentValue(x, y)
			if currentHead > goal {
				return currentHead
			}
		}
		//back to the start!
	}
	return currentHead
}

func (co *Coordinates) addPointWithAdjacentValue(x, y int) int {
	value := co.sumAdjacent(x, y)
	co.addPoint(x, y, value)
	return value
}

func (co *Coordinates) addPoint(x, y, val int) {
	p := &Point{
		X:   x,
		Y:   y,
		Val: val,
	}
	co.xAxis[x] = append(co.xAxis[x], p)
}

func (co *Coordinates) getPointValue(x, y int) int {
	for _, point := range co.xAxis[x] {
		if point.Y == y {
			return point.Val
		}
	}
	return 0
}
