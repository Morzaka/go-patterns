//Пакет abstract_factory є прикладом патерна абстрактної фабрики.
package AbstractFactory

//AbstractFactory -- інтерфейс для створення сімейства пов'язаних об'єктів.
type AbstractFactory interface {
	CreateWater(volume float64) AbstractWater
	CreateBotle(voleme float64) AbstractBotle
}

// AbstractWater інтерфейс Води.
type AbstractWater interface {
	GetVolume() float64
}

// AbstractBotle інтерфейс Пляшки.
type AbstractBotle interface {
	PourWater(water AbstractWater) //Bottle iteract with a water.
	GetBottleVolume() float64
	GetWaterVolume() float64
}

// CocaColaFactory реалізує інтерфейс AbstractFactory.
type CocaColaFactory struct {
}

// NewCocaColaFactory - конструктор CocaColaFactory.
func NewCocaColaFactory() AbstractFactory {
	return &CocaColaFactory{}
}

// CreateWater реалізаця методу інтерфеса AbstractFactory.
func (f *CocaColaFactory) CreateWater(volume float64) AbstractWater {
	return &CocaColaWater{volume: volume}
}

// CreateBottle реалізаця методу інтерфеса AbstractFactory.
func (f *CocaColaFactory) CreateBotle(volume float64) AbstractBotle {
	return &CocaColaBotle{volume: volume}
}

// CocaColaWater реалізує інтерфейс AbstractWater.
type CocaColaWater struct {
	volume float64
}

// GetVolume повертає об'єм напою і крім цього є методом CocaColaWater.
func (w *CocaColaWater) GetVolume() float64 {
	return w.volume
}

// CocaColaBottle реалізує AbstractBottle.
type CocaColaBotle struct {
	water  AbstractWater
	volume float64
}

// PourWater налиття води в пляшку.
func (b *CocaColaBotle) PourWater(water AbstractWater) {
	b.water = water
}

// GetBottleVolume повертає об'єм пляшки.
func (b *CocaColaBotle) GetBottleVolume() float64 {
	return b.volume
}

// GetWaterVolume повертає об'єм води у пляшці.
func (b *CocaColaBotle) GetWaterVolume() float64 {
	return b.water.GetVolume()
}
