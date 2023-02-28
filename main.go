package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sgaunet/gitlab-vars/internal/gitlabapi"
	"github.com/sirupsen/logrus"
)

var version string = "development"

func printVersion() {
	fmt.Println(version)
}

func main() {
	var (
		debugLevel string
		projectId  int
		groupId    int
		vOption    bool
		// err        error
	)
	// Parameters treatment (except src + dest)
	flag.StringVar(&debugLevel, "d", "error", "Debug level (info,warn,debug)")
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
		fmt.Fprintln(os.Stderr, "-p and -g option are incompatible")
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
		_, err := gitlabapi.GetGroup(groupId)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	if projectId != 0 {
		p, err := gitlabapi.GetProject(projectId)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		v, err := p.GetAllVars("preprod-mtrg")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		for _, vv := range v {
			// os.Setenv(vv.Key, vv.Value)
			fmt.Printf("(%s) %20s=%s\n", vv.EnvironmentScope, vv.Key, vv.Value)
		}
		// for _, vv := range v {
		// 	if vv.Raw && vv.EnvironmentScope == "preprod" {
		// 		os.Setenv(vv.Key, vv.Value)
		// 		fmt.Println(vv.Key, "=", vv.Value)
		// 	}
		// }
		// for _, vv := range v {
		// 	if !vv.Raw && vv.EnvironmentScope == "preprod" {
		// 		os.Setenv(vv.Key, vv.Value)
		// 		fmt.Println(vv.Key, "=", ExpandEnv(vv.Value))
		// 	}
		// }
	}

	// _, err = gitlabapi.GetVarsOfProject(projectId)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// os.Setenv("TOTO", "123456")
	// os.Setenv("TOTO_TYU", "123456")
	// fmt.Println(ExpandEnv("GLOBAL=$TOTO"))
	// fmt.Println(ExpandEnv("GLOBAL=$TOTO_TYU"))
	// fmt.Println(ExpandEnv("GLOBAL=${TOTO_TYU}"))
}

// ExpandEnv replaces ${var} or $var in the string based on the values of the
// current environment variables. The replacement is case-sensitive. References
// to undefined variables are replaced by the empty string. A default value can
// be given by using the form ${var:-default value}. The default value is used
// only if var is unset or empty. A different value can be given by using the
// form ${var:default value}. The default value is used if var is unset.
// References to other variables are expanded as the string is processed.
// Recursive references are not allowed. If there is an error in the syntax of
// the variable reference, the reference is replaced by the empty string.
func ExpandEnv(s string) string {
	return os.Expand(s, func(v string) string {
		return os.Getenv(v)
	})
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
