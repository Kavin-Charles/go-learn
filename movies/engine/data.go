package engine
import (
	"net/http"
	"encoding/json"
	"fmt"
)


func ParseData(name, t string) any {

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

	for _, i := range searchData {
		if i.(map[string]any)["Type"].(string) == ty {
			final = i
			break
		} 
	}
	i := final.(map[string]any)
	if i["Type"] == "movie" {
		model := MovieModel {
		  i["Title"].(string),
			i["imdbID"].(string),
		}
		return model

	} else {
		model := SeriesModel {
			i["Title"].(string),
			i["imdbID"].(string),
			1,
			1,
		}
		return model
	}

}
