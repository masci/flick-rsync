package command

import (
	"fmt"
	"github.com/masci/flick-rsync/flickr"
	"github.com/masci/flick-rsync/flickr/test"
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
	fmt.Println("Would sync", *src, *dest)
	fmt.Println("Apikey:", *api_key, "Apisec:", *api_secret)

	// get flickr client
	client := flickr.NewFlickrClient(*api_key, *api_secret)

	var accessTok *flickr.OAuthToken

	if config.OAuthToken == "" {
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
		fmt.Println(accessTok)
		fmt.Println(err)
	} else {
		accessTok = &flickr.OAuthToken{
			OAuthToken:       config.OAuthToken,
			OAuthTokenSecret: config.OAuthTokenSecret,
		}
	}

	// test
	resp, err := test.Login(client, accessTok)
	fmt.Println(resp.Status, resp.User)
	fmt.Println(err)
}
