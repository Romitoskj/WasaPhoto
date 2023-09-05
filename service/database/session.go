package database

import (
	// "wasaphoto/service/types"
	"database/sql"
	"errors"
)

// return the user id given the username
func (db *appdbimpl) GetId(username string) (int64, error) {
	var id int64
	err := db.c.QueryRow("SELECT id FROM user WHERE username=?", username).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// create a user with the given username
func (db *appdbimpl) createUser(username string) (int64, error) {
	var id int64
	row, err := db.c.Exec("INSERT INTO user (username) VALUES (?) RETURNING id", username)

	if err != nil {
		return 0, err
	}

	id, err = row.LastInsertId()
	return id, err
}

// return the user id given the username and create the user if not exists
func (db *appdbimpl) Login(username string) (int64, error) {

	id, err := db.GetId(username)

	if errors.Is(err, sql.ErrNoRows) {
		// if there isn't a user with that username
		id, err = db.createUser(username)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}
