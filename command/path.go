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
	// compile the regex matching flickr paths
	re = regexp.MustCompile(`(\w+)@flickr:(/[\w|/]*)?`)
}

// determine if a string contains a valid flickr Set Identifier
func isSetId(set string) bool {
	s := strings.Trim(set, "/")
	_, err := strconv.Atoi(s)
	return err == nil
}

// Parse component of a flickr path of the form:
//
// `<username>@flickr:<setId>`
//
// setId can be omitted, like in:
//
// `<username>@flickr:<setId>`
//
// or set to `/`:
//
// `<username>@flickr:/`
//
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
