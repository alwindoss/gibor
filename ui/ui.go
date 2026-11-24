package ui

import (
	"embed"
	"io/fs"
)

//go:embed dist
var embeddedFiles embed.FS

func FS() (fs.FS, error) {
	return fs.Sub(embeddedFiles, "dist")
}
