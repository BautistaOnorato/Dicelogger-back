package item

var (
	QueryCreateItem      = `INSERT INTO item (name, weight, price, description, campaign_id) VALUES (?, ?, ?, ?, ?);`
	QueryGetAll          = `SELECT * FROM item;`
	QueryGetAllGeneric   = `SELECT * FROM item WHERE campaign_id IS NULL;`
	QueryGetByCampaignId = `SELECT * FROM item WHERE campaign_id = ?;`
	QueryGetById         = `SELECT * FROM item WHERE item_id = ?;`
	QueryUpdate          = `UPDATE item SET name = ?, weight = ? , price = ? , description = ? , campaign_id = ? WHERE item_id = ? ;`
	QueryDelete          = `DELETE from item where item_id = ?;`
)
