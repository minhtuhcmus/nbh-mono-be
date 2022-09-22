package services

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/constant"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
	"github.com/minhtuhcmus/nbh-mono-be/domain/repositories"
	"github.com/minhtuhcmus/nbh-mono-be/graph/model"
)

type ItemService interface {
	GetItems(ctx context.Context, filter *model.Pagination) (*[]*model.OverviewItem, error)
	GetItemAttribute(ctx context.Context) (*[]*model.OverviewLabel, error)
	CreateItem(ctx context.Context, itemDetail *model.NewItem) (*model.OverviewItem, error)
}

type itemService struct {
	itemRepository          *repositories.ItemRepository
	collectionRepository    *repositories.CollectionRepository
	labelRepository         *repositories.LabelRepository
	itemAttributeRepository *repositories.ItemAttributeRepository
	itemImageRepository     *repositories.ItemImageRepository
}

func (i itemService) GetItems(ctx context.Context, filter *model.Pagination) (*[]*model.OverviewItem, error) {
	var items []*models.Item
	var err error
	if filter.Keyword != nil {
		err = i.itemRepository.SearchItemByKeyword(ctx, *filter.Keyword, &items)
	} else {
		if filter.Collections != nil {
			err = i.collectionRepository.GetItemsInCollections(ctx, filter, &items)
		} else {
			err = i.itemRepository.SearchItemByFilter(ctx, filter, &items)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("error itemService.GetItems %v", err)
	}

	var itemIds []int
	for _, item := range items {
		itemIds = append(itemIds, item.ID)
	}

	var itemAvatars []*models.ItemAvatar

	err = i.itemRepository.GetAvatarOfItems(ctx, itemIds, &itemAvatars)
	if err != nil {
		return nil, fmt.Errorf("error itemService.GetItems %v", err)
	}

	itemAvatarMap := make(map[int]*model.OverviewImage)
	for _, itemAvatar := range itemAvatars {
		itemAvatarMap[itemAvatar.FkItem] = &model.OverviewImage{
			ID:   itemAvatar.FkImage,
			Link: itemAvatar.Link,
		}
	}

	var overviewItems []*model.OverviewItem

	for _, item := range items {
		overviewItems = append(overviewItems, &model.OverviewItem{
			ID:     item.ID,
			Name:   item.Name,
			Avatar: itemAvatarMap[item.ID],
		})
	}

	return &overviewItems, nil
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
		Name:       itemDetail.Name,
		SearchKeys: *itemDetail.SearchKeys,
	}

	err := i.itemRepository.CreateItem(ctx, newItem)
	if err != nil {
		return nil, fmt.Errorf("error itemService.CreateItem %v", err)
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
) ItemService {
	return &itemService{
		itemRepository:          itemRepository,
		collectionRepository:    collectionRepository,
		labelRepository:         labelRepository,
		itemAttributeRepository: itemAttributeRepository,
		itemImageRepository:     itemImageRepository,
	}
}
