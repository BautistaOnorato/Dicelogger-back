package characterdata

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrNotFound            = errors.New("characters not found")
)

type CharacterDataMySqlRepository struct {
	db *sql.DB
}

func NewCharacterDataRepository(db *sql.DB) RepositoryCharacterData {
	return &CharacterDataMySqlRepository{db}
}

// GetGenerics implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetGenerics() ([]dto.CharacterCardDto, error) {
	rows, err := r.db.Query(QueryGetGenerics)
	if err != nil {
		return []dto.CharacterCardDto{}, err
	}
	defer rows.Close()

	var characterDtoList []dto.CharacterCardDto

	for rows.Next() {
		var characterDto dto.CharacterCardDto
		err := ScanCharacterCardDto(rows, &characterDto)
		if err != nil {
			return []dto.CharacterCardDto{}, err
		}
		characterDtoList = append(characterDtoList, characterDto)
	}
	if err := rows.Err(); err != nil {
		return []dto.CharacterCardDto{}, err
	}

	return characterDtoList, nil
}

// Create implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) Create(character domain.CharacterData) (domain.CharacterData, error) {
	statement, err := r.db.Prepare(QueryCreateCharacter)
	if err != nil {
		log.Println(1, err)
		return domain.CharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()
	result, err := statement.Exec(
		character.User_Id,
		character.Campaign_Id,
		character.Race.RaceID,
		character.Class.ClassId,	
		character.Background.BackgroundID,
		character.Name,
		character.Story,
		character.Alignment,
		character.Age,
		character.Hair,
		character.Eyes,
		character.Skin,
		character.Height,
		character.Weight,
		character.ImgUrl,
		character.Str,
		character.Dex,
		character.Int,
		character.Con,
		character.Wiz,
		character.Cha,
		character.Hitpoints,
		character.HitDice,
		character.Speed,
		character.Armor_Class,
		character.Level,
		character.Exp,
	)
	if err != nil {
		log.Println(2, err)
		return domain.CharacterData{}, err
	}
	
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(3, err)
		return domain.CharacterData{}, ErrGettingLastInsertId
	}
	character.Character_Id = int(lastId)
	return character, nil
}

// Delete implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) Delete(id int) error {
	result, err := r.db.Exec(QueryDelete, id)
	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected < 1 {
		return ErrNotFound
	}

	return nil
}

// GetAll implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetAll() ([]dto.CharacterCardDto, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []dto.CharacterCardDto{}, err
	}
	defer rows.Close()

	var characterDtoList []dto.CharacterCardDto

	for rows.Next() {
		var characterDto dto.CharacterCardDto
		err := ScanCharacterCardDto(rows, &characterDto)
		if err != nil {
			return []dto.CharacterCardDto{}, err
		}
		characterDtoList = append(characterDtoList, characterDto)
	}
	if err := rows.Err(); err != nil {
		return []dto.CharacterCardDto{}, err
	}

	return characterDtoList, nil
}

// GetByCampaignId implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetByCampaignId(campaignid int) ([]dto.CharacterCardDto, error) {
	rows, err := r.db.Query(QueryGetByCampaignId, campaignid)
	if err != nil {
		return []dto.CharacterCardDto{}, err
	}
	defer rows.Close()

	var characterDtoList []dto.CharacterCardDto

	for rows.Next() {
		var characterDto dto.CharacterCardDto
		err := ScanCharacterCardDto(rows, &characterDto)
		if err != nil {
			return []dto.CharacterCardDto{}, err
		}
		characterDtoList = append(characterDtoList, characterDto)
	}
	if err := rows.Err(); err != nil {
		return []dto.CharacterCardDto{}, err
	}

	return characterDtoList, nil
}

// GetById implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetById(id int) (domain.CharacterData, error) {
	row := r.db.QueryRow(QueryGetById, id)
	var character domain.CharacterData
	err := ScanCharacterData(row, &character)
	if err != nil {
		fmt.Println(err)
		return domain.CharacterData{}, ErrNotFound
	}

	return character, nil
}

// GetByUserId implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetByUserId(userid string) ([]dto.CharacterCardDto, error) {
	rows, err := r.db.Query(QueryGetByUserId, userid)
	if err != nil {
		return []dto.CharacterCardDto{}, err
	}
	defer rows.Close()

	var characterDtoList []dto.CharacterCardDto

	for rows.Next() {
		var characterDto dto.CharacterCardDto

		err := ScanCharacterCardDto(rows, &characterDto)
		if err != nil {
			return []dto.CharacterCardDto{}, err
		}
		characterDtoList = append(characterDtoList, characterDto)
	}
	if err := rows.Err(); err != nil {
		return []dto.CharacterCardDto{}, err
	}

	return characterDtoList, nil
}

// GetByUserIdAndCampaignId implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetByUserIdAndCampaignId(userid string, campaignid int) ([]dto.CharacterCardDto, error) {
	rows, err := r.db.Query(QueryGetByUserIdAndCampaignId, userid, campaignid)
	if err != nil {
		return []dto.CharacterCardDto{}, err
	}
	defer rows.Close()

	var characterDtoList []dto.CharacterCardDto

	for rows.Next() {
		var characterDto dto.CharacterCardDto

		err := ScanCharacterCardDto(rows, &characterDto)
		if err != nil {
			return []dto.CharacterCardDto{}, err
		}
		characterDtoList = append(characterDtoList, characterDto)
	}
	if err := rows.Err(); err != nil {
		return []dto.CharacterCardDto{}, err
	}

	return characterDtoList, nil
}

// GetByAttackEventId implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) GetByAttackEventId(attackeventid int) ([]dto.CharacterCardDto, error) {
	rows, err := r.db.Query(QueryGetByAttackEventId, attackeventid)
	if err != nil {
		return []dto.CharacterCardDto{}, err
	}
	defer rows.Close()

	var characterDtoList []dto.CharacterCardDto

	for rows.Next() {
		var characterDto dto.CharacterCardDto

		err := ScanCharacterCardDto(rows, &characterDto)
		if err != nil {
			return []dto.CharacterCardDto{}, err
		}
		characterDtoList = append(characterDtoList, characterDto)
	}
	if err := rows.Err(); err != nil {
		return []dto.CharacterCardDto{}, err
	}

	return characterDtoList, nil
}

// Update implements RepositoryCharacterData.
func (r *CharacterDataMySqlRepository) Update(character domain.CharacterData) (domain.CharacterData, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.CharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()
	_, err = statement.Exec(
		character.User_Id,
		character.Campaign_Id,
		character.Race.RaceID,
		character.Class.ClassId,
		character.Background.BackgroundID,
		character.Name,
		character.Story,
		character.Alignment,
		character.Age,
		character.Hair,
		character.Eyes,
		character.Skin,
		character.Height,
		character.Weight,
		character.ImgUrl,
		character.Str,
		character.Dex,
		character.Int,
		character.Con,
		character.Wiz,
		character.Cha,
		character.Hitpoints,
		character.HitDice,
		character.Speed,
		character.Armor_Class,
		character.Level,
		character.Exp,
		character.Character_Id,
	)
	if err != nil {
		return domain.CharacterData{}, err
	}
	return character, nil
}

type scannable interface {
	Scan(dest ...any) error
}

func ScanCharacterCardDto(rows scannable, characterCardDto *dto.CharacterCardDto) error {
	err := rows.Scan(
		&characterCardDto.CharacterId,
		&characterCardDto.UserId,
		&characterCardDto.CampaignID,
		&characterCardDto.ImageUrl,
		&characterCardDto.Name,
		&characterCardDto.Race,
		&characterCardDto.Class,
		&characterCardDto.Level,
		&characterCardDto.HitPoints,
	)
	return err
}

func ScanCharacterData(rows scannable, characterData *domain.CharacterData) error {
	err := rows.Scan(
		&characterData.Character_Id,
		&characterData.User_Id,
		&characterData.Campaign_Id,
		&characterData.Race.RaceID,
		&characterData.Race.Name,
		&characterData.Race.Description,
		&characterData.Race.Speed,
		&characterData.Race.Str,
		&characterData.Race.Dex,
		&characterData.Race.Int,
		&characterData.Race.Con,
		&characterData.Race.Wiz,
		&characterData.Race.Cha,
		&characterData.Class.ClassId,
		&characterData.Class.Name,
		&characterData.Class.Description,
		&characterData.Class.ProficiencyBonus,
		&characterData.Class.HitDice,
		&characterData.Class.ArmorProficiencies,
		&characterData.Class.WeaponProficiencies,
		&characterData.Class.ToolProficiencies,
		&characterData.Class.SpellcastingAbility,
		&characterData.Background.BackgroundID,
		&characterData.Background.Name,
		&characterData.Background.Languages,
		&characterData.Background.PersonalityTraits,
		&characterData.Background.Ideals,
		&characterData.Background.Bond,
		&characterData.Background.Flaws,
		&characterData.Background.Trait,
		&characterData.Background.ToolProficiencies,
		&characterData.Name,
		&characterData.Story,
		&characterData.Alignment,
		&characterData.Age,
		&characterData.Hair,
		&characterData.Eyes,
		&characterData.Skin,
		&characterData.Height,
		&characterData.Weight,
		&characterData.ImgUrl,
		&characterData.Str,
		&characterData.Dex,
		&characterData.Int,
		&characterData.Con,
		&characterData.Wiz,
		&characterData.Cha,
		&characterData.Hitpoints,
		&characterData.HitDice,
		&characterData.Speed,
		&characterData.Armor_Class,
		&characterData.Level,
		&characterData.Exp,
	)
	return err
}
