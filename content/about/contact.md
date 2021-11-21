---
title: Contacter le responsable du site
linkTitle: Contact
description: Spécificités du projet, contact de l’auteur du site, licence des contenus & informations techniques.
weight: 1000
---

Pourquoi lancer un nouveau site?
Qui contacter pour des questions, des remarques ou pour proposer une nouvelle église?
Quel est le statut des contenus publiés?

## Le projet

Ce site est né de la frustration de ne pas toujours trouver les informations utiles dans durant mes voyages.
À la suite de mon article [Églises ouvertes (ou pas)](https://theologique.ch/eglises-ouvertes-ou-pas/), j’ai acheté le nom de domaine `eglises-ouvertes.ch` le 26 octobre 2021.

Des listes d’églises et des indications d’ouverture existent en ligne. 
Toutefois, elles sont rarement systématiques et leur navigation est souvent pénible, notamment sur un téléphone.

Ce site est destiné aux personnes qui se déplacent ou cherchent des bâtiment à visiter:

- il présente des édifices ouverts
- il dépasse les frontières cantonales

Le beau succès de [Trouver ma paroisse](https://ma-paroisse.ch/) me fait penser qu’un site efficace (simple, précis & pertinent) est possible dans le domaine des bâtiments protestants réformés.

## Contact

Le propriétaire et responsable bénévole du site est:

Nicolas Friedli  
Rue du Château 3  
2013 Colombier (NE)  
Suisse

[nicolas.friedli+eo@gmail.com](mailto:nicolas.friedli+eo@gmail.com)  
[+41 32 841 48 74](tel:+41328414874) (urgences uniquement)  
[+41 79 344 33 82](tel:+41793443382) (Threema, Signal, SMS, Telegram, ~~WhatsApp~~)

<script type="application/ld+json">
{
"@context": "http://www.schema.org",
"@type": "person",
"name": "Nicolas Friedli",
"jobTitle": "Consultant web indépendant",
"gender": "male",
"url": "https://nicolasfriedli.ch/",
"sameAs": [
"https://theologique.ch/",
"https://frdl.ch",
"https://ma-paroisse.ch/",
"https://eglises-ouvertes.ch"
],
"image": "https://frdl.ch/nicolas-friedli.jpg",
"address": {
"@type": "PostalAddress",
"streetAddress": "Rue du Château 3",
"addressLocality": "Colombier",
"addressRegion": "NE",
"postalCode": "2013",
"addressCountry": "Suisse"
},
"email": "nicolasfriedli@protonmail.ch",
"alumniOf": "Université de Neuchâtel",
"nationality": "Suisse",
"telephone": "+41328414874"
}
</script>

Je suis disponible pour vous aider à:
- [gérer votre fiche Google Maps](/about/google/) ou
- [évaluer votre page de présentation sur votre site](/about/site/)

## Licence

Sauf mention contraire explicite, tout le contenu de ce site est placé sous licence Creative Commons Zero (CC0).
Elle est équivalente au domaine public. 
Le code source sera mis à disposition.

Le script `fuse.js` utilisé pour la recherche et la police d’écriture sont sous d’autres licences libres.
Merci d’aller les chercher sur leurs sites propres.

## Technique

Les édifices sont saisis en format `json`, selon ce modèle:

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

Les **données sont disponibles** à l’adresse [`/index.json`](/index.json).
Elles sont toujours actualisées et le *hotlinking* est permis.
Je suis prêt à générer des `json` sur mesure ou à créer une API en cas de demande précise.

Le site `eglises-ouvertes.ch` est servi par Netlify, compilé par génial [Hugo](https://gohugo.io/) et maintenu sur GitHub.

Il respecte au mieux les standards HTML et CSS.
Il vise les meilleurs scores d’accessibilité (a11y) WAVE.

Il devrait obtenir de très bons résultats de performances au test Lighthouse.

Le site utilise Google Analytics pour obtenir quelques informations de navigation.
Il utilise la polices libre *Fraunces*.
La recherche utilise le splendide `fuse.js`.


## Remerciements

Pour leur soutien, leurs avis, leurs contributions et leurs outils libres, merci à:

- [Maëlle Bader](https://www.referguel.ch/paroisses/courtelary-cormoret/)
- [Joël Burri](https://www.reformes.ch/blog/joel-burri)
- [Diane Friedli](https://dianefriedli.ch/)
- [David Kneubühler](https://www.referguel.ch/paroisses/corgemont-cortebert/)
- [Bjørn Erik Pedersen](https://bep.is/en/)
- [Kiro Risk](https://fusejs.io/team.html)
- [Zach Leatherman](https://www.zachleat.com/)