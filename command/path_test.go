package command

import (
	"testing"

	. "github.com/masci/flick-rsync/testutils"
)

func TestParseFilckrPath(t *testing.T) {
	var check = func(path string) bool {
		_, _, err := ParseFilckrPath(path)
		return err == nil
	}

	Expect(t, check("/path/to/somewhere"), false)
	Expect(t, check("masci@flickr:/123456"), true)
	Expect(t, check("masci@flickr:/foo"), false)
	Expect(t, check("masci@flickr:/"), true)
	Expect(t, check("masci@flickr:"), true)
	Expect(t, check("masci@flickr"), false)

	user, set, _ := ParseFilckrPath("masci@flickr:/123456/")
	Expect(t, user, "masci")
	Expect(t, set, "/123456/")

	user, set, _ = ParseFilckrPath("masci@flickr:")
	Expect(t, user, "masci")
	Expect(t, set, "")

	user, set, _ = ParseFilckrPath("masci@flickr:/")
	Expect(t, user, "masci")
	Expect(t, set, "/")
}

func TestIsSetId(t *testing.T) {
	Expect(t, isSetId("/123/"), true)
	Expect(t, isSetId("/123"), true)
	Expect(t, isSetId("123"), true)
	Expect(t, isSetId("/foo"), false)
}
