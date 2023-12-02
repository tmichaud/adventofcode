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
	//process1(text)
	process2(text)
}


func getColor(num string, color string, r int, b int, g int) (int, int, int) {
	n, _ := strconv.Atoi(num)
	if "green" == color {
		if n > g {
			return r, b, n
		}
	}
	if "blue" == color {
		if n > b {
			return r, n, g
		}
	}
	if "red" == color {
		if n > r {
			return n, b, g
		}
	}
	if "green" != color && "blue" != color && "red" != color {
		fmt.Printf("How the hell did we get here?  (%v) (%v) %d %d %d\n", num, color, r, b, g)
	}
	return r, b, g
}

func getPower(s string) int {
	r,b,g := 0,0,0

	a := strings.Split(strings.ToLower(strings.ReplaceAll(s, ";", ",")), ":")
	cubes := strings.Split(a[1], ",")
	for _, c := range cubes {
		c = strings.TrimSpace(c)
		f := strings.Split(c, " ")
		r,b,g = getColor(f[0], f[1], r, b, g)
	}
	return r * b * g
}

func getGameID(s string) int {
	a := strings.Split(s, ":")
	x, _ := strconv.Atoi(a[0][5:])
	if !testCubes(s) {
		return 0
	}
	return x
}


func testCubes(s string) bool {
	a := strings.Split(strings.ReplaceAll(s, ";", ","), ":")
	cubes := strings.Split(a[1], ",")
	for _, c := range cubes {
		c = strings.TrimSpace(c)
		f := strings.Split(c, " ")
		if "blue" == strings.ToLower(f[1]) {
			n, _ := strconv.Atoi(f[0])
			if n > 14 {
				return false
			}
		} else {
			if "red" == strings.ToLower(f[1]) {
				n, _ := strconv.Atoi(f[0])
				if n > 12 {
					return false
				}
			} else {
				if "green" == strings.ToLower(f[1]) {
					n, _ := strconv.Atoi(f[0])
					if n > 13 {
						return false
					}
				} else {
					fmt.Printf("How did we get here??? (%v)\n", s)
				}
			}
		}
	}
	return true
}

func process1(text []string) {
	fmt.Println("-------------\n")
	sum := 0
	for _, s := range text {
		sum = sum + getGameID(s)
		fmt.Printf("%v\n", getGameID(s))
	}
	fmt.Printf("Sum is (%v)\n", sum)
}

func process2(text []string) {
	fmt.Println("-------------\n")
	sum := 0
	for _, s := range text {
		sum = sum + getPower(s)
		fmt.Printf("%v\n", getPower(s))
	}
	fmt.Printf("Sum is (%v)\n", sum)
}

////////////////////////////////////////////////////////////////////////////////
// Helper functions
////////////////////////////////////////////////////////////////////////////////
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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

////////////////////////////////////////////////////////////////////////////////
// Old stuff
////////////////////////////////////////////////////////////////////////////////

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

	n, v = first(s, "0", 0, n, v)
	n, v = first(s, "1", 1, n, v)
	n, v = first(s, "2", 2, n, v)
	n, v = first(s, "3", 3, n, v)
	n, v = first(s, "4", 4, n, v)
	n, v = first(s, "5", 5, n, v)
	n, v = first(s, "6", 6, n, v)
	n, v = first(s, "7", 7, n, v)
	n, v = first(s, "8", 8, n, v)
	n, v = first(s, "9", 9, n, v)

	n, v = first(s, "one", 1, n, v)
	n, v = first(s, "two", 2, n, v)
	n, v = first(s, "three", 3, n, v)
	n, v = first(s, "four", 4, n, v)
	n, v = first(s, "five", 5, n, v)
	n, v = first(s, "six", 6, n, v)
	n, v = first(s, "seven", 7, n, v)
	n, v = first(s, "eight", 8, n, v)
	n, v = first(s, "nine", 9, n, v)

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

	n, v = last(s, "0", 0, n, v)
	n, v = last(s, "1", 1, n, v)
	n, v = last(s, "2", 2, n, v)
	n, v = last(s, "3", 3, n, v)
	n, v = last(s, "4", 4, n, v)
	n, v = last(s, "5", 5, n, v)
	n, v = last(s, "6", 6, n, v)
	n, v = last(s, "7", 7, n, v)
	n, v = last(s, "8", 8, n, v)
	n, v = last(s, "9", 9, n, v)

	n, v = last(s, "one", 1, n, v)
	n, v = last(s, "two", 2, n, v)
	n, v = last(s, "three", 3, n, v)
	n, v = last(s, "four", 4, n, v)
	n, v = last(s, "five", 5, n, v)
	n, v = last(s, "six", 6, n, v)
	n, v = last(s, "seven", 7, n, v)
	n, v = last(s, "eight", 8, n, v)
	n, v = last(s, "nine", 9, n, v)

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
