package user_campaign

var (
	QueryCreateUserCampaign = `
		INSERT INTO user_campaign (campaign_id, user_id, character_id, is_owner)
		VALUES (?, ?, ?, ?);
	`

	QueryGetAll = `
		SELECT * FROM user_campaign;
	`

	QueryGetById = `
		SELECT * FROM user_campaign WHERE user_campaign = ?;
	`

	QueryGetByCampaignId = `
		SELECT * FROM user_campaign WHERE campaign_id = ?;
	`

	QueryGetByUserId = `
		SELECT * FROM user_campaign WHERE user_id = ?;
	`

	QueryDelete = `
		DELETE FROM user_campaign WHERE user_campaign = ?;
	`

	QueryDeleteByCampaignID = `
		DELETE FROM user_campaign WHERE campaign_id = ?;
		`

	QueryCreateMultipleUserCampaign = `
	INSERT INTO user_campaign (campaign_id, user_id, character_id, is_owner)
	VALUES
	`

	QueryUpdateUserCharacter = `
	UPDATE user_campaign SET character_id = ? WHERE user_id = ? and campaign_id = ?;
	`

	QueryUpdateCharacterCampaign = `
	UPDATE character_data SET campaign_id = ? WHERE character_id = ?;
	`

)