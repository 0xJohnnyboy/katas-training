# Bank Account (Compte Bancaire Concurrent)

## Règles

Implémenter un compte bancaire supportant des opérations concurrentes sécurisées.

**Opérations :**
- `Deposit(amount)` : déposer de l'argent
- `Withdraw(amount)` : retirer de l'argent
- `Balance()` : consulter le solde
- `Transfer(to, amount)` : transférer vers un autre compte

**Contraintes :**
- Plusieurs goroutines peuvent accéder au même compte
- Pas de race conditions
- Pas de solde négatif (sauf si autorisé)
- Transfers atomiques (tout ou rien)
- Éviter les deadlocks

## Cas de test

### Opérations séquentielles

| Opération       | Solde avant | Solde après |
|-----------------|-------------|-------------|
| Deposit(100)    | 0           | 100         |
| Deposit(50)     | 100         | 150         |
| Withdraw(30)    | 150         | 120         |
| Withdraw(20)    | 120         | 100         |
| Balance()       | 100         | 100         |

### Opérations concurrentes - Deposits

10 goroutines déposent 10€ chacune simultanément

| Solde initial | Dépôts simultanés | Solde final attendu |
|---------------|-------------------|---------------------|
| 0             | 10 × 10€          | 100€                |
| 50            | 10 × 10€          | 150€                |
| 100           | 100 × 1€          | 200€                |

### Opérations concurrentes - Withdrawals

| Solde initial | Retraits simultanés | Solde final attendu |
|---------------|---------------------|---------------------|
| 100           | 10 × 5€             | 50€                 |
| 100           | 5 × 10€             | 50€                 |
| 50            | 50 × 1€             | 0€                  |

### Opérations concurrentes - Mixtes

Solde initial : 100€
Simultanément : 5 deposits de 10€ + 5 withdrawals de 10€

| Opération               | Solde final |
|-------------------------|-------------|
| 5×Deposit(10) + 5×Withdraw(10) | 100€   |

### Withdraw avec solde insuffisant

| Solde | Retrait | Résultat           |
|-------|---------|-------------------|
| 50    | 100     | Erreur, solde=50  |
| 0     | 10      | Erreur, solde=0   |
| 30    | 30      | OK, solde=0       |
| 30    | 29      | OK, solde=1       |

### Race condition (doit être évitée)

Solde initial : 100€
2 goroutines tentent de retirer 60€ simultanément

| Résultat attendu                                |
|-------------------------------------------------|
| Une réussit (solde=40), l'autre échoue (solde insuffisant) |

**Résultat INCORRECT** (si race condition) : les deux réussissent, solde=-20

### Transfer simple

| Compte A | Compte B | Transfer A→B | A final | B final |
|----------|----------|--------------|---------|---------|
| 100      | 50       | 30           | 70      | 80      |
| 100      | 0        | 100          | 0       | 100     |
| 50       | 50       | 25           | 25      | 75      |

### Transfer avec échec

| Compte A | Compte B | Transfer A→B | Résultat                    |
|----------|----------|--------------|----------------------------|
| 100      | 50       | 150          | Erreur, A=100, B=50        |
| 0        | 50       | 10           | Erreur, A=0, B=50          |

Le transfer est atomique : soit les deux comptes sont modifiés, soit aucun.

### Transfers concurrents (pas de deadlock)

Comptes A=100, B=100, C=100

Simultanément :
- A → B : 50
- B → C : 50
- C → A : 50

| Compte | Solde final |
|--------|-------------|
| A      | 100         |
| B      | 100         |
| C      | 100         |

Tous les transfers doivent réussir sans deadlock.

### Deadlock classique (à éviter)

Goroutine 1 : Transfer A → B (50)
Goroutine 2 : Transfer B → A (50)

Si mal implémenté :
- G1 lock A, attend B
- G2 lock B, attend A
- **Deadlock**

Solution : ordre cohérent de locking (ex: toujours lock le compte avec l'ID le plus petit en premier)

### Historique des transactions

| Transaction          | Nouveau solde |
|----------------------|---------------|
| Deposit(100)         | 100           |
| Deposit(50)          | 150           |
| Withdraw(30)         | 120           |
| Transfer out(20)     | 100           |
| Transfer in(10)      | 110           |

Chaque opération doit être enregistrée avec timestamp.

## Scénarios de stress

### 1000 opérations concurrentes

Solde initial : 1000€

Opérations simultanées :
- 500 deposits de 1€
- 500 withdrawals de 1€

Solde final attendu : 1000€

### Race sur Balance()

10 goroutines lisent Balance() pendant que 10 autres font Deposit(10)

Balance() doit toujours retourner une valeur cohérente (jamais de valeur intermédiaire corrompue).

### Transfer circulaire

A → B → C → A (chacun transfère 10)

| Compte | Initial | Final |
|--------|---------|-------|
| A      | 100     | 100   |
| B      | 100     | 100   |
| C      | 100     | 100   |

Aucun deadlock, tous les transfers réussissent.

## Concepts de concurrence

- **sync.Mutex** : protéger le solde
- **sync.RWMutex** : lectures multiples, écriture exclusive
- **atomic.Int64** : opérations atomiques simples
- **Lock ordering** : éviter deadlocks dans transfers
- **Transaction log** : historique thread-safe
- **Conditional variables** : attendre solde suffisant

## Variantes

### Avec découvert autorisé
```
Withdraw peut mettre le solde à -100€ max
```

### Avec intérêts
```
Goroutine qui ajoute 1% toutes les secondes
```

### Avec limites de taux
```
Max 3 withdrawals par minute
```
