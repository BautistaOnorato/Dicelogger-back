package character_feature

var (
	QueryCreateCharacterFeature = `
		INSERT INTO character_feature (feature_id, character_id)
		VALUES (?, ?)
	`

	QueryGetAll = `
		SELECT * FROM character_feature;
	`

	QueryGetByFeatureId = `
		SELECT * FROM character_feature
		WHERE feature_id = ?;
	`

	QueryGetByCharacterId = `
		SELECT * FROM character_feature
		WHERE character_id = ?;
	`

	QueryDelete = `
		DELETE FROM character_feature WHERE feature_id = ? AND character_id = ?;
	`

	QueryDeleteByCharacterId = `
		DELETE FROM character_feature WHERE character_id = ?;
	`
)