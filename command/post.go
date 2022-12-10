package command

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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

	postRoot := filepath.Join(p.global.dir, fmt.Sprintf("%d_%s", api.Post.Fanclub.ID, api.Post.Fanclub.FanclubName), fmt.Sprintf("%d_%s", api.Post.ID, api.Post.Title))
	os.MkdirAll(postRoot, fs.ModeDir)

	for _, content := range api.Post.PostContents {
		root := filepath.Join(postRoot, fmt.Sprintf("%d_%s", content.Plan.Price, content.Title))
		os.Mkdir(root, fs.ModeDir)

		switch content.Category {
		case "file":
			if content.DownloadURI == "" {
				continue
			}
			core.DownloadContent(agent, root, core.BaseUrl+content.DownloadURI, content.Title)
		case "photo_gellery":
			for _, photo := range content.PostContentPhotos {
				core.DownloadContent(agent, root, photo.URL.Original, strconv.Itoa(photo.ID))
			}
		}
	}

	return nil
}
