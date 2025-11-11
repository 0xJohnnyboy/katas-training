# Karate Chop (Binary Search)

## Règles

Implémenter une recherche binaire (binary chop) dans un tableau trié d'entiers.

**Fonction à implémenter:**
```
chop(target, array) → index
```

**Comportement:**
- Cherche `target` dans `array` trié
- Retourne l'index de `target` si trouvé
- Retourne `-1` si `target` n'est pas dans le tableau
- Le tableau est toujours trié par ordre croissant

**Défi:**
Implémenter cette fonction de 5 façons complètement différentes :
1. Itérative classique
2. Récursive
3. Avec slicing
4. Avec des pointeurs/indices
5. Fonctionnelle

## Cas de test

### Tableau vide

| target | array | résultat |
|--------|-------|----------|
| 3      | []    | -1       |
| 0      | []    | -1       |

### Un élément

| target | array | résultat |
|--------|-------|----------|
| 1      | [1]   | 0        |
| 3      | [1]   | -1       |

### Plusieurs éléments - trouvé

| target | array           | résultat | position    |
|--------|-----------------|----------|-------------|
| 1      | [1, 3, 5]       | 0        | début       |
| 3      | [1, 3, 5]       | 1        | milieu      |
| 5      | [1, 3, 5]       | 2        | fin         |
| 1      | [1, 3, 5, 7]    | 0        | début       |
| 3      | [1, 3, 5, 7]    | 1        | milieu-g    |
| 5      | [1, 3, 5, 7]    | 2        | milieu-d    |
| 7      | [1, 3, 5, 7]    | 3        | fin         |

### Plusieurs éléments - non trouvé

| target | array           | résultat | raison              |
|--------|-----------------|----------|---------------------|
| 0      | [1, 3, 5]       | -1       | < minimum           |
| 2      | [1, 3, 5]       | -1       | entre 1 et 3        |
| 4      | [1, 3, 5]       | -1       | entre 3 et 5        |
| 6      | [1, 3, 5]       | -1       | > maximum           |
| 0      | [1, 3, 5, 7]    | -1       | < minimum           |
| 2      | [1, 3, 5, 7]    | -1       | entre éléments      |
| 4      | [1, 3, 5, 7]    | -1       | entre éléments      |
| 6      | [1, 3, 5, 7]    | -1       | entre éléments      |
| 8      | [1, 3, 5, 7]    | -1       | > maximum           |

### Tableaux plus grands

| target | array                          | résultat |
|--------|--------------------------------|----------|
| 1      | [1, 2, 3, 4, 5, 6, 7, 8, 9]   | 0        |
| 5      | [1, 2, 3, 4, 5, 6, 7, 8, 9]   | 4        |
| 9      | [1, 2, 3, 4, 5, 6, 7, 8, 9]   | 8        |
| 0      | [1, 2, 3, 4, 5, 6, 7, 8, 9]   | -1       |
| 10     | [1, 2, 3, 4, 5, 6, 7, 8, 9]   | -1       |
| 3      | [1, 2, 3, 4, 5, 6, 7, 8, 9]   | 2        |
| 7      | [1, 2, 3, 4, 5, 6, 7, 8, 9]   | 6        |

### Nombres négatifs

| target | array              | résultat |
|--------|--------------------|----------|
| -5     | [-5, -3, -1, 0, 2] | 0        |
| -3     | [-5, -3, -1, 0, 2] | 1        |
| 0      | [-5, -3, -1, 0, 2] | 3        |
| 2      | [-5, -3, -1, 0, 2] | 4        |
| -4     | [-5, -3, -1, 0, 2] | -1       |
| 1      | [-5, -3, -1, 0, 2] | -1       |

### Doublons (si autorisés)

Si le tableau contient des doublons, retourner n'importe quel index valide :

| target | array              | résultat possible |
|--------|--------------------|-------------------|
| 3      | [1, 3, 3, 3, 5]    | 1, 2 ou 3         |
| 5      | [1, 3, 5, 5, 7]    | 2 ou 3            |

### Cas limites

| target | array                    | résultat | note               |
|--------|--------------------------|----------|--------------------|
| 1      | [1]                      | 0        | tableau à 1 élément|
| 1      | [1, 1]                   | 0 ou 1   | tous identiques    |
| 50     | [10, 20, 30, 40, 50]     | 4        | multiples de 10    |
| 100    | [10, 20, 30, 40, 50]     | -1       | au-delà            |
