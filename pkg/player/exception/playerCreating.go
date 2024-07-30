package exception

import "fmt"

type PlayerCreating struct {
	PlayerID string
}

func (e *PlayerCreating) Error() string {
	return fmt.Sprintf("Player ID: %s already exists", e.PlayerID)
}
