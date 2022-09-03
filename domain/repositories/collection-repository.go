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
	collectionInfos *[]model.OverviewCollection,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT " +
			"c.id, " +
			"c.name, " +
			"c.`order`, " +
			"count(*) AS totalItem " +
			"FROM collections c " +
			"INNER JOIN item_collection ic " +
			"ON ic.fk_collection=c.id " +
			"GROUP BY c.id, c.`order` " +
			"ORDER BY `order`").
		Scan(&collectionInfos).Error
	if err != nil {
		collectionInfos = nil
		return fmt.Errorf("error CollectionRepository.GetCollectionsInfo: %v", err)
	}
	return nil
}

func (c *CollectionRepository) GetItemsInCollections(
	ctx context.Context,
	pagination *model.Pagination,
	items *[]*models.Item,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT * "+
			"FROM items INNER JOIN item_collection ON item.id = item_collection.fk_item "+
			"WHERE item_collection.fk_collection IN ? AND active = TRUE "+
			"LIMIT ?, ?", pagination.Collections, pagination.Page*pagination.Size, pagination.Size).
		Scan(&items).Error
	if err != nil {
		items = nil
		return fmt.Errorf("error CollectionRepository.GetItemsInCollections %v", err)
	}
	return nil
}
