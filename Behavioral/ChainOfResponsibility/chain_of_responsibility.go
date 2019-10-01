// ChainOfResponsibility пакет є прикладом схеми відповідальності.
package ChainOfResponsibility

// Handler забезпечує інтерфейс обробника.
type Handler interface {
	SendRequest(message int) string
}

// ConcreteHandlerA реалізує обробник "A".
type ConcreteHandlerA struct {
	next Handler
}

// ConcreteHandlerA реалізує обробник "B".
type ConcreteHandlerB struct {
	next Handler
}

// ConcreteHandlerA реалізує обробник "C".
type ConcreteHandlerC struct {
	next Handler
}

// Реалізація SendRequest.
func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// Реалізація SendRequest.
func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// Реалізація SendRequest.
func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}