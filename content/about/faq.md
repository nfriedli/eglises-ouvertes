---
title: Foire aux questions (FAQ)
linkTitle: Questions-réponses
description: Une liste de questions-réponses pour essayer d’expliquer les enjeux et les choix du projet eglises-ouvertes.ch.
weight: 30
---

Une liste de questions-réponses pour essayer d’expliquer les enjeux et les choix du projet `eglises-ouvertes.ch`.
N’hésitez pas à en poser d’autres afin que cette foire aux questions s’améliore.

{{< toc >}}

## Qu’est-ce qu’une «église ouverte»?

La définition utilisée ici est simple:

> Une «église ouverte» est un *bâtiment* qu’il est possible de visiter librement et gratuitement, hors de toute manifestation cultuelle ou culturelle.

Ce site ne mentionne pas les *manifestation* religieuses ou spirituelles appelées «Églises ouvertes» (souvent avec majuscule).

## Quel est le but du projet?

Ce site est né de la frustration de ne pas toujours trouver les informations utiles dans durant mes voyages.
Le nom de domaine a été réservé le 26 octobre 2021.

Des listes d’églises et des indications d’ouverture existent en ligne. 
Toutefois, elles sont rarement systématiques et leur navigation est souvent pénible, notamment sur un téléphone.

Ce site est destiné aux personnes qui se déplacent ou cherchent des bâtiment à visiter:

- il présente des édifices ouverts
- il dépasse les frontières cantonales

## Pourquoi la chapelle de X et le temple de Y ne sont pas dans la liste?

Il est regrettable que des églises ouvertes réformées de Suisse romande n’apparaissent pas sur ce site.
Quelques explications possibles:

- l’église n’a pas été trouvée
- la page officielle et Google Maps de présentent pas d’horaire
- le temps a manqué au responsable du site

## Comment ajouter une église ouverte?

Il suffit d’envoyer un message au responsable du site, en indiquant:

- le nom du lieu
- le lien vers sa page officielle (avec horaire d’ouverture) ou la mention d’une fiche Google Maps (avec horaires d’ouverture)
- toute autre information utile (par exemple une inscription à l’inventaire PBC)

Les églises ouvertes doivent présenter des **horaires simples et compréhensibles**. 
L’ouverture doit être 7/7 et les heures identiques tous les jours.

Si l’horaire n’est présent ni sur la page officielle, ni sur Google Maps, l’édifice ne sera pas publié.

## Comment transmettre des informations précises?

Idéalement, les données sont transmises au format `json`.
Les édifices sont saisis selon ce modèle:

```
"title":        "Temple de Petauchnok",
"site":         "https://www.petaouchnok.ch/temple/",
"maps":         "https://goo.gl/maps/3KSuvNTqXsdSFmnc9",
"rue":          "Rue de Petropavlovsk 666",
"NPA":          1234,
"localite":     "Petaouchnok",
"commune":      "St-Creux-des-Bas-Fonds",
"canton":       "Perpette",
"latitude":     46.9666066,
"longitude":    6.8623137,
"ouverture":    "7/7",
"horaire":      "8:00-20:00"
"pbc":          "B",
"vitraux":      ["Michael Meyers", "Jamie Lee Curtis"]
```

Mais toute transmission d’informations sera utile!

## Quel tarif pour publier «mon» bâtiment?

La publication de «votre» église ouverte est gratuite.
Certains sites similaires facturent plusieurs centaines d’euros pour une présence en ligne.
Le travail fourni pour `eglises-ouvertes.ch` est bénévole et réalisé en toute indépendance.

## Comment modifier une page existante?

Vous ne pouvez pas modifier une page, mais vous pouvez suggérer des modifications à l’auteur de site.
Merci, en particulier, de **signaler au plus vite** des changements d’horaires d’ouverture ou des erreurs.

## Pourquoi certains édifices ne proposent pas de liens vers une page web?

L’absence de lien vers une page web de présentation peut signifier:

- qu’il n’existe pas de présentation en ligne (ou qu’elle n’a pas été trouvée) ou
- que la page ne présente pas d’horaires d’ouverture

Il est possible de proposer une [bonne page de présentation d’une église](/about/site/) en se concentrant sur l’essentiel.
Merci pour les internautes qui la liront.

## Pourquoi certains bâtiments ne renvoient pas à Google Maps?

L’absence de lien vers une fiche Google Maps peut signifier:

- que l’église ne question n’existe pas dans Google Maps
- que la fiche Google Maps ne présente pas d’horaires ou qu’elle est trop sommaire
- qu’il y a une incohérence entre le site officiel et la fiche Google Maps

Dans tous les cas, il est conseillé d’[améliorer sa fiche Google Maps](/about/google/).
C’est possible que l’on en soit propriétaire ou non.

## Pourquoi ne pas proposer systématiquement un lien vers la paroisse responsable d’une église?

Personne ne sait *pourquoi* les internautes souhaitent trouver une église ouverte. 
Intérêt architectural? 
Passion pour l’art du vitrail? 
Recherche d’un lieu de concert? 
Souhait d’un temps de méditation personnelle?

Pourtant, un lien vers un site paroissial est toujours proposé si la page de présentation de l’édifice de qualité (et mentionne des horaires d’ouverture).

Pour les recherches spécifiquement paroissiales, il existe le site [Trouver ma paroisse](https://ma-paroisse.ch/).

## Pourquoi se limiter aux Églises protestantes réformées de Suisse romande?

Parce que le *corpus* semble raisonnable et gérable pour un *webmaster* dilettante et bénévole.
Il est bien assez difficile de lister toutes les églises ouvertes protestantes réformées, souvent mal documentées, pour ne pas viser plus grand.

## Où trouver des informations sur les Églises ouvertes?

Plusieurs documentations de qualité existent:

- ~~[brochure EERV](https://vcc.eerv.ch/wp-content/uploads/sites/287/2015/02/170920_Brochure_Eglises_Ouvertes_LIGHT_V00.pdf)~~ (n’est plus disponible en ligne)
- [brochure RefBeJuSo](https://www.refbejuso.ch/fr/activites/eglise-lieu-daccueil/comment-devenir-une-eglise-lieu-daccueil/)
- [brochure EERS](https://www.evref.ch/fr/publications/accueil-et-confiance-pour-des-eglises-ouvertes/)

Merci d’enrichir la liste par vos suggestions.

## Pourquoi proposer des contenus de Berne et du Jura alors qu’un excellent site local existe?

[Visite d’église](https://visitedeglise.ch/) a une vocation patrimoniale. 
C’est un site remarquable par sa richesse.
Par sa visée particulière, il n’est pas optimal pour parcourir des églises ouvertes de proche en proche.

## Pourquoi le site ne présente pas de carte?

La carte n’est pas forcément le meilleur moyen de découvrir de nouvelles églises ouvertes.
C’est souvent de l’esbroufe.
Sa pratique peut être difficile sur un smartphone.
Au fond, l’application Google Maps est souvent meilleure que des cartes intégrées.


## Pourquoi est-il si sobre? 

C’est un site d’information, qui peut être utilisé par quiconque sur un *smartphone*.
Il est volontairement simple et léger.

## Pourquoi ce site est-il sans images?

Il n’est pas possible de trouver des images libres de droits de tous les édifices présentés.
Mais la fourniture d’images libre de qualité pourrait faire changer le responsable d’avis.

Toutefois, des images sont automatiquement générées pour les partages sur les réseaux sociaux.

## Pourquoi ce site est aussi rapide?

Le site `eglises-ouvertes.ch` est servi par Netlify, compilé par génial [Hugo](https://gohugo.io/) et maintenu sur [GitHub](https://github.com/nfriedli/eglises-ouvertes).
Il respecte au mieux les standards HTML et CSS.
Il vise les meilleurs scores d’accessibilité (a11y) WAVE.
Il devrait obtenir de très bons résultats de performances au test Lighthouse.

## Comment réutiliser les données automatiquement?

Les **données sont disponibles** à l’adresse [`/index.json`](/index.json).
Elles sont toujours actualisées et le *hotlinking* est permis.
Des fichiers `json` sur mesure ou une API peuvent être mise à disposition en cas de demande précise.

Des fichiers `json` par canton sont disponibles en ajoutant `index.json` à l’URL.

## Comment soutenir l’initiative?

C’est tout simple! En publiant un lien vers `eglises-ouvertes.ch` sur votre propre site. 
Vous pouvez pointer vers la page d’accueil ou vers une page précise, au choix.
N’hésitez pas à partager des pages sur les réseaux sociaux, dans un mail ou par une application de messagerie.

## Qui remercier pour ce site?

Pour leur soutien, leurs avis et leurs contributions, merci à:

- Constantin Bacha, pasteur
- [Maëlle Bader](https://www.referguel.ch/paroisses/courtelary-cormoret/), pasteure
- [Joël Burri](https://www.reformes.ch/blog/joel-burri), journaliste
- [Diane Friedli](https://dianefriedli.ch/), pasteure
- [Elio Jaillet](https://eliojaillet.ch/), théologien
- [David Kneubühler](https://www.referguel.ch/paroisses/corgemont-cortebert/), pasteur
- [Jean-Marc Leresche](https://jeanmarcleresche.ch/), diacre

Pour leurs outils libres et performants, merci à:

- [Bjørn Erik Pedersen](https://bep.is/en/), créateur de Hugo, générateur de site statique utilisé actuellement
- [Kiro Risk](https://fusejs.io/team.html), créateur de fuse.js, script de recherche
- [Zach Leatherman](https://www.zachleat.com/), créateur d’Eleventy (11ty), générateur utilisé pour la première version du site
