package main

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultSize = 5
const neighbourCount = defaultSize * (defaultSize - 1) / 2

func calcCost(assignment Assignment, m1 IntMat, m2 IntMat) (result int) {
	for i := 0; i < defaultSize; i++ {
		for j := 0; j < defaultSize; j++ {
			result += m1[assignment[i]][assignment[j]] * m2[i][j]
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

func randomPermutation() Assignment {
	_range := makeRange(0, defaultSize)
	var result [defaultSize]int
	for i := 0; i < defaultSize; i++ {
		j := rand.Intn(defaultSize - i)
		result[i] = _range[j]
		_range[j] = _range[len(_range)-1-i]
	}
	return result
}

func createNeighbours(assignment Assignment) (result [neighbourCount]Assignment) {
	index := 0
	for i := 0; i < defaultSize-1; i++ {
		for j := i + 1; j < defaultSize; j++ {
			tmp := assignment
			tmp[i], tmp[j] = tmp[j], tmp[i]
			result[index] = tmp
			index++
		}
	}
	return
}

func main() {
	var timeSplits []int64
	maxRange := 5
	start := time.Now()

	m1 := NewRandomMatrix(maxRange)
	m2 := NewRandomMatrix(maxRange)

	stop := time.Since(start)
	timeSplits = append(timeSplits, stop.Microseconds())

	fmt.Println(timeSplits)

	testAssignment := randomPermutation()
	testAssignment = Assignment{1,0,2,3,4}
	fmt.Println(testAssignment)
	fmt.Println(testAssignment.translateAssignment())

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(testAssignment)
	fmt.Println(calcCost(testAssignment, m1, m2))
	//fmt.Println(fileReader("instances/chr12a.dat"))

	fmt.Println(testAssignment)
	for _,v := range createNeighbours(testAssignment) {
		fmt.Println(v)
	}
	fmt.Println(testAssignment.any(func(a int) bool {return a == 4 }))
}
