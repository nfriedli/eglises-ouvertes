const fs = require("fs");
const { getDistance } = require("geolib");

const eglises = JSON.parse(fs.readFileSync("./public/index.json"))

distances = new Object()

for (var eglise1 in eglises) {
    
    url = eglises[eglise1].url

    distances[url] = []

    for (var eglise2 in eglises) {

        let distance = getDistance(
            {latitude: eglises[eglise1].latitude, longitude: eglises[eglise1].longitude},
            {latitude: eglises[eglise2].latitude, longitude: eglises[eglise2].longitude},
            1
            ) 

        distance = Math.round(distance / 100) / 10

        var temp = {
            "url": eglises[eglise2].url,
            "nom": eglises[eglise2].title,
            "canton": eglises[eglise2].canton,
            "distance": distance
        }

        distances[url].push(temp)
    }
    distances[url].sort( function (a,b) {
        return (b.distance-a.distance)
    })

    distances[url].reverse()

    let complet = distances[url]
    selection = complet.slice(1,6)
    distances[url] = selection


}

distances = JSON.stringify(distances)

fs.writeFileSync("./data/distances.json", distances)
