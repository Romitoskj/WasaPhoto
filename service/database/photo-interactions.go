package database

import "wasaphoto/service/types"

// LikeExists check if a user likes a picture
func (db *appdbimpl) LikeExists(liker int64, photo int64) (bool, error) {
	var exists int
	err := db.c.QueryRow("SELECT EXISTS (SELECT * FROM like WHERE liker=? AND photo=?)",
		liker, photo).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

// LikePhoto adds a like to a specified photo
func (db *appdbimpl) LikePhoto(user int64, photo int64) error {
	_, err := db.c.Exec(`INSERT OR IGNORE INTO like (liker, photo) VALUES (?, ?)`, user, photo)
	return err
}

// UnlikePhoto removes a like to a specified photo
func (db *appdbimpl) UnlikePhoto(user int64, photo int64) error {
	_, err := db.c.Exec(`DELETE FROM like WHERE liker=? AND photo=?`, user, photo)
	return err
}

// GetLikers returns the list of followers that like the specified photo
func (db *appdbimpl) GetLikers(photo int64, auth int64) ([]types.User, error) {
	var users []types.User
	users = []types.User{}

	// Get all the users that like the specified photo that have not banned the authenticated user
	rows, err := db.c.Query(
		`SELECT id, username
				FROM user
				JOIN like ON user.id = like.liker
				WHERE like.photo=?
				AND id NOT IN (
				    SELECT banner
				    FROM ban
				    WHERE banned = ?
				)`,
		photo,
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

// CommentPhoto insert a photo comment in db and return its id
func (db *appdbimpl) CommentPhoto(photo int64, content string, author int64) (int64, error) {
	row, err := db.c.Exec(`
		INSERT INTO comment (photo, author, content) VALUES (?, ?, ?)
	`, photo, author, content)
	if err != nil {
		return 0, err
	}
	var id int64
	id, err = row.LastInsertId()
	return id, err
}

// UncommentPhoto removes a comment to a specified photo
func (db *appdbimpl) UncommentPhoto(comment int64) error {
	_, err := db.c.Exec(`DELETE FROM comment WHERE id=?`, comment)
	return err
}

// GetComments returns the list of comments of the specified photo
func (db *appdbimpl) GetComments(photo int64, auth int64) ([]types.Comment, error) {
	var comments []types.Comment
	comments = []types.Comment{}

	// Get all the comments of the specified photo of the users that have not banned the authenticated user
	rows, err := db.c.Query(
		`SELECT c.id, c.author, u.username, c.created_at, c.content
				FROM comment c
				JOIN user u ON c.author = u.id
				WHERE c.photo=?
				AND u.id NOT IN (
				    SELECT banner
				    FROM ban
				    WHERE banned = ?
				)
				ORDER BY c.created_at`,
		photo,
		auth,
	)
	if err != nil {
		return comments, err
	}

	// append each query row to the users list
	for rows.Next() {
		var comment types.Comment
		err = rows.Scan(
			&comment.Identifier,
			&comment.Author.Id,
			&comment.Author.Name,
			&comment.CreatedAt,
			&comment.Content,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	// check if the iteration ended prematurely
	if rows.Err() != nil {
		return comments, rows.Err()
	}

	return comments, nil
}

// GetComment get comment from db given a its id
func (db *appdbimpl) GetComment(id int64) (types.Comment, error) {
	var comment types.Comment
	err := db.c.QueryRow(`
		SELECT c.id, c.created_at, c.author, u.username, c.content
		FROM comment c
		JOIN user u ON c.author = u.id
		WHERE c.id = ?
	`, id).Scan(
		&comment.Identifier,
		&comment.CreatedAt,
		&comment.Author.Id,
		&comment.Author.Name,
		&comment.Content,
	)
	if err != nil {
		return comment, err
	}

	return comment, err
}

// CommentAuthor returns the author id of a comment given its id
func (db *appdbimpl) CommentAuthor(id int64) (int64, error) {
	var author int64
	err := db.c.QueryRow(`
		SELECT author
		FROM comment
		WHERE id = ?
	`, id).Scan(&author)
	return author, err
}

// CommentBelongsToPhoto returns true if the comment belongs to the given photo
func (db *appdbimpl) CommentBelongsToPhoto(photo int64, comment int64) (bool, error) {
	var exists int
	err := db.c.QueryRow("SELECT EXISTS (SELECT * FROM comment WHERE id=? AND photo=?)",
		comment, photo).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}
