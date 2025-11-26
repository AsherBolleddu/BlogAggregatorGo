package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete users: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
