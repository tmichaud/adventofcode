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
	process1(text)
	fmt.Printf("-------------------------------\n")
	process2(text)
}

// Interesting search pattern
func process2(s []string) {
	oxy := subprocess2(s, 0, 1)
	co2 := subprocess2(s, 0, 0)
	fmt.Printf("\noxy (%v) co2(%v) -- (%v) \n", oxy, co2, oxy*co2)
}

// s is our input (we will be changing it)
// bit is our bit...left ot right
// mm is oxygen (1) or c02 (0)
func subprocess2(s []string, bit int, mm int) int {
	m0 := 0
	m1 := 0
	ret := []string{}
	ans := 0

	// For all of our input
	for _, v := range s {
		b := strings.Split(v, "")
		// This time were only looking at a single bit
		if b[bit] == "0" {
			m0++
		} else {
			m1++
		}
	}

	for _, v := range s {
		b := strings.Split(v, "")

		// If oxygen
		if mm == 1 {
			if m1 >= m0 {
				if b[bit] == "1" {
					ret = append(ret, v)
				}
			} else {
				if b[bit] == "0" {
					ret = append(ret, v)
				}
			}

		} else {
			if m0 <= m1 {
				if b[bit] == "0" {
					ret = append(ret, v)
				}
			} else {
				if b[bit] == "1" {
					ret = append(ret, v)
				}
			}
		}
	}

	if len(ret) == 1 {
		fmt.Printf("Found (%v)\n", ret[0])
		x := strings.Split(ret[0], "")
		for i, e := range x {
			if e == "1" {
				ans = ans + pow(len(x)-1-i)
			}
		}
		return ans
	} else {
		return subprocess2(ret, bit+1, mm)
	}
}

// Ugly function...and there may be better ways to do it
//
// m0 contains the count of 0s per position (left to right)
// m1 contains the count of 1s per position (left to right)
func process1(s []string) {
	m0 := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0}
	m1 := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0}
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	// Take each line of input
	// Split it into digits, left to right
	// Walk the digits, setting m0 and m1 as necessary
	for _, v := range s {
		b := strings.Split(v, "")
		for i, e := range b {
			if e == "0" {
				m0[i]++
			} else {
				m1[i]++
			}
		}
	}

	// Now, walk a count 0-4  (could'e done a for loop)
	// if m1 > m0...add the value to x. (Gamma)   -- Makes sure you invert the pow (len(digits)-1-i)
	// if m0 > m1...add teh value to y. (epsilon) -- Makes sure you invert the pow (len(digits)-1-i)
	x := 0
	y := 0

	for _, i := range digits {
		if m0[i] > m1[i] {
			y = y + pow(len(digits)-1-i)
		} else {
			x = x + pow(len(digits)-1-i)
		}
	}
	fmt.Printf(" %d %d -- (%d) \n", x, y, x*y)
}

func pow(i int) int {
	switch i {
	case 0:
		return 1
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 8
	case 4:
		return 16
	case 5:
		return 32
	case 6:
		return 64
	case 7:
		return 128
	case 8:
		return 256
	case 9:
		return 512
	case 10:
		return 1024
	case 11:
		return 2048
	case 12:
		return 4096

	}
	fmt.Printf(" Shouldn't get here!  ERROR! \n")
	return -1
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
