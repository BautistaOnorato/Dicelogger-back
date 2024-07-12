package charactertrade

var (
	QueryBulkInsert = "INSERT INTO character_trade (trade_event_id, weapon_id, item_id, armor_id, item_owner, item_receiver, quantity, item_name, item_type) VALUES "
	QueryGetByTradeEventId = "SELECT character_trade_id, trade_event_id, weapon_id, item_id, armor_id, item_owner, item_receiver, quantity, item_name, item_type FROM character_trade WHERE trade_event_id = ?"
	QueryDeleteByTradeEventId = "DELETE FROM character_trade WHERE trade_event_id = ?"
)