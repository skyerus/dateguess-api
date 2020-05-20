package api

import (
	"net/http"

	"github.com/skyerus/history-api/pkg/article/articleservice"
)

func (router router) randomArticle(w http.ResponseWriter, r *http.Request) {
	articleService := articleservice.NewArticleService()
	article, customErr := articleService.RandomArticle("world")
	if customErr != nil {
		handleError(w, customErr)
		return
	}

	respondJSON(w, http.StatusOK, article)
}
