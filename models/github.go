package models

import (
	"encoding/json"
	"time"
)

type Repos struct {
	Data struct {
		User struct {
			Pinned struct {
				Nodes []node `json:"nodes"`
			} `json:"pinnedItems"`
		} `json:"user"`
	} `json:"data"`
}

type Commit struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Commits []struct {
			Message string `json:"message"`
		}
	}
	CreatedAt time.Time `json:"created_at"`
}

type node struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Languages   struct {
		Node []struct {
			Name string `json:"name"`
		} `json:"nodes"`
	} `json:"languages"`
	Forks int    `json:"forkCount"`
	Stars int    `json:"stargazerCount"`
	URL   string `json:"url"`
}

func (data *Repos) MarshalJSON() ([]byte, error) {
	res := map[string]any{
		"projects": data.Data.User.Pinned.Nodes,
	}
	return json.Marshal(res)
}

func (data *node) MarshalJSON() ([]byte, error) {
	var langs []string
	for _, lang := range data.Languages.Node {
		langs = append(langs, lang.Name)
	}
	res := map[string]any{
		"name":        data.Name,
		"description": data.Description,
		"languages":   langs,
		"forks":       data.Forks,
		"stars":       data.Stars,
		"repository":  data.URL,
	}
	return json.Marshal(res)
}

func (data *Commit) MarshalJSON() ([]byte, error) {
	res := map[string]any{
		"type":           data.Type,
		"repository":     "https://github.com/" + data.Repo.Name,
		"commit_message": data.Payload.Commits,
		"timestamp":      data.CreatedAt,
	}
	return json.Marshal(res)
}
