// Calcul des distances entre les Églises ouvertes.
// Nécessite une première compilation par Hugo.
// Crée un fichier distances.json dans le répertoire data.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
)

// définition pour reprise du json général
type Paroisses []struct {
	Title     string  `json:"title"`
	URL       string  `json:"url"`
	Canton    string  `json:"canton"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// définition d’une paroisse cible
type Cible struct {
	Title    string  `json:"nom"`
	URL      string  `json:"url"`
	Canton   string  `json:"canton"`
	Distance float64 `json:"distance"`
}

// Calcul de distance par haversine
// source: https://gist.github.com/cdipaolo/d3f8db3848278b49db68
// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378.1 // Earth radius in km

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

// Le programme principal
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

		var cibles []Cible

		// parcours de chaque paroisse cible pour une source donnée
		for j := 0; j < len(paroisses); j++ {

			// calcul de la distance source-cible
			distance := Distance(paroisses[i].Latitude, paroisses[i].Longitude, paroisses[j].Latitude, paroisses[j].Longitude)

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
