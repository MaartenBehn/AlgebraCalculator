package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type EditorItem struct {
	content fyne.CanvasObject
	editor  *Editor

	entry       *widget.Entry
	label       *widget.Label
	closeButton *widget.Button
}

func NewEditorItem(editor *Editor) *EditorItem {
	editorItem := &EditorItem{editor: editor}

	editorItem.entry = widget.NewEntry()
	editorItem.entry.OnChanged = (editorItem).onChange

	editorItem.label = widget.NewLabel("")

	editorItem.closeButton = widget.NewButton("X", (editorItem).onClose)

	editorItem.content = container.NewBorder(nil, nil, nil,
		container.NewCenter(editorItem.closeButton),
		container.NewVBox(editorItem.entry, editorItem.label))
	return editorItem
}

func (e *EditorItem) setText(text string) {
	e.entry.SetText(text)
	e.onChange(text)
}

func (e *EditorItem) getText() string {
	return e.entry.Text
}

func (e *EditorItem) setResult(text string) {
	e.label.SetText(text)
}

func (e *EditorItem) onChange(string) {
	e.editor.update()
}

func (e *EditorItem) onClose() {
	if len(e.editor.items) <= 1 {
		e.setText("")
		return
	}
	e.editor.removeItem(e)
}
