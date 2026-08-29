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

	dirty := false
	scroll := container.NewScroll(editor)

  saveButton := widget.NewButton("Save", func (){
		FileSaveAs(window, []byte(editor.Text))
		dirty = false
	})

	openButton := widget.NewButton("Open", func (){
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {

			if reader == nil || err != nil {
				return
			}

			data, er := io.ReadAll(reader)
			defer reader.Close()

			if er != nil {
				dialog.NewError(er, window)
			}
			editor.SetText(string(data))
			window.SetTitle("Text editor " + reader.URI().Name())


		}, window)

		fd.Show()

	})

	window.SetCloseIntercept(func() {

		if dirty {
			fd := dialog.NewConfirm("Are you sure!?", "This file is not saved, do you want to quit", func(res bool){
				if res {
					window.Close()
				}
			}, window)
			fd.Show()
		}

	})

	toolbar := container.NewHBox(openButton, saveButton)
	content := container.NewBorder(toolbar, nil, nil, nil, scroll)

	editor.OnChanged = func(s string) {
		
		dirty = true
		if dirty && "*" != string(window.Title()[len(window.Title()) - 1]) {
			window.SetTitle(window.Title() + " *")
		}
	}

	window.SetContent(content)
	window.ShowAndRun()
}


func FileSaveAs(window fyne.Window, text []byte) {
	fd := dialog.NewFileSave(func(w fyne.URIWriteCloser, err error){
			
			if w == nil || err != nil {
				return
			}

			defer w.Close()
			w.Write(text)
			window.SetTitle("Text Editor " + w.URI().Name())

		}, window)	

		fd.Show()

}
