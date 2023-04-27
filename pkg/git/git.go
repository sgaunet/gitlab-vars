package git

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

func FindGitRepository() (string, error) {
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

func GetRemoteOrigin(gitConfigFile string) (string, error) {
	cfg, err := ini.Load(gitConfigFile)
	if err != nil {
		return "", fmt.Errorf("fail to read file %v: %w", gitConfigFile, err)
	}
	url := cfg.Section("remote \"origin\"").Key("url").String()
	return url, nil
}
