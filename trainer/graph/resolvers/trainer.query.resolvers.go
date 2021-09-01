package resolvers

import (
	"context"
	"trainer/graph/model"
)

func (r *Resolver) findTrainerByID(ctx context.Context, trainerID int) (*model.Trainer, error) {
  trainer := &model.Trainer{}
	baseQuery := r.POKEMON
	addPreload(ctx, baseQuery)
	err := baseQuery.First(&trainer).Error
	if err != nil {
		return nil, err
	}
	return trainer, nil
}
