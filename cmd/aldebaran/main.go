package main

import (
	"github.com/irvanherz/aldebaran/internal/config"
	"github.com/irvanherz/aldebaran/internal/routes"
)

func main() {
	print("ALDEBARAN v1")
	print("API Server is starting...")

	config.SetupDB()
	r := routes.Setup()
	r.Run(":3001")
}
