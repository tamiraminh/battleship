package main

import (
	"battleships/model"
	"fmt"
)

func main() {
	var m, s, t int
	fmt.Scan(&m)
	fmt.Scan(&s)

	var p1ShipPositionStr, p2ShipPositionStr string
	fmt.Scan(&p1ShipPositionStr)
	fmt.Scan(&p2ShipPositionStr)

	fmt.Scan(&t)

	var p1MissilePositionStr, p2MissilePositionStr string
	fmt.Scan(&p1MissilePositionStr)
	fmt.Scan(&p2MissilePositionStr)

	P1 := model.NewPlayer("P1", s, p1ShipPositionStr, m, t)
	P2 := model.NewPlayer("P2", s, p2ShipPositionStr, m, t)

	P1.AttackPlayer(p1MissilePositionStr, P2)
	P2.AttackPlayer(p2MissilePositionStr, P1)

	// print Map
	fmt.Println()
	fmt.Println("Player 1")
	P1.Map.ShowMap()
	fmt.Println("Player 2")
	P2.Map.ShowMap()

	fmt.Println()
	fmt.Printf("P1: %d\n", s-P1.ActiveShips)
	fmt.Printf("P2: %d\n", s-P2.ActiveShips)
	if P1.ActiveShips > P2.ActiveShips {
		fmt.Println("Player 1 wins")
	} else if P1.ActiveShips < P2.ActiveShips {
		fmt.Println("Player 2 wins")
	} else {
		fmt.Println("It is a draw")
	}

}
