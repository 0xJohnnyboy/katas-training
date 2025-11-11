# Reversi (Othello)

## Règles

Jeu de stratégie sur plateau 8x8 avec pions bicolores (noir/blanc).

**Configuration initiale:**
```
. . . . . . . .
. . . . . . . .
. . . . . . . .
. . . W B . . .
. . . B W . . .
. . . . . . . .
. . . . . . . .
. . . . . . . .
```
Centre : 4 pions en diagonale (2 blancs, 2 noirs)
Le joueur noir commence.

**Coup légal:**
Une case est jouable si elle forme une ligne continue de pions adverses terminée par un pion allié dans au moins une des 8 directions (↑ ↓ ← → ↖ ↗ ↙ ↘).

**Capture:**
Placer un pion retourne automatiquement tous les pions adverses encadrés dans toutes les directions valides simultanément.

**Passer son tour:**
Si aucun coup légal n'existe, le joueur passe.

**Fin de partie:**
- Aucun joueur ne peut jouer, OU
- Plateau complètement rempli

**Victoire:**
Le joueur avec le plus de pions de sa couleur.

## Cas de test

### Coups légaux - Position initiale

```
. . . . . . . .
. . . . . . . .
. . . . . . . .
. . . W B . . .
. . . B W . . .
. . . . . . . .
. . . . . . . .
. . . . . . . .
```

Coups légaux pour Noir :

| Position | Direction | Ligne encadrée | Valide |
|----------|-----------|----------------|--------|
| (3, 2)   | ↓         | W en (3,3)     | Oui    |
| (2, 3)   | →         | W en (3,3)     | Oui    |
| (4, 5)   | ↑         | W en (4,4)     | Oui    |
| (5, 4)   | ←         | W en (4,4)     | Oui    |
| (3, 3)   | -         | Occupé         | Non    |
| (0, 0)   | -         | Aucune ligne   | Non    |

### Placement et captures

Noir joue en (2, 3) :
```
Avant:               Après:
. . . . . . . .      . . . . . . . .
. . . . . . . .      . . . . . . . .
. . . . . . . .      . . . . . . . .
. . . W B . . .      . . B B B . . .
. . . B W . . .      . . . B W . . .
. . . . . . . .      . . . . . . . .
. . . . . . . .      . . . . . . . .
. . . . . . . .      . . . . . . . .

W en (3,3) retourné en B (encadré entre (2,3) et (4,3))
```

### Captures multiples directions

```
Avant:               Après coup B en (2,2):
. . . . . . . .      . . . . . . . .
. . . . . . . .      . . . . . . . .
. . B . . . . .      . . B . . . . .
. . W W B . . .      . . B B B . . .
. . W B W . . .      . . B B W . . .
. . W . . . . .      . . B . . . . .
. . . . . . . .      . . . . . . . .
. . . . . . . .      . . . . . . . .

Capture dans 3 directions : → (W en 3,3), ↓ (W en 2,3 et 2,4), ↘ (W en 3,4)
```

### Aucun coup légal - Passer

```
B B B B B B B B
B B B B B B B B
B B B B B B B B
B B B B B B B W
B B B B B B W .
B B B B B W . .
B B B B W . . .
B B B B . . . .
```

Blanc ne peut pas jouer (aucune case ne forme de ligne valide) → passe son tour
Noir joue à nouveau

### Fin de partie

```
B B B B B B B B
B B B B B B B B
B B B B B W W W
B B B B W W W W
B B W W W W W W
B W W W W W W W
W W W W W W W W
W W W W W W W W

Plateau rempli
Noir : 18 pions
Blanc : 46 pions
Victoire : Blanc
```

### Cas d'égalité

```
B B B B W W W W
B B B B W W W W
B B B B W W W W
B B B B W W W W
B B B B W W W W
B B B B W W W W
B B B B W W W W
B B B B W W W W

Noir : 32 pions
Blanc : 32 pions
Résultat : Égalité
```
