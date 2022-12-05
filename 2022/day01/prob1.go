package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	text := getText()
	showText(text)
	ints := getENum(text)
	showInt(ints)
	x := processInts(ints)
	fmt.Printf("Max calories is %d\n", x)
	y := process2(ints)
	fmt.Printf("%v, %v, %v - (%v)\n", y[0], y[1], y[2], y[0]+y[1]+y[2])
}

// Create a slice of three-measurement sliding window
func process2(ints []int) []int {
	ret := []int{}
	calories := 0
	for _, v := range ints {
		if v == -1 {
			ret = append(ret, calories)
			calories = 0
		} else {
			calories += v
		}
	}
	ret = append(ret, calories)
	sort.Sort(sort.Reverse(sort.IntSlice(ret)))
	fmt.Printf("%v", ret)
	return ret
}

// Find elf carrying the most calories
func processInts(ints []int) int {
	lastelf := 0
	maxcalories := 0
	calories := 0

	for _, v := range ints {
		calories += v
		if v == -1 {
			calories += 1
			fmt.Printf("Elf = (%v) Calories = (%v)  MaxCalories = (%v)\n", lastelf, calories, maxcalories)
			if calories >= maxcalories {
				maxcalories = calories
			}
			calories = 0
			lastelf++
		}
	}
	fmt.Printf("Elf = (%v) Calories = (%v)  MaxCalories = (%v)\n", lastelf, calories, maxcalories)
	if calories >= maxcalories {
		maxcalories = calories
	}

	return maxcalories
}

func getENum(s []string) []int {
	ret := []int{}
	for _, v := range s {
		if len(v) == 0 {
			ret = append(ret, -1)
		} else {
			i, _ := strconv.Atoi(v)
			ret = append(ret, i)
		}
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
// Helper functions
////////////////////////////////////////////////////////////////////////////////
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
