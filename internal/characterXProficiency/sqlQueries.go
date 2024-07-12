package characterXproficiency

var (
	QueryInsert = "INSERT INTO character_proficiency (character_id, proficiency_id) values(?,?);"
	QueryDelete = "DELETE FROM character_proficiency WHERE character_proficiency_id=?;"
	QueryDeleteByCharacterId = "DELETE FROM character_proficiency WHERE character_id =?;"
)

