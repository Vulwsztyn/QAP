package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func fileReader(fileName string) (m1 IntMat, m2 IntMat, instanceSize int) {
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
	instanceSize, _ = strconv.Atoi(strings.TrimSpace(scanner.Text()))

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
