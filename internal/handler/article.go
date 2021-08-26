package handler

import (
	"dateguess-api/internal/model"
	"fmt"
	"net/http"
)

type Article struct {
	router router
	logger logger
	as     articleService
}

type articleService interface {
	RandomArticle(category string) (model.GuardianResult, error)
}

func NewArticle(r router, l logger, as articleService) *Article {
	return &Article{
		router: r,
		logger: l,
		as:     as,
	}
}

func (h *Article) Random(w http.ResponseWriter, r *http.Request) {
	article, err := h.as.RandomArticle("world")
	if err != nil {
		h.router.Error(r.Context(), w, http.StatusInternalServerError, "INTERNAL SERVER ERROR",
			fmt.Errorf("failed to get random article: %w", err),
		)
		return
	}

	h.router.Respond(r.Context(), w, http.StatusOK, article)
}
