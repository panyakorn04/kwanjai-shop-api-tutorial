package exception

import "fmt"

type ItemArchiving struct {
	ItemID uint64
}

func (e *ItemArchiving) Error() string {
	return fmt.Sprintf("Item archiving failed for item with ID %d", e.ItemID)
}
