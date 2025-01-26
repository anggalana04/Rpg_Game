package main

import (
	"Rpg_Game/logic"
	"fmt"
)

func main() {
	logic.ClearScreen()
	fmt.Println("==============")
	logic.Greet()
	logic.ShowMainMenu()
	fmt.Println("Press Enter To Play The Game")
	fmt.Scanln()
	logic.ClearScreen()
	logic.AddCharacter()
}