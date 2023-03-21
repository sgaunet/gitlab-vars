package gitlabvarsfile

import "github.com/sgaunet/gitlab-vars/internal/gitlabapi"

type GitLabVarsFile struct {
	Variables []gitlabapi.Variable `json:"variables"`
}
