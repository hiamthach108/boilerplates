package aggregate

import (
	"time"

	"github.com/hiamthach108/dreon-backend-service/internal/model"
)

type CreateExampleReq struct {
	Name        string `json:"name" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"omitempty,max=4000"`
	IsActive    *bool  `json:"isActive,omitempty"`
}

type UpdateExampleReq struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=1,max=255"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=4000"`
	IsActive    *bool   `json:"isActive,omitempty"`
}

type ExampleAggregate struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (a *ExampleAggregate) FromModel(m *model.Example) {
	if m == nil || a == nil {
		return
	}
	a.ID = m.ID
	a.Name = m.Name
	a.Description = m.Description
	a.IsActive = m.IsActive
	a.CreatedAt = m.CreatedAt
	a.UpdatedAt = m.UpdatedAt
}
