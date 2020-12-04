package source

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/rafedramzi/fetch-gitignore/internal/config"
	log "github.com/sirupsen/logrus"
)

type SourceGitRepository struct {
	name     string
	url      string
	config   *config.Config
	cacheDir string
}

// NewGitSource create new cache dir
func NewGitSource(name string, url string, conf *config.Config) *SourceGitRepository {
	return &SourceGitRepository{
		name:     name,
		url:      url,
		config:   conf,
		cacheDir: path.Join(conf.CacheDir, name),
	}
}

// CacheDir return git source cache director
func (s *SourceGitRepository) CacheDir() string {
	return s.cacheDir
}

// Sync clone or update git repository to local cache directory
func (s *SourceGitRepository) Sync(force bool) error {
	repo, err := git.PlainOpen(s.CacheDir())
	if err != nil {
		if err != git.ErrRepositoryNotExists {
			return err
		}
		// Try to clone if not exists
		repo, err = git.PlainCloneContext(context.Background(), s.CacheDir(), false, &git.CloneOptions{
			URL:      s.url,
			Progress: os.Stdout, // TODO: Remove this
		})

		if err != nil {
			return err
		}
		return nil
	}

	// Check Expiration
	finfo, err := os.Stat(s.CacheDir())
	if HasExpired(finfo.ModTime(), s.config.ExpireDuration) && !force {
		log.Infof("Source %s exists and not expired yet", s.name)
		// Not Expired Yet! TODO: Probably should throw error intead
		return nil
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	log.Infof("Pulling latest %s repository", s.name)
	pullOptions := &git.PullOptions{}
	worktree.Pull(pullOptions)
	log.Infof("Source %s is up to date", s.name)
	return nil
}

// GetFile get file from git repository that stored in local
func (s *SourceGitRepository) GetFile(fileName string) ([]byte, error) {
	if !strings.HasSuffix(fileName, ".gitignore") {
		fileName += ".gitignore"
	}
	data, err := ioutil.ReadFile(path.Join(s.CacheDir(), fileName))
	if err != nil {
		return nil, err
	}
	return data, nil
}
