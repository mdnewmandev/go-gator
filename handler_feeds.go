package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to retrieve feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Println("User Feeds:")
	for _, feed := range feeds {
		createdByRow, err := s.db.GetUserByID(context.Background(), feed.UserID)
		var createdBy string
		if err != nil {
			createdBy = "Unknown"
		} else {
			createdBy = createdByRow.Name
		}
		fmt.Printf("- %s %s\n", feed.Name, feed.Url)
		fmt.Printf("- %s\n", createdBy)
	}
	return nil
}