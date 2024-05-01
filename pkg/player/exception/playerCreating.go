package exception

import "fmt"

type PlayerCreating struct {
	PlayerID string
}

func (e *PlayerCreating) Error() string {
	return fmt.Sprintf("player with ID %s already exists", e.PlayerID)
}
