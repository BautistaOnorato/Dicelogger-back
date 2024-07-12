package backgroundxSkill

import (
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatementBackgroundSkills    = errors.New("error preparing statement for background skills")
	ErrGettingLastInsertIdBackgroundSkills = errors.New("error getting last insert id for background skills")
	ErrNotFoundBackgroundSkills            = errors.New("background skill not found")
)

type backgroundSkillsSqlRepository struct {
	db *sql.DB
}

func (r *backgroundSkillsSqlRepository) CreateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error) {
	statement, err := r.db.Prepare(QueryCreateBackgroundXSkills)
	if err != nil {
		return domain.BackgroundXSkills{}, ErrPrepareStatementBackgroundSkills
	}
	defer statement.Close()

	result, err := statement.Exec(
		data.Background_Id,
		data.Skill_Id,
	)

	if err != nil {
		return domain.BackgroundXSkills{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.BackgroundXSkills{}, ErrGettingLastInsertIdBackgroundSkills
	}

	data.BackgroundXSkills_Id = int(lastId)

	return data, nil
}

func (r *backgroundSkillsSqlRepository) DeleteBackgroundXSkills(id int) error {
	result, err := r.db.Exec(QueryDeleteBackgroundXSkills, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return ErrNotFoundBackgroundSkills
	}
	return nil
}

func (r *backgroundSkillsSqlRepository) DeleteByBackgroundId(id int) error {
	result, err := r.db.Exec(QueryDeleteByBackgroundId, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return ErrNotFoundBackgroundSkills
	}
	return nil
}

func (r *backgroundSkillsSqlRepository) GetAllBackgroundXSkills() ([]domain.BackgroundXSkills, error) {
	rows, err := r.db.Query(QueryGetAllBackgroundXSkills)

	if err != nil {
		return []domain.BackgroundXSkills{}, err
	}

	defer rows.Close()

	var backgroundXSkillsList []domain.BackgroundXSkills

	for rows.Next() {
		var data domain.BackgroundXSkills
		err := rows.Scan(
			&data.BackgroundXSkills_Id,
			&data.Background_Id,
			&data.Skill_Id,
		)
		if err != nil {
			return []domain.BackgroundXSkills{}, err
		}
		backgroundXSkillsList = append(backgroundXSkillsList, data)
	}
	if err := rows.Err(); err != nil {
		return []domain.BackgroundXSkills{}, err
	}
	return backgroundXSkillsList, nil
}

func (r *backgroundSkillsSqlRepository) GetByIdBackgroundXSkills(id int) (domain.BackgroundXSkills, error) {
	row := r.db.QueryRow(QueryGetByIdBackgroundXSkills, id)
	var data domain.BackgroundXSkills
	err := row.Scan(
		&data.BackgroundXSkills_Id,
		&data.Background_Id,
		&data.Skill_Id,
	)
	if err != nil {
		return domain.BackgroundXSkills{}, err
	}

	return data, nil
}

func (r *backgroundSkillsSqlRepository) GetByBackgroundId(id int) ([]domain.BackgroundXSkills, error) {
	rows, err := r.db.Query(QueryGetByBackgroundId, id)

	if err != nil {
		return []domain.BackgroundXSkills{}, err
	}

	defer rows.Close()

	var backgroundXSkillsList []domain.BackgroundXSkills

	for rows.Next() {
		var data domain.BackgroundXSkills
		err := rows.Scan(
			&data.BackgroundXSkills_Id,
			&data.Background_Id,
			&data.Skill_Id,
		)
		if err != nil {
			return []domain.BackgroundXSkills{}, err
		}
		backgroundXSkillsList = append(backgroundXSkillsList, data)
	}
	if err := rows.Err(); err != nil {
		return []domain.BackgroundXSkills{}, err
	}
	return backgroundXSkillsList, nil
}

func (r *backgroundSkillsSqlRepository) UpdateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error) {
	query := QueryUpdateBackgroundXSkills

	statement, err := r.db.Prepare(query)
	if err != nil {
		return domain.BackgroundXSkills{}, ErrPrepareStatementBackgroundSkills
	}
	defer statement.Close()

	_, err = statement.Exec(
		data.Background_Id,
		data.Skill_Id,
		data.Background_Id,
		data.Skill_Id,
	)

	if err != nil {
		return domain.BackgroundXSkills{}, err
	}

	return data, nil
}
