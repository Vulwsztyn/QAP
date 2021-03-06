package main

import (
	"fmt"
	"math/rand"
)

type IntMat [defaultSize][defaultSize]int

func NewRandomMatrix(maxRange int, minRange int) (matrix IntMat) {
	delta := maxRange - minRange
	for i := 0; i < defaultSize; i++ {
		for j := 0; j < defaultSize; j++ {
			matrix[i][j] = rand.Intn(delta) + minRange
		}
	}
	return
}

func (m1 IntMat) permuteMatrix(assignment Assignment) (matrix IntMat) {
	for i := 0; i < defaultSize; i++ {
		for j := 0; j < defaultSize; j++ {
			matrix[i][j] = m1[assignment[i]][assignment[j]]
		}
	}
	return
}

func (m1 IntMat) sum() (result int) {
	for i := 0; i < defaultSize; i++ {
		for j := 0; j < defaultSize; j++ {
			result += m1[i][j]
		}
	}
	return
}

func (m1 IntMat) String() string {
	s := ""
	for _, row := range m1 {
		for _, n := range row {
			s += fmt.Sprintf("%d ", n)
		}
		s += fmt.Sprintln()
	}
	return s
}
