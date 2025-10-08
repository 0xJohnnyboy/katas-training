# Guide d'Entraînement Intensif - Katas Shodo

> Programme intensif de préparation aux entretiens techniques Shodo Lille
> Focus : Go, TDD, Software Craftsmanship

## 🎯 Objectif

Se préparer efficacement aux katas techniques Shodo en 3-5 jours, avec une approche progressive macro-mezzo-micro.

## 📋 Structure du Programme

- **Durée totale** : 3-5 jours (adaptable selon disponibilité)
- **Approche** : Macro → Mezzo → Micro
- **Langage** : Go
- **Focus** : TDD, Clean Code, Software Craftsmanship

## 🗓️ Plan de Formation

### [Jour 1 : MACRO - Comprendre le Contexte](./day1.md)
**Objectif** : Comprendre la philosophie, identifier les patterns communs
- Philosophie Software Craftsmanship chez Shodo
- Tour d'horizon des katas principaux
- Patterns récurrents
- Ce qui est évalué en entretien
- **Durée** : 3-4h

### [Jour 2 : MEZZO - Patterns & Méthodologie](./day2.md)
**Objectif** : Maîtriser les patterns techniques et la méthodologie TDD
- Patterns techniques (grilles 2D, state machines, commands)
- TDD en Go (testing, testify, table-driven tests)
- Comment démarrer un kata
- Premiers katas guidés (String Calculator, Tennis)
- **Durée** : 4-6h

### [Jours 3-4 : MICRO - Pratique Intensive](./day3_4.md)
**Objectif** : Pratiquer, répéter, automatiser
- Mars Rover (prioritaire)
- Game of Life
- Sailboat
- Gilded Rose
- Répétition avec contraintes
- **Durée** : 6-8h par jour

### [Jour 5 : Consolidation & Mock Interview](./day5.md)
**Objectif** : Valider l'acquis, simuler les conditions réelles
- Refactoring sous contraintes
- Mock interview chronométré
- Débriefing et axes d'amélioration
- **Durée** : 3-4h

## 📚 Ressources

### Repositories Principaux
- [Shodo Code Retreat](https://shodo-nantes.github.io/code-retreat/)
- [Kata Starters](https://github.com/shodo-nantes/kata-starters)

### Katas Prioritaires
1. **Mars Rover** ⭐⭐⭐ (très probable)
2. **Game of Life** ⭐⭐
3. **String Calculator** ⭐⭐ (bon pour débuter)
4. **Tennis** ⭐
5. **Sailboat** ⭐
6. **Gilded Rose** ⭐ (refactoring)

### Documentation Go
- [Go by Example](https://gobyexample.com/)
- [Testing package](https://pkg.go.dev/testing)
- [Testify](https://github.com/stretchr/testify)

## 🎓 Principes Clés à Retenir

### Ce qui est Évalué
- ✅ **Démarche TDD** : Red → Green → Refactor
- ✅ **Baby steps** : avancer par petits incréments
- ✅ **Clean Code** : nommage, lisibilité, simplicité
- ✅ **Communication** : expliquer sa démarche
- ✅ **Refactoring** : améliorer continuellement

### Ce qui n'est PAS Évalué (ou moins)
- ❌ Performance optimale
- ❌ Solution complète et exhaustive
- ❌ Connaissance de tous les patterns avancés

## 💡 Conseils Généraux

### Pendant l'Entraînement
1. **Chronométrez-vous** : visez 45-60min par kata
2. **Recommencez** : refaire un kata 2-3 fois est normal
3. **Verbalisez** : parlez à voix haute pendant que vous codez
4. **Commencez simple** : le test le plus simple possible
5. **Refactorez tôt** : dès que vous voyez de la duplication

### Le Jour J
1. **Posez des questions** : clarifiez les specs
2. **Proposez un plan** : "je vais commencer par..."
3. **Tests d'abord** : toujours
4. **Itérez visiblement** : montrez votre processus de pensée
5. **Restez calme** : mieux vaut un code simple qui marche qu'un code complexe incomplet

## 📊 Auto-Évaluation

Après chaque kata, évaluez-vous sur ces critères (1-5) :

| Critère | Score | Notes |
|---------|-------|-------|
| J'ai commencé par un test | /5 | |
| J'ai fait des baby steps | /5 | |
| Mon code est lisible | /5 | |
| J'ai refactoré régulièrement | /5 | |
| J'ai terminé dans les temps | /5 | |

**Score cible** : 20/25 minimum

## 🚀 Comment Utiliser ce Guide

1. **Suivez l'ordre** : Jour 1 → 2 → 3-4 → 5
2. **Adaptez le rythme** : selon votre disponibilité
3. **Documentez** : prenez des notes dans chaque fichier
4. **Pratiquez** : plus que lire, FAIRE
5. **Répétez** : la répétition crée l'automatisme

## 📝 Tracking de Progression

```markdown
### Jour 1 - MACRO
- [ ] Lecture philosophie Software Craftsmanship (30min)
- [ ] Tour d'horizon des katas (1h)
- [ ] Identification des patterns (1h)
- [ ] Notes sur critères d'évaluation (30min)

### Jour 2 - MEZZO
- [ ] Étude patterns techniques (1h)
- [ ] Setup projet Go + tests (30min)
- [ ] String Calculator guidé (1h)
- [ ] Tennis guidé (1h30)

### Jours 3-4 - MICRO
- [ ] Mars Rover - tentative 1
- [ ] Mars Rover - tentative 2
- [ ] Mars Rover - tentative 3
- [ ] Game of Life - tentative 1
- [ ] Game of Life - tentative 2
- [ ] Sailboat - tentative 1

### Jour 5 - CONSOLIDATION
- [ ] Mars Rover sous contrainte "no if"
- [ ] Mock interview chronométré
- [ ] Débriefing personnel
```

---

**Prêt à commencer ?** → [Jour 1 : Comprendre le Contexte](./day1.md)

**Questions ?** N'hésite pas à adapter ce programme selon tes besoins et contraintes !
