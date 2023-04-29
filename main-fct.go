package main

import (
	"github.com/sgaunet/gitlab-vars/internal/gitlabapi"
	"github.com/sgaunet/gitlab-vars/internal/gitlabvarsfile"
)

func printVarsOfProject(projectId int, environment string) error {
	p, err := gitlabapi.GetProject(projectId)
	if err != nil {
		return err
	}
	v, err := p.GetAllVars(environment)
	if err != nil {
		return err
	}
	v = gitlabapi.FilterVars(v, environment)
	v, err = gitlabvarsfile.UpdateVarsWithGitlabVarsFileIfExist(v, environment)
	if err != nil {
		return err
	}
	gitlabapi.ExpandAndPrintVars(v)
	return nil
}

func printVarsOfGroup(groupId int, environment string) error {
	g, err := gitlabapi.GetGroup(groupId)
	if err != nil {
		return err
	}
	v, err := g.GetAllVars(environment)
	if err != nil {
		return err
	}
	v = gitlabapi.FilterVars(v, environment)
	v, err = gitlabvarsfile.UpdateVarsWithGitlabVarsFileIfExist(v, environment)
	if err != nil {
		return err
	}
	gitlabapi.ExpandAndPrintVars(v)
	return nil
}
