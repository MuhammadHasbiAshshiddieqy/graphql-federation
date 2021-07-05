package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"order/graph/generated"
	"order/graph/model"
	"fmt"
)

func (r *queryResolver) Order(ctx context.Context, id int) (*model.Order, error) {
	orderPayment := &model.OrderPayment{}
	orderAddress := &model.OrderAddress{}
	order := &model.Order{}

	var profileIDInt *int
	profileIDInt = new(int)

	profileID := ctx.Value(Key{}).(string)
	*profileIDInt = fixProfileID(profileID)
	if (*profileIDInt == 0) {
		return nil, nil
	}

	err := r.ORDER.Where("profile_id = ? AND id = ?", *profileIDInt, intToString(id)).First(&order).Error
	if err != nil {
		return nil, err
	}
	if len(order.ID) == 0 {
		return nil, nil
	}

	err = r.ORDER.Where("order_id = ?", intToString(id)).First(&orderAddress).Error
	if err != nil {
		return nil, err
	}

	err = r.ORDER.Where("order_id = ?", intToString(id)).First(&orderPayment).Error
	if err != nil {
		return nil, err
	}

	order.Payment = orderPayment
	order.Address = orderAddress
	return order, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
