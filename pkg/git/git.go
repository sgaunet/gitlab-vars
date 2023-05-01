package git

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sgaunet/gitlab-vars/internal/gitlabapi"
	"gopkg.in/ini.v1"
)

func findGitRepository() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for cwd != "/" {
		stat, err := os.Stat(cwd + string(os.PathSeparator) + ".git")
		if err == nil {
			// !TODO: submodule case
			if stat.IsDir() {
				return cwd, err
			}
		}
		cwd = filepath.Dir(cwd)
	}
	return "", errors.New(".git not found")
}

func getRemoteOrigin(gitConfigFile string) (string, error) {
	cfg, err := ini.Load(gitConfigFile)
	if err != nil {
		return "", fmt.Errorf("fail to read file %v: %w", gitConfigFile, err)
	}
	url := cfg.Section("remote \"origin\"").Key("url").String()
	return url, nil
}

func retrieveRemoteOriginFromGitConfig() (string, error) {
	gitFolder, err := findGitRepository()
	if err != nil {
		return "", err
	}
	remoteOrigin, err := getRemoteOrigin(gitFolder + string(os.PathSeparator) + ".git" + string(os.PathSeparator) + "config")
	if err != nil {
		return "", err
	}
	return remoteOrigin, nil
}

func TryToFindGitlabProjectFromGitConfig() (int, error) {
	remoteOrigin, err := retrieveRemoteOriginFromGitConfig()
	if err != nil {
		return 0, err
	}
	project, err := gitlabapi.FindProject(remoteOrigin)
	if err != nil {
		return 0, err
	}
	return project.Id, nil
}
