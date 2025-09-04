package collections

import (
	"design-patterns/behavioural/iterator/iterators"
	"design-patterns/behavioural/iterator/models"
)

type Log struct {
	entries []models.LogEntry
}

func NewLog(entries []models.LogEntry) *Log {
	return &Log{entries: entries}
}

func (l *Log) CreateIterator() iterators.IIterator {
	return iterators.NewLogIterator(l.entries)
}
