package main

import (
	"design-patterns/behavioural/iterator/collections"
	"design-patterns/behavioural/iterator/models"
	"fmt"
)

func main() {
	l := collections.NewLog([]models.LogEntry{
		{Level: "INFO", Msg: "server started"},
		{Level: "WARN", Msg: "cache miss"},
		{Level: "ERROR", Msg: "db timeout"},
	})

	it := l.CreateIterator()
	for it.HasNext() {
		e := it.Next()
		fmt.Println(e.Level+":", e.Msg)
	}
}
