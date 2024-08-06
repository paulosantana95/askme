package main

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(0)

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		panic(err)
	}

	var stderr bytes.Buffer
	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Println("Error running tern migrate:", stderr.String())
		panic(err)
	}

	log.Println("Migration completed! 🎉")
}
