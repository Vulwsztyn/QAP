package main

import (
	"math/rand"
	"sort"
	"time"
)

type masterListElement struct {
	indexes [2]int
	cost    int
	used    bool
}

const masterListSize = (defaultSize) + 1
const tabuListSize = (defaultSize / 4) + 1

func sortMasterList(masterListSlice []masterListElement) {
	sort.Slice(masterListSlice, func(i, j int) bool {
		return masterListSlice[i].cost < masterListSlice[j].cost
	})
}

func swapAssignment(assignment Assignment, i int, j int) (newAssignment Assignment) {
	newAssignment = assignment
	newAssignment[i], newAssignment[j] = newAssignment[j], newAssignment[i]
	return
}

func calcMasterList(assignment Assignment, m1 IntMat, m2 IntMat, previousCostMatrix IntMat, previousCost int, startIndex int) (masterList [masterListSize]masterListElement) {
	var index, iCount int
	var allMasterListCandidates [neighbourCount]masterListElement
	for i := startIndex; index < neighbourCount; i = (i + 1) % defaultSize {
		for j := (i + 1) % defaultSize; j != positiveReminder(i-iCount, defaultSize); j = (j + 1) % defaultSize {
			tmpAssignment := swapAssignment(assignment, i, j)
			tmpCost, _ := reCalcCost(tmpAssignment, m1, m2, previousCostMatrix, previousCost, [2]int{i, j})
			allMasterListCandidates[index] = masterListElement{[2]int{i, j}, tmpCost, false}
			index++
		}
		iCount++
	}
	//for _, v := range allMasterListCandidates {
	//	println(v.indexes[0], v.indexes[1], v.cost)
	//}
	sortMasterList(allMasterListCandidates[:])
	//for _, v := range allMasterListCandidates {
	//	println(v.indexes[0], v.indexes[1], v.cost)
	//}
	copy(masterList[:], allMasterListCandidates[:masterListSize])
	return
}
func reCalcMasterListCosts(masterList [masterListSize]masterListElement, assignment Assignment, m1 IntMat, m2 IntMat, previousCostMatrix IntMat, previousCost int) ([masterListSize]masterListElement, int) {
	var solutionsExplored int
	for index, v := range masterList {
		if v.used {
			continue
		}
		i, j := v.indexes[0], v.indexes[1]
		tmpAssignment := swapAssignment(assignment, i, j)
		masterList[index].cost, _ = reCalcCost(tmpAssignment, m1, m2, previousCostMatrix, previousCost, [2]int{i, j})
		solutionsExplored++
	}
	sortMasterList(masterList[:])
	return masterList, solutionsExplored
}

func TS(assignment Assignment, m1, m2 IntMat) (Assignment, int, int, int, int64) {

	start := time.Now()
	var masterList [masterListSize]masterListElement
	var solutionsExplored, threshold, noBettermentStepCount, tmpSolutionsExplored int
	var tabuMatrix IntMat
	currentAssignment := assignment
	currentCost, currentCostMatrix := calcCost(assignment, m1, m2)
	bestAssignment, bestCost := assignment, currentCost
	shouldCalculateMasterList := true
	stepCount := 1

	for shouldContinueLoop := true; shouldContinueLoop; {
		//for _, v := range masterList {
		//	fmt.Println(v.indexes[0], v.indexes[1], v.cost, v.used)
		//}
		//fmt.Println(shouldCalculateMasterList)
		//fmt.Println()
		if shouldCalculateMasterList {
			shouldCalculateMasterList = false
			masterList = calcMasterList(currentAssignment, m1, m2, currentCostMatrix, currentCost, rand.Intn(defaultSize))
			threshold = masterList[masterListSize-1].cost
			solutionsExplored += neighbourCount
		} else {
			masterList, tmpSolutionsExplored = reCalcMasterListCosts(masterList, currentAssignment, m1, m2, currentCostMatrix, currentCost)
			solutionsExplored += tmpSolutionsExplored
		}

		//for _, v := range masterList {
		//	fmt.Println(v.indexes[0], v.indexes[1], v.cost, v.used)
		//}
		//fmt.Println(shouldCalculateMasterList)
		//fmt.Println()
		//fmt.Println(tabuMatrix)
		shouldCalculateMasterList = true
		noBettermentStepCount++
		//fmt.Println(noBettermentStepCount)

		for indexInMasterList, v := range masterList {
			if v.used {
				continue
			}
			if v.cost > threshold {
				break
			}
			i, j := v.indexes[0], v.indexes[1]
			lastUsedInStep := tabuMatrix[i][j]
			//if lastUsedInStep > 0 && stepCount-lastUsedInStep <= tabuListSize && v.cost >= currentCost{
			//	fmt.Println(v.cost, currentCost,v.cost < currentCost)
			//	fmt.Println("tabu")
			//}
			if lastUsedInStep == 0 || stepCount-lastUsedInStep > tabuListSize || v.cost < currentCost {
				if v.cost < bestCost {
					noBettermentStepCount = 0
				}
				currentAssignment[i], currentAssignment[j] = currentAssignment[j], currentAssignment[i]
				currentCost = v.cost
				currentCostMatrix = reCalcCostMatrixOnly(currentAssignment, m1, m2, currentCostMatrix, [2]int{i, j})
				if currentCost < bestCost {
					bestCost, bestAssignment = currentCost, currentAssignment
				}
				masterList[indexInMasterList].used = true
				tabuMatrix[i][j], tabuMatrix[j][i] = stepCount, stepCount
				stepCount++
				shouldCalculateMasterList = false
				break
			}
		}

		shouldContinueLoop = noBettermentStepCount < 10*defaultSize
	}
	stop := time.Since(start)
	//fmt.Println(bestCost)
	return bestAssignment, bestCost, stepCount - 1, solutionsExplored, stop.Microseconds()
}
