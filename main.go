package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"

	"github.com/KleinChiu/fantia-dl/core"
)

func main() {
	session := flag.String("session", "", "Session ID which has been logged in to Fantia")
	dir := flag.String("dir", "", "Directory to store the downloaded contents")
	postId := flag.Int("post", 0, "Post ID")

	flag.Parse()

	if len(*session) < 1 {
		panic("session ID cannot be empty")
	}

	if len(*dir) == 0 {
		*dir = filepath.Dir("")
	}
	if path, err := filepath.Abs(*dir); err != nil {
		panic(err.Error())
	} else {
		*dir = path
	}
	if stat, err := os.Stat(*dir); err != nil {
		panic(err.Error())
	} else if !stat.IsDir() {
		panic(fmt.Sprintf("dir %s is not a directory", *dir))
	}

	if *postId == 0 {
		panic("post id is empty")
	}

	agent := core.NewAgent(*session)
	api, err := core.FetchPost(agent, *postId)
	if err != nil {
		panic(err)
	}

	postRoot := filepath.Join(*dir, fmt.Sprintf("%d_%s", api.Post.Fanclub.ID, api.Post.Fanclub.FanclubName), fmt.Sprintf("%d_%s", api.Post.ID, api.Post.Title))
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
				go core.DownloadContent(agent, root, photo.URL.Original, strconv.Itoa(photo.ID))
			}
		}
	}
}
