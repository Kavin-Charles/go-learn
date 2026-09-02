package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Editor struct {
	widget.BaseWidget

	Text string
	RichText *widget.RichText
}
func main(){
	a := app.NewWithID("editor")
	window := a.NewWindow("Editor")

	cont := container.NewCenter(NewEditor("FIRST TEXT?"))
	window.SetContent(cont)
	window.ShowAndRun()
}

func (e *Editor) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(e.RichText)
}

func NewEditor(code string) *Editor{
	editor := &Editor{ Text: code }
	editor.RichText = widget.NewRichTextFromMarkdown(code)
	editor.ExtendBaseWidget(editor)
	return editor
}
