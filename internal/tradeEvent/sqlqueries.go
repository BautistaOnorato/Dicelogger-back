package tradeevent

var (
	QueryInsert = "INSERT INTO trade_event (session_id, sender, receiver, description, timestamp) VALUES (?, ?, ?, ?, ?)"
	QueryGetBySessionId = "SELECT trade_event_id, session_id, sender, receiver, description, timestamp FROM trade_event WHERE session_id = ?"
	QueryGetBySender = "SELECT trade_event_id, session_id, sender, receiver, description, timestamp FROM trade_event WHERE sender = ?"
	QueryGetByReceiver = "SELECT trade_event_id, session_id, sender, receiver, description, timestamp FROM trade_event WHERE receiver = ?"
	QueryGetBySenderOrReciever = "SELECT trade_event_id, session_id, sender, receiver, description, timestamp FROM trade_event WHERE sender = ? OR receiver = ?"
	QueryDelete = "DELETE FROM trade_event WHERE trade_event_id = ?"
)