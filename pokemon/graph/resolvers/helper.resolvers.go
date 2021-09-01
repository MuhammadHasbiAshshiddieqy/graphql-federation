package resolvers

import (
	"pokemon/graph/helper"
	"gorm.io/gorm"
	"context"
	"strings"
  "strconv"
	"fmt"
)

func numberForm(num int) string {
	slip := strconv.Itoa(num)
	if len(slip) != 5 {
		zeroAdd := strings.Repeat("0", 3-len(slip))
		slip = fmt.Sprintf("%s%s", zeroAdd, slip)
	}
	return slip
}

func addPreload(ctx context.Context, baseQuery *gorm.DB) {
	for _, field := range helper.GetPreloads(ctx) {
		*baseQuery = *baseQuery.Preload(field)
	}
}
