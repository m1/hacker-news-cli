package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/m1/hacker-news-cli/pkg/hn/posts"
	"github.com/m1/hacker-news-cli/pkg/hn/storage"
	"os"
)

var (
	savedCmd = cobra.Command{
		Use:     "saved",
		Short:   "Shows saved posts",
		Long:    "Shows saved posts using the save command",
		Example: "saved",
		Run:     saved,
	}
)

func saved(_ *cobra.Command, _ []string) {
	db, err := storage.CreateDbIfNotExist(storage.DbPath)
	if err != nil {
		fmt.Printf("Error gettings local db: %v\n", err.Error())
		os.Exit(1)
	}

	var ps posts.Posts
	db.Order("saved DESC").Find(&ps)

	var js []byte
	if indent {
		js, err = json.MarshalIndent(ps, "", "  ")
	} else {
		js, err = json.Marshal(ps)
	}

	if err != nil {
		fmt.Printf("Error getting your saved posts: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(js))
	os.Exit(0)
}
