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

const defaultSize = 26

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

func randomPermutation(size int) []int {
	_range := makeRange(0, size)
	result := make([]int, size)
	for i := 0; i < size; i++ {
		j := rand.Intn(size - i)
		result[i] = _range[j]
		_range[j] = _range[len(_range)-1]
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
	maxRange := 100
	start := time.Now()

	m1 := NewRandomMatrix(maxRange)
	m2 := NewRandomMatrix(maxRange)

	stop := time.Since(start)
	timeSplits = append(timeSplits, stop.Microseconds())

	fmt.Println(m1)
	fmt.Println(m2)

	fmt.Println(timeSplits)

	fmt.Println(randomPermutation(8))
	fmt.Println(randomPermutation(8))
	fmt.Println(randomPermutation(8))

	timeOfGeneration := measureTime(func() {
		randomPermutation(999999)
	})
	fmt.Println(timeOfGeneration)

	fmt.Println(fileReader("instances/bur26a.dat"))
}
