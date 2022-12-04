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
	items *[]*models.Item,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT * FROM items WHERE MATCH (name, search_keys, description) "+
			"AGAINST (? IN NATURAL LANGUAGE MODE)", keyword).
		Order("items.`order`").
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
	items *[]*models.Item,
) error {
	var err error
	if pagination.Filter != nil && pagination.Filter.Attributes != nil && len(pagination.Filter.Attributes) > 0 {
		err = datastore.
			GetDB().
			WithContext(ctx).
			Raw("SELECT i.* "+
				"FROM items i "+
				"INNER JOIN item_attribute ia ON i.id = ia.fk_item "+
				"WHERE i.active = ? AND ia.fk_label IN ? LIMIT ?, ?", true, pagination.Filter.Attributes, pagination.Page*pagination.Size, pagination.Size).
			Order("i.`order`").
			Scan(&items).Error
	} else {
		err = datastore.
			GetDB().
			WithContext(ctx).
			Raw("SELECT i.* "+
				"FROM items i "+
				"WHERE i.active = ? "+
				"LIMIT ?, ?", true, pagination.Page*pagination.Size, pagination.Size).
			Order("i.`order`").
			Scan(&items).Error
	}

	if err != nil {
		items = nil
		return fmt.Errorf("error ItemRepository.SearchItemByKeyword %v", err)
	}
	return nil
}

func (i *ItemRepository) GetAvatarOfItems(
	ctx context.Context,
	itemIds []int,
	itemAvatars *[]*models.ItemAvatar,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT items.id as fk_item, images.id as fk_image, images.link "+
			"FROM items INNER JOIN item_images ON items.id = item_images.fk_item "+
			"INNER JOIN images ON item_images.fk_image = images.id "+
			"WHERE items.active = ? AND item_images.is_avatar = TRUE AND items.id IN ? ", true, itemIds).
		Order("items.`order`, item_images.`order`").
		Scan(&itemAvatars).Error
	if err != nil {
		itemAvatars = nil
		return fmt.Errorf("error ItemRepository.GetAvatarOfItems %v", err)
	}

	return nil
}

func (i *ItemRepository) CreateItem(
	ctx context.Context,
	itemDetail *models.Item,
) error {
	err := datastore.GetDB().WithContext(ctx).Create(itemDetail).Error
	if err != nil {
		return fmt.Errorf("error ItemRepository.CreateItem %v", err)
	}

	return nil
}

func (i *ItemRepository) GetListItem(
	ctx context.Context,
	pagination *model.Pagination,
	count *int,
	listItem *[]*models.DetailItem,
) error {
	listItemQuery := datastore.GetDB().WithContext(ctx).Raw("SELECT DISTINCT "+
		"i.id, i.name, i.description, i.`order`, "+
		"GROUP_CONCAT(JSON_OBJECT('id', l.id, 'code', l.code, 'value', l.value, 'subLabels', NULL)) as attributes, "+
		"GROUP_CONCAT(JSON_OBJECT('id', img.id, 'link', img.link)) as images, "+
		"JSON_OBJECT('id', c.id, 'name', c.name, 'order', c.`order`) as collection, "+
		"ic.`order` as orderInCollection "+
		"FROM items i "+
		"LEFT JOIN item_collections ic ON i.id = ic.fk_item "+
		"LEFT JOIN collections c on ic.fk_collection = c.id "+
		"LEFT JOIN item_attributes ia on i.id = ia.fk_item "+
		"LEFT JOIN labels l on ia.fk_label = l.id "+
		"LEFT JOIN item_images ii on i.id = ii.fk_item "+
		"LEFT JOIN images img on ii.fk_image = img.id "+
		"WHERE i.active AND ic.active AND ia.active AND c.active AND l.active AND ii.active "+
		"GROUP BY i.id, i.name, i.description, i.name, i.id, i.`order`, "+
		"JSON_OBJECT('id', c.id, 'name', c.name, 'order', c.`order`), ic.`order` "+
		"LIMIT ?, ?", pagination.Page*pagination.Size, pagination.Size)

	countItemQuery := datastore.GetDB().WithContext(ctx).Raw("SELECT " +
		"COUNT(*) " +
		"FROM items i " +
		"LEFT JOIN item_collections ic ON i.id = ic.fk_item " +
		"LEFT JOIN item_attributes ia on i.id = ia.fk_item " +
		"WHERE i.active AND ic.active AND ia.active",
	)

	if pagination.Collections != nil && len(pagination.Collections) > 0 {
		listItemQuery = listItemQuery.Where("ic.fk_collection IN ?", pagination.Collections)
		countItemQuery = countItemQuery.Where("ic.fk_collection IN ?", pagination.Collections)
	}
	if pagination.Filter != nil && pagination.Filter.Attributes != nil && len(pagination.Filter.Attributes) > 0 {
		listItemQuery = listItemQuery.Where("ia.fk_label IN ?", pagination.Filter.Attributes)
		countItemQuery = countItemQuery.Where("ia.fk_label IN ?", pagination.Filter.Attributes)
	}

	err := countItemQuery.Scan(count).Error
	if err != nil {
		return err
	}

	if pagination.Page*pagination.Size >= *count {
		pagination.Page = *count / pagination.Page
	}

	err = listItemQuery.Scan(listItem).Error
	if err != nil {
		return err
	}

	return nil
}
