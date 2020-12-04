package source

import (
	"io"
	"net/http"
	"path"

	"github.com/rafedramzi/fetch-gitignore/internal/config"
)

// fetch from gitignore.io similar to https://github.com/tj/git-extras/blob/master/Commands.md#git-ignore

const GITIGNORE_URL = `https://www.toptal.com/developers/gitignore/api/{{.name}}`

type SourceGitignoreio struct {
	sourceUrl *SourceURL
}

func NewGitignoreioSource(name string, url string, client *http.Client, conf *config.Config) *SourceGitignoreio {
	return &SourceGitignoreio{
		sourceUrl: &SourceURL{
			name:     name,
			client:   client,
			url:      url,
			config:   conf,
			cacheDir: path.Join(conf.CacheDir, name),
		},
	}
}

func (c *SourceGitignoreio) GetFile(name string) (io.Reader, error) {
	return c.sourceUrl.GetFile(name)
}

func (c *SourceGitignoreio) GetFiles(names []string) (io.Reader, error) {
	// TODO:
	// for _, name := range names {
	// get the file individually
	// }
	// return the merged file
	return nil, nil
}
