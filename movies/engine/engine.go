package engine
import (
	"fmt"
	"net/http"
	"html/template"
)

func RunMovie(MovieData MovieModel, w http.ResponseWriter) {
	templ, err := template.ParseFiles("movie.html")
	if err != nil {
		fmt.Println("cannot load movie tempalte!")
	}
	err = templ.Execute(w,MovieData)
	
}

func RunSeries(SeriesData SeriesModel,w http.ResponseWriter){
	templ, err := template.ParseFiles("series.html")
	if err != nil {
		fmt.Println("cannot load movie tempalte!")
	}

	err = templ.Execute(w, SeriesData)

}

func OpenWeb(data string, t string, season int, episode int) {



	if t == "m" {
		entryData := ParseData(data, t)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {RunMovie(entryData.(MovieModel), w)})
	}	else {
		Data := ParseData(data, t).(SeriesModel)
		seriesData := SeriesModel {
			Data.Name,
			Data.ImdbId,
			season,
			episode,
		}
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {RunSeries(seriesData, w)})

	}

	http.ListenAndServe(":8000", nil)
}


