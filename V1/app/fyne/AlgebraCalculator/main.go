package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"time"
)

const (
	panelIdEditor = 1
	panelIdMax    = 2
	panelIdStart  = 1
)

var window fyne.Window
var basePanel fyne.CanvasObject
var panels []*panel
var currentPanel int

func main() {
	a := app.New()

	a.Settings().SetTheme(theme.DarkTheme())

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

	panels = make([]*panel, panelIdMax)
	panels[panelIdEditor] = newEditorPanel().panel
	currentPanel = panelIdStart

	go updateLoop()
	window.ShowAndRun()

	running = false
}

var running bool
var fps float64

const maxFPS float64 = 10

func updateLoop() {
	startTime := time.Now()
	var startDuration time.Duration
	wait := time.Duration(1000000000 / int(maxFPS))

	running = true
	for running {
		startDuration = time.Since(startTime)
		// All update Calls

		checkLayout()

		diff := time.Since(startTime) - startDuration
		if diff > 0 {
			fps = (wait.Seconds() / diff.Seconds()) * maxFPS
		} else {
			fps = 10000
		}
		if diff < wait {
			time.Sleep(wait - diff)
		}
	}
}

func checkLayout() {
	newLayout := getLayout(window.Content().Size())
	if newLayout != currentLayout {
		currentLayout = newLayout

		changeContent(panels[currentPanel].content[currentLayout])
	}
}

func changeContent(content fyne.CanvasObject) {
	window.SetContent(content)
	window.Canvas().Content().Refresh()
	window.Show()
}
