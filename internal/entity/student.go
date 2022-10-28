package entity

type Student struct {
	ID             int64  `json:"id"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
	Name           string `json:"name"`
}
