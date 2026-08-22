package main 
import (
	"flag"
	"movies/engine"
) 

var entryName, entryType string
var Season, Episode int
func main() {
	getInput()
	engine.OpenWeb(entryName, entryType, Season, Episode)
}

func getInput() {
	flag.StringVar(&entryName, "n", "Avengers", "Movie or Series Name")
	flag.StringVar(&entryType, "t", "m", "Enter series or movie")
	flag.IntVar(&Episode, "e", 1, "Enter series or movie")
	flag.IntVar(&Season, "s", 1, "Enter series or movie")
	flag.Parse()
}

 
