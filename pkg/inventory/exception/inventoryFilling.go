package exception

import "fmt"

type InventoryFilling struct {
	PlayerID string
	ItemID   uint64
}

func (e *InventoryFilling) Error() string {
	return fmt.Sprintf("filling inventory playerID: %s itemID: %d failed", e.PlayerID, e.ItemID)
}
