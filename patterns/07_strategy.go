package main

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

import "fmt"

type Vector3 struct {
	X, Y, Z float32
}

type Path struct {
	Points []Vector3
}

const (
	AStarPathfinding      = 0
	BruteForcePathfinding = 1
)

type IPathfindingStrategy interface {
	FindPath(point1, point2 Vector3) Path
}

type AStarPathfinder struct {
}

func (p *AStarPathfinder) FindPath(point1, point2 Vector3) Path {
	// Making the whole stuff for A* algorithm...
	return Path{}
}

type BruteForcePathfinder struct {
}

func (p *BruteForcePathfinder) FindPath(point1, point2 Vector3) Path {
	// Check all points and try to find out the way from the source to the destination...
	return Path{}
}

type LabyrinthNavigator struct {
	pathfinder IPathfindingStrategy
}

func InitNavigator(pathfinder IPathfindingStrategy) *LabyrinthNavigator {
	return &LabyrinthNavigator{
		pathfinder: pathfinder,
	}
}

func (n *LabyrinthNavigator) SwitchPathfinder(pathfinder IPathfindingStrategy) {
	n.pathfinder = pathfinder
}
func (n *LabyrinthNavigator) BuildPath(point1, point2 Vector3) Path {
	return n.pathfinder.FindPath(point1, point2)
}

func main() {
	pathfinder := &AStarPathfinder{}

	navigator := InitNavigator(pathfinder)

	path := navigator.BuildPath(Vector3{0, 1, 5}, Vector3{2, 3, 7})

	fmt.Println(path)
}
