// Package Bridge є прикладом шаблону Міст.
package Bridge

// Carer забезпечує інтерфейс автомобіля.
type Carer interface {
	Rase() string
}

// Enginer забезпечує інтерфейс двигуна.
type Enginer interface {
	GetSound() string
}

// Car реалізація
type Car struct {
	engine Enginer
}

// NewCar конструктор нового авта
func NewCar(engine Enginer) Carer{
	return &Car{
		engine: engine,
	}
}

// Rase реалізація
func (c *Car) Rase() string {
	return c.engine.GetSound()
}

// EngineSuzuki реалізація двигуна Сузукі
type EngineSuzuki struct {
}

// GetSound звук двигуна певного автомобіля.
func (e *EngineSuzuki) GetSound() string {
	return "Suuuzzukkiii"
}

// EngineHonda реалізація двигуна Хонда
type EngineHonda struct {
}

// GetSound звук двигуна певного автомобіля.
func (e *EngineHonda) GetSound() string {
	return "HhoooNnnnDddaaaa"
}

