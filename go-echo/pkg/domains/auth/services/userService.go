package services

import (
	"boilerplates/libs/crypto"
	"boilerplates/pkg/domains/auth/dtos"
	"boilerplates/pkg/domains/auth/interfaces"
	"boilerplates/pkg/infrastructures/models"
	"boilerplates/pkg/infrastructures/repositories"
	"boilerplates/shared/enums"
	sharedI "boilerplates/shared/interfaces"
	"context"

	"github.com/devfeel/mapper"
	"github.com/golobby/container/v3"
)

type userSvc struct {
	logger   sharedI.ILogger
	userRepo interfaces.IUserRepo
	mapper   mapper.IMapper
	crypto   crypto.IPasswordEncoder
}

type IUserSvc interface {
	FindAllUser(ctx context.Context, page, pageSize *int32, status *enums.UserStatus, search *string) (result *[]dtos.UserDto, total int64, err error)
	IsExistUserByEmail(ctx context.Context, email string) (result bool, err error)
	FindUserByEmail(ctx context.Context, email string) (result *dtos.UserDto, err error)
	FindUserById(ctx context.Context, id string) (result *dtos.UserDto, err error)
	CreateOAuthUser(ctx context.Context, req *dtos.UserDto, authenType enums.AuthenType) (result *dtos.UserDto, err error)
}

func NewUserSvc() *userSvc {
	var (
		logger sharedI.ILogger
		mapper mapper.IMapper
		crypto crypto.IPasswordEncoder
	)
	container.Resolve(&logger)
	container.Resolve(&mapper)
	container.Resolve(&crypto)

	userRepo := repositories.NewUserRepo()

	return &userSvc{
		logger:   logger,
		userRepo: userRepo,
		mapper:   mapper,
		crypto:   crypto,
	}
}

func (s *userSvc) FindAllUser(ctx context.Context, page, pageSize *int32, status *enums.UserStatus, search *string) (result *[]dtos.UserDto, total int64, err error) {
	action := "userSvc.FindAllUser"

	r, total, err := s.userRepo.FindAllUser(ctx, page, pageSize, status, search)
	if err != nil {
		s.logger.Errorf("%s error %v", action, err)
		return
	}

	err = s.mapper.Mapper(&r, &result)
	if err != nil {
		return
	}

	return
}

func (s *userSvc) IsExistUserByEmail(ctx context.Context, email string) (result bool, err error) {
	action := "userSvc.IsExistUserByEmail"

	result, err = s.userRepo.ExistUserByEmail(ctx, email)
	if err != nil {
		s.logger.Errorf("%s error %v", action, err)
		return
	}

	return
}

func (s *userSvc) FindUserByEmail(ctx context.Context, email string) (result *dtos.UserDto, err error) {
	action := "userSvc.FindUserByEmail"

	r, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		s.logger.Errorf("%s error %v", action, err)
		return
	}

	result = &dtos.UserDto{}

	err = s.mapper.Mapper(r, result)
	if err != nil {
		return
	}

	return result, nil
}

func (s *userSvc) FindUserById(ctx context.Context, id string) (result *dtos.UserDto, err error) {
	action := "userSvc.FindUserById"

	r, err := s.userRepo.FindUserById(ctx, id)
	if err != nil {
		s.logger.Errorf("%s error %v", action, err)
		return
	}

	result = &dtos.UserDto{}

	err = s.mapper.Mapper(r, result)
	if err != nil {
		return
	}

	return result, nil
}

func (s *userSvc) CreateOAuthUser(ctx context.Context, req *dtos.UserDto, authenType enums.AuthenType) (result *dtos.UserDto, err error) {
	action := "userSvc.CreateUser"

	model := &models.User{}
	err = s.mapper.Mapper(req, model)
	if err != nil {
		return nil, err
	}

	model.AuthType = string(authenType)
	model.Status = enums.USER_STATUS_ACTIVE

	new, err := s.userRepo.Create(ctx, model)
	if err != nil {
		s.logger.Errorf("[%s] error %v", action)
		return nil, err
	}

	req.Id = new.Id

	result = req

	return
}
