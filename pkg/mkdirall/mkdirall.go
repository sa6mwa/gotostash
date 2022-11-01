package mkdirall

import (
	"os"
	"strconv"
	"syscall"
)

// this is a modified version of os.MkdirAll (add setting uid and gid of created directories)
func MkdirAllAndChown(path string, perm os.FileMode, uid, gid string) error {
	// Resolve uid and gid as integers.
	u, err := strconv.Atoi(uid)
	if err != nil {
		return err
	}
	g, err := strconv.Atoi(gid)
	if err != nil {
		return err
	}
	// Fast path: if we can tell whether path is a directory or file, stop with success or error.
	dir, err := os.Stat(path)
	if err == nil {
		if dir.IsDir() {
			return nil
		}
		return &os.PathError{Op: "mkdir", Path: path, Err: syscall.ENOTDIR}
	}
	// Slow path: make sure parent exists and then call Mkdir for path.
	i := len(path)
	for i > 0 && os.IsPathSeparator(path[i-1]) { // Skip trailing path separator.
		i--
	}
	j := i
	for j > 0 && !os.IsPathSeparator(path[j-1]) { // Scan backward over element.
		j--
	}
	if j > 1 {
		// Create parent.
		err = MkdirAllAndChown(fixRootDirectory(path[:j-1]), perm, uid, gid)
		if err != nil {
			return err
		}
	}
	// Parent now exists; invoke Mkdir and use its result.
	err = os.Mkdir(path, perm)
	if err != nil {
		// Handle arguments like "foo/." by
		// double-checking that directory doesn't exist.
		dir, err1 := os.Lstat(path)
		if err1 == nil && dir.IsDir() {
			return nil
		}
		return err
	}
	err = os.Chown(path, u, g)
	if err != nil {
		return err
	}
	return nil
}
