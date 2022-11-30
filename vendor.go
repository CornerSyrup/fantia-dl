package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseUrl = "https://fantia.jp"
)

// Endpoints
const (
	backnumberEndpoint = baseUrl + "/api/v1/fanclub/backnumbers/monthly_contents/plan/%d/month/%d%02d"
	postEndpoint       = baseUrl + "/api/v1/posts/%d"
)

func FetchBacknumber(agent *http.Client, plan int, year int, month int) (*BackNumberApi, error) {
	buf, err := FetchApi(agent, fmt.Sprintf(backnumberEndpoint, plan, year, month))
	if err != nil {
		return nil, err
	}

	api := new(BackNumberApi)
	err = json.Unmarshal(buf, api)
	if err != nil {
		return nil, err
	}

	return api, err
}

func FetchPost(agent *http.Client, id int) (*PostApi, error) {
	buf, err := FetchApi(agent, fmt.Sprintf(postEndpoint, id))
	if err != nil {
		return nil, err
	}

	api := new(PostApi)
	err = json.Unmarshal(buf, api)
	if err != nil {
		return nil, err
	}

	return api, nil
}

func FetchApi(agent *http.Client, url string) ([]byte, error) {
	res, err := agent.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
