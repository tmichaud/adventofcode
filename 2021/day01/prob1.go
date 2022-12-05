package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	text := getText()
	//showText(text)
	ints := getNum(text)
	//showInt(ints)
	x := processInts(ints)
	fmt.Printf("Count is %d\n", x)
	x = processInts(process2(ints))
	fmt.Printf("Count2 is %d\n", x)
}

// Create a slice of three-measurement sliding window
func process2(ints []int) []int  {
	r := []int{}
	a,b,c := ints[0], ints[1], ints[2]
	r = append(r, a+b+c)

	for _,e := range ints[3:len(ints)] {
		a,b,c = b,c,e
		r = append(r, a+b+c)
	}
	return r
}

// Return the count of decreasing values
func processInts(ints []int) int {
	count := 0
	last := ints[0]
	for _,v := range ints[1:len(ints)] {
		if v > last {
			count++
		}
		last = v
	}
	return count
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
	for _,e := range ints {
		fmt.Printf("%d\n", e)
	}
}

func getNum(s []string) []int {
	ret := []int{}
	for _,x := range s {
		i,_ := strconv.Atoi(x)
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
