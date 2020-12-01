package source

import (
	"io"
	"strings"
)

type Source interface {
	// Will clone or refresh the source
	Sync() error
	// Will getch the stile
	Get(string) (io.Reader, error)
}

// Fetch get file + refresh the repository if needed
func Fetch(source Source, files []string) (io.Reader, error) {
	err := source.Sync()
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		_, err := source.Get(file)
		// TODO: Accumulate Error
		if err != nil {
			return nil, err
		}
	}

	// TODO:
	return nil, nil
}

// RemoveSpecialCharacter remove ,, /, \ from string
// Assumption derived from https://en.wikipedia.org/wiki/Filename#Reserved_characters_and_words
func RemoveSpecialCharacter(input string) string {
	return strings.NewReplacer(
		`,`, `-`,
		`/`, `-`,
		`\`, `-`,
	).Replace(input)
}
