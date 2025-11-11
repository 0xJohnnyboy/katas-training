# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

Collection de katas de programmation classiques du Software Craftsmanship. Chaque kata contient ses règles complètes et cas de test de manière agnostique du langage.

## Structure

```
katas-training/
├── katas/                    # Règles des katas (markdown)
│   ├── game_of_life.md
│   ├── karate_chop.md
│   ├── klondike.md
│   ├── mars_rover.md
│   ├── pacman.md
│   ├── reversi.md
│   ├── sailboat.md
│   ├── string_calculator.md
│   ├── supermarket_checkout.md
│   └── tennis.md
├── mars_rover/               # Implémentations en Go
│   ├── v1/                   # Première tentative
│   └── v2/                   # Seconde tentative
├── game_of_life/
│   ├── v1/
│   └── v2/
├── string_calculator/
├── tennis/
└── cmd/                      # Applications CLI (si présentes)
```

## Katas Disponibles

**Navigation et Grilles 2D:** Mars Rover, Game of Life, Sailboat, Pac-Man

**Jeux de Plateau:** Reversi (Othello), Klondike (Solitaire)

**Logique et Calcul:** String Calculator, Tennis, Supermarket Checkout

**Algorithmes:** Karate Chop (Binary Search)

## Running Tests

Les implémentations Go utilisent le package standard `testing` :

```bash
# Tests pour un kata spécifique
cd <kata_name>/v1
go test -v

# Tests pour une version spécifique
cd mars_rover/v2
go test -v

# Tous les tests
go test ./...

# Test spécifique
go test -v -run TestName
```

## Format des Règles

Chaque fichier markdown dans `katas/` suit cette structure :
- **Règles** : Description complète des règles du kata
- **Cas de test** : Tableaux avec inputs/outputs attendus
- **Exemples** : Cas concrets, patterns, séquences de jeu

Les règles sont volontairement agnostiques du langage pour permettre l'implémentation dans n'importe quel langage.

## Implémentations Go

- **Module Name**: `katas` (défini dans `go.mod`)
- **Go Version**: 1.24.2
- **Dépendances**: Aucune (standard library uniquement)
- **Pattern**: Versions multiples (v1, v2) pour différentes approches ou tentatives
- **Tests**: Subtests avec `t.Run()`, table-driven tests

## Notes de Développement

- Les fichiers markdown de règles ne doivent PAS être modifiés sauf pour corrections
- Les implémentations sont organisées par version pour permettre différentes approches
- Pas de dépendances externes dans les implémentations Go
- TDD (Test-Driven Development) est l'approche recommandée
