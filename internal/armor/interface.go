package armor

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type ArmorHandler interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetById() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
}

type ArmorService interface {
	CreateArmor(dto.CreateArmorDto) (domain.Armor, error)
	GetAllArmor() ([]domain.Armor, error)
	GetArmorByID(id int) (domain.Armor, error)
	UpdateArmor(Armor dto.CreateArmorDto, id int) (domain.Armor, error)
	DeleteArmor(id int) error
}

type ArmorRepository interface {
	Create(Armor domain.Armor) (domain.Armor, error)
	GetAllArmors() ([]domain.Armor, error)
	GetArmorById(id int) (domain.Armor, error)
	UpdateArmor(Armor domain.Armor, id int) (domain.Armor, error)
	DeleteArmor(id int) error
}
