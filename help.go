package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type arguments struct {
	urlMode   bool
	query     string
	resources []string
}

func (args *arguments) parse() {
	flag.Usage = func() {
		h := []string{
			fmt.Sprintf("HTMLquery %v", version),
			"A fairly simple html selector built around the goquery golang package.",
			"",
			"Usage:",
			fmt.Sprintf("   %v [OPTIONS] <query> [<resources>]", filepath.Base(os.Args[0])),
			"",
			"Options:",
			"   -u, --url                   enable URL mode",
			"   -h, --help                  show help",
			"",
		}
		fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
		os.Exit(1)
	}

	flag.BoolVar(&args.urlMode, "url", args.urlMode, "")
	flag.BoolVar(&args.urlMode, "u", args.urlMode, "")

	flag.Parse()
}

var args arguments

func init() {
	args.parse()
	args.query = flag.Arg(0)
	inputs := flag.Args()
	if len(inputs) > 1 {
		args.resources = flag.Args()[1:]
	}

	// basic validations
	if len(args.query) == 0 {
		println("query argument must not be empty")
		os.Exit(1)
	}
	log.SetFlags(0)
}
