package internal

import (
	"path/filepath"
	"testing"
)

func TestRewriteImagePaths(t *testing.T) {
	tests := []struct {
		name        string
		markdown    string
		basePath    string
		expected    string
		expectError bool
	}{
		{
			name:     "relative image path",
			markdown: "![alt text](image.png)",
			basePath: "/path/to/file.md",
			expected: "![alt text](/path/to/image.png)",
		},
		{
			name:     "relative image path with dot slash",
			markdown: "![alt text](./image.png)",
			basePath: "/path/to/file.md",
			expected: "![alt text](/path/to/image.png)",
		},
		{
			name:     "relative image path with parent dir",
			markdown: "![alt text](../image.png)",
			basePath: "/path/to/sub/file.md",
			expected: "![alt text](/path/to/image.png)",
		},
		{
			name:     "absolute image path",
			markdown: "![alt text](/image.png)",
			basePath: "/path/to/file.md",
			expected: "![alt text](/image.png)",
		},
		{
			name:     "no image path",
			markdown: "This is a markdown without an image.",
			basePath: "/path/to/file.md",
			expected: "This is a markdown without an image.",
		},
		{
			name:     "multiple image paths",
			markdown: "![alt1](img1.png) and ![alt2](/img2.png)",
			basePath: "/path/to/file.md",
			expected: "![alt1](/path/to/img1.png) and ![alt2](/img2.png)",
		},
		{
			name:     "image path with subdirectory",
			markdown: "![alt text](./images/image.png)",
			basePath: "/path/to/file.md",
			expected: "![alt text](/path/to/images/image.png)",
		},
		{
			name:     "empty markdown",
			markdown: "",
			basePath: "/path/to/file.md",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// The function under test uses filepath.Join, which will use the OS-specific path separator.
			// For testing purposes, we want to normalize this to always use forward slashes,
			// so the tests are consistent across platforms.
			expected := filepath.ToSlash(tt.expected)

			result, err := rewriteImagePaths(tt.markdown, tt.basePath)
			if (err != nil) != tt.expectError {
				t.Errorf("rewriteImagePaths() error = %v, wantErr %v", err, tt.expectError)
				return
			}

			// Normalize result to use forward slashes for comparison
			result = filepath.ToSlash(result)

			if result != expected {
				t.Errorf("rewriteImagePaths() = %q, want %q", result, expected)
			}
		})
	}
}
