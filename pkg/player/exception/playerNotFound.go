package exception

import "fmt"

type PlayerNotFound struct {
	PlayerID string
}

func (e *PlayerNotFound) Error() string {
	return fmt.Sprintf("player with ID %s not found", e.PlayerID)
}
