package weapon

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

type weaponMySqlRepository struct {
	db *sql.DB
}

func NewWeaponRepository(db *sql.DB) RepositoryWeapon {
	return &weaponMySqlRepository{db}
}

// GetAllGeneric implements RepositoryWeapon.
func (r *weaponMySqlRepository) GetAllGeneric() ([]domain.Weapon, error) {
	rows, err := r.db.Query(QueryGetGeneric)
	if err != nil {
		return []domain.Weapon{}, err
	}
	defer rows.Close()
	var weapons []domain.Weapon
	for rows.Next() {
		var weapon domain.Weapon
		err := rows.Scan(
			&weapon.Weapon_Id,
			&weapon.Weapon_Type,
			&weapon.Name,
			&weapon.Weight,
			&weapon.Price,
			&weapon.Category,
			&weapon.Reach,
			&weapon.Description,
			&weapon.Damage,
			&weapon.Versatile_Damage,
			&weapon.Ammunition,
			&weapon.Damage_Type,
			&weapon.Campaign_Id,
		)
		if err != nil {

			return []domain.Weapon{}, err
		}
		weapons = append(weapons, weapon)
	}
	if err := rows.Err(); err != nil {
		return []domain.Weapon{}, err
	}
	return weapons, nil
}

// Create implements RepositoryWeapon.
func (r *weaponMySqlRepository) Create(weapon domain.Weapon) (domain.Weapon, error) {
	statement, err := r.db.Prepare(QueryCreateWeapon)
	if err != nil {
		return domain.Weapon{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		weapon.Weapon_Type,
		weapon.Name,
		weapon.Weight,
		weapon.Price,
		weapon.Category,
		weapon.Reach,
		weapon.Description,
		weapon.Damage,
		weapon.Versatile_Damage,
		weapon.Ammunition,
		weapon.Damage_Type,
		weapon.Campaign_Id,
	)
	if err != nil {
		return domain.Weapon{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Weapon{}, ErrGettingLastInsertId
	}

	weapon.Weapon_Id = int(lastId)

	return weapon, nil
}

// Delete implements RepositoryWeapon.
func (r *weaponMySqlRepository) Delete(id int) error {
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

// GetAll implements RepositoryWeapon.
func (r *weaponMySqlRepository) GetAll() ([]domain.Weapon, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.Weapon{}, err
	}
	defer rows.Close()
	var weapons []domain.Weapon

	for rows.Next() {
		var weapon domain.Weapon
		err := rows.Scan(
			&weapon.Weapon_Id,
			&weapon.Weapon_Type,
			&weapon.Name,
			&weapon.Weight,
			&weapon.Price,
			&weapon.Category,
			&weapon.Reach,
			&weapon.Description,
			&weapon.Damage,
			&weapon.Versatile_Damage,
			&weapon.Ammunition,
			&weapon.Damage_Type,
			&weapon.Campaign_Id,
		)
		if err != nil {
			return []domain.Weapon{}, err
		}
		weapons = append(weapons, weapon)
	}
	if err := rows.Err(); err != nil {
		return []domain.Weapon{}, err
	}
	return weapons, nil
}

// GetByCampaignId implements RepositoryWeapon.
func (r *weaponMySqlRepository) GetByCampaignId(campaignId int) ([]domain.Weapon, error) {
	rows, err := r.db.Query(QueryGetByCampaignId, campaignId)
	if err != nil {
		return []domain.Weapon{}, err
	}
	defer rows.Close()
	var weapons []domain.Weapon

	for rows.Next() {
		var weapon domain.Weapon
		err := rows.Scan(
			&weapon.Weapon_Id,
			&weapon.Weapon_Type,
			&weapon.Name,
			&weapon.Weight,
			&weapon.Price,
			&weapon.Category,
			&weapon.Reach,
			&weapon.Description,
			&weapon.Damage,
			&weapon.Versatile_Damage,
			&weapon.Ammunition,
			&weapon.Damage_Type,
			&weapon.Campaign_Id,
		)
		if err != nil {
			return []domain.Weapon{}, err
		}
		weapons = append(weapons, weapon)
	}
	if err := rows.Err(); err != nil {
		return []domain.Weapon{}, err
	}

	if len(weapons) < 1 {
		return []domain.Weapon{}, ErrNotFound
	}
	return weapons, nil
}

// GetById implements RepositoryWeapon.
func (r *weaponMySqlRepository) GetById(id int) (domain.Weapon, error) {
	row := r.db.QueryRow(QueryGetById, id)
 
	var weapon domain.Weapon
	err := row.Scan(
		&weapon.Weapon_Id,
		&weapon.Weapon_Type,
		&weapon.Name,
		&weapon.Weight,
		&weapon.Price,
		&weapon.Category,
		&weapon.Reach,
		&weapon.Description,
		&weapon.Damage,
		&weapon.Versatile_Damage,
		&weapon.Ammunition,
		&weapon.Damage_Type,
		&weapon.Campaign_Id,
	)
	if err != nil {
		return domain.Weapon{}, ErrNotFound
	}

	return weapon, nil
}

// Update implements RepositoryWeapon.
func (r *weaponMySqlRepository) Update(weapon domain.Weapon) (domain.Weapon, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return domain.Weapon{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		weapon.Weapon_Type,
		weapon.Name,
		weapon.Weight,
		weapon.Price,
		weapon.Category,
		weapon.Reach,
		weapon.Description,
		weapon.Damage,
		weapon.Versatile_Damage,
		weapon.Ammunition,
		weapon.Damage_Type,
		weapon.Campaign_Id,
		weapon.Weapon_Id,
	)

	if err != nil {
		return domain.Weapon{}, err
	}

	return weapon, nil
}

