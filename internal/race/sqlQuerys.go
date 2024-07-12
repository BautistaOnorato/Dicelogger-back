package race

var (
	QueryCreateRace = "INSERT INTO race (name, description, speed, str, dex, `int`, con, wiz, cha) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);	"

	QueryGetAllRaces = `
		SELECT * FROM race;
	`

	QueryGetRaceById = `
		SELECT * FROM race WHERE race_id = ?;
	`

	QueryUpdateRace = "UPDATE race SET name = ?, description = ?, speed = ?, str = ?, dex = ?, `int` = ?, con = ?, wiz = ?, cha = ? WHERE race_id = ?;"

	QueryDeleteRace = `DELETE FROM race WHERE race_id = ?;`
)
