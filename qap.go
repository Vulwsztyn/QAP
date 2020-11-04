package main

import (
	"fmt"
	"math/rand"
	"time"
)

var instances = []string{"tai256c", "tho150", "wil50", "sko100c", "lipa80a", "nug30", "rou20", "kra32", "chr12c", "bur26e"}

//var instanceSizes = []int{256,150,50,100,80,30,20,32,12,26}

const defaultSize = 256
const neighbourCount = defaultSize * (defaultSize - 1) / 2

func steepest(assignment Assignment, m1 IntMat, m2 IntMat) (Assignment, int, int, int64) {
	start := time.Now()
	currentAssignment := assignment
	var bestCost, bestNeighbourCost, bestNeighbourIndex, stepCount int
	var costMatrix IntMat
	for ok := true; ok; ok = bestNeighbourCost < bestCost {
		bestCost, costMatrix = calcCost(currentAssignment, m1, m2)
		neighbours, neighboursCosts := createNeighbours(currentAssignment, m1, m2, costMatrix, bestCost, rand.Intn(defaultSize))
		bestNeighbourCost, bestNeighbourIndex = min(neighboursCosts[:])
		currentAssignment = neighbours[bestNeighbourIndex]
		println(bestCost)
		stepCount++
	}
	stop := time.Since(start)
	return currentAssignment, bestCost, stepCount, stop.Microseconds()
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
			if i != j {
				result -= previousCostMatrix[i][j]
				result -= previousCostMatrix[j][i]

				costMatrix[i][j] = m1[assignment[i]][assignment[j]] * m2[i][j]
				costMatrix[j][i] = m1[assignment[j]][assignment[i]] * m2[j][i]

				result += costMatrix[i][j]
				result += costMatrix[j][i]
			}
		}
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

func greedy(assignment Assignment, m1, m2 IntMat) (Assignment, int, int, int64) {
	start := time.Now()
	bestAssignment := assignment
	var exists bool
	var cost, stepCount int
	for ok := true; ok; {
		bestAssignment, cost, exists = bestAssignment.getFirstBetterNeighbour(m1, m2)
		ok = exists
		stepCount++
	}
	stop := time.Since(start)
	return bestAssignment, cost, stepCount, stop.Microseconds()
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

func measureTime(filename string, times int) {
	m1, m2, _ := fileReader("instances/" + filename + ".dat")
	var SArray, GArray, RWArray, RArray [][2]int
	for i := 0; i < times; i++ {
		fmt.Println(i, "iteration...")
		assignment := randomPermutation()
		fmt.Println("Steepest")
		_, costS, _, timeS := steepest(assignment, m1, m2)
		fmt.Println("Greedy")
		_, costG, _, timeG := greedy(assignment, m1, m2)
		timeLimit := (timeS + timeG) / 2
		fmt.Println("Random Walk")
		_, costRW, _, timeRW := randomWalk(assignment, timeLimit, m1, m2)
		fmt.Println("Random")
		_, costR, _, timeR := random(timeLimit, m1, m2)

		SArray = append(SArray, [2]int{costS, int(timeS)})
		GArray = append(GArray, [2]int{costG, int(timeG)})
		RWArray = append(RWArray, [2]int{costRW, int(timeRW)})
		RArray = append(RArray, [2]int{costR, int(timeR)})
	}
	writeFile(SArray, "S_"+filename)
	writeFile(GArray, "G_"+filename)
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

	measureTime(instances[0], 10)
}
