package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KleinChiu/fantia-dl/command"
)

func main() {
	var cmd command.Command
	var fs *flag.FlagSet

	switch os.Args[1] {
	case "post":
		fs = flag.NewFlagSet("post", flag.ExitOnError)
		cmd = command.NewPostCommand(fs)
	default:
		fmt.Fprintf(os.Stderr, "Usage:\n\n\tfantia-dl post [arguments]\n")
		os.Exit(0)
	}

	fs.Parse(os.Args[2:])

	if err := cmd.Sanitize(); err != nil {
		fs.Usage()
	}
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stdout, err.Error()+"\n")
	}
}
