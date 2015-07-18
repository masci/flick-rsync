package command

import (
	"errors"
	"fmt"
	"regexp"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`(\w+)@flickr:(/)?(\d+)?`)
}

func ParseFilckrPath(path string) (string, string, error) {
	match := re.FindStringSubmatch(path)

	if len(match) == 4 {
		return match[1], match[3], nil
	} else if len(match) > 1 {
		return match[1], "", nil
	}

	return "", "", errors.New(fmt.Sprintf("Not a valid Flickr path: %s", path))
}
