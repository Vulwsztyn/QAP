package main

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultSize = 10

type IntMat [defaultSize][defaultSize]int

func matrix2String(t IntMat) string {
	s := ""
	for i := range t {
		for _, n := range t[i] {
			s += fmt.Sprintf("%d ", n)
		}
		s += fmt.Sprintln()
	}
	return s
}

func genMatrix(size int, maxRange int) (matrix IntMat) {
	for i := 0; i < size; i++ {
		var row [defaultSize]int
		for j := 0; j < size; j++ {
			row[j] = rand.Intn(maxRange)
		}
		matrix[i] = row
	}
	return
}

func makeRange(min, max int) []int {
	a := make([]int, max-min)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func randomPermutation(size int) []int {
	_range := makeRange(0, size)
	result := make([]int, size)
	for i := 0; i < size; i++ {
		j := rand.Intn(size - i)
		result[i] = _range[j]
		_range[j] = _range[len(_range)-1]
		_range = _range[:len(_range)-1]
	}
	return result
}

func main() {
	var timeSplits []int64
	maxRange := 100
	start := time.Now()

	m1 := genMatrix(defaultSize, maxRange)
	m2 := genMatrix(defaultSize, maxRange)

	stop := time.Since(start)
	timeSplits = append(timeSplits, stop.Microseconds())

	fmt.Println(matrix2String(m1))
	fmt.Println(matrix2String(m2))

	fmt.Println(timeSplits)

	fmt.Println(randomPermutation(8))
	fmt.Println(randomPermutation(8))
	fmt.Println(randomPermutation(8))
}
