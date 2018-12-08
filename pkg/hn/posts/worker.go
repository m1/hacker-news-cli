package posts

import (
	"fmt"
	"github.com/go-siris/siris/core/errors"
	"github.com/m1/hacker-news-cli/pkg/hn/config"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	classSpacer      = "spacer"
	classMoreSpace   = "morespace"
	classMainContent = "athing"
	classSubContent  = "subtext"

	advertIdentifier = "advert"

	maxStrLength = 256
)

type Job struct {
	Page   int
	Posts  Posts
	Config *config.Config
	Post   Post
}

func FetchPagePosts(page int, config *config.Config) Posts {
	c := colly.NewCollector()

	// init now to use later
	posts := Posts{}
	post := NewPost()

	selector := fmt.Sprintf(".itemlist > tbody > tr:not(.%v):not(.%v)", classSpacer, classMoreSpace)
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		if e.Attr("class") == classMainContent {
			post = NewPost()
			parseMainContent(e, post, config.UrlBase)
			return
		}

		isSubContent := parseSubContent(e, post)
		if isSubContent {
			p := *post
			posts = append(posts, &p)
			*post = Post{}
		}
	})

	err := c.Visit(fmt.Sprintf("%v/news?p=%v", config.UrlBase, page))
	if err != nil {
		// FIXME: make it a non-silent fail
		return nil
	}

	return posts
}

func FetchPost(id int, config *config.Config) (*Post, error) {
	c := colly.NewCollector()
	post := NewPost()
	selector := fmt.Sprintf(".fatitem > tbody > tr")
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		if e.Attr("class") == classMainContent {
			parseMainContent(e, post, config.UrlBase)
			return
		}

		parseSubContent(e, post)
	})

	err := c.Visit(fmt.Sprintf("%v/item?id=%v", config.UrlBase, id))
	if err != nil {
		return nil, errors.New("post not found")
	}

	return post, nil
}

// This function parses the .athing tr, this holds bulk of the
// data for the post, e.g. title, url etc
func parseMainContent(e *colly.HTMLElement, post *Post, urlBase string) {
	var err error
	id := e.Attr("id")
	post.Id, err = strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	rankStr := strings.Replace(e.ChildText("span.rank"), ".", "", -1)
	*post.Rank, err = strconv.Atoi(rankStr)
	if err != nil {
		*post.Rank = 0
	}

	post.Title = trimStr(e.ChildText(".storylink"))
	storyLink := e.ChildAttr(".storylink", "href")
	post.Uri, err = url.ParseRequestURI(storyLink)
	if err != nil {
		uri := fmt.Sprintf("%v/%v", urlBase, storyLink)
		post.Uri, err = url.ParseRequestURI(uri)
		if err != nil {
			panic(err)
		}
	}
}

// This function parses the tr below .athing, this holds the data
// such as time posted, number of comments etc
func parseSubContent(e *colly.HTMLElement, post *Post) bool {
	// for parsing of the time strings
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	// need to know if its actually a <tr> below .athing
	isSubContent := false

	// should have .subtext due to being being tr below .athing
	e.ForEach("."+classSubContent, func(i int, sub *colly.HTMLElement) {
		isSubContent = true
		author := e.ChildText(".hnuser")
		post.Author = advertIdentifier
		if author != "" {
			post.Author = trimStr(author)
		}

		*post.Points = 0
		points := e.ChildText(".score")
		if points != "" {
			// match point or points
			re := regexp.MustCompile("points?")
			points := strings.TrimSpace(re.ReplaceAllString(points, ""))
			p, err := strconv.Atoi(points)
			if err == nil {
				post.Points = &p
			}
		}

		*post.Comments = 0
		sub.ForEach("a:last-child", func(i int, commentsEl *colly.HTMLElement) {
			commentStr := strings.TrimSpace(strings.Replace(commentsEl.Text, "comments", "", -1))
			comments, err := strconv.Atoi(commentStr)
			if err == nil {
				post.Comments = &comments
			}
		})

		age := e.ChildText(".age a")
		r, err := w.Parse(age, time.Now())
		if err == nil {
			posted := r.Time
			post.Posted = &posted
		}
	})

	return isSubContent
}

// FIXME: probably should add to another package
func trimStr(str string) string {
	if len(str) > maxStrLength {
		str = str[0:maxStrLength-3] + "..."
	}

	return str
}
