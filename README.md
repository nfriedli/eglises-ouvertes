# eglises-ouvertes

Liste d’églises protestantes réformées ouvertes en Suisse romande.

Site: https://eglises-ouvertes.ch/

Contact: nicolas.friedli+eo@gmail.com

## API

L’ensemble des données du site est disponible en `json` en lecture seule: 

- https://eglises-ouvertes.ch/index.json

Un `json` est produit par canton, par exemple:

- https://eglises-ouvertes.ch/berne/index.json
- https://eglises-ouvertes.ch/jura/index.json

Le `hotlinking` (raisonnable) est possible.

## Compilation

Le site est généré par Hugo: https://gohugo.io/ (version 0.91 minimale)

### Calcul des distances (nouveau)

Le calcul des distances est effectué par un petit programme.
Il se trouve à la racine du projet.

Windows:

    distances.exe

Linux:

    ./distances

Mac (Intel):

    ./distances-intel.app

Mac (M1):

    ./distances-m1.app

La source est disponible en Go (`distances.go`).
Elle peut être lancée en direct (`go run distances.go`) ou compilée pour d’autres plateformes.

### Calcul des distances (en js)

**L’utilisation de la méthode ci-dessus est recommandée.**

Le calcul des distances est effectué par le script `distances.js` à la racine du projet.
Il requiert l’installation de `node.js`.
`npm install` permet d’installer `geolib`.

Pour lancer le calcul:

    node distances.js

### Calcul des distances (en Python)

**L’utilisation de la première méthode est recommandée.**

Un script de calcul en Python existe.
Il est lancé par 

    ./distances.py

Évidemment, il exige un Python fonctionnel sur l’ordinateur.

### Compilation complète

- `hugo`  
  (crée notamment le fichier `/public/index.json`)
- `./distances.exe` ou `./distances` ou `./distances-intel.app` ou `./distances-m1.app`  
  (utilise le fichier `/public/index.json` et crée le fichier `/data/distances.json`)
- `hugo`  
  (utilise le fichier `/data/distances.json`)

## Fichiers spéciaux

- `/data/vitraux.json` permet d’entrer des dates pour les artistes
- `/config/_default/config.yaml` permet d’éditer le titre du site et le pied de page

## Layouts

- `/layouts/_default/list.html` est utilisé pour les pages cantonales
- `/layouts/_default/single.html` est utilisé pour les pages d’églises
- `/layouts/about/search.html` est utilisé pour la page de recherche

## Shortcodes

- `cantons.html` renvoie une liste des cantons du site
- `pbc.html` doit être utilisé avec un paramètres (`A` ou `B`)
- `sitemap.html` renvoie une liste de toutes les pages du site
- `vitraux.html` renvoie une liste des artistes et églises (utilise: `/data/vitraux.json`)

## Réutilisation

L’ensemble des fichiers peut être réutilisé librement, sauf:

- `fuse.js` pour la recherche
- *Fraunces*, la police d’écriture des titres
- ~~le code d’image `opengraph` chez Cloudinary~~
- ~~`geolib` à installer avant calcul des distances pour l’ancienne version en `js`~~

Merci de vérifier les licences ;)