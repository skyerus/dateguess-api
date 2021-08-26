package article

import (
	"dateguess-api/internal/model"
	rand2 "dateguess-api/pkg/rand"
	"errors"
	"fmt"
	"math/rand"
)

// Service provides an API for performing actions on articles
type Service struct {
	repo repository
}

// NewService constructs a Service
func NewService(repo repository) *Service {
	return &Service{repo}
}

func (s *Service) RandomArticle(category string) (model.GuardianResult, error) {
	from := rand2.Time()
	to := from.AddDate(0, 6, 0)
	searchParams := model.SearchParams{
		PageSize: 1,
		Page:     1,
		Section:  category,
		From:     from,
		To:       to,
	}
	article, err := s.repo.Search(searchParams)
	if err != nil {
		return model.GuardianResult{}, fmt.Errorf("failed to search article: %w", err)
	}

	var randPage int
	switch article.Response.Total {
	case 0:
		return s.RandomArticle(category)
	case 1:
		randPage = 1
	default:
		randPage = rand.Intn(article.Response.Total) + 1
	}
	if randPage != 1 {
		searchParams.Page = randPage
		article, err = s.repo.Search(searchParams)
		if err != nil {
			return model.GuardianResult{}, fmt.Errorf("failed to retrieve specific page: %w", err)
		}
	}
	if len(article.Response.Results) != 1 {
		return model.GuardianResult{}, errors.New("guardian API behaviour changed")
	}

	return article.Response.Results[0], nil
}
