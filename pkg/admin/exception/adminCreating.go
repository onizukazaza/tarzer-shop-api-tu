package exception

import "fmt"

type AdminCreating struct {
	AdminID string
}

func (e *AdminCreating) Error() string {
	return fmt.Sprintf("Player ID: %s already exists", e.AdminID)
}
