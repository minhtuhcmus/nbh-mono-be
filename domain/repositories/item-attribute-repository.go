package repositories

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
)

type ItemAttributeRepository struct{}

var itemAttributeRepository *ItemAttributeRepository

func NewItemAttributeRepository() *ItemAttributeRepository {
	if itemAttributeRepository == nil {
		itemAttributeRepository = &ItemAttributeRepository{}
	}
	return itemAttributeRepository
}

func (ia *ItemAttributeRepository) SaveItemAttributes(ctx context.Context, itemAttributes []*models.ItemAttribute) error {
	err := datastore.GetDB().WithContext(ctx).CreateInBatches(itemAttributes, len(itemAttributes)).Error
	if err != nil {
		return fmt.Errorf("error ItemAttributeRepository.SaveItemAttributes %v", err)
	}
	return nil
}
