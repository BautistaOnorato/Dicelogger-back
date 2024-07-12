package race

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatementRace    = errors.New("error preparing statement for race")
	ErrGettingLastInsertIdRace = errors.New("error getting last insert id for race")
)

type raceMySqlRepository struct {
	db *sql.DB
}

func NewRaceRepository(db *sql.DB) RaceRepository {
	return &raceMySqlRepository{db: db}
}

func (r *raceMySqlRepository) Create(race domain.Race) (domain.Race, error) {
	statement, err := r.db.Prepare(QueryCreateRace)
	if err != nil {
		fmt.Println(err)
		return domain.Race{}, ErrPrepareStatementRace
	}

	defer statement.Close()
	result, err := statement.Exec(
		race.Name,
		race.Description,
		race.Speed,
		race.Str,
		race.Dex,
		race.Int,
		race.Con,
		race.Wiz,
		race.Cha,
	)
	if err != nil {
		return domain.Race{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return domain.Race{}, ErrGettingLastInsertIdRace
	}
	race.RaceID = int(lastID)

	return race, nil
}

func (r *raceMySqlRepository) GetAllRaces() ([]domain.Race, error) {
	rows, err := r.db.Query(QueryGetAllRaces)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var races []domain.Race
	for rows.Next() {
		var race domain.Race
		if err := rows.Scan(
			&race.RaceID,
			&race.Name,
			&race.Description,
			&race.Speed,
			&race.Str,
			&race.Dex,
			&race.Int,
			&race.Con,
			&race.Wiz,
			&race.Cha,
		); err != nil {
			return nil, err
		}
		races = append(races, race)
	}

	return races, nil
}

func (r *raceMySqlRepository) GetRaceById(id int) (domain.Race, error) {
	var race domain.Race
	err := r.db.QueryRow(QueryGetRaceById, id).Scan(
		&race.RaceID,
		&race.Name,
		&race.Description,
		&race.Speed,
		&race.Str,
		&race.Dex,
		&race.Int,
		&race.Con,
		&race.Wiz,
		&race.Cha,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Race{}, errors.New("race not found")
		}
		return domain.Race{}, err
	}

	return race, nil
}

func (r *raceMySqlRepository) UpdateRace(race domain.Race, id int) (domain.Race, error) {
	statement, err := r.db.Prepare(QueryUpdateRace)
	if err != nil {
		return domain.Race{}, ErrPrepareStatementRace
	}
	defer statement.Close()

	_, err = statement.Exec(
		race.Name,
		race.Description,
		race.Speed,
		race.Str,
		race.Dex,
		race.Int,
		race.Con,
		race.Wiz,
		race.Cha,
		id,
	)
	if err != nil {
		return domain.Race{}, err
	}

	race.RaceID = id
	return race, nil
}

func (r *raceMySqlRepository) DeleteRace(id int) error {
	statement, err := r.db.Prepare(QueryDeleteRace)
	if err != nil {
		return ErrPrepareStatementRace
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}
