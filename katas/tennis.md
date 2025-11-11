# Tennis

## Règles

Simuler le score d'un jeu de tennis entre deux joueurs.

**Système de score:**
- 0 point = "Love"
- 1 point = "Fifteen"
- 2 points = "Thirty"
- 3 points = "Forty"

**Cas particuliers:**
- Score identique (sauf à 0) : "X-All" (ex: "Fifteen-All", "Thirty-All")
- Score identique à 0 : "Love-All"
- À partir de 40-40 : "Deuce"
- Un joueur mène après deuce : "Advantage [joueur]"
- Gagner le jeu : "Win for [joueur]"

**Pour gagner:**
- Avoir au moins 4 points ET 2 points d'avance sur l'adversaire

## Cas de test

### Scores classiques

Points J1, Points J2 → Score affiché

| J1 | J2 | Score                |
|----|----|----------------------|
| 0  | 0  | Love-All             |
| 1  | 0  | Fifteen-Love         |
| 0  | 1  | Love-Fifteen         |
| 1  | 1  | Fifteen-All          |
| 2  | 0  | Thirty-Love          |
| 0  | 2  | Love-Thirty          |
| 2  | 1  | Thirty-Fifteen       |
| 1  | 2  | Fifteen-Thirty       |
| 2  | 2  | Thirty-All           |
| 3  | 0  | Forty-Love           |
| 0  | 3  | Love-Forty           |
| 3  | 1  | Forty-Fifteen        |
| 1  | 3  | Fifteen-Forty        |
| 3  | 2  | Forty-Thirty         |
| 2  | 3  | Thirty-Forty         |
| 3  | 3  | Deuce                |
| 4  | 0  | Win for player1      |
| 0  | 4  | Win for player2      |
| 4  | 2  | Win for player1      |
| 2  | 4  | Win for player2      |

### Deuce et advantage

| J1 | J2 | Score                |
|----|----|----------------------|
| 3  | 3  | Deuce                |
| 4  | 3  | Advantage player1    |
| 3  | 4  | Advantage player2    |
| 4  | 4  | Deuce                |
| 5  | 4  | Advantage player1    |
| 5  | 5  | Deuce                |
| 6  | 4  | Win for player1      |
| 4  | 6  | Win for player2      |
| 10 | 8  | Win for player1      |

### Séquences de jeu

Séquence de points gagnés → Scores successifs

**Exemple 1:** J1 gagne 4 points d'affilée
```
0-0: Love-All
1-0: Fifteen-Love
2-0: Thirty-Love
3-0: Forty-Love
4-0: Win for player1
```

**Exemple 2:** Alternance puis J1 gagne
```
0-0: Love-All
1-0: Fifteen-Love
1-1: Fifteen-All
2-1: Thirty-Fifteen
2-2: Thirty-All
3-2: Forty-Thirty
4-2: Win for player1
```

**Exemple 3:** Deuce avec rebondissements
```
3-3: Deuce
4-3: Advantage player1
4-4: Deuce
4-5: Advantage player2
5-5: Deuce
6-5: Advantage player1
7-5: Win for player1
```
