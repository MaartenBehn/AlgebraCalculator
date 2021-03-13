package main

import (
	"AlgebraCalculator/V4"
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
			widget.NewButton("Clear All", e.clearAll),
			widget.NewButton("Log", func() {
				changeContent(log.content)
			}),
		),
	)

	e.list = container.NewVBox()
	scroll := container.NewVScroll(e.list)

	e.content = container.NewBorder(header, nil, nil, nil, scroll)

	e.update()

	return e
}

func (e *Editor) update() {
	if len(e.items) == 0 || e.items[len(e.items)-1].getText() != "" {
		e.items = append(e.items, NewEditorItem(e))
	}

	var items []fyne.CanvasObject
	for _, item := range e.items {
		items = append(items, item.content)
	}
	e.list.Objects = items

	result := V4.Calculate(e.getAllTexts()...)
	for i, item := range e.items {
		if result.TermStrings[i] == "" {
			item.setResult("")
		} else {
			item.setResult(result.TermStrings[i])
		}
	}

	if log != nil {
		log.setTexts(result.Log)
	}
}

func (e *Editor) getAllTexts() []string {
	var texts []string
	for _, item := range e.items {
		texts = append(texts, item.getText())
	}
	return texts
}

func (e *Editor) clearAll() {
	e.items = []*EditorItem{}
	e.update()
}
