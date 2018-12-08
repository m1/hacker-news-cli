package posts

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/url"
	"time"
)

type Post struct {
	Id        int        `json:"id" gorm:"primary_key;column:id"`
	Rank      *int       `json:"rank,omitempty" gorm:"-"`
	Author    string     `json:"author" gorm:"column:author"`
	Comments  *int       `json:"comments,omitempty" gorm:"-"`
	Points    *int       `json:"points,omitempty" gorm:"-"`
	Posted    *time.Time `json:"posted" gorm:"column:posted"`
	Saved     *time.Time `json:"saved,omitempty" gorm:"column:saved"`
	Title     string     `json:"title" gorm:"column:title"`
	Uri       *url.URL   `json:"uri" gorm:"-"`
	UriString *string    `json:"-" gorm:"column:uri"`
}

func NewPost() *Post {
	return &Post{
		// init these due to omitempty and pointers
		Rank:     new(int),
		Comments: new(int),
		Points:   new(int),
	}
}

func (p *Post) MarshalJSON() ([]byte, error) {
	type Alias Post

	var savedUnix *int64 = nil
	if p.Saved != nil {
		unix := p.Saved.Unix()
		savedUnix = &unix
	}

	return json.Marshal(&struct {
		*Alias
		Posted int64  `json:"posted"`
		Saved  *int64 `json:"saved,omitempty"`
		Uri    string `json:"uri"`
	}{
		Alias:  (*Alias)(p),
		Posted: p.Posted.Unix(),
		Saved:  savedUnix,
		Uri:    p.Uri.String(),
	})
}

func (p *Post) UnmarshallJSON(data []byte) error {
	type Alias Post

	t := &struct {
		*Alias
		Posted int64  `json:"posted"`
		Uri    string `json:"uri"`
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	posted := time.Unix(t.Posted, 0)
	p.Posted = &posted

	var err error
	p.Uri, err = url.ParseRequestURI(t.Uri)
	if err != nil {
		return err
	}

	return nil
}

func (p *Post) BeforeSave(scope *gorm.Scope) error {
	_ = scope.SetColumn("saved", time.Now())
	_ = scope.SetColumn("uri", p.Uri.String())
	return nil
}

func (p *Post) AfterFind(scope *gorm.Scope) error {
	p.Uri, _ = url.ParseRequestURI(*p.UriString)
	return nil
}

type Posts []*Post
