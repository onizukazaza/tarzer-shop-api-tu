package exception

type Logout struct{}

func (e *Logout) Error() string {
    return "User logged out failed"
}