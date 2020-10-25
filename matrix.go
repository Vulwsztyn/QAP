package main

import (
	"fmt"
	"math/rand"
)

type IntMat [defaultSize][defaultSize]int

func NewRandomMatrix(maxRange int) (matrix IntMat) {
	for i := 0; i < defaultSize; i++ {
		for j := 0; j < defaultSize; j++ {
			matrix[i][j] = rand.Intn(maxRange)
		}
	}
	return
}

func (t IntMat) String() string {
	s := ""
	for _, row := range t {
		for _, n := range row {
			s += fmt.Sprintf("%d ", n)
		}
		s += fmt.Sprintln()
	}
	return s
}
