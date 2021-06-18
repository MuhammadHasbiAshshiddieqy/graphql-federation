package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"customer/graph/generated"
	"customer/graph/model"
)

func (r *entityResolver) FindOrderByID(ctx context.Context, id string) (*model.Order, error) {
	return &model.Order{
		ID: id,
		Customer: &model.Customer{
			ID:          "1",
			FirstName:   "Hasbi",
			LastName:    "Ashshiddieqy",
			CompanyName: "Forstok",
			Email:       "hasbi@forstok.com",
			Phone:       "0918283",
			ProfileID:   12,
		},
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
