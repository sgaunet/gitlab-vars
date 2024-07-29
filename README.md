[![Go Report Card](https://goreportcard.com/badge/github.com/sgaunet/gitlab-vars)](https://goreportcard.com/report/github.com/sgaunet/gitlab-vars)
[![GitHub release](https://img.shields.io/github/release/sgaunet/gitlab-vars.svg)](https://github.com/sgaunet/gitlab-vars/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/sgaunet/gitlab-vars)](https://goreportcard.com/report/github.com/sgaunet/gitlab-vars)
![GitHub Downloads](https://img.shields.io/github/downloads/sgaunet/gitlab-vars/total)
[![Maintainability](https://api.codeclimate.com/v1/badges/061be3219efb765b5461/maintainability)](https://codeclimate.com/github/sgaunet/gitlab-vars/maintainability)
[![GoDoc](https://godoc.org/github.com/sgaunet/gitlab-vars?status.svg)](https://godoc.org/github.com/sgaunet/gitlab-vars)
[![License](https://img.shields.io/github/license/sgaunet/gitlab-vars.svg)](LICENSE)


# gitlab-vars

Tool to get the environment variables of a gitlab project or gitlab group.

* This tool does not initiliaze the [predefinid variables.](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html)
* It does not care of variables declared in .gitlab-ci.yml file
* But you can simulate some vars by creating a .gitlab-vars.json file (see below)

```
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
```

# Install 

Copy the binary to /usr/local/bin for example. (or another directory which is in your PATH).

## Configuration

2 environment variables can be set :

* GITLAB_TOKEN: used to access to the gitlab API
* GITLAB_URI: to specify another instance of Gitlab (if not set, GITLAB_URI is set to https://gitlab.com)


# Examples

## print variables of the current gitlab project (where the project has been cloned)

```
$ git clone git@gitlab.com:group/awesome-project.git
$ cd awesome-project
$ gitlab-vars
```

## get variables of a gitlab group

```
gitlab-vars -g XXXXXX
```

## get variables of a gitlab project

```
gitlab-vars -p XXXXXX
```

## get variables of a gitlab project

```
cd .../gitlab-cloned-project
gitlab-vars
```

## Filter for a specific env

```
cd .../gitlab-cloned-project
gitlab-vars -e production
```

## Add some vars with .gitlab-vars.json

```
cat > .gitlab-vars.json <<EOF
{
  "variables": [
    {
      "variable_type": "string",
      "key": "variable_key",
      "value": "value",
      "protected": false,
      "masked": false,
      "raw": false,
      "environment_scope": ""
    },
    {
      "variable_type": "string",
      "key": "variable_key2",
      "value": "value2",
      "protected": false,
      "masked": false,
      "raw": false,
      "environment_scope": "prod"
    }
  ]
}
EOF

$ gitlab-vars
...
```


# Possibility to use it with envtemplate

[envtemplate](https://github.com/sgaunet/envtemplate) is a project to generate files from templates according to environment variables.

Example:

```
$ cat tmpl 
{{ .REGISTRY }}
$ cat t.sh 
#!/usr/bin/env bash
eval "$(gitlab-vars -p XXXXX)"
envtemplate -i tmpl > tmpl.out
$ ls
tmpl  t.sh
$ bash t.sh 
$ cat tmpl.out 
registry.gitlab.com
```
