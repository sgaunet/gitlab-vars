package gitlabapi

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type MockClientSuccess struct {
}

func (m MockClientSuccess) Do(req *http.Request) (*http.Response, error) {
	responseBody := ioutil.NopCloser(bytes.NewReader([]byte(`fixed response`)))
	return &http.Response{
		StatusCode: 200,
		Body:       responseBody,
	}, nil
}

type MockClientFailed struct {
}

func (m MockClientFailed) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("failed")
}

func TestRequestSuccess(t *testing.T) {
	g := NewGitlapApiClient()
	g.HttpClient = MockClientSuccess{}

	body, err := g.Request("test")
	if err != nil {
		t.Errorf("Expected err=nil, got %s", err.Error())
	}
	wantedBody := []byte(`fixed response`)
	if !reflect.DeepEqual(body, wantedBody) {
		t.Errorf("Expected value=%v , got %v", wantedBody, body)
	}
}

func TestRequestFailed(t *testing.T) {
	g := NewGitlapApiClient()
	g.HttpClient = MockClientFailed{}

	_, err := g.Request("test")
	if err == nil {
		t.Errorf("Expected err!=nil")
	}
}
