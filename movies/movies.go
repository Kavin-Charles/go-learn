package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
)

var (
	movieName string = ""
	imdb      string
)

func main() {
	flag.StringVar(&movieName, "movie", "Avengers", "Movie name you want to search")
	flag.Parse()
	fmt.Println(movieName)

	url := "https://omdbapi.com/?apikey=a7c2263a&s="
	url = url + movieName

	fmt.Println(url)
	data, err := http.Get(url)
	defer data.Body.Close()
	if err != nil {
		fmt.Println("Something went wrong when getting data")
	}

	var moviesData map[string]any
	err = json.NewDecoder(data.Body).Decode(&moviesData)
	firstMovie := moviesData["Search"].([]interface{})[0].(map[string]any)
	imdb = firstMovie["imdbID"].(string)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

type pageData struct {
	ImdbId string
}

func handler(w http.ResponseWriter, _ *http.Request) {
	movieInfo := pageData{imdb}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Something went wrong while loading html")
		return
	}

	err = tmpl.Execute(w, movieInfo)
	if err != nil {
		fmt.Println("something went wrong while executing movieInfo: ", err)
	}
}
