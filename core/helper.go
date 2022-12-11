package core

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
)

func NewAgent(session string) *http.Client {
	url, _ := url.Parse(BaseUrl)
	jar, _ := cookiejar.New(nil)

	jar.SetCookies(url, []*http.Cookie{{
		Name:       "_session_id",
		Value:      session,
		Path:       "/",
		Domain:     url.Host,
		RawExpires: "Session",
		Secure:     false,
		SameSite:   http.SameSiteLaxMode,
	}})

	return &http.Client{
		Jar: jar,
	}
}

func DownloadContent(agent *http.Client, dir string, url string, filename string, overwrite bool) (int64, error) {
	res, err := agent.Get(url)
	if err != nil {
		return 0, err
	} else if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code of %d", res.StatusCode)
	}
	defer res.Body.Close()

	fp := filepath.Join(dir, filename+filepath.Ext(res.Request.URL.Path))
	if _, err := os.Stat(fp); !errors.Is(err, os.ErrNotExist) && !overwrite {
		return 0, nil
	}

	f, err := os.Create(fp)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	if n, err := io.Copy(f, res.Body); err != nil {
		return 0, err
	} else {
		return n, nil
	}
}

func (p Post) JoinBasePath(base string) string {
	return filepath.Join(fmt.Sprintf("%d_%s", p.Fanclub.ID, p.Fanclub.FanclubName), fmt.Sprintf("%d_%s", p.ID, p.Title))
}

func (c PostApiPostContent) JoinBasePath(base string) string {
	return filepath.Join(base, fmt.Sprintf("%d_%s", c.Plan.Price, c.Title))
}

func (p Backnumber) JoinBasePath(base string) string {
	return filepath.Join(base, fmt.Sprintf("%d_%s", p.Fanclub.ID, p.Fanclub.FanclubName))
}

func (p BacknumberContent) JoinBasePath(base string) string {
	var postId int
	fmt.Sscanf(p.ParentPost.URL, "/posts/%d", &postId)

	return filepath.Join(base, fmt.Sprintf("%d_%s", postId, p.ParentPost.Title), fmt.Sprintf("%d_%s", p.Plan.Price, p.Title))
}
