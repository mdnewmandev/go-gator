package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mdnewmandev/go-gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("register command requires a name argument")
	}

	name := cmd.Arguments[0]
	_, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code.Name() == "unique_violation" {
			log.Printf("user with name '%s' already exists", name)
			os.Exit(1)
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	err = s.Config.SetUser(name)
	if err != nil {
		return fmt.Errorf("failed to set current user: %w", err)
	}

	fmt.Printf("User '%s' registered successfully.\n", name)
	log.Printf("Created user: %+v\n", name) // Log user data for debugging

	return nil
}
