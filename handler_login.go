package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("login command requires a username argument")
	}
	
	userName := cmd.Arguments[0]
	err := s.Config.SetUser(userName)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}
	
	fmt.Printf("User '%s' logged in successfully.\n", userName)
	return nil
}