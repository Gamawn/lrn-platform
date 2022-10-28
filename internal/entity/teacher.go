package entity

type Teacher struct {
	TeacherID      int64  `json:"teacher_id"`
	HashedPassword string `json:"password"`
	Email          string `json:"email"`
	Username       string `json:"username"`
}
