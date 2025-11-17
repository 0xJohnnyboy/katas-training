# Prime Sieve (Crible d'Ératosthène Concurrent)

## Règles

Générer les nombres premiers en utilisant un pipeline de goroutines avec channels.

**Principe :**
- Chaque goroutine filtre les multiples d'un nombre premier
- Les goroutines sont chaînées via channels
- Premier nombre non filtré = nouveau nombre premier

**Architecture :**
```
Générateur → Filtre(2) → Filtre(3) → Filtre(5) → Filtre(7) → ...
```

**Algorithme :**
1. Générateur envoie 2, 3, 4, 5, 6...
2. Premier filtre (2) : bloque multiples de 2, laisse passer impairs
3. Nouveau filtre (3) : bloque multiples de 3
4. Continue jusqu'au nombre cible

## Cas de test

### Nombres premiers jusqu'à N

| N    | Premiers                                           |
|------|----------------------------------------------------|
| 10   | 2, 3, 5, 7                                         |
| 20   | 2, 3, 5, 7, 11, 13, 17, 19                         |
| 30   | 2, 3, 5, 7, 11, 13, 17, 19, 23, 29                 |
| 50   | 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47 |
| 100  | 25 premiers (2, 3, 5, 7, 11... 89, 97)             |

### Les N premiers nombres premiers

| N  | Résultat                                    |
|----|---------------------------------------------|
| 1  | 2                                           |
| 5  | 2, 3, 5, 7, 11                              |
| 10 | 2, 3, 5, 7, 11, 13, 17, 19, 23, 29          |
| 20 | 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71 |

### Filtrage étape par étape

Exemple : Nombres premiers jusqu'à 15

```
Générateur : 2 3 4 5 6 7 8 9 10 11 12 13 14 15

Filtre(2) reçoit: 2 3 4 5 6 7 8 9 10 11 12 13 14 15
Filtre(2) envoie: 2 3 5 7 9 11 13 15  (bloque 4,6,8,10,12,14)

Filtre(3) reçoit: 3 5 7 9 11 13 15
Filtre(3) envoie: 3 5 7 11 13  (bloque 9,15)

Filtre(5) reçoit: 5 7 11 13
Filtre(5) envoie: 5 7 11 13  (rien à bloquer)

Résultat final: 2, 3, 5, 7, 11, 13
```

### Test de non-premiers

Vérifier que ces nombres ne sont PAS générés :

| Nombre | Raison            |
|--------|-------------------|
| 4      | 2 × 2             |
| 6      | 2 × 3             |
| 8      | 2 × 4             |
| 9      | 3 × 3             |
| 10     | 2 × 5             |
| 12     | 2 × 6             |
| 15     | 3 × 5             |
| 20     | 2 × 10            |
| 21     | 3 × 7             |
| 25     | 5 × 5             |

### Cas limites

| N    | Résultat | Note                      |
|------|----------|---------------------------|
| 0    | []       | Aucun premier             |
| 1    | []       | 1 n'est pas premier       |
| 2    | [2]      | Premier nombre premier    |
| 3    | [2, 3]   | Deux premiers             |

## Concepts de concurrence à implémenter

- **Générateur** : goroutine qui produit 2, 3, 4, 5...
- **Filtre** : goroutine qui reçoit un channel, filtre les multiples, envoie au suivant
- **Pipeline** : chaînage de filtres via channels
- **Fermeture** : close() des channels quand terminé
