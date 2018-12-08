package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/m1/hacker-news-cli/pkg/hn"
	"github.com/m1/hacker-news-cli/pkg/hn/posts"
	"github.com/m1/hacker-news-cli/pkg/hn/storage"
	"os"
	"strconv"
)

var (
	saveCmd = cobra.Command{
		Use:   "save",
		Short: "Saves post",
		Long:  "Saves post <int>",

		Example: "save 120",
		Args:    cobra.ExactArgs(1),

		Run: save,
	}
)

func save(_ *cobra.Command, args []string) {
	id := args[0]
	postId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Post ID must be an int, input given: %v\n", id)
		os.Exit(1)
	}

	db, err := storage.CreateDbIfNotExist(storage.DbPath)
	if err != nil {
		fmt.Printf("Error gettings local db: %v\n", err.Error())
		os.Exit(1)
	}

	c := hn.NewClient(Cfg)
	post := &posts.Post{}
	post, err = c.GetPost(postId)
	if err != nil {
		fmt.Printf("Error gettings post: %v\n", err.Error())
		os.Exit(1)
	}

	db.NewRecord(post)
	err = db.Create(&post).Error
	if err != nil {
		fmt.Printf("Problem saving this post: %v\n", err.Error())
		os.Exit(1)
	}
}
