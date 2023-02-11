package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/constant"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
	"github.com/minhtuhcmus/nbh-mono-be/domain/repositories"
	"github.com/minhtuhcmus/nbh-mono-be/graph/model"
	"github.com/minhtuhcmus/nbh-mono-be/utils"
	"sort"
	"time"
)

type ItemService interface {
	GetItems(ctx context.Context, filter *model.PaginationFilter) (*model.ListItem, error)
	GetItem(ctx context.Context) *model.DetailItem
	GetItemAttribute(ctx context.Context) (*[]*model.OverviewLabel, error)
	GetAllItemAttribute(ctx context.Context) (*model.ItemAttributes, error)
	CreateItem(ctx context.Context, itemDetail *model.NewItem) (*model.OverviewItem, error)
	GetListDetailItem(ctx context.Context, pagination *model.PaginationFilter) (*model.ListDetailItem, error)
}

type itemService struct {
	itemRepository           *repositories.ItemRepository
	collectionRepository     *repositories.CollectionRepository
	labelRepository          *repositories.LabelRepository
	itemAttributeRepository  *repositories.ItemAttributeRepository
	itemImageRepository      *repositories.ItemImageRepository
	itemCollectionRepository *repositories.ItemCollectionRepository
	stockRepository          *repositories.StockRepository
}

func (i itemService) GetAllItemAttribute(ctx context.Context) (*model.ItemAttributes, error) {
	var itemAttributeList []*models.ItemAttributeWithSubLabels
	err := i.labelRepository.FetchAllItemAttributes(ctx, &itemAttributeList)
	if err != nil {
		return nil, err
	}

	itemAttributes := &model.ItemAttributes{}
	for _, itAttr := range itemAttributeList {
		switch itAttr.Code {
		case constant.ITEM_COLOR_CODE:
			{
				err = json.Unmarshal([]byte("["+itAttr.Labels+"]"), &itemAttributes.Colors)
				if err != nil {
					return nil, err
				}
				sort.Slice(itemAttributes.Colors, func(i, j int) bool {
					return itemAttributes.Colors[i].ID < itemAttributes.Colors[j].ID
				})
				break
			}
		case constant.ITEM_ORIGIN_CODE:
			{
				err = json.Unmarshal([]byte("["+itAttr.Labels+"]"), &itemAttributes.Origins)
				if err != nil {
					return nil, err
				}
				sort.Slice(itemAttributes.Origins, func(i, j int) bool {
					return itemAttributes.Origins[i].ID < itemAttributes.Origins[j].ID
				})
				break
			}
		case constant.ITEM_SIZE_CODE:
			{
				err = json.Unmarshal([]byte("["+itAttr.Labels+"]"), &itemAttributes.Sizes)
				if err != nil {
					return nil, err
				}
				sort.Slice(itemAttributes.Sizes, func(i, j int) bool {
					return itemAttributes.Sizes[i].ID < itemAttributes.Sizes[j].ID
				})
				break
			}
		case constant.ITEM_PRICE_CODE:
			{
				err = json.Unmarshal([]byte("["+itAttr.Labels+"]"), &itemAttributes.Prices)
				if err != nil {
					return nil, err
				}
				sort.Slice(itemAttributes.Prices, func(i, j int) bool {
					return itemAttributes.Prices[i].ID < itemAttributes.Prices[j].ID
				})
				break
			}
		case constant.ITEM_AVAILABILITY_CODE:
			{
				err = json.Unmarshal([]byte("["+itAttr.Labels+"]"), &itemAttributes.Availability)
				if err != nil {
					return nil, err
				}
				sort.Slice(itemAttributes.Availability, func(i, j int) bool {
					return itemAttributes.Availability[i].ID < itemAttributes.Availability[j].ID
				})
				break
			}
		}
	}

	return itemAttributes, nil
}

func (i itemService) GetListDetailItem(ctx context.Context, pagination *model.PaginationFilter) (*model.ListDetailItem, error) {
	var listItem = &model.ListDetailItem{
		Data:      nil,
		Page:      0,
		Size:      0,
		Total:     0,
		IsEndPage: false,
	}

	var details []*models.DetailItem

	err := i.itemRepository.GetListDetailItem(ctx, pagination, &listItem.Total, &details)
	if err != nil {
		return nil, err
	}

	if pagination.Page*pagination.Size >= listItem.Total {
		listItem.IsEndPage = true
	} else {
		listItem.IsEndPage = false
	}

	for _, obj := range details {
		var attributesRaw []model.OverviewLabel
		var collectionRaw model.OverviewCollection
		var imagesRaw []model.OverviewImage

		var attributes []*model.OverviewLabel
		var collection *model.OverviewCollection
		var images []*model.OverviewImage

		if obj.Images != nil {
			err = json.Unmarshal([]byte("["+*obj.Images+"]"), &imagesRaw)
			if err != nil {
				return nil, err
			}
			for _, img := range imagesRaw {
				images = append(images, &model.OverviewImage{
					ID:   img.ID,
					Link: img.Link,
				})
			}
		}

		if obj.Collection != nil {
			err = json.Unmarshal([]byte(*obj.Collection), &collectionRaw)
			if err != nil {
				return nil, err
			}
			collection = &collectionRaw
		}

		if obj.Attributes != nil {
			err = json.Unmarshal([]byte("["+*obj.Attributes+"]"), &attributesRaw)
			if err != nil {
				return nil, err
			}
			for _, attr := range attributesRaw {
				attributes = append(attributes, &model.OverviewLabel{
					ID:        attr.ID,
					Code:      attr.Code,
					Value:     attr.Value,
					SubLabels: nil,
				})
			}
		}

		listItem.Data = append(listItem.Data, &model.DetailItem{
			ID:                obj.ID,
			Name:              obj.Name,
			Description:       obj.Description,
			Order:             obj.Order,
			Attributes:        attributes,
			Images:            images,
			Collection:        collection,
			OrderInCollection: obj.OrderInCollection,
		})
	}

	return listItem, err
}

func (i itemService) GetItem(ctx context.Context) *model.DetailItem {
	//TODO implement me
	panic("implement me")
}

func (i itemService) GetItems(ctx context.Context, filter *model.PaginationFilter) (*model.ListItem, error) {
	listItem := &model.ListItem{
		Data:      []*model.OverviewItem{},
		Page:      filter.Page,
		Size:      filter.Size,
		Total:     0,
		IsEndPage: false,
	}

	if filter.Attributes == nil {
		filter.Attributes = &model.AttributesFilter{
			Colors:       nil,
			Origins:      nil,
			Sizes:        nil,
			Availability: nil,
			Prices:       nil,
		}
	}

	if filter.Collections == nil {
		var collections []int
		err := i.collectionRepository.GetAllCollectionIDs(ctx, &collections)
		if err != nil {
			return nil, err
		}
		filter.Collections = collections
	}

	if filter.Attributes.Sizes == nil {
		var sizes []int
		err := i.labelRepository.FetchAllItemAttributeIDsByCode(ctx, constant.ITEM_SIZE, &sizes)
		if err != nil {
			return nil, err
		}
		filter.Attributes.Sizes = sizes
	}

	if filter.Attributes.Prices == nil {
		var prices []int
		err := i.labelRepository.FetchAllItemAttributeIDsByCode(ctx, constant.ITEM_PRICE, &prices)
		if err != nil {
			return nil, err
		}
		filter.Attributes.Prices = prices
	}

	if filter.Attributes.Colors == nil {
		var colors []int
		err := i.labelRepository.FetchAllItemAttributeIDsByCode(ctx, constant.ITEM_COLOR, &colors)
		if err != nil {
			return nil, err
		}
		filter.Attributes.Colors = colors
	}

	if filter.Attributes.Origins == nil {
		var origins []int
		err := i.labelRepository.FetchAllItemAttributeIDsByCode(ctx, constant.ITEM_ORIGIN, &origins)
		if err != nil {
			return nil, err
		}
		filter.Attributes.Origins = origins
	}

	if filter.Attributes.Availability == nil {
		var availability []int
		err := i.labelRepository.FetchAllItemAttributeIDsByCode(ctx, constant.ITEM_AVAILABILITY, &availability)
		if err != nil {
			return nil, err
		}
		filter.Attributes.Availability = availability
	}

	var overviewItems []*models.OverviewItem
	err := i.itemRepository.SearchItemByFilter(ctx, filter, &overviewItems, &listItem.Total)
	if err != nil {
		return nil, fmt.Errorf("error itemService.GetItems %v", err)
	}

	listItem.Filter = &model.AttributeFilter{}

	if listItem.Total == 0 {
		listItem.Data = []*model.OverviewItem{}
		listItem.IsEndPage = true
		listItem.Page = 0
		listItem.Size = 0
		return listItem, nil
	}

	//var itemIds []int
	//

	var stockAmountList []models.StockAmount
	if err = i.stockRepository.GetStockLogsAvailableOn(ctx, time.Now(), &stockAmountList); err != nil {
		return nil, err
	}

	stockAmountMap := map[int]int{}
	for _, sa := range stockAmountList {
		stockAmountMap[sa.FkItem] = sa.AvailableStock
	}

	for _, it := range overviewItems {
		//itemIds = append(itemIds, it.ID)
		var imageRaw model.OverviewImage
		if it.Avatar != nil {
			err = json.Unmarshal([]byte(*it.Avatar), &imageRaw)
			if err != nil {
				return nil, err
			}
		}
		listItem.Data = append(listItem.Data, &model.OverviewItem{
			ID:          it.ID,
			Name:        it.Name,
			Avatar:      &imageRaw,
			Price:       nil,
			StockAmount: utils.ToPointerInt(stockAmountMap[it.ID]),
		})
	}

	//attributeFilterRaw := models.AttributeFilter{}
	//err = i.itemRepository.GetItemAttributeFilter(ctx, itemIds, &attributeFilterRaw)
	//if err != nil {
	//	return nil, err
	//}

	//err = json.Unmarshal([]byte("["+attributeFilterRaw.Colors+"]"), &attributeFilter.Colors)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = json.Unmarshal([]byte("["+attributeFilterRaw.Origins+"]"), &attributeFilter.Origins)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = json.Unmarshal([]byte("["+attributeFilterRaw.Sizes+"]"), &attributeFilter.Sizes)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = json.Unmarshal([]byte("["+attributeFilterRaw.Availability+"]"), &attributeFilter.Availability)
	//if err != nil {
	//	return nil, err
	//}
	//
	//listItem.Filter = &attributeFilter

	return listItem, nil
}

func (i itemService) GetItemAttribute(ctx context.Context) (*[]*model.OverviewLabel, error) {
	var mainAttributes []*models.Label
	err := i.labelRepository.FetchLabelByCode(ctx, constant.ITEM_ATTRIBUTE_CODES, &mainAttributes)
	if err != nil {
		return nil, fmt.Errorf("error itemService.GetItemAttribute %v", err)
	}

	var mainAttributeIds []int
	for _, mainAttr := range mainAttributes {
		mainAttributeIds = append(mainAttributeIds, mainAttr.ID)
	}

	var subAttributes []*models.Label
	err = i.labelRepository.GetAllSubAttributesOfGroups(ctx, mainAttributeIds, &subAttributes)
	if err != nil {
		return nil, fmt.Errorf("error itemService.GetItemAttribute %v", err)
	}

	var attrMap = map[int][]*model.OverviewLabel{}

	for _, subAttr := range subAttributes {
		if attrMap[*subAttr.FkLabel] != nil {
			attrMap[*subAttr.FkLabel] = append(
				attrMap[*subAttr.FkLabel],
				&model.OverviewLabel{
					ID:        subAttr.ID,
					Code:      subAttr.Code,
					Value:     subAttr.Value,
					SubLabels: nil,
				},
			)
		} else {
			attrMap[*subAttr.FkLabel] = []*model.OverviewLabel{{
				ID:        subAttr.ID,
				Code:      subAttr.Code,
				Value:     subAttr.Value,
				SubLabels: nil,
			}}
		}
	}

	var resLabels []*model.OverviewLabel

	for _, mainAttrs := range mainAttributes {
		resLabels = append(resLabels, &model.OverviewLabel{
			ID:        mainAttrs.ID,
			Code:      mainAttrs.Code,
			Value:     mainAttrs.Value,
			SubLabels: attrMap[mainAttrs.ID],
		})
	}

	return &resLabels, nil
}

func (i itemService) CreateItem(ctx context.Context, itemDetail *model.NewItem) (*model.OverviewItem, error) {
	newItem := &models.Item{
		Name:       *itemDetail.Name,
		SearchKeys: *itemDetail.SearchKeys,
		Active:     true,
	}

	err := i.itemRepository.CreateItem(ctx, newItem)
	if err != nil {
		return nil, fmt.Errorf("error itemService.CreateItem %v", err)
	}

	var newItemCollection []*models.ItemCollection
	newItemCollection = append(newItemCollection, &models.ItemCollection{
		FkItem:       newItem.ID,
		FkCollection: *itemDetail.Type,
		Active:       true,
	})
	err = i.itemCollectionRepository.SaveItemCollections(ctx, newItemCollection)
	if err != nil {
		return nil, fmt.Errorf("error itemService.SaveItemCollection %v", err)
	}

	item := &model.OverviewItem{
		ID:     newItem.ID,
		Name:   newItem.Name,
		Avatar: nil,
	}

	if len(itemDetail.Images) != 0 {
		avatar := &models.ItemImage{
			FkImage:  itemDetail.Images[0],
			FkItem:   item.ID,
			Order:    1,
			IsAvatar: true,
			Active:   true,
		}
		err = i.itemImageRepository.SaveItemImage(ctx, avatar)
		if err != nil {
			return item, fmt.Errorf("error itemService.CreateItem.SaveItemImage %v", err)
		}
		item.Avatar = &model.OverviewImage{
			ID: avatar.FkImage,
		}
	}

	if len(itemDetail.Attributes) != 0 {
		var itemAttributes []*models.ItemAttribute
		for _, attr := range itemDetail.Attributes {
			itemAttributes = append(itemAttributes, &models.ItemAttribute{
				FkLabel: attr,
				FkItem:  newItem.ID,
				Active:  true,
			})
		}
		err = i.itemAttributeRepository.SaveItemAttributes(ctx, itemAttributes)
		if err != nil {
			return item, fmt.Errorf("error itemService.CreateItem.SaveItemAttributes %v", err)
		}
	}

	return item, nil
}

func NewItemService(
	itemRepository *repositories.ItemRepository,
	collectionRepository *repositories.CollectionRepository,
	labelRepository *repositories.LabelRepository,
	itemAttributeRepository *repositories.ItemAttributeRepository,
	itemImageRepository *repositories.ItemImageRepository,
	stockRepository *repositories.StockRepository,
) ItemService {
	return &itemService{
		itemRepository:          itemRepository,
		collectionRepository:    collectionRepository,
		labelRepository:         labelRepository,
		itemAttributeRepository: itemAttributeRepository,
		itemImageRepository:     itemImageRepository,
		stockRepository:         stockRepository,
	}
}
