package exception

type ItemCounting struct {
}

func (e *ItemCounting) Error() string {
	return "Error while counting items"
}
