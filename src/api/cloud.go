package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Application struct
type Application struct {
	Id   string `json:"id"`
	Slug string `json:"slug"`
}

// Environment struct
type Environment struct {
	Id   string `json:"id"`
	Slug string `json:"slug"`
}

type Dep struct {
	Activity struct {
		Id string `json:"id"`
	} `json:"activity"`
}

func baseURL() string {
	return "https://api.v2.vapor.cloud"
}

func GetAppEnv(appSlug string, envSlug string, token string) (*Application, *Environment, *http.Response) {
	app, errApp := GetApplication(appSlug, token)
	if errApp != nil || app == nil {
		return nil, nil, errApp
	}

	env, errEnv := GetEnvironment(app.Id, envSlug, token)
	if errEnv != nil || env == nil {
		return nil, nil, errEnv
	}

	return app, env, nil
}

// GetApplication get specific application from api-apps
func GetApplication(slug string, token string) (*Application, *http.Response) {
	url := baseURL() + "/v2/apps/applications?slug=" + slug + "&exact=true"

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	res, getErr := client.Do(req)
	if getErr != nil || res.StatusCode != http.StatusOK {
		return nil, res
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	application := []Application{}
	jsonErr := json.Unmarshal([]byte(body), &application)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if len(application) < 1 {
		fmt.Println("Application not found")
		os.Exit(1)
	}

	app := application[0]

	return &app, nil
}

// GetEnvironment get specific environment from api-apps
func GetEnvironment(appID string, slug string, token string) (*Environment, *http.Response) {
	url := baseURL() + "/v2/apps/applications/" + appID + "/environments"

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	res, getErr := client.Do(req)
	if getErr != nil || res.StatusCode != http.StatusOK {
		return nil, res
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	environment := []Environment{}
	jsonErr := json.Unmarshal(body, &environment)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	for _, element := range environment {
		if element.Slug == slug {
			return &element, nil
		}
	}

	return nil, nil
}

// Deploy trigger a deploy
func Deploy(envID string, branch string, token string) (*Dep, *http.Response) {
	url := baseURL() + "/v2/apps/environments/" + envID + "/deploy"

	client := http.Client{
		Timeout: time.Second * 2,
	}

	var jsonStr = []byte(`{}`)

	if branch != "" {
		jsonStr = []byte("{\"branch\": \"" + branch + "\"}")
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	res, getErr := client.Do(req)
	if getErr != nil || res.StatusCode != http.StatusOK {
		return nil, res
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	activity := &Dep{}
	jsonErr := json.Unmarshal(body, &activity)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return activity, nil
}
