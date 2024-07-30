package exception

type Uauthorized struct {}

func (e *Uauthorized) Error() string {
    return "Unauthorized access"
}