package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"user/graph/generated"
	"user/graph/model"
)

func (r *entityResolver) FindOrderByID(ctx context.Context, id string) (*model.Order, error) {
	account := &model.ProfileChannelAssociation{}
	channel := &model.Channel{}
	relationId := &model.Order{}

	err := r.ORDER.Select("channel_id, account_id").Where("id = ?", id).First(&relationId).Error
	if err != nil {
		return nil, err
	}

	err = r.APRA.Where("id = ?", intToString(relationId.ChannelID)).First(&channel).Error
	if err != nil {
		return nil, err
	}

	err = r.APRA.Where("id = ?", intToString(relationId.AccountID)).First(&account).Error
	if err != nil {
		return nil, err
	}

	return &model.Order{
		ID:        id,
		AccountID: relationId.AccountID,
		ChannelID: relationId.ChannelID,
		Channel:   channel,
		Account:   account,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
