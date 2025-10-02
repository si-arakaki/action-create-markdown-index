package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v6"
)

func main() {
	workspace := "/github/workspace"

	repo, err := git.PlainOpen(workspace)
	if err != nil {
		panic(err)
	}

	tree, err := repo.Worktree()
	if err != nil {
		panic(err)
	}

	status, err := tree.Status()
	if err != nil {
		panic(err)
	}

	fmt.Printf("status: %+v\n", status)

	err = filepath.Walk(workspace, func(path string, info os.FileInfo, err error) error {
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
