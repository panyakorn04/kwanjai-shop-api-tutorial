package exception

type PlayerCoinShowing struct{}

func (e *PlayerCoinShowing) Error() string {
	return "Cannot show player coin"
}
