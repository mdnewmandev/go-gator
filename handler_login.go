package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("login command requires a username argument")
	}

	userName := cmd.Arguments[0]	
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		log.Printf("user with name '%s' does not exist", userName)
		os.Exit(1)
	}

	err = s.Config.SetUser(userName)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	
	fmt.Printf("User '%s' logged in successfully.\n", userName)
	return nil
}