package command

import (
	"flag"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/KleinChiu/fantia-dl/core"
)

type BacknumberParams struct {
	year   int
	month  int
	plan   int
	global GlobalParams
}

func NewBacknumberCommand(fs *flag.FlagSet) *BacknumberParams {
	cmd := new(BacknumberParams)
	addGlobalFlags(fs, &cmd.global)

	fs.IntVar(&cmd.year, "year", 0, "Year of the back number issue")
	fs.IntVar(&cmd.month, "month", 0, "Month of the back number issue")
	fs.IntVar(&cmd.plan, "plan", 0, "Plan of the back number")
	fs.Func("url", "Backnumber URL", func(s string) error {
		raw, err := url.Parse(s)
		if err != nil {
			return fmt.Errorf("invalid url")
		}

		path := strings.Split(raw.Path, "/")
		if path[3] != "backnumbers" {
			return fmt.Errorf("invalid backnumber url")
		}

		m := raw.Query().Get("month")
		if len(m) != 6 {
			return fmt.Errorf("invalid backnumber url")
		}
		cmd.year, _ = strconv.Atoi(m[0:4])
		cmd.month, _ = strconv.Atoi(m[4:])
		cmd.plan, _ = strconv.Atoi(raw.Query().Get("plan"))

		return nil
	})

	return cmd
}

func (p BacknumberParams) Sanitize() error {
	if err := sanitizeGlobalParam(p.global); err != nil {
		return err
	}

	if p.year < 2016 || p.year > time.Now().Year() {
		return fmt.Errorf("invalid year of %d", p.year)
	}
	if p.month < 0 || p.month > 12 {
		return fmt.Errorf("invalid month of %d", p.month)
	}
	if p.plan == 0 {
		return fmt.Errorf("invalid plan of %d", p.plan)
	}

	return nil
}

func (p BacknumberParams) Execute() error {
	wg := new(sync.WaitGroup)

	agent := core.NewAgent(p.global.session)
	api, err := core.FetchBacknumber(agent, p.plan, p.year, p.month)
	if err != nil {
		return err
	}

	bnRoot := api.JoinBasePath(p.global.dir)
	os.MkdirAll(bnRoot, fs.ModeDir)

	for _, content := range api.BacknumberContents {
		root := content.JoinBasePath(bnRoot)
		os.MkdirAll(root, fs.ModeDir)

		src := content.ParentPost.Title
		if len(content.Title) > 0 {
			src += " - " + content.Title
		}

		switch content.Category {
		case "file":
			if content.DownloadURI == "" {
				continue
			}

			if p.global.dryRun {
				fmt.Fprintf(os.Stdout, "Will download file %s from %s\n", content.Filename, src)
				continue
			}

			go func(content core.BacknumberContent) {
				_, path, _ := core.DownloadContent(wg, agent, root, core.BaseUrl+content.DownloadURI, content.Title, p.global.overwrite)
				os.Chtimes(path, time.Now(), content.ParentPost.Date)
			}(content)
		case "photo_gallery":
			for _, photo := range content.PostContentPhotos {
				if p.global.dryRun {
					fmt.Fprintf(os.Stdout, "Will download photo %d from %s\n", photo.ID, src)
					continue
				}

				go func(content core.BacknumberContent, photo core.PostContentPhoto) {
					_, path, _ := core.DownloadContent(wg, agent, root, photo.URL.Original, strconv.Itoa(photo.ID), p.global.overwrite)
					os.Chtimes(path, time.Now(), content.ParentPost.Date)
				}(content, photo)
			}
		}
	}

	wg.Wait()

	return nil
}
