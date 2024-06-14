package main

import (
	"fmt"
	"log"
	"os/exec"
)

// executeCommand runs a shell command and returns its output or an error
func executeCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	err := cmd.Run()
	return err
}

func main() {
	fmt.Println("Starting environment setup...")

	// 1. Clone the GitHub repository
	cloneGitRepository()

	// Change directory to the cloned repository
	/// setup mqtt
	changeDirectory("Sakura-Gateway")
	// 2. Start the initial Docker compose
	fmt.Println("Starting initial Docker Compose...")
	if err := executeCommand("docker-compose", "up", "-d"); err != nil {
		log.Fatalf("Failed to start initial Docker Compose: %v", err)
	}

	// 3. Start the DB and API Docker compose
	fmt.Println("Starting DB and API Docker Compose...")
	if err := executeCommand("docker-compose", "-f", "docker-compose.db.yml", "up", "-d"); err != nil {
		log.Fatalf("Failed to start DB and API Docker Compose: %v", err)
	}

	// 4. Enter the API container and run migration
	apiContainerName := "api_container"          // Replace with your API container name
	migrationCommand := "your_migration_command" // Replace with your migration command
	fmt.Println("Running migration in API container:", apiContainerName)
	if err := executeCommand("docker", "exec", apiContainerName, migrationCommand); err != nil {
		log.Fatalf("Failed to run migration: %v", err)
	}

	fmt.Println("Environment setup completed successfully!")
}

func cloneGitRepository() {
	sakura_pi_node := "https://github.com/niwaniwa/Sakura-Pi-Node.git"
	sakura_api := "https://github.com/niwaniwa/Sakura-API.git"
	sakura_web := "https://github.com/niwaniwa/Sakura-Web.git"
	Gateway := "https://github.com/niwaniwa/Sakura-Gateway.git"

	fmt.Println("Cloning repository:", sakura_pi_node)
	if err := executeCommand("git", "clone", sakura_pi_node); err != nil {
		log.Fatalf("Failed to clone repository: %v", err)
	}

	fmt.Println("Cloning repository:", sakura_api)
	if err := executeCommand("git", "clone", sakura_api); err != nil {
		log.Fatalf("Failed to clone repository: %v", err)
	}

	fmt.Println("Cloning repository:", sakura_web)
	if err := executeCommand("git", "clone", sakura_web); err != nil {
		log.Fatalf("Failed to clone repository: %v", err)
	}

	fmt.Println("Cloning repository:", Gateway)
	if err := executeCommand("git", "clone", Gateway); err != nil {
		log.Fatalf("Failed to clone repository: %v", err)
	}

}

func changeDirectory(path string) {
	if err := executeCommand("cd", path); err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}
}
