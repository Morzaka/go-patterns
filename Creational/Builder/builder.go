//builder пакет є прикладом патерну Builder.
package Builder

// Builder реалізує інтерфейс будівельника
type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

//Director реалізує менеджера
type Director struct {
	builder Builder
}

// Construct говорить будівельнику, що робити і в якому порядку.
func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

//ConcreteBuilder реалізує інтерфейс Builder.
type ConcreteBuilder struct {
	product *Product
}

//MakeHeader створює заголовок документа.
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.Content += "<header>" + str + "</header>"
}

//MakeBody будує тіло документа.
func (b *ConcreteBuilder) MakeBody(str string) {
	b.product.Content += "<article>" + str + "</article>"
}

//MakeFooter створює нижній колонтитул документа.
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.Content += "<footer>" + str + "</footer>"
}

//Product імплементує складний об'єкт.
type Product struct {
	Content string
}

//Show повертає готовий продукт(складний об'єкт)
func (p *Product) Show() string {
	return p.Content
}
