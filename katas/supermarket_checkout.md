# Supermarket Checkout

## Règles

Implémenter une caisse de supermarché qui calcule le prix total d'articles scannés.

**Articles:**
- Identifiés par SKU (Stock Keeping Unit) : lettres A, B, C, D, etc.
- Prix unitaire pour chaque article
- Offres spéciales : "N pour X€" (ex: 3 pour 130)

**Contraintes:**
- Les articles peuvent être scannés dans n'importe quel ordre
- La caisse doit automatiquement appliquer les meilleures offres
- Les règles de prix sont configurables (changent fréquemment)
- Scanner plusieurs fois le même article doit être géré correctement

**Exemple de tarifs:**
- A : 50, offre 3 pour 130
- B : 30, offre 2 pour 45
- C : 20
- D : 15

## Cas de test

### Prix unitaires sans offre

| Articles scannés | Calcul      | Total |
|------------------|-------------|-------|
| ""               | 0           | 0     |
| A                | 50          | 50    |
| C                | 20          | 20    |
| D                | 15          | 15    |
| AC               | 50 + 20     | 70    |
| CD               | 20 + 15     | 35    |
| ACDC             | 50+20+15+20 | 105   |

### Offres simples

Tarifs : A=50 (3 pour 130), B=30 (2 pour 45)

| Articles scannés | Calcul          | Total |
|------------------|-----------------|-------|
| AA               | 50 + 50         | 100   |
| AAA              | 130             | 130   |
| AAAA             | 130 + 50        | 180   |
| AAAAA            | 130 + 50 + 50   | 230   |
| AAAAAA           | 130 + 130       | 260   |
| B                | 30              | 30    |
| BB               | 45              | 45    |
| BBB              | 45 + 30         | 75    |
| BBBB             | 45 + 45         | 90    |

### Ordre de scan quelconque

Tarifs : A=50 (3 pour 130), B=30 (2 pour 45)

| Articles scannés | Détail              | Total |
|------------------|---------------------|-------|
| ABABA            | 3A + 2B = 130 + 45  | 175   |
| BABAA            | 3A + 2B = 130 + 45  | 175   |
| AABBA            | 3A + 2B = 130 + 45  | 175   |
| BAAAB            | 3A + 2B = 130 + 45  | 175   |

### Mixte avec et sans offres

Tarifs : A=50 (3 pour 130), B=30 (2 pour 45), C=20, D=15

| Articles scannés | Détail                    | Total |
|------------------|---------------------------|-------|
| AAAB             | 130 + 30                  | 160   |
| AAABB            | 130 + 45                  | 175   |
| AAABBD           | 130 + 45 + 15             | 190   |
| DABABA           | 3A + 2B + D = 130+45+15   | 190   |
| AAABBCD          | 130 + 45 + 20 + 15        | 210   |

### Scan incrémental

Scanner article par article et afficher le total courant :

Tarifs : A=50 (3 pour 130), B=30 (2 pour 45)

| Scan | Total courant | Explication       |
|------|---------------|-------------------|
| -    | 0             | Début             |
| A    | 50            | 1 A               |
| B    | 80            | 1 A + 1 B         |
| A    | 130           | 2 A + 1 B         |
| B    | 160           | 2 A + 2 B (offre) |
| A    | 210           | 3 A (offre) + 2 B |

### Offres multiples

Tarifs : A=50 (3 pour 130), B=30 (2 pour 45, 5 pour 100)

| Articles scannés | Détail              | Total |
|------------------|---------------------|-------|
| BB               | 45                  | 45    |
| BBB              | 45 + 30             | 75    |
| BBBB             | 45 + 45             | 90    |
| BBBBB            | 100                 | 100   |
| BBBBBB           | 100 + 30            | 130   |
| BBBBBBB          | 100 + 45            | 145   |

Note : Appliquer toujours l'offre la plus avantageuse (5 pour 100 > 2×(2 pour 45))
