package main

// Article - Our struct for all articles
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article = []Article{
	{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}
