package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Article1", Content: "Article1 body"},
	article{ID: 2, Title: "Article2", Content: "Article2 body"},
}

func getAllArticles() []article {
	return articleList
}
