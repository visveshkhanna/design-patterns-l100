package main

func main() {
	shapes := []Element{&Circle{radius: 3}, &Square{side: 4}}
	visitor := AreaCalculator{}
	for _, s := range shapes {
		s.accept(visitor)
	}
}
