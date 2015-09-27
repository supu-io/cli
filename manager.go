package main

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/fatih/color"
)

// Manager is a type to encapsulate issue actions
type Manager struct {
	URL string
}

func (m *Manager) list(status string) {
	resp, err := http.Get(m.URL + "/issues?status=" + status)
	if err != nil {
		color.Red("Couldn't connect to the server")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	printIssuesList(body)
}

func (m *Manager) details(issue string) {
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

func (m *Manager) move(issue string, spin string) {
	color.Red("Moving " + spin + " issue " + issue)
}

func (m *Manager) comment(issue string, text string) {
	resp, err := http.PostForm(m.URL+"/issues/"+issue+"/comment",
		url.Values{"body": {text}})
	if err != nil {
		color.Red("Couldn't connect to the server")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	color.Green(string(body))
}
