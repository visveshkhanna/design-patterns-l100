package main

type Iterator interface {
	hasNext() bool
	next() LogEntry
}
