package posts

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
)

func TestPost_MarshalJSON_Success(t *testing.T) {
	now := time.Now().UTC()
	uri, _ := url.ParseRequestURI("http://google.com")
	post := Post{
		Posted: &now,
		Uri: uri,
	}

	js, err := json.Marshal(post)
	if err != nil {
		assert.Error(t, err, "marshalling failed")
	}

	var parsedPost Post
	err = json.Unmarshal(js, &parsedPost)
	if err != nil {
		assert.Error(t, err, "marshalling failed")
	}


	assert.Equal(t, post.Posted, parsedPost.Posted, "must be equal")
	assert.Equal(t, post.Uri, parsedPost.Uri, "must be equal")
}
