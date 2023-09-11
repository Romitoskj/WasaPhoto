package database

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
