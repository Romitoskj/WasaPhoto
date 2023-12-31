package database

import "wasaphoto/service/types"

// UploadPhoto insert a photo in db and return its id
func (db *appdbimpl) UploadPhoto(img []byte, author int64) (int64, error) {
	row, err := db.c.Exec(`
		INSERT INTO photo (author, image) VALUES (?, ?)
	`, author, img)
	if err != nil {
		return 0, err
	}
	var id int64
	id, err = row.LastInsertId()
	return id, err
}

// GetPhoto get photo information from db given a photo id
func (db *appdbimpl) GetPhoto(id int64, auth int64) (types.Photo, error) {
	var photo types.Photo
	err := db.c.QueryRow(`
		SELECT p.id, p.created_at, p.author, u.username, COUNT(l.liker) likes_n, COUNT(c.id) comments_n
		FROM photo p
				 LEFT JOIN like l ON p.id = l.photo
				 LEFT JOIN comment c on p.id = c.photo
				 JOIN user u ON p.author = u.id
		WHERE p.id = ?
		GROUP BY p.id
	`, id).Scan(
		&photo.Identifier,
		&photo.CreatedAt,
		&photo.Author.Id,
		&photo.Author.Name,
		&photo.LikesN,
		&photo.CommentsN,
	)
	if err != nil {
		return photo, err
	}

	photo.Liked, err = db.LikeExists(auth, id)
	return photo, err
}

// GetImage get an image from db given a photo id
func (db *appdbimpl) GetImage(id int64) ([]byte, error) {
	var photo []byte
	err := db.c.QueryRow(`SELECT image FROM photo WHERE id = ?`, id).Scan(&photo)
	return photo, err
}

// PhotoExists returns true if a photo exists in db given its id
func (db *appdbimpl) PhotoExists(id int64) (bool, error) {
	var exists int
	err := db.c.QueryRow("SELECT EXISTS (SELECT * FROM photo WHERE id=?)", id).Scan(&exists)
	if err != nil {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

// PhotoAuthor returns the author id of a photo given its id
func (db *appdbimpl) PhotoAuthor(id int64) (int64, error) {
	var author int64
	err := db.c.QueryRow(`
		SELECT author
		FROM photo
		WHERE id = ?
	`, id).Scan(&author)
	return author, err
}

// DeletePhoto removes a photo given its id
func (db *appdbimpl) DeletePhoto(id int64) error {
	_, err := db.c.Exec(`DELETE FROM photo WHERE id=?`, id)
	return err
}

// GetUserPhotos returns the list of photos of the specified user
func (db *appdbimpl) GetUserPhotos(user int64, auth int64) ([]types.Photo, error) {
	var photos []types.Photo
	photos = []types.Photo{}

	// Get all the users photos
	rows, err := db.c.Query(
		`SELECT p.id, p.created_at, p.author, u.username, COUNT(DISTINCT l.liker) likes_n, COUNT(DISTINCT c.id) comments_n
		FROM photo p
			LEFT JOIN like l ON p.id = l.photo
			LEFT JOIN comment c on p.id = c.photo
			JOIN user u ON p.author = u.id
		WHERE p.author = ?
		GROUP BY p.id, p.created_at, p.author
		ORDER BY p.created_at DESC
		`, user,
	)
	if err != nil {
		return nil, err
	}

	// append each query row to the photos list
	for rows.Next() {
		var photo types.Photo
		err = rows.Scan(
			&photo.Identifier,
			&photo.CreatedAt,
			&photo.Author.Id,
			&photo.Author.Name,
			&photo.LikesN,
			&photo.CommentsN,
		)
		if err != nil {
			return nil, err
		}
		photo.Liked, err = db.LikeExists(auth, photo.Identifier)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	// check if the iteration ended prematurely
	if rows.Err() != nil {
		return photos, rows.Err()
	}

	return photos, nil
}
