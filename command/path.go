package command

import (
	"errors"
	"regexp"
)

var validPath *regexp.Regexp

func init() {
	validPath = regexp.MustCompile(`https://www.flickr.com/photos/\w+(/sets/\d+)?`)
}

func ParseFilckrPath(path string) (string, string, error) {
	match := validPath.FindStringSubmatch(path)
	if len(match) == 2 {
		return match[0], match[1], nil
	} else if len(match) == 1 {
		return match[0], "", nil
	}

	return "", "", errors.New(path, ": not a valid Flickr path.")
}
