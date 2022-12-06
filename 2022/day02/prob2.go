package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	text := getText()
	showText(text)
	processStrategy(text)
	processStrategy2(text)
	//	ints := getENum(text)
	//	showInt(ints)
	//	x := processInts(ints)
	//	fmt.Printf("Max calories is %d\n", x)
	//	y := process2(ints)
	//	fmt.Printf("%v, %v, %v - (%v)\n", y[0], y[1], y[2], y[0]+y[1]+y[2])
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

func calculateScore(a string, b string) int {
	if a == "A" { // Opponent chooses Rock
		if b == "X" { // You choose Rock - Draw (3) + Rock (1)
			return 3 + 1
		}
		if b == "Y" { // You choose Paper - Win (6) + Paper (2)
			return 6 + 2
		}
		if b == "Z" { // You choos Scissor - Loss (0) + Scissor (3)
			return 0 + 3
		}
	}
	if a == "B" { // Opponent chooses Paper
		if b == "X" { // You choose Rock - Loss (0) + Rock (1)
			return 0 + 1
		}
		if b == "Y" { // You choose Paper - Draw (3) + Paper (2)
			return 3 + 2
		}
		if b == "Z" { // You choose Scissor - Win (6) + Scissor (3)
			return 6 + 3
		}
	}
	if a == "C" { // Opponent chooses Scissor
		if b == "X" { // You choose Rock - Win (6) + Rock (1)
			return 6 + 1
		}
		if b == "Y" { // You choose Paper - Loss (0) + Paper (2)
			return 0 + 2
		}
		if b == "Z" { // You choose Scissor - Draw (3) + Scissor (3)
			return 3 + 3
		}
	}
	return -1
}

func processStrategy(s []string) int {
	ret := 0
	f := []string{}
	for _, v := range s {
		f = strings.Split(v, " ")
		ret += calculateScore(f[0], f[1])
		fmt.Printf("(%v) (%v) - (%v) -- (%v) \n", f[0], f[1], calculateScore(f[0], f[1]), ret)
	}
	return ret
}

func figureReturn(a string, b string) string {
	if a == "A" { // Opponent chooses Rock
		if b == "X" { // X is lose
			return "Z" // Choose Scissor
		}
		if b == "Y" { // Y is draw
			return "X" // Choose Rock
		}
		if b == "Z" { // Z is win
			return "Y" // Choose Paper
		}
	}
	if a == "B" { // Opponent chooses Paper
		if b == "X" { // X is lose
			return "X" // Choose Rock
		}
		if b == "Y" { // Y is draw
			return "Y" // Choose Paper
		}
		if b == "Z" { // Z is win
			return "Z" // Choose Scissor
		}
	}
	if a == "C" { // Opponent chooses Scissor
		if b == "X" { // X is lose
			return "Y" // Choose Paper
		}
		if b == "Y" { // Y is draw
			return "Z" // Choose Scissor
		}
		if b == "Z" { // Z is win
			return "X" // Choose Rock
		}
	}
	return ""
}

func processStrategy2(s []string) int {
	ret := 0
	f := []string{}
	for _, v := range s {
		f = strings.Split(v, " ")
		ret += calculateScore(f[0], figureReturn(f[0], f[1]))
		fmt.Printf("(%v) (%v) (%v) - (%v) -- (%v) \n", f[0], f[1], figureReturn(f[0], f[1]), calculateScore(f[0], figureReturn(f[0], f[1])), ret)
	}
	return ret
}

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
