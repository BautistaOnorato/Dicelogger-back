package backgroundXproficiency

var(
	QueryInsert = "INSERT INTO background_proficiency (background_id, proficiency_id) values(?,?);"
	QueryDelete = "DELETE FROM background_proficiency WHERE background_id=? AND proficiency_id=?;"
)