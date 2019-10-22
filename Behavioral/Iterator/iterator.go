// Iterator пакет є прикладом шаблону ітератора.
package Iterator

// Iterator забезпечує інтерфейс ітератора.
type Iterator interface {
	Index() int
	Value() interface{}
	Has() bool
	Next()
	Prev()
	Reset()
	End()
}

// Aggregate забезпечує інтерфейс колекції.
type Aggregate interface {
	Iterator() Iterator
}

// Book реалізує предмет колекції.
type Book struct {
	Name string
}

// BookShelf реалізує інтерфейс Aggregate.
type BookShelf struct {
	Books []*Book
}

// Iterator творює та повертає ітератор над колекцією.
func (b *BookShelf) Iterator() Iterator {
	return &BookIterator{shelf: b}
}

// Add додає елемент до колекції.
func (b *BookShelf) Add(book *Book) {
	b.Books = append(b.Books, book)
}

// BookIterator реалізує інтерфейс Iterator.
type BookIterator struct {
	shelf    *BookShelf
	index    int
	internal int
}

// Index повертає поточний індекс.
func (i *BookIterator) Index() int {
	return i.index
}

// Value повертає поточне значення.
func (i *BookIterator) Value() interface{} {
	return i.shelf.Books[i.index]
}

// Has перевіряє на наявність в колекції
func (i *BookIterator) Has() bool {
	if i.internal < 0 || i.internal >= len(i.shelf.Books) {
		return false
	}
	return true
}

// Next перехід до наступного елементу.
func (i *BookIterator) Next() {
	i.internal++
	if i.Has() {
		i.index++
	}
}

// Prev переходить до попереднього елементу.
func (i *BookIterator) Prev() {
	i.internal--
	if i.Has() {
		i.index--
	}
}

// Reset скидає ітератор.
func (i *BookIterator) Reset() {
	i.index = 0
	i.internal = 0
}

// End переходить до останнього пункту.
func (i *BookIterator) End() {
	i.index = len(i.shelf.Books) - 1
	i.internal = i.index
}
