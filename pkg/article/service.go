package article

import (
	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/models"
)

// Service ...
type Service interface {
	RandomArticle(category string) (models.Article, customerror.Error)
}
