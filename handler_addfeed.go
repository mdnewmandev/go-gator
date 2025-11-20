package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mdnewmandev/go-gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("command structure: addfeed <feed_name> <feed_url>")
	}

	feedName := cmd.Arguments[0]
	feedURL := cmd.Arguments[1]
	userID, err := uuid.Parse(s.Config.CurrentUserID)
	if err != nil {
		return fmt.Errorf("failed to parse user ID: %w", err)
	}

	feed, err := s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID:   		uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Name: 		feedName,
		Url:  		feedURL,
		UserID:     userID,
	})
	if err != nil {
		return fmt.Errorf("failed to add feed: %w", err)
	}

	fmt.Printf("Feed '%s' added successfully.\n", feed)
	return nil
}