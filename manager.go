package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/fatih/color"
)

// Manager is a type to encapsulate issue actions
type Manager struct {
	URL string
}

func (m *Manager) list(status string, org string, repo string) {
	url := m.URL + "/issues/search?status=" + status + "&org=" + org
	if repo != "" {
		url = url + "&repo=" + repo
	}
	resp, err := http.Get(url)
	if err != nil {
		color.Red("Couldn't connect to the server" + m.URL)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := printIssuesList(body); err != nil {
		m.printError(body)
	}
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

func (m *Manager) move(issue string, status string) {
	if status == "" {
		m.printListOfPossibleStatuses()
	} else {
		var jsonStr = []byte(`{"status":"` + status + `"}`)
		req, err := http.NewRequest("PUT", m.URL+"/issues/"+issue, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("X-AUTH-TOKEN", "token")

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			color.Red("Couldn't connect to the server")
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		color.Green(string(body))
	}
}

func (m *Manager) setup(org string, repo string) {
	url := m.URL + "/setup?org=" + org + "&repo=" + repo

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("X-AUTH-TOKEN", "token")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		color.Red("Couldn't connect to the server")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	println(string(body))
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

func (m *Manager) printError(body []byte) {
	e := Error{}
	err := json.Unmarshal(body, &e)
	if err != nil {
		color.Red(err.Error())
	}

	color.Red(e.Message)
}

func (m *Manager) printListOfPossibleStatuses() {
	req, err := http.NewRequest("GET", m.URL+"/statuses", nil)
	req.Header.Add("X-AUTH-TOKEN", "token")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		color.Red("Couldn't connect to the server")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	statuses := []string{}
	json.Unmarshal(body, &statuses)
	color.Yellow("You should provide a valid status to move your issue")
	color.Yellow("valid statuses for this project are:")
	for _, val := range statuses {
		color.Yellow(" - " + val)
	}
}
