package exception

type OAuth2Processing struct {
}

func (e *OAuth2Processing) Error() string {
	return "OAuth2 processing error"
}
