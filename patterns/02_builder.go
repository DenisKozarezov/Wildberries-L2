package main

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

import "fmt"

const (
	Disel    = 0
	Fuel     = 1
	Electric = 2
)

type ICarBuilder interface {
	SetSeats(seatsCount int8)
	SetEngine()
	SetGPS()
	GetResult() *Product
}

type Product struct {
	Seats  int8
	Engine int8
	GPS    string
}

type TeslaBuilder struct {
	Seats  int8
	Engine int8
	GPS    string
}

func (b *TeslaBuilder) SetSeats(seatsCount int8) {
	b.Seats = seatsCount
}
func (b *TeslaBuilder) SetEngine() {
	b.Engine = Electric
}
func (b *TeslaBuilder) SetGPS() {
	b.GPS = "Tesla GPS v0.0.1"
}
func (b *TeslaBuilder) GetResult() *Product {
	return &Product{
		Seats:  b.Seats,
		Engine: b.Engine,
		GPS:    b.GPS,
	}
}

type FuelCarBuilder struct {
	Seats  int8
	Engine int8
	GPS    string
}

func (b *FuelCarBuilder) SetSeats(seatsCount int8) {
	b.Seats = seatsCount
}
func (b *FuelCarBuilder) SetEngine() {
	b.Engine = Fuel
}
func (b *FuelCarBuilder) SetGPS() {
	b.GPS = "Fuel GPS v0.0.5"
}
func (b *FuelCarBuilder) GetResult() *Product {
	return &Product{
		Seats:  b.Seats,
		Engine: b.Engine,
		GPS:    b.GPS,
	}
}

type Director struct {
	builder ICarBuilder
}

func NewDirector(b ICarBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) CreateCar(seatsCount int8) *Product {
	d.builder.SetSeats(seatsCount)
	d.builder.SetEngine()
	d.builder.SetGPS()
	return d.builder.GetResult()
}
func (d *Director) SetBuilder(b ICarBuilder) {
	d.builder = b
}

func main() {
	teslaBuilder := &TeslaBuilder{}
	director := NewDirector(teslaBuilder)
	tesla := director.CreateCar(2)
	fmt.Println(tesla)

	fmt.Println()

	fuelCarBuilder := &FuelCarBuilder{}
	director.SetBuilder(fuelCarBuilder)
	fuelCar := director.CreateCar(4)
	fmt.Println(fuelCar)
}
