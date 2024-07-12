package feature

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
)

type featureMySqlRepository struct {
	db *sql.DB
}

func NewFeatureRepository(db *sql.DB) FeatureRepository {
	return &featureMySqlRepository{db: db}
}

func (r *featureMySqlRepository) Create(feature dto.CreateFeatureDto) (domain.Feature, error) {
	var newFeature domain.Feature
	newFeature.Name = feature.Name
	newFeature.Description = feature.Description

	statement, err := r.db.Prepare(QueryCreateFeature)
	if err != nil {
		return domain.Feature{}, ErrPrepareStatement
	}
	
	defer statement.Close()
	result, err := statement.Exec(
		newFeature.Name,
		newFeature.Description,
	)
	if err != nil {
		return domain.Feature{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Feature{}, ErrGettingLastInsertId
	}
	newFeature.FeatureId = int(lastId)

	statement, err = r.db.Prepare(QueryCreateCharacterFeature)
	if err != nil {
		return domain.Feature{}, ErrPrepareStatement
	}
	defer statement.Close()
	result, err = statement.Exec(feature.CharacterId, lastId)
	if err != nil {
		return domain.Feature{}, err
	}
	
	return newFeature, nil
}

func (r *featureMySqlRepository) GetAll() ([]domain.Feature, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	features := []domain.Feature{}
	for rows.Next() {
		var feature domain.Feature
		err := rows.Scan(
			&feature.FeatureId,
			&feature.Name,
			&feature.Description,
		)
		if err != nil {
			return nil, err
		}
		features = append(features, feature)
	}
	
	return features, nil
}

func (r *featureMySqlRepository) GetAllByCharacterId(id int) ([]domain.Feature, error) {
	rows, err := r.db.Query(QueryGetAllByCharacterId, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	features := []domain.Feature{}
	for rows.Next() {
		var feature domain.Feature
		err := rows.Scan(
			&feature.FeatureId,
			&feature.Name,
			&feature.Description,
		)
		if err != nil {
			return nil, err
		}
		features = append(features, feature)
	}
	
	return features, nil
}

func (r *featureMySqlRepository) GetById(id int) (domain.Feature, error) {
	row := r.db.QueryRow(QueryGetById, id)
	var feature domain.Feature
	err := row.Scan(
		&feature.FeatureId,
		&feature.Name,
		&feature.Description,
	)
	if err != nil {
		return domain.Feature{}, err
	}
	
	return feature, nil
}

func (r *featureMySqlRepository) Update(feature domain.Feature, id int) (domain.Feature, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.Feature{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(feature.Name, feature.Description, id)
	if err != nil {
		return domain.Feature{}, err
	}

	feature.FeatureId = id

	return feature, nil
}

func (r *featureMySqlRepository) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()
	
	_, err = statement.Exec(id)
	return err
}