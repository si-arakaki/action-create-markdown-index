package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	workspace := "/github/workspace"

	err := filepath.Walk(workspace, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".md" {
			// Markdown ファイルを出力
			fmt.Println(path)
		}
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "error walking the path %q: %v\n", workspace, err)
		os.Exit(1)
	}
}
