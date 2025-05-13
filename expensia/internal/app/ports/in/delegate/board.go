package delegate

type BoardDelegate struct {
}

func Create() *BoardDelegate {
	return &BoardDelegate{}
}
