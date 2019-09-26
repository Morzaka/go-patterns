// Flyweight приклад патерну "Легковажний".
package Flyweight

// Flyweighter
type Flyweighter interface {
	GetName() string
	SetName(name string)
}

// FlyweightFactory реалізує фабрику.
// Якщо у пулі є відповідний примірник, то він повертає його.
type FlyweightFactory struct {
	pool map[int]Flyweighter
}

// GetFlyweight створює або повертає відповідний Flyweighter за станом.
func (f *FlyweightFactory) GetFlyweight(state int) Flyweighter {
	if f.pool == nil {
		f.pool = make(map[int]Flyweighter)
	}
	if _, ok := f.pool[state]; !ok {
		f.pool[state] = &ConcreteFlyweight{state: state}
	}
	return f.pool[state]
}

// ConcreteFlyweight реалізує інтерфейс Flyweighter.
type ConcreteFlyweight struct {
	state int
	name  string
}

// GetName повертає ім'я
func (c ConcreteFlyweight) GetName() string {
	return "My name: " + c.name
}

// SetName встановлює ім'я
func (c ConcreteFlyweight) SetName(name string) {
	c.name = name
}
