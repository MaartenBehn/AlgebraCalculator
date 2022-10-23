package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	AC "github.com/MaartenBehn/AlgebraCalculator"
)

type editorPanel struct {
	*panel
	list     *fyne.Container
	logPanel *logPanel

	Items []*editorListItem
}

func newEditorPanel() *editorPanel {
	e := &editorPanel{
		panel: newPanel(),
	}

	headerDesktop := container.NewBorder(nil, nil, nil,
		container.NewHBox(
			widget.NewButton("Clear All", nil),
		),
	)
	headerMobile := container.NewBorder(nil, nil, nil,
		container.NewHBox(
			widget.NewButton("Clear All", nil),
			widget.NewButton("Log", func() {
				changeContent(e.logPanel.content[layoutMobile])
			}),
		),
	)

	e.list = container.NewVBox()
	scroll := container.NewVScroll(e.list)

	e.logPanel = newLogPanel(e)

	e.content[layoutDesktop] = container.NewHSplit(
		container.NewBorder(headerDesktop, nil, nil, nil, scroll),
		e.logPanel.content[layoutDesktop])

	e.content[layoutMobile] = container.NewBorder(headerMobile, nil, nil, nil, scroll)

	e.updateItems()
	e.updateContent()

	return e
}

func (e *editorPanel) updateItems() {
	if len(e.Items) == 0 || e.Items[len(e.Items)-1].getText() != "" {
		e.Items = append(e.Items, newEditorItem(e))
	}

	for len(e.Items) >= 2 &&
		e.Items[len(e.Items)-1].getText() == "" &&
		e.Items[len(e.Items)-2].getText() == "" {
		e.Items = e.Items[:len(e.Items)-1]
	}
}
func (e *editorPanel) updateContent() {
	var itemContents []fyne.CanvasObject
	for _, item := range e.Items {
		itemContents = append(itemContents, item.content)
	}
	e.list.Objects = itemContents
}
func (e *editorPanel) updateCalculations() {
	result := AC.Calculate(e.getAllTexts()...)
	for i, item := range e.Items {
		item.setResult(result.TermStrings[i])
	}

	e.logPanel.setTexts(result.Log)
}
func (e *editorPanel) getAllTexts() []string {
	var texts []string
	for _, item := range e.Items {
		texts = append(texts, item.getText())
	}
	return texts
}

type editorListItem struct {
	content fyne.CanvasObject
	editor  *editorPanel

	entry       *widget.Entry
	label       *widget.Label
	closeButton *widget.Button

	Text   string
	Result string
}

func newEditorItem(editor *editorPanel) *editorListItem {
	e := &editorListItem{editor: editor}

	e.entry = widget.NewEntry()
	e.entry.OnChanged = e.onChanged

	e.label = widget.NewLabel("")

	e.closeButton = widget.NewButton("X", e.onCloseButton)

	e.content = container.NewBorder(nil, nil, nil,
		container.NewCenter(e.closeButton),
		container.NewVBox(e.entry, e.label))
	return e
}

func (e *editorListItem) setText(text string) {
	e.entry.SetText(text)
}
func (e *editorListItem) getText() string {
	return e.entry.Text
}
func (e *editorListItem) setResult(text string) {
	e.label.SetText(text)
}
func (e *editorListItem) onChanged(s string) {
	e.editor.updateCalculations()
	e.editor.updateItems()
	e.editor.updateContent()
}
func (e *editorListItem) onCloseButton() {
	index := 0
	for i, item := range e.editor.Items {
		if item == e {
			index = i
			break
		}
	}
	e.editor.Items = append(e.editor.Items[:index], e.editor.Items[index+1:]...)
	e.editor.updateItems()
	e.editor.updateContent()
	e.editor.updateCalculations()
}
