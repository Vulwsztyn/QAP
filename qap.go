package main

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultSize = 10

type IntMat [defaultSize][defaultSize]int

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
		var row [defaultSize]int
		for j:=0; j<size; j++ {
			row[j] = rand.Intn(maxRange)
		}
		matrix[i] = row
	}
	return
}

func main() {
	var timeSplits []int64
	maxRange := 100
	start := time.Now()

	m1 := genMatrix(defaultSize, maxRange)
	m2 := genMatrix(defaultSize, maxRange)

	stop := time.Since(start)
	timeSplits = append(timeSplits,stop.Microseconds())

	fmt.Println(matrix2String(m1))
	fmt.Println(matrix2String(m2))

	fmt.Println(timeSplits)
}
