package repositories

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
	"github.com/minhtuhcmus/nbh-mono-be/graph/model"
)

type ItemRepository struct{}

var itemRepository *ItemRepository

func NewItemRepository() *ItemRepository {
	if itemRepository == nil {
		itemRepository = &ItemRepository{}
	}
	return itemRepository
}

func (i *ItemRepository) GetItemByID(
	ctx context.Context,
	itemID int,
	itemDetail *models.Item,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT * "+
			"FROM items "+
			"WHERE id = ?", itemID).
		Scan(&itemDetail).Error
	if err != nil {
		itemDetail = nil
		return fmt.Errorf("error ItemRepository.GetItemByID %v", err)
	}
	return nil
}

func (i *ItemRepository) SearchItemByKeyword(
	ctx context.Context,
	keyword string,
	items []*models.Item,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("MATCH (name, search_keys, description) "+
			"AGAINST (? IN NATURAL LANGUAGE MODE)", keyword).
		Scan(&items).Error
	if err != nil {
		items = nil
		return fmt.Errorf("error ItemRepository.SearchItemByKeyword %v", err)
	}
	return nil
}

func (i *ItemRepository) SearchItemByFilter(
	ctx context.Context,
	pagination *model.Pagination,
	items []*models.Item,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT items.* "+
			"FROM items i "+
			"INNER JOIN item_attribute ia ON i.id = ia.fk_item "+
			"WHERE ia.fk_label IN ? LIMIT ?, ?", pagination.Filter.Attributes, pagination.Page*pagination.Size, pagination.Size).
		Scan(&items).Error
	if err != nil {
		items = nil
		return fmt.Errorf("error ItemRepository.SearchItemByKeyword %v", err)
	}
	return nil
}

func (i *ItemRepository) GetAvatarOfItems(
	ctx context.Context,
	itemIds []int,
	itemAvatars []*models.ItemAvatar,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT items.id, images.id, images.link "+
			"FROM items INNER JOIN item_image ON items.id = item_image.fk_item "+
			"INNER JOIN item_image.fk_image = images.id "+
			"WHERE item_image.is_avatar = TRUE AND items.id IN ?", itemIds).
		Scan(&itemAvatars).Error
	if err != nil {
		itemAvatars = nil
		return fmt.Errorf("error ItemRepository.GetAvatarOfItems %v", err)
	}

	return nil
}
