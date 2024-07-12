package skill

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrNotFound            = errors.New("skill not found")
)

type skillMySqlRepository struct {
	db *sql.DB
}

// Create implements RepositorySkill.
func (r *skillMySqlRepository) Create(skill domain.Skill) (domain.Skill, error) {
	statement, err := r.db.Prepare(QueryCreateSkill)
	if err != nil {
		return domain.Skill{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		skill.Name,
		skill.Stat,
	)

	if err != nil {
		return domain.Skill{}, err
	}
	lastId, err := result.LastInsertId()

	if err != nil {
		return domain.Skill{}, ErrGettingLastInsertId
	}

	skill.SkillId = int(lastId)

	return skill, nil
}

// Delete implements RepositorySkill.
func (r *skillMySqlRepository) Delete(id int) error {
	result, err := r.db.Exec(QueryDelete, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}

// GetAll implements RepositorySkill.
func (r *skillMySqlRepository) GetAll() ([]domain.Skill, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.Skill{}, err
	}

	defer rows.Close()
	skills := []domain.Skill{}
	for rows.Next() {
		var skill domain.Skill
		err := rows.Scan(
			&skill.SkillId,
			&skill.Name,
			&skill.Stat,
		)
		if err != nil {
			return []domain.Skill{}, err
		}
		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		return []domain.Skill{}, err
	}

	return skills, nil

}

// GetByClassId implements RepositorySkill.
func (r *skillMySqlRepository) GetByClassId(classId int) ([]domain.Skill, error) {
	rows, err := r.db.Query(QueryGetByClassId, classId)
	if err != nil {
		return []domain.Skill{}, err
	}

	defer rows.Close()
	skills := []domain.Skill{}
	for rows.Next() {
		var skill domain.Skill
		err := rows.Scan(
			&skill.SkillId,
			&skill.Name,
			&skill.Stat,
		)
		if err != nil {
			return []domain.Skill{}, err
		}
		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		return []domain.Skill{}, err
	}

	return skills, nil
}

// GetById implements RepositorySkill.
func (r *skillMySqlRepository) GetById(id int) (domain.Skill, error) {
	row := r.db.QueryRow(QueryGetById, id)

	var skill domain.Skill
	err := row.Scan(
		&skill.SkillId,
		&skill.Name,
		&skill.Stat,
	)
	if err != nil {
		return domain.Skill{}, ErrNotFound
	}

	return skill, nil
}

// GetByUserId implements RepositorySkill.
func (r *skillMySqlRepository) GetByCharacterId(characterId int) ([]domain.Skill, error) {
	rows, err := r.db.Query(QueryGetByCharacterId, characterId)
	if err != nil {
		return []domain.Skill{}, err
	}

	defer rows.Close()
	skills := []domain.Skill{}
	for rows.Next() {
		var skill domain.Skill
		err := rows.Scan(
			&skill.SkillId,
			&skill.Name,
			&skill.Stat,
		)
		if err != nil {
			return []domain.Skill{}, err
		}
		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		return []domain.Skill{}, err
	}

	return skills, nil
}

// Update implements RepositorySkill.
func (r *skillMySqlRepository) Update(skill domain.Skill) (domain.Skill, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.Skill{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		skill.Name,
		skill.Stat,
		skill.SkillId,
	)

	if err != nil {
		return domain.Skill{}, err
	}

	return skill, nil
}

func NewSkillRepository(db *sql.DB) RepositorySkill {
	return &skillMySqlRepository{db}
}
