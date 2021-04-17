package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGamer(gamerName string) {
	gamer.Games = append(gamer.Games, gamerName)
}

func main() {
	gamer := Gamer{Name: "Zelda"}

	gamer.AddGamer("Mario")
	gamer.AddGamer("Fifa 2020")
	gamer.AddGamer("Soccer 2020")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
