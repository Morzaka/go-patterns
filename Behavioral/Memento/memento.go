// Memento пакет - приклад патерну Memento.
package Memento

// Originator реалізує стан господаря.
type Originator struct {
	State string
}

// NewMemento повертає стан зберігача
func (o *Originator) NewMemento() *Memento {
	return &Memento{state: o.State}
}

// SetMemento задає (сатрий) стан
func (o *Originator) SetMemento(memento *Memento) {
	o.State = memento.GetState()
}

// Memento реалізує сховище для стану Originator
type Memento struct {
	state string
}

// GetState повертає стан.
func (m *Memento) GetState() string {
	return m.state
}

// Caretaker зберігає Мементо до тих пір, поки він не знадобиться Originator.
type Caretaker struct {
	Memento *Memento
}