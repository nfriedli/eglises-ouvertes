import json
import folium

with open("public/index.json") as source:
    paroisses = json.load(source)

# Carte positionnée chez moi ;)
m = folium.Map(location=[46.96661, 6.86231])

for i in paroisses:
    folium.Marker([i["latitude"], i["longitude"]],
    popup=i["title"]).add_to(m)

m.save("public/carte.html")