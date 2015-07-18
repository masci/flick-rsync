package command

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`(\w+)@flickr:(/[\w|/]*)?`)
}

func isSetId(set string) bool {
	s := strings.Trim(set, "/")
	_, err := strconv.Atoi(s)
	return err == nil
}

func ParseFilckrPath(path string) (string, string, error) {
	match := re.FindStringSubmatch(path)

	if len(match) == 3 {
		u := match[1]
		s := match[2]
		if s == "" || s == "/" || isSetId(s) {
			return u, s, nil
		}
		return "", "", errors.New(fmt.Sprintf("Not a valid Flickr Set id: %s", s))
	} else if len(match) > 1 {
		return match[1], "", nil
	}

	return "", "", errors.New(fmt.Sprintf("Not a valid Flickr path: %s", path))
}
