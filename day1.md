# Jour 1 - MACRO : Comprendre le Contexte

> **Durée** : 3-4 heures  
> **Objectif** : Comprendre la philosophie Shodo, identifier les patterns communs dans les katas, savoir ce qui est évalué

---

## 🎯 Objectifs de la Journée

À la fin de ce jour, tu dois pouvoir :
- ✅ Expliquer ce qu'est un kata et pourquoi Shodo les utilise
- ✅ Lister les 5-6 katas principaux et leur essence
- ✅ Identifier 3-4 patterns techniques qui reviennent
- ✅ Comprendre les critères d'évaluation en entretien

---

## 📖 Partie 1 : Philosophie Software Craftsmanship (30-45min)

### Qu'est-ce qu'un Kata ?

Un **kata** (terme emprunté aux arts martiaux) est un exercice de code que l'on **répète** pour :
- Améliorer sa technique (pas pour résoudre un nouveau problème)
- Pratiquer le TDD
- Expérimenter le refactoring
- Travailler le clean code

**Analogie** : C'est comme un musicien qui fait ses gammes. Ce n'est pas un concert, c'est de l'entraînement.

### Pourquoi Shodo Utilise les Katas ?

Shodo se positionne sur le **Software Craftsmanship** (artisanat logiciel) :
- Focus sur la **qualité du code** et la **démarche**
- Valeurs : apprentissage continu, excellence technique, partage
- Les katas permettent d'évaluer **comment** tu codes, pas juste **ce que** tu codes

### Les 3 Piliers du Software Craftsmanship

1. **TDD (Test-Driven Development)**
   - Red → Green → Refactor
   - Les tests guident le design
   - Confiance dans les refactorings

2. **Clean Code**
   - Nommage explicite
   - Fonctions courtes et ciblées
   - Principe DRY (Don't Repeat Yourself)
   - KISS (Keep It Simple, Stupid)

3. **Refactoring Continu**
   - Améliorer sans changer le comportement
   - Éliminer la duplication
   - Simplifier la complexité

### 💡 Ce qu'on Évalue en Entretien

| ✅ Évalué | ❌ Moins Important |
|----------|-------------------|
| Démarche TDD (écrire les tests d'abord) | Performance optimale |
| Baby steps (petits incréments) | Finir à 100% |
| Communication (expliquer sa pensée) | Connaître tous les design patterns |
| Qualité du code (lisibilité, nommage) | Vitesse brute |
| Capacité à refactorer | Solution la plus élégante du premier coup |

**À retenir** : Mieux vaut un code simple, testé et explicite qu'un code complexe incomplet.

---

## 🗺️ Partie 2 : Tour d'Horizon des Katas (60-90min)

### Kata 1 : String Calculator ⭐⭐ (Débutant)

**Description** : Créer une calculatrice qui additionne des nombres depuis une string.

**Exemple** :
```
"" → 0
"1" → 1
"1,2" → 3
"1,2,5" → 8
```

**Évolutions** :
- Support de `\n` comme séparateur
- Support de délimiteurs personnalisés
- Gestion des nombres négatifs (exception)

**Pourquoi il est bon pour débuter** :
- Spécifications claires et incrémentales
- Force à écrire des tests simples d'abord
- Pas de concepts complexes (pas de grille, pas d'état)

**Compétences travaillées** :
- TDD pur
- Baby steps
- Gestion des edge cases

---

### Kata 2 : Tennis ⭐⭐ (Débutant-Intermédiaire)

**Description** : Calculer le score d'un match de tennis.

**Règles du tennis** :
- Points : 0 ("love"), 15, 30, 40
- Égalité à 40-40 → "Deuce"
- Après deuce : "Advantage" + joueur
- Gagner depuis advantage ou depuis 40-X

**Exemple** :
```
Joueur A marque → "15-love"
Joueur B marque → "15-15"
...
```

**Pièges** :
- Modélisation de l'état du match
- Cas "deuce" et "advantage"

**Compétences travaillées** :
- State machine simple
- Nommage expressif (éviter les `if score == 0`)
- Tests de transitions d'états

---

### Kata 3 : Mars Rover ⭐⭐⭐ (Intermédiaire) - **PRIORITAIRE**

**Description** : Simuler un robot (rover) se déplaçant sur Mars (grille 2D).

**Commandes** :
- `F` : Forward (avancer)
- `B` : Backward (reculer)
- `L` : Left (tourner à gauche)
- `R` : Right (tourner à droite)

**Caractéristiques** :
- Position (x, y)
- Orientation (N, S, E, W)
- Grille avec wrapping (edges → côté opposé)
- Obstacles (arrêt du rover)

**Exemple** :
```
Rover en (2,2) face Nord
Commandes: "FFRFF"
→ Position finale: (2,4) face Nord
```

**Extensions** :
- Détection d'obstacles
- Grille avec visualisation
- Multiple rovers

**Compétences travaillées** :
- Grille 2D
- State (position + orientation)
- Command pattern
- Gestion des transformations de coordonnées

**Pourquoi il est probable chez Shodo** :
- Très utilisé dans les Code Retreats
- Riche en possibilités d'extensions
- Permet d'évaluer plusieurs compétences

---

### Kata 4 : Game of Life ⭐⭐⭐ (Intermédiaire)

**Description** : Implémenter le jeu de la vie de Conway.

**Règles** :
- Grille 2D de cellules (vivantes/mortes)
- Chaque cellule évolue selon ses 8 voisins :
  - Vivante + 2-3 voisins vivants → reste vivante
  - Vivante + <2 ou >3 voisins → meurt
  - Morte + exactement 3 voisins vivants → naît

**Exemple** :
```
Generation 0:    Generation 1:
. . .            . . .
. ■ ■            . ■ ■
. ■ .            . ■ ■
```

**Compétences travaillées** :
- Grille 2D infinie ou finie
- Calcul de voisins
- Transformation d'état global
- Immutabilité (nouvelle génération sans modifier l'ancienne)

---

### Kata 5 : Sailboat ⭐⭐ (Intermédiaire)

**Description** : Simuler un voilier naviguant selon le vent.

**Règles** :
- Grille 2D, voilier avec orientation, île à atteindre
- Vent change chaque tour
- Vitesse dépend de l'alignement bateau/vent :
  - Alignés → 4 cases
  - Opposés → 1 case
  - Orthogonaux → 2 cases + décalage latéral

**Compétences travaillées** :
- Grille 2D (encore)
- Calculs vectoriels simples
- État complexe (position + orientation + vent)

---

### Kata 6 : Gilded Rose ⭐⭐⭐ (Refactoring)

**Description** : Refactorer du legacy code (pas un kata from scratch).

**Contexte** : Système de gestion d'inventaire avec des règles complexes de qualité/péremption.

**Objectif** : Ajouter une nouvelle fonctionnalité sans casser l'existant.

**Compétences travaillées** :
- Refactoring de legacy
- Tests de caractérisation
- Elimination de duplication
- Stratégie de refactoring sécurisé

---

## 🧩 Partie 3 : Identification des Patterns Communs (45-60min)

### Pattern 1 : Grilles 2D

**Présent dans** : Mars Rover, Game of Life, Sailboat

**Concepts clés** :
- Représentation : struct `Position { X, Y int }`
- Opérations : déplacement, voisins, wrapping
- Edge cases : bords, wrapping, grille infinie vs finie

**Structure type en Go** :
```go
type Position struct {
    X int
    Y int
}

type Grid struct {
    Width  int
    Height int
    Cells  map[Position]CellType
}

func (p Position) Move(direction Direction) Position {
    // ...
}

func (p Position) Neighbors() []Position {
    // ...
}
```

**Réflexions design** :
- Séparer Position de l'entité (Rover, Cell)
- Rendre Position immutable
- Wrapping : modulo ou vérifications explicites ?

---

### Pattern 2 : State Machines

**Présent dans** : Tennis, Mars Rover (orientation)

**Concepts clés** :
- État actuel
- Transitions
- Actions déclenchées par événements

**Structure type** :
```go
type State int

const (
    StateNormal State = iota
    StateDeuce
    StateAdvantage
)

type Match struct {
    state       State
    playerScore map[Player]int
}

func (m *Match) PlayerScores(player Player) {
    // Logique de transition
}
```

---

### Pattern 3 : Command Pattern

**Présent dans** : Mars Rover, Sailboat

**Concepts clés** :
- Commande = objet/fonction qui encapsule une action
- Séparation entre parsing des commandes et exécution
- Permet de faire des undo, logs, etc.

**Structure type** :
```go
type Command interface {
    Execute(rover *Rover)
}

type MoveForward struct{}

func (c MoveForward) Execute(rover *Rover) {
    rover.Position = rover.Position.Move(rover.Orientation)
}

func ParseCommands(input string) []Command {
    // ...
}
```

---

### Pattern 4 : Value Objects / Immutabilité

**Présent dans** : Presque tous

**Concepts clés** :
- Objets définis par leurs valeurs, pas leur identité
- Immutables → thread-safe, faciles à tester
- Exemples : Position, Orientation, Score

**Exemple** :
```go
// ✅ Bon : immutable
func (p Position) Move(dx, dy int) Position {
    return Position{X: p.X + dx, Y: p.Y + dy}
}

// ❌ Éviter : mutable
func (p *Position) Move(dx, dy int) {
    p.X += dx
    p.Y += dy
}
```

**Pourquoi** : Évite les effets de bord, facilite le raisonnement et les tests.

---

### Pattern 5 : Table-Driven Tests (Spécifique Go)

**Présent dans** : Tous les katas

**Concepts clés** :
- Définir plusieurs cas de test dans un slice
- Boucler dessus avec `t.Run()`
- DRY : le code de test est factorisé

**Structure type** :
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected int
    }{
        {"empty string", "", 0},
        {"single number", "1", 1},
        {"two numbers", "1,2", 3},
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

---

## 🎯 Partie 4 : Synthèse et Préparation (30min)

### Tableau Récapitulatif des Katas

| Kata | Difficulté | Patterns Clés | Temps Estimé | Priorité Shodo |
|------|-----------|---------------|--------------|----------------|
| String Calculator | ⭐⭐ | Parsing, edge cases | 30-45min | Moyenne |
| Tennis | ⭐⭐ | State machine | 45-60min | Moyenne |
| Mars Rover | ⭐⭐⭐ | Grille 2D, Command, State | 60-90min | **Haute** |
| Game of Life | ⭐⭐⭐ | Grille 2D, Voisins | 60-90min | Haute |
| Sailboat | ⭐⭐ | Grille 2D, Calculs | 60min | Moyenne |
| Gilded Rose | ⭐⭐⭐ | Refactoring | 60-90min | Moyenne |

### Stratégie d'Apprentissage pour Demain

**Ordre recommandé pour les jours suivants** :
1. **String Calculator** (J2) → Pratique TDD pure
2. **Tennis** (J2) → Introduction aux states
3. **Mars Rover** (J3-4) → Kata principal, à répéter 2-3 fois
4. **Game of Life** (J3-4) → Grille 2D alternative
5. **Sailboat** (J4-5) → Si temps disponible
6. **Gilded Rose** (J5) → Refactoring si temps

---

## ✅ Checklist Jour 1

Avant de passer au Jour 2, tu dois pouvoir répondre à ces questions :

- [ ] **Qu'est-ce que TDD en une phrase ?**
  - _Réponse attendue : Écrire le test avant le code, faire passer le test, puis refactorer._

- [ ] **Quels sont les 3 piliers du Software Craftsmanship ?**
  - _Réponse : TDD, Clean Code, Refactoring continu_

- [ ] **Pourquoi fait-on des katas ?**
  - _Réponse : Pour s'entraîner, pas pour résoudre un nouveau problème. Répétition pour améliorer la technique._

- [ ] **Cite 3 katas et leur essence en une phrase**
  - _Exemple : Mars Rover = robot sur grille avec commandes ; Tennis = score avec state machine ; String Calculator = parsing et addition_

- [ ] **Quel est LE kata à maîtriser en priorité pour Shodo ?**
  - _Réponse : Mars Rover_

- [ ] **Cite 3 patterns techniques qui reviennent**
  - _Réponse : Grilles 2D, State machines, Command pattern_

- [ ] **Qu'est-ce qui est le PLUS évalué en entretien kata ?**
  - _Réponse : La démarche TDD, les baby steps, la communication, le clean code_

---

## 📝 Notes Personnelles

_Espace pour tes réflexions, questions, insights du jour :_

```
Qu'est-ce qui m'a surpris aujourd'hui ?


Quels sont les 2-3 points que je veux vraiment retenir ?


Qu'est-ce qui me fait peur / me semble difficile ?


Quel kata m'attire le plus ?

```

---

## 🚀 Prochaine Étape

Bravo pour avoir posé les fondations ! 🎉

**Demain** : On passe au concret avec les patterns techniques et les premiers katas guidés.

→ [Jour 2 : MEZZO - Patterns & Méthodologie](./day2.md)

---

**Tips pour ce soir** :
- Repose-toi, tu vas coder intensément demain
- Lis rapidement la doc du package `testing` de Go si tu n'es pas à l'aise
- Regarde 1-2 vidéos courtes de TDD en Go si tu veux (optionnel)
