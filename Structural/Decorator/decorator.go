// Пакет Decorator є простеньким прикладом патерну Decorator
package Decorator

// Component забезпечує інтерфейс для декоратора та компонени.
type Component interface {
	Operation() string
}

// ConcreteComponent реалізує компоненту.
type ConcreteComponent struct {
}

// ConcreteDecorator реалізує декоратор.
type ConcreteDecorator struct {
	component Component
}

// Operation реалізація метода компоненти
func (c *ConcreteComponent) Operation() string {
	return "Я є компонента!"
}

// Operation це метод деокоратора, який огортає метод компоненти
func (c *ConcreteDecorator) Operation() string {
	return "<strong>" + c.component.Operation() + "</strong>"
}
