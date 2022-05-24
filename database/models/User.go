package models

type User struct {
	Base

	Username string
	Password string
}

//func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//	u.ID =
//}
