package iterators

import "design-patterns/behavioural/iterator/models"

type IIterator interface {
	HasNext() bool
	Next() models.LogEntry
}
