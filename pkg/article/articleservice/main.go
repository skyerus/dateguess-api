package articleservice

import (
	"errors"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/skyerus/history-api/pkg/article"
	"github.com/skyerus/history-api/pkg/customerror"
	"github.com/skyerus/history-api/pkg/guardian/guardianservice"
	"github.com/skyerus/history-api/pkg/models"
)

type articleService struct{}

// NewArticleService ...
func NewArticleService() article.Service {
	return &articleService{}
}

func (as articleService) RandomArticle(category string) (*models.GuardianResult, customerror.Error) {
	fromTime := randomTime()
	toTime := fromTime.AddDate(0, 6, 0)
	formattedFromTime := fromTime.Format("2006-01-02")
	formattedToTime := toTime.Format("2006-01-02")
	u := url.URL{}
	q := u.Query()
	q.Add("page-size", "1")
	q.Add("page", "1")
	q.Add("section", category)
	q.Add("from-date", formattedFromTime)
	q.Add("to-date", formattedToTime)

	guardianArticle, customErr := guardianservice.Search(q.Encode())
	if customErr != nil {
		return nil, customErr
	}
	if guardianArticle.Response.Total == 0 {
		return as.RandomArticle(category)
	}
	var randPage int
	switch guardianArticle.Response.Total {
	case 0:
		return as.RandomArticle(category)
	case 1:
		randPage = 1
	default:
		randPage = rand.Intn(guardianArticle.Response.Total) + 1
	}
	if randPage != 1 {
		q.Set("page", strconv.Itoa(randPage))
		guardianArticle, customErr = guardianservice.Search(q.Encode())
		if customErr != nil {
			return nil, customErr
		}
	}
	if len(guardianArticle.Response.Results) != 1 {
		return nil, customerror.NewGenericHTTPError(errors.New("Guardian API behaviour changed"))
	}
	return &guardianArticle.Response.Results[0], nil
}

func randomTime() time.Time {
	var start int64
	start = -2208944868 // Jan 1st 1900
	randomUnix := rand.Int63n(time.Now().Unix()-start) + start
	return time.Unix(randomUnix, 0)
}
