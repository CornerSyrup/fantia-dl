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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-CSRF-Token", "COiUCik3fbywpa8YBjAkC3tKwYo3E7qf8lOO0OIrBUHhNGugft-3jqx2zryoEV6XhoSgZn3oglpmgG1NRgyN-g")

	res, err := agent.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to query Fantia API")
	}
	defer res.Body.Close()

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
