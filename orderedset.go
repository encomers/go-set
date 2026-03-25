package go_set

// OrderedSet is a set that provides additional methods for ordered types (int, float, string, etc.).
type OrderedSet[T Ordered] struct {
	ISet[T]
}

func NewOrderedSet[T Ordered](elements ...T) *OrderedSet[T] {
	return &OrderedSet[T]{ISet: NewSet(elements...)}
}

// Min возвращает минимальный элемент (натуральный порядок).
func (s *OrderedSet[T]) Min() T {
	return Min(s) // Min из base.go принимает ISet[T] с констрейном Ordered
}

// Max возвращает максимальный элемент.
func (s *OrderedSet[T]) Max() T {
	return Max(s)
}

// Sum возвращает сумму всех элементов.
func (s *OrderedSet[T]) Sum() T {
	return Sum(s)
}

// Sorted возвращает отсортированный слайс по натуральному порядку
// (удобная обёртка, которой нет в базовых функциях).
func (s *OrderedSet[T]) Sorted() []T {
	return Sort(s, func(a, b T) bool { return a < b })
}

// Sort оставляем как есть (пользовательская сортировка).
func (s *OrderedSet[T]) Sort(sortFunc func(a, b T) bool) []T {
	return Sort(s, sortFunc)
}
