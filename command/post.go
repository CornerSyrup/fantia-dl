package command

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strconv"

	"github.com/KleinChiu/fantia-dl/core"
)

type PostParams struct {
	postId int
	global GlobalParams
}

func NewPostCommand(fs *flag.FlagSet) *PostParams {
	cmd := new(PostParams)
	addGlobalFlags(fs, &cmd.global)

	fs.IntVar(&cmd.postId, "post", 0, "Post ID")

	return cmd
}

func (p PostParams) Sanitize() error {
	if err := sanitizeGlobalParam(p.global); err != nil {
		return err
	}

	if p.postId == 0 {
		return fmt.Errorf("invalid post id %d", p.postId)
	}

	return nil
}

func (p PostParams) Execute() error {
	agent := core.NewAgent(p.global.session)
	api, err := core.FetchPost(agent, p.postId)
	if err != nil {
		return err
	}

	postRoot := api.JoinBasePath(p.global.dir)
	os.MkdirAll(postRoot, fs.ModeDir)

	for _, content := range api.PostContents {
		root := content.JoinBasePath(postRoot)
		os.Mkdir(root, fs.ModeDir)

		switch content.Category {
		case "file":
			if content.DownloadURI == "" {
				continue
			}

			if p.global.dryRun {
				fmt.Fprintf(os.Stdout, "Will download file %s from %s\n", content.Filename, api.Title)
				continue
			}

			core.DownloadContent(agent, root, core.BaseUrl+content.DownloadURI, content.Title, p.global.overwrite)
		case "photo_gallery":
			for _, photo := range content.PostContentPhotos {
				if p.global.dryRun {
					fmt.Fprintf(os.Stdout, "Will download photo %d from %s\n", photo.ID, api.Title)
					continue
				}

				core.DownloadContent(agent, root, photo.URL.Original, strconv.Itoa(photo.ID), p.global.overwrite)
			}
		}
	}

	return nil
}
