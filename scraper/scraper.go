package scraper

import (
	"github.com/Satr10/1cak-api/model"
	"github.com/gocolly/colly"
)

const (
	onecakBaseURL  = "https://1cak.com"
	randomMemeURL  = "https://1cak.com/shuffle"
	popularMemeURL = "https://1cak.com/legendary"
)

type Scraper struct {
	collector *colly.Collector
}

func NewScraper() *Scraper {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"),
		colly.MaxDepth(1),
	)

	return &Scraper{collector: c}
}

// RandomMeme scrapes a random meme from 1cak.com/shuffle.
//
// It only scrapes the first table element found on the page, and then
// extracts the required information from the HTML elements.
//
// The fields of the returned OneCakPost struct are:
// - PostURL: the URL of the meme post
// - ImageUrl: the URL of the meme image
// - Title: the title of the meme
// - Like: the number of likes the meme has
// - Uploader: the username of the meme uploader
// - UploadTime: the time the meme was uploaded
func (s *Scraper) RandomMeme() (data model.OneCakPost, err error) {
	alreadyFind := false

	s.collector.OnHTML(`table`, func(e *colly.HTMLElement) {
		if alreadyFind {
			return
		}
		alreadyFind = true

		data.PostURL = e.ChildAttr("fb\\:comments-count", "href")
		data.ImageUrl = onecakBaseURL + "/" + e.ChildAttr("img", "src")
		data.Title = e.ChildText("h3")
		data.Like = e.ChildText("span[id]")
		data.Uploader = e.ChildText("b")
		data.UploadTime = e.ChildAttr("abbr", "title")

	})

	err = s.collector.Visit(randomMemeURL)
	if err != nil {
		return model.OneCakPost{}, err
	}
	return data, nil
}

// PopularMemeImages scrapes a list of popular meme images [[7]]
func (s *Scraper) PopularMemeImages() (posts []model.OneCakPost, err error) {

	s.collector.OnHTML(`div[id^='posts']`, func(e *colly.HTMLElement) {
		posts = append(posts, model.OneCakPost{
			PostURL:    "test",
			Title:      e.ChildText("h3"),
			ImageUrl:   onecakBaseURL + e.ChildAttr("a", "href"),
			Like:       e.ChildText("span[id^='span_vote_']"),
			Uploader:   e.ChildText("b"),
			UploadTime: e.ChildAttr("abbr", "title"),
		})
	})

	err = s.collector.Visit(popularMemeURL)
	if err != nil {
		return posts, err
	}

	return posts, err
}
