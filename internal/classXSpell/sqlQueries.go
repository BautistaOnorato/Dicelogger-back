package classXspell

var (
	QueryInsert = `INSERT INTO class_spell (class_id, spell_id) values(?,?);`
	QueryDelete = `DELETE FROM class_spell WHERE class_id=? AND spell_id=?;`
)
