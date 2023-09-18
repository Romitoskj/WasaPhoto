package types

type Username struct {
	Username string `json:"name"`
}

type Id struct {
	Identifier int64 `json:"identifier"`
}

type User struct {
	Id   int64  `json:"identifier"`
	Name string `json:"name"`
}

type Follow struct {
	User     int64 `json:"user"`
	Follower int64 `json:"follower"`
}

type Ban struct {
	User       int64 `json:"user"`
	BannedUser int64 `json:"banned_user"`
}

type Photo struct {
	Identifier int64  `json:"identifier"`
	Author     User   `json:"author"`
	CreatedAt  string `json:"created_at"`
	LikesN     int64  `json:"likes_n"`
	CommentsN  int64  `json:"comments_n"`
	Liked      bool   `json:"liked"`
}

type Profile struct {
	Name       string  `json:"name"`
	PhotoN     int64   `json:"photos_n"`
	FollowersN int64   `json:"followers_n"`
	FollowingN int64   `json:"following_n"`
	Photos     []Photo `json:"photos"`
	Followed   bool    `json:"followed"`
}

type Error struct {
	Message string `json:"message"`
}
