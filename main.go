package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a new scanner for reading input
	scanner := bufio.NewScanner(os.Stdin)

	// Introduction
	fmt.Println("Welcome to the GitHub Collaboration Role Selector!")
	fmt.Println("Choose a role: Owner, Admin, Collaborator, or Viewer.")
	fmt.Print("Enter your role: ")

	// Capture user input
	scanner.Scan()
	role := scanner.Text() // Get the input as a string

	// Trim any leading/trailing whitespace and convert input to lowercase
	role = strings.TrimSpace(role)
	role = strings.ToLower(role)

	// Switch statement to describe the role's permissions
	fmt.Println("\nRole Permissions:")
	switch role {
	case "owner":
		fmt.Println("As an Owner, you have full control over the repository, including managing collaborators and settings!")
		// fallthrough
	case "admin":
		fmt.Println("As an Admin, you can manage repository settings, but you don't have full control over billing and settings like an Owner.")
		// fallthrough
	case "collaborator":
		fmt.Println("As a Collaborator, you can push code to the repository and contribute to the project!")
		// fallthrough
	case "viewer":
		fmt.Println("As a Viewer, you can only view the repository. No changes are allowed.")
	default:
		fmt.Println("Hmm... that's not a valid role. Perhaps you want to create your own role?")
	}

	// Bonus feature: Fallthrough to see additional details about permissions
	fmt.Println("\nBonus: What can you do with your role?")
	switch role {
	case "owner":
		fmt.Println("You have access to billing, repository settings, and the ability to remove anyone from the repository!")
		fallthrough
	case "admin":
		fmt.Println("You can invite new collaborators and change repository visibility.")
		fallthrough
	case "collaborator":
		fmt.Println("You can push code and create branches but cannot change repository settings.")
		fallthrough
	case "viewer":
		fmt.Println("You can only view the repository and open issues, but cannot contribute code.")
	default:
		fmt.Println("You're an unknown entity, so no permissions have been set for you!")
	}
}
