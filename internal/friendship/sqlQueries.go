package friendship

var (
	QueryCreate            = `INSERT INTO friend_list (user1_id, user2_id) VALUES(?,?);`
	QueryGetFriends        = `SELECT * FROM friend_list WHERE user1_id = ?;`
	QueryGetAllFriendships = `SELECT * FROM friend_list;`
	QuerySearchFollowers   = `SELECT * FROM friend_list WHERE user1_id = ?;`
	QueryCheckFriendship   = `SELECT count(*) FROM friend_list WHERE (user1_id = ? AND user2_id = ?) OR (user1_id = ? AND user2_id = ?);`
	QueryDelete            = `DELETE FROM friend_list WHERE user1_id = ? AND user2_id = ?;`
)
