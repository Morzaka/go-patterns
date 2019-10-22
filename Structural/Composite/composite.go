// Composite пакет є прикладом композиційного патерна.
package Composite

// Component забезпечує інтерфейс для гілок та листя дерева.
type Component interface {
	Add(child Component)
	Name() string
	Child() []Component
	Print(prefix string) string
}

// Directory реалізує гілки дерева.
type Directory struct {
	name   string
	childs []Component
}

// Add додає елемент до гілки дерева.
func (d *Directory) Add(child Component) {
	d.childs = append(d.childs, child)
}

// Name повертає ім'я компонента.
func (d *Directory) Name() string {
	return d.name
}

// Child повертає дочірні елементи.
func (d *Directory) Child() []Component {
	return d.childs
}

// Print повертає гілку в рядковому поданні.
func (d *Directory) Print(prefix string) string {
	result := prefix + "/" + d.Name() + "\n"
	for _, val := range d.Child() {
		result += val.Print(prefix + "/" + d.Name())
	}
	return result
}

// File реалізує листя дерева.
type File struct {
	name string
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Child() []Component {
	return []Component{}
}

func (f *File) Print(prefix string) string {
	return prefix + "/" + f.Name() + "\n"
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
	}
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}
