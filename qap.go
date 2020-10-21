package main

import (
	"fmt"
	"math/rand"
	"time"
)

type IntMat [][]int

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

func genMatrix(size int, maxRange int)(matrix IntMat){
	for i:=0; i<size; i++ {
		var row []int
		for j:=0; j<size; j++ {
			row = append(row,rand.Intn(maxRange))
		}
		matrix = append(matrix, row)
	}
	return
}

func permutateIndexes(array *[]int)(res []int){
	res = rand.Perm(len(*array))
	return
}

func chooseTwo(array *[]int)(a int, b int){
	a = rand.Intn(len(*array))
	b = rand.Intn(len(*array))
	for a==b{
		b = rand.Intn(len(*array))
	}
	return
}

func main() {
	var timeSplits []int64
	size := 10
	maxRange := 10
	start := time.Now()

	m1 := genMatrix(size, maxRange)
	m2 := genMatrix(size, maxRange)

	stop := time.Since(start)
	timeSplits = append(timeSplits,stop.Microseconds())

	fmt.Println(matrix2String(m1))
	fmt.Println(matrix2String(m2))

	fmt.Println(timeSplits)

	fmt.Println(permutateIndexes(&m1[0]))
	fmt.Println(chooseTwo(&m1[0]))
}
