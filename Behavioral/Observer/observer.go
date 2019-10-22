// Observer приклад патерну спостерігача.
// Метод проштовхування.
package Observer

// Publisher interface.
type Publisher interface {
	Attach(observer Observer)
	SetState(state string)
	Notify()
}

// Observer це реалізація інтерфесу спостерігача(підписника)
type Observer interface {
	Update(state string)
}

// ConcretePublisher реалізація конкретного Publisher(Видавця)
type ConcretePublisher struct {
	observers []Observer
	state     string
}

// NewPublisher так званий конструктор повертає нового Видавця
func NewPublisher() Publisher {
	return &ConcretePublisher{}
}

// Attach додає нового підписника(спостерігача)
func (p *ConcretePublisher) Attach(observer Observer) {
	p.observers = append(p.observers, observer)
}

// SetState задає новий стан видавця
func (p *ConcretePublisher) SetState(state string) {
	p.state = state
}

// Notify повідомляє всіх підписників(спостерігаців) про зміну стану
// Метод проштовхування
func (p *ConcretePublisher) Notify() {
	for _, observer := range p.observers {
		observer.Update(p.state)
	}
}

// ConcreteObserver реалізація конкретного підписника(спостерігача)
type ConcreteObserver struct {
	state string
}

// Update встановлює новий стан для підписника
func (o *ConcreteObserver) Update(state string) {
	o.state = state
}
