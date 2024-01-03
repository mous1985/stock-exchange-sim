package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Chemin du fichier à lire
	fileName := "./examples/" + os.Args[1]

	// Lecture du contenu du fichier
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	// Conversion du contenu du fichier en une chaîne de caractères
	input := string(content)

	// Séparation de la chaîne d'entrée par lignes
	lines := strings.Split(input, "\n")

	// Parcours des lignes pour extraire les informations pertinentes
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Ignorer les lignes vides ou les commentaires
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Analyse des différentes parties de la ligne
		if strings.HasPrefix(line, "optimize:") {
			// Extraction des éléments à optimiser
			parts := strings.Split(strings.TrimPrefix(line, "optimize:"), "|")
			for _, part := range parts {
				fmt.Println("Optimiser :", part)
			}
		} else if strings.Contains(line, ":") {
			// Extraction des informations de stock et de processus
			data := strings.Split(line, ":")
			fmt.Println("Nom :", data[0])
			fmt.Println("Valeur :", data[1])
		}
	}
}
