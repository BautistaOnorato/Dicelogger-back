package itemxcharacterdata

import (
	"context"
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrNotFound            = errors.New("item not found")
)

type itemXCharacterDataSqlRepository struct {
	db *sql.DB
}

// Create implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) Create(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error) {
	tempContext := context.Background()
	tx, err := r.db.BeginTx(tempContext, nil)
	if err != nil {
		return domain.ItemXCharacterData{}, err
	}
	defer tx.Rollback()
	row := tx.QueryRowContext(tempContext, QueryGetByCharacterDataIdAndItemId, itemXCharacterData.CharacterData_Id, itemXCharacterData.Item.Item_Id)
	var tempItemXCharacterData domain.ItemXCharacterData
	err = row.Scan(
		&tempItemXCharacterData.Character_Item_Id,
		&tempItemXCharacterData.CharacterData_Id,
		&tempItemXCharacterData.Item.Item_Id,
		&tempItemXCharacterData.Quantity,
	)
	if err != nil {
		statement, err := tx.PrepareContext(tempContext, QueryCreateItemXCharacterData)
		if err != nil {
			return domain.ItemXCharacterData{}, ErrPrepareStatement
		}
		result, err := statement.ExecContext(
			tempContext,
			itemXCharacterData.CharacterData_Id,
			itemXCharacterData.Item.Item_Id,
			itemXCharacterData.Quantity,
		)
		if err != nil {
			return domain.ItemXCharacterData{}, err
		}
		lastId, err := result.LastInsertId()
		if err != nil {
			return domain.ItemXCharacterData{}, ErrGettingLastInsertId
		}
		itemXCharacterData.Character_Item_Id = int(lastId)
		statement.Close()

	} else {
		statement, err := tx.PrepareContext(tempContext, QueryUpdate)
		if err != nil {
			return domain.ItemXCharacterData{}, ErrPrepareStatement
		}

		_, err = statement.ExecContext(
			tempContext,
			itemXCharacterData.CharacterData_Id,
			tempItemXCharacterData.Item.Item_Id,
			itemXCharacterData.Quantity+tempItemXCharacterData.Quantity,
			tempItemXCharacterData.Character_Item_Id,
		)
		itemXCharacterData.Character_Item_Id = tempItemXCharacterData.Character_Item_Id
		itemXCharacterData.Quantity = itemXCharacterData.Quantity + tempItemXCharacterData.Quantity
		if err != nil {
			return domain.ItemXCharacterData{}, err
		}
	}

	if err = tx.Commit(); err != nil {
		return domain.ItemXCharacterData{}, err
	}
	return itemXCharacterData, nil
}

// Delete implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) Delete(id int) error {
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

// DeleteByTableCharacterId implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) DeleteByCharacterDataId(id int) error {
	result, err := r.db.Exec(QueryDeleteByCharacterDataId, id)
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

// GetAll implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) GetAll() ([]domain.ItemXCharacterData, error) {
	rows, err := r.db.Query(QueryGetAll)

	if err != nil {
		return []domain.ItemXCharacterData{}, err
	}

	defer rows.Close()

	itemXCharacterDataList := []domain.ItemXCharacterData{}

	for rows.Next() {
		var itemXCharacterData domain.ItemXCharacterData
		err := rows.Scan(
			&itemXCharacterData.Character_Item_Id,
			&itemXCharacterData.CharacterData_Id,
			&itemXCharacterData.Item.Item_Id,
			&itemXCharacterData.Item.Name,
			&itemXCharacterData.Item.Weight,
			&itemXCharacterData.Item.Price,
			&itemXCharacterData.Item.Description,
			&itemXCharacterData.Item.Campaign_Id,
			&itemXCharacterData.Quantity,
		)
		if err != nil {
			return []domain.ItemXCharacterData{}, err
		}
		itemXCharacterDataList = append(itemXCharacterDataList, itemXCharacterData)
	}
	if err := rows.Err(); err != nil {
		return []domain.ItemXCharacterData{}, err
	}
	return itemXCharacterDataList, nil
}

// GetById implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) GetById(id int) (domain.ItemXCharacterData, error) {
	row := r.db.QueryRow(QueryGetById, id)
	var itemXCharacterData domain.ItemXCharacterData
	err := row.Scan(
		&itemXCharacterData.Character_Item_Id,
		&itemXCharacterData.CharacterData_Id,
		&itemXCharacterData.Item.Item_Id,
		&itemXCharacterData.Item.Name,
		&itemXCharacterData.Item.Weight,
		&itemXCharacterData.Item.Price,
		&itemXCharacterData.Item.Description,
		&itemXCharacterData.Item.Campaign_Id,
		&itemXCharacterData.Quantity,
	)
	if err != nil {
		return domain.ItemXCharacterData{}, ErrNotFound
	}

	return itemXCharacterData, nil
}

// GetByTableCharacterId implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) GetByCharacterDataId(id int) ([]domain.ItemXCharacterData, error) {
	rows, err := r.db.Query(QueryGetByCharacterDataId, id)

	if err != nil {
		return []domain.ItemXCharacterData{}, err
	}

	defer rows.Close()

	itemXCharacterDataList := []domain.ItemXCharacterData{}

	for rows.Next() {
		var itemXCharacterData domain.ItemXCharacterData
		err := rows.Scan(
			&itemXCharacterData.Character_Item_Id,
			&itemXCharacterData.CharacterData_Id,
			&itemXCharacterData.Item.Item_Id,
			&itemXCharacterData.Item.Name,
			&itemXCharacterData.Item.Weight,
			&itemXCharacterData.Item.Price,
			&itemXCharacterData.Item.Description,
			&itemXCharacterData.Item.Campaign_Id,
			&itemXCharacterData.Quantity,
		)
		if err != nil {
			return []domain.ItemXCharacterData{}, err
		}
		itemXCharacterDataList = append(itemXCharacterDataList, itemXCharacterData)
	}
	if err := rows.Err(); err != nil {
		return []domain.ItemXCharacterData{}, err
	}
	return itemXCharacterDataList, nil
}

// Update implements RepositoryItemXTableCharacter.
func (r *itemXCharacterDataSqlRepository) Update(itemXCharacterData domain.ItemXCharacterData) (domain.ItemXCharacterData, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.ItemXCharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		itemXCharacterData.CharacterData_Id,
		itemXCharacterData.Item.Item_Id,
		itemXCharacterData.Quantity,
		itemXCharacterData.Character_Item_Id,
	)

	if err != nil {
		return domain.ItemXCharacterData{}, err
	}

	return itemXCharacterData, nil
}

func (r *itemXCharacterDataSqlRepository) UpdateOwnership(itemXCharacterData domain.ItemXCharacterData) error {
	tempContext := context.Background()
	tx, err := r.db.BeginTx(tempContext, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	row := tx.QueryRowContext(tempContext, QueryGetByCharacterDataIdAndItemId, itemXCharacterData.CharacterData_Id, itemXCharacterData.Item.Item_Id)
	var tempItemXCharacterData domain.ItemXCharacterData
	err = row.Scan(
		&tempItemXCharacterData.Character_Item_Id,
		&tempItemXCharacterData.CharacterData_Id,
		&tempItemXCharacterData.Item.Item_Id,
		&tempItemXCharacterData.Quantity,
	)
	if err != nil {
		statement, err := tx.PrepareContext(tempContext, QueryUpdateOwnership)
		if err != nil {
			return ErrPrepareStatement
		}
		_, err = statement.ExecContext(
			tempContext,
			itemXCharacterData.CharacterData_Id,
			itemXCharacterData.Quantity,
			itemXCharacterData.Character_Item_Id,
		)
		if err != nil {
			return err
		}
		statement.Close()
	} else {
		statement, err := tx.PrepareContext(tempContext, QueryUpdateOwnership)
		if err != nil {
			return ErrPrepareStatement
		}
		_, err = statement.ExecContext(
			tempContext,
			itemXCharacterData.CharacterData_Id,
			itemXCharacterData.Quantity+tempItemXCharacterData.Quantity,
			tempItemXCharacterData.Character_Item_Id,
		)
		if err != nil {
			return err
		}

		result, err := tx.ExecContext(tempContext, QueryDelete, itemXCharacterData.Character_Item_Id)
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

		statement.Close()
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func NewItemXCharacterDataSqlRepository(db *sql.DB) RepositoryItemXCharacterData {
	return &itemXCharacterDataSqlRepository{db}
}
