package models

type Message struct {
	Text     string
	Flags    []string
	Rejected bool
	Reason   string
}