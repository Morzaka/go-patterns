//Пакет factory_method є прикладом патерну FactoryMethod.
package FactoryMethod

import (
	"log"
)

// Creater реалізує заводський інтерфейс.
type Creater interface {
	CreateProduct(action string) Producter
}

// Producter реалізує інтерфейс продукту.
// Всі продукти, що повертаються фабрикою, повинні реалізувати єдиний інтерфейс.
type Producter interface {
	Use() string
}

// ConcreteCreater імплементує інтерфейс Creater.
type ConcreteCreater struct {
}

// NewCreater - це конструктор ConcreteCreater.
func NewCreater() Creater {
	return &ConcreteCreater{}
}

// CreateProduct є заводським методом
func (p *ConcreteCreater) CreateProduct(action string) Producter {
	var product Producter

	switch action {
	case "A":
		product = &ConcreteProductA{action}
	case "B":
		product = &ConcreteProductB{action}
	case "C":
		product = &ConcreteProductC{action}
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

// ConcreteProductA реалізує продукт "А"
type ConcreteProductA struct {
	action string
}

//Use повертає дії продукту
func (p *ConcreteProductA) Use() string {
	return p.action
}

// ConcreteProductB реалізує продукт "B"
type ConcreteProductB struct {
	action string
}

//Use повертає дії продукту
func (p *ConcreteProductB) Use() string {
	return p.action
}

// ConcreteProductС реалізує продукт "C"
type ConcreteProductC struct {
	action string
}

//Use повертає дії продукту
func (p *ConcreteProductC) Use() string {
	return p.action
}
