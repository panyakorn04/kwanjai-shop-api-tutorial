package exception

type CoinAdding struct{}

func (e *CoinAdding) Error() string {
	return "Cannot add coin"
}
