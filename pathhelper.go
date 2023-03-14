package pathhelper

import (
	gopath "path"
	"strings"
)

// Split splits 'path' (using slash as separator)
// into slice of strings containing directory names and filename.
// (E.g. "a/b/c.d" is splitted into {"a", "b", "c.d"} slice.)
func Split(path string) []string {
	path = gopath.Clean(path)
	if path == "." || path == "/" {
		return []string{}
	}
	return strings.Split(strings.Trim(path, "/"), "/")
}

// StartSlash returns 'path' guaranteed to start with a slash.
// If 'path' is empty, empty string is returned.
func StartSlash(path string) string {
	if path == "" {
		return ""
	}
	if path[0] == '/' {
		return path
	}
	return "/" + path
}

// EndSlash returns 'path' guaranteed to end with a slash.
// If 'path' is empty, empty string is returned.
func EndSlash(path string) string {
	if path == "" {
		return ""
	}
	if path[len(path)-1] == '/' {
		return path
	}
	return path + "/"
}
