package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var window fyne.Window
var editor *Editor
var log *Log

func main() {
	a := app.New()

	window = a.NewWindow("AlgebraCalculator")
	/*
		window.SetMainMenu(&fyne.MainMenu{
			Items: []*fyne.Menu{
				fyne.NewMenu("File",
					fyne.NewMenuItem("Save", saveFile),
					fyne.NewMenuItem("Load", loadFile),
				),
			}})
	*/

	editor = NewEditor()
	log = NewLog()

	window.SetContent(editor.content)
	window.ShowAndRun()
}

func changeContent(content fyne.CanvasObject) {
	window.SetContent(content)
	window.Show()
}

func saveFile() {

}

func loadFile() {

}

func onRunButton() {

}
