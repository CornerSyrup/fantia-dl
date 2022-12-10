package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseUrl = "https://fantia.jp"
)

// Endpoints
const (
	backnumberEndpoint = BaseUrl + "/api/v1/fanclub/backnumbers/monthly_contents/plan/%d/month/%d%02d"
	postEndpoint       = BaseUrl + "/api/v1/posts/%d"
)

func FetchBacknumber(agent *http.Client, plan int, year int, month int) (*Backnumber, error) {
	buf, err := fetchApi(agent, fmt.Sprintf(backnumberEndpoint, plan, year, month))
	if err != nil {
		return nil, err
	}

	api := new(BacknumberApi)
	err = json.Unmarshal(buf, api)
	if err != nil {
		return nil, err
	}

	return &api.Backnumber, err
}

func FetchPost(agent *http.Client, id int) (*Post, error) {
	buf, err := fetchApi(agent, fmt.Sprintf(postEndpoint, id))
	if err != nil {
		return nil, err
	}

	api := new(PostApi)
	err = json.Unmarshal(buf, api)
	if err != nil {
		return nil, err
	}

	return &api.Post, nil
}

func fetchApi(agent *http.Client, url string) ([]byte, error) {
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
