package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug = kingpin.Flag("debug", "Enable debug mode.").Bool()
	src   = kingpin.Arg("SRC", "Flickr SET to sync").Required().String()
	dest  = kingpin.Arg("DEST", "Destination path").String()
)

func main() {
	kingpin.CommandLine.Help = "A Flickr syncing tool"
	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Printf("Would sync: %s to %s", *src, *dest)
}
