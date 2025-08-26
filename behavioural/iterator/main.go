package main

import "fmt"

func main() {
	l := &Log{entries: []LogEntry{
		{Level: "INFO", Msg: "server started"},
		{Level: "WARN", Msg: "cache miss"},
		{Level: "ERROR", Msg: "db timeout"},
	}}
	it := l.createIterator()
	for it.hasNext() {
		e := it.next()
		fmt.Println(e.Level+":", e.Msg)
	}
}
