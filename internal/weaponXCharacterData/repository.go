package weaponxcharacterdata

import (
	"database/sql"
	"errors"
	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrNotFound            = errors.New("weapon not found")
)

type weaponXCharacterDataSqlRepository struct {
	db *sql.DB
}

// Create implements RepositoryItemXTableCharacter.
func (r *weaponXCharacterDataSqlRepository) Create(weaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error) {
	statement, err := r.db.Prepare(QueryCreateWeaponXCharacterData)
	if err != nil {
		return domain.WeaponXCharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		weaponXCharacterData.CharacterData_Id,
		weaponXCharacterData.Weapon.Weapon_Id,
		weaponXCharacterData.Equipped,
	)

	if err != nil {
		return domain.WeaponXCharacterData{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.WeaponXCharacterData{}, ErrGettingLastInsertId
	}

	weaponXCharacterData.Character_Weapon_Id = int(lastId)

	return weaponXCharacterData, nil
}

// Delete implements RepositoryItemXTableCharacter.
func (r *weaponXCharacterDataSqlRepository) Delete(id int) error {
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
func (r *weaponXCharacterDataSqlRepository) DeleteByCharacterDataId(id int) error {
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
func (r *weaponXCharacterDataSqlRepository) GetAll() ([]domain.WeaponXCharacterData, error) {
	rows, err := r.db.Query(QueryGetAll)

	if err != nil {
		return []domain.WeaponXCharacterData{}, err
	}

	defer rows.Close()

	weaponXCharacterDataList := []domain.WeaponXCharacterData{}


	for rows.Next() {
		var weaponXCharacterData domain.WeaponXCharacterData
		err := rows.Scan(
			&weaponXCharacterData.Character_Weapon_Id,
			&weaponXCharacterData.CharacterData_Id,
			&weaponXCharacterData.Weapon.Weapon_Id,
			&weaponXCharacterData.Weapon.Weapon_Type,
			&weaponXCharacterData.Weapon.Name,
			&weaponXCharacterData.Weapon.Weight,
			&weaponXCharacterData.Weapon.Price,
			&weaponXCharacterData.Weapon.Category,
			&weaponXCharacterData.Weapon.Reach,
			&weaponXCharacterData.Weapon.Description,
			&weaponXCharacterData.Weapon.Damage,
			&weaponXCharacterData.Weapon.Versatile_Damage,
			&weaponXCharacterData.Weapon.Ammunition,
			&weaponXCharacterData.Weapon.Damage_Type,
			&weaponXCharacterData.Weapon.Campaign_Id,
			&weaponXCharacterData.Equipped,
		)
		if err != nil {
			return []domain.WeaponXCharacterData{}, err
		}
		weaponXCharacterDataList = append(weaponXCharacterDataList, weaponXCharacterData)
	}

	if err := rows.Err(); err != nil {
		return []domain.WeaponXCharacterData{}, err
	}
	return weaponXCharacterDataList, nil
}

// GetById implements RepositoryItemXTableCharacter.
func (r *weaponXCharacterDataSqlRepository) GetById(id int) (domain.WeaponXCharacterData, error) {
	row := r.db.QueryRow(QueryGetById, id)
	var weaponXCharacterData domain.WeaponXCharacterData
	err := row.Scan(
		&weaponXCharacterData.Character_Weapon_Id,
		&weaponXCharacterData.CharacterData_Id,
		&weaponXCharacterData.Weapon.Weapon_Id,
		&weaponXCharacterData.Weapon.Weapon_Type,
		&weaponXCharacterData.Weapon.Name,
		&weaponXCharacterData.Weapon.Weight,
		&weaponXCharacterData.Weapon.Price,
		&weaponXCharacterData.Weapon.Category,
		&weaponXCharacterData.Weapon.Reach,
		&weaponXCharacterData.Weapon.Description,
		&weaponXCharacterData.Weapon.Damage,
		&weaponXCharacterData.Weapon.Versatile_Damage,
		&weaponXCharacterData.Weapon.Ammunition,
		&weaponXCharacterData.Weapon.Damage_Type,
		&weaponXCharacterData.Weapon.Campaign_Id,
		&weaponXCharacterData.Equipped,
	)
	if err != nil {
		return domain.WeaponXCharacterData{}, err
	}

	return weaponXCharacterData, nil
}

// GetByTableCharacterId implements RepositoryItemXTableCharacter.
func (r *weaponXCharacterDataSqlRepository) GetByCharacterDataId(id int) ([]domain.WeaponXCharacterData, error) {
	rows, err := r.db.Query(QueryGetByCharacterDataId, id)
	if err != nil {
		return []domain.WeaponXCharacterData{}, err
	}

	defer rows.Close()

	weaponXCharacterDataList := []domain.WeaponXCharacterData{}

	for rows.Next() {
		var weaponXCharacterData domain.WeaponXCharacterData
		err := rows.Scan(
			&weaponXCharacterData.Character_Weapon_Id,
			&weaponXCharacterData.CharacterData_Id,
			&weaponXCharacterData.Weapon.Weapon_Id,
			&weaponXCharacterData.Weapon.Weapon_Type,
			&weaponXCharacterData.Weapon.Name,
			&weaponXCharacterData.Weapon.Weight,
			&weaponXCharacterData.Weapon.Price,
			&weaponXCharacterData.Weapon.Category,
			&weaponXCharacterData.Weapon.Reach,
			&weaponXCharacterData.Weapon.Description,
			&weaponXCharacterData.Weapon.Damage,
			&weaponXCharacterData.Weapon.Versatile_Damage,
			&weaponXCharacterData.Weapon.Ammunition,
			&weaponXCharacterData.Weapon.Damage_Type,
			&weaponXCharacterData.Weapon.Campaign_Id,
			&weaponXCharacterData.Equipped,
		)
		if err != nil {
			return []domain.WeaponXCharacterData{}, err
		}
		weaponXCharacterDataList = append(weaponXCharacterDataList, weaponXCharacterData)
	}
	if err := rows.Err(); err != nil {
		return []domain.WeaponXCharacterData{}, err
	}
	return weaponXCharacterDataList, nil
}

// Update implements RepositoryItemXTableCharacter.
func (r *weaponXCharacterDataSqlRepository) Update(weaponXCharacterData domain.WeaponXCharacterData) (domain.WeaponXCharacterData, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.WeaponXCharacterData{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		weaponXCharacterData.CharacterData_Id,
		weaponXCharacterData.Weapon.Weapon_Id,
		weaponXCharacterData.Equipped,
		weaponXCharacterData.Character_Weapon_Id,
	)

	if err != nil {
		return domain.WeaponXCharacterData{}, err
	}

	return weaponXCharacterData, nil
}

func (r *weaponXCharacterDataSqlRepository) UpdateOwnership(weaponXCharacterData domain.WeaponXCharacterData) (error) {
	statement, err := r.db.Prepare(QueryUpdateOwnership)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		weaponXCharacterData.CharacterData_Id,
		weaponXCharacterData.Equipped,
		weaponXCharacterData.Character_Weapon_Id,
	)

	if err != nil {
		return err
	}

	return nil
}

func NewWeaponXCharacterDataSqlRepository(db *sql.DB) RepositoryWeaponXCharacterData {
	return &weaponXCharacterDataSqlRepository{db}
}
