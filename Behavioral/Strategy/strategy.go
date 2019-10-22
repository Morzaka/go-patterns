// Strategy -- приклад шаблону стратегії.
package Strategy

// StrategySort надає інтерфейс для алгоритмів сортування.
type StrategySort interface {
	Sort([]int)
}

// BubbleSort реалізує бульбашковий алгоритм сортування.
type BubbleSort struct{}

// Sort сортує дані
func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// InsertionSort реалізує алгоритм "вставного" сортування.
type InsertionSort struct{}

// Sort сортує дані
func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

// Context містить контекст для виконання стратегії.
type Context struct {
	strategy StrategySort
}

// Algorithm підміняє стратегії.
func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

// Sort сортує дані відповідно до обраної стратегії.
func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}
