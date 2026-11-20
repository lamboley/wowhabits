// Package webstatic serves the browser assets embedded in the binary under content-hashed names.
package webstatic

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"path"

	"github.com/benbjohnson/hashfs"
)

// URLPrefix is the server route the shared assets are served under.
const URLPrefix = "/_wowhabits/static"

//go:embed static/*
var embedded embed.FS

// FS is the hashed view of the embedded assets, rooted at the static dir.
var FS = hashfs.NewFS(mustSub(embedded, "static"))

func mustSub(fsys fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err)
	}

	return sub
}

// URL returns the content-hashed URL of an asset.
func URL(name string) string {
	return path.Join(URLPrefix, FS.HashName(name))
}

// Funcs returns the template functions for pages.
func Funcs() template.FuncMap {
	return template.FuncMap{"static": URL}
}

// Handler serves the assets.
func Handler() http.Handler {
	return http.StripPrefix(URLPrefix, hashfs.FileServer(FS))
}
