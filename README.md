# eglises-ouvertes

Liste d’églises protestantes réformées ouvertes en Suisse romande.

J'ai donné l'ancien nom de domaine eglises-ouvertes.ch. Le site a été ensuite abandonné.

À titre d'archive, il ressemblait à ceci: https://eglises-ouvertes.netlify.app/

## API

L’ensemble des données du site est disponible en `json` en lecture seule: 

- https://eglises-ouvertes.netlify.app/index.json

Un `json` est produit par canton, par exemple:

- https://eglises-ouvertes.netlify.app/berne/index.json
- https://eglises-ouvertes.netlify.app/jura/index.json

Le `hotlinking` (raisonnable) est possible.

## Compilation

Le site est généré par Hugo: https://gohugo.io/ (version 0.91 minimale)

### Calcul des distances

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

Merci de vérifier les licences ;)