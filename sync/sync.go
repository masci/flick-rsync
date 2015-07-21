package sync

import (
	"fmt"

	"github.com/masci/flickr.go/flickr"
)

func SyncRemote(client *flickr.FlickrClient, setId string, destPath string) {

	// default destination path if not provided
	if destPath == "" {
		destPath = "."
	}

	fmt.Println("Sync flickr set", setId, "with local folder", destPath)

}
