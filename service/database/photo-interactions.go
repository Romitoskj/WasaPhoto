package database

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
