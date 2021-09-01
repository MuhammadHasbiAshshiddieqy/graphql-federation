package resolvers

import (
	"trainer/graph/helper"
	"gorm.io/gorm"
	"context"
)

func addPreload(ctx context.Context, baseQuery *gorm.DB) {
	for _, field := range helper.GetPreloads(ctx) {
		*baseQuery = *baseQuery.Preload(field)
	}
}
