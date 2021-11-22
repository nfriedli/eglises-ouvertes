# eglises-ouvertes

Liste d’églises protestantes réformées ouvertes en Suisse romande.

Site: https://eglises-ouvertes.ch/

Contact: nicolas.friedli+eo@gmail.com

## Compilation

Le site est généré par Hugo: https://gohugo.io/ (version standard suffisante)

Le calcul des distances est effectué par le script `distances.js` à la racine du projet.
Il requiert l'installation de `node.js`.
`npm install` permet d'installer `geolib`.

Pour une compilation complète:

- `hugo` (crée notamment le fichier `/public/index.json`)
- `node distances.js` (utilise le fichier `/public/index.json` et crée le fichier `/data/distances.json`)
- `hugo` (utilise le fichier `/data/distances.json`)

## Fichiers spéciaux

- `/data/vitraux.json` permet d'entrer des dates pour les artistes
- `/config/_default/config.yaml` permet d'éditer le titre du site et le pied de page

## Layouts

- `/layouts/_default/list.html` est utilisé pour les pages cantonales
- `/layouts/_default/single.html` est utilisé pour les pages d'églises
- `/layouts/about/search.html` est utilisé pour la page de recherche

## Shortcodes

- `cantons.html` renvoie une liste des cantons du site
- `pbc.html` doit être utilisé avec un paramètres (`A` ou `B`)
- `sitemap.html` renvoie une liste de toutes les pages du site
- `vitraux.html` renvoie une liste des artistes et églises (utilise: `/data/vitraux.json`)

## Réutilisation

L'ensemble des fichiers peut être réutilisé librement, sauf:

- `fuse.js` pour la recherche
- *Fraunces*, la police d'écriture des titres
- le code d'image `opengraph` chez Cloudinary
- `geolib` à installer avant calcul des distances

Merci de vérifier les licences ;)