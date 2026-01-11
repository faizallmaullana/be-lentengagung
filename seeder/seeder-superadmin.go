package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	database "github.com/faizallmaullana/lenteng-agung/backend/db/db_connection"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/domains/repo"
	"github.com/faizallmaullana/lenteng-agung/backend/internal/models"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}
}

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func main() {
	email := getenv("SUPERADMIN_EMAIL", "superadmin@example.com")
	password := getenv("SUPERADMIN_PASSWORD", "superadmin")
	nik := getenv("SUPERADMIN_NIK", "0000000000")
	name := getenv("SUPERADMIN_NAME", "Super Admin")

	prov, err := database.NewGormProviderFromEnv()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer prov.Close()

	r := repo.NewAuthRepo(prov)

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}

	u := &models.User{
		Email:        email,
		PasswordHash: string(hashed),
		IsActive:     true,
		Role:         "superadmin",
		ApprovedAt:   time.Now(),
	}

	if err := r.CreateUser(u); err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	p := &models.Profile{
		UserID: u.ID,
		NIK:    nik,
		Name:   name,
	}

	if err := r.CreateProfile(p); err != nil {
		log.Fatalf("failed to create profile: %v", err)
	}

	if err := r.ApproveUser(u.ID); err != nil {
		log.Fatalf("failed to approve user: %v", err)
	}

	fmt.Printf("superadmin created/updated: %s\n", email)
}
