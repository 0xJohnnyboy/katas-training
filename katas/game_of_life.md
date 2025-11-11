# Game of Life

## Règles

Grille 2D où chaque cellule est vivante ou morte. À chaque génération :

1. Une cellule vivante avec 2 ou 3 voisins vivants survit
2. Une cellule vivante avec moins de 2 voisins meurt (sous-population)
3. Une cellule vivante avec plus de 3 voisins meurt (surpopulation)
4. Une cellule morte avec exactement 3 voisins devient vivante (reproduction)

Voisins = 8 cellules adjacentes (horizontales, verticales, diagonales)

## Cas de test

### Évolution d'une cellule

| État actuel | Voisins vivants | État suivant | Raison |
|-------------|-----------------|--------------|---------|
| Vivante     | 0               | Morte        | Sous-population |
| Vivante     | 1               | Morte        | Sous-population |
| Vivante     | 2               | Vivante      | Survie |
| Vivante     | 3               | Vivante      | Survie |
| Vivante     | 4               | Morte        | Surpopulation |
| Vivante     | 5+              | Morte        | Surpopulation |
| Morte       | 2               | Morte        | - |
| Morte       | 3               | Vivante      | Reproduction |
| Morte       | 4               | Morte        | - |

### Patterns classiques

**Blinker (oscillateur - période 2)**
```
Génération 0:     Génération 1:
. . . . .         . . . . .
. . X . .         . . . . .
. . X . .         . X X X .
. . X . .         . . . . .
. . . . .         . . . . .

Positions G0: (2,1) (2,2) (2,3)
Positions G1: (1,2) (2,2) (3,2)
```

**Block (nature morte)**
```
. . . .
. X X .
. X X .
. . . .

Positions: (1,1) (2,1) (1,2) (2,2)
Reste stable à chaque génération
```

**Glider (vaisseau - se déplace)**
```
Génération 0:    Génération 4:
. X . . . .      . . . . . .
. . X . . .      . . X . . .
X X X . . .      . . . X . .
. . . . . .      . X X X . .
. . . . . .      . . . . . .

G0: (1,0) (2,1) (0,2) (1,2) (2,2)
G4: (2,1) (3,2) (1,3) (2,3) (3,3)
Le pattern se déplace d'une case en diagonale toutes les 4 générations
```
