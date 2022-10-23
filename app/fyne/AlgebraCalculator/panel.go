package main

import "fyne.io/fyne/v2"

type panel struct {
	content []fyne.CanvasObject
}

func newPanel() *panel {
	return &panel{
		content: make([]fyne.CanvasObject, layoutMax),
	}
}
