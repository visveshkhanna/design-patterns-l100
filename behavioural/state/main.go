package main

func main() {
	d := &Document{state: Draft{}, title: "Proposal"}
	d.submit()  // -> Review
	d.approve() // -> Published
	d.reject()  // noop
}
