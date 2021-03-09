package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type EditorItem struct {
	content     fyne.CanvasObject
	entry       *widget.Entry
	label       *widget.Label
	closeButton *widget.Button
	index       int
}

func NewEditorItem() *EditorItem {
	editorItem := &EditorItem{}
	editorItem.entry = widget.NewEntry()
	editorItem.label = widget.NewLabel("")

	editorItem.content = container.NewVBox(editorItem.entry, editorItem.label)
	return editorItem
}

func (e *EditorItem) setText(text string) {
	e.entry.SetText(text)
}

func (e *EditorItem) getText() string {
	return e.entry.Text
}

func (e *EditorItem) setResult(text string) {
	e.label.SetText(text)
}
