package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Editor struct {
	content fyne.CanvasObject
	list    *fyne.Container
	items   []*EditorItem
}

func NewEditor() *Editor {
	e := &Editor{}

	header := container.NewBorder(nil, nil, nil,
		container.NewHBox(
			widget.NewButton("Run", onRunButton),
			widget.NewButton("Log", func() {
				changeContent(log.content)
			}),
		),
	)

	e.list = container.NewVBox()
	scroll := container.NewVScroll(e.list)

	e.content = container.NewBorder(header, nil, nil, nil, scroll)

	e.addItem()
	e.addItem()
	e.addItem()

	return e
}

func (e *Editor) addItem() {
	editorItem := NewEditorItem()
	editorItem.index = len(e.list.Objects)
	e.list.Add(editorItem.content)
	e.items = append(e.items, editorItem)
}

func (e *Editor) removeItem(editorItem EditorItem) {
	e.list.Remove(editorItem.content)
	e.items = append(e.items[:editorItem.index], e.items[editorItem.index+1:]...)
}
