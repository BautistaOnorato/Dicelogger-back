package user_campaign

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type UserCampaignHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerGetByCampaignId() gin.HandlerFunc
	HandlerGetByUserId() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
	HandlerAddFriendsToCampaign() gin.HandlerFunc
	HandlerAddCharacterToCampaign() gin.HandlerFunc
}

type UserCampaignService interface {
	CreateUserCampaign(dto.CreateUserCampaignDto) (domain.UserCampaign, error)
	GetAllUserCampaigns() ([]domain.UserCampaign, error)
	GetUserCampaignById(id int) (domain.UserCampaign, error)
	GetUserCampaignByCampaignId(id int) ([]domain.UserCampaign, error)
	GetUserCampaignByUserId(id string) ([]domain.UserCampaign, error)
	DeleteUserCampaign(id int) error
	DeleteUserCampaignByCampaignId(id int) error
	AddFriendsToUserCampaign(userIds []string, campaignId int) error
	AddCharacterToCampaign(characterId int, campaignId int, cookie string) error
}

type UserCampaignRepository interface {
	Create(UserCampaign domain.UserCampaign) (domain.UserCampaign, error)
	GetAll()([]domain.UserCampaign, error)
	GetById(id int)(domain.UserCampaign, error)
	GetByCampaignId(id int)([]domain.UserCampaign, error)
	GetByUserId(id string)([]domain.UserCampaign, error)
	Delete(id int)error
	DeleteUserCampaignByCampaignId(id int) error
	AddFriendsToUserCampaign(userIds []string, campaignId int) error
	AddCharacterToCampaign(characterId int, campaignId int, userId string) error
}