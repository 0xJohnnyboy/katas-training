# Jour 5 - Consolidation & Mock Interview

> **Durée** : 3-4 heures  
> **Objectif** : Valider les acquis, simuler les conditions réelles, se préparer mentalement

---

## 🎯 Objectifs de la Journée

À la fin de ce jour, tu dois :
- ✅ Avoir fait un mock interview chronométré
- ✅ Identifier tes points forts et axes d'amélioration
- ✅ Être confiant pour l'entretien réel
- ✅ Avoir une checklist de ce qu'il faut faire le jour J

---

## 📋 Planning de la Journée

- **9h-10h** : Révision rapide des patterns clés
- **10h-11h30** : Mock Interview - Mars Rover
- **11h30-12h** : Débriefing & analyse
- **12h-13h30** : Pause déjeuner
- **13h30-14h30** : Kata challenge au choix
- **14h30-15h30** : Préparation mentale & checklist finale
- **15h30-16h** : Repos et visualisation

---

## 📖 Partie 1 : Révision Express (60min)

### Patterns Clés à Connaître

#### 1. Grille 2D avec Position

```go
type Position struct {
    X, Y int
}

func (p Position) Move(dx, dy int) Position {
    return Position{X: p.X + dx, Y: p.Y + dy}
}

// Wrapping
func (p Position) Wrap(width, height int) Position {
    x := ((p.X % width) + width) % width
    y := ((p.Y % height) + height) % height
    return Position{X: x, Y: y}
}

// Voisins (8 directions)
func (p Position) Neighbors() []Position {
    neighbors := []Position{}
    for dx := -1; dx <= 1; dx++ {
        for dy := -1; dy <= 1; dy++ {
            if dx == 0 && dy == 0 {
                continue
            }
            neighbors = append(neighbors, Position{X: p.X + dx, Y: p.Y + dy})
        }
    }
    return neighbors
}
```

#### 2. Directions avec Rotations

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
    delta := deltas[d]
    return delta[0], delta[1]
}
```

#### 3. Table-Driven Tests

```go
func TestSomething(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected int
    }{
        {"case 1", "input1", 1},
        {"case 2", "input2", 2},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Function(tt.input)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}
```

### Checklist Mentale TDD

Avant de coder, répète-toi :
1. ✅ Je commence par le test le **plus simple**
2. ✅ Je fais passer le test avec le **code minimum**
3. ✅ Je refactore **immédiatement** si je vois de la duplication
4. ✅ J'avance par **baby steps**
5. ✅ Je **verbalise** ma démarche

---

## 🎤 Partie 2 : Mock Interview (90min)

### Format de l'Interview

**Durée totale** : 90 minutes
- 5min : Introduction et questions
- 60-75min : Coding
- 10min : Débriefing avec l'intervieweur

### Simulation : Mars Rover

**Tu vas jouer les deux rôles** (ou demande à quelqu'un de t'observer).

---

### Rôle : Intervieweur

**Script d'introduction** :

> "Bonjour ! Aujourd'hui, nous allons faire un kata ensemble. L'objectif n'est pas de finir à 100%, mais de voir ta démarche, comment tu codes, comment tu testes.
>
> Le kata du jour est **Mars Rover**. Tu connais peut-être déjà, mais je vais te présenter le contexte :
>
> Nous envoyons un robot sur Mars pour l'explorer. Ce robot, appelé Rover, se déplace sur une grille rectangulaire. Il a une position (x, y) et une orientation (Nord, Sud, Est, Ouest).
>
> Le rover reçoit une chaîne de commandes :
> - 'F' : avancer d'une case
> - 'B' : reculer d'une case
> - 'L' : tourner à gauche (90°)
> - 'R' : tourner à droite (90°)
>
> Exemple : si le rover est en (0, 0) face au Nord et reçoit 'F', il se retrouve en (0, 1).
>
> **Première étape** : commence par créer un rover avec une position et une orientation initiales, et implémente les commandes de rotation.
>
> Des questions avant de commencer ?"

**Grille d'évaluation** (noter mentalement ou sur papier) :

| Critère | Score (1-5) | Observations |
|---------|-------------|--------------|
| **TDD** : Commence par les tests | | |
| **Baby steps** : Avance progressivement | | |
| **Communication** : Explique sa démarche | | |
| **Nommage** : Variables et fonctions explicites | | |
| **Refactoring** : Améliore le code régulièrement | | |
| **Structure** : Code bien organisé | | |
| **Gestion du temps** : Reste calme et avance | | |

---

### Rôle : Candidat

**Setup** :
```bash
# Crée un nouveau dossier pour le mock
mkdir mock_interview
cd mock_interview
go mod init mock
mkdir marsrover
cd marsrover
```

**Lance le chrono !** ⏱️

**Étapes à suivre** :

#### Phase 1 : Questions et Clarifications (2-3min)

Pose des questions avant de coder :
- "La grille a-t-elle des dimensions fixes ?"
- "Que se passe-t-il si le rover sort de la grille ?"
- "Y a-t-il des obstacles ?"
- "Quel format de sortie veux-tu ?"

**Même si tu connais déjà les réponses, pose les questions !**

#### Phase 2 : Plan d'Approche (2min)

Énonce ton plan à voix haute :
> "Je vais commencer par créer la structure du Rover avec sa position et son orientation. Ensuite, j'implémenterai les rotations, puis les mouvements. Je ferai tout en TDD, en commençant par le test le plus simple."

#### Phase 3 : Coding (55-60min)

**IMPORTANT** : Verbalise en continu !

Exemples de ce qu'il faut dire :
- "Je vais écrire un premier test pour créer un rover..."
- "Ce test échoue, maintenant je vais écrire le code minimum..."
- "Bon, là je vois de la duplication, je vais refactorer..."
- "Je pense qu'il serait plus clair de créer une fonction Move()..."

**Checklist des fonctionnalités** :
```markdown
- [ ] Création du rover avec position initiale (5min)
- [ ] Rotation left (5min)
- [ ] Rotation right (5min)
- [ ] Mouvement forward (10min)
- [ ] Mouvement backward (5min)
- [ ] Séquence de commandes (5min)
- [ ] Wrapping / grille toroïdale (15min)
- [ ] Obstacles (10min) - optionnel
```

**Règles du jeu** :
- ✅ Verbalise à chaque étape
- ✅ Écris le test AVANT le code
- ✅ Commit mental à chaque étape qui passe
- ❌ Ne te précipite pas
- ❌ N'efface pas de code qui fonctionne sans raison

#### Phase 4 : Débriefing (5-10min)

Questions à te poser :
- Qu'est-ce qui s'est bien passé ?
- Qu'est-ce qui était difficile ?
- Qu'est-ce que j'aurais pu faire différemment ?
- Ai-je respecté le TDD ?
- Mon code est-il clair ?

---

## 📊 Auto-Évaluation Détaillée

### Grille d'Auto-Évaluation

**Remplis honnêtement après le mock** :

#### 1. TDD (1-5) : _____

**Critères** :
- J'ai écrit le test avant le code à chaque fois
- J'ai fait échouer le test (RED) avant de le faire passer
- J'ai refactoré après avoir fait passer le test
- Je n'ai pas écrit de code "au cas où"

**Commentaire** :
```
Ce qui a bien marché :


Ce qui a été difficile :

```

---

#### 2. Baby Steps (1-5) : _____

**Critères** :
- Mes tests avançaient par petits incréments
- Je n'ai pas essayé de tout implémenter d'un coup
- Chaque test ajoutait UNE seule fonctionnalité
- J'ai résisté à la tentation de "sauter" des étapes

**Commentaire** :
```
Exemples de bons baby steps :


Moments où j'ai fait des sauts trop grands :

```

---

#### 3. Communication (1-5) : _____

**Critères** :
- J'ai verbalisé ma démarche en continu
- J'ai expliqué POURQUOI je faisais mes choix
- J'ai partagé mes hésitations et réflexions
- J'ai posé des questions quand nécessaire

**Commentaire** :
```
Qu'est-ce que j'ai bien communiqué ?


Qu'est-ce que j'aurais dû dire plus ?

```

---

#### 4. Qualité du Code (1-5) : _____

**Critères** :
- Noms de variables/fonctions explicites
- Fonctions courtes et ciblées
- Pas de duplication
- Code facile à lire

**Commentaire** :
```
Meilleurs nommages :


Code à améliorer :

```

---

#### 5. Gestion du Temps (1-5) : _____

**Critères** :
- Je n'ai pas paniqué
- J'ai avancé régulièrement
- Je n'ai pas bloqué trop longtemps sur un problème
- J'ai priorisé les fonctionnalités

**Commentaire** :
```
Temps passé sur chaque partie :
- Setup : ___min
- Tests basiques : ___min
- Mouvements : ___min
- Wrapping : ___min
- Total : ___min

Où ai-je perdu du temps ?

```

---

**SCORE TOTAL** : _____/25

- **20-25** : Excellent ! Tu es prêt 🎉
- **15-19** : Très bien, quelques ajustements
- **10-14** : Bon, mais travaille les points faibles
- **<10** : Refais un mock demain

---

## 🔥 Partie 3 : Kata Challenge (60min)

**Objectif** : Tester ta vitesse et ta fluidité.

### Challenge : Mars Rover Speed Run

**Règles** :
- Timer : 45 minutes MAX
- Fonctionnalités obligatoires :
  - Création rover
  - Rotations (L, R)
  - Mouvements (F, B)
  - Séquence de commandes
  - Wrapping

**Contraintes additionnelles** (choisis-en UNE) :

#### Option A : Code Golf
Fais le kata avec le moins de lignes possible (tout en restant lisible).

#### Option B : Zero Bugs
Une seule règle : AUCUN test ne doit jamais échouer après avoir été écrit.
Si un test échoue, tu as le droit de supprimer le test (mais pas le code).

#### Option C : Live Documentation
Chaque fonction doit avoir un commentaire GoDoc qui explique son comportement.

---

## 🧠 Partie 4 : Préparation Mentale (60min)

### Checklist du Jour J

**La veille de l'entretien** :
- [ ] Relis ce guide rapidement (30min max)
- [ ] Refais un kata rapide (45min) pour te remettre en jambes
- [ ] Prépare ton environnement : éditeur, terminal, tests
- [ ] Dors bien (vraiment !)

**Le matin de l'entretien** :
- [ ] Petit-déjeuner léger
- [ ] Relis les patterns clés (15min)
- [ ] Fais quelques exercices de respiration
- [ ] Arrive 5-10min en avance (physiquement ou en ligne)

**Juste avant de commencer** :
- [ ] Respire profondément 3 fois
- [ ] Rappelle-toi : ils veulent voir ta démarche, pas la perfection
- [ ] Souris, sois toi-même

---

### Phrases à se Répéter

**Quand tu doutes** :
- "Baby steps. Un test à la fois."
- "TDD : RED → GREEN → REFACTOR."
- "Commence par le plus simple."

**Quand tu bloques** :
- "Qu'est-ce que je peux tester de plus simple ?"
- "Est-ce que je peux décomposer ce problème ?"
- "Je vais verbaliser ma réflexion."

**Quand ça va trop vite** :
- "Je prends 30 secondes pour réfléchir."
- "Laisse-moi écrire ce test d'abord."
- "Je vais refactorer avant de continuer."

---

### Gestion des Situations Difficiles

#### Si tu bloques sur un test
1. Verbalise : "Je réfléchis à comment tester cette partie..."
2. Propose : "Je pourrais commencer par tester un cas plus simple..."
3. Demande : "Est-ce que tu as une préférence sur comment aborder ça ?"

#### Si tu ne finis pas
**C'EST NORMAL !** Beaucoup de candidats ne finissent pas.

Dis simplement :
> "Ok, je n'ai pas eu le temps de finir les obstacles, mais si j'avais continué, j'aurais testé d'abord le cas où le rover rencontre un obstacle en (x, y), puis j'aurais modifié la méthode Move() pour vérifier avant de déplacer..."

**Montre que tu sais où tu allais.**

#### Si tu fais une erreur
1. **Ne panique pas** : "Ah, je vois l'erreur..."
2. **Corrige calmement** : Refais le test ou le code
3. **Explique** : "J'avais oublié de gérer le cas négatif..."

**Les erreurs sont humaines et montrent ton processus de réflexion !**

---

### Visualisation Positive (10min)

Ferme les yeux et visualise-toi pendant l'entretien :

1. Tu poses des questions pertinentes
2. Tu écris ton premier test avec confiance
3. Il passe au vert, tu souris
4. Tu refactores naturellement
5. Tu communiques clairement ta démarche
6. L'intervieweur hoche la tête, satisfait
7. Tu termines sereinement

**Ressens les émotions positives** de cette situation.

---

## 📋 Checklist Finale - Es-tu Prêt ?

### Compétences Techniques

- [ ] Je peux créer un projet Go et lancer des tests
- [ ] Je connais les patterns Position et Direction par cœur
- [ ] Je sais faire du wrapping de grille
- [ ] Je peux écrire des tests table-driven
- [ ] Je maîtrise le cycle TDD

### Katas

- [ ] J'ai fait Mars Rover au moins 3 fois
- [ ] Je peux démarrer Mars Rover sans réfléchir
- [ ] J'ai fait au moins un autre kata (Game of Life, Tennis, String Calculator)
- [ ] Mon temps sur Mars Rover est <60min

### Soft Skills

- [ ] Je sais poser des questions de clarification
- [ ] Je verbalise naturellement ma démarche
- [ ] Je reste calme sous pression
- [ ] Je sais reconnaître et corriger mes erreurs

### Mindset

- [ ] Je comprends qu'on évalue la démarche, pas le résultat final
- [ ] Je sais qu'il est normal de ne pas finir
- [ ] Je suis prêt à partager mes réflexions
- [ ] Je me sens confiant (ou au moins... moins stressé)

---

## 🎯 Derniers Conseils

### Ce que les Intervieweurs de Shodo Cherchent

**✅ Ils ADORENT voir** :
- Du TDD naturel et fluide
- De la communication claire
- Des refactorings courageux
- De la curiosité et des questions
- Un code propre et lisible

**❌ Ils n'aiment PAS** :
- Coder en silence
- Ignorer les tests
- Du code brouillon qu'on ne nettoie jamais
- Se précipiter sans réfléchir
- Prétendre qu'on ne fait jamais d'erreurs

### La Vraie Clé du Succès

**Ce n'est pas la vitesse.**
**Ce n'est pas la perfection.**

**C'est de montrer comment tu penses, comment tu travailles, et comment tu améliores ton code en continu.**

Shodo cherche des craftsmen/craftswomen, pas des machines à coder.

**Sois toi-même, partage ton processus, et tout ira bien.** 🚀

---

## 📝 Notes Finales

```
Qu'est-ce que j'ai appris cette semaine ?




Quelle est ma plus grande fierté ?




Qu'est-ce que je veux encore améliorer ?




Mon kata préféré et pourquoi :




Message à moi-même pour le jour de l'entretien :




```

---

## 🎉 Conclusion

**Bravo !** Tu as terminé ce programme intensif. Tu es maintenant prêt pour l'entretien Shodo.

### Rappels Importants

1. **TDD, TDD, TDD** : C'est la base de tout
2. **Communique** : Verbalise sans cesse
3. **Baby steps** : Petit à petit
4. **Sois toi-même** : L'authenticité compte
5. **Respire** : Tu as préparé, tu es prêt

### Le Jour de l'Entretien

**Rappelle-toi** :
- Tu connais Mars Rover par cœur
- Tu maîtrises le TDD
- Tu as les patterns en tête
- Tu es capable de bien communiquer

**Et surtout** :
- Amuse-toi ! Les katas sont là pour ça
- Montre ton process, pas la perfection
- Pose des questions
- **Crois en toi** 💪

---

## 🔗 Récapitulatif du Programme

- [README principal](./README.md)
- [Jour 1 - MACRO](./day1.md) : Philosophie et vue d'ensemble
- [Jour 2 - MEZZO](./day2.md) : Patterns et méthodologie
- [Jours 3-4 - MICRO](./day3_4.md) : Pratique intensive
- [Jour 5 - CONSOLIDATION](./day5.md) : Mock interview et préparation   

---

## 🍀 Bonne Chance !

Tu vas assurer. On croit en toi ! 🚀

N'oublie pas de revenir partager ton expérience après l'entretien ! 😊

---

**Last but not least** : Si tu réussis l'entretien (ce qui va arriver !), n'oublie pas de contribuer à la communauté en partageant ton expérience et en aidant les prochains candidats. C'est ça, l'esprit Software Craftsmanship ! ✨
