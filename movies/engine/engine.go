package engine
import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func OpenHomePage(w fyne.Window) {
	fmt.Println("Opening Home Page")
	movieContent := widget.NewLabel("Movie Label")
	w.SetContent(widget.NewCard("Movie 1", "Action Movie", movieContent))
}


