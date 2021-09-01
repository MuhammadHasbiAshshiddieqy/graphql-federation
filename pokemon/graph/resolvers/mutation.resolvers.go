package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"pokemon/graph/generated"
	"pokemon/graph/model"
	"strings"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreatePokemon(ctx context.Context, input model.CreatePokemonInput) (*model.CreateResponse, error) {
	pokemon_types := []*model.Type{}
	var cnt int64
	err := r.POKEMON.Table("pokemons").Count(&cnt).Error
	if err != nil {
		return &model.CreateResponse{err.Error(), nil}, err
	}
	for _, tp := range input.Types {
		pokemon_type := &model.Type{
			TypeName: strings.ToUpper(tp),
		}
		pokemon_types = append(pokemon_types, pokemon_type)
	}

	pokemon := &model.Pokemon{
		ID:     uuid.New().String(),
		Number: numberForm(int(cnt) + 1),
		Name:   input.Name,
		Types:  pokemon_types,
	}

	err = r.POKEMON.Create(&pokemon).Error
	if err != nil {
		return &model.CreateResponse{err.Error(), nil}, err
	}

	return &model.CreateResponse{"Success", pokemon}, nil
}

func (r *mutationResolver) UpdatePokemon(ctx context.Context, input model.UpdatePokemonInput) (*model.BasicResponse, error) {
	if input.Name != nil {
		statment := fmt.Sprintf("UPDATE pokemons SET name = '%s' WHERE id = '%s'", *input.Name, input.ID)
		err := r.POKEMON.Exec(statment).Error
		if err != nil {
			return &model.BasicResponse{err.Error()}, err
		}
	}

	if input.Types != nil {
		exist_types := []*model.Type{}
		err := r.POKEMON.Where("pokemon_id = ?", input.ID).Find(&exist_types).Error
		if err != nil {
			return &model.BasicResponse{err.Error()}, err
		}
		pokemon_types := []*model.Type{}
	outerloop:
		for _, tp := range input.Types {
			for _, xtp := range exist_types {
				if xtp.TypeName == strings.ToUpper(*tp) {
					continue outerloop
				}
			}
			pokemon_type := &model.Type{
				PokemonID: input.ID,
				TypeName:  strings.ToUpper(*tp),
			}
			pokemon_types = append(pokemon_types, pokemon_type)
		}
		if len(pokemon_types) > 0 {
			err = r.POKEMON.Create(&pokemon_types).Error
			if err != nil {
				return &model.BasicResponse{err.Error()}, err
			}
		}
	}

	return &model.BasicResponse{"Success"}, nil
}

func (r *mutationResolver) DeletePokemon(ctx context.Context, input model.DeletePokemonInput) (*model.BasicResponse, error) {
	statment := fmt.Sprintf("UPDATE pokemons SET deleted_at = NOW() WHERE id = '%s'", input.ID)
	err := r.POKEMON.Exec(statment).Error
	if err != nil {
		return &model.BasicResponse{err.Error()}, err
	}

	return &model.BasicResponse{"Success"}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
