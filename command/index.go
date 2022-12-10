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
	session string
	dir     string
}

func addGlobalFlags(fs *flag.FlagSet, gp *GlobalParams) {
	fs.StringVar(&gp.session, "session", "", "Session ID which has been logged in to Fantia")
	fs.StringVar(&gp.dir, "dir", "", "Directory to store the downloaded contents")
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
