package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/spf13/pflag"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var _ = template.Must
var cacheDuration = time.Duration(30*24) * time.Hour
var UserAgent = "Fetch-Gitignore/0.0"
var CACHE_FILE_FORMAT = "%s.gitignore"
var DefaultClient = &http.Client{}

var (
	pflagCache         = pflag.Bool("no-cache", false, "Ignore local cache")
	pflagHelp          = pflag.Bool("help", false, "Print this page")
	pflagCacheLifetime = pflag.Duration("cache-duration", cacheDuration, "Set cache duration, this option will override the configuration file, if the file is over the duration it will refetch the file")
)

type fetchTemplate struct {
	Lang string
}

type repository struct {
	name     string
	patterns *template.Template
}

func createRepository(name string, patterns string) *repository {
	repo := &repository{}
	repo.name = name
	repo.patterns = template.Must(template.New("letter").Parse(patterns))
	return repo
}

// GetURL render the url pattern
func (repo *repository) GetURL(tpl *fetchTemplate) (string, error) {
	var buf bytes.Buffer
	err := repo.patterns.Execute(&buf, tpl)

	if err != nil {
		return "", err
	}

	return buf.String(), nil

}

func printUsage() {
	pflag.Usage()
	os.Exit(1)
}

func printFlag() {
	log.Debug("pflagCache:\t", *pflagCache)
	log.Debug("pflagHelp:\t", *pflagHelp)
	log.Debug("pflagCacheLifetime:\t", *pflagCacheLifetime)
}

func fetchIgnoreFile(url string) ([]byte, error) {
	// TODO: Use context and integrate with ctrl+c
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Debugf("Fetching ignore file at %s", req.URL)

	if resp.StatusCode < http.StatusOK && resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("Http status code not not within 2xx, got: %d while fetching %s", resp.StatusCode, req.URL)
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func storeCacheFile(repo *repository, tpl *fetchTemplate, result []byte) error {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	repoCacheDir := filepath.Join(cacheDir, "fetch-gitignore", repo.name)
	err = os.MkdirAll(repoCacheDir, 0755)
	if err != nil {
		return err
	}

	fileName := filepath.Join(repoCacheDir, fmt.Sprintf(CACHE_FILE_FORMAT, tpl.Lang))

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	if _, err = f.Write(result); err != nil {
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	return nil
}

func getIgnoreFile(repo *repository, tpl *fetchTemplate, storeCache bool) ([]byte, error) {
	url, err := repo.GetURL(tpl)
	if err != nil {
		return nil, err
	}

	result, err := fetchIgnoreFile(url)
	if err != nil {
		return nil, err
	}

	if storeCache {
		if err = storeCacheFile(repo, tpl, result); err != nil {
			return nil, err
		}
	}

	return result, nil
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		pflag.PrintDefaults()
	}
	viper.SetConfigName("example_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/fetch-ingore")
	viper.AddConfigPath("$HOME/.config/fetch-ignore")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Debugf("Error config file: %s \n", err)
		log.Info("Using default config file")
	}

	// TODO: Set set default config
	// TODO: Set cache

	// TODO: Invalidate cache

	// TODO: Fetch File
	repo := createRepository("github", "https://raw.githubusercontent.com/github/gitignore/master/{{.Lang}}.gitignore")
	tpl := &fetchTemplate{
		Lang: "Go",
	}
	result, err := getIgnoreFile(repo, tpl, *pflagCache)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))

	pflag.Parse()

	if pflag.NArg() < 1 || *pflagHelp {
		printUsage()
	}

	printFlag()
}
