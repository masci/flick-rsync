package command

import (
	"fmt"
	"github.com/masci/flick-rsync/flickr"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

func Main() {
	var (
		src        = kingpin.Arg("SRC", "Flickr SET to sync").Required().String()
		dest       = kingpin.Arg("DEST", "Destination path").String()
		api_key    = kingpin.Flag("api_key", "Flickr API key").String()
		api_secret = kingpin.Flag("api_secret", "Flickr API secret").String()
	)

	kingpin.CommandLine.Help = "A Flickr syncing tool"
	kingpin.Version("0.0.1")

	// read configuration in priority order, first the config file
	config, err := parseConfigFile(getConfigFilePath())
	if err == nil {
		*api_key = config.ApiKey
		*api_secret = config.ApiSecret
	}

	// then the environment vars
	if apik := os.Getenv("FLICKRSYNC_API_KEY"); apik != "" {
		*api_key = apik
	}

	if apisec := os.Getenv("FLICKRSYNC_API_SECRET"); apisec != "" {
		*api_secret = apisec
	}

	// then the command line
	kingpin.Parse()

	// give up if api keys were not provided
	if *api_key == "" || *api_secret == "" {
		fmt.Println("Flickr API keys not found, exiting...")
		os.Exit(1)
	}
	fmt.Println("Would sync", *src, *dest)
	fmt.Println("Apikey:", *api_key, "Apisec:", *api_secret)

	// get flickr client
	client := flickr.NewFlickrClient(*api_key, *api_secret)

	tok, err := flickr.GetRequestToken(client)
	fmt.Println(tok)
	url, _ := flickr.GetAuthorizeUrl(client, tok)
	fmt.Println(url)
}
