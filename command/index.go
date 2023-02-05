package command

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Command interface {
	Sanitize() error
	Execute() error
}

type GlobalParams struct {
	session   string
	token     string
	dir       string
	dryRun    bool
	overwrite bool
}

func addGlobalFlags(fs *flag.FlagSet, gp *GlobalParams) {
	fs.StringVar(&gp.session, "session", "", "Session ID which has been logged in to Fantia")
	fs.StringVar(&gp.token, "token", "", "CSRF token that use for interacting with Fantia")
	fs.StringVar(&gp.dir, "dir", "", "Directory to store the downloaded contents")
	fs.BoolVar(&gp.dryRun, "dry-run", false, "Only list out all content to be downloaded, will not download anything")
	fs.BoolVar(&gp.overwrite, "overwrite", false, "Overwrite existing content")
}

func sanitizeGlobalParam(g GlobalParams) error {
	if len(g.session) < 1 {
		return fmt.Errorf("session ID cannot be empty")
	}

	if path, err := filepath.Abs(g.dir); err != nil {
		return fmt.Errorf("invalid path %s", path)
	} else if stat, err := os.Stat(path); err != nil {
		return err
	} else if !stat.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	} else {
		g.dir = path
	}

	return nil
}
