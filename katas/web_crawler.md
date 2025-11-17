# Web Crawler

## Règles

Crawler un site web en parallèle en suivant les liens.

**Contraintes :**
- Crawler plusieurs pages en parallèle
- Ne pas crawler la même URL deux fois
- Limiter la profondeur de crawl
- Limiter le nombre de goroutines actives
- Extraire les liens de chaque page
- Gérer les erreurs (404, timeout, etc.)

**Algorithme :**
1. Commencer avec URL de départ
2. Télécharger la page
3. Extraire tous les liens
4. Ajouter liens non visités à la queue
5. Répéter jusqu'à profondeur max ou queue vide

## Cas de test

### Crawl simple

Site avec structure :
```
A → B, C
B → D
C → D, E
D → F
E → (rien)
F → (rien)
```

| URL départ | Profondeur max | URLs visitées |
|------------|----------------|---------------|
| A          | 0              | [A]           |
| A          | 1              | [A, B, C]     |
| A          | 2              | [A, B, C, D, E] |
| A          | 3              | [A, B, C, D, E, F] |
| A          | 999            | [A, B, C, D, E, F] |

### Détection de cycles

```
A → B
B → C
C → A  (cycle)
```

| URL départ | Profondeur max | URLs visitées | Visites |
|------------|----------------|---------------|---------|
| A          | 10             | [A, B, C]     | A:1, B:1, C:1 |

Chaque URL ne doit être visitée qu'une seule fois malgré le cycle.

### Graphe avec liens multiples

```
A → B, C, D
B → C
C → D
D → E
```

Toutes les URLs doivent être visitées exactement une fois :
- A découvre B, C, D
- B découvre C (déjà vu, ignoré)
- C découvre D (déjà vu, ignoré)
- etc.

### Limites de concurrence

Crawler 100 URLs avec limite de 10 workers :

| Mesure                  | Valeur attendue  |
|-------------------------|------------------|
| Workers max simultanés  | 10               |
| URLs crawlées           | 100              |
| Durée si séquentiel     | ~100s (1s/URL)   |
| Durée avec 10 workers   | ~10s             |

### Gestion d'erreurs

| URL          | Résultat  | Comportement                |
|--------------|-----------|-----------------------------|
| /valid       | 200 OK    | Crawler les liens           |
| /notfound    | 404       | Logger erreur, continuer    |
| /timeout     | Timeout   | Logger erreur, continuer    |
| /servererror | 500       | Logger erreur, continuer    |

### Filtrage de domaine

URL de départ : `https://example.com/page1`

| Lien trouvé                    | Crawler ? | Raison                |
|--------------------------------|-----------|-----------------------|
| /page2                         | Oui       | Même domaine          |
| https://example.com/page3      | Oui       | Même domaine          |
| https://other.com/page         | Non       | Domaine différent     |
| https://sub.example.com/page   | Non       | Sous-domaine différent|

### Résultats attendus

Pour chaque URL crawlée, collecter :

```
{
    url: string
    statusCode: int
    links: []string
    depth: int
    error: error (si échec)
}
```

## Exemples de scénarios

### Petit site (5 pages)
```
home → about, contact
about → team
contact → (rien)
team → (rien)

Résultat: 4 URLs visitées
```

### Site avec pagination
```
/posts?page=1 → /posts?page=2, /post/1, /post/2
/posts?page=2 → /posts?page=3, /post/3, /post/4
...
```

Limiter profondeur ou nombre max de pages.

### Site réel simulé

100 pages organisées en :
- 1 homepage → 10 catégories
- Chaque catégorie → 10 articles
- Chaque article → liens vers 3 articles reliés

Total : ~100 pages uniques avec beaucoup de duplication de liens.

## Métriques

| Métrique              | Description                          |
|-----------------------|--------------------------------------|
| URLs visitées         | Nombre total d'URLs crawlées         |
| URLs en erreur        | 404, 500, timeout, etc.              |
| Profondeur moyenne    | Profondeur moyenne des pages         |
| Profondeur max        | Page la plus profonde atteinte       |
| Durée totale          | Temps total de crawl                 |
| Vitesse               | URLs/seconde                         |
| Liens trouvés         | Total de liens extraits              |
| Liens uniques         | Liens uniques après déduplication    |

## Concepts de concurrence

- **sync.Map** ou **mutex** : URLs visitées (accès concurrent)
- **Worker pool** : limiter goroutines actives
- **Channels** : queue d'URLs à crawler
- **sync.WaitGroup** : attendre fin du crawl
- **context.Context** : timeout global et annulation
- **Rate limiting** : ne pas surcharger le serveur
