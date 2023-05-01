package main

import (
	"fmt"
)

func printHelp() {
	fmt.Println(`
Usage: gitlab-vars [options]
Options:
	-g <group_id>		  Group id
	-p <project_id>		Project id
	-e <environment>	Environment
	-h			          Help
	-d <debuglevel>		Debug level (info, error, debug)
	-v		          	Version

gitlab-vars is a tool to print the variables of a gitlab project. You can use it to check how variables will be overrided.
But don't expect from this tool to give you the value of the variables after a workflow for example.
It won't read .gitlab-ci.yml file.

Every parameters are optionals.

By default, gitlab-vars try to find the project id from the current git repository.
How it works ?
	1. Find the .git directory
	2. Read the git config file
	3. Get the remote origin url
	4. Find the project id from the remote origin url

If the project id is not found, you can specify the group id or the project id.
	`)
}
