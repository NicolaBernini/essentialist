package internal

import (
	"path/filepath"
	"regexp"
)

var mdImageRegex = regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)

// rewriteImagePaths takes a markdown string and a base path, and rewrites any relative image paths to be absolute based on the provided base path.
func rewriteImagePaths(markdown string, basePath string) (string, error) {
	if basePath == "" {
		return markdown, nil
	}
	return mdImageRegex.ReplaceAllStringFunc(markdown, func(match string) string {
		submatches := mdImageRegex.FindStringSubmatch(match)
		// submatches[0] is the full match
		// submatches[1] is the alt text
		// submatches[2] is the path
		dest := submatches[2]

		if !filepath.IsAbs(dest) && dest != "" {
			newDest := filepath.Join(filepath.Dir(basePath), dest)
			return "![" + submatches[1] + "](" + newDest + ")"
		}
		return match
	}), nil
}
