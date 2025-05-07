package main

import (
	"errors"
	"fmt"
)

type Animal struct {
	Name string
}

func (animal *Animal) Speak() error {
	return errors.New("not implemented")
}

type Dog Animal
type Cat Animal

func (dog *Dog) Speak() error {
	fmt.Println("Guf!")
	return nil
}
func (cat *Cat) Speak() error {
	fmt.Println("Myau!")
	return nil
}

type AnimalInterface interface {
	Speak() error
}

func main() {
	dog := Dog{}
	dog.Speak()

	cat := Cat{}
	cat.Speak()
}
