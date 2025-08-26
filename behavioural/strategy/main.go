package main

func main() {
	data := []int{5, 3, 8}
	s := &Sorter{strategy: BubbleSort{}}
	s.sort(data)
	s.setStrategy(QuickSort{})
	s.sort(data)
}
