package characterdata

import (
	"fmt"
	"log"
	"sync"

	"github.com/proyecto-dnd/backend/internal/armorXCharacterData"
	"github.com/proyecto-dnd/backend/internal/attackEvent"
	characterXproficiency "github.com/proyecto-dnd/backend/internal/characterXProficiency"
	characterXspell "github.com/proyecto-dnd/backend/internal/characterXSpell"
	"github.com/proyecto-dnd/backend/internal/character_feature"
	"github.com/proyecto-dnd/backend/internal/dice_event"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
	"github.com/proyecto-dnd/backend/internal/feature"
	"github.com/proyecto-dnd/backend/internal/itemXCharacterData"
	"github.com/proyecto-dnd/backend/internal/proficiency"
	"github.com/proyecto-dnd/backend/internal/skill"
	skillxcharacterdata "github.com/proyecto-dnd/backend/internal/skillXCharacterData"
	"github.com/proyecto-dnd/backend/internal/spell"
	tradeevent "github.com/proyecto-dnd/backend/internal/tradeEvent"
	"github.com/proyecto-dnd/backend/internal/user"
	weaponxcharacterdata "github.com/proyecto-dnd/backend/internal/weaponXCharacterData"
)

// To do: Optimize query quantity

type service struct {
	characterRepo                RepositoryCharacterData
	itemService                  itemxcharacterdata.ServiceItemXCharacterData
	weaponService                weaponxcharacterdata.ServiceWeaponXCharacterData
	armorService                 armorXCharacterData.ServiceArmorXCharacterData
	skillService                 skill.ServiceSkill
	skillXCharacterService       skillxcharacterdata.ServiceSkillXCharacter
	featureService               feature.FeatureService
	featureXCharacterService     character_feature.CharacterFeatureService
	spellService                 spell.ServiceSpell
	spellXCharacterService       characterXspell.ServiceCharacterXSpell
	proficiencyService           proficiency.ProficiencyService
	proficiencyXCharacterService characterXproficiency.CharacterXProficiencyService
	tradeEventService            tradeevent.ServiceTradeEvent
	attackEventService           attackEvent.AttackEventService
	diceEventService             dice_event.DiceEventService
	userService                  user.ServiceUsers
}

func NewServiceCharacterData(characterRepo RepositoryCharacterData, itemService itemxcharacterdata.ServiceItemXCharacterData, weaponService weaponxcharacterdata.ServiceWeaponXCharacterData, armorService armorXCharacterData.ServiceArmorXCharacterData, skillService skill.ServiceSkill, skillXCharacterService skillxcharacterdata.ServiceSkillXCharacter, featureService feature.FeatureService, featureXCharacterService character_feature.CharacterFeatureService, spellService spell.ServiceSpell, spellXCharacterService characterXspell.ServiceCharacterXSpell, proficiencyService proficiency.ProficiencyService, proficiencyXCharacterService characterXproficiency.CharacterXProficiencyService, tradeEventService tradeevent.ServiceTradeEvent, attackEventService attackEvent.AttackEventService, diceEventService dice_event.DiceEventService, userService user.ServiceUsers) ServiceCharacterData {
	return &service{characterRepo: characterRepo, itemService: itemService, weaponService: weaponService, armorService: armorService, skillService: skillService, skillXCharacterService: skillXCharacterService, featureService: featureService, featureXCharacterService: featureXCharacterService, spellService: spellService, spellXCharacterService: spellXCharacterService, proficiencyService: proficiencyService, proficiencyXCharacterService: proficiencyXCharacterService, tradeEventService: tradeEventService, attackEventService: attackEventService, diceEventService: diceEventService, userService: userService}
}

// GetGenerics implements ServiceCharacterData.
func (s *service) GetGenerics() ([]dto.CharacterCardDto, error) {
	return s.characterRepo.GetGenerics()
}

// Create implements ServiceCharacterData.
func (s *service) Create(character domain.CharacterData) (dto.FullCharacterData, error) {
	newCharacter, err := s.characterRepo.Create(character)
	if err != nil {
		log.Println("characterrepo", err)
		return dto.FullCharacterData{}, err
	}
	newCharacterDto, err := s.GetById(newCharacter.Character_Id)
	if err != nil {
		log.Println("get", err)
		return dto.FullCharacterData{}, err
	}
	return newCharacterDto, nil
}

// Delete implements ServiceCharacterData.
func (s *service) Delete(id int) error {
	errChan := make(chan error, 10)
	maxWorkers := make(chan bool, 3)
	var wg sync.WaitGroup
	wg.Add(10)
	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		err := s.itemService.DeleteByCharacterDataId(id)

		errChan <- err
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		err := s.weaponService.DeleteByCharacterDataId(id)
		errChan <- err

	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		err := s.featureXCharacterService.DeleteByCharacterDataId(id)
		errChan <- err
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		err := s.spellXCharacterService.DeleteByCharacterDataId(id)
		errChan <- err
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		err := s.proficiencyXCharacterService.DeleteByCharacterDataId(id)
		errChan <- err
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		err := s.skillXCharacterService.DeleteByCharacterDataId(id)

		errChan <- err

	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		err := s.armorService.DeleteByCharacterDataIdArmor(id)
		errChan <- err
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		err := s.tradeEventService.DeleteBySenderOrReciever(id)
		errChan <- err
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		err := s.attackEventService.DeleteByProtagonistAndAffectedId(id, id)
		errChan <- err
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			wg.Done()
		}()
		err := s.diceEventService.DeleteByProtagonistId(id)
		errChan <- err
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			fmt.Println(err)
		}
	}
	return s.characterRepo.Delete(id)

}

// GetAll implements ServiceCharacterData.
func (s *service) GetAll() ([]dto.CharacterCardDto, error) {
	return s.characterRepo.GetAll()

}

// GetByCampaignId implements ServiceCharacterData.
func (s *service) GetByCampaignId(campaignid int) ([]dto.CharacterCardDto, error) {
	return s.characterRepo.GetByCampaignId(campaignid)
}

// GetById implements ServiceCharacterData.
func (s *service) GetById(id int) (dto.FullCharacterData, error) {
	character, err := s.characterRepo.GetById(id)
	if err != nil {
		return dto.FullCharacterData{}, err
	}

	fullCharacter, err := s.fetchAndConvertToFullCharacterData(&character)
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	return fullCharacter, nil
}

// GetByUserId implements ServiceCharacterData.
func (s *service) GetByUserId(userid string) ([]dto.CharacterCardDto, error) {
	return s.characterRepo.GetByUserId(userid)
}

// GetByUserIdAndCampaignId implements ServiceCharacterData.
func (s *service) GetByUserIdAndCampaignId(userid string, campaignid int) ([]dto.CharacterCardDto, error) {
	return s.characterRepo.GetByUserIdAndCampaignId(userid, campaignid)
}

// GetByAttackEventId implements ServiceCharacterData.
func (s *service) GetByAttackEventId(attackeventid int) ([]dto.CharacterCardDto, error) {
	return s.characterRepo.GetByAttackEventId(attackeventid)

}

// Update implements ServiceCharacterData.
func (s *service) Update(character domain.CharacterData) (dto.FullCharacterData, error) {
	updatedCharacter, err := s.characterRepo.Update(character)
	if err != nil {
		return dto.FullCharacterData{}, err
	}

	updatedFullCharacter, err := s.GetById(updatedCharacter.Character_Id)
	if err != nil {
		return dto.FullCharacterData{}, err
	}
	return updatedFullCharacter, nil
}

// GetByUser implements ServiceCharacterData.

func (s *service) GetByUser(cookie string) ([]dto.CharacterCardDto, error) {
	var uid string
	user, err := s.userService.GetJwtInfo(cookie)
	if err != nil {
		return nil, err
	}
	uid = user.Id
	return s.characterRepo.GetByUserId(uid)
}

func characterDataToFullCharacterData(character domain.CharacterData, items []domain.ItemXCharacterData, weapons []domain.WeaponXCharacterData, armor []domain.ArmorXCharacterData, skills []domain.Skill, features []domain.Feature, spells []domain.Spell, proficiencies []domain.Proficiency) dto.FullCharacterData {
	return dto.FullCharacterData{
		Character_Id:  character.Character_Id,
		User_Id:       character.User_Id,
		Campaign_Id:   character.Campaign_Id,
		Race:          character.Race,
		Class:         character.Class,
		Background:    character.Background,
		Name:          character.Name,
		Story:         character.Story,
		Alignment:     character.Alignment,
		Age:           character.Age,
		Hair:          character.Hair,
		Eyes:          character.Eyes,
		Skin:          character.Skin,
		Height:        character.Height,
		Weight:        character.Weight,
		ImgUrl:        character.ImgUrl,
		Str:           character.Str,
		Dex:           character.Dex,
		Int:           character.Int,
		Con:           character.Con,
		Wiz:           character.Wiz,
		Cha:           character.Cha,
		Hitpoints:     character.Hitpoints,
		HitDice:       character.HitDice,
		Speed:         character.Speed,
		Armor_Class:   character.Armor_Class,
		Level:         character.Level,
		Exp:           character.Exp,
		Items:         items,
		Weapons:       weapons,
		Armor:         armor,
		Skills:        skills,
		Features:      features,
		Spells:        spells,
		Proficiencies: proficiencies,
	}
}

// TO DO: Finish Armor and skill implementation
func (s *service) fetchAndConvertToFullCharacterData(character *domain.CharacterData) (dto.FullCharacterData, error) {
	errChan := make(chan error, 7)
	itemChan := make(chan []domain.ItemXCharacterData, 1)
	weaponChan := make(chan []domain.WeaponXCharacterData, 1)
	armorChan := make(chan []domain.ArmorXCharacterData, 1)
	featureChan := make(chan []domain.Feature, 1)
	spellChan := make(chan []domain.Spell, 1)
	skillChan := make(chan []domain.Skill, 1)
	proficiencyChan := make(chan []domain.Proficiency, 1)

	maxWorkers := make(chan bool, 3)
	var wg sync.WaitGroup
	wg.Add(7)

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			close(itemChan)
			wg.Done()
		}()
		items, err := s.itemService.GetByCharacterDataId(character.Character_Id)
		errChan <- err
		itemChan <- items
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			close(weaponChan)
			wg.Done()
		}()
		weapons, err := s.weaponService.GetByCharacterDataId(character.Character_Id)
		errChan <- err
		weaponChan <- weapons
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			close(featureChan)
			wg.Done()
		}()
		featuresDto, err := s.featureService.GetAllFeaturesByCharacterId(character.Character_Id)
		errChan <- err
		featureChan <- featuresDto.Features
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			close(spellChan)
			wg.Done()
		}()
		spells, err := s.spellService.GetByCharacterDataId(character.Character_Id)
		errChan <- err
		spellChan <- spells
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			close(proficiencyChan)
			wg.Done()
		}()
		proficiencies, err := s.proficiencyService.GetByCharacterDataId(character.Character_Id)
		errChan <- err
		proficiencyChan <- proficiencies
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			close(skillChan)
			wg.Done()
		}()
		skills, err := s.skillService.GetByCharacterId(character.Character_Id)
		errChan <- err
		skillChan <- skills
	}()

	go func() {
		maxWorkers <- true
		defer func() {
			<-maxWorkers
			close(armorChan)
			wg.Done()
		}()
		armors, err := s.armorService.GetByCharacterDataIdArmor(character.Character_Id)
		errChan <- err
		armorChan <- armors
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			fmt.Println(err)
			return dto.FullCharacterData{}, err
		}
	}

	return characterDataToFullCharacterData(*character, <-itemChan, <-weaponChan, <-armorChan, <-skillChan, <-featureChan, <-spellChan, <-proficiencyChan), nil
}
