package app

import (
	"math"
	"sort"
)

// Считаем моду mode
func CalculateMode(arr []int) float64 {

	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	mode := arr[0]
	maxFreq := freq[mode]

	for num, count := range freq {
		if count > maxFreq || (count == maxFreq && num < mode) {
			mode = num
			maxFreq = count
		}
	}
	return float64(mode)
}

// Считаем медиану median
func CalculateMedian(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	median := 0.0
	if n%2 == 0 {
		median = float64(arr[n/2-1]+arr[n/2]) / 2.0
	} else {
		median = float64(arr[n/2])
	}
	return median
}

// Считаем среднее mean
func CalculateMean(arr []int) float64 {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	mean := float64(sum) / float64(len(arr))
	return mean
}

// Считаем стандартное отклонение SD
func CalculateSd(arr []int) float64 {
	mean := CalculateMean(arr)
	var varSum float64
	for _, num := range arr {
		varSum += math.Pow(float64(num)-mean, 2)
	}
	return math.Sqrt(varSum / float64(len(arr)))
}
