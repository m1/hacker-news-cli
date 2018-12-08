package hn_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/m1/hacker-news-cli/pkg/hn"
	"github.com/m1/hacker-news-cli/pkg/hn/config"
	"github.com/m1/hacker-news-cli/pkg/hn/mock"
	"testing"
)

func NewTestClient(config *config.Config, url string) *hn.Client {
	config.UrlBase = url

	return &hn.Client{
		Config: config,
	}

}

func TestClient_GetPosts_Success(t *testing.T) {
	s := mock.NewTestServer()
	defer s.Close()

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	c := NewTestClient(cfg, s.URL)
	posts := c.GetPosts(1)

	mockedPosts := mock.GetMockedPosts()

	assert.Len(t, posts, 1)
	assert.Equal(t, posts[0].Id, mockedPosts[0].Id)
	assert.Equal(t, posts[0].Rank, mockedPosts[0].Rank)
	assert.Equal(t, posts[0].Author, mockedPosts[0].Author)
}
