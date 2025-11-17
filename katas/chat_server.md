# Chat Server

## Règles

Implémenter un serveur de chat concurrent avec plusieurs clients et rooms.

**Fonctionnalités :**
- Clients peuvent se connecter/déconnecter
- Envoyer messages à tous (broadcast)
- Créer et rejoindre des rooms
- Messages privés entre utilisateurs
- Liste des utilisateurs connectés

**Architecture :**
- Serveur central qui gère les connexions
- Chaque client = goroutine
- Broadcast via channels
- Synchronisation des états

## Cas de test

### Connexion/Déconnexion

| Action                    | Utilisateurs connectés |
|---------------------------|------------------------|
| Alice se connecte         | [Alice]                |
| Bob se connecte           | [Alice, Bob]           |
| Charlie se connecte       | [Alice, Bob, Charlie]  |
| Bob se déconnecte         | [Alice, Charlie]       |
| Alice se déconnecte       | [Charlie]              |

### Broadcast simple

Utilisateurs : Alice, Bob, Charlie

| Émetteur | Message      | Destinataires      |
|----------|--------------|-------------------|
| Alice    | "Hello all"  | Bob, Charlie      |
| Bob      | "Hi Alice"   | Alice, Charlie    |
| Charlie  | "Bye"        | Alice, Bob        |

Note : L'émetteur ne reçoit pas son propre message

### Rooms (Salons)

Utilisateurs : Alice (room1), Bob (room1), Charlie (room2), Dave (room2)

| Émetteur | Room  | Message  | Destinataires |
|----------|-------|----------|---------------|
| Alice    | room1 | "Hello"  | Bob           |
| Charlie  | room2 | "Hi"     | Dave          |
| Bob      | room1 | "World"  | Alice         |

### Changement de room

| Action                    | Room1         | Room2         |
|---------------------------|---------------|---------------|
| Alice rejoint room1       | [Alice]       | []            |
| Bob rejoint room1         | [Alice, Bob]  | []            |
| Alice rejoint room2       | [Bob]         | [Alice]       |
| Charlie rejoint room2     | [Bob]         | [Alice, Charlie] |

### Messages privés

| Émetteur | Destinataire | Message    | Qui reçoit |
|----------|--------------|------------|------------|
| Alice    | Bob          | "Secret"   | Bob        |
| Bob      | Alice        | "Reply"    | Alice      |
| Charlie  | Bob          | "Private"  | Bob        |

### Commandes

| Commande              | Action                          |
|-----------------------|---------------------------------|
| /join room1           | Rejoindre room1                 |
| /leave                | Quitter la room actuelle        |
| /users                | Lister utilisateurs connectés   |
| /rooms                | Lister rooms disponibles        |
| /msg Bob Hello        | Message privé à Bob             |
| /quit                 | Déconnexion                     |

### Ordre des messages

Alice envoie rapidement 5 messages :

```
Alice: msg1
Alice: msg2
Alice: msg3
Alice: msg4
Alice: msg5
```

Bob et Charlie doivent recevoir les messages dans le même ordre.

### Connexions simultanées

| Test                        | Attendu                     |
|-----------------------------|-----------------------------|
| 10 clients se connectent    | Tous connectés              |
| 100 messages simultanés     | Tous délivrés               |
| 5 clients se déconnectent   | Serveur reste stable        |

### Gestion d'erreurs

| Scénario                      | Comportement                        |
|-------------------------------|-------------------------------------|
| Rejoindre room inexistante    | Créer la room ou erreur             |
| Message à user inexistant     | Erreur "user not found"             |
| Déconnexion brutale client    | Nettoyer ressources, notifier autres|
| Double connexion même nom     | Rejeter ou suffixer (Alice, Alice2) |

## Format des messages

### Message broadcast
```
{
    type: "broadcast"
    from: "Alice"
    message: "Hello everyone"
    timestamp: "2024-01-01T12:00:00Z"
}
```

### Message de room
```
{
    type: "room"
    from: "Alice"
    room: "general"
    message: "Hello room"
    timestamp: "2024-01-01T12:00:00Z"
}
```

### Message privé
```
{
    type: "private"
    from: "Alice"
    to: "Bob"
    message: "Secret message"
    timestamp: "2024-01-01T12:00:00Z"
}
```

### Message système
```
{
    type: "system"
    message: "Alice has joined"
    timestamp: "2024-01-01T12:00:00Z"
}
```

## Scénario complet

```
1. Alice se connecte → broadcast "Alice joined"
2. Bob se connecte → broadcast "Bob joined"
3. Alice: "Hello" → Bob reçoit
4. Bob: "Hi Alice" → Alice reçoit
5. Charlie se connecte → broadcast "Charlie joined"
6. Alice: "/join room1" → Alice dans room1
7. Bob: "/join room1" → Bob dans room1
8. Alice: "Hello room" → seul Bob reçoit
9. Charlie: "Anyone here?" → personne ne reçoit (pas dans room)
10. Bob: "/leave" → Bob quitte room1
11. Alice: "Bob left?" → personne ne reçoit
12. Alice: "/join lobby" → Alice dans lobby
13. Charlie: "Hi" → Alice reçoit (tous deux dans lobby)
```

## Métriques

| Métrique                  | Description                      |
|---------------------------|----------------------------------|
| Clients connectés         | Nombre actuel                    |
| Messages totaux           | Depuis démarrage serveur         |
| Messages/seconde          | Débit actuel                     |
| Rooms actives             | Nombre de rooms avec ≥1 user     |
| Latence moyenne           | Temps émission → réception       |

## Concepts de concurrence

- **Hub central** : goroutine qui gère broadcast
- **Client goroutines** : une par connexion
- **Channels** : communication hub ↔ clients
- **sync.Map** ou **mutex** : liste des clients
- **select** : multiplexage des messages
- **context.Context** : shutdown gracieux
- **Buffered channels** : éviter blocages
