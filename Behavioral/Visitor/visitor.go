// Visitor -- це пакет, що реалізує патерн Візитера
package Visitor

// Visitor забезпечує інтерфейс відвідувачів.
type Visitor interface {
	VisitSushiBar(p *SushiBar) string
	VisitPizzeria(p *Pizzeria) string
	VisitBurgerBar(p *BurgerBar) string
}

// Place надає інтерфейс для місця, яке відвідувач повинен відвідати.
type Place interface {
	Accept(v Visitor) string
}

// People реалізує інтерфейс Visitor.
type People struct{}

func (*People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

func (*People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

func (*People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

// City реалізує колекцію місць для відвідування.
type City struct {
	places []Place
}

// Add додає Місце до колекції.
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

// Accept здійснює відвідування всіх місць у місті.
func (c *City) Accept(v Visitor) string {
	var result string
	for _, p := range c.places {
		result += p.Accept(v)
	}
	return result
}

// SushiBar реалізує інтерфейс Place.
type SushiBar struct{}

func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}

// Pizzeria реалізує інтерфейс Place.
type Pizzeria struct{}

func (s *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(s)
}

func (s *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}

// BurgerBar реалізує інтерфейс Place.
type BurgerBar struct{}

func (s *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(s)
}

func (s *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}

