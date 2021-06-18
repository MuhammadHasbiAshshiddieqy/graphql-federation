package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"order/graph/generated"
	"order/graph/model"
)

func (r *queryResolver) Order(ctx context.Context) (*model.Order, error) {
	return &model.Order{
		ID:   "2",
		LocalID: "Tom",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
