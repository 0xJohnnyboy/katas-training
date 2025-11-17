# Worker Pool (Pool de Workers)

## Règles

Implémenter un pool de workers qui traite des jobs en parallèle.

**Architecture :**
- Queue de jobs (channel)
- N workers qui consomment les jobs
- Collecte des résultats
- Attente de la fin de tous les jobs

**Composants :**
```
Jobs → [Worker 1]
    → [Worker 2] → Results
    → [Worker 3]
```

**Contraintes :**
- Nombre de workers configurable
- Jobs peuvent échouer (gestion d'erreurs)
- Attendre que tous les jobs soient terminés
- Pouvoir annuler tous les jobs en cours

## Cas de test

### Pool basique

| Jobs          | Workers | Résultats attendus     |
|---------------|---------|------------------------|
| [1, 2, 3]     | 1       | [2, 4, 6]              |
| [1, 2, 3, 4]  | 2       | [2, 4, 6, 8]           |
| [1..10]       | 3       | [2, 4, 6... 20]        |
| []            | 5       | []                     |

Fonction de test : doubler chaque nombre

### Jobs avec durée variable

Jobs : télécharger URLs avec temps variable

| Job ID | Durée (ms) | Workers | Temps total max |
|--------|------------|---------|-----------------|
| 5 jobs | 100ms      | 1       | ~500ms          |
| 5 jobs | 100ms      | 5       | ~100ms          |
| 10 jobs| 100ms      | 2       | ~500ms          |
| 10 jobs| 100ms      | 5       | ~200ms          |

### Gestion d'erreurs

Jobs : diviser 100 par N

| Jobs           | Résultats                          |
|----------------|------------------------------------|
| [10, 5, 0, 2]  | [10, 20, error, 50]                |
| [0, 0, 0]      | [error, error, error]              |
| [1, 2, 3]      | [100, 50, 33]                      |

### Annulation (Context)

```
Lancer 100 jobs de 1s chacun avec 5 workers
Annuler après 500ms
Résultat : ~5 jobs complétés, 95 annulés
```

### Ordre des résultats

| Mode        | Jobs      | Résultats possibles      |
|-------------|-----------|--------------------------|
| Non ordonné | [1,2,3,4] | [2,4,6,8] ou [4,2,8,6]   |
| Ordonné     | [1,2,3,4] | [2,4,6,8] (toujours)     |

### Utilisation mémoire

Pool avec buffer limité :

| Jobs | Buffer taille | Comportement                    |
|------|---------------|---------------------------------|
| 100  | 10            | Max 10 jobs en attente          |
| 1000 | 50            | Bloque après 50 jobs non traités|

## Exemples de jobs

### Simple : calcul mathématique
```
Job: number int
Traitement: number * 2
Résultat: int
```

### Réaliste : téléchargement URL
```
Job: {url string, id int}
Traitement: HTTP GET + parse
Résultat: {id int, data []byte, error error}
```

### Complexe : traitement de fichiers
```
Job: filepath string
Traitement: lire + compresser + sauvegarder
Résultat: {filepath string, size int, error error}
```

## Métriques à suivre

| Métrique                | Description                        |
|-------------------------|------------------------------------|
| Jobs complétés          | Nombre de jobs terminés avec succès|
| Jobs échoués            | Nombre de jobs avec erreur         |
| Temps moyen par job     | Durée moyenne de traitement        |
| Workers actifs          | Nombre de workers occupés          |
| Queue size              | Jobs en attente                    |

## Concepts de concurrence

- **sync.WaitGroup** : attendre tous les workers
- **Buffered channel** : queue de jobs
- **Worker goroutines** : pool de workers
- **Result channel** : collecter résultats
- **context.Context** : annulation
- **select** : timeout et annulation
