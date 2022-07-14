package model

type User struct {
	Uid      int    `gorm:"primaryKey;AUTO_INCREMENT=1;not null"`
	Username string `gorm:"not null;unique;type:varchar(20)"`
	Password string `gorm:"type:varchar(100)"`
	Nickname string `gorm:"type:varchar(20)"`
}
