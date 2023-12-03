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
	testFunctions(text)
	//process1(text)
	process2(text)
}

// We completely ignore runes actually
func getByte(text []string, x int, y int) byte {
	line := text[x]
	return line[y]
}

func isDot(b byte) bool {
	if b == 46 {
		return true
	}
	return false
}

func isNumber(b byte) bool {
	if b >= 48 && b <= 57 {
		return true
	}
	return false
}

func isGear(b byte) bool {
	if b == 42 {
		return true
	}
	return false
}

func isSymbol(b byte) bool {
	if isDot(b) {
		return false
	}
	if isNumber(b) {
		return false
	}
	return true
}

func testFunctions(text []string) {
	showText(text)
	x, y := 0, 1
	fmt.Printf("Expected: 0,1 (.) isDot(T) isNumber(F) isSymbol(F) --- Actual %d,%d (%v) isDot(%v) isNumber(%v) isSymbol(%v) \n", x, y, string(getByte(text, x, y)), isDot(getByte(text, x, y)), isNumber(getByte(text, x, y)), isSymbol(getByte(text, x, y)))

	x, y = 1, 0
	fmt.Printf("Expected: 1,0 (4) isDot(F) isNumber(T) isSymbol(F) --- Actual %d,%d (%v) isDot(%v) isNumber(%v) isSymbol(%v) \n", x, y, string(getByte(text, x, y)), isDot(getByte(text, x, y)), isNumber(getByte(text, x, y)), isSymbol(getByte(text, x, y)))

	x, y = 2, 3
	fmt.Printf("Expected: 2,3 (*) isDot(F) isNumber(F) isSymbol(T) --- Actual %d,%d (%v) isDot(%v) isNumber(%v) isSymbol(%v) \n", x, y, string(getByte(text, x, y)), isDot(getByte(text, x, y)), isNumber(getByte(text, x, y)), isSymbol(getByte(text, x, y)))

	x, y = 1, 3
	fmt.Printf("Expected: 1,3 (7) isDot(F) isNumber(T) isSymbol(F) --- Actual %d,%d (%v) isDot(%v) isNumber(%v) isSymbol(%v) \n", x, y, string(getByte(text, x, y)), isDot(getByte(text, x, y)), isNumber(getByte(text, x, y)), isSymbol(getByte(text, x, y)))
	fmt.Printf("testForSymbol: Yes --- Actual %v\n", testForSymbol(text, 1, 3))


	x,y = 1,0
	fmt.Printf("Expect: 467 - Actual %d\n", getGearNum(text, x, y))

	x,y = 1,1
	fmt.Printf("Expect: 467 - Actual %d\n", getGearNum(text, x, y))

	x,y = 1,2
	fmt.Printf("Expect: 467 - Actual %d\n", getGearNum(text, x, y))

	x,y = 1,3
	fmt.Printf("Expect: ??? - Actual %d\n", getGearNum(text, x, y))

	fmt.Printf("---------------------------\n")
	fmt.Printf("---------------------------\n")
	fmt.Printf("---------------------------\n")
}

func getNum(s string, y0 int, y1 int) int {
	n, _ := strconv.Atoi(s[y0:y1])
	return n
}

func testForSymbol(text []string, x int, y int) bool {
	// Test upper left, upper and upper right
	if x != 0 && y != 0 {
		if isSymbol(getByte(text, x-1, y-1)) {
			return true
		}
	}
	if x != 0 {
		if isSymbol(getByte(text, x-1, y)) {
			return true
		}
	}
	if x != 0 && y != len(text[x-1])-1 {
		if isSymbol(getByte(text, x-1, y+1)) {
			return true
		}
	}

	// Test left and right
	if y != 0 {
		if isSymbol(getByte(text, x, y-1)) {
			return true
		}
	}

	if y != len(text[x])-1 {
		if isSymbol(getByte(text, x, y+1)) {
			return true
		}
	}

	// Test lower left, lower and lower right
	if x != len(text)-1 && y != 0 {
		if isSymbol(getByte(text, x+1, y-1)) {
			return true
		}
	}

	if x != len(text)-1 {
		if isSymbol(getByte(text, x+1, y)) {
			return true
		}
	}

	if x != len(text)-1 && y != len(text[x-1])-1 {
		if isSymbol(getByte(text, x+1, y+1)) {
			return true
		}
	}
	return false
}

func testForSymbolRange(text []string, x int, y0 int, y1 int) bool {
	for i := y0; i < y1; i++ {
		if testForSymbol(text, x, i) {
			return true
		}
	}
	return false
}

// Walk thorugh the 1:len(text)-1 lines for numbers.  Ignore the first and last lines used as . buffers.
// If we find the start of a number (flag), record x1,y1
// Continue walking through numbers to end of number (x1,y2)
// for each x1,y1 ... x1,y2 - test if a symbol is next to it.
//  --  if it is, convert the number
//  --  and add it to the sum
func process1(text []string) {
	sum := 0

	for i, line := range text[1 : len(text)-1] {
		fmt.Printf("%d %v \n", i, line)
		soan := false
		x1 := i + 1
		y1 := 0
		y2 := 0
		for j := 0; j < len(line); j++ {
			if isNumber(line[j]) {
				if !soan {
					soan = true
					y1 = j
				}
			} else {
				if soan {
					soan = false
					y2 = j
					symbolExists := testForSymbolRange(text, x1, y1, y2)
					num := getNum(line, y1, y2)
					fmt.Printf(" -- Num(%s) Num(%d) SymbolExists(%v) \n", string(line[y1:y2]), num, symbolExists)
					if symbolExists {
						sum = sum + num
					}
				}
			}
		}
	}
	fmt.Printf("\n Sum is (%d)\n", sum)
}

func getGearNum(text []string, x int, y int) int {
	// Find start of number
	y0 := y
	for y0 != 0 && isNumber(getByte(text, x, y0-1)) {
		y0 -= 1
	}

	// Find end of number
	y1 := y
	for y1 != len(text[x])-1 && isNumber(getByte(text, x, y1+1)) {
		y1 += 1
	}

	//convert and return number
	line := text[x]
	//fmt.Printf("Line : (%v)\n", line)
	//fmt.Printf(" -- %d:%d (%v)\n", y0, y1, line[y0:y1+1])

	num,_ := strconv.Atoi(line[y0:y1+1])
	return num
}

// A gear is a * symbol adjancent to exactly 2 part numbers
func getGearNums(text []string, x int, y int) int {
	ret := []int{}

	// Check left
	if y != 0 {
		if isNumber(getByte(text, x, y-1)) {
			ret = append(ret, getGearNum(text, x, y-1))
		}
	}

	// Check right
	if y != len(text[x])-1 {
		if isNumber(getByte(text, x, y+1)) {
			ret = append(ret, getGearNum(text, x, y+1))
		}
	}

	// Check up
	if x != 0 {
		if isNumber(getByte(text, x-1, y)) {
			ret = append(ret, getGearNum(text, x-1, y))
		} else {
			// Now check up & left and up & right
			if y != 0 {
				if isNumber(getByte(text, x-1, y-1)) {
					ret = append(ret, getGearNum(text, x-1, y-1))
				}
			}
			if y != len(text[x])-1 {
				if isNumber(getByte(text, x-1, y+1)) {
					ret = append(ret, getGearNum(text, x-1, y+1))
				}
			}
		}
	}

	// Check down
	if x != len(text)-1 {
		if isNumber(getByte(text, x+1, y)) {
			ret = append(ret, getGearNum(text, x+1, y))
		} else {
			// Now check down & left and down & right
			if y != 0 {
				if isNumber(getByte(text, x+1, y-1)) {
					ret = append(ret, getGearNum(text, x+1, y-1))
				}
			}
			if y != len(text[x])-1 {
				if isNumber(getByte(text, x+1, y+1)) {
					ret = append(ret, getGearNum(text, x+1, y+1))
				}
			}
		}
	}

	fmt.Printf("ret size is (%d) -- (%v) \n", len(ret), ret)
	if len(ret) == 2 {
		return ret[0]*ret[1]
	}
	return 0
}

// Walk thorugh the 1:len(text)-1 lines for gears.  Ignore the first and last lines used as . buffers.
// If we find a gear, look for 2 numbers adjacent to it
// -- if we find the nubmers, figure the start/end of the numbers and then convert them into numbers.
//    multiply the numbers together
//    return the sum of the multiplied numbers
func process2(text []string) {
	sum := 0

	for i, line := range text[1 : len(text)-1] {
		fmt.Printf("%d %v\n", i, line)
		for j := 0; j < len(line); j++ {
			if isGear(line[j]) {
				fmt.Printf(" -- Gear found at %d,%d\n", i+1, j)
				sum += getGearNums(text, i+1, j)
			}
		}
	}

	fmt.Printf("\n Sum is (%d)\n", sum)
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

//func getNum(s []string) []int {
//	ret := []int{}
//	for _, x := range s {
//		i, _ := strconv.Atoi(x)
//		ret = append(ret, i)
//	}
//	return ret
//}

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
