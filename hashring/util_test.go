package hashring

type fakeMember struct {
	address  string
	identity string
	labels   map[string]string
}

func (f fakeMember) GetAddress() string {
	return f.address
}

func (f fakeMember) Label(key string) (value string, has bool) {
	if f.labels == nil {
		return "", false
	}

	l, ok := f.labels[key]
	return l, ok
}

func (f fakeMember) Identity() string {
	if f.identity != "" {
		return f.identity
	}
	return f.address
}
