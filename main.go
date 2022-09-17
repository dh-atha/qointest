package main

import (
	"fmt"
	"math/rand"
)

func main() {
	game(3, 4)
	game(3, 2)
}

func game(n int, m int) {
	resultMap := make(map[int]int, n)
	daduMap := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		resultMap[i] = 0
		daduMap[i] = m
	}

	players := n
	steps := 1
	for players > 1 {
		fmt.Println(fmt.Sprint("Permainan ke ", steps))
		currentDaduGame := make(map[int][]int)
		for i := 1; i <= n; i++ {
			var currentResult []int
			currentTotalDadu := daduMap[i]
			for j := 0; j < currentTotalDadu; j++ {
				currentDadu := lemparDadu()
				currentResult = append(currentResult, currentDadu)
				currentDaduGame[i] = append(currentDaduGame[i], currentDadu)
			}
			fmt.Println(fmt.Sprint("Pemain ", i, ": ", currentResult))
		}

		for key := range currentDaduGame {
			for _, dadu := range currentDaduGame[key] {
				if dadu == 1 {
					daduMap[key]--
					if key == n {
						daduMap[1]++
					} else {
						daduMap[key+1]++
					}
				} else if dadu == 6 {
					daduMap[key]--
					resultMap[key]++
				}
			}

			if daduMap[key] == 0 {
				players--
			}
		}

		fmt.Println()
		steps++
	}

	fmt.Print("hasil: ")
	fmt.Println(resultMap)
	var result int
	var highest int
	for i := 1; i <= n; i++ {
		if highest < resultMap[i] {
			highest = resultMap[i]
			result = i
		}
	}
	fmt.Println(fmt.Sprint("Pemenangnya adalah pemain ", result))
}

func lemparDadu() (result int) {
	for result == 0 {
		result = rand.Intn(7)
	}
	return
}
