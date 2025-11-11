# String Calculator

## Règles

Créer une fonction `Add(string)` qui retourne la somme des nombres dans une chaîne.

### Étapes progressives

1. **Chaîne vide** : retourne 0
2. **Un nombre** : retourne ce nombre
3. **Deux nombres** : séparés par une virgule, retourne leur somme
4. **N nombres** : séparés par des virgules, retourne leur somme
5. **Nouvelle ligne** : accepter `\n` comme séparateur en plus de la virgule
6. **Délimiteur personnalisé** : format `//[délimiteur]\n[nombres]`
7. **Nombres négatifs** : lancer une exception avec le message "negatives not allowed: [liste]"
8. **Ignorer grands nombres** : les nombres > 1000 sont ignorés
9. **Délimiteur multi-caractères** : format `//[***]\n[nombres]`
10. **Multiples délimiteurs** : format `//[*][%]\n[nombres]`

## Cas de test

### Basiques

| Input            | Output | Note                      |
|------------------|--------|---------------------------|
| ""               | 0      | Chaîne vide               |
| "1"              | 1      | Un nombre                 |
| "1,2"            | 3      | Deux nombres              |
| "1,2,3"          | 6      | Plusieurs nombres         |

### Séparateurs

| Input            | Output | Note                      |
|------------------|--------|---------------------------|
| "1\n2,3"         | 6      | Virgule et nouvelle ligne |
| "1,\n"           | Erreur | Séparateur invalide       |

### Délimiteur personnalisé

| Input            | Output | Note                      |
|------------------|--------|---------------------------|
| "//;\n1;2"       | 3      | Point-virgule             |
| "//\|\n1\|2\|3"  | 6      | Pipe                      |
| "//***\n1***2***3" | 6    | Multi-caractères          |
| "//[*][%]\n1*2%3" | 6     | Multiples délimiteurs     |

### Nombres négatifs

| Input            | Exception                           |
|------------------|-------------------------------------|
| "1,-2"           | "negatives not allowed: -2"         |
| "-1,-2,3"        | "negatives not allowed: -1, -2"     |

### Ignorer grands nombres

| Input            | Output | Note                      |
|------------------|--------|---------------------------|
| "2,1001"         | 2      | 1001 est ignoré           |
| "1000,1001,2"    | 1002   | 1000 est inclus           |
