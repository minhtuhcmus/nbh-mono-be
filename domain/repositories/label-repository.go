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
	labels *[]*models.Label,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Where("code IN ?", code).
		Where("active = ?", true).
		Order("id").
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
	labels *[]*models.Label,
) error {
	err := datastore.
		GetDB().
		WithContext(ctx).
		Raw("SELECT * "+
			"FROM labels "+
			"WHERE fk_label IN ? "+
			"ORDER BY fk_label, labels.id", groupIds).
		Scan(&labels).Error
	if err != nil {
		labels = nil
		return fmt.Errorf("error LabelRepository.GetLabelsByFkLabel: %v", err)
	}
	return nil
}
