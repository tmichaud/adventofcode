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

func processFish(fish map[int]int, day int) {
	dayIndex := day % 7

	// Get Float (10 + dayIndex) and reset Float to 0
	floatVal := fish[dayIndex+10]
	fish[dayIndex+10] = 0

	// Get the count of fish for the current day
	val := fish[dayIndex]

	// Set new float value
	fish[ ((dayIndex+2) % 7) + 10] = val

	// Set current day
	fish[dayIndex] = floatVal+val

	//fmt.Printf("day = (%d) dayIndex = (%d)     floatVal = (%d) val = (%d)   map = (%v) \n", day, dayIndex, floatVal, val, fish)
}

func process1(s []string) {
	f := getNum(strings.Fields(strings.Replace(s[0], ",", " ", -1)))
	fish := map[int]int{}

	// Set a bunch of it to 0.
	for x := 0; x < 20; x++ {
		fish[x] = 0
	}

	// Populate the initial data
	for _, e:= range f {
		fish[e] += 1
	}

	for x := 0; x < 80; x++ {
		processFish(fish, x)
	}

	cnt := 0
	for x := 0; x < 20; x++ {
		if val, ok := fish[x]; ok {
			cnt += val
		}
	}

	fmt.Printf("count %d\n", cnt)
}

func process2(s []string) {
	f := getNum(strings.Fields(strings.Replace(s[0], ",", " ", -1)))
	fish := map[int]int{}

	// Set a bunch of it to 0.
	for x := 0; x < 20; x++ {
		fish[x] = 0
	}

	// Populate the initial data
	for _, e:= range f {
		fish[e] += 1
	}

	for x := 0; x < 256; x++ {
		processFish(fish, x)
	}

	cnt := 0
	for x := 0; x < 20; x++ {
		if val, ok := fish[x]; ok {
			cnt += val
		}
	}

	fmt.Printf("count %d\n", cnt)
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
