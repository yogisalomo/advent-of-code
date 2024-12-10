package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	input, err := os.Open("./input/day_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var firstList []int
	var secondList []int

	for scanner.Scan() {
		inputLine := scanner.Text()
		listItem := strings.Split(inputLine, "   ")
		if len(listItem) != 2 {
			log.Fatal("invalid parsed input", listItem, len(listItem))
		}

		firstItem, _ := strconv.Atoi(listItem[0])
		secondItem, _ := strconv.Atoi(listItem[1])
		firstList = append(firstList, firstItem)
		secondList = append(secondList, secondItem)
	}

	slices.Sort(firstList)
	slices.Sort(secondList)

	// PART 1
	totalDif := 0
	for i, firstListItem := range firstList {
		if firstListItem > secondList[i] {
			totalDif += firstListItem - secondList[i]
		} else {
			totalDif += secondList[i] - firstListItem
		}
	}
	fmt.Println("Total Dif: ", totalDif)

	// PART 2
	similarityScore := 0
	appearanceMap := make(map[int]int)
	for _, secondListItem := range secondList {
		if _, ok := appearanceMap[secondListItem]; !ok {
			appearanceMap[secondListItem] = 1
		} else {
			appearanceMap[secondListItem]++
		}
	}

	for _, firstListItem := range firstList {
		if _, ok := appearanceMap[firstListItem]; ok {
			similarityScore += firstListItem * appearanceMap[firstListItem]
		}
	}
	fmt.Println(similarityScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
