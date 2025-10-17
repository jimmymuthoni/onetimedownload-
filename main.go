package main

type VideoRequest struct {
	URL string `json:"url"`
}

type VideoResponse struct {
	URL       string `json:"url"`
	Source    string `json:"source"`
	ID        string `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Medias    []struct {
		URL     string `json:"url"`
		Quality string `json:"quality"`
		Width   int    `json:"width"`
		Height  int    `json:"height"`
		Ext     string `json:"ext"`
	} `json:"medias"`
	Error bool `json:"error"`
}

type YTDLPOutput struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Uploader  string `json:"uploader"`
	Thumbnail string `json:"thumbnail"`
	Formats   []struct {
		URL    string `json:"url"`
		Ext    string `json:"ext"`
		Format string `json:"format"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Acodec string `json:"acodec"`
		Vcodec string `json:"vcodec"`
	} `json:"formats"`
}