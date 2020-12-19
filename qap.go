package main

import (
	"fmt"
	"math/rand"
)

//var instances = []string{"tai256c", "tho150", "wil50", "sko100c", "lipa80a", "nug30", "rou20", "kra32", "chr12c", "bur26e"}
var instances = []string{"bur26d", "kra30a", "tho40", "wil50", "lipa60a", "lipa70a", "tai80a", "sko90", "sko100a", "esc128"}

const defaultSize = int(
	(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((0-1)*(0-2)*(0-3)*(0-4)*(0-5)*(0-6)*(0-7)*(0-8)*(0-9))*26 +
		(instanceIndex-0)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((1-0)*(1-2)*(1-3)*(1-4)*(1-5)*(1-6)*(1-7)*(1-8)*(1-9))*30 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((2-0)*(2-1)*(2-3)*(2-4)*(2-5)*(2-6)*(2-7)*(2-8)*(2-9))*40 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((3-0)*(3-1)*(3-2)*(3-4)*(3-5)*(3-6)*(3-7)*(3-8)*(3-9))*50 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((4-0)*(4-1)*(4-2)*(4-3)*(4-5)*(4-6)*(4-7)*(4-8)*(4-9))*60 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((5-0)*(5-1)*(5-2)*(5-3)*(5-4)*(5-6)*(5-7)*(5-8)*(5-9))*70 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-7)*(instanceIndex-8)*(instanceIndex-9)/((6-0)*(6-1)*(6-2)*(6-3)*(6-4)*(6-5)*(6-7)*(6-8)*(6-9))*80 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-8)*(instanceIndex-9)/((7-0)*(7-1)*(7-2)*(7-3)*(7-4)*(7-5)*(7-6)*(7-8)*(7-9))*90 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-9)/((8-0)*(8-1)*(8-2)*(8-3)*(8-4)*(8-5)*(8-6)*(8-7)*(8-9))*100 +
		(instanceIndex-0)*(instanceIndex-1)*(instanceIndex-2)*(instanceIndex-3)*(instanceIndex-4)*(instanceIndex-5)*(instanceIndex-6)*(instanceIndex-7)*(instanceIndex-8)/((9-0)*(9-1)*(9-2)*(9-3)*(9-4)*(9-5)*(9-6)*(9-7)*(9-8))*128)

//var instanceSizes = []int{256,150,50,100,80,30,20,32,12,26}

const neighbourCount = defaultSize * (defaultSize - 1) / 2

func optimalAssignment() (assignment Assignment) {
	bestAssignments := [][]int{
		{17, 3, 1, 22, 18, 20, 23, 6, 9, 24, 21, 10, 0, 16, 8, 4, 25, 12, 11, 13, 7, 19, 14, 2, 15, 5},
		{25, 23, 22, 15, 19, 18, 5, 9, 10, 1, 21, 17, 6, 29, 14, 20, 24, 28, 11, 8, 4, 16, 0, 7, 12, 27, 13, 2, 3, 26},
		{12, 1, 7, 25, 34, 21, 5, 29, 24, 13, 14, 22, 35, 9, 11, 19, 15, 16, 2, 10, 28, 26, 3, 4, 31, 36, 37, 17, 30, 20, 39, 27, 18, 6, 38, 23, 33, 32, 8, 0},
		{0, 28, 14, 45, 11, 33, 49, 37, 8, 31, 34, 18, 38, 10, 4, 23, 44, 48, 7, 29, 32, 30, 17, 15, 19, 39, 9, 41, 21, 22, 3, 12, 5, 46, 20, 47, 35, 16, 36, 2, 24, 42, 27, 13, 25, 40, 6, 1, 26, 43},
		{8, 30, 17, 3, 49, 31, 57, 18, 16, 25, 51, 15, 5, 46, 38, 6, 2, 41, 59, 50, 35, 10, 27, 55, 0, 44, 11, 43, 28, 36, 22, 48, 19, 13, 26, 37, 9, 14, 39, 24, 7, 56, 23, 45, 53, 12, 20, 1, 34, 4, 32, 42, 54, 21, 52, 33, 29, 58, 40, 47},
		{40, 11, 13, 67, 7, 29, 66, 58, 44, 8, 17, 22, 9, 36, 60, 35, 57, 45, 21, 10, 26, 32, 54, 24, 2, 27, 28, 68, 33, 39, 16, 18, 34, 4, 5, 64, 0, 43, 1, 69, 41, 49, 59, 56, 48, 61, 55, 38, 3, 53, 37, 51, 65, 62, 14, 63, 42, 12, 15, 52, 6, 31, 47, 30, 25, 50, 23, 46, 19, 20},
		{10, 18, 63, 50, 77, 53, 30, 79, 76, 13, 36, 58, 69, 26, 33, 25, 61, 59, 28, 16, 68, 48, 66, 78, 6, 17, 72, 38, 14, 55, 54, 57, 4, 22, 52, 47, 51, 62, 12, 43, 73, 5, 60, 32, 0, 40, 2, 42, 29, 8, 21, 46, 35, 19, 49, 24, 27, 41, 1, 65, 3, 75, 44, 23, 64, 70, 37, 15, 56, 11, 67, 71, 9, 39, 7, 20, 45, 74, 34, 31},
		{58, 59, 33, 57, 89, 46, 34, 9, 2, 62, 80, 42, 37, 47, 32, 14, 38, 81, 35, 83, 68, 70, 40, 88, 19, 65, 1, 44, 49, 21, 60, 76, 84, 50, 75, 41, 43, 7, 86, 16, 5, 79, 28, 15, 71, 82, 55, 10, 22, 20, 45, 6, 39, 56, 72, 51, 24, 18, 11, 61, 17, 77, 26, 69, 53, 25, 87, 73, 31, 52, 30, 85, 67, 27, 12, 4, 48, 36, 8, 23, 66, 29, 54, 13, 74, 0, 78, 3, 63, 64},
		{6, 48, 90, 34, 11, 29, 97, 53, 41, 60, 2, 82, 12, 27, 39, 32, 99, 9, 46, 69, 92, 35, 94, 61, 16, 74, 20, 49, 93, 65, 25, 36, 3, 37, 91, 87, 8, 85, 10, 62, 89, 50, 26, 96, 80, 77, 75, 95, 47, 81, 67, 33, 88, 22, 13, 55, 38, 70, 51, 40, 84, 31, 17, 78, 0, 72, 63, 52, 68, 71, 4, 83, 66, 14, 24, 98, 28, 7, 18, 21, 56, 57, 1, 58, 64, 42, 59, 30, 45, 79, 19, 86, 54, 23, 15, 44, 5, 43, 76, 73},
		{79, 74, 65, 78, 70, 72, 76, 68, 19, 5, 48, 24, 114, 80, 121, 118, 84, 90, 31, 27, 47, 86, 8, 4, 122, 125, 15, 25, 46, 93, 119, 91, 42, 1, 10, 71, 11, 26, 127, 45, 61, 56, 20, 85, 29, 57, 97, 9, 124, 77, 115, 44, 83, 69, 40, 53, 105, 98, 37, 17, 51, 33, 6, 104, 35, 109, 34, 62, 30, 89, 32, 110, 54, 59, 95, 22, 82, 116, 7, 81, 66, 28, 113, 96, 50, 107, 75, 112, 100, 123, 13, 64, 52, 87, 106, 58, 39, 43, 126, 36, 41, 23, 111, 103, 63, 73, 14, 16, 94, 18, 49, 12, 3, 101, 102, 38, 55, 67, 108, 117, 99, 120, 0, 92, 2, 60, 88, 21},
	}
	copy(assignment[:], bestAssignments[instanceIndex])
	return
}

func measureTime(filename string, times int) {
	m1, m2, _ := fileReader("instances/" + filename + ".dat")
	var SAArray, SArray, GArray, HArray, RWArray, RArray [][5]int
	var SADists, SDists, GDists, HDists, RWDists, RDists []float64
	for i := 0; i < times; i++ {
		fmt.Println(i, "iteration...")
		assignment := randomPermutation()
		assignmentCost, _ := calcCost(assignment, m1, m2)
		fmt.Println("SA")
		bestSA, costSA, stepsSA, exploreSolutionsSA, timeSA := SA(assignment, m1, m2, neighbourCount, 0.95, 10, 0.01)
		fmt.Println("Steepest")
		bestS, costS, stepsS, exploreSolutionsS, timeS := steepest(assignment, m1, m2)
		fmt.Println("Greedy")
		bestG, costG, stepsG, exploreSolutionsG, timeG := greedy(assignment, m1, m2)
		fmt.Println("Heuristic")
		bestH, costH, stepsH, exploreSolutionsH, timeH := heuristic(assignment, m1, m2)
		timeLimit := (timeS + timeG) / 2
		fmt.Println("Random Walk")
		bestRW, costRW, exploreSolutionsRW, timeRW := randomWalk(assignment, timeLimit, m1, m2)
		fmt.Println("Random")
		bestR, costR, exploreSolutionsR, timeR := random(timeLimit, m1, m2)

		SAArray = append(SArray, [5]int{costSA, stepsSA, exploreSolutionsSA, int(timeSA), assignmentCost})
		GArray = append(GArray, [5]int{costG, stepsG, exploreSolutionsG, int(timeG), assignmentCost})
		SArray = append(SArray, [5]int{costS, stepsS, exploreSolutionsS, int(timeS), assignmentCost})
		HArray = append(HArray, [5]int{costH, stepsH, exploreSolutionsH, int(timeH), assignmentCost})
		RWArray = append(RWArray, [5]int{costRW, -1, exploreSolutionsRW, int(timeRW), assignmentCost})
		RArray = append(RArray, [5]int{costR, -1, exploreSolutionsR, int(timeR), assignmentCost})

		SADists = append(SADists, distance(bestSA, optimalAssignment()))
		GDists = append(GDists, distance(bestG, optimalAssignment()))
		SDists = append(SDists, distance(bestS, optimalAssignment()))
		HDists = append(HDists, distance(bestH, optimalAssignment()))
		RWDists = append(RWDists, distance(bestRW, optimalAssignment()))
		RDists = append(RDists, distance(bestR, optimalAssignment()))
	}
	writeFile("SA_"+filename+"_"+fmt.Sprintf("%d", times), SAArray, SADists)
	writeFile("S_"+filename+"_"+fmt.Sprintf("%d", times), SArray, SDists)
	writeFile("G_"+filename+"_"+fmt.Sprintf("%d", times), GArray, GDists)
	writeFile("H_"+filename+"_"+fmt.Sprintf("%d", times), HArray, HDists)
	writeFile("RW_"+filename+"_"+fmt.Sprintf("%d", times), RWArray, RWDists)
	writeFile("R_"+filename+"_"+fmt.Sprintf("%d", times), RArray, RDists)
	fmt.Println("done")
}

func similarities(filename string, times int) {
	m1, m2, _ := fileReader("instances/" + filename + ".dat")
	var SArray, GArray []Assignment
	//var SDists, GDists, HDists, RWDists, RDists []float64
	for i := 0; i < times; i++ {
		fmt.Println(i, "iteration...")
		assignment := randomPermutation()
		//assignmentCost, _ := calcCost(assignment, m1, m2)
		fmt.Println("Steepest")
		bestS, _, _, _, _ := steepest(assignment, m1, m2)
		fmt.Println("Greedy")
		bestG, _, _, _, _ := greedy(assignment, m1, m2)
		//fmt.Println("Heuristic")
		//bestH, costH, stepsH, exploreSolutionsH, timeH := heuristic(assignment, m1, m2)
		//timeLimit := (timeS + timeG) / 2
		//fmt.Println("Random Walk")
		//bestRW, costRW, exploreSolutionsRW, timeRW := randomWalk(assignment, timeLimit, m1, m2)
		//fmt.Println("Random")
		//bestR, costR, exploreSolutionsR, timeR := random(timeLimit, m1, m2)

		GArray = append(GArray, bestG)
		SArray = append(SArray, bestS)
		//HArray = append(HArray, bestH)
		//RWArray = append(RWArray, bestRW)
		//RArray = append(RArray, bestR)
		//
		//GDists = append(GDists, distance(bestG,optimalAssignment()))
		//SDists = append(SDists, distance(bestS,optimalAssignment()))
		//HDists = append(HDists, distance(bestH,optimalAssignment()))
		//RWDists = append(RWDists, distance(bestRW,optimalAssignment()))
		//RDists = append(RDists, distance(bestR,optimalAssignment()))
	}
	//writeFile("S_"+filename, SArray, SDists)
	//writeFile("G_"+filename, GArray, GDists)
	//writeFile("H_"+filename, HArray, HDists)
	//writeFile("RW_"+filename, RWArray, RWDists)
	//writeFile("R_"+filename, RArray, RDists)
	fmt.Println(GArray)
	fmt.Println(SArray)
	println("S_Sim" + filename)
	writeSimilaritiesFile("S_Sim_"+filename, SArray)
	writeSimilaritiesFile("G_Sim_"+filename, GArray)
	fmt.Println("done")
}

func main() {
	rand.Seed(123)
	//filename := "instances/chr12a.dat"
	//m1, m2, _ := fileReader("instances/" + instances[instanceIndex] + ".dat")
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

	measureTime(instances[instanceIndex], 1)

	//fmt.Println(distance(Assignment{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, Assignment{0, 1, 2, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 19}))
	//fmt.Println(distance(Assignment{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, Assignment{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}))
}
