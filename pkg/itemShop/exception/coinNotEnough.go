package exception

import "fmt"

type CoinNotEnough struct {
	PlayerID string
}

func (e *CoinNotEnough) Error() string {
	return fmt.Sprintf("playerID: %s coin not enough", e.PlayerID)
}
