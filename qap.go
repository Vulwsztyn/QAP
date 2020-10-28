package main

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultSize = 5
const neighbourCount = defaultSize * (defaultSize - 1) / 2

func steepest(m1 IntMat, m2 IntMat) (Assignment, int) {
	currentAssignment := randomPermutation()
	var bestCost, bestNeighbourCost, bestNeighbourIndex int
	for ok := true; ok; ok = bestNeighbourCost < bestCost {
		bestCost, _ = calcCost(currentAssignment, m1, m2)
		neighbours, neighboursCosts := createNeighbours(currentAssignment, m1, m2, rand.Intn(defaultSize))
		bestNeighbourCost, bestNeighbourIndex = min(neighboursCosts[:])
		currentAssignment = neighbours[bestNeighbourIndex]
	}
	fmt.Println(bestCost, currentAssignment)
	return currentAssignment, bestCost
}

func positiveReminder(a, b int) (result int) {
	result = a % b
	if result < 0 {
		result += b
	}
	return
}

func createNeighbours(assignment Assignment, m1 IntMat, m2 IntMat, startIndex int) (result [neighbourCount]Assignment, costs [neighbourCount]int) {
	index := 0
	iCount := 0
	for i := startIndex; index < neighbourCount; i = (i + 1) % defaultSize {
		fmt.Println(i)
		for j := (i + 1) % defaultSize; j != positiveReminder(i-iCount, defaultSize); j = (j + 1) % defaultSize {
			tmp := assignment
			tmp[i], tmp[j] = tmp[j], tmp[i]
			fmt.Println(i, j)
			costs[index], _ = calcCost(tmp, m1, m2)
			result[index] = tmp
			index++
		}
		iCount++
	}
	return
}

func calcCost(assignment Assignment, m1 IntMat, m2 IntMat) (result int, costMatrix IntMat) {
	for i := 0; i < defaultSize; i++ {
		for j := 0; j < defaultSize; j++ {
			costMatrix[i][j] = m1[assignment[i]][assignment[j]] * m2[i][j]
			result += costMatrix[i][j]
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

func equals(value int) func(int) bool {
	// returns a function checking whether its argument is equal to the argument of this function
	return func(a int) bool { return a == value }
}

func notEquals(value int) func(int) bool {
	// jestem pewny, że da się to zrobić lepiej
	// returns a function checking whether its argument is not equal to the argument of this function
	return func(a int) bool { return a != value }
}

func and(fn1, fn2 func(int) bool) func(int) bool {
	// returns a function checking whether its argument returns true for both parameter functions
	return func(a int) bool { return fn1(a) && fn2(a) }
}

func greedy(assignment Assignment, m1, m2 IntMat) (result Assignment, cost int) {
	bestAssignment := assignment
	var exists bool
	for ok := true; ok; {
		bestAssignment, cost, exists = bestAssignment.getFirstBetterNeighbour(m1, m2)
		ok = exists
	}
	return bestAssignment, cost
}

func main() {
	rand.Seed(123)
	var timeSplits []int64
	maxRange := 1000
	minRange := 100
	start := time.Now()
	m1 := NewRandomMatrix(maxRange, minRange)
	m2 := NewRandomMatrix(maxRange, minRange)

	stop := time.Since(start)
	timeSplits = append(timeSplits, stop.Microseconds())

	fmt.Println(timeSplits)

	testAssignment := randomPermutation()
	//testAssignment := Assignment{1, 4, 2, 3, 0}
	//fmt.Println(testAssignment)
	//fmt.Println(testAssignment.translateAssignment())
	//
	fmt.Println(m1)

	fmt.Println(m2)
	fmt.Println(m1.permuteMatrix(testAssignment))
	//fmt.Println(testAssignment)
	//fmt.Println(calcCost(testAssignment, m1, m2))
	//fmt.Println(fileReader("instances/chr12a.dat"))

	//fmt.Println(testAssignment)
	//for _,v := range createNeighbours(testAssignment) {
	//	fmt.Println(v)
	//}
	//fmt.Println(testAssignment.any(func(a int) bool {return a == 4 }))
	fmt.Println(testAssignment)
	fmt.Println(calcCost(testAssignment, m1, m2))
	fmt.Println(greedy(testAssignment, m1, m2))
}
