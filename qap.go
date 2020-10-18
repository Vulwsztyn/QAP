package main

import (
	"fmt"
	"math/rand"
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

func main() {
	size := 10
	maxRange := 10
	m1 := genMatrix(size, maxRange)
	m2 := genMatrix(size, maxRange)
	fmt.Println(matrix2String(m1))
	fmt.Println(matrix2String(m2))
}
