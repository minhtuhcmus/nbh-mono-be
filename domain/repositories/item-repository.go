package repositories

import (
	"context"
	"fmt"
	sq "github.com/elgris/sqrl"
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
		Raw("SELECT * "+
			"FROM items "+
			"WHERE MATCH (name, search_keys, description) "+
			"AGAINST (? IN NATURAL LANGUAGE MODE) "+
			"LITMIT ?, ?", keyword).
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
	filter *model.PaginationFilter,
	listItem *[]*models.OverviewItem,
	count *int,
) error {
	//selectQuery := sq.Select("DISTINCT `items`.`id`, `items`.`name`, GROUP_CONCAT(JSON_OBJECT('id', images.id, 'link', images.link)) as avatar").From("`items`")
	//countQuery := sq.Select("COUNT(DISTINCT `items`.`id`)").From("`items`")
	//
	//whereFilter := sq.And{}
	//whereFilter = append(whereFilter, sq.Eq{"`items`.`active`": true})
	//
	//if filter.Collections != nil && len(filter.Collections) > 0 {
	//	selectQuery = selectQuery.Join("`item_collections` ON `item_collections`.`fk_item` = `items`.`id`")
	//	countQuery = countQuery.Join("`item_collections` ON `item_collections`.`fk_item` = `items`.`id`")
	//	whereFilter = append(whereFilter, sq.And{
	//		sq.Eq{"`item_collections`.`fk_collection`": filter.Collections},
	//		sq.Eq{"`item_collections`.`active`": true},
	//	})
	//}
	//
	//if filter.Attributes != nil {
	//	var subTables []sq.Sqlizer
	//
	//	values := reflect.ValueOf(*filter.Attributes)
	//	for i := 0; i < values.NumField(); i++ {
	//		if !values.Field(i).IsZero() {
	//			subTables = append(subTables, sq.Expr("SELECT * FROM item_attributes WHERE fk_label IN ?", values.Field(i)))
	//		}
	//	}
	//
	//	selectQuery = selectQuery.Join("`item_attributes` ON `item_attributes`.`fk_item` = `items`.`id`")
	//	countQuery = countQuery.Join("`item_attributes` ON `item_attributes`.`fk_item` = `items`.`id`")
	//
	//	attrQuery := sq.And{}
	//	if filter.Attributes.Origins != nil && len(filter.Attributes.Origins) > 0 {
	//		attrQuery = append(attrQuery, sq.Eq{"`item_attributes`.`fk_label`": filter.Attributes.Origins})
	//	}
	//	if filter.Attributes.Colors != nil && len(filter.Attributes.Colors) > 0 {
	//		attrQuery = append(attrQuery, sq.Eq{"`item_attributes`.`fk_label`": filter.Attributes.Colors})
	//	}
	//	if filter.Attributes.Sizes != nil && len(filter.Attributes.Sizes) > 0 {
	//		attrQuery = append(attrQuery, sq.Eq{"`item_attributes`.`fk_label`": filter.Attributes.Sizes})
	//	}
	//	if filter.Attributes.Availability != nil && len(filter.Attributes.Availability) > 0 {
	//		attrQuery = append(attrQuery, sq.Eq{"`item_attributes`.`fk_label`": filter.Attributes.Availability})
	//	}
	//	if filter.Attributes.Prices != nil && len(filter.Attributes.Prices) > 0 {
	//		attrQuery = append(attrQuery, sq.Eq{"`item_attributes`.`fk_label`": filter.Attributes.Prices})
	//	}
	//
	//	whereFilter = append(whereFilter, sq.And{
	//		attrQuery,
	//		sq.Eq{"`item_attributes`.`active`": true},
	//	})
	//}
	//
	//if filter.Keyword != nil {
	//	whereFilter = append(whereFilter, sq.Expr("MATCH (name, search_keys, description) "+
	//		"AGAINST (? IN NATURAL LANGUAGE MODE)", filter.Keyword))
	//}
	//
	//countQuery = countQuery.Where(whereFilter)
	//countQueryRaw, countQueryParams, err := countQuery.ToSql()
	//if err != nil {
	//	return err
	//}
	//
	//err = datastore.GetDB().WithContext(ctx).
	//	Raw(countQueryRaw, countQueryParams...).Scan(count).Error
	//if err != nil {
	//	return err
	//}
	//
	//if *count == 0 {
	//	listItem = nil
	//	return nil
	//}
	//
	//if filter.Page*filter.Size >= *count {
	//	filter.Page = *count / filter.Page
	//}
	//
	//selectQuery = selectQuery.
	//	LeftJoin("`item_images` ON `item_images`.`fk_item` = `items`.`id`").
	//	LeftJoin("`images` ON `images`.`id` = `item_images`.`fk_image` AND `item_images`.`is_avatar` = ?", true).
	//	Where(whereFilter).
	//	GroupBy("`items`.`id`, `items`.`name`").
	//	Offset(uint64(filter.Page * filter.Size)).
	//	Limit(uint64(filter.Size))
	//
	//selectQueryRaw, selectQueryParams, err := selectQuery.ToSql()
	//if err != nil {
	//	return err
	//}
	//
	//err = datastore.GetDB().WithContext(ctx).
	//	Raw(selectQueryRaw, selectQueryParams...).Scan(&listItem).Error
	//if err != nil {
	//	return err
	//}
	//
	//return nil

	//selectQuery := sq.Select("DISTINCT `items`.`id`, `items`.`name`, GROUP_CONCAT(JSON_OBJECT('id', images.id, 'link', images.link)) as avatar").From("`items`")
	//countQuery := sq.Select("COUNT(DISTINCT `ud`.`id`)").From("`items`")

	var fromBases []sq.Sqlizer

	if filter.Attributes.Origins != nil && len(filter.Attributes.Origins) > 0 {
		fromBases = append(
			fromBases,
			sq.Expr("SELECT i.id "+
				"FROM items i "+
				"LEFT JOIN item_attributes ia ON i.id = ia.fk_item "+
				"WHERE fk_label IN ?",
				filter.Attributes.Origins))
	}
	if filter.Attributes.Colors != nil && len(filter.Attributes.Colors) > 0 {
		fromBases = append(
			fromBases,
			sq.Expr("SELECT i.id "+
				"FROM items i "+
				"LEFT JOIN item_attributes ia ON i.id = ia.fk_item "+
				"WHERE fk_label IN ?",
				filter.Attributes.Colors))
	}
	if filter.Attributes.Sizes != nil && len(filter.Attributes.Sizes) > 0 {
		fromBases = append(
			fromBases,
			sq.Expr("SELECT i.id "+
				"FROM items i "+
				"LEFT JOIN item_attributes ia ON i.id = ia.fk_item "+
				"WHERE fk_label IN ?",
				filter.Attributes.Sizes))
	}
	if filter.Attributes.Availability != nil && len(filter.Attributes.Availability) > 0 {
		fromBases = append(
			fromBases,
			sq.Expr("SELECT i.id "+
				"FROM items i "+
				"LEFT JOIN item_attributes ia ON i.id = ia.fk_item "+
				"WHERE fk_label IN ?",
				filter.Attributes.Availability))
	}
	if filter.Attributes.Prices != nil && len(filter.Attributes.Prices) > 0 {
		fromBases = append(
			fromBases,
			sq.Expr("SELECT i.id "+
				"FROM items i "+
				"LEFT JOIN item_attributes ia ON i.id = ia.fk_item "+
				"WHERE fk_label IN ?",
				filter.Attributes.Prices))
	}

	var fromBaseQueries []string
	var fromBaseParams []interface{}

	var err error
	for idx, fromBase := range fromBases {
		fromBaseQueries[idx], fromBaseParams[idx], err = fromBase.ToSql()
		if err != nil {
			return err
		}
	}

	//baseQuery := ""
	for i := 0; i < len(fromBaseQueries); i++ {

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

func (i *ItemRepository) GetListDetailItem(
	ctx context.Context,
	pagination *model.PaginationFilter,
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
	if pagination.Attributes != nil {
		listItemQuery = listItemQuery.Where("ia.fk_label IN ?", pagination.Attributes)
		countItemQuery = countItemQuery.Where("ia.fk_label IN ?", pagination.Attributes)
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
