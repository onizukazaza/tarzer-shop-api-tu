package exception

import "fmt"

type AdminNotFound struct {
	AdminID string
}

func (e *AdminNotFound) Error() string {
	return fmt.Sprintf("Player with ID: %s was not found", e.AdminID)
}
