package command

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
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

			core.DownloadContent(agent, root, core.BaseUrl+content.DownloadURI, content.Title, p.global.overwrite)
		case "photo_gallery":
			for _, photo := range content.PostContentPhotos {
				if p.global.dryRun {
					fmt.Fprintf(os.Stdout, "Will download photo %d from %s\n", photo.ID, src)
					continue
				}

				core.DownloadContent(agent, root, photo.URL.Original, strconv.Itoa(photo.ID), p.global.overwrite)
			}
		}
	}

	return nil
}
