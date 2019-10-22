//
package Facade

import (
	"strings"
)

// House ревлізує підсистему "House"
type House struct{}

// Build метод підсистеми
func (h *House) Build() string {
	return "Build house"
}

// Tree ревлізує підсистему "Tree"
type Tree struct{}

// Grow метод підсистеми
func (t *Tree) Grow() string {
	return "Tree grow"
}

// Child ревлізує підсистему "Child"
type Child struct{}

// Born метод підсистеми
func (c *Child) Born() string {
	return "Child born"
}

// Man структура, що надає уніфікований доступ підсистемі(Реалізація фасаду)
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// newMan створює нового Man
func newMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// MustDo повертає можливості, які надають підсистеми
func (m *Man) MustDo() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
		m.child.Born(),
	}

	return strings.Join(result, "\n")
}
