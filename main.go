package main

import (
	"github.com/m1/hacker-news-cli/cmd"
	"log"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
