package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"customer/graph/generated"
	"customer/graph/model"
)

func (r *entityResolver) FindOrderByID(ctx context.Context, id string) (*model.Order, error) {
	customer := &model.Customer{}
	customerId := &model.Order{}

	var profileIDInt *int
	profileIDInt = new(int)

	profileID := ctx.Value(Key{}).(string)
	*profileIDInt = fixProfileID(profileID)
	if (*profileIDInt == 0) {
		return nil, nil
	}

	err := r.ORDER.Select("customer_id").Where("profile_id = ? AND id = ?", *profileIDInt, id).First(&customerId).Error
	if err != nil {
		return nil, err
	}

	err = r.APRA.Where("profile_id = ? AND id = ?", *profileIDInt, intToString(customerId.CustomerID)).First(&customer).Error
	if err != nil {
		return nil, err
	}

	return &model.Order{
		ID: id,
		Customer: customer,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
