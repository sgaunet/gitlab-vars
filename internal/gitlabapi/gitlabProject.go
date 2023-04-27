package gitlabapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

func FindProject(remoteOrigin string) (*project, error) {
	projectName := filepath.Base(remoteOrigin)
	projectName = strings.ReplaceAll(projectName, ".git", "")
	// log.Infof("Try to find project %s in %s\n", projectName, os.Getenv("GITLAB_URI"))

	_, res, err := Request("search?scope=projects&search=" + projectName)
	if err != nil {
		return nil, err
	}

	var p []project
	err = json.Unmarshal(res, &p)
	if err != nil {
		return nil, err
	}

	for _, project := range p {
		// log.Debugln(project.Name)
		// log.Debugln(project.Id)
		// log.Debugln(project.HttpUrlToRepo)
		// log.Debugln(project.SshUrlToRepo)

		if project.SshUrlToRepo == remoteOrigin {
			return &project, err
		}
	}
	return nil, errors.New("project not found")
}

func GetProject(projectId int) (*GitlabProject, error) {
	var g GitlabProject
	uri := fmt.Sprintf("projects/%d", projectId)
	_, body, _ := Request(uri)
	if err := json.Unmarshal(body, &g); err != nil {
		return nil, err
	}
	// fmt.Println(g)
	return &g, nil
}

func (r *GitlabProject) GetGroupId() int {
	return r.Namespace.Id
}

func (r *GitlabProject) IsRootGroup() bool {
	return r.GetGroupId() == 0
}

func GetAllGroupParentId(groupId int) ([]int, error) {
	// retrieve group
	g, err := GetGroup(groupId)
	if err != nil {
		return nil, err
	}
	if g.IsRootGroup() {
		return []int{g.Id}, nil
	}
	parentGids, err := GetAllGroupParentId(g.ParentId)
	if err != nil {
		return nil, err
	}
	parentGids = append(parentGids, g.Id)
	return parentGids, nil
}

func (p *GitlabProject) GetAllVars(scope string) ([]Variable, error) {
	var v []Variable
	parents, err := GetAllGroupParentId(p.GetGroupId())
	if err != nil {
		return nil, err
	}

	for p := range parents {
		g, err := GetGroup(parents[p])
		if err != nil {
			return nil, err
		}
		// fmt.Println(g.Id, g.Name)
		vTmp, err := g.GetVarsOfGroup(scope)
		if err != nil {
			return nil, err
		}
		v = MergeVars(v, vTmp)
	}

	// Get all vars of project
	vTmp, err := p.GetVarsOfProject(scope)
	if err != nil {
		return nil, err
	}
	v = MergeVars(v, vTmp)

	return v, nil
}

func (p *GitlabProject) GetVarsOfProject(scope string) (Variables, error) {
	var v, vResult []Variable
	uri := fmt.Sprintf("projects/%d/variables", p.Id)
	_, body, _ := Request(uri)
	if err := json.Unmarshal(body, &v); err != nil {
		return nil, err
	}
	// fmt.Println(v)
	for _, varRange := range v {
		if IsVarPartOfScope(scope, varRange.EnvironmentScope) {
			vResult = append(vResult, varRange)
		}
	}
	return vResult, nil
}
