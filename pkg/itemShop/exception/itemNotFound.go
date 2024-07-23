package exception

import "fmt"

type ItemNotFound struct {
	itemID uint64
}

func (e *ItemNotFound) Error() string {
	return fmt.Sprintf("itemID: %d was not found", e.itemID)
}