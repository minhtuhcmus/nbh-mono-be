package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/minhtuhcmus/nbh-mono-be/graph/generated"
	"github.com/minhtuhcmus/nbh-mono-be/graph/model"
)

// Item is the resolver for the item field.
func (r *mutationResolver) Item(ctx context.Context, id *int, newItem model.NewItem) (*model.OverviewItem, error) {
	item, err := r.itemService.CreateItem(ctx, &newItem)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Images is the resolver for the images field.
func (r *mutationResolver) Images(ctx context.Context, newImage []*model.NewImage) ([]*model.OverviewImage, error) {
	images, err := r.imageService.CreateImages(ctx, newImage)
	if err != nil {
		return nil, err
	}
	return images, nil
}

// Stock is the resolver for the stock field.
func (r *mutationResolver) Stock(ctx context.Context, id *int, newStock model.NewStock) (bool, error) {
	if id != nil {
		return false, nil
	}
	if err := r.stockService.SaveStock(ctx, id, newStock); err != nil {
		return false, err
	}
	return true, nil
}

// StockLog is the resolver for the stockLog field.
func (r *mutationResolver) StockLog(ctx context.Context, id *int, newStockLog model.NewStockLogs) (bool, error) {
	if id != nil {
		return false, nil
	}
	if err := r.stockService.CreateStockLogs(ctx, newStockLog); err != nil {
		return false, err
	}
	return true, nil
}

// Items is the resolver for the items field.
func (r *queryResolver) Items(ctx context.Context, pagination model.PaginationFilter) (*model.ListItem, error) {
	items, err := r.itemService.GetItems(ctx, &pagination)
	if err != nil {
		return nil, fmt.Errorf("error queryResolver.Items %v", err)
	}

	return items, nil
}

// ItemAttributes is the resolver for the itemAttributes field.
func (r *queryResolver) ItemAttributes(ctx context.Context) (*model.ItemAttributes, error) {
	attributes, err := r.itemService.GetAllItemAttribute(ctx)
	if err != nil {
		return nil, fmt.Errorf("error queryResolver.ItemAttributes %v", err)
	}
	return attributes, nil
}

// Collections is the resolver for the collections field.
func (r *queryResolver) Collections(ctx context.Context) ([]*model.OverviewCollection, error) {
	collections, err := r.collectionService.GetCollections(ctx)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// ListDetailItem is the resolver for the listDetailItem field.
func (r *queryResolver) ListDetailItem(ctx context.Context, pagination model.PaginationFilter) (*model.ListDetailItem, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
