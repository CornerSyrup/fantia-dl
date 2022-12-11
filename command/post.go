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

type PostParams struct {
	postId int
	global GlobalParams
}

func NewPostCommand(fs *flag.FlagSet) *PostParams {
	cmd := new(PostParams)
	addGlobalFlags(fs, &cmd.global)

	fs.IntVar(&cmd.postId, "post", 0, "Post ID")
	fs.Func("url", "Post url", func(s string) error {
		raw, err := url.Parse(s)
		if err != nil {
			return fmt.Errorf("invalid url")
		}

		path := strings.Split(raw.Path, "/")
		if path[1] != "posts" {
			return fmt.Errorf("invalid post url")
		}

		n, err := strconv.Atoi(path[2])
		if err != nil {
			return fmt.Errorf("invalid post url")
		}

		cmd.postId = n

		return nil
	})

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
	wg := new(sync.WaitGroup)

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

			_, path, _ := core.DownloadContent(wg, agent, root, core.BaseUrl+content.DownloadURI, content.Title, p.global.overwrite)
			os.Chtimes(path, time.Now(), content.ParentPost.Date)
		case "photo_gallery":
			for _, photo := range content.PostContentPhotos {
				if p.global.dryRun {
					fmt.Fprintf(os.Stdout, "Will download photo %d from %s\n", photo.ID, api.Title)
					continue
				}

				_, path, _ := core.DownloadContent(wg, agent, root, photo.URL.Original, strconv.Itoa(photo.ID), p.global.overwrite)
				os.Chtimes(path, time.Now(), content.ParentPost.Date)
			}
		}
	}

	wg.Wait()

	return nil
}
