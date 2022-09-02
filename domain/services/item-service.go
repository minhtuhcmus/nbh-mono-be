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
	GetItems(ctx context.Context, filter *model.Pagination) ([]*model.OverviewItem, error)
	GetItemAttribute(ctx context.Context) ([]*model.OverviewLabel, error)
}

type itemService struct {
	itemRepository       *repositories.ItemRepository
	collectionRepository *repositories.CollectionRepository
	labelRepository      *repositories.LabelRepository
}

func (i itemService) GetItems(ctx context.Context, filter *model.Pagination) ([]*model.OverviewItem, error) {
	var items []*models.Item
	var err error
	if filter.Keyword != "" {
		err = i.itemRepository.SearchItemByKeyword(ctx, filter.Keyword, items)
	} else {
		err = i.itemRepository.SearchItemByFilter(ctx, filter, items)
	}
	if err != nil {
		return nil, fmt.Errorf("error itemService.GetItems %v", err)
	}

	var itemIds []int
	for idx, item := range items {
		itemIds[idx] = item.ID
	}

	var itemAvatars []*models.ItemAvatar

	err = i.itemRepository.GetAvatarOfItems(ctx, itemIds, itemAvatars)
	if err != nil {
		return nil, fmt.Errorf("error itemService.GetItems %v", err)
	}

	var itemAvatarMap map[int]*model.OverviewImage
	for _, itemAvatar := range itemAvatars {
		itemAvatarMap[itemAvatar.FkItem] = &model.OverviewImage{
			ID:   itemAvatar.FkImage,
			Link: itemAvatar.Link,
		}
	}

	var overviewItems []*model.OverviewItem

	for idx, item := range items {
		overviewItems[idx] = &model.OverviewItem{
			ID:     item.ID,
			Name:   item.Name,
			Avatar: itemAvatarMap[item.ID],
		}
	}

	return overviewItems, nil
}

func (i itemService) GetItemAttribute(ctx context.Context) ([]*model.OverviewLabel, error) {
	var mainAttributes []*models.Label
	err := i.labelRepository.FetchLabelByCode(ctx, constant.ITEM_ATTRIBUTE_CODES, mainAttributes)
	if err != nil {
		return nil, fmt.Errorf("error itemService.GetItemAttribute %v", err)
	}

	var mainAttributeIds []int
	for idx, mainAttr := range mainAttributes {
		mainAttributeIds[idx] = mainAttr.ID
	}

	var subAttributes []*models.Label
	err = i.labelRepository.GetAllSubAttributesOfGroups(ctx, mainAttributeIds, subAttributes)
	if err != nil {
		return nil, fmt.Errorf("error itemService.GetItemAttribute %v", err)
	}

	var attrMap map[int][]*model.OverviewLabel

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

	for idx, mainAttrs := range mainAttributes {
		resLabels[idx] = &model.OverviewLabel{
			ID:        mainAttrs.ID,
			Code:      mainAttrs.Code,
			Value:     mainAttrs.Value,
			SubLabels: attrMap[mainAttrs.ID],
		}
	}

	return resLabels, nil
}

func NewItemService(
	itemRepository *repositories.ItemRepository,
	collectionRepository *repositories.CollectionRepository,
	labelRepository *repositories.LabelRepository,
) ItemService {
	return &itemService{
		itemRepository:       itemRepository,
		collectionRepository: collectionRepository,
		labelRepository:      labelRepository,
	}
}
