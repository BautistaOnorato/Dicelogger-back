package itemxcharacterdata

var (
	QueryCreateItemXCharacterData = `INSERT INTO character_item (character_id, item_id, quantity) values (?, ?, ?)`
    QueryGetAll = `SELECT character_item_id, character_id, item.item_id , name, weight, price, description, campaign_id, quantity FROM character_item LEFT JOIN item ON character_item.item_id = item.item_id;`
    QueryGetById = `SELECT character_item_id, character_id, item.item_id , name, weight, price, description, campaign_id, quantity FROM character_item LEFT JOIN item ON character_item.item_id = item.item_id WHERE character_item_id = ?;`
    QueryGetByCharacterDataId = `SELECT character_item_id, character_id, item.item_id , name, weight, price, description, campaign_id, quantity FROM character_item LEFT JOIN item ON character_item.item_id = item.item_id WHERE character_id = ?;`
    QueryGetByCharacterDataIdAndItemId = `SELECT * FROM character_item WHERE character_id = ? and item_id = ?;`
    QueryUpdate = `UPDATE character_item SET character_id = ? , item_id = ? , quantity = ? WHERE character_item_id = ?`
    QueryUpdateOwnership = `UPDATE character_item SET character_id = ? , quantity = ? WHERE character_item_id = ?`
    QueryDelete = `DELETE FROM character_item WHERE character_item_id = ?;`
    QueryDeleteByCharacterDataId = `DELETE FROM character_item where character_id = ?;`
)