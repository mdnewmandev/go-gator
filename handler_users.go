package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to retrieve users: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("No users found.")
		return nil
	}

	fmt.Println("Registered Users:")
	for _, user := range users {
		current := ""
		if user == s.Config.CurrentUserName {
			current = " (current)"
		}
		fmt.Printf("* %s%s\n", user, current)
	}
	return nil
}