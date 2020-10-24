package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const defaultSize = 3

type IntMat [defaultSize][defaultSize]int

func fileReader(fileName string) (m1 IntMat, m2 IntMat) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	rowNum := 0
	whichMatrix := 1
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	instanceSize, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	for scanner.Scan() {
		row := strings.Fields(strings.TrimSpace(scanner.Text()))
		if len(row) < instanceSize {
			continue
		}
		for i, v := range row {
			if whichMatrix == 1 {
				m1[rowNum][i], _ = strconv.Atoi(v)
			} else {
				m2[rowNum][i], _ = strconv.Atoi(v)
			}
		}
		rowNum++
		if rowNum >= instanceSize {
			rowNum = 0
			whichMatrix++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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

func translateAssignment(assignment [defaultSize]int) (result [defaultSize]int){
	for i:=0; i<defaultSize; i++ {
		result[assignment[i]] = i
	}
	return
}

func calcCost(assignment [defaultSize]int ,m1 IntMat, m2 IntMat) (result int){
	for i:=0; i<defaultSize; i++ {
		for j:=0; j<defaultSize; j++ {
			result += m1[assignment[i]][assignment[j]] * m2[i][j]
		}
	}
	return
}

func NewRandomMatrix(maxRange int) (matrix IntMat) {
	for i := 0; i < defaultSize; i++ {
		for j := 0; j < defaultSize; j++ {
			matrix[i][j] = rand.Intn(maxRange)
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

func randomPermutation() [defaultSize]int {
	_range := makeRange(0, defaultSize)
	var result [defaultSize]int
	for i := 0; i < defaultSize; i++ {
		j := rand.Intn(defaultSize - i)
		result[i] = _range[j]
		_range[j] = _range[len(_range)-1-i]
	}
	return result
}

func measureTime(fn func()) int64 {
	start := time.Now()
	fn()
	stop := time.Since(start)
	return stop.Microseconds()
}

func main() {
	var timeSplits []int64
	maxRange := 5
	start := time.Now()

	m1 := NewRandomMatrix(maxRange)
	m2 := NewRandomMatrix(maxRange)

	stop := time.Since(start)
	timeSplits = append(timeSplits, stop.Microseconds())

	fmt.Println(timeSplits)

	testAssignment:=randomPermutation()
	fmt.Println(testAssignment)
	fmt.Println(translateAssignment(testAssignment))

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(testAssignment)
	fmt.Println(calcCost(testAssignment,m1,m2))
	//fmt.Println(fileReader("instances/chr12a.dat"))
}
