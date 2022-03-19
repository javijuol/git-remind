package infra

import (
	"github.com/javijuol/git-remind/domain"
	"os"
	"path/filepath"
)

var FilesystemRepos domain.GetReposByPathPattern = func(patterns domain.GetPathPatterns) (repos []string, err error) {
	pathPatterns, err := patterns()
	if err != nil {
		return
	}
	for _, pathPattern := range pathPatterns {
		paths, err2 := filepath.Glob(string(pathPattern))
		if err2 != nil {
			return repos, err2
		}
		for _, path := range paths {
			pathExists, _ := pathExists(path + "/.git")
			if pathExists {
				repos = append(repos, path)
			}
		}
	}
	return
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return err == nil, err
}
