package character_feature

import (
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
)

type characterFeatureRepository struct {
	db *sql.DB
}

func NewCharacterFeatureRepository(db *sql.DB) CharacterFeatureRepository {
	return &characterFeatureRepository{db: db}
}

// DeleteByCharacterDataId implements CharacterFeatureRepository.
func (r *characterFeatureRepository) DeleteByCharacterDataId(idCharacter int) error {
	statement, err := r.db.Prepare(QueryDeleteByCharacterId)
	if err != nil {
		return ErrPrepareStatement
	}

	defer statement.Close()
	_, err = statement.Exec(idCharacter)
	if err != nil {
		return err
	}

	return nil
}


func (r *characterFeatureRepository) Create(characterFeature domain.CharacterFeature) (domain.CharacterFeature, error) {
	statement, err := r.db.Prepare(QueryCreateCharacterFeature)
	if err != nil {
		return domain.CharacterFeature{}, ErrPrepareStatement
	}

	defer statement.Close()
	result, err := statement.Exec(
		characterFeature.CharacterId,
		characterFeature.FeatureId,
	)
	if err != nil {
		return domain.CharacterFeature{}, err
	}

	result.LastInsertId()

	return characterFeature, nil
}

func (r *characterFeatureRepository) GetAll() ([]domain.CharacterFeature, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var characterFeatures []domain.CharacterFeature
	for rows.Next() {
		var characterFeature domain.CharacterFeature
		err := rows.Scan(
			&characterFeature.CharacterId,
			&characterFeature.FeatureId,
		)
		if err != nil {
			return nil, err
		}
		characterFeatures = append(characterFeatures, characterFeature)
	}

	return characterFeatures, nil
}

func (r *characterFeatureRepository) GetByFeatureId(id int) ([]domain.CharacterFeature, error) {
	rows, err := r.db.Query(QueryGetByFeatureId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var characterFeatures []domain.CharacterFeature
	for rows.Next() {
		var characterFeature domain.CharacterFeature
		err := rows.Scan(
			&characterFeature.CharacterId,
			&characterFeature.FeatureId,
		)
		if err != nil {
			return nil, err
		}
		characterFeatures = append(characterFeatures, characterFeature)
	}

	return characterFeatures, nil
}

func (r *characterFeatureRepository) GetByCharacterId(id int) ([]domain.CharacterFeature, error) {
	rows, err := r.db.Query(QueryGetByCharacterId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var characterFeatures []domain.CharacterFeature
	for rows.Next() {
		var characterFeature domain.CharacterFeature
		err := rows.Scan(
			&characterFeature.CharacterId,
			&characterFeature.FeatureId,
		)
		if err != nil {
			return nil, err
		}
		characterFeatures = append(characterFeatures, characterFeature)
	}

	return characterFeatures, nil
}

func (r *characterFeatureRepository) Delete(characterId int, featureId int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}

	defer statement.Close()
	_, err = statement.Exec(characterId, featureId)
	if err != nil {
		return err
	}

	return nil
}
