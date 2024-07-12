package saving_throws

var (
	QueryInsert  = "INSERT INTO saving_thows (class_id, str, dex, `int`, con, wiz, cha) VALUES (?,?,?,?,?,?,?);"
	QueryGetAll  = `SELECT * FROM saving_throws;`
	QueryGetById = `SELECT * FROM saving_throws WHERE saving_throw_id = ?;`
	QueryUpdate  = "UPDATE saving_throws SET str=?,dex=?,`int`=?,con=?,wiz=?,cha=? WHERE saving_throw_id = ?;"
	QueryDelete  = `DELETE FROM saving_throws WHERE saving_throw_id = ?;`
)
