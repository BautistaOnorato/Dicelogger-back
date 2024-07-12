package user

import (
	"github.com/proyecto-dnd/backend/internal/domain"
)

type RepositoryUsers interface {
	Create(user domain.User) (domain.UserResponse, error)
	GetAll() ([]domain.UserResponse, error)
	GetById(id string) (domain.UserResponse, error)
	Update(user domain.UserUpdate, id string) (domain.UserUpdate, error)
	Delete(id string) error
	Patch(user domain.UserUpdate, id string) (domain.UserResponse, error)
	Login(userInfo domain.UserLoginInfo) (string, error)
	GetJwtInfo(cookieToken string) (domain.UserTokenClaims, error)
	TransferDataToSql(users []domain.User) (string, error)
	SubscribeToPremium(id string, date string) (string, error)
	CheckSubExpiration(userId string) error
	SendVerificationEmail(emailAddress string) error
}

type ServiceUsers interface {
	Create(user domain.User) (domain.UserResponse, error)
	GetAll() ([]domain.UserResponse, error)
	GetById(id string) (domain.UserResponse, error)
	Update(user domain.UserUpdate, id string) (domain.UserResponse, error)
	Patch(user domain.UserUpdate, id string) (domain.UserResponse, error)
	Delete(id string) error
	Login(userInfo domain.UserLoginInfo) (string, error)
	GetJwtInfo(cookieToken string) (domain.UserTokenClaims, error)
	TransferDataToSql(users []domain.User) (string, error)
	SubscribeToPremium(id string, date string) (string, error)
	CheckSubExpiration(userId string) error
	SendVerificationEmail(emailAddress string) error
}
