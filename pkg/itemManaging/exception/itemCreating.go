package exception

type ItemCreating struct {
}

func (e *ItemCreating) Error() string {
	return "Item creation failed"
}
