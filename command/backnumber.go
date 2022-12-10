package command

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
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
	agent := core.NewAgent(p.global.session)
	api, err := core.FetchBacknumber(agent, p.plan, p.year, p.month)
	if err != nil {
		return err
	}

	postRoot := api.JoinBasePath(p.global.dir)
	os.MkdirAll(postRoot, fs.ModeDir)

	for _, post := range api.Backnumber.BacknumberContents {
		contRoot := filepath.Join(postRoot, fmt.Sprintf("%d_%s", post.Plan.Price, post.Title))
		os.MkdirAll(contRoot, fs.ModeDir)

		switch post.Category {
		case "file":
			if post.DownloadURI == "" {
				continue
			}
			core.DownloadContent(agent, contRoot, core.BaseUrl+post.DownloadURI, post.Title)
		case "photo_gallery":
			for _, photo := range post.PostContentPhotos {
				core.DownloadContent(agent, contRoot, photo.URL.Original, strconv.Itoa(photo.ID))
			}
		}
	}

	return nil
}
