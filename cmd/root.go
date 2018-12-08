package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/m1/hacker-news-cli/pkg/hn"
	"github.com/m1/hacker-news-cli/pkg/hn/config"
	"os"
)

var (
	rootCmd  *cobra.Command
	numPosts int
	indent   bool
	Cfg      *config.Config
)

func RootCommand() *cobra.Command {
	var err error
	Cfg, err = config.NewConfig()
	if err != nil {
		panic(err)
	}

	rootCmd = &cobra.Command{
		Run: getPosts,
	}

	rootCmd.Flags().IntVarP(
		&numPosts,
		"posts",
		"p",
		30,
		`Number of posts to return, max is 100`,
	)

	rootCmd.Flags().BoolVarP(
		&indent,
		"indent",
		"i",
		false,
		`Ident the json (pretty print)`,
	)

	rootCmd.AddCommand(&saveCmd)
	rootCmd.AddCommand(&savedCmd)
	savedCmd.Flags().BoolVarP(
		&indent,
		"indent",
		"i",
		false,
		`Ident the json (pretty print)`,
	)

	return rootCmd
}

func getPosts(cmd *cobra.Command, args []string) {
	if numPosts > 100 {
		fmt.Printf("Number of posts to return must be 100 or below, number given: %v", numPosts)
		os.Exit(1)
	}

	c := hn.NewClient(Cfg)
	posts := c.GetPosts(numPosts)
	var js []byte
	var err error
	if indent {
		js, err = json.MarshalIndent(posts, "", "  ")
	} else {
		js, err = json.Marshal(posts)
	}

	if err != nil {
		fmt.Printf("Error converting the posts into json: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(js))
	os.Exit(0)
}
