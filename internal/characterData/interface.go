package characterdata

import (
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

type RepositoryCharacterData interface {
	Create(character domain.CharacterData) (domain.CharacterData, error)
	GetAll()([]dto.CharacterCardDto, error)
	GetById(id int)(domain.CharacterData, error)
	GetByUserId(userid string)([]dto.CharacterCardDto, error)
	GetByUserIdAndCampaignId(userid string, campaignid int)([]dto.CharacterCardDto, error)
	GetGenerics()([]dto.CharacterCardDto, error)
	GetByCampaignId(campaignid int)([]dto.CharacterCardDto, error)
	GetByAttackEventId(attackeventid int)([]dto.CharacterCardDto, error)
	Update(character domain.CharacterData) (domain.CharacterData, error)
	Delete(id int)error
}

type ServiceCharacterData interface {
	Create(character domain.CharacterData) (dto.FullCharacterData, error)
	GetAll()([]dto.CharacterCardDto, error)
	GetById(id int)(dto.FullCharacterData, error)
	GetByUserId(userid string)([]dto.CharacterCardDto, error)
	GetByUser(cookie string)([]dto.CharacterCardDto, error)
	GetByUserIdAndCampaignId(userid string, campaignid int)([]dto.CharacterCardDto, error)
	GetGenerics()([]dto.CharacterCardDto, error)
	GetByCampaignId(campaignid int)([]dto.CharacterCardDto, error)
	GetByAttackEventId(attackeventid int)([]dto.CharacterCardDto, error)
	Update(character domain.CharacterData) (dto.FullCharacterData, error)
	Delete(id int)error
}