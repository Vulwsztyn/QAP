package main

import "time"

func measureTime(fn func()) int64 {
	start := time.Now()
	fn()
	stop := time.Since(start)
	return stop.Microseconds()
}

func min(a []int) (min int, minIndex int) {
	min = a[0]
	minIndex = 0
	for index, value := range a {
		if value < min {
			min = value
			minIndex = index
		}
	}
	return min, minIndex
}
