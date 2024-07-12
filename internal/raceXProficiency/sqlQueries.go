package raceXproficiency

var (
	QueryInsert = "INSERT INTO race_proficiency (race_id, proficiency_id) values(?,?);"
	QueryDelete = "DELETE FROM race_proficiency WHERE race_id=? AND proficiency_id=?;"
)
