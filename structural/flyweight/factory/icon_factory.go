package factory

import "design-patterns/structural/flyweight/flyweight"

type IconFactory struct {
	cache map[string]*flyweight.Icon
}

func NewIconFactory() *IconFactory {
	return &IconFactory{cache: make(map[string]*flyweight.Icon)}
}

func (f *IconFactory) Get(name string) *flyweight.Icon {
	if icon, ok := f.cache[name]; ok {
		return icon
	}

	icon := flyweight.NewIcon(name, 8*1024)
	f.cache[name] = icon
	return icon
}

func (f *IconFactory) Count() int {
	return len(f.cache)
}
