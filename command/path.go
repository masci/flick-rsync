package command

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var re *regexp.Regexp

func init() {
	// compile the regex matching flickr paths
	re = regexp.MustCompile(`^(\w+)@flickr:(\d+)?$`)
}

// Parse component of a flickr path of the form:
//
// `<username>@flickr:<setId>`
//
// setId can be omitted, like in:
//
// `<username>@flickr:`
//
// an empty setId has different meanings, depending on the context
//
func ParseFilckrPath(path string) (string, string, error) {
	var isSetId = func(set string) bool {
		_, err := strconv.Atoi(set)
		return err == nil
	}

	match := re.FindStringSubmatch(path)

	if len(match) == 3 {
		u := match[1]
		s := match[2]
		if s == "" || isSetId(s) {
			return u, s, nil
		}
		return "", "", errors.New(fmt.Sprintf("Not a valid Flickr Set id: %s", s))
	}

	return "", "", errors.New(fmt.Sprintf("Not a valid Flickr path: %s", path))
}
