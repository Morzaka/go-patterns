// Пакет Mediator є прикладом моделі Медіатора(посередника).
package Mediator

// Mediator реалізкація інтерфейсу медіатору.
type Mediator interface {
	Notify(msg string)
}

// ConcreteMediator реалізація посередника
type ConcreteMediator struct {
	*Farmer
	*Cannery
	*Shop
}

// Notify реалізація методу медіатора
func (m *ConcreteMediator) Notify(msg string) {
	if msg == "Farmer: Tomato complete..." {
		m.Cannery.AddMoney(-15000.00)
		m.Farmer.AddMoney(15000.00)
		m.Cannery.MakeKetchup(m.Farmer.GetTomato())
	} else if msg == "Cannery: Ketchup complete..." {
		m.Shop.AddMoney(-30000.00)
		m.Cannery.AddMoney(30000.00)
		m.Shop.SellKetchup(m.Cannery.GetKetchup())
	}
}

// ConnectColleagues з'єднує всіх колег.
func ConnectColleagues(farmer *Farmer, cannery *Cannery, shop *Shop) {
	mediator := &ConcreteMediator{
		Farmer:  farmer,
		Cannery: cannery,
		Shop:    shop,
	}

	mediator.Farmer.SetMediator(mediator)
	mediator.Cannery.SetMediator(mediator)
	mediator.Shop.SetMediator(mediator)
}

// Farmer -- це реалізація конткретного колеги, а саме Фермера що вирощує томати.
type Farmer struct {
	mediator Mediator
	tomato   int
	money    float64
}

func (f *Farmer) SetMediator(mediator Mediator) {
	f.mediator = mediator
}

func (f *Farmer) AddMoney(m float64) {
	f.money += m
}

func (f *Farmer) GrowTomato(tomato int) {
	f.tomato = tomato
	f.money -= 7500.00
	f.mediator.Notify("Farmer: Tomato complete...")
}

func (f *Farmer) GetTomato() int {
	return f.tomato
}

// Cannery -- це реалізація конткретного колеги, а саме Консервної фабрики що переробляє томати.
type Cannery struct {
	mediator Mediator
	ketchup  int
	money    float64
}

func (c *Cannery) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *Cannery) AddMoney(m float64) {
	c.money += m
}

func (c *Cannery) MakeKetchup(tomato int) {
	c.ketchup = tomato
	c.mediator.Notify("Cannery: Ketchup complete...")
}

func (c *Cannery) GetKetchup() int {
	return c.ketchup
}

// Shop -- це реалізація конткретного колеги, а саме Магазину що продає кетчуп.
type Shop struct {
	mediator Mediator
	money float64
}

func (s *Shop )SetMediator(mediator Mediator) {
	s.mediator = mediator
}

func (s *Shop ) AddMoney(m float64) {
	s.money += m
}

func (s *Shop ) SellKetchup(ketchup int) {
	s.money = float64(ketchup) * 54.75
}

func (s *Shop ) GetMoney() float64 {
	return s.money
}