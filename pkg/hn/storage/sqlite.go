package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/m1/hacker-news-cli/pkg/hn/posts"
	"os"
)

const (
	DbPath = "./.hn.db"
)

type Db struct {
	*gorm.DB
}

func GetDb(path string) (db *Db, err error) {
	db = &Db{}
	db.DB, err = gorm.Open("sqlite3", path)
	db.LogMode(false)
	return
}

func CreateDbIfNotExist(path string) (*Db, error) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		db, err := GetDb(path)
		if err != nil {
			return nil, err
		}

		hasTable := db.HasTable(&posts.Post{})
		if hasTable {
			return db, nil
		}

		// remove db due to not having tables
		if err = os.Remove(path); err != nil {
			return nil, err
		}
	}

	if _, err := os.Create(path); err != nil {
		return nil, err
	}

	db, err := GetDb(path)
	if err != nil {
		return nil, err
	}

	// could possibly use gorm migration tools rather
	// than hard coding the db schema.
	postTableQuery := `CREATE TABLE posts (
  						id     INTEGER       NOT NULL PRIMARY KEY,
  						author VARCHAR(300)  NOT NULL,
  						title  VARCHAR(300)  NOT NULL,
  						uri    VARCHAR(2083) NOT NULL,
  						posted DATETIME      NOT NULL,
  						saved  DATETIME      NOT NULL
					)`

	if err = db.Exec(postTableQuery).Error; err != nil {
		return nil, err
	}

	return db, nil
}
