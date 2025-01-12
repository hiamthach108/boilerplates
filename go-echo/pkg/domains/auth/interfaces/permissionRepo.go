package interfaces

import (
	"boilerplates/pkg/domains/auth/dtos"
	"boilerplates/shared/enums"
	"context"
)

type IPermissionRepo interface {
	GetAllPermissions(ctx context.Context, page, pageSize *int32, search *string) (result []*dtos.PermissionDto, total int64, err error)
	GetPermissionById(ctx context.Context, permissionId string) (result *dtos.PermissionDto, err error)

	CreatePermission(ctx context.Context, permission *[]dtos.PermissionDto) (result int32, err error)
	UpdatePermission(ctx context.Context, permission *dtos.PermissionDto) (result *dtos.PermissionDto, err error)
	UpdateStatus(ctx context.Context, permissionId string, status enums.GeneralStatus) (err error)
}
