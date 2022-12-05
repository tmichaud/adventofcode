package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	text := getText()
	//showText(text)
	process1(text)
	fmt.Printf("-------------------------------\n")
	process2(text)
}

func processMatrix2(matrix [][]int, line []int) [][]int {
	fmt.Printf("%d\n", line)
	if line[0] == line[2] {
		//fmt.Printf("Procesing vertical %d %d\n", line[0], line[2]);
		if line[1] > line[3] {
			line[1], line[3] = line[3], line[1]
		}
		for x := line[1]; x <= line[3]; x++ {
			matrix[x][line[0]]++
		}
	} else {

		if line[1] == line[3] {
			//fmt.Printf("Procesing horizontal %d %d\n", line[1], line[3]);
			if line[0] > line[2] {
				line[0], line[2] = line[2], line[0]
			}
			for x := line[0]; x <= line[2]; x++ {
				//fmt.Printf("%d %d %v\n", line[1], x, matrix[line[1]])
				matrix[line[1]][x]++
			}
		} else {
			//fmt.Printf("Procesing diagonal %d %d  %d %d\n", line[0], line[1], line[2], line[3]);
			// Is it a slash 0,0 -> 4,4 or a  4,4 -> 0,0
			if (((line[0] > line[2]) && (line[1] > line[3])) || ((line[2] > line[0]) && (line[3] > line[1]))) {
				if line[0] > line[2] && line[1] > line[3] {
					line[0], line[2] = line[2], line[0]
					line[1], line[3] = line[3], line[1]
				}
				//fmt.Printf("Processing slash - %d, %d to %d, %d\n", line[0], line[1], line[2], line[3])
				for x := 0; x+line[0] <= line[2]; x++ {
					//fmt.Printf("x = %d\n")
					matrix[line[1]+x][line[0]+x]++
				}
			}
			// Or a backslash 4,0 -> 0,4 or a 0,4 -> 4,0
			if (((line[0] > line[2]) && (line[1] < line[3])) || ((line[2] > line[0]) && (line[3] < line[1]))) {
				if line[0] > line[2] && line[1] < line[3] {
					line[0], line[2] = line[2], line[0]
					line[1], line[3] = line[3], line[1]
				}
				//fmt.Printf("Processing backslash - %d, %d to %d, %d\n", line[0], line[1], line[2], line[3])
				for x := 0; x+line[0] <= line[2]; x++ {
					matrix[line[1]-x][line[0]+x]++
				}
			}

		}
	}
	return matrix
}

func process2(s []string) {
	size := getNum(s[0:1])[0]

	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}

	//showMatrix(matrix, size)

	for _, e := range s[1:] {
		matrix = processMatrix2(matrix, getNum(strings.Fields(strings.Replace(strings.Replace(e, "->", "", -1), ",", " ", -1))))
		//showMatrix(matrix, size)
	}

	//showMatrix(matrix, size)

	fmt.Printf("Matrix count = %d\n", countMatrix(matrix, size))

}

func showMatrix(matrix [][]int, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("%v\n", matrix[i])
	}
}

func countMatrix(matrix [][]int, size int) int {
	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[i][j] > 1 {
				count++
			}
		}
	}
	return count
}

func processMatrix(matrix [][]int, line []int) [][]int {
	//fmt.Printf("%d\n", line)
	if line[0] == line[2] {
		//fmt.Printf("Procesing vertical %d %d\n", line[0], line[2]);
		if line[1] > line[3] {
			line[1], line[3] = line[3], line[1]
		}
		for x := line[1]; x <= line[3]; x++ {
			matrix[x][line[0]]++
		}
	}
	if line[1] == line[3] {
		//fmt.Printf("Procesing horizontal %d %d\n", line[1], line[3]);
		if line[0] > line[2] {
			line[0], line[2] = line[2], line[0]
		}
		for x := line[0]; x <= line[2]; x++ {
			//fmt.Printf("%d %d %v\n", line[1], x, matrix[line[1]])
			matrix[line[1]][x]++
		}
	}
	return matrix
}

func process1(s []string) {
	size := getNum(s[0:1])[0]

	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}

	//showMatrix(matrix, size)

	for _, e := range s[1:] {
		matrix = processMatrix(matrix, getNum(strings.Fields(strings.Replace(strings.Replace(e, "->", "", -1), ",", " ", -1))))
		//showMatrix(matrix, size)
	}

	fmt.Printf("Matrix count = %d\n", countMatrix(matrix, size))
}

////////////////////////////////////////////////////////////////////////////////
// Helper functions
////////////////////////////////////////////////////////////////////////////////
func getInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func showText(text []string) {
	for _, line := range text {
		fmt.Printf("%v\n", line)
	}
}

func showInt(ints []int) {
	for _, e := range ints {
		fmt.Printf("%d\n", e)
	}
}

func getNum(s []string) []int {
	ret := []int{}
	for _, x := range s {
		i, _ := strconv.Atoi(x)
		ret = append(ret, i)
	}
	return ret
}

func getText() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}
