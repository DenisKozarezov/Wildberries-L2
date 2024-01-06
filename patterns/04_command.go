package main

/*
	Реализовать паттерн «команда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

import "fmt"

type ICommand interface {
	Execute()
	Undo()
}

type Character struct {
	PosX, PosY float32
	Health     int
	Level      int
}

type TranslateMoveCommand struct {
	Hero       *Character
	AddX, AddY float32
}

func (c *TranslateMoveCommand) Execute() {
	fmt.Println("Executing move command...")
	c.Hero.PosX += c.AddX
	c.Hero.PosY += c.AddY
	fmt.Printf("Current position: X: %f, Y: %f\n", c.Hero.PosX, c.Hero.PosY)
}
func (c *TranslateMoveCommand) Undo() {
	fmt.Println("Undo move command...")
	c.Hero.PosX -= c.AddX
	c.Hero.PosY -= c.AddY
	fmt.Printf("Reverted position: X: %f, Y: %f\n", c.Hero.PosX, c.Hero.PosY)
}
func NewMoveCommand(hero *Character, x, y float32) *TranslateMoveCommand {
	return &TranslateMoveCommand{Hero: hero, AddX: x, AddY: y}
}

type SetHealthCommand struct {
	Hero      *Character
	oldHealth int
	NewHealth int
}

func (c *SetHealthCommand) Execute() {
	fmt.Println("Executing set health command...")
	c.oldHealth = c.Hero.Health
	c.Hero.Health = c.NewHealth
	fmt.Println("Current health:", c.Hero.Health)
}
func (c *SetHealthCommand) Undo() {
	fmt.Println("Undo set health command...")
	c.Hero.Health = c.oldHealth
	fmt.Println("Reverted health:", c.Hero.Health)
}
func NewHealthCommand(hero *Character, health int) *SetHealthCommand {
	return &SetHealthCommand{Hero: hero, NewHealth: health}
}

type SetLevelCommand struct {
	Hero     *Character
	oldLevel int
	NewLevel int
}

func (c *SetLevelCommand) Execute() {
	fmt.Println("Executing set level command...")
	c.oldLevel = c.Hero.Level
	c.Hero.Level = c.NewLevel
	fmt.Println("Current level:", c.Hero.Level)
}
func (c *SetLevelCommand) Undo() {
	fmt.Println("Undo set level command...")
	c.Hero.Level = c.oldLevel
	fmt.Println("Reverted level:", c.Hero.Level)
}
func NewLevelCommand(hero *Character, level int) *SetLevelCommand {
	return &SetLevelCommand{Hero: hero, NewLevel: level}
}

func main() {
	hero := &Character{Health: 100, Level: 2}

	commands := []ICommand{
		NewMoveCommand(hero, 6, 7),
		NewHealthCommand(hero, 125),
		NewLevelCommand(hero, 3),
	}

	for _, command := range commands {
		command.Execute()
	}

	for _, command := range commands {
		command.Undo()
	}
}
