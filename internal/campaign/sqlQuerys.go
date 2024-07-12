package campaign

var (
	QueryCreateCampaign = `
		INSERT INTO campaign (dungeon_master, name, description, image, notes, status, images)
		VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	QueryGetAll = `
		SELECT * FROM campaign;
	`

	QueryGetById = `
		SELECT * FROM campaign WHERE campaign_id = ?;
	`

	QueryGetByUserId = `
		SELECT c.*
		FROM campaign c
		JOIN user_campaign uc ON c.campaign_id = uc.campaign_id
		WHERE uc.user_id = ?;
	
	`

	QueryGetUserData = `
		SELECT u.uid, u.name, u.email, u.display_name, u.image from user u 
		LEFT JOIN user_campaign uc ON u.uid  = uc.user_id 
		WHERE uc.campaign_id = ?;
	`

	QueryUpdate = `
		UPDATE campaign
		SET dungeon_master = ?, name = ?, description = ?, image = ? , notes = ?, status = ?, images = ?
		WHERE campaign_id = ?;
	`

	QueryDelete = `
		DELETE FROM campaign WHERE campaign_id = ?;
	`
)
