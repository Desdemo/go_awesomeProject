package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)

	//文章详情页
	router.GET("/article/view/:article_id", getArticle)
}
