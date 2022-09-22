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
func (r *mutationResolver) Item(ctx context.Context, newItem model.NewItem) (*model.OverviewItem, error) {
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

// Items is the resolver for the items field.
func (r *queryResolver) Items(ctx context.Context, pagination model.Pagination) ([]*model.OverviewItem, error) {
	items, err := r.itemService.GetItems(ctx, &pagination)
	if err != nil {
		return nil, fmt.Errorf("error queryResolver.Items %v", err)
	}

	return *items, nil
}

// ItemAttributes is the resolver for the itemAttributes field.
func (r *queryResolver) ItemAttributes(ctx context.Context) ([]*model.OverviewLabel, error) {
	attributes, err := r.itemService.GetItemAttribute(ctx)
	if err != nil {
		return nil, fmt.Errorf("error queryResolver.ItemAttributes %v", err)
	}
	return *attributes, nil
}

// Collections is the resolver for the collections field.
func (r *queryResolver) Collections(ctx context.Context) ([]*model.OverviewCollection, error) {
	collections, err := r.collectionService.GetCollections(ctx)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) Image(ctx context.Context, newImage model.NewImage) (*model.OverviewImage, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) CreateRole(ctx context.Context, newRole model.NewRole) (*model.OverviewRole, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) DisableRole(ctx context.Context, roleID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) Label(ctx context.Context, newLabel model.NewLabel) (*model.OverviewLabel, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) DisableLabel(ctx context.Context, labelID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) DisableLabels(ctx context.Context, labelID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) User(ctx context.Context, newUser model.NewUser) (*model.OverviewUser, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) DisableUser(ctx context.Context, userID int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Role(ctx context.Context, roleID int) (*model.OverviewRole, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Roles(ctx context.Context, isActive *bool) ([]*model.OverviewRole, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Labels(ctx context.Context, mainLabelID *int) ([]*model.OverviewLabel, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) User(ctx context.Context, id int) (*model.OverviewUser, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Users(ctx context.Context, pagination model.Pagination) ([]*model.OverviewUser, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Item(ctx context.Context, id int) (*model.OverviewItem, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Me(ctx context.Context) (*model.OverviewUser, error) {
	panic(fmt.Errorf("not implemented"))
}
