package campaign

import (
	"sync"

	"github.com/proyecto-dnd/backend/internal/characterData"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/session"
	"github.com/proyecto-dnd/backend/internal/user"
	"github.com/proyecto-dnd/backend/internal/user_campaign"
)

type service struct {
	campaignRepository   CampaignRepository
	sessionService       session.SessionService
	userCampaignService  user_campaign.UserCampaignService
	characterDataService characterdata.ServiceCharacterData
	userService          user.ServiceUsers
}

func NewCampaignService(campaignRepository CampaignRepository, sessionService session.SessionService, userCampaignService user_campaign.UserCampaignService, characterDataService characterdata.ServiceCharacterData, userService user.ServiceUsers) CampaignService {
	return &service{campaignRepository: campaignRepository, sessionService: sessionService, userCampaignService: userCampaignService, characterDataService: characterDataService, userService: userService}
}

func (s *service) CreateCampaign(campaignDto dto.CreateCampaignDto, userId string) (domain.Campaign, error) {
	campaignDomain := domain.Campaign{
		DungeonMaster: userId,
		Name:          campaignDto.Name,
		Description:   campaignDto.Description,
		Image:         campaignDto.Image,
		Notes:         campaignDto.Notes,
		Status:        campaignDto.Status,
		Images:        campaignDto.Images,
	}

	createdCampaign, err := s.campaignRepository.Create(campaignDomain)
	if err != nil {
		return domain.Campaign{}, err
	}

	userCampaignDto := dto.CreateUserCampaignDto{
		CampaignId:  createdCampaign.CampaignId,
		UserId:      userId,
		CharacterId: nil,
		IsOwner:     1,
	}

	if _, err := s.userCampaignService.CreateUserCampaign(userCampaignDto); err != nil {
		return domain.Campaign{}, err
	}

	return createdCampaign, nil
}

func (s *service) GetAllCampaigns() ([]domain.Campaign, error) {
	// campaigns, err := s.campaignRepository.GetAll()
	// if err != nil {
	// 	return nil, err
	// }

	// var responseCampaigns []dto.ResponseCampaignDto
	// for _, campaign := range campaigns {
	// 	sessions, err := s.sessionService.GetSessionsByCampaignId(campaign.CampaignId)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// responseCampaign := dto.ResponseCampaignDto{
	// 	DungeonMaster: campaign.DungeonMaster,
	// 	Name:          campaign.Name,
	// 	Description:   campaign.Description,
	// 	Image:         campaign.Image,
	// 	Notes:         campaign.Notes,
	// 	Status:        campaign.Status,
	// 	Sessions:      sessions,
	// }

	// 	responseCampaigns = append(responseCampaigns, responseCampaign)

	// }

	// return responseCampaigns, nil
	return s.campaignRepository.GetAll()
}

func (s *service) GetCampaignByID(id int) (dto.ResponseCampaignDto, error) {
	campaign, err := s.campaignRepository.GetById(id)
	if err != nil {
		return dto.ResponseCampaignDto{}, err
	}

	errChan := make(chan error, 5)
	sessionsChan := make(chan []domain.Session, 1)
	usersChan := make(chan []domain.UserResponse, 1)

	maxWorkers := make(chan bool, 3)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		sessions, err := s.sessionService.GetSessionsByCampaignId(campaign.CampaignId)

		if err != nil {
			errChan <- err
			return
		}
		sessionsChan <- sessions
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		users, err := s.campaignRepository.GetUsersData(campaign.CampaignId)
		if err != nil {
			errChan <- err
			return
		}
		usersChan <- users
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return dto.ResponseCampaignDto{}, err
		}
	}

	var usersWithCharacters []dto.UserCharacterCampaignDto

	for _, user := range <-usersChan {

		userCharacter, err := s.characterDataService.GetByUserIdAndCampaignId(user.Id, campaign.CampaignId)

		if err != nil {
			return dto.ResponseCampaignDto{}, err
		}

		var character *dto.CharacterCardDto = nil
		if len(userCharacter) > 0 {
			character = &userCharacter[0]
		}

		userWithCharacter := dto.UserCharacterCampaignDto{
			Id:          user.Id,
			Username:    user.Username,
			Email:       user.Email,
			DisplayName: user.DisplayName,
			Image:       user.Image,
			Character:   character,
		}
		usersWithCharacters = append(usersWithCharacters, userWithCharacter)
	}

	responseCampaign := dto.ResponseCampaignDto{
		CampaignId:    campaign.CampaignId,
		DungeonMaster: campaign.DungeonMaster,
		Name:          campaign.Name,
		Description:   campaign.Description,
		Image:         campaign.Image,
		Notes:         campaign.Notes,
		Status:        campaign.Status,
		Sessions:      <-sessionsChan,
		Images:        campaign.Images,
		Users:         usersWithCharacters,
	}

	return responseCampaign, nil
}

func (s *service) GetCampaignsByUserId(cookie string) ([]dto.ResponseCampaignDto, error) {
	var uid string

	user, err := s.userService.GetJwtInfo(cookie)
	if err != nil {
		return nil, err
	}

	uid = user.Id

	campaigns, err := s.campaignRepository.GetCampaignsByUserId(uid)
	if err != nil {
		return nil, err
	}

	var responseCampaigns []dto.ResponseCampaignDto
	for _, campaign := range campaigns {
		sessions, err := s.sessionService.GetSessionsByCampaignId(campaign.CampaignId)
		if err != nil {
			return nil, err
		}

		responseCampaign := dto.ResponseCampaignDto{
			CampaignId:    campaign.CampaignId,
			DungeonMaster: campaign.DungeonMaster,
			Name:          campaign.Name,
			Description:   campaign.Description,
			Image:         campaign.Image,
			Notes:         campaign.Notes,
			Status:        campaign.Status,
			Sessions:      sessions,
		}

		responseCampaigns = append(responseCampaigns, responseCampaign)
	}

	return responseCampaigns, nil
}

func (s *service) UpdateCampaign(campaignDto dto.CreateCampaignDto, id int) (dto.ResponseCampaignDto, error) {
	campaignDomain := domain.Campaign{
		DungeonMaster: campaignDto.DungeonMaster,
		Name:          campaignDto.Name,
		Description:   campaignDto.Description,
		Image:         campaignDto.Image,
		Notes:         campaignDto.Notes,
		Status:        campaignDto.Status,
		Images:        campaignDto.Images,
	}

	updatedCampaign, err := s.campaignRepository.Update(campaignDomain, id)
	if err != nil {
		return dto.ResponseCampaignDto{}, err
	}

	sessions, err := s.sessionService.GetSessionsByCampaignId(updatedCampaign.CampaignId)
	if err != nil {
		return dto.ResponseCampaignDto{}, err
	}

	responseCampaign := dto.ResponseCampaignDto{
		CampaignId:    updatedCampaign.CampaignId,
		DungeonMaster: updatedCampaign.DungeonMaster,
		Name:          updatedCampaign.Name,
		Description:   updatedCampaign.Description,
		Image:         updatedCampaign.Image,
		Notes:         updatedCampaign.Notes,
		Status:        updatedCampaign.Status,
		Images:        updatedCampaign.Images,
		Sessions:      sessions,
	}

	return responseCampaign, nil
}

func (s *service) DeleteCampaign(id int) error {
	err := s.userCampaignService.DeleteUserCampaignByCampaignId(id)
	if err != nil {
		return err
	}
	return s.campaignRepository.Delete(id)
}
