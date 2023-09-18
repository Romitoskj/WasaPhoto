/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	// "wasaphoto/service/types"
	"database/sql"
	"errors"
	"fmt"
	"wasaphoto/service/types"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// session
	Login(username string) (int64, error)
	GetId(username string) (int64, error)

	// profile interactions
	ChangeUsername(id int64, newName string) error
	UsernameExists(username string) (bool, error)
	UserExists(user int64) (bool, error)
	Search(search string, id int64) ([]types.User, error)
	GetUserProfile(id int64) (types.Profile, error)

	// user relationships
	FollowUser(user int64, follower int64) error
	UnfollowUser(user int64, follower int64) error
	GetFollowers(id int64, auth int64) ([]types.User, error)
	GetFollowing(id int64, auth int64) ([]types.User, error)
	BanUser(user int64, bannedUser int64) error
	UnbanUser(user int64, bannedUser int64) error
	UserBanned(user int64, auth int64) (bool, error)
	UserIsFollowed(user int64, auth int64) (bool, error)

	// photos
	UploadPhoto(img []byte, authorId int64) (int64, error)
	GetPhoto(id int64, auth int64) (types.Photo, error)
	GetImage(id int64) ([]byte, error)
	PhotoExists(user int64) (bool, error)
	PhotoAuthor(id int64) (int64, error)
	DeletePhoto(id int64) error
	GetUserPhotos(user int64, auth int64) ([]types.Photo, error)

	// photo interactions
	LikeExists(liker int64, photo int64) (bool, error)

	// stream
	GetStream(author string) ([]types.Photo, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='user';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `
			CREATE TABLE "user" (
				"id"	INTEGER NOT NULL,
				"username"	TEXT NOT NULL UNIQUE,
				PRIMARY KEY("id" AUTOINCREMENT)
			);

			CREATE TABLE "follow" (
				"follower"	INTEGER NOT NULL,
				"followed"	INTEGER NOT NULL,
				FOREIGN KEY("followed") REFERENCES "user"("id") ON DELETE CASCADE,
				FOREIGN KEY("follower") REFERENCES "user"("id") ON DELETE CASCADE,
				PRIMARY KEY("follower","followed")
			);

			CREATE TABLE "ban" (
				"banner"	INTEGER NOT NULL,
				"banned"	INTEGER NOT NULL,
				FOREIGN KEY("banner") REFERENCES "user"("id") ON DELETE CASCADE,
				FOREIGN KEY("banned") REFERENCES "user"("id") ON DELETE CASCADE,
				PRIMARY KEY("banner","banned")
			);

			CREATE TABLE "photo" (
				"id"	INTEGER NOT NULL,
				"created_at"	DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				"image"	BLOB NOT NULL,
				"author"	INTEGER NOT NULL,
				PRIMARY KEY("id" AUTOINCREMENT),
				FOREIGN KEY("author") REFERENCES "user"("id") ON DELETE CASCADE
			);

			CREATE TABLE "like" (
				"liker"	INTEGER NOT NULL,
				"photo"	INTEGER NOT NULL,
				FOREIGN KEY("liker") REFERENCES "user"("id") ON DELETE CASCADE,
				FOREIGN KEY("photo") REFERENCES "photo"("id") ON DELETE CASCADE,
				PRIMARY KEY("liker","photo")
			);

			CREATE TABLE "comment" (
				"id"	INTEGER NOT NULL,
				"created_at"	DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				"content"	TEXT NOT NULL,
				"author"	INTEGER NOT NULL,
				"photo"	INTEGER NOT NULL,
				FOREIGN KEY("author") REFERENCES "user"("id") ON DELETE CASCADE,
				FOREIGN KEY("photo") REFERENCES "photo"("id") ON DELETE CASCADE,
				PRIMARY KEY("id" AUTOINCREMENT)
			);
		`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	_, err = db.Exec(`PRAGMA foreign_keys = ON;`) // enable foreign keys
	if err != nil {
		return nil, fmt.Errorf("error enabling foreign keys: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
