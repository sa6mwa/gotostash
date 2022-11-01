package mkdirall

import (
	"os"
	"os/user"
	"testing"
)

func TestMkdirAllAndChown(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	directories := []string{"/tmp/mkdirall/testing/directory", "/tmp/mkdirall/another/directory"}
	for _, d := range directories {
		err := MkdirAllAndChown(d, 0755, usr.Uid, usr.Gid)
		if err != nil {
			t.Fatal(err)
		}
	}
	directoriesToRemove := []string{"/tmp/mkdirall"}
	for _, d := range directoriesToRemove {
		err := os.RemoveAll(d)
		if err != nil {
			t.Fatal(err)
		}
	}
}
