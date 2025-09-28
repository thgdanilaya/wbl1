package main

import "fmt"

type Pig struct {
}

func (p *Pig) Grunt() {
	fmt.Println("Хрюкаем хрю хрю епта")
}

type Dog struct {
}

func (d *Dog) Bark() {
	fmt.Println("Гавкаем на хрюкающую свинью гав гав чики пики")
}

type AnimalTranslator interface {
	Talk()
}

type PigAdapter struct {
	*Pig
}

func (pigAdapter *PigAdapter) Talk() {
	pigAdapter.Pig.Grunt()
}

func NewPigAdapter(pig *Pig) *PigAdapter {
	return &PigAdapter{pig}
}

type DogAdapter struct {
	*Dog
}

func (dogAdapter *DogAdapter) Talk() {
	dogAdapter.Dog.Bark()
}

func NewDogAdapter(dog *Dog) *DogAdapter {
	return &DogAdapter{dog}
}

func main() {
	animals := []AnimalTranslator{NewPigAdapter(&Pig{}), NewDogAdapter(&Dog{})}
	for _, animal := range animals {
		animal.Talk()
	}
}
