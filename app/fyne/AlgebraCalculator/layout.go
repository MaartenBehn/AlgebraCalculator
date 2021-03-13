package main

import "fyne.io/fyne/v2"

const (
	layoutDesktop = 1
	layoutMobile  = 2
	layoutMax     = 3
)

var currentLayout int

func getLayout(size fyne.Size) int {
	if size.Width/size.Height < 1 {
		return layoutMobile
	}
	return layoutDesktop
}
