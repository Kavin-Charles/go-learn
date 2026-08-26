package main 
import (
	"fyne.io/fyne/v2/app"
	"movies/engine"
) 

func main() {
	a := app.NewWithID("com.kavin.movies")
	w := a.NewWindow("Kavin Movies")

	engine.OpenHomePage(w)

	//engine.RenderHomePage(w)

	w.Show()
	a.Run()
}
