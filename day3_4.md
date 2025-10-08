# Jours 3-4 - MICRO : Pratique Intensive

> **Durée** : 2 jours (6-8h par jour)  
> **Objectif** : Maîtriser Mars Rover et Game of Life par la répétition

---

## 🎯 Objectifs des Journées

À la fin de ces 2 jours, tu dois pouvoir :
- ✅ Faire Mars Rover en moins de 60 minutes
- ✅ Démarrer un kata sans hésitation
- ✅ Automatiser le cycle TDD
- ✅ Avoir pratiqué au moins 5-6 itérations complètes

---

## 📋 Planning Suggéré

### Jour 3
- **9h-10h30** : Mars Rover - Tentative 1 (guidée)
- **10h30-11h** : Pause + Débriefing
- **11h-12h30** : Mars Rover - Tentative 2 (autonome)
- **12h30-14h** : Pause déjeuner
- **14h-15h30** : Game of Life - Tentative 1 (guidée)
- **15h30-16h** : Pause
- **16h-17h30** : Game of Life - Tentative 2

### Jour 4
- **9h-10h** : Mars Rover - Tentative 3 (sous contraintes)
- **10h-11h** : Sailboat - Découverte
- **11h-12h30** : Mars Rover - Tentative finale (chrono)
- **12h30-14h** : Pause déjeuner
- **14h-15h** : Game of Life - Tentative 3
- **15h-16h** : Kata au choix (révision)
- **16h-17h** : Synthèse et préparation J5

---

## 🚀 Mars Rover - Guide Complet

### Spécifications Complètes

**Contexte** : Un rover se déplace sur Mars (grille 2D rectangulaire).

**État du Rover** :
- Position (x, y)
- Orientation (N, S, E, W)

**Commandes** :
- `F` : Avancer d'une case (forward)
- `B` : Reculer d'une case (backward)
- `L` : Tourner 90° à gauche
- `R` : Tourner 90° à droite

**Contraintes** :
- La grille est "toroïdale" (wrapping : si on sort par un bord, on réapparaît de l'autre côté)
- Les obstacles bloquent le mouvement

**Format de sortie** :
- Position initiale : `"2:3:N"` → x=2, y=3, face Nord
- Après commandes : `"4:1:E"` → nouvelle position

---

### 🎯 Mars Rover - Tentative 1 (Guidée - 90min)

#### Étape 1 : Setup & Premier Test

**Structure** :
```
mars_rover/
├── rover.go
├── position.go
├── direction.go
└── rover_test.go
```

**rover_test.go** :
```go
package marsrover

import "testing"

func TestRover(t *testing.T) {
    t.Run("creates rover with initial position", func(t *testing.T) {
        rover := NewRover(2, 3, North)
        
        if rover.X != 2 || rover.Y != 3 {
            t.Errorf("expected position (2,3), got (%d,%d)", 
                rover.X, rover.Y)
        }
        if rover.Direction != North {
            t.Errorf("expected North, got %v", rover.Direction)
        }
    })
}
```

**rover.go** :
```go
package marsrover

type Rover struct {
    X         int
    Y         int
    Direction Direction
}

func NewRover(x, y int, dir Direction) *Rover {
    return &Rover{X: x, Y: y, Direction: dir}
}
```

**direction.go** (réutilise le code du Jour 2) :
```go
package marsrover

type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"N", "E", "S", "W"}[d]
}

func (d Direction) TurnLeft() Direction {
    return Direction((int(d) + 3) % 4)
}

func (d Direction) TurnRight() Direction {
    return Direction((int(d) + 1) % 4)
}

func (d Direction) Delta() (int, int) {
    deltas := map[Direction][2]int{
        North: {0, 1},
        East:  {1, 0},
        South: {0, -1},
        West:  {-1, 0},
    }
    delta := deltas[d]
    return delta[0], delta[1]
}
```

---

#### Étape 2 : Commandes de Rotation

**Test** :
```go
t.Run("turn left", func(t *testing.T) {
    rover := NewRover(0, 0, North)
    rover.Execute("L")
    
    if rover.Direction != West {
        t.Errorf("expected West, got %v", rover.Direction)
    }
})

t.Run("turn right", func(t *testing.T) {
    rover := NewRover(0, 0, North)
    rover.Execute("R")
    
    if rover.Direction != East {
        t.Errorf("expected East, got %v", rover.Direction)
    }
})
```

**Implémentation** :
```go
func (r *Rover) Execute(commands string) {
    for _, cmd := range commands {
        switch cmd {
        case 'L':
            r.Direction = r.Direction.TurnLeft()
        case 'R':
            r.Direction = r.Direction.TurnRight()
        }
    }
}
```

---

#### Étape 3 : Mouvement Forward

**Test** :
```go
t.Run("move forward facing North", func(t *testing.T) {
    rover := NewRover(0, 0, North)
    rover.Execute("F")
    
    if rover.X != 0 || rover.Y != 1 {
        t.Errorf("expected (0,1), got (%d,%d)", rover.X, rover.Y)
    }
})

t.Run("move forward facing East", func(t *testing.T) {
    rover := NewRover(0, 0, East)
    rover.Execute("F")
    
    if rover.X != 1 || rover.Y != 0 {
        t.Errorf("expected (1,0), got (%d,%d)", rover.X, rover.Y)
    }
})
```

**Implémentation** :
```go
func (r *Rover) Execute(commands string) {
    for _, cmd := range commands {
        switch cmd {
        case 'L':
            r.Direction = r.Direction.TurnLeft()
        case 'R':
            r.Direction = r.Direction.TurnRight()
        case 'F':
            dx, dy := r.Direction.Delta()
            r.X += dx
            r.Y += dy
        }
    }
}
```

---

#### Étape 4 : Mouvement Backward

**Test** :
```go
t.Run("move backward facing North", func(t *testing.T) {
    rover := NewRover(0, 2, North)
    rover.Execute("B")
    
    if rover.X != 0 || rover.Y != 1 {
        t.Errorf("expected (0,1), got (%d,%d)", rover.X, rover.Y)
    }
})
```

**Implémentation** :
```go
case 'B':
    dx, dy := r.Direction.Delta()
    r.X -= dx
    r.Y -= dy
```

---

#### Étape 5 : Séquence de Commandes

**Test** :
```go
t.Run("sequence of commands", func(t *testing.T) {
    rover := NewRover(2, 2, North)
    rover.Execute("FFRFF")
    
    if rover.X != 4 || rover.Y != 4 || rover.Direction != East {
        t.Errorf("expected (4,4,E), got (%d,%d,%v)", 
            rover.X, rover.Y, rover.Direction)
    }
})
```

Le code actuel devrait déjà fonctionner !

---

#### Étape 6 : Wrapping (Grille Toroïdale)

**Test** :
```go
t.Run("wraps around right edge", func(t *testing.T) {
    rover := NewRover(9, 5, East, 10, 10) // grille 10x10
    rover.Execute("F")
    
    if rover.X != 0 || rover.Y != 5 {
        t.Errorf("expected (0,5), got (%d,%d)", rover.X, rover.Y)
    }
})
```

**Modification du Rover** :
```go
type Rover struct {
    X          int
    Y          int
    Direction  Direction
    GridWidth  int
    GridHeight int
}

func NewRover(x, y int, dir Direction, width, height int) *Rover {
    return &Rover{
        X: x, Y: y, Direction: dir,
        GridWidth: width, GridHeight: height,
    }
}

func (r *Rover) wrap() {
    r.X = ((r.X % r.GridWidth) + r.GridWidth) % r.GridWidth
    r.Y = ((r.Y % r.GridHeight) + r.GridHeight) % r.GridHeight
}

func (r *Rover) Execute(commands string) {
    for _, cmd := range commands {
        switch cmd {
        case 'L':
            r.Direction = r.Direction.TurnLeft()
        case 'R':
            r.Direction = r.Direction.TurnRight()
        case 'F':
            dx, dy := r.Direction.Delta()
            r.X += dx
            r.Y += dy
            r.wrap()
        case 'B':
            dx, dy := r.Direction.Delta()
            r.X -= dx
            r.Y -= dy
            r.wrap()
        }
    }
}
```

---

#### Étape 7 : Obstacles

**Test** :
```go
t.Run("stops at obstacle", func(t *testing.T) {
    obstacles := []Position{{X: 0, Y: 2}}
    rover := NewRover(0, 0, North, 10, 10)
    rover.SetObstacles(obstacles)
    
    rover.Execute("FFF")
    
    // Devrait s'arrêter à (0,1) car obstacle en (0,2)
    if rover.X != 0 || rover.Y != 1 {
        t.Errorf("expected (0,1), got (%d,%d)", rover.X, rover.Y)
    }
})
```

**Implémentation** :
```go
type Position struct {
    X int
    Y int
}

type Rover struct {
    X          int
    Y          int
    Direction  Direction
    GridWidth  int
    GridHeight int
    Obstacles  map[Position]bool
}

func NewRover(x, y int, dir Direction, width, height int) *Rover {
    return &Rover{
        X: x, Y: y, Direction: dir,
        GridWidth: width, GridHeight: height,
        Obstacles: make(map[Position]bool),
    }
}

func (r *Rover) SetObstacles(obstacles []Position) {
    for _, obs := range obstacles {
        r.Obstacles[obs] = true
    }
}

func (r *Rover) hasObstacle(x, y int) bool {
    return r.Obstacles[Position{X: x, Y: y}]
}

func (r *Rover) Execute(commands string) {
    for _, cmd := range commands {
        switch cmd {
        case 'L':
            r.Direction = r.Direction.TurnLeft()
        case 'R':
            r.Direction = r.Direction.TurnRight()
        case 'F':
            dx, dy := r.Direction.Delta()
            newX := r.X + dx
            newY := r.Y + dy
            r.wrapCoords(&newX, &newY)
            
            if !r.hasObstacle(newX, newY) {
                r.X = newX
                r.Y = newY
            }
        case 'B':
            dx, dy := r.Direction.Delta()
            newX := r.X - dx
            newY := r.Y - dy
            r.wrapCoords(&newX, &newY)
            
            if !r.hasObstacle(newX, newY) {
                r.X = newX
                r.Y = newY
            }
        }
    }
}

func (r *Rover) wrapCoords(x, y *int) {
    *x = ((*x % r.GridWidth) + r.GridWidth) % r.GridWidth
    *y = ((*y % r.GridHeight) + r.GridHeight) % r.GridHeight
}
```

---

#### ♻️ Étape 8 : Refactoring

**Extraire la logique de mouvement** :

```go
func (r *Rover) Execute(commands string) {
    for _, cmd := range commands {
        switch cmd {
        case 'L':
            r.Direction = r.Direction.TurnLeft()
        case 'R':
            r.Direction = r.Direction.TurnRight()
        case 'F':
            r.moveForward()
        case 'B':
            r.moveBackward()
        }
    }
}

func (r *Rover) moveForward() {
    dx, dy := r.Direction.Delta()
    r.tryMove(dx, dy)
}

func (r *Rover) moveBackward() {
    dx, dy := r.Direction.Delta()
    r.tryMove(-dx, -dy)
}

func (r *Rover) tryMove(dx, dy int) {
    newX := r.X + dx
    newY := r.Y + dy
    r.wrapCoords(&newX, &newY)
    
    if !r.hasObstacle(newX, newY) {
        r.X = newX
        r.Y = newY
    }
}
```

---

### 📊 Débriefing Tentative 1

**Questions à te poser** :
- Ai-je commencé par les tests les plus simples ?
- Ai-je fait des baby steps ou des gros sauts ?
- Mon code est-il lisible ?
- Combien de temps ai-je pris ?

**Prends 15 minutes de notes** avant de passer à la tentative 2.

---

### 🔄 Mars Rover - Tentative 2 (Autonome - 60-90min)

**Objectif** : Refaire le kata **du début** (fichiers vides), mais cette fois **sans guide**.

**Règles** :
- ❌ Ne regarde PAS ton code de la tentative 1
- ✅ Essaie de te souvenir de l'ordre des étapes
- ✅ Chronométre-toi
- ✅ Note les difficultés rencontrées

**Checklist des fonctionnalités** :
```markdown
- [ ] Création du rover avec position initiale
- [ ] Rotation left/right
- [ ] Mouvement forward
- [ ] Mouvement backward
- [ ] Séquence de commandes
- [ ] Wrapping
- [ ] Obstacles
```

**Temps cible** : 75 minutes

---

### ⚡ Mars Rover - Tentative 3 (Contraintes - 60min)

**Objectif** : Refaire le kata avec des **contraintes** pour sortir de ta zone de confort.

**Choisis UNE contrainte** :

#### Contrainte 1 : "No Primitives"
Interdiction d'utiliser des types primitifs directement dans les signatures.

**Exemple** :
```go
// ❌ Interdit
func NewRover(x, y int, dir Direction) *Rover

// ✅ Autorisé
func NewRover(pos Position, dir Direction) *Rover
```

#### Contrainte 2 : "No If"
Interdiction d'utiliser `if` (utilise des maps, switch, ou polymorphisme).

**Exemple** :
```go
// ❌ Interdit
if r.hasObstacle(newX, newY) {
    return
}

// ✅ Autorisé avec early return dans switch ou map de fonctions
```

#### Contrainte 3 : "Immutable Rover"
Le rover ne se modifie jamais, chaque commande retourne un nouveau rover.

**Exemple** :
```go
// Au lieu de r.Execute("F")
newRover := r.Execute("F")
```

**Bénéfice** : Ces contraintes forcent à penser différemment et à découvrir de nouveaux patterns.

---

## 🎮 Game of Life - Guide Complet

### Spécifications

**Règles de Conway** :
- Grille 2D de cellules (vivantes ou mortes)
- À chaque génération, chaque cellule évolue selon ses **8 voisins** :
  1. Cellule vivante avec 2-3 voisins vivants → reste vivante
  2. Cellule vivante avec <2 ou >3 voisins → meurt (sous/surpopulation)
  3. Cellule morte avec exactement 3 voisins vivants → naît (reproduction)

**Exemple** :
```
Génération 0:        Génération 1:
. . . . .            . . . . .
. . ■ . .            . . . . .
. . ■ . .    →       . ■ ■ ■ .
. . ■ . .            . . . . .
. . . . .            . . . . .
```

---

### 🎯 Game of Life - Tentative 1 (Guidée - 90min)

#### Étape 1 : Modéliser une Cellule

**Test** :
```go
func TestCell(t *testing.T) {
    t.Run("cell can be alive", func(t *testing.T) {
        cell := NewCell(true)
        if !cell.IsAlive() {
            t.Error("expected cell to be alive")
        }
    })
    
    t.Run("cell can be dead", func(t *testing.T) {
        cell := NewCell(false)
        if cell.IsAlive() {
            t.Error("expected cell to be dead")
        }
    })
}
```

**Implémentation** :
```go
package gameoflife

type Cell struct {
    alive bool
}

func NewCell(alive bool) Cell {
    return Cell{alive: alive}
}

func (c Cell) IsAlive() bool {
    return c.alive
}
```

---

#### Étape 2 : Grille Basique

**Test** :
```go
func TestGrid(t *testing.T) {
    t.Run("creates empty grid", func(t *testing.T) {
        grid := NewGrid(3, 3)
        if grid.Width != 3 || grid.Height != 3 {
            t.Errorf("expected 3x3 grid")
        }
    })
    
    t.Run("sets cell alive", func(t *testing.T) {
        grid := NewGrid(3, 3)
        grid.SetAlive(1, 1)
        
        if !grid.IsAlive(1, 1) {
            t.Error("expected cell at (1,1) to be alive")
        }
    })
}
```

**Implémentation** :
```go
type Position struct {
    X int
    Y int
}

type Grid struct {
    Width  int
    Height int
    Cells  map[Position]bool
}

func NewGrid(width, height int) *Grid {
    return &Grid{
        Width:  width,
        Height: height,
        Cells:  make(map[Position]bool),
    }
}

func (g *Grid) SetAlive(x, y int) {
    g.Cells[Position{X: x, Y: y}] = true
}

func (g *Grid) IsAlive(x, y int) bool {
    return g.Cells[Position{X: x, Y: y}]
}
```

---

#### Étape 3 : Compter les Voisins

**Test** :
```go
func TestCountNeighbors(t *testing.T) {
    t.Run("cell with no neighbors", func(t *testing.T) {
        grid := NewGrid(3, 3)
        grid.SetAlive(1, 1)
        
        count := grid.CountLiveNeighbors(1, 1)
        if count != 0 {
            t.Errorf("expected 0 neighbors, got %d", count)
        }
    })
    
    t.Run("cell with 3 neighbors", func(t *testing.T) {
        grid := NewGrid(3, 3)
        grid.SetAlive(0, 0)
        grid.SetAlive(0, 1)
        grid.SetAlive(1, 0)
        
        count := grid.CountLiveNeighbors(1, 1)
        if count != 3 {
            t.Errorf("expected 3 neighbors, got %d", count)
        }
    })
}
```

**Implémentation** :
```go
func (g *Grid) CountLiveNeighbors(x, y int) int {
    count := 0
    
    for dx := -1; dx <= 1; dx++ {
        for dy := -1; dy <= 1; dy++ {
            if dx == 0 && dy == 0 {
                continue // skip self
            }
            
            nx := x + dx
            ny := y + dy
            
            // Vérifier les bords (grille finie)
            if nx >= 0 && nx < g.Width && ny >= 0 && ny < g.Height {
                if g.IsAlive(nx, ny) {
                    count++
                }
            }
        }
    }
    
    return count
}
```

---

#### Étape 4 : Règles d'Evolution d'une Cellule

**Test** :
```go
func TestCellEvolution(t *testing.T) {
    tests := []struct {
        name            string
        isAlive         bool
        neighborCount   int
        expectedAlive   bool
    }{
        {"alive with 2 neighbors stays alive", true, 2, true},
        {"alive with 3 neighbors stays alive", true, 3, true},
        {"alive with 1 neighbor dies", true, 1, false},
        {"alive with 4 neighbors dies", true, 4, false},
        {"dead with 3 neighbors becomes alive", false, 3, true},
        {"dead with 2 neighbors stays dead", false, 2, false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := WillBeAlive(tt.isAlive, tt.neighborCount)
            if result != tt.expectedAlive {
                t.Errorf("expected %v, got %v", tt.expectedAlive, result)
            }
        })
    }
}
```

**Implémentation** :
```go
func WillBeAlive(currentlyAlive bool, liveNeighbors int) bool {
    if currentlyAlive {
        return liveNeighbors == 2 || liveNeighbors == 3
    }
    return liveNeighbors == 3
}
```

---

#### Étape 5 : Génération Suivante

**Test** :
```go
func TestNextGeneration(t *testing.T) {
    t.Run("blinker pattern", func(t *testing.T) {
        // Pattern "blinker" : ligne horizontale → verticale
        grid := NewGrid(5, 5)
        grid.SetAlive(2, 1)
        grid.SetAlive(2, 2)
        grid.SetAlive(2, 3)
        
        next := grid.NextGeneration()
        
        // Devrait être vertical
        if !next.IsAlive(1, 2) || !next.IsAlive(2, 2) || !next.IsAlive(3, 2) {
            t.Error("expected vertical blinker")
        }
        
        // Ancienne ligne devrait être morte
        if next.IsAlive(2, 1) || next.IsAlive(2, 3) {
            t.Error("expected old cells to be dead")
        }
    })
}
```

**Implémentation** :
```go
func (g *Grid) NextGeneration() *Grid {
    next := NewGrid(g.Width, g.Height)
    
    for x := 0; x < g.Width; x++ {
        for y := 0; y < g.Height; y++ {
            neighbors := g.CountLiveNeighbors(x, y)
            alive := g.IsAlive(x, y)
            
            if WillBeAlive(alive, neighbors) {
                next.SetAlive(x, y)
            }
        }
    }
    
    return next
}
```

---

#### Étape 6 : Patterns Classiques

**Tests pour valider** :

```go
func TestClassicPatterns(t *testing.T) {
    t.Run("block is stable", func(t *testing.T) {
        // Block : carré 2x2 stable
        grid := NewGrid(4, 4)
        grid.SetAlive(1, 1)
        grid.SetAlive(1, 2)
        grid.SetAlive(2, 1)
        grid.SetAlive(2, 2)
        
        next := grid.NextGeneration()
        
        // Devrait rester identique
        if !next.IsAlive(1, 1) || !next.IsAlive(1, 2) ||
           !next.IsAlive(2, 1) || !next.IsAlive(2, 2) {
            t.Error("block should be stable")
        }
    })
    
    t.Run("glider moves", func(t *testing.T) {
        // Glider : pattern qui "bouge"
        grid := NewGrid(6, 6)
        grid.SetAlive(1, 0)
        grid.SetAlive(2, 1)
        grid.SetAlive(0, 2)
        grid.SetAlive(1, 2)
        grid.SetAlive(2, 2)
        
        // Après 4 générations, devrait s'être déplacé
        for i := 0; i < 4; i++ {
            grid = grid.NextGeneration()
        }
        
        // Position déplacée (simplifié, test visuel recommandé)
        if grid.CountLiveCells() != 5 {
            t.Error("glider should maintain 5 cells")
        }
    })
}

func (g *Grid) CountLiveCells() int {
    return len(g.Cells)
}
```

---

#### ♻️ Étape 7 : Refactoring & Visualisation

**Ajouter l'affichage** :

```go
func (g *Grid) String() string {
    var sb strings.Builder
    
    for y := g.Height - 1; y >= 0; y-- {
        for x := 0; x < g.Width; x++ {
            if g.IsAlive(x, y) {
                sb.WriteString("■ ")
            } else {
                sb.WriteString(". ")
            }
        }
        sb.WriteString("\n")
    }
    
    return sb.String()
}
```

**Utilisation** :
```go
func main() {
    grid := NewGrid(10, 10)
    // Setup initial pattern
    grid.SetAlive(5, 5)
    grid.SetAlive(5, 6)
    grid.SetAlive(5, 7)
    
    for i := 0; i < 10; i++ {
        fmt.Printf("Generation %d:\n%s\n", i, grid)
        grid = grid.NextGeneration()
        time.Sleep(500 * time.Millisecond)
    }
}
```

---

### 🔄 Game of Life - Tentatives 2 & 3

**Tentative 2** (Autonome - 60min) :
- Refais le kata sans regarder le code
- Essaie d'optimiser l'ordre des tests
- Chronométre-toi

**Tentative 3** (Avec contrainte - 60min) :
- **Contrainte suggérée** : Grille infinie (utilise une map au lieu de boucler sur width/height)
- Ou : Implémenter avec des cellules actives uniquement (optimisation mémoire)

---

## 🌊 Sailboat - Découverte (Optionnel - 60min)

Si tu as le temps, essaie le kata Sailboat en autonomie.

**Indices** :
- Réutilise tes structures Position et Direction
- La vitesse dépend de l'angle entre le bateau et le vent
- Utilise des produits scalaires/vectoriels pour calculer l'alignement

---

## 📈 Tracking de Progression

| Kata | Tentative | Temps | Complété ? | Qualité Code (1-5) | Notes |
|------|-----------|-------|------------|-------------------|-------|
| Mars Rover | 1 (guidée) | | | | |
| Mars Rover | 2 (autonome) | | | | |
| Mars Rover | 3 (contrainte) | | | | |
| Game of Life | 1 (guidée) | | | | |
| Game of Life | 2 (autonome) | | | | |
| Game of Life | 3 (contrainte) | | | | |
| Sailboat | 1 | | | | |

---

## ✅ Checklist Jours 3-4

Avant de passer au Jour 5 :

- [ ] J'ai fait Mars Rover au moins 2 fois complètement
- [ ] Je peux démarrer Mars Rover sans hésitation
- [ ] Je comprends le wrapping et les obstacles
- [ ] J'ai fait Game of Life au moins 1 fois complètement
- [ ] Je sais compter les voisins d'une cellule
- [ ] Je peux faire un kata en moins de 75 minutes
- [ ] Mon code est lisible et bien structuré

**Temps de référence à viser** :
- Mars Rover : 45-60 minutes
- Game of Life : 60-75 minutes

---

## 💡 Conseils pour la Répétition

### Quand Refaire un Kata ?

**Refais le kata si** :
- Tu as pris >90 minutes
- Ton code est confus/désordonné
- Tu as sauté des étapes TDD
- Tu veux essayer une approche différente

**Passe au suivant si** :
- Tu l'as fait en <60min avec du bon code
- Tu comprends bien les patterns
- Tu veux découvrir d'autres katas

### Comment Améliorer à Chaque Itération ?

1. **Vitesse** : Chronomètre et essaie de battre ton temps
2. **Qualité** : Focus sur le nommage et la structure
3. **Contraintes** : Sors de ta zone de confort
4. **Variantes** : Ajoute des features (visualisation, multi-rovers, etc.)

---

## 📝 Notes Personnelles - Jours 3-4

```
Quel kata ai-je préféré ? Pourquoi ?


Quelle a été ma plus grande difficulté ?


Qu'est-ce qui est devenu automatique ?


Qu'est-ce qui nécessite encore de la pratique ?


Meilleur temps sur Mars Rover : _____ minutes

Meilleur temps sur Game of Life : _____ minutes
```

---

## 🚀 Prochaine Étape

Dernière ligne droite ! Demain, on consolide et on simule un vrai entretien.

→ [Jour 5 : Consolidation & Mock Interview](./day5.md)
