package days

import (
	"fmt"
)

func Day6() {
	buckets := []int{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11}

	variationCount := make(map[string]int)
	iterations := 0
	var currentVariation string
	for {
		currentVariation = fmt.Sprint(buckets)
		_, exists := variationCount[currentVariation]
		if exists {
			break
		}
		variationCount[currentVariation] = iterations
		redistribute, index := getMaxBucketAndSetToZero(buckets)
		levelOutBuckets(buckets, redistribute, index)
		iterations++
	}

	fmt.Printf("Final buckets: %v\n", buckets)
	fmt.Printf("Took %d iterations to reach repeating buckets sizes.", iterations)
	fmt.Printf("Took %d iterations after first allocation to identify repeating sizes.", iterations-variationCount[currentVariation])
}

func getMaxBucketAndSetToZero(buckets []int) (int, int) {
	max := buckets[0]
	maxIndex := 0
	for i := 1; i < len(buckets); i++ {
		if buckets[i] > max {
			max = buckets[i]
			maxIndex = i
		}
	}
	buckets[maxIndex] = 0
	return max, maxIndex
}

func levelOutBuckets(buckets []int, dist int, startIndex int) {
	length := len(buckets)
	for i := 0; i < dist; i++ {
		buckets[(i+startIndex+1)%length]++
	}
}
