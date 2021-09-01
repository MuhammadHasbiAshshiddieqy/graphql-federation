package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"trainer/graph/generated"
	"trainer/graph/model"
)

func (r *entityResolver) FindTrainerByID(ctx context.Context, id int) (*model.Trainer, error) {
	return r.findTrainerByID(ctx, id)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
