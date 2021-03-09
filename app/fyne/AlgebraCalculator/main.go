package main

import (
	"AlgebraCalculator"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var window fyne.Window
var editor *Editor
var log *Log

func main() {

	var ruleStrings []string
	ruleStrings = append(ruleStrings, string(resourceSimpRulesExpandTxt.Content()))
	ruleStrings = append(ruleStrings, string(resourceSimpRulesSumUpTxt.Content()))
	AlgebraCalculator.Init(ruleStrings)

	a := app.New()

	window = a.NewWindow("AlgebraCalculator")
	window.Resize(fyne.NewSize(1600, 900))

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
