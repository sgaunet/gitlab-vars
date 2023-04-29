package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sgaunet/gitlab-vars/pkg/git"
)

var version string = "development"

func printVersion() {
	fmt.Println(version)
}

func main() {
	var (
		projectId   int
		groupId     int
		vOption     bool
		helpOption  bool
		environment string
		err         error
	)
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

	// Check environment variables
	if len(os.Getenv("GITLAB_TOKEN")) == 0 {
		fmt.Fprintln(os.Stderr, "Set GITLAB_TOKEN environment variable")
		os.Exit(1)
	}
	if len(os.Getenv("GITLAB_URI")) == 0 {
		os.Setenv("GITLAB_URI", "https://gitlab.com")
	}

	// No projectid or groupid specified, try to find it from git config
	if groupId == 0 && projectId == 0 {
		projectId, err = git.TryToFindGitlabProjectFromGitConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "gitlab project from gitconfig not found: %s", err.Error())
			os.Exit(1)
		}
	}
	if groupId != 0 {
		err = printVarsOfGroup(groupId, environment)
		if err != nil {
			fmt.Fprintf(os.Stderr, "printVarsOfGroup: %s\n", err.Error())
			os.Exit(1)
		}
	}
	if projectId != 0 {
		err = printVarsOfProject(projectId, environment)
		if err != nil {
			fmt.Fprintf(os.Stderr, "printVarsOfProject: %s\n", err.Error())
			os.Exit(1)
		}
	}
}
