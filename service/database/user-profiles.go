package database

import "wasaphoto/service/types"

// ChangeUsername update the username of the provided user
func (db *appdbimpl) ChangeUsername(id int64, newName string) error {
	_, err := db.c.Exec("UPDATE user SET username=? WHERE id=?", newName, id)
	return err
}

// UsernameExists return true if a username exists
func (db *appdbimpl) UsernameExists(username string) (bool, error) {
	var exists int
	err := db.c.QueryRow("SELECT EXISTS (SELECT * FROM user WHERE username=?)", username).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

// UserExists return true if a user exists
func (db *appdbimpl) UserExists(id int64) (bool, error) {
	var exists int
	err := db.c.QueryRow("SELECT EXISTS (SELECT * FROM user WHERE id=?)", id).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

// Search return a list of user that match the search string
func (db *appdbimpl) Search(search string, id int64) ([]types.User, error) {
	var users []types.User

	// Get all the users that starts with the search string that have not banned the authenticated user
	rows, err := db.c.Query(
		`SELECT id, username
				FROM user
				WHERE username LIKE ?
				AND id NOT IN (
				    SELECT banner
				    FROM ban
				    WHERE banned = ?
				)`,
		search+"%",
		id,
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

// GetUserProfile return the user profile of the provided user id
func (db *appdbimpl) GetUserProfile(id int64) (types.Profile, error) {
	var profile types.Profile
	err := db.c.QueryRow(`
		SELECT u.username,
		   (
			   SELECT count(*)
			   FROM follow f
			   WHERE f.followed = u.id
		   ) count_followers,
		   (
			   SELECT COUNT(*)
			   FROM follow f
			   WHERE f.follower = u.id
		   ) count_followings,
		   (
			   SELECT count(*)
			   FROM photo p
			   WHERE u.id = p.author
		   ) count_posts
		FROM user u
		WHERE u.id == ?
	`, id).Scan(&profile.Name, &profile.FollowersN, &profile.FollowingN, &profile.PhotoN)
	return profile, err
}
