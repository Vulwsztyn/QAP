package main

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultSize = 10

type IntMat [defaultSize][defaultSize]int

func (t IntMat) String() string {
	s := ""
	for _, row := range t {
		for _, n := range row {
			s += fmt.Sprintf("%d ", n)
		}
		s += fmt.Sprintln()
	}
	return s
}

func genMatrix(maxRange int) (matrix IntMat) {
	for i := 0; i < defaultSize; i++ {
		for j := 0; j < defaultSize; j++ {
			matrix[i][j] = rand.Intn(maxRange)
		}
	}
	return
}

func makeRange(min, max int) []int {
	_range := make([]int, max-min)
	for i := range _range {
		_range[i] = min + i
	}
	return _range
}

func randomPermutation(size int) []int {
	_range := makeRange(0, size)
	result := make([]int, size)
	for i := 0; i < size; i++ {
		j := rand.Intn(size - i)
		result[i] = _range[j]
		_range[j] = _range[len(_range)-1]
	}
	return result
}

func main() {
	var timeSplits []int64
	maxRange := 100
	start := time.Now()

	m1 := genMatrix(maxRange)
	m2 := genMatrix(maxRange)

	stop := time.Since(start)
	timeSplits = append(timeSplits, stop.Microseconds())

	fmt.Println(m1)
	fmt.Println(m2)

	fmt.Println(timeSplits)

	fmt.Println(randomPermutation(8))
	fmt.Println(randomPermutation(8))
	fmt.Println(randomPermutation(8))
}
