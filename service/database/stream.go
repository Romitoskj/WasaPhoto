package database

import "wasaphoto/service/types"

// GetStream returns the list of photos of the specified user
func (db *appdbimpl) GetStream(author string) ([]types.Photo, error) {
	var photos []types.Photo

	// Get all the users photos
	rows, err := db.c.Query(
		`SELECT p.id, p.created_at, p.author, COUNT(l.liker) likes_n, COUNT(c.id) comments_n
		FROM photo p
			LEFT JOIN like l ON p.id = l.photo
			LEFT JOIN comment c on p.id = c.photo
		WHERE p.author = ?
		GROUP BY p.id, p.created_at, p.author
		ORDER BY p.created_at DESC
		`, author,
	)
	if err != nil {
		return nil, err
	}

	// append each query row to the photos list
	for rows.Next() {
		var photo types.Photo
		err = rows.Scan(&photo.Identifier, &photo.CreatedAt, &photo.Author, &photo.LikesN, &photo.CommentsN)
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
