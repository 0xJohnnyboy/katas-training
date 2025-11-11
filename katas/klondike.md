# Klondike (Solitaire)

## Règles

Jeu de solitaire avec un jeu de 52 cartes.

**Disposition initiale:**
- 7 colonnes (tableau) : colonne 1 a 1 carte, colonne 2 a 2 cartes, etc.
- Seule la carte du dessus de chaque colonne est visible
- 4 piles de fondation (une par couleur : ♠ ♥ ♦ ♣)
- 1 pile de pioche (stock) avec les cartes restantes
- 1 pile de défausse (waste)

**Mouvements dans les colonnes:**
- Empiler en ordre décroissant (K, Q, J, 10... A)
- Alterner les couleurs : rouge (♥ ♦) et noir (♠ ♣)
- Déplacer un groupe de cartes visibles ensemble
- Colonne vide : accepte uniquement un Roi
- Retirer une carte révèle la carte cachée en dessous

**Mouvements vers les fondations:**
- Commencer par un As
- Empiler par couleur en ordre croissant (A, 2, 3... K)
- Une carte à la fois

**Pioche:**
- Retourner 3 cartes (ou 1 selon variante)
- Seule la carte du dessus de la défausse est jouable

**Victoire:** Les 4 fondations complètes (A à K)

## Cas de test

### Mouvements valides dans les colonnes

| Carte du dessus | Carte à placer | Valide | Raison                    |
|-----------------|----------------|--------|---------------------------|
| 7♠              | 6♥             | Oui    | Décroissant, rouge/noir   |
| 7♠              | 6♦             | Oui    | Décroissant, rouge/noir   |
| 7♥              | 6♠             | Oui    | Décroissant, rouge/noir   |
| 7♥              | 6♣             | Oui    | Décroissant, rouge/noir   |
| 7♠              | 6♠             | Non    | Même couleur              |
| 7♠              | 6♣             | Non    | Même couleur              |
| 7♠              | 5♥             | Non    | Pas décroissant de 1      |
| 7♠              | 8♥             | Non    | Pas décroissant           |
| A♠              | K♥             | Non    | As n'accepte rien         |

### Mouvements vers colonne vide

| Carte | Valide | Raison                |
|-------|--------|-----------------------|
| K♠    | Oui    | Roi accepté           |
| K♥    | Oui    | Roi accepté           |
| Q♠    | Non    | Seul Roi accepté      |
| A♠    | Non    | Seul Roi accepté      |

### Mouvements vers les fondations

| Fondation actuelle | Carte | Valide | Raison                    |
|--------------------|-------|--------|---------------------------|
| Vide               | A♠    | Oui    | Commence par As           |
| Vide               | 2♠    | Non    | Doit commencer par As     |
| A♠                 | 2♠    | Oui    | Même couleur, croissant   |
| A♠                 | 2♥    | Non    | Couleur différente        |
| A♠                 | 3♠    | Non    | Pas croissant de 1        |
| 5♥                 | 6♥    | Oui    | Même couleur, croissant   |
| K♠                 | -     | -      | Fondation complète        |

### Groupes de cartes visibles

Colonne : [X] [X] K♠ Q♥ J♠ (3 dernières visibles)

| Groupe à déplacer | Sur carte | Valide | Raison                    |
|-------------------|-----------|--------|---------------------------|
| Q♥ J♠             | K♦        | Oui    | Q sur K, couleurs OK      |
| K♠ Q♥ J♠          | Vide      | Oui    | Roi sur colonne vide      |
| J♠                | Q♣        | Oui    | J sur Q, couleurs OK      |
| K♠                | Vide      | Oui    | Roi seul sur vide         |

### Révéler une carte

Colonne : [5♦ caché] 4♠ visible

- Déplacer 4♠ → 5♦ devient visible et jouable

### Exemples de jeu

**Situation initiale:**
```
Col1: K♠
Col2: [X] Q♥
Col3: [X] [X] J♣
Fondations: Vides
Stock: ...
Waste: Vide
```

**Mouvements possibles:**
1. Q♥ sur K♠ (valide, révèle carte col2)
2. J♣ sur Q♥ (valide, révèle 2ème carte col3)
3. K♠ Q♥ J♣ vers colonne vide (si existe)

**Séquence vers fondation:**
```
Waste: A♠
Col1: 2♠

1. A♠ vers fondation ♠ (commence fondation)
2. 2♠ vers fondation ♠ (sur A♠)
```
