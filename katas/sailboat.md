# Sailboat

## Règles

Simuler un bateau naviguant sur une grille 2D infinie pour atteindre une île. Le vent change à chaque tour.

**Orientation du bateau:**
- Le bateau a une direction (N, E, S, W)
- Peut tourner d'un quart de tour à la fois (horaire ou antihoraire)
- Le navigateur choisit la rotation optimale avant chaque déplacement

**Vitesse selon le vent:**
1. **Aligné avec le vent** : avance de 4 cases
2. **Opposé au vent** : avance de 1 case
3. **Orthogonal au vent** : avance de 2 cases + dévie de 1 case dans la direction du vent

## Cas de test

### Rotations

Direction actuelle → Rotation → Direction finale

| Direction | Rotation horaire | Rotation antihoraire |
|-----------|------------------|----------------------|
| N         | E                | W                    |
| E         | S                | N                    |
| S         | W                | E                    |
| W         | N                | S                    |

### Déplacement aligné avec le vent

Bateau en (0, 0), direction N, vent N

| Direction bateau | Direction vent | Cases avancées | Position finale |
|------------------|----------------|----------------|-----------------|
| N                | N              | 4              | (0, 4)          |
| E                | E              | 4              | (4, 0)          |
| S                | S              | 4              | (0, -4)         |
| W                | W              | 4              | (-4, 0)         |

### Déplacement opposé au vent

Bateau en (0, 0), vent N

| Direction bateau | Direction vent | Cases avancées | Position finale |
|------------------|----------------|----------------|-----------------|
| S                | N              | 1              | (0, -1)         |
| N                | S              | 1              | (0, 1)          |
| W                | E              | 1              | (-1, 0)         |
| E                | W              | 1              | (1, 0)          |

### Déplacement orthogonal au vent

Bateau en (0, 0)

| Direction bateau | Direction vent | Avance | Dérive | Position finale |
|------------------|----------------|--------|--------|-----------------|
| N                | E              | 2N     | 1E     | (1, 2)          |
| N                | W              | 2N     | 1W     | (-1, 2)         |
| E                | N              | 2E     | 1N     | (2, 1)          |
| E                | S              | 2E     | 1S     | (2, -1)         |
| S                | E              | 2S     | 1E     | (1, -2)         |
| S                | W              | 2S     | 1W     | (-1, -2)        |
| W                | N              | 2W     | 1N     | (-2, 1)         |
| W                | S              | 2W     | 1S     | (-2, -1)        |

### Séquence de navigation

Position départ: (0, 0), direction N, destination: (5, 5)

| Tour | Vent | Rotation | Direction après | Déplacement | Position |
|------|------|----------|-----------------|-------------|----------|
| 1    | E    | Horaire  | E               | 4E          | (4, 0)   |
| 2    | N    | Aucune   | E               | 2E + 1N     | (6, 1)   |
| 3    | S    | Aucune   | E               | 2E + 1S     | (8, 0)   |

### Choix de rotation optimal

Bateau en (0, 0), direction N, vent E, destination à l'Est

- **Option 1** : tourner horaire → E (aligné) → avance 4 cases
- **Option 2** : rester N → orthogonal → avance 2N + 1E
- **Optimal** : Option 1 (4 cases vers destination vs 1 case)
