package main

import "fmt"

type Player struct {
	health, maxHealth int
	energy, maxEnergy int
	name              string
}

func (player *Player) setHealth(newHealth int) {
	player.health = newHealth
}

func (player *Player) setEnergy(newEnergy int) {
	player.energy = newEnergy
}

func main() {
	player1 := Player{
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 100,
		name:      "Akash",
	}
	fmt.Println(player1)

	player1.setHealth(89)
	player1.setEnergy(60)
	fmt.Println(player1)
}
