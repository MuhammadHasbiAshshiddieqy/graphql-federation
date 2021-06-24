package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"order/graph/generated"
	"order/graph/model"
)

func (r *entityResolver) FindOrderByID(ctx context.Context, id string) (*model.Order, error) {
	return &model.Order{
		ID: id,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
