package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"strings"
)

type Log struct {
	content fyne.CanvasObject
	list    *widget.List
	data    binding.StringList
}

func NewLog() *Log {
	l := &Log{}
	l.data = binding.NewStringList()

	header := container.NewBorder(nil, nil, nil,
		container.NewHBox(
			widget.NewButton("Back", func() {
				changeContent(editor.content)
			}),
		),
	)

	l.newList(nil)
	l.content = container.NewBorder(header, nil, nil, nil, l.list)
	return l
}

func (l *Log) newList(lines []string) {
	l.data.Set(lines)
	l.list = widget.NewListWithData(l.data, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(item binding.DataItem, object fyne.CanvasObject) {
		text, err := item.(binding.String).Get()
		if err != nil {
			panic(err)
		}
		object.(*widget.Label).Text = text
	})
}

func (l *Log) setText(text string) {
	lines := strings.Split(text, "\n")
	l.newList(lines)
}
