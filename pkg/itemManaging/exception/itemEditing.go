package exception

import "fmt"

type ItemEditing struct {
	ItemID uint64
}

func (e *ItemEditing) Error() string {
	return fmt.Sprintf("Item editing failed for item with ID %d", e.ItemID)
}
