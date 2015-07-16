package command

import (
	"testing"

	. "github.com/masci/flick-rsync/testutils"
)

func TestParseFilckrPath(t *testing.T) {
	var IsFilckrPath = func(path string) bool {
		_, _, err := ParseFilckrPath(path)
		return err == nil
	}
	Expect(t, IsFilckrPath("/path/to/somewhere"), false)
	Expect(t, IsFilckrPath("https://www.flickr.com/photos/masci/sets/123456"), true)
	Expect(t, IsFilckrPath("https://www.flickr.com/photos/masci/sets"), true)
	Expect(t, IsFilckrPath("https://www.flickr.com/photos/masci"), true)
}
