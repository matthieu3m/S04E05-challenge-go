package main

import "fmt"

func main() {
	var name string
	var task string
	var question string
	fmt.Println("Hello, merci d'inserer votre nom:")
	fmt.Scanln(&name)

	fmt.Println("Souhaitez-vous ajouter une tache ? y/n:")
	fmt.Scanln(&question)

	if question == "y" {
		fmt.Println("Merci, maintenant inseré une nouvelle tâche:")
		fmt.Scanln(&task)
		fmt.Printf("Vous êtes donc %v et vous avez introduit la tâche suivante: %v\n", name, task)
	} else {
		fmt.Printf("%v vous n'avez pas voulu inséré de tâche\n", name)
	}

}
