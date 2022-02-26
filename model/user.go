package model

// User basic information
type User struct {
	ID       int64  `xorm:"pk autoincr" json:"id"`
	Username string `xorm:"unique" json:"username"`
	Name     string `json:"name"`
}
