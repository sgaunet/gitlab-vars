
# gitlab-vars

Tool to simulate the initialization of environment variables during gitlab pipeline.
**The tool is in beta actually, the command line can change, the options too...**

* Actually, this tool does not initiliaze the [predefinid variables.](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html) (TODO: initialize them with a default value or add the possibility to let the user initialize it)
* It does not care of variables declared in .gitlab-ci.yml file too (TOOD)

# Install 

Copy the binary to /usr/local/bin for example. (or another directory which is in your PATH).

# Usage

```
gitlab-vars:
  -d string
        Debug level (info,warn,debug) (default "error")
  -e string
        environment to filter variables (default "*")
  -g int
        Group ID to get variables from (not compatible with -p option)
  -p int
        Project ID to get variables from
  -v    Get version
```

## Configuration

2 environement variables can be set :

* GITLAB_TOKEN: used to access to the gitlab API
* GITLAB_URI: to specify another instance of Gitlab (if not set, GITLAB_URI is set to https://gitlab.com)


# Examples

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

# Possibility to use it with envtemplate

[envtemplate](https://github.com/sgaunet/envtemplate) is a project to generate files from templates according to environment variables.

Example:

```
$ cat tmpl 
{{ .REGISTRY }}
$ cat t.sh 
#!/usr/bin/env bash
eval "$(go run .. -p XXXXX)"
envtemplate -i tmpl > tmpl.out
$ ls
tmpl  t.sh
$ bash t.sh 
$ cat tmpl.out 
registry.gitlab.com
```
