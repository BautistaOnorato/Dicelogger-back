package user

import (
	"log"

	"github.com/proyecto-dnd/backend/internal/domain"
)

type service struct {
	repositoryFirebase RepositoryUsers
}

func NewServiceUser(repositoryFirebase RepositoryUsers) ServiceUsers {
	return &service{repositoryFirebase: repositoryFirebase}
}

func userToUserResponse(user domain.User) domain.UserResponse {
	var userResponse domain.UserResponse
	userResponse.Id = user.Id
	userResponse.Username = user.Username
	userResponse.Email = user.Email

	return userResponse
}

func (s *service) Create(user domain.User) (domain.UserResponse, error) {
	return s.repositoryFirebase.Create(user)
}

func (s *service) GetAll() ([]domain.UserResponse, error) {
	return s.repositoryFirebase.GetAll()
}

func (s *service) GetById(id string) (domain.UserResponse, error) {
	return s.repositoryFirebase.GetById(id)
}

func (s *service) Update(user domain.UserUpdate, id string) (domain.UserResponse, error) {

	updatedUser, err := s.repositoryFirebase.Update(user, id)
	if err != nil {
		return domain.UserResponse{}, err
	}

	var updatedUserResponse domain.UserResponse

	updatedUserResponse.Id = updatedUser.Id
	updatedUserResponse.Username = updatedUser.Username
	updatedUserResponse.Email = updatedUser.Email
	updatedUserResponse.Image = updatedUser.Image
	updatedUserResponse.DisplayName = updatedUser.DisplayName

	return updatedUserResponse, nil
}

func (s *service) Patch(user domain.UserUpdate, id string) (domain.UserResponse, error) {
	return s.repositoryFirebase.Patch(user, id)
}

func (s *service) Delete(id string) error {
	err := s.repositoryFirebase.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) Login(userInfo domain.UserLoginInfo) (string, error) {

	cookie, err := s.repositoryFirebase.Login(userInfo)
	if err != nil {
		log.Printf(err.Error())
		return "", err
	}

	return cookie, nil
}

func (s *service) GetJwtInfo(cookieToken string) (domain.UserTokenClaims, error) {
	return s.repositoryFirebase.GetJwtInfo(cookieToken)
}

func (s *service) TransferDataToSql(users []domain.User) (string, error) {

	return s.repositoryFirebase.TransferDataToSql(users)
}

func (s *service) SubscribeToPremium(id string, date string) (string, error) {
	return s.repositoryFirebase.SubscribeToPremium(id, date)
}

func (s *service) CheckSubExpiration(userId string) error {
	return s.repositoryFirebase.CheckSubExpiration(userId)
}

func (s *service) SendVerificationEmail(emailAddress string) error {
	return s.repositoryFirebase.SendVerificationEmail(emailAddress)
}
