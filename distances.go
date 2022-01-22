// Calcul des distances entre les Églises ouvertes.
// Nécessite une première compilation par Hugo.
// Crée un fichier distances.json dans le répertoire data.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/connerdouglass/go-geo"
)

// définition pour reprise du json général
type Paroisses []struct {
	Title     string  `json:"title"`
	URL       string  `json:"url"`
	Canton    string  `json:"canton"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// définition d'une paroisse cible
type Cible struct {
	Title    string  `json:"nom"`
	URL      string  `json:"url"`
	Canton   string  `json:"canton"`
	Distance float64 `json:"distance"`
}

func main() {

	// ouverture, validation et fermeture du json général
	jsonFile, err := os.Open("public/index.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var paroisses Paroisses

	json.Unmarshal(byteValue, &paroisses)

	distances := make(map[string][]Cible)

	// parcours des paroisses en tant que source
	for i := 0; i < len(paroisses); i++ {

		a := geo.Coordinate{
			Latitude:  paroisses[i].Latitude,
			Longitude: paroisses[i].Longitude,
		}

		var cibles []Cible

		// parcours de chaque paroisse cible pour une source donnée
		for j := 0; j < len(paroisses); j++ {

			// calcul de la distance source-cible
			b := geo.Coordinate{
				Latitude:  paroisses[j].Latitude,
				Longitude: paroisses[j].Longitude,
			}

			distance := geo.DistanceBetween(a, b).Kilometers()

			// création de la cible et ajout à la liste
			var cible Cible

			cible.Title = paroisses[j].Title
			cible.URL = paroisses[j].URL
			cible.Canton = paroisses[j].Canton
			cible.Distance = distance

			cibles = append(cibles, cible)
		}

		// tri des cibles par distances
		sort.SliceStable(cibles, func(a, b int) bool {
			return cibles[a].Distance < cibles[b].Distance
		})

		// raccourcissement de la liste
		cibles = cibles[1:6]

		// ajoute à la table générale
		distances[paroisses[i].URL] = cibles

	}

	// création, validation et écriture du json de sortie
	output, err := json.Marshal(distances)
	if err != nil {
		fmt.Println(err)
	}

	ioutil.WriteFile("data/distances.json", output, os.ModePerm)
}
