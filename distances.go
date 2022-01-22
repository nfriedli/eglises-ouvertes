package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/connerdouglass/go-geo"
)

type Paroisses []struct {
	Title     string  `json:"title"`
	URL       string  `json:"url"`
	Canton    string  `json:"canton"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Cible struct {
	Title    string  `json:"nom"`
	URL      string  `json:"url"`
	Canton   string  `json:"canton"`
	Distance float64 `json:"distance"`
}

func main() {

	jsonFile, err := os.Open("public/index.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var paroisses Paroisses

	json.Unmarshal(byteValue, &paroisses)

	distances := make(map[string][]Cible)

	for i := 0; i < len(paroisses); i++ {

		a := geo.Coordinate{
			Latitude:  paroisses[i].Latitude,
			Longitude: paroisses[i].Longitude,
		}

		var cibles []Cible

		for j := 0; j < len(paroisses); j++ {

			b := geo.Coordinate{
				Latitude:  paroisses[j].Latitude,
				Longitude: paroisses[j].Longitude,
			}

			distance := geo.DistanceBetween(a, b).Kilometers()
			var cible Cible

			cible.Title = paroisses[j].Title
			cible.URL = paroisses[j].URL
			cible.Canton = paroisses[j].Canton
			cible.Distance = distance

			//println("-> ", cible.Title, cible.Distance)

			cibles = append(cibles, cible)
		}

		sort.SliceStable(cibles, func(a, b int) bool {
			return cibles[a].Distance < cibles[b].Distance
		})
		cibles = cibles[1:6]
		distances[paroisses[i].URL] = cibles

	}

	// output, err := json.MarshalIndent(distances, "", "  ")
	output, err := json.Marshal(distances)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile("data/distances.json", output, os.ModePerm)
}
