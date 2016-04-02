package forms

import "time"

// User Data
type User struct {
	FullName  string    `json:"full_name" form:"full_name"`
	Username  string    `json:"username" form:"username"`
	Address   string    `json:"address" form:"address"`
	BirthDate string    `json:"birth_date" form:"birth_date"`
	Job       string    `json:"job"`
	Created   time.Time `json:"created_time" form:"created_time"`
	Deleted   bool      `json:"-"`
}
