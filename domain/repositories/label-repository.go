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

func (l *LabelRepository) FetchAllItemAttributes(
	ctx context.Context,
	itemAttributeList *[]*models.ItemAttributeWithSubLabels,
) error {
	err := datastore.GetDB().WithContext(ctx).Raw("SELECT ml.id, ml.code, ml.value, GROUP_CONCAT(JSON_OBJECT('id', sl.id, 'code', sl.code, 'value', sl.value)) AS labels " +
		"FROM labels ml LEFT JOIN labels sl ON ml.id = sl.fk_label " +
		"WHERE ml.fk_label IS NULL AND ml.active = true AND sl.active = true " +
		"GROUP BY ml.id, ml.code, ml.value").Scan(itemAttributeList).Error
	if err != nil {
		return err
	}
	return nil
}
