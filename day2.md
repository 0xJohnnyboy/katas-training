# Jour 2 - MEZZO : Patterns & Méthodologie

> **Durée** : 4-6 heures  
> **Objectif** : Maîtriser les patterns techniques, le TDD en Go, et réaliser 2 katas guidés

---

## 🎯 Objectifs de la Journée

À la fin de ce jour, tu dois pouvoir :
- ✅ Implémenter une grille 2D en Go
- ✅ Écrire des tests table-driven efficaces
- ✅ Suivre le cycle TDD (Red/Green/Refactor) naturellement
- ✅ Avoir fait String Calculator et Tennis en mode guidé

---

## 🛠️ Partie 1 : Setup Go & Testing (30min)

### Structure de Projet Kata

```
kata-project/
├── go.mod
├── string_calculator/
│   ├── calculator.go
│   └── calculator_test.go
├── tennis/
│   ├── game.go
│   └── game_test.go
└── mars_rover/
    ├── rover.go
    ├── position.go
    └── rover_test.go
```

### Initialisation

```bash
mkdir kata-training && cd kata-training
go mod init github.com/ton-username/kata-training

# Optionnel : testify pour des assertions plus lisibles
go get github.com/stretchr/testify
```

### Template de Test Go

```go
package example

import "testing"

func TestSomething(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"cas 1", 1, 1},
        {"cas 2", 2, 4},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Function(tt.input)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}
```

### Commandes Essentielles

```bash
# Lancer tous les tests
go test ./...

# Lancer avec verbose
go test -v ./...

# Lancer un package spécifique
go test ./string_calculator

# Coverage
go test -cover ./...
```

---

## 🧩 Partie 2 : Patterns Techniques Détaillés (90min)

### Pattern 2.1 : Grille 2D - Position

**Implémentation recommandée** :

```go
package grid

// Position représente des coordonnées 2D
type Position struct {
    X int
    Y int
}

// NewPosition crée une nouvelle position
func NewPosition(x, y int) Position {
    return Position{X: x, Y: y}
}

// Move retourne une nouvelle position déplacée (immutable)
func (p Position) Move(dx, dy int) Position {
    return Position{
        X: p.X + dx,
        Y: p.Y + dy,
    }
}

// Equals compare deux positions
func (p Position) Equals(other Position) bool {
    return p.X == other.X && p.Y == other.Y
}
```

**Tests** :

```go
func TestPosition(t *testing.T) {
    t.Run("creates position", func(t *testing.T) {
        p := NewPosition(3, 5)
        if p.X != 3 || p.Y != 5 {
            t.Errorf("expected (3,5), got (%d,%d)", p.X, p.Y)
        }
    })
    
    t.Run("moves position", func(t *testing.T) {
        p := NewPosition(2, 3)
        moved := p.Move(1, -1)
        
        if moved.X != 3 || moved.Y != 2 {
            t.Errorf("expected (3,2), got (%d,%d)", moved.X, moved.Y)
        }
        
        // Vérifier l'immutabilité
        if p.X != 2 || p.Y != 3 {
            t.Error("original position was modified")
        }
    })
}
```

---

### Pattern 2.2 : Directions & Orientations

**Pour Mars Rover, Sailboat, etc.**

```go
package grid

type Direction int

const (
    North Direction = iota
    East
    South
    West
)

// String pour affichage
func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}

// TurnLeft tourne de 90° à gauche
func (d Direction) TurnLeft() Direction {
    return Direction((int(d) + 3) % 4) // -1 en modulo 4
}

// TurnRight tourne de 90° à droite
func (d Direction) TurnRight() Direction {
    return Direction((int(d) + 1) % 4)
}

// Delta retourne le déplacement (dx, dy) pour cette direction
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

**Tests** :

```go
func TestDirection(t *testing.T) {
    t.Run("turn left from North", func(t *testing.T) {
        dir := North
        turned := dir.TurnLeft()
        if turned != West {
            t.Errorf("expected West, got %v", turned)
        }
    })
    
    t.Run("turn right from North", func(t *testing.T) {
        dir := North
        turned := dir.TurnRight()
        if turned != East {
            t.Errorf("expected East, got %v", turned)
        }
    })
    
    t.Run("delta for each direction", func(t *testing.T) {
        tests := []struct {
            dir      Direction
            expectedDx int
            expectedDy int
        }{
            {North, 0, 1},
            {East, 1, 0},
            {South, 0, -1},
            {West, -1, 0},
        }
        
        for _, tt := range tests {
            dx, dy := tt.dir.Delta()
            if dx != tt.expectedDx || dy != tt.expectedDy {
                t.Errorf("%v: expected (%d,%d), got (%d,%d)", 
                    tt.dir, tt.expectedDx, tt.expectedDy, dx, dy)
            }
        }
    })
}
```

---

### Pattern 2.3 : Grille avec Wrapping

**Pour gérer les bords de la grille** :

```go
package grid

type Grid struct {
    Width  int
    Height int
}

func NewGrid(width, height int) Grid {
    return Grid{Width: width, Height: height}
}

// Wrap applique le wrapping (grille toroïdale)
func (g Grid) Wrap(pos Position) Position {
    x := ((pos.X % g.Width) + g.Width) % g.Width
    y := ((pos.Y % g.Height) + g.Height) % g.Height
    return Position{X: x, Y: y}
}

// IsInBounds vérifie si une position est dans la grille
func (g Grid) IsInBounds(pos Position) bool {
    return pos.X >= 0 && pos.X < g.Width &&
           pos.Y >= 0 && pos.Y < g.Height
}
```

**Tests** :

```go
func TestGrid(t *testing.T) {
    grid := NewGrid(5, 5)
    
    t.Run("wraps around right edge", func(t *testing.T) {
        pos := Position{X: 6, Y: 2}
        wrapped := grid.Wrap(pos)
        if wrapped.X != 1 || wrapped.Y != 2 {
            t.Errorf("expected (1,2), got (%d,%d)", wrapped.X, wrapped.Y)
        }
    })
    
    t.Run("wraps around negative", func(t *testing.T) {
        pos := Position{X: -1, Y: 2}
        wrapped := grid.Wrap(pos)
        if wrapped.X != 4 || wrapped.Y != 2 {
            t.Errorf("expected (4,2), got (%d,%d)", wrapped.X, wrapped.Y)
        }
    })
}
```

---

### Pattern 2.4 : Voisins (pour Game of Life)

```go
// Neighbors retourne les 8 voisins d'une position
func (p Position) Neighbors() []Position {
    neighbors := make([]Position, 0, 8)
    
    for dx := -1; dx <= 1; dx++ {
        for dy := -1; dy <= 1; dy++ {
            if dx == 0 && dy == 0 {
                continue // skip self
            }
            neighbors = append(neighbors, Position{
                X: p.X + dx,
                Y: p.Y + dy,
            })
        }
    }
    
    return neighbors
}
```

---

## 🧪 Partie 3 : Méthodologie TDD (45min)

### Le Cycle Red-Green-Refactor

```
🔴 RED    : Écrire un test qui échoue
  ↓
🟢 GREEN  : Écrire le code minimum pour le faire passer
  ↓
♻️ REFACTOR : Améliorer le code sans changer le comportement
  ↓
(répéter)
```

### Les Baby Steps

**Principe** : Avancer par le plus petit incrément possible.

**Exemple** pour String Calculator :
1. Test : `"" → 0`
2. Test : `"1" → 1`
3. Test : `"1,2" → 3`
4. Test : `"1,2,5" → 8`

**❌ À éviter** : Vouloir traiter tous les cas d'un coup.

### La Règle des 3

Avant de généraliser (abstraire, factoriser), attends d'avoir vu le pattern **3 fois**.

**Exemple** :
```go
// Après 2 fois : on garde la duplication
if input == "" { return 0 }
if input == "1" { return 1 }

// À la 3ème : on peut généraliser
parts := strings.Split(input, ",")
sum := 0
for _, part := range parts {
    // ...
}
```

---

## 🎯 Partie 4 : Kata Guidé - String Calculator (60-90min)

### Spécifications

Créer une fonction `Add(input string) int` :

1. ✅ String vide retourne 0
2. ✅ Un seul nombre retourne ce nombre
3. ✅ Deux nombres séparés par virgule retournent leur somme
4. ✅ N nombres séparés par virgule
5. ✅ Support de `\n` comme séparateur (en plus de la virgule)
6. ✅ Les nombres négatifs lancent une exception

### Étape 1 : Setup

```bash
mkdir string_calculator
cd string_calculator
touch calculator.go calculator_test.go
```

### Étape 2 : Premier Test (RED)

**calculator_test.go** :
```go
package stringcalculator

import "testing"

func TestAdd(t *testing.T) {
    t.Run("empty string returns zero", func(t *testing.T) {
        result := Add("")
        if result != 0 {
            t.Errorf("expected 0, got %d", result)
        }
    })
}
```

**Lancer** : `go test` → ❌ Erreur de compilation (fonction n'existe pas)

### Étape 3 : Code Minimum (GREEN)

**calculator.go** :
```go
package stringcalculator

func Add(input string) int {
    return 0
}
```

**Lancer** : `go test` → ✅ Test passe !

### Étape 4 : Deuxième Test (RED)

**calculator_test.go** :
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected int
    }{
        {"empty string", "", 0},
        {"single number", "1", 1}, // ← Nouveau
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.input)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}
```

**Lancer** : `go test` → ❌ Expected 1, got 0

### Étape 5 : Implémentation Naïve (GREEN)

```go
package stringcalculator

import "strconv"

func Add(input string) int {
    if input == "" {
        return 0
    }
    
    num, _ := strconv.Atoi(input)
    return num
}
```

**Lancer** : `go test` → ✅

### Étape 6 : Troisième Test (RED)

```go
{"two numbers", "1,2", 3},
```

**Lancer** : `go test` → ❌ Expected 3, got 1

### Étape 7 : Généralisation (GREEN)

```go
package stringcalculator

import (
    "strconv"
    "strings"
)

func Add(input string) int {
    if input == "" {
        return 0
    }
    
    parts := strings.Split(input, ",")
    sum := 0
    
    for _, part := range parts {
        num, _ := strconv.Atoi(part)
        sum += num
    }
    
    return sum
}
```

**Lancer** : `go test` → ✅

### Étape 8 : Tests Supplémentaires

```go
{"multiple numbers", "1,2,5,10", 18},
{"with newlines", "1\n2,3", 6},
```

### Étape 9 : Support de `\n` (GREEN)

```go
func Add(input string) int {
    if input == "" {
        return 0
    }
    
    // Remplacer \n par , pour uniformiser
    input = strings.ReplaceAll(input, "\n", ",")
    
    parts := strings.Split(input, ",")
    sum := 0
    
    for _, part := range parts {
        if part == "" {
            continue
        }
        num, _ := strconv.Atoi(part)
        sum += num
    }
    
    return sum
}
```

### Étape 10 : Nombres Négatifs (Exception)

**Test** :
```go
t.Run("negative numbers throw", func(t *testing.T) {
    _, err := AddWithError("-1,2")
    if err == nil {
        t.Error("expected error for negative numbers")
    }
})
```

**Implémentation** :
```go
func AddWithError(input string) (int, error) {
    if input == "" {
        return 0, nil
    }
    
    input = strings.ReplaceAll(input, "\n", ",")
    parts := strings.Split(input, ",")
    sum := 0
    negatives := []string{}
    
    for _, part := range parts {
        if part == "" {
            continue
        }
        num, _ := strconv.Atoi(part)
        if num < 0 {
            negatives = append(negatives, part)
        }
        sum += num
    }
    
    if len(negatives) > 0 {
        return 0, fmt.Errorf("negatives not allowed: %s", 
            strings.Join(negatives, ","))
    }
    
    return sum, nil
}
```

### ♻️ Étape 11 : Refactoring Final

**Extraire la logique de parsing** :

```go
func parseNumbers(input string) []int {
    if input == "" {
        return []int{}
    }
    
    input = strings.ReplaceAll(input, "\n", ",")
    parts := strings.Split(input, ",")
    numbers := []int{}
    
    for _, part := range parts {
        if part == "" {
            continue
        }
        num, _ := strconv.Atoi(part)
        numbers = append(numbers, num)
    }
    
    return numbers
}

func Add(input string) (int, error) {
    numbers := parseNumbers(input)
    
    sum := 0
    negatives := []int{}
    
    for _, num := range numbers {
        if num < 0 {
            negatives = append(negatives, num)
        }
        sum += num
    }
    
    if len(negatives) > 0 {
        return 0, fmt.Errorf("negatives not allowed: %v", negatives)
    }
    
    return sum, nil
}
```

---

## 🎾 Partie 5 : Kata Guidé - Tennis (90min)

### Spécifications

Modéliser un match de tennis avec la fonction `Score()` qui retourne le score actuel.

**Règles** :
- Points : 0 (love), 15, 30, 40
- Si 40-40 → "Deuce"
- Après deuce, un joueur mène → "Advantage [Player]"
- Gagner : depuis 40-0/15/30 OU depuis advantage

### Structure de Départ

```go
package tennis

type Player int

const (
    Player1 Player = iota
    Player2
)

type Game struct {
    score1 int
    score2 int
}

func NewGame() *Game {
    return &Game{}
}

func (g *Game) PointWonBy(player Player) {
    // TODO
}

func (g *Game) Score() string {
    // TODO
}
```

### Étape 1 : Premiers Tests

```go
func TestTennis(t *testing.T) {
    t.Run("initial score is love-all", func(t *testing.T) {
        game := NewGame()
        if game.Score() != "love-all" {
            t.Errorf("expected love-all, got %s", game.Score())
        }
    })
    
    t.Run("player1 scores once", func(t *testing.T) {
        game := NewGame()
        game.PointWonBy(Player1)
        if game.Score() != "15-love" {
            t.Errorf("expected 15-love, got %s", game.Score())
        }
    })
}
```

### Étape 2 : Implémentation Basique

```go
func (g *Game) PointWonBy(player Player) {
    if player == Player1 {
        g.score1++
    } else {
        g.score2++
    }
}

func (g *Game) Score() string {
    scoreNames := []string{"love", "15", "30", "40"}
    
    score1Str := scoreNames[g.score1]
    score2Str := scoreNames[g.score2]
    
    if g.score1 == g.score2 {
        if g.score1 == 0 {
            return "love-all"
        }
        return score1Str + "-all"
    }
    
    return score1Str + "-" + score2Str
}
```

### Étape 3 : Tests pour Deuce et Advantage

```go
t.Run("deuce at 40-40", func(t *testing.T) {
    game := NewGame()
    // Score 3 points each
    for i := 0; i < 3; i++ {
        game.PointWonBy(Player1)
        game.PointWonBy(Player2)
    }
    if game.Score() != "deuce" {
        t.Errorf("expected deuce, got %s", game.Score())
    }
})

t.Run("advantage player1", func(t *testing.T) {
    game := NewGame()
    for i := 0; i < 3; i++ {
        game.PointWonBy(Player1)
        game.PointWonBy(Player2)
    }
    game.PointWonBy(Player1)
    if game.Score() != "advantage player1" {
        t.Errorf("expected advantage player1, got %s", game.Score())
    }
})
```

### Étape 4 : Gestion Deuce/Advantage

```go
func (g *Game) Score() string {
    scoreNames := []string{"love", "15", "30", "40"}
    
    // Gestion deuce et advantage
    if g.score1 >= 3 && g.score2 >= 3 {
        diff := g.score1 - g.score2
        if diff == 0 {
            return "deuce"
        } else if diff == 1 {
            return "advantage player1"
        } else if diff == -1 {
            return "advantage player2"
        } else if diff >= 2 {
            return "player1 wins"
        } else {
            return "player2 wins"
        }
    }
    
    // Victoire avant deuce
    if g.score1 >= 4 {
        return "player1 wins"
    }
    if g.score2 >= 4 {
        return "player2 wins"
    }
    
    // Score normal
    score1Str := scoreNames[g.score1]
    score2Str := scoreNames[g.score2]
    
    if g.score1 == g.score2 {
        if g.score1 == 0 {
            return "love-all"
        }
        return score1Str + "-all"
    }
    
    return score1Str + "-" + score2Str
}
```

### ♻️ Étape 5 : Refactoring

**Extraire les cas** :

```go
func (g *Game) Score() string {
    if g.isDeuce() {
        return g.deuceScore()
    }
    
    if g.hasWinner() {
        return g.winnerScore()
    }
    
    return g.normalScore()
}

func (g *Game) isDeuce() bool {
    return g.score1 >= 3 && g.score2 >= 3
}

func (g *Game) deuceScore() string {
    diff := g.score1 - g.score2
    switch diff {
    case 0:
        return "deuce"
    case 1:
        return "advantage player1"
    case -1:
        return "advantage player2"
    default:
        return g.winnerScore()
    }
}

func (g *Game) hasWinner() bool {
    return g.score1 >= 4 || g.score2 >= 4
}

func (g *Game) winnerScore() string {
    if g.score1 > g.score2 {
        return "player1 wins"
    }
    return "player2 wins"
}

func (g *Game) normalScore() string {
    scoreNames := []string{"love", "15", "30", "40"}
    score1Str := scoreNames[g.score1]
    score2Str := scoreNames[g.score2]
    
    if g.score1 == g.score2 {
        if g.score1 == 0 {
            return "love-all"
        }
        return score1Str + "-all"
    }
    
    return score1Str + "-" + score2Str
}
```

---

## ✅ Checklist Jour 2

Avant de passer au Jour 3, vérifie :

- [ ] J'ai configuré mon environnement Go avec tests
- [ ] Je sais écrire des tests table-driven
- [ ] J'ai compris le cycle Red-Green-Refactor
- [ ] J'ai fait String Calculator en TDD du début à la fin
- [ ] J'ai fait Tennis en TDD
- [ ] Je sais implémenter une Position immutable
- [ ] Je sais faire tourner une Direction (left/right)
- [ ] Je comprends le wrapping de grille

---

## 📝 Notes Personnelles

```
Qu'est-ce qui m'a posé problème aujourd'hui ?


Quel pattern ai-je trouvé le plus utile ?


Suis-je à l'aise avec le TDD maintenant ?


Points à revoir demain :

```

---

## 🚀 Prochaine Étape

Demain, on attaque Mars Rover et Game of Life en mode intensif !

→ [Jours 3-4 : MICRO - Pratique Intensive](./day3_4.md)
