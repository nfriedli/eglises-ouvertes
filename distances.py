#!/usr/bin/env python3

# Il est recommandé d'utiliser les script de conversion en go.
# Celui-ci ne sera pas forcément tenu à jour.

import json
from math import radians, cos, sin, asin, sqrt
from operator import itemgetter

# source:
# https://stackoverflow.com/questions/4913349/haversine-formula-in-python-bearing-and-distance-between-two-gps-points


def haversine(lon1, lat1, lon2, lat2):
    """
    Calculate the great circle distance in kilometers between two points
    on the earth (specified in decimal degrees)
    """
    # convert decimal degrees to radians
    lon1, lat1, lon2, lat2 = map(radians, [lon1, lat1, lon2, lat2])

    # haversine formula
    dlon = lon2 - lon1
    dlat = lat2 - lat1
    a = sin(dlat / 2)**2 + cos(lat1) * cos(lat2) * sin(dlon / 2)**2
    c = 2 * asin(sqrt(a))
    # Radius of earth in kilometers. Use 3956 for miles. Determines return
    # value units.
    r = 6371
    return c * r


with open("public/index.json") as source:
    paroisses = json.load(source)

source.close()

# tableau des distances qui sera converti en json
distances = {}

# on parcourt les paroisses sources
for i in paroisses:

    # on crée un liste de cibles pour chaque paroisse
    cibles = []

    # on parcourt les paroisses cibles (potentielles)
    for j in paroisses:
        distance = haversine(
            i["longitude"],
            i["latitude"],
            j["longitude"],
            j["latitude"])

        # pour chaque cible, on crée une donnée structurée
        cible = {}
        cible["nom"] = j["title"]
        cible["url"] = j["url"]
        cible["canton"] = j["canton"]
        cible["distance"] = distance

        # on ajoute la cible à la liste de cibles
        cibles.append(cible)

    # on trie les cibles par distances
    cibles.sort(key=itemgetter("distance"))

    # on coupe la liste en supprimant le doublet (source = cible) et en
    # prenant les 5 premières cibles
    cibles = cibles[1:6]

    # on ajoute les cibles aux distances (avec source comme clé)
    distances[i["url"]] = cibles

with open("data/distances.json", "w", encoding="utf8") as output:
    json.dumps(distances, output, ensure_ascii=False, sort_keys=True)

output.close()
