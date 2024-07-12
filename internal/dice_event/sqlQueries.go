package dice_event

var (
	QueryInsert  = `INSERT INTO dice_event (stat,difficulty,dice_rolled,dice_result,event_protagonist,description,session_id,timestamp) values(?,?,?,?,?,?,?,?);`
	QueryGetAll  = `SELECT * from dice_event;`
	QueryGetById = `SELECT * from dice_event where dice_event_id = ?;`
	QueryGetBySessionId = `SELECT * from dice_event where session_id = ?;`
	QueryUpdate  = `UPDATE dice_event SET stat = ?, difficulty = ?, dice_rolled = ?, dice_result = ?, event_protagonist = ?, description = ?, session_id = ? WHERE dice_event_id = ?;`
	QueryDelete  = `DELETE FROM dice_event WHERE dice_event_id = ?;`
	QueryDeleteByProtagonistId  = `DELETE FROM dice_event WHERE event_protagonist = ?;`
)
