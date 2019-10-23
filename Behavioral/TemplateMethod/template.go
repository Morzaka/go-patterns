package TemplateMethod

type Quoter interface {
	Open() string
	Close() string
}

type Quotes struct {
	Quoter
}

func NewQuotes(qt Quoter) *Quotes {
	return &Quotes{qt}
}

func (q *Quotes) Quotes(str string) string {
	return q.Open() + str + q.Close()
}

type FrenchQuotes struct{}

func (f *FrenchQuotes) Open() string {
	return "«"
}

func (f *FrenchQuotes) Close() string {
	return "»"
}

type GermanQuotes struct{}

func (g *GermanQuotes) Open() string {
	return "„"
}

func (g *GermanQuotes) Close() string {
	return "“"
}
