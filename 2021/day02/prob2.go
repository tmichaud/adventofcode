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
	process2(text)
}

// Create a slice of three-measurement sliding window
func process2(dir []string) {
	position := 0
	depth := 0
	aim := 0

	for _, v := range dir {
		dirs := strings.Fields(v)
		x := getInt(dirs[1])
		switch dirs[0] {
		case "forward":
			position += x
			depth += aim*x
			if depth < 0 {
				depth = 0 // Submarines can't fly
			}
		case "down":
			aim += x
		case "up":
			aim -= x
		default:
			fmt.Printf("ERROR DETECTED!  (%v)\n", dir[0])
			return
		}
		// The example test input walkthrough was off.  
		//fmt.Printf("[%v]   Position = (%d)  Depth = (%d)  Aim = (%v)  Multipled = (%d)\n", dirs, position, depth, aim, position*depth)
	}
	fmt.Printf("Position = (%d)  Depth = (%d)    Multipled = (%d)\n", position, depth, position*depth)
}

// Return the count of decreasing values
func process1(dir []string) {
	position := 0
	depth := 0

	for _, v := range dir {
		dirs := strings.Fields(v)
		x := getInt(dirs[1])
		switch dirs[0] {
		case "forward":
			position += x
		case "down":
			depth += x
		case "up":
			depth -= x
			if depth < 0 {
				depth = 0 // Submarines can't fly
			}
		default:
			fmt.Printf("ERROR DETECTED!  (%v)\n", dir[0])
			return
		}
	}
	fmt.Printf("Position = (%d)  Depth = (%d)    Multipled = (%d)\n", position, depth, position*depth)
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
