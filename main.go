package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sgaunet/gitlab-vars/internal/gitlabapi"
	"github.com/sgaunet/gitlab-vars/internal/gitlabvarsfile"
	"github.com/sirupsen/logrus"
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
		environment string
	)
	flag.StringVar(&debugLevel, "d", "error", "Debug level (info,warn,debug)")
	flag.StringVar(&environment, "e", "*", "environment to filter variables")
	flag.BoolVar(&vOption, "v", false, "Get version")
	flag.IntVar(&projectId, "p", 0, "Project ID to get issues from")
	flag.IntVar(&groupId, "g", 0, "Group ID to get issues from (not compatible with -p option)")
	flag.Parse()

	if vOption {
		printVersion()
		os.Exit(0)
	}

	if debugLevel != "info" && debugLevel != "error" && debugLevel != "debug" {
		logrus.Errorf("debuglevel should be info or error or debug\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if projectId != 0 && groupId != 0 {
		fmt.Fprintln(os.Stderr, "-p and -g options are incompatible")
		flag.PrintDefaults()
		os.Exit(1)
	}
	initTrace(debugLevel)
	if len(os.Getenv("GITLAB_TOKEN")) == 0 {
		logrus.Errorf("Set GITLAB_TOKEN environment variable")
		os.Exit(1)
	}
	if len(os.Getenv("GITLAB_URI")) == 0 {
		os.Setenv("GITLAB_URI", "https://gitlab.com")
	}

	if groupId == 0 && projectId == 0 {
		// Try to find git repository and project
		gitFolder, err := findGitRepository()
		if err != nil {
			logrus.Errorf("Folder .git not found")
			os.Exit(1)
		}
		remoteOrigin := GetRemoteOrigin(gitFolder + string(os.PathSeparator) + ".git" + string(os.PathSeparator) + "config")
		project, err := findProject(remoteOrigin)
		if err != nil {
			logrus.Errorln(err.Error())
			os.Exit(1)
		}
		logrus.Infoln("Project found: ", project.SshUrlToRepo)
		logrus.Infoln("Project found: ", project.Id)
		projectId = project.Id
	}

	if groupId != 0 {
		g, err := gitlabapi.GetGroup(groupId)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		v, err := g.GetAllVars(environment)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		gitlabapi.ExpandAndPrintVars(v, environment)
	}

	if projectId != 0 {
		p, err := gitlabapi.GetProject(projectId)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		v, err := p.GetAllVars(environment)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		currentDir := os.Getenv("PWD")
		gitlabVarsFilePath, err := gitlabvarsfile.FindGitLabVarsFile(currentDir)
		if err == nil && gitlabVarsFilePath != "" {
			additionVars, err := gitlabvarsfile.ReadGitLabVarsFile(gitlabVarsFilePath)
			fmt.Println("trouve ", len(additionVars))
			if err == nil {
				// Merge vars
				v = gitlabapi.MergeVars(v, additionVars)
			}
		}

		gitlabapi.ExpandAndPrintVars(v, environment)
	}

	// os.Setenv("TOTO", "123456")
	// os.Setenv("TOTO_TYU", "123456")
	// fmt.Println(ExpandEnv("GLOBAL=$TOTO"))
	// fmt.Println(ExpandEnv("GLOBAL=$TOTO_TYU"))
	// fmt.Println(ExpandEnv("GLOBAL=${TOTO_TYU}"))
}

func initTrace(debugLevel string) {
	// Log as JSON instead of the default ASCII formatter.
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	// logrus.SetFormatter(&logrus.TextFormatter{
	// 	DisableColors: true,
	// 	FullTimestamp: true,
	// })

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	switch debugLevel {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
}
