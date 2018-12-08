package hn

import (
	"github.com/m1/hacker-news-cli/pkg/hn/config"
	"github.com/m1/hacker-news-cli/pkg/hn/posts"
	"math"
	"sort"
	"sync"
)

const (
	urlBase      = "https://news.ycombinator.com"
	postsPerPage = 30
)

type Client struct {
	Config *config.Config
}

func NewClient(config *config.Config) *Client {
	config.UrlBase = urlBase
	return &Client{Config: config}
}

// GetPosts returns list of number<num> of posts
func (c Client) GetPosts(num int) posts.Posts {
	// work out how many pages we need to scrape
	pages := int(math.Ceil(float64(num) / float64(postsPerPage)))
	pagePosts := make(chan posts.Posts, pages)

	// do this asynchronous to make it a little
	// bit faster
	wg := &sync.WaitGroup{}
	for p := 1; p <= pages; p++ {
		wg.Add(1)
		go func(i int, config *config.Config) {
			defer wg.Done()
			pagePosts <- posts.FetchPagePosts(i, config)
		}(p, c.Config)
	}

	wg.Wait()
	close(pagePosts)

	// sort posts into numeric order dependent on rank
	var ps posts.Posts
	for pagePost := range pagePosts {
		ps = append(ps, pagePost...)
	}
	sort.Slice(ps, func(i, j int) bool {
		return *ps[i].Rank < *ps[j].Rank
	})

	// only return the number<num> of posts asked for
	return ps[:num]
}

func (c Client) GetPost(id int) (*posts.Post, error) {
	post, err := posts.FetchPost(id, c.Config)
	if err != nil {
		return nil, err
	}

	return post, nil
}
