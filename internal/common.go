package internal

const (
	GitRepositorySource = iota + 1
	UrlSource
)

type Source struct {
	Source string
	Type   int
}

func NewSource(name string, sourceType int) (error, *Source) {
	if sourceType != GitRepositorySource && sourceType != UrlSource {
		return nil, nil
	}

	return nil, &Source{
		Source: name,
		Type:   sourceType,
	}
}
