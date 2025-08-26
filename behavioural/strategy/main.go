package main

func main() {
	r := &Ride{strategy: Economy{}}
	r.printQuote(12)
	r.setStrategy(Premium{})
	r.printQuote(12)
}
