package proficiency

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
)

type repositorySqlProficiency struct {
	db *sql.DB
}

func NewProficiencyRepository(db *sql.DB) RepositoryProficiency {
	return &repositorySqlProficiency{db: db}
}

func (r *repositorySqlProficiency) Create(proficiencyDto dto.ProficiencyDto) (domain.Proficiency, error) {
	statement, err := r.db.Prepare(QueryInsertProficiency)
	if err != nil {
		return domain.Proficiency{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		proficiencyDto.Name,
		proficiencyDto.Type,
	)
	if err != nil {
		return domain.Proficiency{}, ErrGettingLastInsertId
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Proficiency{}, ErrGettingLastInsertId
	}

	createdProficiency := domain.Proficiency{
		ProficiencyId: int(lastId),
		Name:          proficiencyDto.Name,
		Type:          proficiencyDto.Type,
	}

	return createdProficiency, nil
}

func (r *repositorySqlProficiency) GetAll() ([]domain.Proficiency, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.Proficiency{}, err
	}
	proficiencyList := []domain.Proficiency{}
	for rows.Next() {
		var proficiency domain.Proficiency
		if err := rows.Scan(&proficiency.ProficiencyId, &proficiency.Name, &proficiency.Type); err != nil {
			return []domain.Proficiency{}, err
		}
		proficiencyList = append(proficiencyList, proficiency)
	}

	return proficiencyList, nil
}

func (r *repositorySqlProficiency) GetById(id int) (domain.Proficiency, error) {
	var proficiency domain.Proficiency
	if err := r.db.QueryRow(QueryGetById, id).Scan(&proficiency.ProficiencyId, &proficiency.Name, &proficiency.Type); err != nil {
		return domain.Proficiency{}, err
	}

	return proficiency, nil
}

func (r *repositorySqlProficiency) Update(proficiencyDto dto.ProficiencyDto, id int) (domain.Proficiency, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.Proficiency{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(proficiencyDto.Name, proficiencyDto.Type)
	if err != nil {
		return domain.Proficiency{}, err
	}
	updatedProficiency := domain.Proficiency{
		ProficiencyId: id,
		Name:          proficiencyDto.Name,
		Type:          proficiencyDto.Type,
	}

	return updatedProficiency, nil
}

func (r *repositorySqlProficiency) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}

	defer statement.Close()
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositorySqlProficiency) GetByCharacterDataId(characterId int) ([]domain.Proficiency, error) {
	rows, err := r.db.Query(QueryGetByCharacterDataId, characterId)
	if err != nil {
		return []domain.Proficiency{}, err
	}
	proficiencyList := []domain.Proficiency{}
	for rows.Next() {
		var proficiency domain.Proficiency
		if err := rows.Scan(&proficiency.ProficiencyId, &proficiency.Name, &proficiency.Type); err != nil {
			return []domain.Proficiency{}, err
		}
		proficiencyList = append(proficiencyList, proficiency)
	}
	
	

	return proficiencyList, nil
}
