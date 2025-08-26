package main

import "fmt"

type Snapshot struct {
	content string
	cursor  int
}

type TextField struct {
	content string
	cursor  int
}

type History struct{ stack []Snapshot }

func (t *TextField) typeText(text string) { t.content += text; t.cursor = len(t.content) }

func (t *TextField) moveCursor(pos int) {
	if pos < 0 {
		pos = 0
	}
	if pos > len(t.content) {
		pos = len(t.content)
	}
	t.cursor = pos
}

func (t *TextField) save() Snapshot { return Snapshot{content: t.content, cursor: t.cursor} }

func (t *TextField) restore(m Snapshot) { t.content, t.cursor = m.content, m.cursor }

func (h *History) push(m Snapshot) { h.stack = append(h.stack, m) }

func (h *History) pop() (Snapshot, bool) {
	if len(h.stack) == 0 {
		return Snapshot{}, false
	}
	m := h.stack[len(h.stack)-1]
	h.stack = h.stack[:len(h.stack)-1]
	return m, true
}

func (t *TextField) print() { fmt.Printf("Content: %q, cursor: %d\n", t.content, t.cursor) }
