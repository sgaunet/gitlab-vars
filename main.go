package main

import (
	"flag"
	"fmt"
	"os"

	tty "github.com/mattn/go-tty"
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

	// Init logger
	t, err := tty.Open()
	if err != nil {
		fmt.Fprintf(os.Stderr, "tty.Open: %s\n", err.Error())
		os.Exit(1)
	}
	defer t.Close()
	l := initTrace(t.Output(), debugLevel)

	// Check environment variables
	if len(os.Getenv("GITLAB_TOKEN")) == 0 {
		l.Errorf("Set GITLAB_TOKEN environment variable")
		os.Exit(1)
	}
	if len(os.Getenv("GITLAB_URI")) == 0 {
		os.Setenv("GITLAB_URI", "https://gitlab.com")
	}

	// No projectid or groupid specified, try to find it from git config
	if groupId == 0 && projectId == 0 {
		projectId, err = git.TryToFindGitlabProjectFromGitConfig(l)
		if err != nil {
			l.Errorf("gitlab project from gitconfig not found: %s", err.Error())
			os.Exit(1)
		}
	}
	if groupId != 0 {
		printVarsOfGroup(groupId, environment, l)
	}
	if projectId != 0 {
		printVarsOfProject(projectId, environment, l)
	}
}
