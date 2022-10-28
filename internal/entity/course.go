package entity

type Course struct {
	CourseID          int64  `json:"course_id"`
	CourseName        string `json:"course_name"`
	CourseDescription string `json:"course_description"`
}
