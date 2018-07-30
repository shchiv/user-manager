package models

type User struct {
	ID   int    `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name string `sql:"type:VARCHAR(50)"`
	Role string
}
