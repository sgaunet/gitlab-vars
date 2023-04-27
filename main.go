package main

import (
	"flag"
	"fmt"
	"os"

	tty "github.com/mattn/go-tty"
	"github.com/sgaunet/gitlab-vars/internal/gitlabapi"
	"github.com/sgaunet/gitlab-vars/internal/gitlabvarsfile"
	"github.com/sgaunet/gitlab-vars/pkg/git"
)

var version string = "development"

func printVersion() {
	fmt.Println(version)
}

func main() {
	var (
		debugLevel  string
		projectId   int
		groupId     int
		vOption     bool
		helpOption  bool
		environment string
	)
	flag.StringVar(&debugLevel, "d", "error", "Debug level (info,warn,debug)")
	flag.StringVar(&environment, "e", "*", "environment")
	flag.BoolVar(&vOption, "v", false, "Get version")
	flag.IntVar(&projectId, "p", 0, "Project ID to get issues from")
	flag.IntVar(&groupId, "g", 0, "Group ID to get issues from (not compatible with -p option)")
	flag.BoolVar(&helpOption, "h", false, "Print help command line")
	flag.Parse()

	if vOption {
		printVersion()
		os.Exit(0)
	}
	if helpOption {
		printHelp()
		os.Exit(0)
	}

	if projectId != 0 && groupId != 0 {
		fmt.Fprintln(os.Stderr, "-p and -g options are incompatible")
		flag.PrintDefaults()
		os.Exit(1)
	}

	t, err := tty.Open()
	if err != nil {
		fmt.Fprintf(os.Stderr, "tty.Open: %s\n", err.Error())
		os.Exit(1)
	}

	defer t.Close()
	l := initTrace(t.Output(), debugLevel)

	if len(os.Getenv("GITLAB_TOKEN")) == 0 {
		l.Errorf("Set GITLAB_TOKEN environment variable")
		os.Exit(1)
	}
	if len(os.Getenv("GITLAB_URI")) == 0 {
		os.Setenv("GITLAB_URI", "https://gitlab.com")
	}

	if groupId == 0 && projectId == 0 {
		// Try to find git repository and project
		remoteOrigin, err := git.RetrieveRemoteOriginFromGitConfig()
		if err != nil {
			l.Errorf(err.Error())
			os.Exit(1)
		}
		project, err := gitlabapi.FindProject(remoteOrigin)
		if err != nil {
			l.Errorln("gitlab project not found")
			os.Exit(1)
		}
		l.Infoln("Project found: ", project.SshUrlToRepo)
		l.Infoln("Project found: ", project.Id)
		projectId = project.Id
	}

	if groupId != 0 {
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
		gitlabapi.ExpandAndPrintVars(v)
	}

	if projectId != 0 {
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

		currentDir, err := os.Getwd()
		if err != nil {
			l.Errorln(err.Error())
			os.Exit(1)
		}
		gitlabVarsFilePath, err := gitlabvarsfile.FindGitLabVarsFile(currentDir)
		if err == nil && gitlabVarsFilePath != "" {
			additionVars, err := gitlabvarsfile.ReadGitLabVarsFile(gitlabVarsFilePath)
			if err == nil {
				vNoneFiltered := gitlabapi.MergeVars(v, additionVars)
				v = gitlabapi.FilterVars(vNoneFiltered, environment)
			}
		}
		gitlabapi.ExpandAndPrintVars(v)
	}
}
