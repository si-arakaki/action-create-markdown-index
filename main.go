package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/si-arakaki/action-create-markdown-index/lib/env"
)

func main() {
	filepath.WalkDir(env.Home(), func(path string, d fs.DirEntry, err error) error {
		fmt.Printf("filepath.WalkDir\n\tpath: %s\n\tfs.DirEntry: %+v\n\terror: %+v\n", path, d, err)
		return nil
	})
}
