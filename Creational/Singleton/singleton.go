// package singleton є прикладом шаблону Singleton.
package singleton

import (
	"sync"
)

// Реалізація Singleton.
type Singleton struct {
}

var (
	instance *Singleton
	once sync.Once
)

// GetInstance повертає синглет
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}