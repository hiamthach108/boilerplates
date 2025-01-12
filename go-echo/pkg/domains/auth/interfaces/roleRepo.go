package interfaces

import (
	"boilerplates/pkg/domains/auth/dtos"
	"boilerplates/shared/enums"
	"context"
)

type IRoleRepo interface {
	GetAllRoles(ctx context.Context, clientId *string) (result []*dtos.RoleDto, err error)
	GetRoleById(ctx context.Context, roleId string) (result *dtos.RoleDto, err error)
	GetRoleByName(ctx context.Context, roleName string) (result *dtos.RoleDto, err error)

	CreateRole(ctx context.Context, role *dtos.RoleDto) (result *dtos.RoleDto, err error)
	UpdateRole(ctx context.Context, role *dtos.RoleDto) (result *dtos.RoleDto, err error)
	UpdateStatus(ctx context.Context, roleId string, status enums.GeneralStatus) (err error)
}
