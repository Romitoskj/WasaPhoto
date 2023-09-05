package types

type Username struct {
	Username string `json:"name"`
}

type Id struct {
	Identifier int64 `json:"identifier"`
}

type Error struct {
	Message string `json:"message"`
}
