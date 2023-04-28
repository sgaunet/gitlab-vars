package main

import (
	"os"

	"github.com/sgaunet/gitlab-vars/internal/gitlabapi"
	"github.com/sgaunet/gitlab-vars/internal/gitlabvarsfile"
	"github.com/sirupsen/logrus"
)

func printVarsOfProject(projectId int, environment string, l *logrus.Logger) {
	p, err := gitlabapi.GetProject(projectId)
	if err != nil {
		l.Errorln(err.Error())
		os.Exit(1)
	}
	v, err := p.GetAllVars(environment)
	if err != nil {
		l.Errorln(err.Error())
		os.Exit(1)
	}

	v = gitlabapi.FilterVars(v, environment)
	v, err = gitlabvarsfile.UpdateVarsWithGitlabVarsFileIfExist(v, environment)
	if err != nil {
		l.Errorln(err.Error())
		os.Exit(1)
	}
	gitlabapi.ExpandAndPrintVars(v)
}

func printVarsOfGroup(groupId int, environment string, l *logrus.Logger) {
	g, err := gitlabapi.GetGroup(groupId)
	if err != nil {
		l.Errorln(err.Error())
		os.Exit(1)
	}
	v, err := g.GetAllVars(environment)
	if err != nil {
		l.Errorln(err.Error())
		os.Exit(1)
	}
	v = gitlabapi.FilterVars(v, environment)
	v, err = gitlabvarsfile.UpdateVarsWithGitlabVarsFileIfExist(v, environment)
	if err != nil {
		l.Errorln(err.Error())
		os.Exit(1)
	}
	gitlabapi.ExpandAndPrintVars(v)
}
