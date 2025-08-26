package main

// IconFactory caches and returns shared icon flyweights.
type IconFactory struct {
	cache map[string]*Icon
}

func NewIconFactory() *IconFactory {
	return &IconFactory{cache: make(map[string]*Icon)}
}

func (f *IconFactory) Get(name string) *Icon {
	if icon, ok := f.cache[name]; ok {
		return icon
	}
	// Simulate that each distinct icon costs 8KB to store.
	icon := newIcon(name, 8*1024)
	f.cache[name] = icon
	return icon
}

func (f *IconFactory) Count() int {
	return len(f.cache)
}
