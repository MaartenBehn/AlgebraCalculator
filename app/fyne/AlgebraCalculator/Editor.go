package main

import (
	"AlgebraCalculator"
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
			widget.NewButton("Clear All", (e).clearAll),
			widget.NewButton("Log", func() {
				changeContent(log.content)
			}),
		),
	)

	e.list = container.NewVBox()
	scroll := container.NewVScroll(e.list)

	e.content = container.NewBorder(header, nil, nil, nil, scroll)

	e.addItem()

	return e
}

func (e *Editor) addItem() {
	editorItem := NewEditorItem(e)
	e.list.Add(editorItem.content)
	e.items = append(e.items, editorItem)
}

func (e *Editor) removeItem(editorItem *EditorItem) {
	e.list.Remove(editorItem.content)

	var index int
	for i, item := range e.items {
		if item == editorItem {
			index = i
			break
		}
	}
	e.items = append(e.items[:index], e.items[index+1:]...)
}

func (e *Editor) update() {
	if e.items[len(e.items)-1].getText() != "" {
		e.addItem()
	}

	results, logged := AlgebraCalculator.Run(e.getAllTexts()...)

	for i, item := range e.items {
		if results[i] == "" {
			item.setResult("")
		} else {
			item.setResult(results[i])
		}
	}

	log.setTexts(logged)
}

func (e *Editor) getAllTexts() []string {
	var texts []string
	for _, item := range e.items {
		texts = append(texts, item.getText())
	}
	return texts
}

func (e *Editor) clearAll() {
	for _, item := range e.items {
		e.removeItem(item)
	}
	e.addItem()
}
