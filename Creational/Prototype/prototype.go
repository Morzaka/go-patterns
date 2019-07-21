// package Prototype - приклад шаблону Prototype.
package Prototype

// Прототип надає інтерфейс клонування.
type Prototyper interface {
	Clone() Prototyper
	GetName() string
}

// ConcreteProduct реалізує продукт "А"
type ConcreteProduct struct {
	name string
}

// NewConcreteProduct - конструктор прототипів.
func NewConcreteProduct(name string) Prototyper {
	return &ConcreteProduct{
		name : name,
	}
}

// GetName повертає ім'я продукту.
func (p *ConcreteProduct) GetName() string {
	return p.name
}

// Clone повертає клонований об'єкт.
func (p *ConcreteProduct) Clone() Prototyper {
	return &ConcreteProduct{p.name}
}