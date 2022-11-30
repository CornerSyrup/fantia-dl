package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
)

func NewAgent(session string) *http.Client {
	url, _ := url.Parse(baseUrl)
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

func DownloadContent(agent *http.Client, dir string, url string, filename string) (int64, error) {
	res, err := agent.Get(url)
	if err != nil {
		return 0, err
	} else if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code of %d", res.StatusCode)
	}

	fp := filepath.Join(dir, filename+filepath.Ext(res.Request.URL.Path))
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