# Pac-Man

## Règles

Reproduire la boucle de jeu de Pac-Man.

**Déplacement de Pac-Man:**
- Avance d'une case par tick en ligne droite
- 4 directions : haut, bas, gauche, droite
- Change de direction instantanément sur commande
- Traverse vers le côté opposé aux bords du plateau
- Bloqué par les murs

**Collecte et score:**
- Pacgums : toutes les cases non-murs (+10 points)
- Super-pacgums : cases spéciales (+50 points, effraie fantômes 24 ticks)
- Victoire : tous les pacgums mangés

**Fantômes mangés** (sous effet super-pacgum) :
- 1er : 200 pts, 2e : 400 pts, 3e : 800 pts, 4e : 1600 pts
- Compteur réinitialisé après le 4e

**Défaite:**
- Collision avec fantôme actif = perte d'une vie
- Réapparition au point de départ
- Jeu terminé si plus de vies

## Modes des fantômes

Chaque fantôme alterne entre 4 modes :

1. **Home** : rotation au point de départ (durée variable par fantôme)
2. **Scattered** : 28 ticks vers case assignée
3. **Chase** : 80 ticks avec IA spécifique
4. **Frightened** : 24 ticks en fuite (effet super-pacgum)
5. **Eaten** : retour au point de départ après capture

## Comportements par fantôme

**Blinky (rouge):**
- Pas de mode Home
- Chase : cible directement la position de Pac-Man

**Pinky (rose):**
- Home : 12 ticks
- Chase : cible 4 cases devant Pac-Man

**Inky (cyan):**
- Home : 24 ticks
- Chase : cible position symétrique complexe

**Clyde (orange):**
- Home : 36 ticks
- Chase : cible Pac-Man si distance > 8, sinon passe en Scattered

## Cas de test

### Déplacement basique

Grille 5x5, Pac-Man en (2, 2), direction E

| Tick | Position | Direction | Commande | Prochaine pos |
|------|----------|-----------|----------|---------------|
| 0    | (2, 2)   | E         | -        | (3, 2)        |
| 1    | (3, 2)   | E         | -        | (4, 2)        |
| 2    | (4, 2)   | E         | N        | (4, 1)        |
| 3    | (4, 1)   | N         | -        | (4, 0)        |

### Wrapping

Grille 5x5 (largeur 0-4)

| Position | Direction | Prochaine position |
|----------|-----------|---------------------|
| (4, 2)   | E         | (0, 2)              |
| (0, 2)   | W         | (4, 2)              |
| (2, 0)   | N         | (2, 4)              |
| (2, 4)   | S         | (2, 0)              |

### Collision avec mur

```
#####
#P..#
#.#.#
#...#
#####

P en (1,1), direction E
- Tick 0: (1,1) → (2,1) OK
- Tick 1: (2,1) → (2,2) mur → reste (2,1)
```

### Score et pacgums

| Événement                  | Points | Total |
|----------------------------|--------|-------|
| Départ                     | 0      | 0     |
| Mange pacgum               | +10    | 10    |
| Mange 3 pacgums            | +30    | 40    |
| Mange super-pacgum         | +50    | 90    |
| Mange fantôme (1er)        | +200   | 290   |
| Mange fantôme (2e)         | +400   | 690   |
| Effet termine, mange autre super-pacgum | +50 | 740 |
| Mange fantôme (1er après reset) | +200 | 940 |

### Modes fantôme (Blinky)

| Tick | Mode      | Durée restante | Cible               |
|------|-----------|----------------|---------------------|
| 0-27 | Scattered | 28→1           | Case assignée       |
| 28-107| Chase    | 80→1           | Position Pac-Man    |
| 108-135| Scattered| 28→1          | Case assignée       |
| N    | Frightened| 24             | Fuite (si super-pacgum mangé) |

### Effet super-pacgum

```
Tick 0: Pac-Man mange super-pacgum
Tick 0-23: Fantômes en mode Frightened (24 ticks)
Tick 10: Pac-Man mange Blinky → 200 pts, Blinky en mode Eaten
Tick 12: Pac-Man mange Pinky → 400 pts
Tick 24: Fin de l'effet, fantômes reprennent leur cycle normal
```
