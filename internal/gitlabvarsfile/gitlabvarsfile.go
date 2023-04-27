package gitlabvarsfile

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sgaunet/gitlab-vars/internal/gitlabapi"
)

// FindGitLabVarsFile is a recursive function to find .gitlab-vars.json file
// in the current directory and in the parent directories
func FindGitLabVarsFile(path string) (string, error) {
	f := filepath.Join(path, ".gitlab-vars.json")
	if _, err := os.Stat(f); err == nil {
		return f, nil
	}
	if path == "/" {
		return "", errors.New("gitLab vars file not found")
	}
	return FindGitLabVarsFile(filepath.Dir(path))
}

func ReadGitLabVarsFile(filepath string) ([]gitlabapi.Variable, error) {
	jsonContent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var gitlabVarsFile GitLabVarsFile
	err = json.Unmarshal(jsonContent, &gitlabVarsFile)
	if err != nil {
		return nil, err
	}
	return gitlabVarsFile.Variables, nil
}
