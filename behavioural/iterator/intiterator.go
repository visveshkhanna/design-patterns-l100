package main

type LogEntry struct {
	Level string
	Msg   string
}

type Log struct{ entries []LogEntry }

type LogIterator struct {
	log   *Log
	index int
}

func (l *Log) createIterator() *LogIterator {
	return &LogIterator{log: l, index: 0}
}

func (it *LogIterator) hasNext() bool { return it.index < len(it.log.entries) }

func (it *LogIterator) next() LogEntry {
	val := it.log.entries[it.index]
	it.index++
	return val
}
