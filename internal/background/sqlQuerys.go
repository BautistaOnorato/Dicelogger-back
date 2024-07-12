package background

var (
	QueryCreateBackground = `
		INSERT INTO background (name, languages, personality_traits, ideals, bond, flaws, trait, tool_proficiencies)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`

	QueryGetAllBackgrounds = `
		SELECT * FROM background;
	`

	QueryGetBackgroundById = `
		SELECT * FROM background WHERE background_id = ?;
	`

	QueryUpdateBackground = `
		UPDATE background
		SET name = ?, languages = ?, personality_traits = ?, ideals = ?, bond = ?, flaws = ?, trait = ?, tool_proficiencies = ?
		WHERE background_id = ?;
	`

	QueryDeleteBackground = `
		DELETE FROM background WHERE background_id = ?;
	`
)
