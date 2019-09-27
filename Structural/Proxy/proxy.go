// Пакет Proxy - приклад шаблону Adapter.
package Proxy

// Subject забезпечує інтерфейс для реального предмета та його сурогату.
type Subject interface {
	Send() string
}

// Proxy реалізує сурогат.
type Proxy struct {
	realSubject Subject
}

// RealSubject
type RealSubject struct {
}

// Send надсилає повідомлення
func (p *Proxy) Send() string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}

	return "<strong>" + p.realSubject.Send() + "</strong>"
}

// Send надсилає повідомлення
func (s *RealSubject) Send() string {
	return "I'll be back!"
}