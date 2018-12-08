# Hacker News CLI

## Install

### Go

If you have go installed can just `go get github.com/m1/hacker-news-cli` and run `hacker-news-cli`

OR

Can clone and build it yourself, e.g. `GO111MODULE=on go build`


### Make

Can run `make` and will use docker to build.


## Usage

Three basic commands:

### Get posts

`hacker-news-cli` Gets the latest 30 posts from hacker news, can add `--posts <int>` 
to determine how many posts are returned (max of 100), can also add `--indent` to pretty
print the json or use something like [jq](https://stedolan.github.io/jq/). Example:

```
$ hacker-news-cli --posts 30 --indent
```

Returns (truncated)
```json
[
  {
    "id": 18636113,
    "rank": 1,
    "author": "sbenitez",
    "comments": 20,
    "points": 104,
    "title": "Rocket v0.4: Typed URIs, Database Support, Revamped Queries, and More",
    "posted": 1544291822,
    "uri": "https://rocket.rs/v0.4/news/2018-12-08-version-0.4/"
  },
  ...
]
```

### Save post

To save a post the command is just `save <id>` (this uses a sqlite db). Example:

```
$ hacker-news-cli save 18636113
```

### View saved posts

To view the saved posts, use the `saved` command. As with getting the posts, you can 
also use the `--indent` flag here. Example:

```
$ hacker-news-cli saved --indent
```

Returns
```json
[
  {
    "id": 18636113,
    "author": "sbenitez",
    "title": "Rocket v0.4: Typed URIs, Database Support, Revamped Queries, and More",
    "posted": 1544288444,
    "saved": 1544295644,
    "uri": "https://rocket.rs/v0.4/news/2018-12-08-version-0.4/"
  }
]
```

## Notes

- Unit tested with ~70% coverage
- Used sqlite due to portable sql.123
- Uses [colly](https://github.com/gocolly/colly) for scraping, due to 
being nice and simple to use
- Uses [gorm](http://gorm.io/) for db access
- Uses [cobra](https://github.com/spf13/cobra) for command line usability.