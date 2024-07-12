package armor

var (
	QueryCreateArmor = `
		INSERT INTO armor (material, name, weight, price, category, protection_type, description, penalty, strength, armor_class, dex_bonus, campaign_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	QueryGetAllArmor = `
		SELECT * FROM armor;
	`

	QueryGetArmorByID = `
		SELECT * FROM armor WHERE armor_id = ?;
	`

	QueryUpdateArmor = `
		UPDATE armor
		SET material = ?, name = ?, weight = ?, price = ?, category = ?, protection_type = ?, description = ?, penalty = ?, strength = ?, armor_class = ?, dex_bonus = ?, campaign_id = ?
		WHERE armor_id = ?;
	`

	QueryDeleteArmor = `
		DELETE FROM armor WHERE armor_id = ?;
	`
)
