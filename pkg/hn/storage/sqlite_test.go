package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/m1/hacker-news-cli/pkg/hn/posts"
	"os"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	TestDbPath = "./.test.db"
)

func TestCreateDbIfNotExist_Success(t *testing.T) {
	if _, err := os.Stat(TestDbPath); os.IsNotExist(err) {
		os.Remove(TestDbPath)
	}

	db, err := CreateDbIfNotExist(TestDbPath)
	if err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(TestDbPath); os.IsNotExist(err) {
		t.Errorf(".test.db should exist: %v", err)
	}

	hasTable := db.HasTable(&posts.Posts{})
	assert.True(t, hasTable, "table should exist")

	if _, err := os.Stat(TestDbPath); !os.IsNotExist(err) {
		os.Remove(TestDbPath)
	}
}
