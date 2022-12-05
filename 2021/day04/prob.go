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

func processCards2(cards [][][]spot, num int, cardBingos []bool) int {
	lastCard := -1
	for c := 0; c < len(cards); c++ {
		if cardBingos[c] == false {
			for x := 0; x < len(cards[0]); x++ {
				for y := 0; y < len(cards[0][0]); y++ {
					if cards[c][x][y].value == num {
						cards[c][x][y].found = 1
						if bingo := checkForBingo(cards, c, x, y, num); bingo != 0 {
							// Might have more than one bingo for a num - choosing LAST card that bingoed
							cardBingos[c] = true
							lastCard = c
						}
					}
				}
			}
		}
	}
	return lastCard
}

func lastBingo(cardBingos []bool) bool {
	for _, e := range cardBingos {
		if e == false {
			return false
		}
	}
	return true
}

func process2(s []string) {
	// Load called numbers
	nums := getNum(strings.Split(s[0], ","))
	//showNums(nums)

	cardBingos := []bool{} // Lazy way of associating cards to found bingos

	cards := make([][][]spot, 0, 0)
	card := make([][]spot, 0, 0)

	// Load cards
	for _, e := range s[1:] {
		if len(e) == 0 && len(card) != 0 {
			cards = append(cards, card)
			cardBingos = append(cardBingos, false)
			card = make([][]spot, 0, 0)

		}
		if len(e) != 0 {
			card = append(card, createSpots(getNum(strings.Fields(e))))
		}
	}
	cards = append(cards, card)
	cardBingos = append(cardBingos, false)

	//showCards(cards)

	for _, num := range nums {
		lastCard := processCards2(cards, num, cardBingos)
		if lastBingo(cardBingos) {
			fmt.Printf("Process2: Bingo %d  -- Card %d -- Value %d\n", num, lastCard, num*getSumOfUnfoundSpotsForACard(cards, lastCard))
			return
		}
	}
}

type spot struct {
	value int
	found int
}

func showNums(nums []int) {
	fmt.Printf("%v\n", nums)
}

func showCards(cards [][][]spot) {
	fmt.Printf("------------\n")
	for c := 0; c < len(cards); c++ {
		for x := 0; x < len(cards[0]); x++ {
			for y := 0; y < len(cards[0][0]); y++ {
				fmt.Printf("[%02d] [%01d] ", cards[c][x][y].value, cards[c][x][y].found)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("------------\n")
	}
}

func showCard(cards [][][]spot, c int) {
	fmt.Printf("------------\n")
	for x := 0; x < len(cards[0]); x++ {
		for y := 0; y < len(cards[0][0]); y++ {
			fmt.Printf("[%02d] [%01d] ", cards[c][x][y].value, cards[c][x][y].found)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("------------\n")
}

func createSpots(nums []int) []spot {
	row := []spot{}
	for _, e := range nums {
		row = append(row, spot{value: e})
	}
	return row
}

func getSumOfUnfoundSpotsForACard(cards [][][]spot, c int) int {
	sum := 0
	for x := 0; x < len(cards[0]); x++ {
		for y := 0; y < len(cards[0][0]); y++ {
			if 0 == cards[c][x][y].found {
				sum += cards[c][x][y].value
			}
		}
	}
	return sum

}

func checkForBingo(cards [][][]spot, c int, x int, y int, num int) int {
	// Check horizontal bingo
	found := true
	for z := 0; found && z < len(cards[0][0]); z++ {
		if 0 == cards[c][x][z].found {
			found = false
		}
	}
	if found {
		return num * getSumOfUnfoundSpotsForACard(cards, c)
	}

	// Check vertical bingo
	found = true
	for z := 0; found && z < len(cards[0]); z++ {
		if 0 == cards[c][z][y].found {
			found = false
		}
	}
	if found {
		return num * getSumOfUnfoundSpotsForACard(cards, c)
	}
	return 0
}

func processCards(cards [][][]spot, num int) int {
	for c := 0; c < len(cards); c++ {
		for x := 0; x < len(cards[0]); x++ {
			for y := 0; y < len(cards[0][0]); y++ {
				if cards[c][x][y].value == num {
					cards[c][x][y].found = 1
					if bingo := checkForBingo(cards, c, x, y, num); bingo != 0 {
						return bingo
					}
				}
			}
		}
	}
	return 0
}

func process1(s []string) {
	// Load called numbers
	nums := getNum(strings.Split(s[0], ","))
	//showNums(nums)

	cards := make([][][]spot, 0, 0)
	card := make([][]spot, 0, 0)

	// Load cards
	for _, e := range s[1:] {
		if len(e) == 0 && len(card) != 0 {
			cards = append(cards, card)
			card = make([][]spot, 0, 0)

		}
		if len(e) != 0 {
			//fmt.Printf(" -- %02d\n", getNum(strings.Fields(e)))
			card = append(card, createSpots(getNum(strings.Fields(e))))
		}
	}
	cards = append(cards, card)

	//showCards(cards)

	for _, num := range nums {
		if bingo := processCards(cards, num); bingo != 0 {
			fmt.Printf("process1: Bingo %d  -- Value %d\n", num, bingo)
			return
		}

	}
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
