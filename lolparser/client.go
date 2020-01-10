package lolparser

import (
	"github.com/alexbyk/panicif"
	"github.com/anaskhan96/soup"
	"io"
	"net/http"
	"strings"
)

func GetDocument(url string) soup.Root{
	soup.Headers = map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36 OPR/65.0.3467.78"}
	response, err := soup.Get(url)
	panicif.Err(err)
	doc := soup.HTMLParse(response)
	return doc
}

func GetLastUpdates() []string {
	url := "https://ru.leagueoflegends.com/ru/news/game-updates/patch"
	doc := GetDocument(url)
	allElements := doc.FindAll("div")
	var href string
	var lastUpdates []string
	for _, element := range allElements {
		if strings.HasPrefix(element.Attrs()["class"], "views-row views-row-"){
			href = element.Find("a").Attrs()["href"]
			lastUpdates = append(lastUpdates, "https://ru.leagueoflegends.com" + href)
		}
	}
	return lastUpdates
}

func GetFile(url string) io.Reader {
	response, err := http.Get(url)
	panicif.Err(err)
	return response.Body
}

func GetUpdateInfo(url string) Update{
	doc := GetDocument(url)
	title := doc.Find("h1", "class", "article-title").Text()
	text := doc.Find("blockquote").FullText()
	imageUrl := "https://ru.leagueoflegends.com" + doc.Find("img").Attrs()["src"]
	image := GetFile(imageUrl)

	return Update{
		Title: title,
		Text: text,
		Image: image,
		ImageUrl: imageUrl,
		Url: url,
	}
}
