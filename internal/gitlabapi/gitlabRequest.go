package gitlabapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const GitlabApiVersion = "v4"

type GitlabApiClient struct {
	HttpClient *http.Client
}

func NewGitlapApiClient() *GitlabApiClient {
	return &GitlabApiClient{
		HttpClient: &http.Client{},
	}
}

func Request(uri string) (body []byte, err error) {
	url := fmt.Sprintf("%s/api/%s/%s", os.Getenv("GITLAB_URI"), GitlabApiVersion, uri)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("PRIVATE-TOKEN", os.Getenv("GITLAB_TOKEN"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	return
}
