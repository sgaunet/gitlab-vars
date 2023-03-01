
# gitlab-vars

Tool to simulate the initialization of environment variables during gitlab pipeline.
**The tool is in beta actually, the command line can change, the options too...**

* Actually, this tool does not initiliaze the [predefinid variables.](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html) (TODO: initialize them with a default value or add the possibility to let the user initialize it)
* It does not care of variables declared in .gitlab-ci.yml file too (TOOD)

# Install 

Copy the binary to /usr/local/bin for example. (or another directory which is in your PATH).

# Usage (TODO)

```
...
```

## Configuration

2 environement variables can be set :

* GITLAB_TOKEN: used to access to private repositories
* GITLAB_URI: to specify another instance of Gitlab (if not set, GITLAB_URI is set to https://gitlab.com)


# Example (TODO)

```
eval "$(go run . -g 6925554)"
```
