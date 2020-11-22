package main

import (
	"fmt"
	"math/rand"
	"time"
)

//var instances = []string{"tai256c", "tho150", "wil50", "sko100c", "lipa80a", "nug30", "rou20", "kra32", "chr12c", "bur26e"}
var instances = []string{"bur26d", "kra30a", "lipa40a", "wil50", "lipa60a","lipa70a","sko81", "sko90","sko100a","esc128"}

const defaultSize = int(
		(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((0-1)*(0-2)*(0-3)*(0-4)*(0-5)*(0-6)*(0-7)*(0-8)*(0-9))*26 +
		(instanceIndex-0)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((1-0)*(1-2)*(1-3)*(1-4)*(1-5)*(1-6)*(1-7)*(1-8)*(1-9))*30 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((2-0)*(2-1)*(2-3)*(2-4)*(2-5)*(2-6)*(2-7)*(2-8)*(2-9))*40 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((3-0)*(3-1)*(3-2)*(3-4)*(3-5)*(3-6)*(3-7)*(3-8)*(3-9))*50 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((4-0)*(4-1)*(4-2)*(4-3)*(4-5)*(4-6)*(4-7)*(4-8)*(4-9))*60 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((5-0)*(5-1)*(5-2)*(5-3)*(5-4)*(5-6)*(5-7)*(5-8)*(5-9))*70 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((6-0)*(6-1)*(6-2)*(6-3)*(6-4)*(6-5)*(6-7)*(6-8)*(6-9))*81 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-8)*(instanceIndex-9)/((7-0)*(7-1)*(7-2)*(7-3)*(7-4)*(7-5)*(7-6)*(7-8)*(7-9))*90 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-9)/((8-0)*(8-1)*(8-2)*(8-3)*(8-4)*(8-5)*(8-6)*(8-7)*(8-9))*100 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)/((9-0)*(9-1)*(9-2)*(9-3)*(9-4)*(9-5)*(9-6)*(9-7)*(9-8))*128)
//var instanceSizes = []int{256,150,50,100,80,30,20,32,12,26}

const neighbourCount = defaultSize * (defaultSize - 1) / 2

func steepest(assignment Assignment, m1 IntMat, m2 IntMat) (Assignment, int, int, int, int64) {
	start := time.Now()
	currentAssignment := assignment
	var bestCost, bestNeighbourCost, bestNeighbourIndex, stepCount, exploredSolutions int
	var costMatrix IntMat
	for ok := true; ok; ok = bestNeighbourCost < bestCost {
		bestCost, costMatrix = calcCost(currentAssignment, m1, m2)
		neighbours, neighboursCosts := createNeighbours(currentAssignment, m1, m2, costMatrix, bestCost, rand.Intn(defaultSize))
		exploredSolutions += neighbourCount
		bestNeighbourCost, bestNeighbourIndex = min(neighboursCosts[:])
		currentAssignment = neighbours[bestNeighbourIndex]
		stepCount++
	}
	stop := time.Since(start)
	return currentAssignment, bestCost, stepCount, exploredSolutions, stop.Microseconds()
}

func positiveReminder(a, b int) (result int) {
	result = a % b
	if result < 0 {
		result += b
	}
	return
}

func createNeighbours(assignment Assignment, m1 IntMat, m2 IntMat, previousCostMatrix IntMat, previousCost int, startIndex int) (result [neighbourCount]Assignment, costs [neighbourCount]int) {
	index := 0
	iCount := 0
	for i := startIndex; index < neighbourCount; i = (i + 1) % defaultSize {
		//fmt.Println(i)
		for j := (i + 1) % defaultSize; j != positiveReminder(i-iCount, defaultSize); j = (j + 1) % defaultSize {
			tmp := assignment
			tmp[i], tmp[j] = tmp[j], tmp[i]
			//fmt.Println(i, j)
			costs[index], _ = reCalcCost(tmp, m1, m2, previousCostMatrix, previousCost, [2]int{i, j})
			result[index] = tmp
			index++
		}
		iCount++
	}
	return
}

func minNeighbour(assignment Assignment, m1 IntMat, m2 IntMat, previousCostMatrix IntMat, previousCost int, startIndex int) (result Assignment, cost int, costMatrix IntMat) {
	index := 0
	iCount := 0
	firstIteration := true
	for i := startIndex; index < neighbourCount; i = (i + 1) % defaultSize {
		for j := (i + 1) % defaultSize; j != positiveReminder(i-iCount, defaultSize); j = (j + 1) % defaultSize {
			tmp := assignment
			tmp[i], tmp[j] = tmp[j], tmp[i]
			costTmp, matrixTmp := reCalcCost(tmp, m1, m2, previousCostMatrix, previousCost, [2]int{i, j})
			if costTmp < cost || firstIteration {
				result, cost, costMatrix = tmp, costTmp, matrixTmp
				firstIteration = false
			}
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

func reCalcCost(assignment Assignment, m1 IntMat, m2 IntMat, previousCostMatrix IntMat, previousCost int, indexes [2]int) (int, IntMat) {
	result := previousCost
	costMatrix := previousCostMatrix
	for _, j := range indexes {
		for i := 0; i < defaultSize; i++ {
			if i != j && !(j == indexes[1] && i == indexes[0]) {
				result -= previousCostMatrix[i][j]
				result -= previousCostMatrix[j][i]

				costMatrix[i][j] = m1[assignment[i]][assignment[j]] * m2[i][j]
				costMatrix[j][i] = m1[assignment[j]][assignment[i]] * m2[j][i]

				result += costMatrix[i][j]
				result += costMatrix[j][i]
			}
		}
		result -= previousCostMatrix[j][j]

		costMatrix[j][j] = m1[assignment[j]][assignment[j]] * m2[j][j]

		result += costMatrix[j][j]
	}
	return result, costMatrix
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

func greedy(assignment Assignment, m1, m2 IntMat) (Assignment, int, int, int, int64) {
	start := time.Now()
	bestAssignment := assignment
	var exists bool
	var cost, stepCount, solutionsExplored int
	for ok := true; ok; {
		bestAssignment, cost, solutionsExplored, exists = bestAssignment.getFirstBetterNeighbour(m1, m2)
		ok = exists
		stepCount++
	}
	stop := time.Since(start)
	return bestAssignment, cost, stepCount, solutionsExplored, stop.Microseconds()
}

func heuristic(assignment Assignment, m1, m2 IntMat) (Assignment, int, int, int, int64) {
	start := time.Now()
	currentAssignment := assignment
	var bestCost, bestNeighbourCost, bestNeighbourIndex, stepCount, exploredSolutions int
	var costMatrix IntMat

	for ok := true; ok; ok = bestNeighbourCost < bestCost {
		bestCost, costMatrix = calcCost(currentAssignment, m1, m2)
		var temp [defaultSize]int
		for i := 0; i < defaultSize; i++ {
			for j := 0; j < defaultSize; j++ {
				temp[i] += costMatrix[i][j] + costMatrix[j][i]
			}
		}

		_, maxi := max(temp[:])
		neighbours, neighboursCosts := createNeighboursHeuristic(currentAssignment, m1, m2, costMatrix, bestCost, maxi)
		exploredSolutions += len(neighbours)
		bestNeighbourCost, bestNeighbourIndex = min(neighboursCosts[:])
		currentAssignment = neighbours[bestNeighbourIndex]
		stepCount++
	}
	stop := time.Since(start)
	return currentAssignment, bestCost, stepCount, exploredSolutions, stop.Microseconds()
}

func createNeighboursHeuristic(assignment Assignment, m1 IntMat, m2 IntMat, previousCostMatrix IntMat, previousCost int, i int) (result [defaultSize - 1]Assignment, costs [defaultSize - 1]int) {
	index := 0
	iCount := 0
	for j := (i + 1) % defaultSize; j != positiveReminder(i-iCount, defaultSize); j = (j + 1) % defaultSize {
		tmp := assignment
		tmp[i], tmp[j] = tmp[j], tmp[i]
		costs[index], _ = reCalcCost(tmp, m1, m2, previousCostMatrix, previousCost, [2]int{i, j})
		result[index] = tmp
		index++
	}
	iCount++
	return
}

func random(timeLimit int64, m1, m2 IntMat) (Assignment, int, int, int64) {
	start := time.Now()
	var bestCost, stepCount int
	var bestAssignment Assignment
	stop := time.Since(start)
	for ok := true; ok; {
		assignment := randomPermutation()
		cost, _ := calcCost(assignment, m1, m2)
		if stepCount == 0 || bestCost > cost {
			bestCost = cost
			bestAssignment = assignment
		}
		stop = time.Since(start)
		ok = stop.Microseconds() < timeLimit
		stepCount++
	}
	return bestAssignment, bestCost, stepCount, stop.Microseconds()
}

func randomWalk(assignment Assignment, timeLimit int64, m1, m2 IntMat) (Assignment, int, int, int64) {
	start := time.Now()
	currentAssignment := assignment
	var bestCost, stepCount int
	var bestAssignment Assignment
	stop := time.Since(start)
	for ok := true; ok; {
		i := rand.Intn(defaultSize)
		j := rand.Intn(defaultSize - 1)
		if j >= i {
			j++
		}
		currentAssignment[i], currentAssignment[j] = currentAssignment[j], currentAssignment[i]
		cost, _ := calcCost(assignment, m1, m2)
		if stepCount == 0 || bestCost > cost {
			bestCost = cost
			bestAssignment = assignment
		}
		//fmt.Println(assignment, currentAssignment, bestCost)
		stop = time.Since(start)
		ok = stop.Microseconds() < timeLimit
		stepCount++

	}
	return bestAssignment, bestCost, stepCount, stop.Microseconds()
}

func distance(assignment Assignment, assignment2 Assignment) (distance float64) {
	for i, v := range assignment {
		if assignment2[i] != v {
			distance++
		}
	}
	distance /= defaultSize
	return
}

func measureTime(filename string, times int) {
	m1, m2, _ := fileReader("instances/" + filename + ".dat")
	var SArray, GArray, HArray, RWArray, RArray [][5]int
	for i := 0; i < times; i++ {
		fmt.Println(i, "iteration...")
		assignment := randomPermutation()
		assignmentCost, _ := calcCost(assignment, m1, m2)
		fmt.Println("Steepest")
		_, costS, stepsS, exploreSolutionsS, timeS := steepest(assignment, m1, m2)
		fmt.Println("Greedy")
		_, costG, stepsG, exploreSolutionsG, timeG := greedy(assignment, m1, m2)
		fmt.Println("Heuristic")
		_, costH, stepsH, exploreSolutionsH, timeH := heuristic(assignment, m1, m2)
		timeLimit := (timeS + timeG) / 2
		fmt.Println("Random Walk")
		_, costRW, exploreSolutionsRW, timeRW := randomWalk(assignment, timeLimit, m1, m2)
		fmt.Println("Random")
		_, costR, exploreSolutionsR, timeR := random(timeLimit, m1, m2)

		GArray = append(GArray, [5]int{costG, stepsG, exploreSolutionsG, int(timeG), assignmentCost})
		SArray = append(SArray, [5]int{costS, stepsS, exploreSolutionsS, int(timeS), assignmentCost})
		HArray = append(HArray, [5]int{costH, stepsH, exploreSolutionsH, int(timeH), assignmentCost})
		RWArray = append(RWArray, [5]int{costRW, -1, exploreSolutionsRW, int(timeRW), assignmentCost})
		RArray = append(RArray, [5]int{costR, -1, exploreSolutionsR, int(timeR), assignmentCost})
	}
	writeFile(SArray, "S_"+filename)
	writeFile(GArray, "G_"+filename)
	writeFile(HArray, "H_"+filename)
	writeFile(RWArray, "RW_"+filename)
	writeFile(RArray, "R_"+filename)
	fmt.Println("done")
}

func main() {
	rand.Seed(123)
	//filename := "instances/chr12a.dat"
	//m1, m2, _ := fileReader(filename)
	//
	//assignment := randomPermutation()
	//assignmentS, costS, stepsS, timeS := steepest(assignment, m1, m2)
	//assignmentG, costG, stepsG, timeG := greedy(assignment, m1, m2)
	//assignmentR, costR, stepsR, timeR := random(timeS, m1, m2)
	//assignmentRW, costRW, stepsRW, timeRW := randomWalk(assignment, 50, m1, m2)
	//fmt.Println(assignmentS, costS, stepsS, timeS)
	//fmt.Println(assignmentG, costG, stepsG, timeG)
	//fmt.Println(assignmentR, costR, stepsR, timeR)
	//fmt.Println(assignmentRW, costRW, stepsRW, timeRW)

	measureTime(instances[5], 10)
	fmt.Println(distance(Assignment{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, Assignment{0,1,2, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3,19}))
	fmt.Println(distance(Assignment{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, Assignment{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}))
}
