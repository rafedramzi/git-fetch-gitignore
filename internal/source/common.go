package source

import (
	"io"
	"strings"
	"time"
)

type Source interface {
	// Will clone or refresh the source
	// Sync(force bool) error
	// Will getch the stile
	GetFile(string) (io.Reader, error)
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

// HasExpired determinte wheter time is expired or not
func HasExpired(t time.Time, expDuration time.Duration) bool {
	if time.Since(t) < expDuration {
		return false
	}
	return true
}
