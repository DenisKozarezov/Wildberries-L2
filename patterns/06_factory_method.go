package main

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import "fmt"

const (
	GoblinType int8 = 0
	OrcType    int8 = 1
	HumanType  int8 = 2
)

type IEnemy interface {
	GetType() int8
	GetAgility() int
	GetStrength() int
	GetIntelligence() int
}

type EnemyBase struct {
	agility      int
	strength     int
	intelligence int
	enemyType    int8
}

func (e *EnemyBase) GetType() int8 {
	return e.enemyType
}
func (e *EnemyBase) GetAgility() int {
	return e.agility
}
func (e *EnemyBase) GetStrength() int {
	return e.strength
}
func (e *EnemyBase) GetIntelligence() int {
	return e.intelligence
}

type Goblin struct {
	EnemyBase
}

func NewGoblin() *Goblin {
	return &Goblin{
		EnemyBase{
			agility:      12,
			strength:     10,
			intelligence: 4,
			enemyType:    GoblinType,
		},
	}
}

type Orc struct {
	EnemyBase
}

func NewOrc() *Orc {
	return &Orc{
		EnemyBase{
			agility:      8,
			strength:     14,
			intelligence: 2,
			enemyType:    OrcType,
		},
	}
}

type Human struct {
	EnemyBase
}

func NewHuman() *Human {
	return &Human{
		EnemyBase{
			agility:      8,
			strength:     8,
			intelligence: 7,
			enemyType:    HumanType,
		},
	}
}

func SpawnEnemy(enemyType int8) (IEnemy, error) {
	switch enemyType {
	case GoblinType:
		return NewGoblin(), nil
	case OrcType:
		return NewOrc(), nil
	case HumanType:
		return NewHuman(), nil
	default:
		return nil, nil
	}
}

func main() {
	goblin, _ := SpawnEnemy(GoblinType)
	fmt.Println(goblin)

	orc, _ := SpawnEnemy(OrcType)
	fmt.Println(orc)

	human, _ := SpawnEnemy(HumanType)
	fmt.Println(human)
}
