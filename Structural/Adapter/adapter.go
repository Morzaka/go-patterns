package adapter

// Target абезпечує інтерфейс, з яким система повинна працювати.
type Target interface {
	Request() string
}

// Adaptee впроваджує систему для адаптації.
type Adaptee struct {
}

// NewAdapter конструктор адаптера.
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// SpecificRequest реалізація
func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Adapter реалізує інтерфейс Target і є адаптером.
type Adapter struct {
	*Adaptee
}

// Request є адаптивним методом.
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}
