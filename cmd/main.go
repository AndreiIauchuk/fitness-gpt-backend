package main

import (
	db "fitness-gpt-backend/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")
	db.Init();
}