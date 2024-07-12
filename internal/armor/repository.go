package armor

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrArmorNotFound       = errors.New("armor not found")
)

type armorMySqlRepository struct {
	db *sql.DB
}

func NewArmorRepository(db *sql.DB) ArmorRepository {
	return &armorMySqlRepository{db: db}
}

func (r *armorMySqlRepository) Create(armor domain.Armor) (domain.Armor, error) {
	statement, err := r.db.Prepare(QueryCreateArmor)
	if err != nil {
		fmt.Println(err)
		return domain.Armor{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		armor.Material,
		armor.Name,
		armor.Weight,
		armor.Price,
		armor.Category,
		armor.ProtectionType,
		armor.Description,
		armor.Penalty,
		armor.Strength,
		armor.ArmorClass,
		armor.DexBonus,
		armor.CampaignId,
	)
	if err != nil {
		return domain.Armor{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return domain.Armor{}, ErrGettingLastInsertId
	}
	armor.ArmorId = int(lastID)

	return armor, nil
}

func (r *armorMySqlRepository) GetAllArmors() ([]domain.Armor, error) {
	rows, err := r.db.Query(QueryGetAllArmor)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var armors []domain.Armor
	for rows.Next() {
		var armor domain.Armor
		if err := rows.Scan(
			&armor.ArmorId,
			&armor.Material,
			&armor.Name,
			&armor.Weight,
			&armor.Price,
			&armor.Category,
			&armor.ProtectionType,
			&armor.Description,
			&armor.Penalty,
			&armor.Strength,
			&armor.ArmorClass,
			&armor.DexBonus,
			&armor.CampaignId,
		); err != nil {
			return nil, err
		}
		armors = append(armors, armor)
	}
	return armors, nil
}

func (r *armorMySqlRepository) GetArmorById(id int) (domain.Armor, error) {
	var armor domain.Armor
	err := r.db.QueryRow(QueryGetArmorByID, id).Scan(
		&armor.ArmorId,
		&armor.Material,
		&armor.Name,
		&armor.Weight,
		&armor.Price,
		&armor.Category,
		&armor.ProtectionType,
		&armor.Description,
		&armor.Penalty,
		&armor.Strength,
		&armor.ArmorClass,
		&armor.DexBonus,
		&armor.CampaignId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Armor{}, ErrArmorNotFound
		}
		return domain.Armor{}, err
	}
	return armor, nil
}

func (r *armorMySqlRepository) UpdateArmor(armor domain.Armor, id int) (domain.Armor, error) {
	statement, err := r.db.Prepare(QueryUpdateArmor)
	if err != nil {
		return domain.Armor{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		armor.Material,
		armor.Name,
		armor.Weight,
		armor.Price,
		armor.Category,
		armor.ProtectionType,
		armor.Description,
		armor.Penalty,
		armor.Strength,
		armor.ArmorClass,
		armor.DexBonus,
		armor.CampaignId,
		id,
	)
	if err != nil {
		return domain.Armor{}, err
	}

	armor.ArmorId = id
	return armor, nil
}

func (r *armorMySqlRepository) DeleteArmor(id int) error {
	statement, err := r.db.Prepare(QueryDeleteArmor)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}
