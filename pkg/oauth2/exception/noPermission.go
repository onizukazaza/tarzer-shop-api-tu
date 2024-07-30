package exception

type NoPermission struct{}

func (e *NoPermission) Error() string {
    return "You do not have permission to perform this action."
}