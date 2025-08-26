package main

import "fmt"

type EditorMemento struct{ content string }

type Editor struct{ content string }

type History struct{ stack []EditorMemento }

func (e *Editor) typeText(text string) { e.content += text }

func (e *Editor) save() EditorMemento { return EditorMemento{content: e.content} }

func (e *Editor) restore(m EditorMemento) { e.content = m.content }

func (h *History) push(m EditorMemento) { h.stack = append(h.stack, m) }

func (h *History) pop() (EditorMemento, bool) {
	if len(h.stack) == 0 {
		return EditorMemento{}, false
	}
	m := h.stack[len(h.stack)-1]
	h.stack = h.stack[:len(h.stack)-1]
	return m, true
}

func (e *Editor) print() { fmt.Println("Content:", e.content) }
