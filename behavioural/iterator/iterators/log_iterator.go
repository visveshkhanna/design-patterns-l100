package iterators

import "design-patterns/behavioural/iterator/models"

type LogIterator struct {
	entries []models.LogEntry
	index   int
}

func NewLogIterator(entries []models.LogEntry) *LogIterator {
	return &LogIterator{entries: entries, index: 0}
}

func (it *LogIterator) HasNext() bool {
	return it.index < len(it.entries)
}

func (it *LogIterator) Next() models.LogEntry {
	val := it.entries[it.index]
	it.index++
	return val
}
