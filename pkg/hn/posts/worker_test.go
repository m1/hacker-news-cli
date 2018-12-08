package posts_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/m1/hacker-news-cli/pkg/hn/config"
	"github.com/m1/hacker-news-cli/pkg/hn/mock"
	"github.com/m1/hacker-news-cli/pkg/hn/posts"
	"testing"
)

func init() {
}
func TestRun_Success(t *testing.T) {
	s := mock.NewTestServer()
	defer s.Close()

	cfg, err := config.NewConfig()
	if err != nil {
		assert.Error(t, err, "couldn't read config")
	}

	cfg.UrlBase = s.URL
	posts := posts.FetchPagePosts(1, cfg)
	mockedPosts := mock.GetMockedPosts()
	assert.Len(t, posts, 30)
	assert.Equal(t, posts[0].Id, mockedPosts[0].Id)
}
