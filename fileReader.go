package main

import (
	"bufio"
	"fmt"
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

	columnNum := 0
	rowNum := 0
	whichMatrix := 1
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	instanceSize, _ = strconv.Atoi(strings.TrimSpace(scanner.Text()))

	for scanner.Scan() {
		row := strings.Fields(strings.TrimSpace(scanner.Text()))
		for _, v := range row {
			if whichMatrix == 1 {
				m1[rowNum][columnNum], _ = strconv.Atoi(v)
			} else {
				m2[rowNum][columnNum], _ = strconv.Atoi(v)
			}
			columnNum++
			if columnNum >= defaultSize {
				columnNum = 0
				rowNum++
			}
		}
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

func writeFile(results [][5]int, filename string) {
	f, err := os.Create("results/" + filename + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, value := range results {
		_, err = f.WriteString(fmt.Sprint(value[0]) + " " + fmt.Sprint(value[1]) + " " + fmt.Sprint(value[2]) + " " + fmt.Sprint(value[3]) + " " + fmt.Sprint(value[4]) + "\n")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
