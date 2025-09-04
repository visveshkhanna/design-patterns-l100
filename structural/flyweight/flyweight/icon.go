package flyweight

type Icon struct {
	Name      string
	SizeBytes int
}

func NewIcon(name string, sizeBytes int) *Icon {
	return &Icon{Name: name, SizeBytes: sizeBytes}
}

func (i *Icon) GetName() string {
	return i.Name
}

func (i *Icon) GetSize() int {
	return i.SizeBytes
}
