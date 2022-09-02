package repositories

import (
	"context"
	"fmt"
	"github.com/minhtuhcmus/nbh-mono-be/database/datastore"
	"github.com/minhtuhcmus/nbh-mono-be/domain/models"
)

type LabelRepository struct{}

var labelRepository *LabelRepository

func NewLabelRepository() *LabelRepository {
	if labelRepository == nil {
		labelRepository = &LabelRepository{}
	}
	return labelRepository
}

func (l *LabelRepository) FetchLabelByCode(
	ctx context.Context,
	code []string,
	labels []*models.Label,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Where("code IN ?", code).
		Where("active = ?", true).
		Find(&labels).Error
	if err != nil {
		labels = nil
		return fmt.Errorf("error LabelRepository.GetItemAttributeList: %v", err)
	}
	return nil
}

func (l *LabelRepository) GetAllSubAttributesOfGroups(
	ctx context.Context,
	groupIds []int,
	labels []*models.Label,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT * "+
			"FROM labels "+
			"WHERE id IN ? "+
			"AND fk_label IS NULL", groupIds).
		Scan(&labels).Error
	if err != nil {
		labels = nil
		return fmt.Errorf("error LabelRepository.GetLabelsByFkLabel: %v", err)
	}
	return nil
}
