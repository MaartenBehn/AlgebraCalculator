package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"strings"
)

type logPanel struct {
	*panel
	list *widget.List
	data binding.StringList
}

func newLogPanel(editor *editorPanel) *logPanel {
	l := &logPanel{panel: newPanel()}
	l.data = binding.NewStringList()

	headerMobile := container.NewBorder(nil, nil, nil,
		container.NewHBox(
			widget.NewButton("Back", func() {
				changeContent(editor.content[layoutMobile])
			}),
		),
	)

	l.newList(nil)
	l.content[layoutDesktop] = l.list
	l.content[layoutMobile] = container.NewBorder(headerMobile, nil, nil, nil, l.list)
	return l
}
func (l *logPanel) newList(lines []string) {
	err := l.data.Set(lines)
	if err != nil {
		panic(err)
	}

	l.list = widget.NewListWithData(l.data, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(item binding.DataItem, object fyne.CanvasObject) {
		text, err := item.(binding.String).Get()
		recover()
		if err != nil {
			return
		}

		object.(*widget.Label).Text = text
	})
}

func (l *logPanel) setTexts(texts []string) {
	var lines []string
	for _, text := range texts {
		lines = append(lines, strings.Split(text, "\n")...)
	}
	l.newList(lines)
}
