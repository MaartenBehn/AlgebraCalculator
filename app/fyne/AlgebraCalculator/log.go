package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Log struct {
	content fyne.CanvasObject
	log     *widget.TextGrid
}

func NewLog() *Log {
	log := &Log{}

	header := container.NewBorder(nil, nil, nil,
		container.NewHBox(
			widget.NewButton("Back", func() {
				changeContent(editor.content)
			}),
		),
	)

	log.log = widget.NewTextGridFromString("Log\nline 1")
	scroll := container.NewVScroll(log.log)
	log.content = container.NewBorder(header, nil, nil, nil, scroll)
	return log
}
