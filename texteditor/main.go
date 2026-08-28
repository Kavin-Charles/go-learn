package main 

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"io"
)

func main(){
	app := app.NewWithID("texteditor")
	window :=  app.NewWindow("Text Editor")

	editor := widget.NewMultiLineEntry()
	scroll := container.NewScroll(editor)

//  saveButton := widget.NewButton("Same", func (){
		
//	})

	openButton := widget.NewButton("Open", func (){
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {

			if reader == nil || err != nil {
				return
			}

			data, er := io.ReadAll(reader)
			defer reader.Close()

			if err != nil {
				dialog.NewError(er, window)
			}
			editor.SetText(string(data))
			window.SetTitle(reader.URI().Name())

		}, window)

		fd.Show()

	})

	toolbar := container.NewHBox(openButton)
	content := container.NewBorder(toolbar, nil, nil, nil, scroll)

	window.SetContent(content)
	window.ShowAndRun()
}
