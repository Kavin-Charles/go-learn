package engine
import (
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
)

var Season, Episode = 0, 0


func ParseData(name, t string) any {
	name = strings.ReplaceAll(name, " ", "%20")
	url := "http://www.omdbapi.com/?apikey=a7c2263a&s=" + name
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println("Error while parsing data: ", err)
	}
	defer resp.Body.Close()
	var rawData map[string]any	
	err = json.NewDecoder(resp.Body).Decode(&rawData)
	if err != nil {
		fmt.Println("error while getting raw data: ", err)
	}

	searchData, ok:= rawData["Search"].([]any)
	if !ok  {
		fmt.Println("Error while getting searchData")
	}
//	fmt.Println(searchData)
//i := searchData[0].(map[string]any)
	ty := "movie"
	if t == "s" {
		ty = "series"
	}

	var final any
	finalList := make([]any, 0)
	fmt.Println(len(searchData))
	for m := 0; m < len(searchData) - 1; m++ {
		if searchData[m].(map[string]any)["Type"].(string) == ty {
			i := searchData[m].(map[string]any)

			if ty == "movie" {
			final = MovieModel {
		i["Title"].(string),
	i["imdbID"].(string),

			}
		} else {
			final = SeriesModel {
			i["Title"].(string),
			i["imdbID"].(string),
			Season,
		Episode,
		}

		}
			finalList = append(finalList, final)
		} 
	}

	fmt.Println("Enter an option :")
	for o, i := range finalList {
		fmt.Printf("%d. %s\n", o, i)
	}
	opt := 0
	fmt.Scan(&opt)

 return finalList[opt]
}
