package command

import (
	"fmt"
	"os"

	"github.com/masci/flick-rsync/sync"
	"github.com/masci/flickr"
	"gopkg.in/alecthomas/kingpin.v2"
)

func Main() {
	var (
		// kingping stuff
		src        = kingpin.Arg("SRC", "Flickr SET to sync").Required().String()
		dest       = kingpin.Arg("DEST", "Destination path").String()
		api_key    = kingpin.Flag("api_key", "Flickr API key").String()
		api_secret = kingpin.Flag("api_secret", "Flickr API secret").String()

		// others
		accessTok *flickr.OAuthToken
	)

	kingpin.CommandLine.Help = "A Flickr syncing tool"
	kingpin.Version("0.0.1")

	// read configuration in priority order, first the config file
	config, err := loadConfigFile(getConfigFilePath())
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

	// get flickr client
	client := flickr.NewFlickrClient(*api_key, *api_secret)

	if config.OAuthToken == "" || config.OAuthTokenSecret == "" {
		// get request token
		tok, err := flickr.GetRequestToken(client)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		// get access token
		url, _ := flickr.GetAuthorizeUrl(client, tok)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}

		// tell user to authorize the app
		var oauthVerifier string
		fmt.Println("Open your browser at this url:", url)
		fmt.Print("Then, insert the code:")
		fmt.Scanln(&oauthVerifier)

		// get the access token
		accessTok, err = flickr.GetAccessToken(client, tok, oauthVerifier)
		client.OAuthToken = accessTok.OAuthToken
		client.OAuthTokenSecret = accessTok.OAuthTokenSecret
	} else {
		client.OAuthToken = config.OAuthToken
		client.OAuthTokenSecret = config.OAuthTokenSecret
	}

	// give up if SRC is not a valid Flickr path
	_, set, err := ParseFilckrPath(*src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sync.SyncRemote(client, set, *dest)
}
