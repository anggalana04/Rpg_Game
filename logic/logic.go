package logic

import (
	"fmt"
)

type Character struct {
	Name string
	Role string
}

var characters []Character


func AddCharacter() {
	fmt.Println("Please give your character a name and roler")
	var character_name string
	var character_role string
	fmt.Println("What is the name of your character?")
	fmt.Scanln(&character_name)
	fmt.Println("What is the role of your character?")
	fmt.Scanln(&character_role)
	newCharacter := Character{
		Name: character_name, 
		Role: character_role,
	}
	characters = append(characters, newCharacter)
	fmt.Println(characters)

}



func Greet() {
	fmt.Println("Hello World")
}

func ShowMainMenu() {
	var name string
	var role string

	fmt.Println("Hello Welcome to the game!")
	fmt.Println("What is thou name, Fellow?")
	fmt.Scanln(&name)
	fmt.Println("ahh ", name, "!")
	fmt.Println("Which Role Do u Choose ?")
	fmt.Println(`
	1. Mage
	2. Fighter 
	3. Assasin
	`)
	fmt.Scanln(&role)
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

