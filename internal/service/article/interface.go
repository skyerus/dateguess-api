package article

import "dateguess-api/internal/model"

type repository interface {
	Search(params model.SearchParams) (model.GuardianContent, error)
}
