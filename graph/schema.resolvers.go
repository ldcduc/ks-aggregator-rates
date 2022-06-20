package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"ks-aggregator-rates/graph/generated"
	"ks-aggregator-rates/graph/model"
	"ks-aggregator-rates/internal/pkg/client/models"
)

func (r *queryResolver) KyberSwapPrices(ctx context.Context) ([]*model.KyberSwapPrice, error) {
	var prices []*model.KyberSwapPrice
	result := r.DB.Scopes(models.KyberSwapPriceTable()).Find(&prices)
	if result.Error != nil {
		return nil, result.Error
	}

	return prices, nil
}

func (r *queryResolver) OneInchPrices(ctx context.Context) ([]*model.OneInchPrice, error) {
	var prices []*model.OneInchPrice
	result := r.DB.Scopes(models.OneInchPriceTable()).Find(&prices)
	if result.Error != nil {
		return nil, result.Error
	}

	return prices, nil
}

func (r *queryResolver) ParaSwapPrices(ctx context.Context) ([]*model.ParaSwapPrice, error) {
	var prices []*model.ParaSwapPrice
	result := r.DB.Scopes(models.ParaSwapPriceTable()).Find(&prices)
	if result.Error != nil {
		return nil, result.Error
	}

	return prices, nil
}

func (r *queryResolver) ZeroXPrices(ctx context.Context) ([]*model.ZeroXPrice, error) {
	var prices []*model.ZeroXPrice
	result := r.DB.Scopes(models.ZeroXPriceTable()).Find(&prices)
	if result.Error != nil {
		return nil, result.Error
	}

	return prices, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }
