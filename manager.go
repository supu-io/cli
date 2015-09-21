package main

import (
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Manager struct {
	Context *cli.Context
	Url     string
}

func (m *Manager) Manage() string {
	if len(m.Context.Args()) == 0 {
		return "Not enough arguments"
	}
	switch m.Context.Args()[0] {
	case "next":
		return m.getList("todo")
	case "doing":
		return m.getList("doing")
	case "details":
		issue := m.Context.Args()[1]
		return m.getDetails(issue)
	case "comment":
		return m.postComment()
	case "start":
		return m.moveTo("doing")
	case "review":
		return m.moveTo("review")
	case "uat":
		return m.moveTo("uat")
	case "done":
		return m.moveTo("done")
	}
	return "Unrecognized option"
}

func (m *Manager) getList(status string) string {
	resp, err := http.Get(m.Url + "/issues")
	if err != nil {
		return "Couldn't connect to the server " + m.Url
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func (m *Manager) getDetails(issue string) string {
	req, err := http.NewRequest("GET", m.Url+"/issues/"+issue, nil)
	req.Header.Add("X-AUTH-TOKEN", "token")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "Couldn't connect to the server " + m.Url
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func (m *Manager) postComment() string {
	resp, err := http.PostForm(m.Url+"/issues/1/comment",
		url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		return "Couldn't connect to the server" + m.Url
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func (m *Manager) moveTo(status string) string {
	return ""
}
