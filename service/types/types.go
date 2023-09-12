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

type Error struct {
	Message string `json:"message"`
}
