package exception

import "fmt"

type PlayerItemsFinding struct {
	PlayerID string
	
}

func (e *PlayerItemsFinding) Error() string {
	return fmt.Sprintf("player with ID %s has no items", e.PlayerID)
}
