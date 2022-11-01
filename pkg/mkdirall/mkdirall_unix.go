// copied from src/os/path_unix.go

// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris

package mkdirall

func fixRootDirectory(p string) string {
	return p
}
