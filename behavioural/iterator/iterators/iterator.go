package iterators

import "design-patterns/behavioural/iterator/models"

type Iterator interface {
	HasNext() bool
	Next() models.LogEntry
}
