package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Devansh3712/portfolio/models"
)

func GetRepos() (*models.Repos, error) {
	api := "https://api.github.com/graphql"
	graphql := map[string]string{
		"query": fmt.Sprintf(`
			{
				user(login: "%s") {
					pinnedItems(first: 6, types: REPOSITORY) {
						nodes {
							... on Repository {
								name
								description
								languages(first: 10) {
									nodes {
										name
									}
								}
								forkCount
								stargazerCount
								url
							}
						}
					}
				}
			}
		`, os.Getenv("GITHUB_USERNAME")),
	}
	postJson, err := json.Marshal(graphql)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", api, bytes.NewBuffer(postJson))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "bearer "+os.Getenv("GITHUB_PAT"))
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var userData models.Repos
	if err = json.Unmarshal(data, &userData); err != nil {
		return nil, err
	}
	return &userData, nil
}

func GetLatestCommit() (*models.Commit, error) {
	api := "https://api.github.com/users/" + os.Getenv("GITHUB_USERNAME") + "/events/public"

	client := &http.Client{}
	request, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var commits []models.Commit
	if err = json.Unmarshal(data, &commits); err != nil {
		return nil, err
	}
	return &commits[0], nil
}
