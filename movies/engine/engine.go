package engine
import (
	"fmt"
	"net/http"
	"html/template"
)

func RunMovie(name string, w http.ResponseWriter) {
	var MovieData MovieModel = ParseData(name, "m").(MovieModel)
	templ, err := template.ParseFiles("movie.html")
	if err != nil {
		fmt.Println("cannot load movie tempalte!")
	}

	err = templ.Execute(w,MovieData)
	
}

func RunSeries(name string, w http.ResponseWriter){
	var SeriesData SeriesModel = ParseData(name, "s").(SeriesModel)
	templ, err := template.ParseFiles("series.html")
	if err != nil {
		fmt.Println("cannot load movie tempalte!")
	}

	err = templ.Execute(w, SeriesData)

}

func OpenWeb(data string, t string) {
	if t == "m" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {RunMovie(data, w)})
	}	else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {RunSeries(data, w)})

	}

	http.ListenAndServe(":8000", nil)
}


