package articleservice

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/skyerus/history-api/pkg/article"
	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/models"
)

type articleService struct{}

// NewArticleService ...
func NewArticleService() article.Service {
	return &articleService{}
}

func (as articleService) RandomArticle(category string) (models.Article, customerror.Error) {

}

func randomTime() {
	var start int64
	start = -2208944868 // Jan 1st 1900
	randomUnix := rand.Int63n(time.Now().Unix()-start) + start
	randomTime := time.Unix(randomUnix, 0)
	fmt.Println(randomTime.String())
}
