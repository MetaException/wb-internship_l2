package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Преимущества:
		* Горячая замена алгоритмов на лету
		* Изолирует код и данные алгоритмов от остальных классов
		* Уход от наследования к делегированию
		* Реализует принцип открытости/закрытости

	Недостатки:
		* Усложняет программу за счёт дополнительных классов
		* Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую

	Применение:
		* Выбор метода оплаты
		* Сжатие файлов (выбор алгоритма)
*/

type SortStrategy interface {
	Sort([]int)
}

type BubbleSortStrategy struct{}

func (s BubbleSortStrategy) Sort(data []int) {
	n := len(data)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Println("Bubble sorted:", data)
}

type QuickSortStrategy struct{}

func (s QuickSortStrategy) Sort(data []int) {
	if len(data) < 2 {
		return
	}
	left, right := 0, len(data)-1
	pivot := data[len(data)/2]
	for left <= right {
		for data[left] < pivot {
			left++
		}
		for data[right] > pivot {
			right--
		}
		if left <= right {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}
	s.Sort(data[:right+1])
	s.Sort(data[left:])
	fmt.Println("Quick sorted:", data)
}

type Context struct {
	strategy SortStrategy
}

func (c *Context) SetStrategy(strategy SortStrategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(data []int) {
	c.strategy.Sort(data)
}

/*
func main() {
	data := []int{33, 10, 55, 71, 29, 5, 70}

	context := Context{}

	context.SetStrategy(BubbleSortStrategy{})
	context.ExecuteStrategy(append([]int{}, data...))

	context.SetStrategy(QuickSortStrategy{})
	context.ExecuteStrategy(append([]int{}, data...))
}
*/
