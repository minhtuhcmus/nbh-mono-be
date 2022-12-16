package repositories

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
	"github.com/minhtuhcmus/nbh-mono-be/graph/model"
)

type CollectionRepository struct{}

var collectionRepository *CollectionRepository

func NewCollectionRepository() *CollectionRepository {
	if collectionRepository == nil {
		collectionRepository = &CollectionRepository{}
	}
	return collectionRepository
}

func (c *CollectionRepository) GetCollectionsInfo(
	ctx context.Context,
	collectionInfos *[]*model.OverviewCollection,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT " +
			"c.id, " +
			"c.name, " +
			"c.`order`, " +
			"count(ic.fk_item) AS totalItem " +
			"FROM collections c " +
			"LEFT JOIN item_collections ic " +
			"ON ic.fk_collection=c.id " +
			"GROUP BY c.id, c.`order` " +
			"ORDER BY `order`").
		Scan(collectionInfos).Error
	if err != nil {
		collectionInfos = nil
		return fmt.Errorf("error CollectionRepository.GetCollectionsInfo: %v", err)
	}
	return nil
}

func (c *CollectionRepository) GetItemsInCollections(
	ctx context.Context,
	pagination *model.PaginationFilter,
	items *[]*models.Item,
) error {
	var err error
	if pagination.Attributes != nil {
		err = datastore.
			GetDB().
			WithContext(ctx).
			Raw("SELECT DISTINCT i.* "+
				"FROM items i "+
				"INNER JOIN item_collections ic ON i.id = ic.fk_item "+
				"INNER JOIN item_attributes ia ON ia.fk_item = i.id "+
				"INNER JOIN collections c ON ic.fk_collection = c.id "+
				"WHERE ic.fk_collection in ? AND ia.fk_label IN ? LIMIT ?, ?",
				pagination.Collections,
				pagination.Attributes,
				pagination.Page*pagination.Size,
				pagination.Size,
			).Order("c.`order`, ic.`order`, i.`order`").Scan(&items).Error
	} else {
		err = datastore.
			GetDB().
			WithContext(ctx).
			Raw("SELECT DISTINCT items.* "+
				"FROM items "+
				"INNER JOIN item_collections ON items.id = item_collections.fk_item "+
				"INNER JOIN collections ON item_collections.fk_collection = collections.id "+
				"WHERE item_collections.fk_collection IN ? AND items.active = TRUE AND item_collections.active = TRUE "+
				"LIMIT ?, ?",
				pagination.Collections,
				pagination.Page*pagination.Size,
				pagination.Size,
			).Order("collections.`order`, item_collections.`order`, items.`order`").Scan(&items).Error
	}

	if err != nil {
		items = nil
		return fmt.Errorf("error CollectionRepository.GetItemsInCollections %v", err)
	}
	return nil
}
