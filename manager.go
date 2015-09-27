package main

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

// Manager is a type to encapsulate issue actions
type Manager struct {
	Context *cli.Context
	URL     string
}

// Manage is a method to execute actions on a single context
func (m *Manager) Manage() {
	if len(m.Context.Args()) == 0 {
		color.Yellow("Not enough arguments")
	}
	switch m.Context.Args()[0] {
	case "next":
		m.getList("todo")
		return
	case "doing":
		m.getList("doing")
		return
	case "details":
		issue := m.Context.Args()[1]
		m.getDetails(issue)
		return
	case "comment":
		m.postComment()
		return
	case "start":
		m.moveTo("doing")
		return
	case "review":
		m.moveTo("review")
		return
	case "uat":
		m.moveTo("uat")
		return
	case "done":
		m.moveTo("done")
		return
	}
	color.Red("Unrecognized option")
}

func (m *Manager) getList(status string) {
	resp, err := http.Get(m.URL + "/issues?status=" + status)
	if err != nil {
		color.Red("Couldn't connect to the server")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	printIssuesList(body)
}

func (m *Manager) getDetails(issue string) {
	req, err := http.NewRequest("GET", m.URL+"/issues/"+issue, nil)
	req.Header.Add("X-AUTH-TOKEN", "token")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		color.Red("Couldn't connect to the server")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	printIssueDetails(body)
}

func (m *Manager) postComment() string {
	resp, err := http.PostForm(m.URL+"/issues/1/comment",
		url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		return "Couldn't connect to the server" + m.URL
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func (m *Manager) moveTo(status string) string {
	// req, err := http.NewRequest("PUT", m.Url+"/issues/"+issue, nil)
	// req.Header.Add("X-AUTH-TOKEN", "token")

	return ""
}
