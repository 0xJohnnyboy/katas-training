# Todo CLI (Refactoring Kata)

Petit outil CLI de gestion de todos avec stockage CSV.

## Ce qui existe

- `add <YYYY-MM-DD> <title> <description>`
- `update <id> <YYYY-MM-DD> <title> <description> <true|false>`
- `delete <id>`
- `list` (affiche seulement les 20 derniers)

Chaque todo contient:

- id
- date (`YYYY-MM-DD`)
- titre
- description
- état complété (`true/false`)

Le starter contient un `todos.csv` avec 200 entrées.

## Objectif du kata

Refactorer du code legacy (grosses fonctions, duplication, mélange parsing/I/O/logique), puis ajouter une feature de recherche avec filtres.

Le client nous a envoyé un message pour l'ajout d'une nouvelle feature:

> Bonjour l'équipe,  
> D'abord merci pour votre travail, c'est très bien.  
> Le problème c'est que quand je liste les taches j'ai forcément les 20 dernières dans l'ordre ascendant.  
> J'aimerais bien pouvoir filtrer et trier.  
> Quelque chose comme `go run . list --filter "date > 'YYYY-MM-DD'; completed = true; title like '%implemenent%'" --order "date desc"`  
> Enfin peu importe le format, mais qu'on puisse filtrer et trier en somme.  
> Merci d'avance.  
> Cordialement,  
> Le client

## Run

```bash
go test ./...
go run . list
```
