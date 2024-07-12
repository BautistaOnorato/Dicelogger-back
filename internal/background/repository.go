package background

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatementBackground    = errors.New("error preparing statement for background")
	ErrGettingLastInsertIdBackground = errors.New("error getting last insert id for background")
)

type backgroundMySqlRepository struct {
	db *sql.DB
}

func NewBackgroundRepository(db *sql.DB) BackgroundRepository {
	return &backgroundMySqlRepository{db: db}
}

func (r *backgroundMySqlRepository) Create(background domain.Background) (domain.Background, error) {
	statement, err := r.db.Prepare(QueryCreateBackground)
	if err != nil {
		fmt.Println(err)
		return domain.Background{}, ErrPrepareStatementBackground
	}

	defer statement.Close()
	result, err := statement.Exec(
		background.Name,
		background.Languages,
		background.PersonalityTraits,
		background.Ideals,
		background.Bond,
		background.Flaws,
		background.Trait,
		background.ToolProficiencies,
	)
	if err != nil {
		return domain.Background{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return domain.Background{}, ErrGettingLastInsertIdBackground
	}
	background.BackgroundID = int(lastID)

	return background, nil
}

func (r *backgroundMySqlRepository) GetAllBackgrounds() ([]domain.Background, error) {
	rows, err := r.db.Query(QueryGetAllBackgrounds)
	if err != nil {
		log.Println(1, err)
		return nil, err
	}
	defer rows.Close()
	
	var backgrounds []domain.Background
	for rows.Next() {
		var background domain.Background
		if err := rows.Scan(
			&background.BackgroundID,
			&background.Name,
			&background.Languages,
			&background.PersonalityTraits,
			&background.Ideals,
			&background.Bond,
			&background.Flaws,
			&background.Trait,
			&background.ToolProficiencies,
			); err != nil {
			log.Println(2, err)
			return nil, err
		}
		backgrounds = append(backgrounds, background)
	}

	return backgrounds, nil
}

func (r *backgroundMySqlRepository) GetBackgroundById(id int) (domain.Background, error) {
	var background domain.Background
	err := r.db.QueryRow(QueryGetBackgroundById, id).Scan(
		&background.BackgroundID,
		&background.Name,
		&background.Languages,
		&background.PersonalityTraits,
		&background.Ideals,
		&background.Bond,
		&background.Flaws,
		&background.Trait,
		&background.ToolProficiencies,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Background{}, errors.New("background not found")
		}
		return domain.Background{}, err
	}

	return background, nil
}

func (r *backgroundMySqlRepository) UpdateBackground(background domain.Background, id int) (domain.Background, error) {
	statement, err := r.db.Prepare(QueryUpdateBackground)
	if err != nil {
		return domain.Background{}, ErrPrepareStatementBackground
	}
	defer statement.Close()

	_, err = statement.Exec(
		background.Name,
		background.Languages,
		background.PersonalityTraits,
		background.Ideals,
		background.Bond,
		background.Flaws,
		background.Trait,
		background.ToolProficiencies,
		id,
	)
	if err != nil {
		return domain.Background{}, err
	}

	background.BackgroundID = id
	return background, nil
}

func (r *backgroundMySqlRepository) DeleteBackground(id int) error {
	statement, err := r.db.Prepare(QueryDeleteBackground)
	if err != nil {
		return ErrPrepareStatementBackground
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}
