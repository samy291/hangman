# Jeu du Pendu

Ce projet est le jeu du pendu codé en Golang.

## Fonctionnalités

- Jouer au jeu du pendu avec des mots aléatoires
- Possibilité de sauvegarder les parties mais pas de les charger
- Deux modes d'affichage : classique et ASCII

## Utilisation

Pour lancer le jeu, exécutez les deux commandes suivante :

go build -o hangman main/main.go
./hangman --classic "chemin du dictionnaire"
./hangman --ascii "chemin du dictionnaire"
