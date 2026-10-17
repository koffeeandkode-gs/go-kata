package main

import "fmt"

type Stack[T any] struct {
	items []T
}

type Product struct {
	Name  string
	Price int
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := len(s.items) - 1
	v := s.items[last]
	s.items[last] = zero
	s.items = s.items[:last]
	return v, true

}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func main() {

	numbers := Stack[int]{}
	numbers.Push(10)
	numbers.Push(20)

	names := Stack[string]{}
	names.Push("Go")
	names.Push("Python")

	products := Stack[Product]{}
	products.Push(Product{"Notebook", 50})

	fmt.Println(numbers.Peek()) // 20 true
	fmt.Println(numbers.Pop())  // 20 true
	fmt.Println(numbers.Pop())  // 10 true
	fmt.Println(numbers.Pop())  // 0 false

	fmt.Println("num items:", numbers)
	fmt.Println("names items:", names)
	fmt.Println("product items:", products)
}
