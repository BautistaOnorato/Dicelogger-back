package campaign

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type CampaignHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerGetByUserId() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type CampaignService interface {
	CreateCampaign(dto.CreateCampaignDto, string) (domain.Campaign, error)
	GetAllCampaigns() ([]domain.Campaign, error)
	GetCampaignByID(id int) (dto.ResponseCampaignDto, error)
	GetCampaignsByUserId(cookie string) ([]dto.ResponseCampaignDto, error)
	UpdateCampaign(Campaign dto.CreateCampaignDto, id int) (dto.ResponseCampaignDto, error)
	DeleteCampaign(id int) error
}

type CampaignRepository interface {
	Create(Campaign domain.Campaign) (domain.Campaign, error)
	GetAll() ([]domain.Campaign, error)
	GetById(id int) (domain.Campaign, error)
	GetCampaignsByUserId(id string) ([]domain.Campaign, error)
	GetUsersData(id int) ([]domain.UserResponse, error)
	Update(Campaign domain.Campaign, id int) (domain.Campaign, error)
	Delete(id int) error
}
