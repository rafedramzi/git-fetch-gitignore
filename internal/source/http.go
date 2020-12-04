package source

import (
	"bytes"
	"io"
	"net/http"
	"path"
	"text/template"

	"github.com/rafedramzi/fetch-gitignore/internal/config"
)

type SourceURL struct {
	name     string
	client   *http.Client
	url      string
	config   *config.Config
	cacheDir string
}

type URLTemplate struct {
	name string
}

// NewGitSource create new cache dir
func NewUrlSource(name string, url string, client *http.Client, conf *config.Config) *SourceURL {
	return &SourceURL{
		name:     name,
		client:   client,
		url:      url,
		config:   conf,
		cacheDir: path.Join(conf.CacheDir, name),
	}
}

func (c *SourceURL) GetFile(name string) (io.Reader, error) {
	t := template.Must(template.New("url").Parse(c.url))
	url := new(bytes.Buffer)
	err := t.Execute(url, URLTemplate{name})
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Get(url.String())
	if err != nil {
		return nil, err
	}
	// TODO: Cache
	return resp.Body, nil
}
