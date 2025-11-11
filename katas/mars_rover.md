# Mars Rover

## Règles

Un rover se déplace sur une grille 2D. Il a une position (x, y) et une direction (N, E, S, W).

**Commandes:**
- `L` : tourner à gauche
- `R` : tourner à droite
- `F` : avancer d'une case dans la direction
- `B` : reculer d'une case

**Contraintes:**
- La grille est toroïdale (wrapping) : sortir d'un côté = réapparaître de l'autre
- Des obstacles peuvent bloquer le déplacement
- Si un obstacle est rencontré, le rover s'arrête avant l'obstacle

## Cas de test

### Rotations

Direction initiale → Commande → Direction finale

| Initial | Commande | Final |
|---------|----------|-------|
| N       | L        | W     |
| N       | R        | E     |
| E       | L        | N     |
| E       | R        | S     |
| S       | L        | E     |
| S       | R        | W     |
| W       | L        | S     |
| W       | R        | N     |

### Mouvements simples

Grille 5x5, position initiale (2, 2)

| Direction | Commande | Position finale |
|-----------|----------|-----------------|
| N         | F        | (2, 1)          |
| N         | B        | (2, 3)          |
| E         | F        | (3, 2)          |
| E         | B        | (1, 2)          |
| S         | F        | (2, 3)          |
| S         | B        | (2, 1)          |
| W         | F        | (1, 2)          |
| W         | B        | (3, 2)          |

### Séquences de commandes

Grille 5x5, position initiale (2, 2, N)

| Séquence       | Position finale |
|----------------|-----------------|
| LFLFRFRFF      | (0, 2, N)       |
| FFRFF          | (2, 0, N)       |
| RFFLFFLF       | (3, 1, N)       |

### Wrapping

Grille 5x5

| Position | Direction | Commande | Position finale |
|----------|-----------|----------|-----------------|
| (0, 0)   | N         | F        | (0, 4)          |
| (0, 0)   | W         | F        | (4, 0)          |
| (4, 4)   | S         | F        | (4, 0)          |
| (4, 4)   | E         | F        | (0, 4)          |

### Obstacles

Grille 5x5, obstacles à (2, 1) et (3, 3)

| Position départ | Direction | Commande | Position finale | Raison               |
|-----------------|-----------|----------|-----------------|----------------------|
| (2, 2)          | N         | F        | (2, 2)          | Obstacle en (2, 1)   |
| (2, 2)          | N         | LFF      | (1, 2)          | Pas d'obstacle       |
| (3, 2)          | N         | F        | (3, 2)          | Obstacle en (3, 3) ? |

Note: Les coordonnées utilisent (x, y) où x=colonne, y=ligne. Origine en haut à gauche.
