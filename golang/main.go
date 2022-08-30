package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run *.go randomseed1 randomseed2")
		fmt.Println("Example: go run *.go 314 15")
		os.Exit(0)
	}
	seed1, _ := strconv.ParseInt(os.Args[0], 10, 64)
	seed2, _ := strconv.ParseInt(os.Args[1], 10, 64)
	player1 := NewPlayer("Taro", NewWinningStrategy(seed1))
	player2 := NewPlayer("Hana", NewProbStrategy(seed2))
	for i := 0; i < 10000; i++ {
		nextHand1 := player1.NextHand()
		nextHand2 := player2.NextHand()
		if nextHand1.IsStrongerThan(nextHand2) {
			fmt.Printf("Winner:%v\n", player1)
			player1.Win()
			player2.Lose()
		} else if nextHand2.IsStrongerThan(nextHand1) {
			fmt.Printf("Winner:%v\n", player2)
			player1.Lose()
			player2.Win()
		} else {
			fmt.Println("Even...")
			player1.Even()
			player2.Even()
		}
	}
	fmt.Println("Total result:")
	fmt.Println(player1)
	fmt.Println(player2)
}
