package repository

import (
	"context"

	"github.com/hiamthach108/dreon-backend-service/internal/model"
	"gorm.io/gorm"
)

type IExampleRepository interface {
	IRepository[model.Example]
	FindByName(ctx context.Context, name string) *model.Example
}

type exampleRepository struct {
	Repository[model.Example]
}

func NewExampleRepository(dbClient *gorm.DB) IExampleRepository {
	return &exampleRepository{Repository: Repository[model.Example]{dbClient: dbClient}}
}

func (r *exampleRepository) FindByName(ctx context.Context, name string) *model.Example {
	var result model.Example
	if err := r.dbClient.WithContext(ctx).Where("name = ?", name).First(&result).Error; err != nil {
		return nil
	}
	return &result
}
