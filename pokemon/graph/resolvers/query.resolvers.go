package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"pokemon/graph/generated"
	"pokemon/graph/model"
)

func (r *queryResolver) Pokemons(ctx context.Context, limit *int) ([]*model.Pokemon, error) {
	pokemon := []*model.Pokemon{}
	baseQuery := r.POKEMON
	addPreload(ctx, baseQuery)
	err := baseQuery.Where("deleted_at IS NULL").Find(&pokemon).Limit(*limit).Error
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
