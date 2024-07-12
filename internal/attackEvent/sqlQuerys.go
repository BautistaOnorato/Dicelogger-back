package attackEvent

var (
	QueryCreateAttackEvent = `
		INSERT INTO attack_event (type, environment, session_id, event_protagonist_id, event_resolution, weapon, spell, dmg_type, description, timestamp)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	QueryGetAll = `
		SELECT * from attack_event ae;
	`

	QueryGetById = `
		SELECT ae.*, s.*, cd.character_id, cd.user_id, cd.campaign_id, cd.name AS character_name, r.name AS race_name, cl.name AS class_name, cd.level, cd.hitpoints FROM attack_event ae
		INNER JOIN session s ON ae.session_id = s.session_id
		INNER join character_data cd ON ae.event_protagonist_id = cd.character_id 
		LEFT JOIN race r ON cd.race_id = r.race_id
		LEFT JOIN class cl ON cd.class_id = cl.class_id
		WHERE ae.event_id = ?;
	`

	QueryGetBySessionId = `
		SELECT ae.*, s.*, cd.character_id, cd.user_id, cd.campaign_id, cd.name AS character_name, r.name AS race_name, cl.name AS class_name, cd.level, cd.hitpoints FROM attack_event ae
		INNER JOIN session s ON ae.session_id = s.session_id
		INNER join character_data cd ON ae.event_protagonist_id = cd.character_id 
		LEFT JOIN race r ON cd.race_id = r.race_id
		LEFT JOIN class cl ON cd.class_id = cl.class_id
		WHERE ae.session_id = ?
	`

	QueryGetByProtagonistId = `
		SELECT ae.*, s.*, cd.character_id, cd.user_id, cd.campaign_id, cd.name AS character_name, r.name AS race_name, cl.name AS class_name, cd.level, cd.hitpoints FROM attack_event ae
		INNER JOIN session s ON ae.session_id = s.session_id
		INNER join character_data cd ON ae.event_protagonist_id = cd.character_id 
		LEFT JOIN race r ON cd.race_id = r.race_id
		LEFT JOIN class cl ON cd.class_id = cl.class_id
		WHERE ae.event_protagonist_id = ?;
	`

	QueryGetByAffectedId = `
		SELECT ae.*, s.*, cd.character_id, cd.user_id, cd.campaign_id, cd.name AS character_name, r.name AS race_name, cl.name AS class_name, cd.level, cd.hitpoints
		FROM character_data cd
		LEFT JOIN race r ON cd.race_id = r.race_id
		LEFT JOIN class cl ON cd.class_id = cl.class_id
		LEFT JOIN character_attack_event cae ON cd.character_id = cae.character_id
		INNER JOIN attack_event ae ON cae.event_id = ae.event_id AND ae.event_protagonist_id = cd.character_id
		INNER JOIN session s ON ae.session_id = s.session_id
		WHERE cae.character_id = ?
	`

	QueryGetByProtagonistIdAndAffectedId = `
		SELECT ae.*, s.*, cd.character_id, cd.user_id, cd.campaign_id, cd.name AS character_name, r.name AS race_name, cl.name AS class_name, cd.level, cd.hitpoints
		FROM character_data cd
		LEFT JOIN race r ON cd.race_id = r.race_id
		LEFT JOIN class cl ON cd.class_id = cl.class_id
		LEFT JOIN character_attack_event cae ON cd.character_id = cae.character_id
		INNER JOIN attack_event ae ON cae.event_id = ae.event_id AND ae.event_protagonist_id = cd.character_id
		INNER JOIN session s ON ae.session_id = s.session_id
		WHERE cae.character_id = ? OR ae.event_protagonist_id = ?;
	`

	QueryGetCharacterDataByAttackEventId = `
		SELECT character_data.character_id, character_data.user_id, character_data.campaign_id, character_data.img_url,
		race.name,
		class.name,
		character_data.name,
		character_data.hitpoints,
		character_data.level
		FROM character_data left join race on character_data.race_id = race.race_id left join class on character_data.class_id = class.class_id INNER JOIN character_attack_event cae on character_data.character_id = cae.character_id WHERE cae.event_id = ?;
	`

	QueryUpdate = `
		UPDATE attack_event
		SET type = ?, environment = ?, session_id = ?, event_protagonist_id = ?, event_resolution = ?, dmg_type = ?, weapon = ?, spell = ?, description = ?, timestamp = ?
		WHERE event_id = ?;
	`

	QueryDelete = `
		DELETE FROM attack_event WHERE event_id = ?;
	`
)
