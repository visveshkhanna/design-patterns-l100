package main

type IntCollection struct{ items []int }

type IntIterator struct {
	collection *IntCollection
	index      int
}

func (c *IntCollection) createIterator() *IntIterator {
	return &IntIterator{collection: c, index: 0}
}

func (it *IntIterator) hasNext() bool { return it.index < len(it.collection.items) }

func (it *IntIterator) next() int {
	val := it.collection.items[it.index]
	it.index++
	return val
}
