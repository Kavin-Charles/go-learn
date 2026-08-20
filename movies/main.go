package main 
import (
	"flag"
	"movies/engine"
) 

var entryName, entryType string

func main() {
	getInput()
	engine.OpenWeb(entryName, entryType)
}

func getInput() {
	flag.StringVar(&entryName, "n", "Avengers", "Movie or Series Name")
	flag.StringVar(&entryType, "t", "m", "Enter series or movie")
	flag.Parse()
}

 
