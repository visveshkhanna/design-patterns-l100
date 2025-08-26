package main

// Icon is the flyweight: intrinsic state shared across many markers.
type Icon struct {
	name      string
	sizeBytes int
}

func newIcon(name string, sizeBytes int) *Icon {
	return &Icon{name: name, sizeBytes: sizeBytes}
}
