package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//	"sort"
	"strconv"
	"strings"
)

func main() {
	text := getText()
	showText(text)
//	process1(text)
	process2(text)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


func first(s string, sub string, x int, n int, v int) (int, int) {
	t := strings.Index(s, sub)
	if t != -1 && (n == -1 || t < n) {
		return t, x
	}
	return n, v
}


func last(s string, sub string, x int, n int, v int) (int, int) {
	t := strings.LastIndex(s, sub)
	if t != -1 && t > n {
		return t, x
	}
	return n, v
}


func getFirstDigitValue(s string) int {
	n := -1
	v := -1

	n,v = first(s, "0", 0, n, v)
	n,v = first(s, "1", 1, n, v)
	n,v = first(s, "2", 2, n, v)
	n,v = first(s, "3", 3, n, v)
	n,v = first(s, "4", 4, n, v)
	n,v = first(s, "5", 5, n, v)
	n,v = first(s, "6", 6, n, v)
	n,v = first(s, "7", 7, n, v)
	n,v = first(s, "8", 8, n, v)
	n,v = first(s, "9", 9, n, v)

	n,v = first(s, "one", 1, n, v)
	n,v = first(s, "two", 2, n, v)
	n,v = first(s, "three", 3, n, v)
	n,v = first(s, "four", 4, n, v)
	n,v = first(s, "five", 5, n, v)
	n,v = first(s, "six", 6, n, v)
	n,v = first(s, "seven", 7, n, v)
	n,v = first(s, "eight", 8, n, v)
	n,v = first(s, "nine", 9, n, v)

	if n == -1 {
		fmt.Printf("-- Error!\n")
	} else {
		fmt.Printf("-- %d, %d - (%v)\n", n, v, s)
	}
	return v
}


func getLastDigitValue(s string) int {
	n := -1
	v := -1

	n,v = last(s, "0", 0, n, v)
	n,v = last(s, "1", 1, n, v)
	n,v = last(s, "2", 2, n, v)
	n,v = last(s, "3", 3, n, v)
	n,v = last(s, "4", 4, n, v)
	n,v = last(s, "5", 5, n, v)
	n,v = last(s, "6", 6, n, v)
	n,v = last(s, "7", 7, n, v)
	n,v = last(s, "8", 8, n, v)
	n,v = last(s, "9", 9, n, v)

	n,v = last(s, "one", 1, n, v)
	n,v = last(s, "two", 2, n, v)
	n,v = last(s, "three", 3, n, v)
	n,v = last(s, "four", 4, n, v)
	n,v = last(s, "five", 5, n, v)
	n,v = last(s, "six", 6, n, v)
	n,v = last(s, "seven", 7, n, v)
	n,v = last(s, "eight", 8, n, v)
	n,v = last(s, "nine", 9, n, v)

	if n == -1 {
		fmt.Printf("-- Error!\n")
	} else {
		fmt.Printf("-- %d, %d - (%v)\n", n, v, s)
	}
	return v
}

func getFirstDigit(s string) int {
	for _, i := range s {
		num, err := strconv.Atoi(string(i))
		if err == nil {
			fmt.Printf("-- %d - %v  (%v) (%v)\n", num, i, string(i), s)
			return num
		}
	}
	fmt.Printf("-- ERROR -- %v\n", s)
	return -1
}

func getLastDigit(s string) int {
	return getFirstDigit(reverse(s))
}

func process1(text []string) {
	fmt.Println("-------------\n")
	x := 0
	for _, line := range text {
		if len(line) > 0 {
			num := getFirstDigit(line)*10 + getLastDigit(line)
			fmt.Printf("%v\n", num)
			x = x + num
		}
	}
	fmt.Printf("%v\n", x)
}

func process2(text []string) {
	fmt.Println("-------------\n")
	x := 0
	for _, line := range text {
		if len(line) > 0 {
			num := getFirstDigitValue(line)*10 + getLastDigitValue(line)
			fmt.Printf("%v\n", num)
			x = x + num
		}
	}
	fmt.Printf("%v\n", x)
}

/*
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
*/

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
