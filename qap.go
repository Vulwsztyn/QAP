package main

import (
	"fmt"
	"math/rand"
)

const defaultSize = 12
const neighbourCount = defaultSize * (defaultSize - 1) / 2

func steepest(m1 IntMat, m2 IntMat) (Assignment, int) {
	currentAssignment := randomPermutation()
	bestCost, _ := calcCost(currentAssignment, m1, m2)
	bestNeighbourCost := bestCost - 1
	bestNeighbourIndex := 0
	for bestNeighbourCost < bestCost {
		bestCost, _ = calcCost(currentAssignment, m1, m2)
		neighbours, neighboursCosts := createNeighbours(currentAssignment, m1, m2)
		bestNeighbourCost, bestNeighbourIndex = min(neighboursCosts[:])
		currentAssignment = neighbours[bestNeighbourIndex]
	}
	fmt.Println(bestCost, currentAssignment)
	return currentAssignment, bestCost
}

func createNeighbours(assignment Assignment, m1 IntMat, m2 IntMat) (result [neighbourCount]Assignment, costs [neighbourCount]int) {
	index := 0
	for i := 0; i < defaultSize-1; i++ {
		for j := i + 1; j < defaultSize; j++ {
			tmp := assignment
			tmp[i], tmp[j] = tmp[j], tmp[i]
			costs[index], _ = calcCost(tmp, m1, m2)
			result[index] = tmp
			index++
		}
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

func greedy(m1, m2 IntMat) (Assignment, Assignment) {
	elemsMax := m1.elemsSorted()
	elemsMin := m2.elemsSorted()
	//fmt.Println(elemsMax)
	//fmt.Println(elemsMin)
	indexMax := defaultSize * defaultSize
	indexMin := -1
	indexMinDiagonal := -1
	assignment := NewAssignment(-1)
	assignment2 := NewAssignment(-1)
	alreadyAssigned := notEquals(-1)
	alreadyAssignedAndNotEqual := func(a int) func(int) bool { return and(alreadyAssigned, func(x int) bool { return x != a }) }
	for assignment.count(equals(-1)) > 1 {
		// Okazało się, że te 2 linijki są konieczne
		// pewnie da sie teraz wywalać zużyte czy coś, ale najpierw ma działać, a potem działać dobrze
		indexMin = -1
		indexMinDiagonal = -1
		var iMax, jMax int
		for toBeContinued := true; toBeContinued; toBeContinued = alreadyAssigned(assignment2[iMax]) && alreadyAssigned(assignment2[jMax]) {
			//checks if a value's pair isn't already defined
			indexMax--
			iMax, jMax = elemsMax[indexMax][1], elemsMax[indexMax][2]
			//fmt.Println(iMax, jMax)
			//fmt.Println(assignment2[iMax])
			//fmt.Println(assignment2[jMax])
		}

		var iMin, jMin int
		if iMax == jMax {
			//diagonal matches only with diagonal
			for toBeContinued := true; toBeContinued; {
				indexMinDiagonal++
				iMin, jMin = elemsMin[indexMinDiagonal][1], elemsMin[indexMinDiagonal][2]
				toBeContinued =
					iMin != jMin ||
						alreadyAssignedAndNotEqual(iMin)(assignment2[iMax]) ||
						alreadyAssignedAndNotEqual(iMax)(assignment[iMin])
			}
		} else {
			for toBeContinued := true; toBeContinued; {
				indexMin++
				iMin, jMin = elemsMin[indexMin][1], elemsMin[indexMin][2]
				toBeContinued =
					iMin == jMin ||
						alreadyAssignedAndNotEqual(iMin)(assignment2[iMax]) ||
						alreadyAssignedAndNotEqual(jMin)(assignment2[jMax]) ||
						alreadyAssignedAndNotEqual(iMax)(assignment[iMin]) ||
						alreadyAssignedAndNotEqual(jMax)(assignment[jMin])
			}
		}

		assignment[iMin] = iMax
		assignment[jMin] = jMax
		assignment2[iMax] = iMin
		assignment2[jMax] = jMin
		//fmt.Println(iMax, jMax)
		//fmt.Println(iMin, jMin)
		//fmt.Println(assignment)
		//fmt.Println(assignment2)
	}
	if index := assignment.findIndex(equals(-1)); index > -1 {
		value := assignment2.findIndex(equals(-1))
		assignment[index] = value
		assignment2[value] = index
	}
	return assignment, assignment2
}

func main() {
	m1, m2 := fileReader("instances/nug12.dat")
	fmt.Println(m1)
	fmt.Println(m2)
	for i := 0; i < 100; i++ {
		steepest(m1, m2)
	}
	cost, _ := calcCost(Assignment{11, 6, 8, 2, 3, 7, 10, 0, 4, 5, 9, 1}.translate(), m1, m2)
	fmt.Println(cost)
}
