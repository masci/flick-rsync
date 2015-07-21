package sync

import (
	"fmt"
	"path/filepath"

	"github.com/masci/flickr.go/flickr"
)

func SyncRemote(client *flickr.FlickrClient, setId string, destPath string) {

	// default destination path if not provided
	if destPath == "" {
		destPath = "."
	}

	// normalize destination path
	if !filepath.IsAbs(destPath) {
		destPath, _ = filepath.Abs(destPath)
	}

	// if setId is empty, assume "all sets" need to be synced
	if setId == "" {
		// retrieve all sets
		fmt.Println("Syncing al flickr sets with local folder", destPath)
	} else {
		fmt.Println("Sync flickr set", setId, "with local folder", destPath)
	}
}
