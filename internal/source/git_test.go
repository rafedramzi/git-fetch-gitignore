package source

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/rafedramzi/fetch-gitignore/internal/config"
	"github.com/stretchr/testify/assert"
)

var debug bool

func TestGetFileGitRepository(t *testing.T) {
	repoName := "dummy"
	workspace := path.Join(os.TempDir(), "fetchignore", "test_fetch_git_repsitory")
	repoDir := utilsCreateDummyGitRepository(t, workspace, repoName)
	source := NewGitSource(repoName, "file://"+workspace, utilsConfig(workspace))
	dummyData := "hello \n hello"
	if !debug {
		defer os.RemoveAll(repoDir)
	}

	subdirPath := path.Join(repoDir, "dir")
	err := os.Mkdir(subdirPath, 0755)
	if err != nil && !os.IsExist(err) {
		t.Fatal(err)
	}

	testCases := []struct {
		name     string
		fileName string
		dirPath  string
		data     string
	}{
		{
			name:     "Retrieve file at root",
			fileName: "boop",
			dirPath:  repoDir,
			data:     dummyData,
		},
		{
			name:     "Retrieve file at root full-name",
			fileName: "boop.ignorecase",
			dirPath:  repoDir,
			data:     dummyData,
		},
		{
			name:     "Retrieve file at sub-directiory",
			fileName: "boop",
			dirPath:  subdirPath,
			data:     dummyData,
		},
		{
			name:     "Retrieve file at sub-directiory full-name",
			fileName: "boop.gitignore",
			dirPath:  subdirPath,
			data:     dummyData,
		},
		// TODO: Handle bad case!
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			fileName := tc.fileName
			if !strings.HasSuffix(fileName, ".gitignore") {
				fileName += ".gitignore"
			}
			err = ioutil.WriteFile(path.Join(tc.dirPath, fileName), []byte(tc.data), 0644)
			if err != nil {
				t.Fatal(err)
			}
			// Test
			ans, err := source.GetFile(tc.fileName)
			assert.NoError(t, err)
			assert.Equal(t, dummyData, string(ans))
		})
	}

}

// NOTE: I feel like this test needs more work
func TestIntegrationFetchGitRepsitory(t *testing.T) {
	repositories := utilsRepositoriesData()

	cacheDir := path.Join(os.TempDir(), string(os.PathSeparator), "fetchignore")
	if !debug {
		defer os.RemoveAll(cacheDir)
	}
	conf := utilsConfig(cacheDir)
	repo := repositories[1]
	source := NewGitSource(repo.name, repo.source, conf)
	err := source.Sync(true)
	assert.NoError(t, err)
}

func TestFetchGitRepository(t *testing.T) {
	t.Skip("TODO: TestFetchGitRepsitory")
}

/*
 > Test Helpers
*/
type TestRepositoryData struct {
	name   string
	source string
}

func utilsCreateDummyGitRepository(t *testing.T, workspace string, repoName string) string {
	// cacheDir, err := ioutil.TempDir(path.Join(os.TempDir(), "fetchignore"), "test_fetch_git_repsitory")
	repoDir := path.Join(workspace, repoName)
	err := os.MkdirAll(repoDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err.Error())
	}
	git.PlainInit(workspace, false)
	return repoDir
}

func utilsRepositoriesData() []TestRepositoryData {
	return []TestRepositoryData{
		{name: "lc-gitignore", source: "localhost:3000/pwyll/gitignore.git"},
		{name: "lc-gitingore", source: "http://localhost:3000/pwyll/gitignore.git"},
		{name: "gh-gitingore", source: "ssh://git@github.com:github/gitignore.git"},
		{name: "gh-gitignore", source: "git@github.com:github/gitignore.git"},
	}
}

func utilsConfig(cacheDir string) *config.Config {
	dur, _ := time.ParseDuration("3s")
	return &config.Config{
		DefaultRepository: "",
		CacheDir:          cacheDir,
		ExpireDuration:    dur,
		Sources:           nil,
	}
}
