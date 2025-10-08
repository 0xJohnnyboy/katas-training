# Cheat Sheet - Référence Rapide Katas

> À relire 15 minutes avant l'entretien

---

## ⚡ TDD en 3 Étapes

```
🔴 RED    → Écris un test qui échoue
🟢 GREEN  → Code minimum pour le faire passer  
♻️  REFACTOR → Améliore sans changer le comportement
```

---

## 🎯 Mars Rover - Ordre des Tests

1. Créer rover avec position initiale
2. Tourner à gauche (L)
3. Tourner à droite (R)
4. Avancer (F) - une direction
5. Avancer (F) - toutes directions
6. Reculer (B)
7. Séquence de commandes ("FFRFF")
8. Wrapping (sortir = réapparaître)
9. Obstacles (stop si bloqué)

**Temps cible** : 45-60 minutes

---

## 🧩 Patterns Go Essentiels

### Position Immutable

```go
type Position struct {
    X, Y int
}

func (p Position) Move(dx, dy int) Position {
    return Position{X: p.X + dx, Y: p.Y + dy}
}
```

### Direction avec Rotation

```go
type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) TurnLeft() Direction {
    return Direction((int(d) + 3) % 4)
}

func (d Direction) TurnRight() Direction {
    return Direction((int(d) + 1) % 4)
}

func (d Direction) Delta() (int, int) {
    deltas := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    return deltas[d][0], deltas[d][1]
}
```

### Wrapping

```go
func wrap(value, max int) int {
    return ((value % max) + max) % max
}

// Usage
x = wrap(x, gridWidth)
y = wrap(y, gridHeight)
```

### Table-Driven Test

```go
func TestSomething(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {"case 1", 1, 2},
        {"case 2", 3, 4},
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

---

## 💬 Phrases Clés à Dire

### Au Début
- "Laisse-moi poser quelques questions pour clarifier..."
- "Je vais commencer par le test le plus simple..."
- "Mon approche sera de..."

### Pendant le Coding
- "Je vais écrire un test pour..."
- "Maintenant je fais le code minimum..."
- "Je vois de la duplication, je refactorise..."
- "Ce serait plus clair si j'extrais une fonction..."

### Si Tu Bloques
- "Laisse-moi réfléchir 30 secondes..."
- "Je pense qu'il faut que je décompose le problème..."
- "Qu'est-ce que tu en penses ?"

### À la Fin
- "Si j'avais plus de temps, j'ajouterais..."
- "Je refactorerais X pour améliorer Y..."

---

## ✅ Checklist Avant de Coder

1. [ ] J'ai posé des questions de clarification
2. [ ] J'ai annoncé mon plan d'approche
3. [ ] Je sais par quel test commencer
4. [ ] Mon éditeur et terminal sont prêts
5. [ ] Je respire calmement

---

## 🚫 Erreurs à Éviter

❌ Coder en silence  
❌ Écrire du code avant le test  
❌ Faire de gros sauts (pas de baby steps)  
❌ Ne jamais refactorer  
❌ Paniquer si on ne finit pas  
❌ Prétendre qu'on ne fait jamais d'erreurs  

---

## 🎯 Ce Qu'ils Évaluent

### Importance Haute ⭐⭐⭐
- TDD (test d'abord)
- Communication
- Baby steps
- Refactoring

### Importance Moyenne ⭐⭐
- Structure du code
- Nommage
- Gestion du temps

### Importance Basse ⭐
- Finir à 100%
- Performance optimale
- Connaissance de patterns avancés

---

## 🧘 Gestion du Stress

**Si tu paniques** :
1. Respire profondément (3x)
2. Verbalise : "Je prends 30 secondes..."
3. Reviens au test le plus simple
4. Continue petit à petit

**Rappelle-toi** :
- Ils veulent voir ta DÉMARCHE
- Les erreurs sont normales
- Ne pas finir est OK
- Tu t'es préparé, tu es prêt

---

## ⏱️ Gestion du Temps (60min)

- **0-5min** : Questions & plan
- **5-15min** : Rotations (L/R)
- **15-30min** : Mouvements (F/B)
- **30-40min** : Séquences
- **40-55min** : Wrapping
- **55-60min** : Obstacles OU recap

**Si en retard** : Passe au suivant, tu pourras y revenir.

---

## 🎓 Les 5 Règles d'Or

1. **TEST D'ABORD** - Toujours
2. **BABY STEPS** - Un incrément à la fois
3. **VERBALISE** - Parle en continu
4. **REFACTORE** - Dès que tu vois de la duplication
5. **RESTE CALME** - Tu es préparé

---

## 💪 Auto-Motivation

Tu as fait :
- ✅ 5 jours de préparation intensive
- ✅ Mars Rover plusieurs fois
- ✅ Plusieurs autres katas
- ✅ Un mock interview

**Tu es prêt. Tu vas assurer. On croit en toi !** 🚀

---

## 📞 Derniers Mots

**N'oublie pas** :
- Sois toi-même
- Amuse-toi (vraiment !)
- Montre ton process
- Pose des questions
- Respire

**Bonne chance !** 🍀

---

_Imprime cette page et garde-la à portée de main le jour J._
