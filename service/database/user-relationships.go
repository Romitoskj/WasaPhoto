package database

import "wasaphoto/service/types"

// FollowUser adds a follow relationship between two specified users
func (db *appdbimpl) FollowUser(user int64, follower int64) error {
	_, err := db.c.Exec(`INSERT OR IGNORE INTO follow (follower, followed) VALUES (?, ?)`, follower, user)
	return err
}

// UnfollowUser removes a follow relationship between two specified users
func (db *appdbimpl) UnfollowUser(user int64, follower int64) error {
	_, err := db.c.Exec(`DELETE FROM follow WHERE follower=? AND followed=?`, follower, user)
	return err
}

// GetFollowers returns the list of users that follows the specified user
func (db *appdbimpl) GetFollowers(id int64, auth int64) ([]types.User, error) {
	var users []types.User
	users = []types.User{}

	// Get all the users that follows the specified user that have not banned the authenticated user
	rows, err := db.c.Query(
		`SELECT id, username
				FROM user
				JOIN follow ON user.id = follow.follower
				WHERE follow.followed=?
				AND id NOT IN (
				    SELECT banner
				    FROM ban
				    WHERE banned = ?
				)`,
		id,
		auth,
	)
	if err != nil {
		return users, err
	}

	// append each query row to the users list
	for rows.Next() {
		var user types.User
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// check if the iteration ended prematurely
	if rows.Err() != nil {
		return users, rows.Err()
	}

	return users, nil
}

// GetFollowing returns the list of users that the specified user follows
func (db *appdbimpl) GetFollowing(id int64, auth int64) ([]types.User, error) {
	var users []types.User
	users = []types.User{}

	// Get all the users that the specified user follows that have not banned the authenticated user
	rows, err := db.c.Query(
		`SELECT id, username
				FROM user
				JOIN follow ON user.id = follow.followed
				WHERE follow.follower=?
				AND id NOT IN (
				    SELECT banner
				    FROM ban
				    WHERE banned = ?
				)`,
		id,
		auth,
	)
	if err != nil {
		return users, err
	}

	// append each query row to the users list
	for rows.Next() {
		var user types.User
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// check if the iteration ended prematurely
	if rows.Err() != nil {
		return users, rows.Err()
	}

	return users, nil
}

// BanUser adds a follow relationship between two specified users
func (db *appdbimpl) BanUser(user int64, bannedUser int64) error {
	_, err := db.c.Exec(`INSERT OR IGNORE INTO ban (banner, banned) VALUES (?, ?)`, user, bannedUser)
	return err
}

// UnbanUser removes a follow relationship between two specified users
func (db *appdbimpl) UnbanUser(user int64, bannedUser int64) error {
	_, err := db.c.Exec(`DELETE FROM ban WHERE banner=? AND banned=?`, user, bannedUser)
	return err
}

// UserBanned returns true if the authenticated user is banned by another user
func (db *appdbimpl) UserBanned(user int64, auth int64) (bool, error) {
	var exists int
	err := db.c.QueryRow("SELECT EXISTS (SELECT * FROM ban WHERE banner=? AND banned=?)", user, auth).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

// UserIsFollowed check if the authenticated user follows another user
func (db *appdbimpl) UserIsFollowed(user int64, auth int64) (bool, error) {
	var exists int
	err := db.c.QueryRow(
		"SELECT EXISTS (SELECT * FROM follow WHERE follower=? AND followed=?)",
		auth, user).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}
