package interfaces

import (
	"boilerplates/pkg/domains/auth/dtos"
	"boilerplates/pkg/infrastructures/models"
	"boilerplates/shared/enums"
	"context"
)

type IUserRepo interface {
	FindAllUser(ctx context.Context, page, pageSize *int32, status *enums.UserStatus, search *string) (result *[]models.User, total int64, err error)
	FindUserByEmail(ctx context.Context, email string) (result *models.User, err error)
	FindUserById(ctx context.Context, userId string) (result *models.User, err error)
	ExistUserByEmail(ctx context.Context, email string) (result bool, err error)
	GetUserAuth(ctx context.Context, userId string) (result *[]models.UserAuth, err error)

	Create(ctx context.Context, user *models.User) (result *models.User, err error)
	UpdateLastLogin(ctx context.Context, userId string) (err error)
	UpdateStatus(ctx context.Context, userId string, status enums.UserStatus) (err error)
	UpdateUser(ctx context.Context, user *models.User) (result *models.User, err error)

	// auth
	UpsertUserAuth(ctx context.Context, userAuth *dtos.UserAuthDto) (result *models.UserAuth, err error)
	DeleteUserAuth(ctx context.Context, userId, clientId, roleId string) (err error)
}
