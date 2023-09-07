package database

/*
import (
	// "wasaphoto/service/types"
	"database/sql"
	"errors"
)
*/

// update the username of the provided user
func (db *appdbimpl) ChangeUsername(id int64, newName string) error {
	_, err := db.c.Exec("UPDATE user SET username=? WHERE id=?", newName, id)
	return err
}

// return true if a username exists
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

// return true if a user exists
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
