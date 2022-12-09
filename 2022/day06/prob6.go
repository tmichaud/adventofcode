package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	text := getText()
	showText(text)
	findStart(text)
	fmt.Printf("---------------------------\n")
	findStart2(text)

}

func findStart(text []string) {
	for _, line := range text {
		found := false
		fmt.Printf("%v\n", line)
		d := strings.Split(line, "")
		for i := 3; i < len(d); i++ {
			test := true
			for y := i-3+0; test && y < i ; y++ {
				for z := y+1; test && z <= i; z++ {
					//fmt.Printf("Testing i (%v)  %v (%v) with %v (%v) \n", i, y, d[y], z, d[z])
					if d[y] == d[z] {
						test = false
					}
				}
			}
			if test {
				found = true
				fmt.Printf("Found! - (%v) \n", i+1)
				break
			}
		}
//		for i := 3; i < len(d); i++ {
//			fmt.Printf("(%v) (%v) (%v) \n", i, i-3, d[i-3:i-3+4])
//			if d[i-3+0] != d[i-3+1] && d[i-3+0] != d[i-3+2] && d[i-3+0] != d[i-3+3] && d[i-3+0] != d[i-3+4] &&
//				d[i-3+1] != d[i-3+2] && d[i-3+1] != d[i-3+3] && d[i-3+1] != d[i-3+4] &&
//				d[i-3+2] != d[i-3+3] && d[i-3+1] != d[i-3+4] { // <- bug!
//				fmt.Printf("Found! - (%v) \n", i+1)
//				found = true
//				break
//			}
//		}
		if !found {
			fmt.Printf("Did not find start!\n")
		}
	}
}

func findStart2(text []string) {
	for _, line := range text {
		found := false
		fmt.Printf("%v\n", line)
		d := strings.Split(line, "")
		for i := 14; i < len(d); i++ {
			test := true
			for y := i-13+0; test && y < i ; y++ {
				for z := y+1; test && z <= i; z++ {
					if d[y] == d[z] {
						test = false
					}
				}
			}
			if test {
				found = true
				fmt.Printf("Found! - (%v) \n", i+1)
				break
			}
		}
		if !found {
			fmt.Printf("Did not find start!\n")
		}
	}
}

func loadCrates(text []string, moveStart int, stacks int) [][]string {
	fmt.Printf("Loading Crates - stacks = (%v)\n", stacks)
	crates := make([][]string, stacks)
	for i := 0; i < stacks; i++ {
		crates[i] = make([]string, 0)
	}

	for _, v := range text[0:moveStart] {
		//fmt.Printf("%v (%v) --------", i, v)
		for x := 1; x <= stacks; x++ {
			s := string(v[(4*(x-1))+1]) // Get the crate value
			//fmt.Printf("(%v)(%v)(%v)", x, 4*(x-1)+1, s)
			if s != " " {
				//arr := crates[x-1]
				//arr = append(arr, s)
				crates[x-1] = prependString(crates[x-1], s)
			}
		}
		//fmt.Printf("\n")
	}

	return crates
}

func prependString(x []string, y string) []string {
	x = append(x, "")
	copy(x[1:], x)
	x[0] = y
	return x
}

func showCrates(crates [][]string) {
	fmt.Printf("-------------------------------------\n")
	for _, arr := range crates {
		for _, v := range arr {
			fmt.Printf("[%v]", v)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("-------------------------------------\n")
}

func showTopMostCrates(crates [][]string) {
	fmt.Printf("-------------------------------------\n")
	fmt.Printf("-------------------------------------\n")
	fmt.Printf("-------------------------------------\n")
	for _, arr := range crates {
		fmt.Printf("%v", arr[len(arr)-1])
	}
	fmt.Printf("\n")
}

func moveCrates(crates [][]string, num int, from int, to int) {
	fmt.Printf("Moving (%v) crates from stack (%v) to stack (%v)\n", num, from, to)
	crates[to] = append(crates[to], crates[from][len(crates[from])-num:]...)
	crates[from] = crates[from][0 : len(crates[from])-num]
}

func moveCratesOneAtATime(crates [][]string, num int, from int, to int) {
	if num == 0 {
		fmt.Printf("Finished\n")
		return
	}
	fmt.Printf("Moving (%v) crates from stack (%v) to stack (%v)\n", num, from, to)
	crates[to] = append(crates[to], crates[from][len(crates[from])-1:]...)
	crates[from] = crates[from][0 : len(crates[from])-1]
	moveCratesOneAtATime(crates, num-1, from, to)
}

func getCratesAndRules(text []string) {
	moveStart := 0
	for i, v := range text {
		if v == "" {
			moveStart = i
		}
	}
	space := regexp.MustCompile(`\s+`)
	r := space.ReplaceAllString(strings.TrimSpace(text[moveStart-1]), " ")
	s := strings.Split(r, " ")
	fmt.Printf("%v (%v)\n", len(s), s)

	crates := loadCrates(text, moveStart, len(s))
	showCrates(crates)

	fmt.Printf("Process moves\n")
	for _, v := range text[moveStart+1:] {
		fmt.Printf("%v\n", v)
		m := strings.Split(v, " ")
		num, _ := strconv.Atoi(m[1])
		from, _ := strconv.Atoi(m[3])
		to, _ := strconv.Atoi(m[5])
		moveCratesOneAtATime(crates, num, from-1, to-1)
		showCrates(crates)
	}

	showTopMostCrates(crates)
}

func getCratesAndRules2(text []string) {
	moveStart := 0
	for i, v := range text {
		if v == "" {
			moveStart = i
		}
	}
	space := regexp.MustCompile(`\s+`)
	r := space.ReplaceAllString(strings.TrimSpace(text[moveStart-1]), " ")
	s := strings.Split(r, " ")
	fmt.Printf("%v (%v)\n", len(s), s)

	crates := loadCrates(text, moveStart, len(s))
	showCrates(crates)

	fmt.Printf("Process moves\n")
	for _, v := range text[moveStart+1:] {
		fmt.Printf("%v\n", v)
		m := strings.Split(v, " ")
		num, _ := strconv.Atoi(m[1])
		from, _ := strconv.Atoi(m[3])
		to, _ := strconv.Atoi(m[5])
		moveCrates(crates, num, from-1, to-1)
		showCrates(crates)
	}

	showTopMostCrates(crates)
}

/********************************************************************************/
func getValue(r rune) int {
	if r >= 'a' {
		return int(r-'a') + 1
	}
	return int(r-'A') + 26 + 1
}

func processRange(text []string) {
	ret := 0
	for _, v := range text {
		arr := strings.SplitN(v, ",", 2)
		a := strings.SplitN(arr[0], "-", 2)
		b := strings.SplitN(arr[1], "-", 2)
		x1, _ := strconv.Atoi(a[0])
		x2, _ := strconv.Atoi(a[1])
		y1, _ := strconv.Atoi(b[0])
		y2, _ := strconv.Atoi(b[1])
		fmt.Printf("x : (%v)-(%v) y : (%v)-(%v) : ", x1, x2, y1, y2)
		if (x1 <= y1) && (x2 >= y2) || (y1 <= x1) && (y2 >= x2) {
			fmt.Printf(" Contains")
			ret++
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Final (%v)\n", ret)
}

func test(x1 int, x2 int, y1 int, y2 int) bool {
	if (x1 >= y1) && (x1 <= y2) {
		fmt.Printf("(@1)")
		return true
	}
	if (x2 >= y1) && (x2 <= y2) {
		fmt.Printf("(@2)")
		return true
	}
	if (y1 >= x1) && (y1 <= x2) {
		fmt.Printf("(@3)")
		return true
	}
	if (y2 >= x1) && (y2 <= x2) {
		fmt.Printf("(@4)")
		return true
	}
	return false
}

func processRange2(text []string) {
	ret := 0
	for _, v := range text {
		arr := strings.SplitN(v, ",", 2)
		a := strings.SplitN(arr[0], "-", 2)
		b := strings.SplitN(arr[1], "-", 2)
		x1, _ := strconv.Atoi(a[0])
		x2, _ := strconv.Atoi(a[1])
		y1, _ := strconv.Atoi(b[0])
		y2, _ := strconv.Atoi(b[1])
		fmt.Printf("x : (%v)-(%v) y : (%v)-(%v) : ", x1, x2, y1, y2)

		if test(x1, x2, y1, y2) {
			fmt.Printf(" Contains")
			ret++
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Final (%v)\n", ret)
}

func processRucksack2(text []string) {
	ret := 0
	for i := 0; i+3 <= len(text); i += 3 {
		c1 := text[i]
		c2 := text[i+1]
		c3 := text[i+2]

		fmt.Printf("%d   (%v) (%v)  (%v) \n", i, c1, c2, c3)
		for _, r := range c1 {
			if strings.Contains(c2, string(r)) && strings.Contains(c3, string(r)) {
				ret += getValue(r)
				fmt.Printf("(%v) found value (%v)\n", string(r), getValue(r))
				break
			}
		}
	}
	fmt.Printf("Final (%v)\n", ret)
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
