package lolparser

import "io"

type Update struct {
	Title string
	Text string
	Url string
	Image io.Reader
	ImageUrl string
}