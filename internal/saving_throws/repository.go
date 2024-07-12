package saving_throws

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

type repositorySqlSavingThrows struct {
	db *sql.DB
}

func NewRepositorySqlSavingThrows(db *sql.DB) SavingThrowsRepository {
	return &repositorySqlSavingThrows{db: db}
}

func (r *repositorySqlSavingThrows) Create(savingThrowDto dto.SavingThrowDto) (domain.SavingThrow, error) {
	statement, err := r.db.Prepare(QueryInsert)
	if err != nil {
		return domain.SavingThrow{}, ErrPrepareStatement
	}
	defer statement.Close()
	result, err := statement.Exec(
		savingThrowDto.ClassId,
		savingThrowDto.Str,
		savingThrowDto.Dex,
		savingThrowDto.Int,
		savingThrowDto.Con,
		savingThrowDto.Wiz,
		savingThrowDto.Cha,
	)

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.SavingThrow{}, ErrGettingLastInsertId
	}
	createdSavingThrow := domain.SavingThrow{
		SavingThrowId: int(lastId),
		ClassId:       savingThrowDto.ClassId,
		Str:           savingThrowDto.Str,
		Dex:           savingThrowDto.Dex,
		Int:           savingThrowDto.Int,
		Con:           savingThrowDto.Con,
		Wiz:           savingThrowDto.Wiz,
		Cha:           savingThrowDto.Cha,
	}

	return createdSavingThrow, nil
}

func (r *repositorySqlSavingThrows) GetAll() ([]domain.SavingThrow, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.SavingThrow{}, err
	}
	var savingThrowList []domain.SavingThrow
	for rows.Next() {
		var savingThrow domain.SavingThrow
		if err := rows.Scan(&savingThrow.SavingThrowId, &savingThrow.ClassId, &savingThrow.Str, &savingThrow.Dex, &savingThrow.Int, &savingThrow.Con, &savingThrow.Wiz, &savingThrow.Cha); err != nil {
			return []domain.SavingThrow{}, err
		}
		savingThrowList = append(savingThrowList, savingThrow)
	}
	return savingThrowList, nil
}

func (r *repositorySqlSavingThrows) GetById(id int) (domain.SavingThrow, error) {
	var savingThrow domain.SavingThrow
	if err := r.db.QueryRow(QueryGetById).Scan(&savingThrow.SavingThrowId, &savingThrow.ClassId, &savingThrow.Str, &savingThrow.Dex, &savingThrow.Int, &savingThrow.Con, &savingThrow.Wiz, &savingThrow.Cha); err != nil {
		return domain.SavingThrow{}, err
	}

	return savingThrow, nil
}

func (r *repositorySqlSavingThrows) Update(savingThrowDto dto.SavingThrowDto, id int) (domain.SavingThrow, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.SavingThrow{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		savingThrowDto.ClassId,
		savingThrowDto.Str,
		savingThrowDto.Dex,
		savingThrowDto.Int,
		savingThrowDto.Con,
		savingThrowDto.Wiz,
		savingThrowDto.Cha,
		id,
	)

	updatedSavingThrow := domain.SavingThrow{
		SavingThrowId: id,
		ClassId:       savingThrowDto.ClassId,
		Str:           savingThrowDto.Str,
		Dex:           savingThrowDto.Dex,
		Int:           savingThrowDto.Int,
		Con:           savingThrowDto.Con,
		Wiz:           savingThrowDto.Wiz,
		Cha:           savingThrowDto.Cha,
	}

	return updatedSavingThrow, nil
}

func (r *repositorySqlSavingThrows) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
