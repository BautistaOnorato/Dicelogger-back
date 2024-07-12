package backgroundxSkill

var (
	QueryCreateBackgroundXSkills = `
		INSERT INTO background_skills (background_id, skill_id)
		VALUES (?, ?);
	`

	QueryGetAllBackgroundXSkills = `
		SELECT * FROM background_skills;
	`

	QueryGetByIdBackgroundXSkills = `
		SELECT * FROM background_skills WHERE id = ?;
	`

	QueryGetByBackgroundId = `
		SELECT * FROM background_skills WHERE background_id = ?;
	`

	QueryUpdateBackgroundXSkills = `
		UPDATE background_skills
		SET background_id = ?, skill_id = ?
		WHERE background_id = ? AND skill_id = ?;
	`

	QueryDeleteBackgroundXSkills = `
		DELETE FROM background_skills WHERE id = ?;
	`

	QueryDeleteByBackgroundId = `
		DELETE FROM background_skills WHERE background_id = ?;
	`
)
